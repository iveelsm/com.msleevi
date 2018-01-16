// @flow

import 'babel-polyfill'
import React from 'react'
import ReactDOM from 'react-dom'
import ReactDOMServer from 'react-dom/server'

import App from './app'

import compression from 'compression'
import express from 'express'
import { Server } from 'http'

import { WEB_PORT, STATIC_PATH } from './constants/config'
import { isProd } from './constants/util'

import {
  HOME_PAGE_ROUTE,
} from './routes/routes'

const app = express()
const http = Server(app)

app.use(compression())
app.use(STATIC_PATH, express.static('dist'))
app.use(STATIC_PATH, express.static('public'))

app.get(HOME_PAGE_ROUTE, (req, res) => res.send(renderApp("foo")))

app.listen(3000, () => console.log('Example app listening on port 3000!'))

app.get('/500', () => {
  throw Error('Fake Internal Server Error')
})

// eslint-disable-next-line no-unused-vars
app.use((err, req, res, next) => {
  // eslint-disable-next-line no-console
  console.error(err.stack)
  res.status(500).send('Something went wrong!')
})

http.listen(WEB_PORT, () => {
  // eslint-disable-next-line no-console
  console.log(`Server running on port ${WEB_PORT}
    ${isProd
    ? '(production)'
    : '(development).\n Keep "yarn dev:wds" running in another terminal'}`)
})

function renderApp(title: string) {
 return (
         `<!doctype html>
        <html>
          <head>
            <title>${title}</title>
            <link rel="stylesheet" href="${STATIC_PATH}/css/style.css">
          </head>
          <body>
            <h1>${title}</h1>
          </body>
        </html>`
    )
}
