import * as React from 'react';

import Common from './Common';

import '../css/styles.scss';

const ClientLayout: React.FC = ({ children }) => {
  return <Common>{children}</Common>;
};

export default ClientLayout;
