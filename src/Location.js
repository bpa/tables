import React from 'react';
import { Button, IconButton, ListItem, ListItemText, TextField, Grid } from '@material-ui/core';
import Edit from '@material-ui/icons/Edit';
import Delete from '@material-ui/icons/Delete';

export default class Location extends React.Component {
  constructor(props) {
    super(props);
    this.change = this.change.bind(this);
    this.delete = this.delete.bind(this);
    this.edit = this.edit.bind(this);
    this.save = this.save.bind(this);
  }

  change(e) {
    this.setState({ location: e.target.value });
  }

  delete() {
    this.props.delete(this.props.location);
  }

  edit() {
    this.setState({ location: this.props.location });
    this.props.edit(this.props.location);
  }

  save() {
    this.props.save(this.state.location);
  }

  editing() {
    return (
      <ListItem>
        <TextField id="location" label="Location" type="text" autoFocus
          value={this.state.location}
          InputLabelProps={{ shrink: true, }}
          onChange={this.change}
        />
        <Button color="primary" onClick={this.save}>Save</Button>
        <Button color="secondary" onClick={this.props.cancel}>Cancel</Button>
      </ListItem>
    );
  }

  normal() {
    return (
      <ListItem>
        <Grid item xs={1}>
          <IconButton onClick={this.delete}><Delete /></IconButton>
        </Grid>
        <Grid item xs={1}>
          <IconButton onClick={this.edit}><Edit /></IconButton>
        </Grid>
        <Grid item xs={10}>
          <ListItemText primary={this.props.location} />
        </Grid>
      </ListItem>
    );
  }

  render() {
    return this.props.editing ? this.editing() : this.normal();
  }
}
