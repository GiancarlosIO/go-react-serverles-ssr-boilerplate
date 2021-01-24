import * as React from 'react';
import * as ReactDOM from 'react-dom';

import App from './Pages';

const mainNode = document.querySelector('#app');
const needSSR = mainNode?.textContent;
const renderer: ReactDOM.Renderer = needSSR
  ? ReactDOM.hydrate
  : ReactDOM.render;

renderer(<App />, mainNode);
