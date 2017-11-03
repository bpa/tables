import React from 'react';
import ReactDOM from 'react-dom';
import Socket from './Socket';

var ws;

function Player(props) {
  return <span>props.player.fullName</span>;
}

function Players(props) {
  return props.players ? (
    <div>
      {props.players.map((p) => <Player key={p.email} player={p}/>)}
    </div>
  ) : null;
}

function Table(props) {
  return (
    <div>
      <div>{props.table.name}</div>
      <Players players={props.table.players}/>
    </div>
  );
}

class Client extends React.Component {
  constructor() {
    super();
    ws = new Socket(this);
    this.state = {tables: []}

    window.onerror =  function(messageOrEvent, source, lineno, colno, error) {
      ws.send({cmd: 'error',
        message: error.message,
        stack: error.stack
      });
    }
  }

  on_tables(msg) {
    this.setState({tables: msg.tables});
  }

  render() {
    return (
      <div>
      {this.state.tables.map((t) => <Table table={t} key={t.name}/>)}
      </div>
      );
  }
}

const root = document.createElement("div");
document.body.appendChild(root);
ReactDOM.render(<Client/>, root);
