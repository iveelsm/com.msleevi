// @ flow

import React from 'react'

class BlogList extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      posts: props.posts,
    }
  }

  render() {
    return (
      <ul>
        { this.state.posts.map(post => <li>{post}</li>)}
      </ul>
    )
  }
}

export default BlogList
