import React from 'react';
import { Divider, List, ListItem, ListItemIcon, ListItemText, Paper, Toolbar, Typography } from 'material-ui';
import People from 'material-ui-icons/People';
import Event from 'material-ui-icons/Event';
import Schedule from 'material-ui-icons/Schedule';
import Room from 'material-ui-icons/Room';
import moment from 'moment';

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

export default function Table(props) {
  let { table } = props, start = moment(table.start, moment.ISO_8601);
  return (
      <Paper elevation={4} style={{padding:4,margin:8}}>
        <List dense>
          <ListItem>
            <ListItemIcon><Event/></ListItemIcon>
            <ListItemText primary={table.game.name}/>
          </ListItem>
          <ListItem>
            <ListItemIcon><Schedule/></ListItemIcon>
            <ListItemText primary={start.format("h:mm A")}/>
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
        <Divider/>
      </Paper>
  );
}

