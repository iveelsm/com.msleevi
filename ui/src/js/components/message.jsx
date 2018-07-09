// @ flow

import React from 'react'

class Message extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      message: props.message || ''
    }
  }

  render() {
    return (
      <p>
        {this.state.message}
      </p>
    )
  }
}

export { Message as default }
