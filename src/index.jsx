
// @flow

import 'babel-polyfill'

import Tether from 'tether'
import Immutable from 'immutable'
import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import { AppContainer } from 'react-hot-loader'

import App from './app'
import { APP_CONTAINER_SELECTOR, JSS_SSR_SELECTOR } from './constants/config'
import { isProd } from './constants/util'

require('bootstrap')

const rootEl = document.querySelector(APP_CONTAINER_SELECTOR)

const wrapApp = (AppComponent) => (
    <BrowserRouter>
      <AppContainer>
        <AppComponent />
      </AppContainer>
    </BrowserRouter>
)

ReactDOM.render(wrapApp(App), rootEl)
