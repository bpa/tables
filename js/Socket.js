export default class Socket {
  constructor(obj) {
    this.on_message = handle_message.bind(obj);
    this.init();
  }

  init() {
    let path = "ws" + location.toString().substr(4) + "/websocket";
    this.ws = new WebSocket(path);
    this.ws.onmessage = this.on_message;
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
} 

function handle_message(m) {
  let msg = JSON.parse(m.data);
  console.info(msg);
  if (msg.cmd) {
    var f, f_name = 'on_' + msg.cmd, o=this;
    while (o) {
      f = o[f_name];
      if (typeof f === 'function') {
        f.call(o, msg);
        return;
      }
      o = Object.getPrototypeOf(o);
    }
  }
}
