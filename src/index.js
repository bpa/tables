import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import send from './Socket';

window.onerror = function (messageOrEvent, source, lineno, colno, error) {
    send({
        cmd: 'error',
        message: error.message,
        stack: error.stack
    });
};

ReactDOM.render(<App />, document.getElementById('root'));
