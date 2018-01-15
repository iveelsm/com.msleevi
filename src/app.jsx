// @flow

import React from 'react'
import Helmet from 'react-helmet'
import { Switch } from 'react-router'
import { Route } from 'react-router-dom'
import { APP_NAME } from './constants/config'
import Nav from './component/nav'
import HomePage from './component/page/home'
import BlogPage from './component/page/blog'
import NotFoundPage from './component/page/not-found'
import {
  HOME_PAGE_ROUTE,
} from './routes'


const App = () => (
  <div style={{ paddingTop: 54 }}>
    <Helmet titleTemplate={`%s | ${APP_NAME}`} defaultTitle={APP_NAME} />
    <Nav />
    <Switch>
      <Route exact path={HOME_PAGE_ROUTE} render={() => <HomePage />} />
      <Route path={BLOG_PAGE} render={() => <BlogPage />} />
      <Route component={NotFoundPage} />
    </Switch>
  </div>
)


export default App
