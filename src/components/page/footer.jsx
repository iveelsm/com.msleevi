// @flow

import React from 'react'
import { APP_NAME } from '../../constants/config'


class Footer extends React.Component {
  componentWillMount() {
    this.state = {
      links: [],
    }
  }

  render() {
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
