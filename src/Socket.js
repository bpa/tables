import { action, state } from './GlobalState';

let open = false;
let buffer = [];
var ws;

let path = window.location.toString()
  .replace(/https?/, 'ws').replace(/\?.*$/, '') + 'websocket';

function init() {
  ws = new WebSocket(path);
  ws.onmessage = handle_message;
  ws.onopen = onOpen;
  ws.onclose = onClose;
}

function onOpen() {
  open = true;
  if (window.localStorage.player) {
    const player = JSON.parse(window.localStorage.player);
    state.player = player;
    send({ cmd: 'login', method: 'trusted', username: player.fullName });
  }
  for (var m of buffer) {
    ws.send(JSON.stringify(m));
  }
  buffer = [];
};

function onClose() {
  open = false;
  setTimeout(init, 1000);
};

export default function send(msg) {
  console.log(msg);
  if (open) {
    ws.send(JSON.stringify(msg));
  }
  else {
    buffer.push(msg);
  }
}

function handle_message(m) {
  let msg = JSON.parse(m.data);
  console.info(msg);
  const callback = action[msg.cmd];
  callback && callback(msg);
}

init();