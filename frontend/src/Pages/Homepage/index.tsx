import * as React from 'react';

import renderApp from '@/Utils/renderApp';
import ClientLayout from '@/Layouts/ClientLayout';

import Homepage from './Homepage';

renderApp(
  <ClientLayout>
    <Homepage />
  </ClientLayout>,
);
