// @flow

import { createLogger, transports } from 'winston'
import { isProd, logLevel, logFile, errorFile } from '../constants/util'

function returnProductionTransports() {
  return ([
    new transports.File({
      filename: logFile,
    }),
    new transports.File({
      level: 'error',
      fileName: errorFile,
    }),
  ])
}

function returnDevTransports() {
  return (
    [
      new transports.Console({
        level: 'debug',
      }),
    ]
  )
}


function generateLogger(status: boolean) {
  const transport = (status)
    ? returnProductionTransports()
    : returnDevTransports()
  return (
    createLogger({
      level: logLevel,
      transports: transport,
    })
  )
}

const logger = generateLogger(isProd)
const consoleLogger = generateLogger(false)

export {
  logger as default,
  consoleLogger,
}
