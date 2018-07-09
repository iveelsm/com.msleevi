// @flow

import React from 'react'


class Post extends React.Component {
  componentWillMount() {
    this.state = {
      message: '',
    }
  }

  render() {
    return (
      <div onClick={this.props.handleClick} className="post">
        {this.state.message}
      </div>
    )
  }
}

export { Post as default }
