// @ flow

import React from 'react'

class BlogList extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      posts: [],
    }
  }

  componentWillMount() {
    console.log('Include fetchPosts shit here')
  }

  render() {
    return (
      <ul>
        <li>{this.state.posts}</li>
      </ul>
    )
  }
}

export default BlogList
