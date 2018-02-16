import React from 'react'

class ProjectList extends React.Component {
  componentWillMount() {
    this.state = {
      projects: [],
    }
  }

  render() {
    return (
      <div>
        <h3>{this.state.projects}</h3>
      </div>
    )
  }
}

export { ProjectList as default }
