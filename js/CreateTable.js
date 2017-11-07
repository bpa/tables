import React from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, FormControl,
  IconButton, Input, InputLabel, MenuItem, Select, TextField } from 'material-ui';

import AddIcon from 'material-ui-icons/Add';
import Edit from 'material-ui-icons/Edit';
import ws from './Socket';
import moment from 'moment';

export default class CreateTable extends React.Component {
  constructor() {
    super();
    this.open = this.open.bind(this);
    this.close = this.close.bind(this);
    this.create = this.create.bind(this);
    this.state = {open: false};
    this.locCb = ws.register('locations', this.on_locations.bind(this));
    this.gameCb = ws.register('games', this.on_games.bind(this));
    this.state = { locations: [], location: '', games: [], game: '' };
  }

  componentWillUnmount() {
    ws.deregister(this.locCb);
    ws.deregister(this.gameCb);
  }

  open() {
    ws.send({cmd:'new_game'});
    this.setState({open: true, game:'', location:'', start: '12:00'});
  }

  close() {
    ws.send({cmd:'cancel_game_creation'});
    this.setState({open: false});
  }

  create() {
    let start = this.state.start.split(':'),
      now = moment();
    now.hour(start[0]);
    now.minute(start[1]);
    console.log(now.hour());
    ws.send({
      cmd:'create_game',
      game:this.state.game,
      location:this.state.location,
      start:now,
    });
    this.setState({open: false});
  }

  on_change(field,b) {
    this.setState({[field]: b.target.value});
  }

  on_locations(msg) {
    this.setState({locations: msg.locations});
  }

  on_games(msg) {
    this.setState({games: Object.keys(msg.games)});
  }

  render() {
    return (
      <div>
        <Button fab style={{alignSelf:"center", margin:"100px"}} color="primary" aria-label="add" onClick={this.open}>
          <AddIcon/>
        </Button>
        <Dialog open={this.state.open} onRequestClose={this.close}>
          <DialogTitle>Create Table</DialogTitle>
          <DialogContent>
            <form noValidate autoComplete="off">
              <FormControl style={{minWidth: 200}}>
                <InputLabel htmlFor="game">Game</InputLabel>
                <Select
                  value={this.state.game}
                  onChange={this.on_change.bind(this, 'game')}
                  input={<Input id="game"/>}
                >
                  {this.state.games.map((g) => (
                        <MenuItem value={g} key={g}>{g}</MenuItem>))}
                </Select>
              </FormControl>
              <br/>
              <FormControl style={{minWidth: 200}}>
                <InputLabel htmlFor="location">Location</InputLabel>
                <Select
                  value={this.state.location}
                  onChange={this.on_change.bind(this, 'location')}
                  input={<Input id="location"/>}
                >
                  {this.state.locations.map((l) => (
                        <MenuItem value={l} key={l}>{l}</MenuItem>))}
                </Select>
              </FormControl>
              <br/>
              <FormControl style={{minWidth: 200}}>
                <TextField id="time" label="Time" type="time"
                  defaultValue="12:00"
                  InputLabelProps={{ shrink: true, }}
                  inputProps={{ step: 900 }} // 15 min
                  onChange={this.on_change.bind(this, 'start')}
                />
              </FormControl>
            </form>
          </DialogContent>
          <DialogActions>
            <Button color="accent" onClick={this.close}>Cancel</Button>
            <Button color="primary" onClick={this.create}>Create</Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}
