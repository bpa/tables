import React, { useState, useContext } from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField, Typography } from '@material-ui/core';
import send from './Socket';
import GlobalContext from './GlobalState';
import { observer } from 'mobx-react-lite';

export default observer(() => {
  let context = useContext(GlobalContext);
  let [name, setName] = useState(context.player.fullName);

  function login(e) {
    e.preventDefault();
    send({ cmd: 'login', method: 'trusted', username: name });
  }
  
  return (
    <>
      <Typography color="inherit" type="title">
        {context.player.fullName}
      </Typography>
      <Dialog open={!context.player.fullName}>
        <DialogTitle>Login</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Enter your name as other people know it.
          </DialogContentText>
          <form noValidate autoComplete="off" onSubmit={login}>
            <TextField id="name" label="Name" margin="normal" style={{ width: '100%' }}
              onChange={e => setName(e.target.value)}
              value={name}
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
});
