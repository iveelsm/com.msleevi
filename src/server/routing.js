// @flow

import renderApp from './render-app'
import logger from '../logging/logger'


import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
  CV_PAGE_ROUTE,
  PROJECTS_PAGE_ROUTE,
  CONTACT_PAGE_ROUTE,
  ADMIN_PAGE_ROUTE,
} from './routes/routes'


function routing(app: Object) {
  app.get(HOME_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url))
  })

  app.get(CONTACT_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url))
  })

  app.get(CV_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url))
  })

  app.get(PROJECTS_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url))
  })

  app.get(BLOG_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url))
  })

  app.get(ADMIN_PAGE_ROUTE, (req, res) => {
    res.send(renderApp(req.url))
  })

  app.get('*', (req, res) => {
    res.status(404).send(renderApp(req.url))
  })

  // eslint-disable-next-line no-unused-vars
  app.use((err, req, res, next) => {
    logger.error(err.stack)
    res.status(500).send('Something went wrong!')
  })
}

export default routing
