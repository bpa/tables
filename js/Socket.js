class Socket {
  constructor(obj) {
    this.initialized = false;
    this.callbacks = {};
    this.callbackIds = {};
    this.id = 0;
  }

  register(cmd, func) {
    let id = this.id++;
    this.callbackIds[id] = cmd;
    if (this.callbacks[cmd] === undefined) {
      this.callbacks[cmd] = {}
    }
    this.callbacks[cmd][id] = func;

    if (!this.initialized) {
      this.init();
      this.initialized = true;
    }
  }

  deregister(id) {
    let cmd = this.callbackIds[id];
    delete this.callbackIds[id];
    delete this.callbacks[cmd][id];
    if (!Object.keys(this.callbacks[cmd]).length) {
      delete this.callbacks[cmd];
    }
  }

  init() {
    let path = location.toString()
      .replace(/https?/, 'ws').replace(/\?.*$/,'') + 'websocket';
    this.ws = new WebSocket(path);
    this.ws.onmessage = this.handle_message.bind(this);
    this.ws.onclose = () => setTimeout(this.init.bind(this), 1000);
  }

  close() {
    this.ws.onclose = undefined;
    this.ws.close();
  };

  send(msg) {
    console.log(msg);
    this.ws.send(JSON.stringify(msg));
  }

  handle_message(m) {
    let msg = JSON.parse(m.data);
    console.info(msg);
    if (msg.cmd) {
      let callbacks = this.callbacks[msg.cmd];
      if (!callbacks) {
        return;
      }
      for (const cb in callbacks) {
        callbacks[cb].call(null, msg);
      }
    }
  }
} 

var ws = new Socket();
export default ws;
