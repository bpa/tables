import React from 'react';
import { render } from 'react-dom';
import ws from './Socket';
import { AppBar, Button, Divider, List, ListItem, ListItemIcon, ListItemText, Paper, Toolbar, Typography } from '@material-ui/core';
import CreateTable from './CreateTable';
import MainMenu from './MainMenu';
import Table from './Table';
import { Login } from './Login';

class Client extends React.Component {
  constructor() {
    super();
    this.state = {tables: []}
    ws.register('tables', this.on_tables.bind(this));

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
            <MainMenu/>
            <Typography type="title" color="inherit" style={{flex: 1}}>
              Tables
            </Typography>
            <Login/>
          </Toolbar>
        </AppBar>
        <div style={{display:'flex',alignItems:'flex-start'}}>
          {this.state.tables.map((t) => <Table table={t} key={t.id}/>)}
          <CreateTable/>
        </div>
      </div>
    );
  }
}

render(<Client/>, document.getElementById('root'));
