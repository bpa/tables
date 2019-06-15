import React, { useState } from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton, ListItem, ListItemText, GridListTile, Grid } from '@material-ui/core';
import Delete from '@material-ui/icons/Delete';
import Edit from '@material-ui/icons/Edit';
import GameDialog from './GameDialog';
import send from './Socket';

export default function GameItem(props) {
  let [open, setOpen] = useState(false);
  let [confirming, setConfirming] = useState(false);

  function delete_confirmed() {
    send({ cmd: 'delete_game', game: props.game.id });
    setConfirming(false);
  }

  let game = props.game,
    players = game.min === game.max
      ? game.min + ' players'
      : game.min + '-' + game.max + ' players';

  return (
    <>
      <ListItem>
        <Grid item xs={1}>
          <IconButton onClick={() => setConfirming(true)}><Delete /></IconButton>
        </Grid>
        <Grid item xs={1}>
          <IconButton onClick={() => setOpen(true)}><Edit /></IconButton>
        </Grid>
        <Grid item xs={4}><ListItemText primary={players} /></Grid>
        <Grid item xs={6}><ListItemText primary={game.name} /></Grid>
      </ListItem>
      <GameDialog game={game} title="Edit Game"
        open={open} onClose={() => setOpen(false)} />
      <Dialog open={confirming} onClose={() => setConfirming(false)}>
        <DialogTitle>Confirm Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete '{game.name}'?
            </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setConfirming(false)} color="secondary">
            Cancel
            </Button>
          <Button onClick={delete_confirmed} color="primary" autoFocus>
            Delete
            </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
