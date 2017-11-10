import React from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton, ListItem, ListItemText } from 'material-ui';
import Delete from 'material-ui-icons/Delete';
import Edit from 'material-ui-icons/Edit';
import GameDialog from './GameDialog';
import ws from './Socket';

export default class GameItem extends React.Component {
  constructor(props) {
    super(props);
    this.cancel = this.cancel.bind(this);
    this.edit = this.edit.bind(this);
    this.delete_game = this.delete_game.bind(this);
    this.delete_confirmed = this.delete_confirmed.bind(this);
    this.delete_cancel = this.delete_cancel.bind(this);
    this.state = {open: false, confirming: false};
  }

  cancel() {
    this.setState({open: false});
  }

  edit() {
    this.setState({open: true});
  }

	delete_game() {
    this.setState({confirming: true});
	}

  delete_confirmed() {
    ws.send({cmd:'delete_game', game: this.props.game.id});
    this.setState({confirming: false});
  }

  delete_cancel() {
    this.setState({confirming: false});
  }

  render() {
    let game = this.props.game,
        players = game.min == game.max
          ? game.min + ' players'
          : game.min + '-' + game.max + ' players';
    return (
      <div>
        <ListItem>
          <IconButton onClick={this.delete_game}><Delete/></IconButton>
          <IconButton onClick={this.edit}><Edit/></IconButton>
          <ListItemText primary={players}/>
          <ListItemText primary={game.name}/>
        </ListItem>
        <GameDialog game={game} title="Edit Game"
          open={this.state.open} onRequestClose={this.cancel}/>
				<Dialog open={this.state.confirming} onRequestClose={this.delete_cancel}>
          <DialogTitle>Confirm Deletion</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Are you sure you want to delete '{game.name}'?
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.delete_cancel} color="accent">
              Cancel
            </Button>
            <Button onClick={this.delete_confirmed} color="primary" autoFocus>
              Delete
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}
