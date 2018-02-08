// @ flow

import React from 'react'

class Message extends React.Component {
  componentWillMount() {
    this.state = {
      message: '',
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

export default Message
