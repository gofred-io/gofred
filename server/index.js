// Load and run the WASM file in the browser.
// Modified from https://tinygo.org/docs/guides/webassembly/

const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = 'main.wasm';
var wasm;

function createWebsocketClient() {
  const livePort = window.env.LIVE_PORT;
  if (!livePort) {
    return;
  }

  const socket = new WebSocket(`ws://localhost:${livePort}/ws`);
  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    if (msg.cmd === "reload") {
      window.location.reload();
    }
  }
  socket.onopen = () => {
  }
  socket.onclose = () => {
    console.log("disconnected from websocket server, reconnecting...");
    createWebsocketClient();
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

  customElements.define('pushstate-anchor', HTMLPushStateAnchorElement, { extends: 'a' });
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

class HTMLPushStateAnchorElement extends HTMLAnchorElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.addEventListener('click', this.pushStateAnchorEventListener, false);
  }

  disconnectedCallback() {
    this.removeEventListener('click', this.pushStateAnchorEventListener, false);
  }

  pushStateAnchorEventListener(event) {
    // open in new tab or open context menu (workaround for Firefox)
    if (event.ctrlKey || event.metaKey || event.which === 2 || event.which === 3) {
      return;
    }

    var href = this.getAttribute('href');
    if (!href) {
      return;
    }

    // don't pushState if the URL is for a different host
    if (href.indexOf('http') === 0 && window.location.host !== new URL(href).host) {
      return;
    }

    if (href !== window.location.pathname) {
      // push state into the history stack if the current path is not the same as the new path
      window.history.pushState(JSON.parse(this.getAttribute('state')) || window.history.state, this.getAttribute('title'), href);
      window.document.getElementById('root').scrollTo(0, 0);
    }

    // dispatch a popstate event
    try {
      var popstateEvent = new PopStateEvent('popstate', {
        bubbles: false,
        cancelable: false,
        state: window.history.state
      });

      if ('dispatchEvent_' in window) {
        // FireFox with polyfill
        window.dispatchEvent_(popstateEvent);
      } else {
        // normal
        window.dispatchEvent(popstateEvent);
      }
    } catch (error) {
      // Internet Explorer
      var evt = document.createEvent('CustomEvent');
      evt.initCustomEvent('popstate', false, false, { state: window.history.state });
      window.dispatchEvent(evt);
    }

    // prevent the default link click
    event.preventDefault();
  }
}