import * as React from 'react';

import Header from '@/Components/Header';

const Common: React.FC = ({ children }) => {
  return (
    <div>
      <Header />
      {children}
    </div>
  );
};

export default Common;
