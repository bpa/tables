import React, { useContext, useState } from 'react';
import { Button, Card, CardContent, CardActions, Divider, List, ListItem, ListItemIcon, ListItemText, Dialog, DialogTitle, DialogContent, DialogActions, DialogContentText } from '@material-ui/core';
import DeleteForever from '@material-ui/icons/DeleteForever';
import People from '@material-ui/icons/People';
import Event from '@material-ui/icons/Event';
import Schedule from '@material-ui/icons/Schedule';
import Room from '@material-ui/icons/Room';
import moment from 'moment';
import send from './Socket';
import { observer } from 'mobx-react-lite';
import GlobalContext from './GlobalState';

function Player(props) {
  let name = (props.player && props.player.fullName) || '';
  return (
    <ListItem dense={true}>
      <ListItemText secondary={props.i + '.'} />
      <ListItemText primary={name} />
    </ListItem>
  );
}

function Players(props) {
  let players = props.players || [],
    p = players.length,
    total = Math.max(props.max, players.length),
    list = [];
  for (var i = 0; i < total; i++) {
    list.push(<Player key={i} i={i + 1} player={i < p && players[i]} />);
  }
  return <List>{list}</List>;
}

function leave_table() {
  send({ cmd: 'leave_table', id: this.id });
}

function join_table() {
  send({ cmd: 'join_table', id: this.id });
}

export default observer((props) => {
  let context = useContext(GlobalContext);
  let [open, setOpen] = useState(false);

  let me = context.player.id;
  let table = props.table,
    start = moment(table.start, moment.ISO_8601).format("h:mm A"),
    seat = table.players.findIndex((p) => p.id === me),
    join = seat === -1,
    owner = seat === 0,
    players = table.game.min === table.game.max
      ? table.game.min + ' players'
      : table.game.min + '-' + table.game.max + ' players';

  function deleteGame() {
    send({cmd: 'delete_table', id: table.id })
  }

  return (
    <Card elevation={4} style={{ padding: 4, margin: 8 }}>
      <CardContent>
        <List dense={true}>
          <ListItem>
            <ListItemIcon><Event /></ListItemIcon>
            <ListItemText primary={table.game.name} />
          </ListItem>
          <ListItem>
            <ListItemIcon><Schedule /></ListItemIcon>
            <ListItemText primary={start} />
          </ListItem>
          <ListItem>
            <ListItemIcon><People /></ListItemIcon>
            <ListItemText primary={players} />
          </ListItem>
          <ListItem>
            <ListItemIcon><Room /></ListItemIcon>
            <ListItemText primary={table.location} />
          </ListItem>
        </List>
        <Divider />
        <Players players={table.players} max={table.game.max} />
      </CardContent>
      <CardActions>
        {join
          ? <Button dense="true" color="primary" onClick={join_table.bind(table)}>Join</Button>
          : <Button dense="true" color="primary" onClick={leave_table.bind(table)}>Leave</Button>
        }
        {owner && <Button dense="true" color="secondary" onClick={() => setOpen(true)}>
          Delete<DeleteForever />
        </Button>
        }
      </CardActions>
      <Dialog open={open}>
        <DialogTitle>Confirm Table Deletion</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete {table.game.name}?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button dense="true" color="primary" onClick={deleteGame}>Delete</Button>
          <Button dense="true" color="secondary" onClick={() => setOpen(false)}>Cancel</Button>
        </DialogActions>
      </Dialog>
    </Card>
  );
});

