import './css/styles.scss';

import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';

import App from './Pages';

const mainNode = document.querySelector('#app');
const needSSR = mainNode?.textContent;
const renderer: ReactDOM.Renderer = needSSR
  ? ReactDOM.hydrate
  : ReactDOM.render;

if (needSSR) {
  console.log(
    '%c [Serverless SSR Enabled]: Using ReactDOM.Hydrate to render the React App! ',
    'color: #1794cd',
  );
}

renderer(
  <Router>
    <App />
  </Router>,
  mainNode,
);
