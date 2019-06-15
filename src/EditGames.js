import React, { useContext, useState } from 'react';
import { AppBar, Dialog, IconButton, List, ListItem, ListItemText, Toolbar, Typography, DialogTitle, DialogContent, DialogContentText, DialogActions, Button } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import Slide from '@material-ui/core/Slide';
import GameItem from './GameItem';
import GameDialog from './GameDialog';
import GlobalContext from './GlobalState';
import { observer } from 'mobx-react-lite';
import send from './Socket';

const Up = React.forwardRef(function Up(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default observer((props) => {
  let context = useContext(GlobalContext);
  let [editing, setEditing] = useState(false);
  let [deleting, setDeleting] = useState(false);

  function delete_confirmed() {
    send({ cmd: 'delete_game', game: deleting.id });
    setDeleting(false);
  }

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
        {context.games.map((g) =>
          <GameItem game={g} key={g.name} delete={() => setDeleting(g)} edit={() => setEditing(g)} />
        )}
        <ListItem button onClick={() => setEditing({ name: '', min: 2, max: 10 })}>
          <ListItemText primary="Add new game" />
        </ListItem>
      </List>

      <GameDialog game={editing} title={editing.id ? "Edit Game" : "New Game"}
        open={!!editing} onClose={() => setEditing(false)} />

      <Dialog open={!!deleting} onClose={() => setDeleting(false)}>
        <DialogTitle>Confirm Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete '{deleting.name}'?
            </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleting(false)} color="secondary">
            Cancel
          </Button>
          <Button onClick={delete_confirmed} color="primary" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </Dialog>
  );
});
