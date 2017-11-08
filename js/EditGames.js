import React from 'react';
import MenuIcon from 'material-ui-icons/Menu';
import { AppBar, Dialog, IconButton, List, ListItem, ListItemIcon, ListItemText, Toolbar, Typography } from 'material-ui';
import Menu, { MenuItem } from 'material-ui/Menu';
import ws from './Socket';
import CloseIcon from 'material-ui-icons/Close';
import Slide from 'material-ui/transitions/Slide';

function Game(props) {
  let game = props.game;
  return (
    <ListItem>
      <ListItemText primary={game.name}/>
    </ListItem>
  );
}

function Up(props) {
  return <Slide direction="up" {...props}/>;
}

export default class EditGames extends React.Component {
  constructor(props) {
    super(props);
    this.gcb = ws.register('games', this.on_games.bind(this));
    this.state = {games: {}}
  }

  on_games(msg) {
    this.setState({games: msg.games});
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
        {Object.keys(games).map((k)=><Game game={games[k]}/>)}
        <ListItem button>
          <ListItemText primary="Add new game"/>
        </ListItem>
        </List>
      </Dialog>
    );
  }
}
