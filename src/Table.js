import React from 'react';
import { Button, Card, CardContent, CardActions, Divider, List, ListItem, ListItemIcon, ListItemText } from '@material-ui/core';
import People from '@material-ui/icons/People';
import Event from '@material-ui/icons/Event';
import Schedule from '@material-ui/icons/Schedule';
import Room from '@material-ui/icons/Room';
import moment from 'moment';
import { player } from './Login';
import ws from './Socket';

function Player(props) {
  let name = props.player ? props.player.fullName || '' : '';
  return (
    <ListItem dense={true}>
      <ListItemText secondary={props.i + '.'} />
      <ListItemText primary={name} />
    </ListItem>
  );
}

function Players(props) {
  let players = props.players || [],
    total = Math.max(props.max, players.length),
    list = [];
  for (var i = 0; i < total; i++) {
    list.push(<Player key={i} i={i + 1} player={players[i]} />);
  }
  return <List>{list}</List>;
}

function leave_table() {
  ws.send({ cmd: 'leave_table', id: this.id });
}

function join_table() {
  ws.send({ cmd: 'join_table', id: this.id });
}

export default function Table(props) {
  let table = props.table,
    start = moment(table.start, moment.ISO_8601).format("h:mm A"),
    seat = table.players.findIndex((p) => player.data && player.data.id === p.id),
    join = seat === -1,
    // owner = seat === 0,
    players = table.game.min === table.game.max
      ? table.game.min + ' players'
      : table.game.min + '-' + table.game.max + ' players';

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
          ? <Button dense="true" onClick={join_table.bind(table)}>Join</Button>
          : <Button dense="true" onClick={leave_table.bind(table)}>Leave</Button>
        }
      </CardActions>
    </Card>
  );
}

