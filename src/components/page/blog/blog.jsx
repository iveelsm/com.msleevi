// @flow

import React from 'react'
import Helmet from 'react-helmet'

import BlogList from '../../container/blog-list'

const title = 'Blog'

const BlogPage = () => (
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
