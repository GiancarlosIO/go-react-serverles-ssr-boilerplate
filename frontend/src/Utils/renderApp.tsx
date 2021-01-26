import * as React from 'react';
import * as ReactDOM from 'react-dom';

const renderApp = (
  App: React.SFCElement<any> | React.SFCElement<any>[],
): void => {
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

  renderer(App, mainNode);
};

export default renderApp;
