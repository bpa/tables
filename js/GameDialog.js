import React from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, FormControl, IconButton, ListItem, ListItemText, TextField } from 'material-ui';
import Edit from 'material-ui-icons/Edit';
import CloseIcon from 'material-ui-icons/Close';
import ws from './Socket';

export default class GameItem extends React.Component {
  constructor(props) {
    super(props);
    this.onEnter = this.onEnter.bind(this);
    this.update = this.update.bind(this);
    this.state = {name:'',min:2,max:10};
  }

  onEnter() {
    this.setState(Object.assign({}, this.props.game));
  }

  update() {
    ws.send({cmd:'save_game', game: {
      id: this.state.id,
      min: parseInt(this.state.min),
      max: parseInt(this.state.max),
      name: this.state.name,
    }});
    this.props.onRequestClose()
  }

  on_change(f, e) {
    let v = e.target.value;
    if (f !== 'name' && v < 2) {
      v = 2;
    }
    this.setState({[f]: v});
  }

  render() {
    let game = this.props.game;
    return (
      <Dialog open={this.props.open} onEnter={this.onEnter}>
        <DialogTitle>{this.props.title}</DialogTitle>
        <DialogContent>
          <form noValidate autoComplete="off">
            <FormControl style={{minWidth: 200}}>
              <TextField id="game_name" label="Game" type="text" autoFocus
                value={this.state.name}
                InputLabelProps={{ shrink: true, }}
                onChange={this.on_change.bind(this, 'name')}
              />
            </FormControl>
            <br/>
            <FormControl margin="normal" style={{maxWidth: 100}}>
              <TextField id="min" label="Min" type="number"
                value={this.state.min}
                InputLabelProps={{ shrink: true, }}
                onChange={this.on_change.bind(this, 'min')}
              />
            </FormControl>
            <FormControl margin="normal" style={{maxWidth: 100, position:'absolute',right:'24px'}}>
              <TextField id="max" label="Max" type="number"
                value={this.state.max}
                InputLabelProps={{ shrink: true, }}
                onChange={this.on_change.bind(this, 'max')}
              />
            </FormControl>
          </form>
        </DialogContent>
        <DialogActions>
          <Button color="accent" onClick={this.props.onRequestClose}>Cancel</Button>
          <Button color="primary" onClick={this.update}>Update</Button>
        </DialogActions>
      </Dialog>
    );
  }
}
