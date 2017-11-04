import React from 'react';
import { render } from 'react-dom';
import Socket from './Socket';
import { AppBar, Button, Divider, IconButton, List, ListItem, ListItemText, Paper, Toolbar, Typography } from 'material-ui';
import MenuIcon from 'material-ui-icons/Menu';
import AddIcon from 'material-ui-icons/Add';

var ws;

function Player(props) {
  let name = props.player ? props.player.fullName || '' : '';
  return (
    <ListItem dense>
      <ListItemText secondary={props.i + '.'}/>
      <ListItemText primary={name}/>
    </ListItem>
  );
}

function Players(props) {
  let players = props.players || [],
      total = Math.max(props.max, players.length),
      list = [];
  for (var i=0; i<total; i++) {
    list.push(<Player key={i} i={i+1} player={players[i]}/>);
  }
  return <List>{list}</List>;
}

function Table(props) {
  let { table } = props;
  return (
      <Paper elevation={4} style={{padding:4,margin:8}}>
        <List>
          <ListItem>
            <ListItemText primary={table.name}/>
          </ListItem>
        </List>
        <Divider/>
        <Players players={table.players} max={table.max}/>
      </Paper>
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
        <AppBar position="static">
          <Toolbar>
            <IconButton color="contrast" aria-label="Menu"
                style={{marginLeft:-12, marginRight:20}}>
              <MenuIcon/>
            </IconButton>
            <Typography type="title" color="inherit" style={{flex: 1}}>
              Tables
            </Typography>
            <Button color="contrast">Login</Button>
          </Toolbar>
        </AppBar>
        <div style={{display:'flex'}}>
        {this.state.tables.map((t) => <Table table={t} key={t.name}/>)}
        <Button fab color="primary" aria-label="add" style={{margin:75}}>
          <AddIcon/>
        </Button>
        </div>
      </div>
    );
  }
}

render(<Client/>, document.getElementById('root'));
