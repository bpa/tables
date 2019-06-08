class Socket {
  constructor(obj) {
    this.open = false;
    this.initialized = false;
    this.buffer = [];
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
    return id;
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
    let self = this;
    let path = window.location.toString()
      .replace(/https?/, 'ws').replace(/\?.*$/, '') + 'websocket';
    this.ws = new WebSocket(path);
    this.ws.onmessage = this.handle_message.bind(this);
    this.ws.onopen = function () {
      self.open = true;
      let buf = self.buffer;
      self.buffer = [];
      if (window.localStorage.player) {
        const player = JSON.parse(window.localStorage.player);
        self.send({ cmd: 'login', method: 'trusted', username: player.fullName });
      }
      for (var m of buf) {
        self.ws.send(JSON.stringify(m));
      }
    }
    this.ws.onclose = function () {
      self.open = false;
      setTimeout(self.init.bind(self), 1000);
    }
  }

  close() {
    this.ws.onclose = undefined;
    this.ws.close();
  };

  send(msg) {
    console.log(msg);
    if (this.open) {
      this.ws.send(JSON.stringify(msg));
    }
    else {
      this.buffer.push(msg);
    }
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
