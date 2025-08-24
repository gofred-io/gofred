// Load and run the WASM file in the browser.
// Modified from https://tinygo.org/docs/guides/webassembly/

const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = 'main.wasm';
var wasm;

function createWebsocketClient() {
  console.log("creating websocket client");
  const socket = new WebSocket('ws://localhost:3001/ws');
  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    if (msg.cmd === "reload") {
      window.location.reload();
    }
  }
  socket.onopen = () => {
    console.log("connected to websocket server");
  }
  socket.onclose = () => {
    console.log("disconnected from websocket server");
  }
  socket.onerror = (event) => {
    console.log("error: ", event);
  }
}

// A function to run after WebAssembly is instantiated.
function postInstantiate(obj) {
  wasm = obj.instance;
  go.run(wasm);
  createWebsocketClient();
}

if ('instantiateStreaming' in WebAssembly) {
  WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(postInstantiate);
} else {
  fetch(WASM_URL).then(resp =>
    resp.arrayBuffer()
  ).then(bytes =>
    WebAssembly.instantiate(bytes, go.importObject).then(postInstantiate)
  )
}