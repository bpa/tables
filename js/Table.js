import React from 'react';
import { Button, Card, CardContent, CardActions, Divider, List, ListItem, ListItemIcon, ListItemText, Toolbar, Typography } from 'material-ui';
import People from 'material-ui-icons/People';
import Event from 'material-ui-icons/Event';
import Schedule from 'material-ui-icons/Schedule';
import Room from 'material-ui-icons/Room';
import moment from 'moment';
import { player } from './Login';
import ws from './Socket';

function Player(props) {
  let name = props.player ? props.player.fullName || '' : '';
  return (
    <ListItem dense>
      <ListItemText secondary={props.i + '.'}/>
      <ListItemText primary={name}/>
    </ListItem>
  );
}

function Players(props) {
  let players = props.players || [],
      total = Math.max(props.max, players.length),
      list = [];
  for (var i=0; i<total; i++) {
    list.push(<Player key={i} i={i+1} player={players[i]}/>);
  }
  return <List>{list}</List>;
}

function leave_table() {
  console.log("leave", this);
  ws.send({cmd:'leave_table', id: this.id, player: player.data});
}

function join_table() {
  console.log("join", this);
  ws.send({cmd:'join_table', id: this.id, player: player.data});
}

export default function Table(props) {
  let table = props.table,
      start = moment(table.start, moment.ISO_8601).format("h:mm A"),
      seat  = table.players.findIndex((p) => player.data && player.data.id === p.id),
      join  = seat === -1,
      owner = seat === 0;

  return (
    <Card elevation={4} style={{padding:4,margin:8}}>
      <CardContent>
        <List dense>
          <ListItem>
            <ListItemIcon><Event/></ListItemIcon>
            <ListItemText primary={table.game.name}/>
          </ListItem>
          <ListItem>
            <ListItemIcon><Schedule/></ListItemIcon>
            <ListItemText primary={start}/>
          </ListItem>
          <ListItem>
            <ListItemIcon><People/></ListItemIcon>
            <ListItemText primary={table.game.min + '-' + table.game.max}/>
          </ListItem>
          <ListItem>
            <ListItemIcon><Room/></ListItemIcon>
            <ListItemText primary={table.location}/>
          </ListItem>
        </List>
        <Divider/>
        <Players players={table.players} max={table.game.max}/>
      </CardContent>
      <CardActions>
      {join
        ? <Button dense onClick={join_table.bind(table)}>Join</Button>
        : <Button dense onClick={leave_table.bind(table)}>Leave</Button>
      }
      </CardActions>
    </Card>
  );
}

