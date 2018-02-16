// @flow

import React from 'react'

class Button extends React.Component {
  componentWillMount() {
    this.state = {
      label: '',
    }
  }

  render() {
    return (
      <button
        onClick={this.props.handleClick}
        className="btn btn-primary"
        type="button"
      >
        {this.state.label}
      </button>
    )
  }
}

export default Button
