// @flow

import React from 'react'
import Helmet from 'react-helmet'

import BlogList from '../../list/blog-list'

const title = 'Blog'

class BlogPage extends React.Component {
  componentWillMount() {
    this.state = {
      posts: [],
      comments: [],
    }
  }

  render() {
    return (
      <div>
        <Helmet
          title={title}
          meta={[
                  { name: 'description', content: 'Blog page' },
                  { property: 'og:title', content: title },
                ]}
        />
        <div className="row">
          <div className="col-12">
            <h1>{title}</h1>
            <BlogList />
          </div>
        </div>
      </div>
    )
  }
}

export { BlogPage as default }
