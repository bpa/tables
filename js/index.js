import React from 'react';
import { render } from 'react-dom';
import Socket from './Socket';
import { AppBar, Button, Divider, IconButton, List, ListItem, ListItemIcon, ListItemText, Paper, Toolbar, Typography } from 'material-ui';
import CreateTable from './CreateTable';
import Table from './Table';
import MenuIcon from 'material-ui-icons/Menu';

var ws;

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
        <div style={{display:'flex',alignItems:'flex-start'}}>
          {this.state.tables.map((t) => <Table table={t} key={t.game.name}/>)}
          <CreateTable/>
        </div>
      </div>
    );
  }
}

render(<Client/>, document.getElementById('root'));
