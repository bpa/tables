import React from 'react';
import { IconButton, ListItem, ListItemText, Grid } from '@material-ui/core';
import Delete from '@material-ui/icons/Delete';
import Edit from '@material-ui/icons/Edit';

export default function GameItem(props) {
  let game = props.game;
  let players = game.min === game.max
    ? game.min + ' players'
    : game.min + '-' + game.max + ' players';

  return (
    <>
      <ListItem>
        <Grid item xs={1}>
          <IconButton onClick={props.delete}><Delete /></IconButton>
        </Grid>
        <Grid item xs={1}>
          <IconButton onClick={props.edit}><Edit /></IconButton>
        </Grid>
        <Grid item xs={4}><ListItemText primary={players} /></Grid>
        <Grid item xs={6}><ListItemText primary={game.name} /></Grid>
      </ListItem>
    </>
  );
}
