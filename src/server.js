// @flow

import compression from 'compression'
import express from 'express'
import { Server } from 'http'
import socketIO from 'socket.io'

import routing from './utils/routing'
import { WEB_PORT, STATIC_PATH } from './constants/config'
import { isProd } from './constants/util'

const app = express()

// // flow-disable-next-line
const http = Server(app)
// const io = socketIO(http)
// setUpSocket(io)

app.use(compression())
app.use(STATIC_PATH, express.static('dist'))
app.use(STATIC_PATH, express.static('public'))

routing(app)

http.listen(WEB_PORT, () => {
  // eslint-disable-next-line no-console
  console.log(`Server running on port ${WEB_PORT}
    ${isProd
    ? '(production)'
    : '(development).\n Keep "yarn dev:wds" running in another terminal'}`)
})
