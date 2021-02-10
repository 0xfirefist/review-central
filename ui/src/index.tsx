import React from 'react';
import ReactDOM from 'react-dom';
import Page from './page';
import { CookiesProvider } from "react-cookie";

ReactDOM.render(
  <CookiesProvider>    
    <Page />
  </CookiesProvider>,
  document.getElementById('root')
);
