import React from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField, Typography } from '@material-ui/core';
import ws from './Socket';

var player = {};

class Login extends React.Component {
  constructor() {
    super();
    this.open = this.open.bind(this);
    this.close = this.close.bind(this);
    this.login = this.login.bind(this);
    ws.register('login', this.on_login.bind(this));
    ws.register('logout', this.on_logout.bind(this));

    let state = {open:false, name:''};
    if (window.localStorage.player) {
      state.player = JSON.parse(window.localStorage.player);
      state.name = state.player.fullName;
      player.data = state.player;
    }
    else {
      state.open = true;
    }
    this.state = state;
  }

  open() {
    this.setState({open: true, name: ''});
  }

  close() {
    this.setState({open: false});
  }

  login(e) {
    e.preventDefault();
    ws.send({cmd: 'login', method: 'trusted', username: this.state.name});
  }

  on_change(k, e) {
    this.setState({[k]: e.target.value});
  }

  on_login(msg) {
    window.localStorage.player = JSON.stringify(msg.player);
    this.setState({player: msg.player, open:false});
    player.data = msg.player;
  }

  on_logout(msg) {
    window.localStorage.removeItem('player');
    this.setState({player: null, open:false});
    delete player.data;
  }

  render() {
    if (this.state.player) {
      return (
        <Typography color="inherit" type="title">
          {this.state.player.fullName}
        </Typography>
      );
    }
    return (
      <div>
        <Button color="inherit" onClick={this.open}>Login</Button>
        <Dialog open={this.state.open}>
          <DialogTitle>Login</DialogTitle>
          <DialogContent>
            <DialogContentText>
              For the short term, we'll just need your name as other people know it.
            </DialogContentText>
            <form noValidate autoComplete="off" onSubmit={this.login}>
              <TextField id="name" label="Name" margin="normal" style={{width:'100%'}}
                value={this.state.name}
                onChange={this.on_change.bind(this, 'name')}
                autoFocus
              />
            </form>
          </DialogContent>
          <DialogActions>
            <Button type="button" color="primary" onClick={this.login}>Login</Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export { player, Login };
