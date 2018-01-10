// @ flow

import React from 'react'

type Props = {
  posts: list,
}

const BlogList = ({ posts }: Props) =>
  <ul>
    <li>{posts}</li>
  </ul>

export default BlogList
