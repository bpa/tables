import React, { useState } from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField, Typography } from '@material-ui/core';
import ws from './Socket';

export var player = {};

export function Login() {
  let [nameInput, setNameInput] = useState('');
  let [name, setName] = useState(() => {
    if (!window.localStorage.player) {
      return '';
    }
    const player = JSON.parse(window.localStorage.player);
    return player.fullName;
  });

  ws.register('login', (msg) => {
    window.localStorage.player = JSON.stringify(msg.player);
    setName(msg.player.fullName);
    setNameInput('');
  });

  ws.register('logout', () => {
    window.localStorage.removeItem('player');
    setNameInput(name);
    setName('');
  });

  function login(e) {
    e.preventDefault();
    ws.send({ cmd: 'login', method: 'trusted', username: nameInput });
  }

  return (
    <>
      <Typography color="inherit" type="title">
        {name}
      </Typography>
      <Dialog open={!name}>
        <DialogTitle>Login</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Enter your name as other people know it.
          </DialogContentText>
          <form noValidate autoComplete="off" onSubmit={login}>
            <TextField id="name" label="Name" margin="normal" style={{ width: '100%' }}
              onChange={e => setNameInput(e.target.value)}
              value={nameInput}
              autoFocus
            />
          </form>
        </DialogContent>
        <DialogActions>
          <Button type="button" color="primary" onClick={login}>Login</Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
