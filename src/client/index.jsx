// @flow

import 'babel-polyfill'

import React from 'react'
import ReactDOM from 'react-dom'
import { AppContainer } from 'react-hot-loader'
import { BrowserRouter } from 'react-router-dom'

import App from '../components/app'
import { APP_CONTAINER_SELECTOR } from '../constants/config'

const rootEl = document.querySelector(APP_CONTAINER_SELECTOR)

function wrapApp(AppComponent) {
  console.log('Attempting to wrapApp')
  return (
    <BrowserRouter>
      <AppContainer>
        <AppComponent />
      </AppContainer>
    </BrowserRouter>
  )
}

ReactDOM.render(wrapApp(App), rootEl)

if (module.hot) {
  // flow-disable-next-line
  module.hot.accept('../components/app', () => {
    // eslint-disable-next-line global-require
    const NextApp = require('../components/app').default
    ReactDOM.render(wrapApp(NextApp), rootEl)
  })
}
