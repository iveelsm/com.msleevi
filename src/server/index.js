/* @flow */

import 'babel-polyfill'
import compression from 'compression'
import express from 'express'
import { Server } from 'http'

import routing from './routing'
import { WEB_PORT, STATIC_PATH } from '../constants/config'
import { isProd } from '../constants/util'
import { consoleLogger } from '../logging/logger'

const app = express()
const http = Server(app)

app.use(compression())
app.use(STATIC_PATH, express.static('dist'))
app.use(STATIC_PATH, express.static('public'))

routing(app)

http.listen(WEB_PORT, () => {
  const logString = `Server running on port ${WEB_PORT}
    ${isProd
    ? '(production)'
    : '(development).\n Keep "yarn dev:wds" running in another terminal'}`
  consoleLogger.info(logString)
})
