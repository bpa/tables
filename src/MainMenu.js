import React, { useState } from 'react';
import MenuIcon from '@material-ui/icons/Menu';
import { IconButton } from '@material-ui/core';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import EditGames from './EditGames';
import EditLocations from './EditLocations';
import send from './Socket';

export default function MainMenu() {
  let [type, setType] = useState('');
  let [anchor, setAnchor] = useState(undefined);

  function edit(t) {
    send({ cmd: 'list_' + t });
    setType(t);
    setAnchor(undefined);
  }

  function stopEditing() {
    setType('');
  }

  function logout() {
    send({ cmd: 'logout' });
    setAnchor(undefined);
  }

  function openMenu(e) {
    setAnchor(e.currentTarget);
  }

  function closeMenu() {
    setAnchor(undefined);
  }

  return (
    <>
      <IconButton color="inherit" aria-label="Menu"
        style={{ marginLeft: -12, marginRight: 20 }}
        onClick={openMenu}>
        <MenuIcon />
      </IconButton>
      <Menu id="main-menu" anchorEl={anchor}
        keepMounted
        open={!!anchor}
        onClose={closeMenu}
      >
        <MenuItem onClick={() => edit('games')}>Edit Games</MenuItem>
        <MenuItem onClick={() => edit('locations')}>Edit Locations</MenuItem>
        <MenuItem onClick={logout}>Log out</MenuItem>
      </Menu>
      <EditGames open={type === 'games'}
        onClose={stopEditing} />
      <EditLocations open={type === 'locations'}
        onClose={stopEditing} />
    </>
  );
}
