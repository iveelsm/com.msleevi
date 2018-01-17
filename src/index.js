/* @flow */

import 'babel-polyfill'
import React from 'react'
import ReactDOM from 'react-dom'
import ReactDOMServer from 'react-dom/server'

import App from './app'
import renderApp from './render-app'

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

app.get(HOME_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url, null))
})

app.use((err, req, res, next) => {
  console.error(err.stack)
  res.status(500).send('Something went wrong!')
})

http.listen(WEB_PORT, () => {
  console.log(`Server running on port ${WEB_PORT}
    ${isProd
    ? '(production)'
    : '(development).\n Keep "yarn dev:wds" running in another terminal'}`)
})
