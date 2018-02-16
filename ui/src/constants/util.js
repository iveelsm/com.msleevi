// @flow

export const isProd = process.env.NODE_END === 'production'
export const logLevel = (process.env.LOG_LEVEL)
  ? process.env.LOG_LEVEL
  : 'info'
export const logFile = '/opt/com/msleevi/ui.log'
export const errorFile = '/opt/com/msleevi/ui.error.log'
