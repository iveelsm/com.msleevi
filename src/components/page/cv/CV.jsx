// @flow

import React from 'react'


class CVPage extends React.Component {
  componentWillMount() {
    this.state = {
      CV: '',
    }
  }

  render() {
    return (
      <div>
        <h3>{this.state.CV}</h3>
      </div>
    )
  }
}

export { CVPage as default }
