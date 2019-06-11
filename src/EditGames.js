import React, { useContext, useState } from 'react';
import { AppBar, Dialog, IconButton, List, ListItem, ListItemText, Toolbar, Typography } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import Slide from '@material-ui/core/Slide';
import GameItem from './GameItem';
import GameDialog from './GameDialog';
import GlobalContext from './GlobalState';
import { observer } from 'mobx-react-lite';

const Up = React.forwardRef(function Up(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default observer((props) => {
  let context = useContext(GlobalContext);
  let [open, setOpen] = useState(false);

  return (
    <Dialog fullScreen
      open={props.open}
      onClose={props.onClose}
      TransitionComponent={Up}
    >
      <AppBar position="static">
        <Toolbar>
          <IconButton color="inherit" onClick={props.onClose}>
            <CloseIcon />
          </IconButton>
          <Typography type="title" color="inherit">
            Edit Games
            </Typography>
        </Toolbar>
      </AppBar>
      <List>
        {context.games.map((g) => <GameItem game={g} key={g.name} />)}
        <ListItem button onClick={() => setOpen(true)}>
          <ListItemText primary="Add new game" />
        </ListItem>
        <GameDialog game={{ name: '', min: 2, max: 10 }} title="New Game"
          open={open} onClose={() => setOpen(false)} />
      </List>
    </Dialog>
  );
});
