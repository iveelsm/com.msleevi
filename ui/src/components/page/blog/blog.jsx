// @flow

import React from 'react'
import Helmet from 'react-helmet'

import { HTTP, API_LOCATION, API_PORT } from '../../../constants/config'
import BlogList from './blogList'
import { getBlogPosts } from './blogPage.controller'
import logger from '../../../logging/logger';

const title = 'Blog'

class BlogPage extends React.Component {
  componentWillMount() {
    this.state = {
      posts : [],
    }
  }

  componentDidMount() {
    console.info("Did mount")
    getBlogPosts(HTTP + API_LOCATION + ":" + API_PORT)
        .then(result => {
          this.setState({posts : results.data})
        })
    logger.info("Data: " + this.state.posts)
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
            <BlogList posts={this.state.posts} />
          </div>
        </div>
      </div>
    )
  }
}

export { BlogPage as default }
