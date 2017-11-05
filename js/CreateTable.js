import React from 'react';
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, TextField } from 'material-ui';

import AddIcon from 'material-ui-icons/Add';

export default class CreateTable extends React.Component {
  constructor() {
    super();
    this.open = this.open.bind(this);
    this.close = this.close.bind(this);
    this.create = this.create.bind(this);
    this.state = {open: false};
  }

  open() {
    this.setState({open: true});
  }

  close() {
    this.setState({open: false});
  }

  create() {
    this.setState({open: false});
  }

  render() {
    return (
      <div>
        <Button fab style={{alignSelf:"center", margin:"100px"}} color="primary" aria-label="add" onClick={this.open}>
          <AddIcon/>
        </Button>
        <Dialog open={this.state.open} onRequestClose={this.close}>
          <DialogTitle>Create Table</DialogTitle>
          <DialogContent>
            <form noValidate>
            <TextField id="time" label="Time" type="time"
            defaultValue="12:00"
            InputLabelProps={{
              shrink: true,
            }}
            inputProps={{
              step: 900, // 15 min
            }}
            />
            </form>
          </DialogContent>
          <DialogActions>
            <Button color="accent" onClick={this.close}>Cancel</Button>
            <Button color="primary" onClick={this.create}>Create</Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}
