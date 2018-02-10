// @flow

import React from 'react'
import { APP_NAME } from '../../constants/config'
import logger from '../../logging/logger'


class Footer extends React.Component {
  componentWillMount() {
    this.state = {
      links: [],
    }
  }

  render() {
    logger.info('Footer')
    return (
      <div className="container mt-5">
        <hr />
        <footer>
          <p>© {APP_NAME} 2018</p>
        </footer>
      </div>
    )
  }
}

export { Footer as default }
