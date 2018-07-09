// @flow

import React from 'react'

class ProjectPost extends React.Component {
  componentWillMount() {
    this.state = {
      post: '',
    }
  }

  render() {
    return (
      <div>
        <h3>{this.state.post}</h3>
      </div>
    )
  }
}

export { ProjectPost as default }
