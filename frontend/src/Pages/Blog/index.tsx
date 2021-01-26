import * as React from 'react';

import renderApp from '@/Utils/renderApp';
import ClientLayout from '@/Layouts/ClientLayout';

import Blog from './Blog';

renderApp(
  <ClientLayout>
    <Blog />
  </ClientLayout>,
);
