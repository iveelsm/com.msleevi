// @flow

import React from 'react'
import Helmet from 'react-helmet'
import { Switch, BrowserRouter } from 'react-router'
import { Route } from 'react-router-dom'
import { APP_NAME } from './constants/config'
import NavigationBar from './components/navigation-bar'

import ProjectsPage from './components/page/projects/projects'
import ContactPage from './components/page/contact/contact'
import HomePage from './components/page/home/home'
import AdminPage from './components/page/admin/admin'
import BlogPage from './components/page/blog/blog'
import CVPage from './components/page/cv/CV'
import Footer from './components/page/footer'

import NotFoundPage from './components/page/not-found'
import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
  CONTACT_PAGE_ROUTE,
  PROJECTS_PAGE_ROUTE,
  CV_PAGE_ROUTE,
  ADMIN_PAGE_ROUTE,
  NOT_FOUND_PAGE_ROUTE,
} from './routes/routes'


class App extends React.Component {
    render() {
       console.log("Rendering App")
       return (
            <div style={{ paddingTop: 54 }}>
              <Helmet titleTemplate={`%s | ${APP_NAME}`} defaultTitle={APP_NAME} />
              <NavigationBar />
              <Switch>
                  <Route exact path={HOME_PAGE_ROUTE} render={() => <HomePage />} />
                  <Route path={BLOG_PAGE_ROUTE} render={() => <BlogPage />} />
                  <Route path={CONTACT_PAGE_ROUTE} render={() => <ContactPage />} />
                  <Route path={PROJECTS_PAGE_ROUTE} render={() => <ProjectsPage />} />
                  <Route path={CV_PAGE_ROUTE} render={() => <CVPage />} />
                  <Route path={ADMIN_PAGE_ROUTE} render={() => <AdminPage />} />
                  <Route component={NotFoundPage} />
              </Switch>
              <Footer />
            </div>
        )
    }
}

export { App as default }
