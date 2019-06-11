import React from 'react';
import { observer } from 'mobx-react-lite';
import { AppBar, Toolbar, Typography } from '@material-ui/core';
import CreateTable from './CreateTable';
import MainMenu from './MainMenu';
import Table from './Table';
import Login from './Login';
import GlobalContext, { state } from './GlobalState';

export default observer(() => {
  return (
    <GlobalContext.Provider value={state}>
      <AppBar position="static">
        <Toolbar>
          <MainMenu />
          <Typography type="title" color="inherit" style={{ flex: 1 }}>
            Tables
          </Typography>
          <Login />
        </Toolbar>
      </AppBar>
      <div style={{ display: 'flex', alignItems: 'flex-start' }}>
        {state.tables.map((t) => <Table table={t} key={t.id} />)}
        <CreateTable />
      </div>
    </GlobalContext.Provider>
  );
});
