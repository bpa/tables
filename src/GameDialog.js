import React, { useState } from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, TextField } from '@material-ui/core';
import send from './Socket';

export default function GameItem(props) {
  let [name, setName] = useState('');
  let [min, setMin] = useState(2);
  let [max, setMax] = useState(10);

  function update() {
    send({
      cmd: 'save_game', game: {
        id: props.game.id,
        min: parseInt(min),
        max: parseInt(max),
        name: name,
      }
    });
    props.onClose()
  }

  function editGame() {
    setName(props.game.name);
    setMin(props.game.min);
    setMax(props.game.max);
  }

  return (
    <Dialog open={props.open} onEnter={editGame}>
      <DialogTitle>{props.title}</DialogTitle>
      <DialogContent>
        <form noValidate autoComplete="off">
          <FormControl style={{ minWidth: 200 }}>
            <TextField id="game_name" label="Game" type="text" autoFocus
              value={name}
              InputLabelProps={{ shrink: true, }}
              onChange={e => setName(e.target.value)}
            />
          </FormControl>
          <br />
          <FormControl margin="normal" style={{ maxWidth: 100 }}>
            <TextField id="min" label="Min" type="number"
              value={min}
              InputLabelProps={{ shrink: true, }}
              onChange={e => setMin(e.target.value < 2 ? 2 : e.target.value)}
            />
          </FormControl>
          <FormControl margin="normal" style={{ maxWidth: 100, position: 'absolute', right: '24px' }}>
            <TextField id="max" label="Max" type="number"
              value={max}
              InputLabelProps={{ shrink: true, }}
              onChange={e => setMax(e.target.value < min ? min : e.target.value)}
            />
          </FormControl>
        </form>
      </DialogContent>
      <DialogActions>
        <Button color="secondary" onClick={props.onClose}>Cancel</Button>
        <Button color="primary" onClick={update}>Update</Button>
      </DialogActions>
    </Dialog>
  );
}
