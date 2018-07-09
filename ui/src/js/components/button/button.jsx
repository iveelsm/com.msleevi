// @flow

import React from 'react'

class Button extends React.Component {
  constructor(props) {
    super(props)
    this.state.label = props.label || ''
  }

  render() {
    return (
      <button onClick={this.props.handleClick} className="btn btn-primary" type="button">
        {this.state.label}
      </button>
    )
  }
}

export default Button
