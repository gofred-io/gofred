package utils

import (
	"fmt"
	"runtime/debug"
)

func SafeGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				debug.PrintStack()
			}
		}()

		f()
	}()
}
