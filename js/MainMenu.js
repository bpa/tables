import React from 'react';
import MenuIcon from '@material-ui/icons/Menu';
import { IconButton } from '@material-ui/core';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import ws from './Socket';
import EditGames from './EditGames';
import EditLocations from './EditLocations';

export default class MainMenu extends React.Component {
  constructor() {
    super();
    this.closeMenu = this.closeMenu.bind(this);
    this.logout = this.logout.bind(this);
    this.openMenu = this.openMenu.bind(this);
    this.stopEditing = this.stopEditing.bind(this);
    this.state = {open: false, anchor: null, type: null};
  }

  edit(type) {
    ws.send({cmd: 'list_' + type});
    this.setState({open: false, type: type});
  }

  stopEditing() {
    this.setState({type: null});
  }

  logout() {
    ws.send({cmd: 'logout'});
    this.setState({open: false});
  }

  openMenu(e) {
    this.setState({open: true, anchor: e.currentTarget});
  }

  closeMenu() {
    this.setState({open: false});
  }

  render() {
    return (
      <div>
        <IconButton color="inherit" aria-label="Menu"
            style={{marginLeft:-12, marginRight:20}}
            onClick={this.openMenu}>
          <MenuIcon/>
        </IconButton>
        <Menu id="main-menu" anchorEl={this.state.anchor}
          open={this.state.open}
          onClose={this.closeMenu}
        >
          <MenuItem onClick={this.edit.bind(this, 'games')}>Edit Games</MenuItem>
          <MenuItem onClick={this.edit.bind(this, 'locations')}>Edit Locations</MenuItem>
          <MenuItem onClick={this.logout}>Log out</MenuItem>
        </Menu>
        <EditGames open={this.state.type === 'games'}
          onClose={this.stopEditing}/>
        <EditLocations open={this.state.type === 'locations'}
          onClose={this.stopEditing}/>
      </div>
    );
  }
}
