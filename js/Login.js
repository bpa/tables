import React from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField, Typography } from 'material-ui';
import ws from './Socket';

var player = {};

class Login extends React.Component {
  constructor() {
    super();
    this.open = this.open.bind(this);
    this.close = this.close.bind(this);
    this.login = this.login.bind(this);
    ws.register('login', this.on_login.bind(this));

    let state = {open:false, name:''};
    if (window.localStorage.player) {
      state.player = JSON.parse(window.localStorage.player);
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

  login() {
    console.log("logging in");
    ws.send({cmd: 'login', player: {fullName:this.state.name}});
  }

  on_change(k, e) {
    this.setState({[k]: e.target.value});
  }

  on_login(msg) {
    console.log(msg);
    window.localStorage.player = JSON.stringify(msg.player);
    this.setState({player: msg.player, open:false});
    player.data = msg.player;
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
        <Button color="contrast" onClick={this.open}>Login</Button>
        <Dialog open={this.state.open}>
          <DialogTitle>Login</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Welcome.  For the short term, we'll just need your name as other people know it.
            </DialogContentText>
            <form noValidate autoComplete="off">
              <TextField id="name" label="Name" margin="normal" style={{width:'100%'}}
                value={this.state.name}
                onChange={this.on_change.bind(this, 'name')}
              />
            </form>
          </DialogContent>
          <DialogActions>
            <Button color="primary" onClick={this.login}>Login</Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export { player, Login };
