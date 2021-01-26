import 'source-map-support/register';

import * as React from 'react';
import * as ReactDOMServer from 'react-dom/server'
import { StaticRouter } from 'react-router-dom'

import type { ValidatedEventAPIGatewayProxyEvent } from '@libs/apiGateway';
import { formatJSONResponse } from '@libs/apiGateway';
import { middyfy } from '@libs/lambda';

import App from '@Frontend/Pages/index';

import schema from './schema';

const homepage: ValidatedEventAPIGatewayProxyEvent<typeof schema> = async (event) => {
  const context: { url: string } = { url: undefined };
  const html = ReactDOMServer.renderToString(
    <StaticRouter location={event.body.url} context={context}>
      <App />
    </StaticRouter>
  )

  // if (context.url) {
  //   redirect(301, context.url)
  //   return
  // }

  return formatJSONResponse({
    // message: `Hello ${event.body.name}, welcome to the exciting Serverless world!`,
    // event,
    metaTags: `<meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mr N</title>`,
    html,
    css: ''
  });
}

export const main = middyfy(homepage);
