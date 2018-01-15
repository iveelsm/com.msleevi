// @flow

import React from 'react'
import Helmet from 'react-helmet'
import { Switch } from 'react-router'
import { Route } from 'react-router-dom'
import { APP_NAME } from './constants/config'
import Nav from './components/nav'
import HomePage from './components/page/home/home'
import BlogPage from './components/page/blog/blog'
import NotFoundPage from './components/page/not-found'
import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
} from './routes/routes'


class App extends React.Component {
    render() {
       return (
            <div style={{ paddingTop: 54 }}>
              <Helmet titleTemplate={`%s | ${APP_NAME}`} defaultTitle={APP_NAME} />
              <Nav />
              <Switch>
                <Route path={HOME_PAGE_ROUTE} component={HomePage} />
              </Switch>
            </div>
        )
    }
}

export default App
