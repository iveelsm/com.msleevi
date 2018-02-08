// @flow

import React from 'react'
import Helmet from 'react-helmet'
import { Link } from 'react-router-dom'
import { HOME_PAGE_ROUTE } from '../../server/routes/routes'

const title = 'Page Not Found!'

class NotFoundPage extends React.Component {
  componentWillMount() {
    this.state = {
      notFound: '',
    }
  }

  render() {
    return (
      <div className="container mt-4">
        <Helmet title={title} />
        <div className="row">
          <div className="col-12">
            <h1>{title}</h1>
            <div><Link to={HOME_PAGE_ROUTE}>Go to the homepage</Link>.</div>
          </div>
        </div>
      </div>
    )
  }
}

export { NotFoundPage as default }
