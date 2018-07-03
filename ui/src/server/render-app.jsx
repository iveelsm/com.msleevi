// @flow

import React from 'react'
import Helmet from 'react-helmet'
import ReactDOMServer from 'react-dom/server'
import { StaticRouter } from 'react-router'

import App from '../components/app'
import { APP_CONTAINER_CLASS, STATIC_PATH, WDS_PORT } from '../constants/config'
import { isProd } from '../constants/util'


const renderApp = (location: string, routerContext: ?Object = {}) => {
  // eslint-disable-next-line
  const appHtml = ReactDOMServer.renderToString(
    <StaticRouter location={location} context={routerContext}>
      <App />
    </StaticRouter>)
  const head = Helmet.rewind()

  return (
    `<!doctype html>
    <html>
      <head>
        ${head.title}
        ${head.meta}
        <link rel="stylesheet" href="${STATIC_PATH}/css/public.css">
        <link rel="stylesheet" href="${STATIC_PATH}/css/home.css">        
      </head>
      <body>
        <div class="${APP_CONTAINER_CLASS}">${appHtml}</div>
        <script src="${isProd ? STATIC_PATH : `http://localhost:${WDS_PORT}/dist`}/js/bundle.js"></script>
      </body>
    </html>`
  )
}

export { renderApp as default }
