import React, { useContext, useState } from 'react';
import {
  Button, Dialog, DialogActions, DialogContent, DialogTitle, FormControl,
  Input, InputLabel, MenuItem, Select, TextField
} from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import { Add } from '@material-ui/icons';
import send from './Socket';
import moment from 'moment';
import GlobalContext from './GlobalState';
import { observer } from 'mobx-react-lite';

export default observer(() => {
  let context = useContext(GlobalContext);
  let [open, setOpen] = useState(false);
  let [game, setGame] = useState('');
  let [location, setLocation] = useState('');
  let [start, setStart] = useState('');

  function openDialog() {
    send({ cmd: 'new_game' });
    setOpen(true);
    setGame('');
    setLocation('');
    setStart('12:00');
  }

  function cancel() {
    send({ cmd: 'cancel_table_creation' });
    setOpen(false);
  }

  function create() {
    const [hour, minute] = start.split(':');
    let now = moment();
    now.hour(hour);
    now.minute(minute);
    send({
      cmd: 'create_table',
      game: game.id,
      location: location,
      start: now,
      player: context.player,
    });
    setOpen(false);
  }

  let g = context.games;

  return (
    <>
      <Fab style={{ alignSelf: "center", margin: "100px" }} color="primary" aria-label="add" onClick={openDialog}>
        <Add />
      </Fab>
      <Dialog open={open}>
        <DialogTitle>Create Table</DialogTitle>
        <DialogContent>
          <form noValidate autoComplete="off">
            <FormControl style={{ minWidth: 200 }}>
              <InputLabel htmlFor="game">Game</InputLabel>
              <Select
                value={game}
                onChange={(e) => setGame(e.target.value)}
                input={<Input id="game" />}
              >
                {Object.keys(g)
                  .sort((a, b) => g[b].name < g[a].name)
                  .map((id) => (
                    <MenuItem value={g[id]} key={id}>{g[id].name}</MenuItem>
                  ))}
              </Select>
            </FormControl>
            <br />
            <FormControl style={{ minWidth: 200 }}>
              <InputLabel htmlFor="location">Location</InputLabel>
              <Select
                value={location}
                onChange={e => setLocation(e.target.value)}
                input={<Input id="location" />}
              >
                {context.locations.map((l) => (
                  <MenuItem value={l} key={l}>{l}</MenuItem>))}
              </Select>
            </FormControl>
            <br />
            <FormControl style={{ minWidth: 200 }}>
              <TextField id="time" label="Time" type="time"
                value={start}
                InputLabelProps={{ shrink: true, }}
                inputProps={{ step: 900 }} // 15 min
                onChange={e => setStart(e.target.value)}
              />
            </FormControl>
          </form>
        </DialogContent>
        <DialogActions>
          <Button color="secondary" onClick={cancel}>Cancel</Button>
          <Button color="primary" onClick={create}>Create</Button>
        </DialogActions>
      </Dialog>
    </>
  );
});
