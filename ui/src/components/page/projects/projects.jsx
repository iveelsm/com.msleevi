// @flow

import React from 'react'

class ProjectsPage extends React.Component {
  componentWillMount() {
    this.state = {
      projects: [],
    }
  }

  render() {
    return (
      <div>
        <h2>{this.state.projects}</h2>
      </div>
    )
  }
}

export { ProjectsPage as default }
