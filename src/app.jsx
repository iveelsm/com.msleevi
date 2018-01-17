// @flow

import React from 'react'
import Helmet from 'react-helmet'
import { Switch, BrowserRouter } from 'react-router'
import { Route } from 'react-router-dom'
import { APP_NAME } from './constants/config'
import NavigationBar from './components/navigation-bar'
import HomePage from './components/page/home/home'
import NotFoundPage from './components/page/not-found'
import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
  NOT_FOUND_PAGE_ROUTE,
} from './routes/routes'


class App extends React.Component {
    render() {
       console.log("Rendering App")
       return (
            <div style={{ paddingTop: 54 }}>
              <Helmet titleTemplate={`%s | ${APP_NAME}`} defaultTitle={APP_NAME} />
              <Switch>
                <Route path={HOME_PAGE_ROUTE} component={HomePage} />
              </Switch>
            </div>
        )
    }
}

export { App as default }
