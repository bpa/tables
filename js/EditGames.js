import React from 'react';
import MenuIcon from 'material-ui-icons/Menu';
import { AppBar, Dialog, IconButton, List, ListItem, ListItemIcon, ListItemText, Toolbar, Typography } from 'material-ui';
import Menu, { MenuItem } from 'material-ui/Menu';
import ws from './Socket';
import CloseIcon from 'material-ui-icons/Close';
import Slide from 'material-ui/transitions/Slide';
import GameItem from './GameItem';
import GameDialog from './GameDialog';

function Up(props) {
  return <Slide direction="up" {...props}/>;
}

export default class EditGames extends React.Component {
  constructor(props) {
    super(props);
    this.gcb = ws.register('games', this.on_games.bind(this));
    this.state = {games: {}, open:false, keys:[]}
  }

  cancel() {
    this.setState({open: false});
  }

  edit() {
    this.setState({open: true});
  }

	delete_game() {
	}

  on_games(msg) {
    let g = msg.games,
        k = Object.keys(g).sort((a,b)=>g[b].name < g[a].name);
    this.setState({games: msg.games, keys: k});
  }

  componentWillUnmount() {
    ws.deregister(this.gcb);
  }

  render() {
    let games = this.state.games;
    return (
      <Dialog fullScreen
        open={this.props.open}
        onRequestClose={this.props.onRequestClose}
        transition={Up}
      >
        <AppBar position="static">
          <Toolbar>
            <IconButton color="contrast" onClick={this.props.onRequestClose}>
              <CloseIcon/>
            </IconButton>
            <Typography type="title" color="inherit">
              Edit Games
            </Typography>
          </Toolbar>
        </AppBar>
        <List>
          {this.state.keys.map((k)=><GameItem game={games[k]} key={k}/>)}
          <ListItem button onClick={this.edit.bind(this)}>
            <ListItemText primary="Add new game"/>
          </ListItem>
          <GameDialog game={{name:'', min:2, max:10}} title="New Game"
            open={this.state.open} onRequestClose={this.cancel.bind(this)} />
        </List>
      </Dialog>
    );
  }
}
