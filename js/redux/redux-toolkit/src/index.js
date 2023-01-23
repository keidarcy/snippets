import React from 'react';
import ReactDOM from 'react-dom';

import { Provider } from 'react-redux';
import { Theme } from '@twilio-paste/core/theme';

import Application from './components/Application';
import { store } from './store';

import './index.css';

ReactDOM.render(
  <Provider store={store}>
    <Theme.Provider theme="default">
      <React.StrictMode>
        <Application />
      </React.StrictMode>
    </Theme.Provider>
  </Provider>,
  document.getElementById('root')
);
