import React, { useContext, useState } from 'react';
import { AppBar, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton, List, ListItem, TextField, Toolbar, Typography } from '@material-ui/core';
import CloseIcon from '@material-ui/icons/Close';
import Slide from '@material-ui/core/Slide';
import Location from './Location';
import GlobalContext from './GlobalState';
import { observer } from 'mobx-react-lite';
import send from './Socket';

const Up = React.forwardRef(function Up(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default observer((props) => {
  let context = useContext(GlobalContext);
  let [open, setOpen] = useState(false);
  let [input, setInput] = useState('');
  let [editing, setEditing] = useState('');
  let [toDelete, setToDelete] = useState('');

  function add_location() {
    send({ cmd: "create_location", location: input });
    setInput('');
  }

  function confirm_delete(l) {
    setToDelete(l);
    setOpen(true);
  }

  function delete_location() {
    send({ cmd: 'delete_location', location: toDelete });
    setOpen(false);
  }

  function complete_edit(new_value) {
    send({ cmd: 'edit_location', from: editing, to: new_value });
    setEditing('');
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
            Edit Locations
          </Typography>
        </Toolbar>
      </AppBar>
      <List>
        {context.locations.map((l) => <Location
          location={l}
          key={l}
          edit={() => setEditing(l)}
          cancel={() => setEditing('')}
          delete={confirm_delete.bind(null, l)}
          save={complete_edit}
          editing={editing === l}
        />)}
        <ListItem>
          <TextField id="location" label="Location" type="text"
            InputLabelProps={{ shrink: true, }}
            value={input}
            onChange={e => setInput(e.target.value)}
          />
          <Button onClick={add_location}>Add</Button>
        </ListItem>
      </List>
      <Dialog open={open} onClose={() => setOpen(false)}>
        <DialogTitle>Confirm Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            {"Are you sure you want to delete '" + toDelete + "'?"}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)} color="secondary">
            Cancel
          </Button>
          <Button onClick={delete_location} color="primary" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </Dialog>
  );
});
