// @flow

import React from 'react'
import Helmet from 'react-helmet'

import BlogList from '../../container/blog-list'

const title = 'Blog'

class BlogPage extends React.Component {
    render() {
        return (
            <div>
              <Helmet
                title={title}
                meta={[
                  { name: 'description', content: 'Blog page' },
                  { property: 'og:title', content: this.props.title },
                ]}
              />
              <div className="row">
                <div className="col-12">
                  <h1>{this.props.title}</h1>
                  <BlogList />
                </div>
              </div>
            </div>
        );
    }
}
