import React from 'react';
import { observable } from 'mobx';

var GlobalContext = React.createContext({});
export default GlobalContext;

export var state = observable({
    player: (window.localStorage.player && JSON.parse(window.localStorage.player)) || {},
    games: [],
    locations: [],
    tables: [],
});

export const action = {
    tables: (msg) => state.tables = msg.tables,

    login: function (msg) {
        state.player = msg.player;
        window.localStorage.player = JSON.stringify(msg.player);
    },

    logout: function () {
        window.localStorage.removeItem('player');
        state.player = {};
    },

    locations: (msg) => state.locations = msg.locations.sort((a, b) => a.localeCompare(b, undefined, { sensitivity: 'base'})),

    games: (msg) => state.games = msg.games.sort((a, b) => a.name.localeCompare(b.name)),
};
