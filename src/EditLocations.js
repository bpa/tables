import React from 'react';
import { AppBar, Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton, List, ListItem, TextField, Toolbar, Typography } from '@material-ui/core';
import ws from './Socket';
import CloseIcon from '@material-ui/icons/Close';
import Slide from '@material-ui/core/Slide';
import Location from './Location';

const Up = React.forwardRef(function Up(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default class EditLocations extends React.Component {
  constructor(props) {
    super(props);
    this.add_location = this.add_location.bind(this);
    this.cancel_edit = this.cancel_edit.bind(this);
    this.cancel_delete = this.cancel_delete.bind(this);
    this.complete_edit = this.complete_edit.bind(this);
    this.delete_location = this.delete_location.bind(this);
    this.lcb = ws.register('locations', this.on_locations.bind(this));
    this.state = { locations: [], input: '', open: false };
  }

  on_change(e) {
    this.setState({ input: e.target.value });
  }

  on_locations(msg) {
    this.setState({ locations: msg.locations });
  }

  add_location() {
    ws.send({ cmd: "create_location", location: this.state.input });
    this.setState({ input: '' });
  }

  edit_location(l) {
    this.setState({ editing: l });
  }

  cancel_edit() {
    this.setState({ editing: null });
  }

  confirm_delete(l) {
    this.setState({ to_delete: l, open: true });
  }

  delete_location() {
    ws.send({ cmd: 'delete_location', location: this.state.to_delete });
    this.setState({ open: false });
  }

  cancel_delete() {
    this.setState({ open: false });
  }

  complete_edit(new_value) {
    ws.send({ cmd: 'edit_location', from: this.state.editing, to: new_value });
    this.setState({ editing: null });
  }

  componentWillUnmount() {
    ws.deregister(this.lcb);
  }

  render() {
    return (
      <Dialog fullScreen
        open={this.props.open}
        onClose={this.props.onClose}
        TransitionComponent={Up}
      >
        <AppBar position="static">
          <Toolbar>
            <IconButton color="inherit" onClick={this.props.onClose}>
              <CloseIcon />
            </IconButton>
            <Typography type="title" color="inherit">
              Edit Locations
            </Typography>
          </Toolbar>
        </AppBar>
        <List>
          {this.state.locations.map((l) => <Location
            location={l}
            key={l}
            edit={this.edit_location.bind(this, l)}
            cancel={this.cancel_edit}
            delete={this.confirm_delete.bind(this, l)}
            save={this.complete_edit}
            editing={this.state.editing === l}
          />)}
          <ListItem>
            <TextField id="location" label="Location" type="text"
              InputLabelProps={{ shrink: true, }}
              value={this.state.input}
              onChange={this.on_change.bind(this)}
            />
            <Button onClick={this.add_location}>Add</Button>
          </ListItem>
        </List>
        <Dialog open={this.state.open} onClose={this.cancel_delete}>
          <DialogTitle>Confirm Deletion</DialogTitle>
          <DialogContent>
            <DialogContentText>
              {"Are you sure you want to delete '" + this.state.to_delete + "'?"}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.cancel_delete} color="secondary">
              Cancel
            </Button>
            <Button onClick={this.delete_location} color="primary" autoFocus>
              Delete
            </Button>
          </DialogActions>
        </Dialog>
      </Dialog>
    );
  }
}
