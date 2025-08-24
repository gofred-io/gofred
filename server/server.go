package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var (
	chReload = make(chan struct{})
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type wsMsg struct {
	Command string `json:"cmd"`
}

func main() {
	go watch()
	go runWebsockerServer()

	http.ListenAndServe(":3000", http.FileServer(http.Dir("./server")))
}

func runWebsockerServer() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var (
			quit = make(chan struct{})
		)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error: %v", err)
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			return
		}

		defer conn.Close()
		defer close(quit)

		go func() {
			for {
				select {
				case <-quit:
					return
				case <-chReload:
					err := conn.WriteJSON(wsMsg{Command: "reload"})
					if err != nil {
						log.Printf("error writing to websocket: %v", err)
					}
				}
			}
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				//log.Printf("error: %v", err)
				return
			}
		}
	})

	http.ListenAndServe(":3001", nil)
}

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("error: %v", err)
	}
	defer watcher.Close()

	dir := "."
	err = watcher.Add(dir)
	if err != nil {
		log.Printf("error: %v", err)
	}

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && !strings.HasPrefix(path, ".") && strings.HasSuffix(path, ".go") {
			watcher.Add(path)
		}

		return nil
	})
	if err != nil {
		log.Printf("error: %v", err)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	hasChanged := atomic.Bool{}

	for {
		select {
		case _, ok := <-watcher.Events:
			if !ok {
				return
			}
			hasChanged.Store(true)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("error: %v", err)
		case <-ticker.C:
			if hasChanged.CompareAndSwap(true, false) {
				log.Printf("detected changes, rebuilding...")
				build()
				chReload <- struct{}{}
			}
		}
	}
}

func build() {
	process := exec.Command("go", "build", "-o", "server/main.wasm", ".")
	process.Env = append(os.Environ(), "GOARCH=wasm", "GOOS=js")
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	err := process.Run()
	if err != nil {
		log.Printf("error: %v", err)
	}
}
