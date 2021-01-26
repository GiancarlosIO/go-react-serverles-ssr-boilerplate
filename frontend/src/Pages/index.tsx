import * as React from 'react';
import { Switch, Route } from 'react-router-dom';

import Homepage from './Homepage';

const Pages = () => {
  return (
    <div>
      <nav>MainMenu</nav>
      <Switch>
        <Route path="/" exact>
          <Homepage />
        </Route>
      </Switch>
    </div>
  );
};

export default Pages;
