// @flow

import React from 'react'

type Props = {
  message: string,
  handleClick: Function,
}

const Post = ( {message, handleClick }: Props) =>
  <div onClick={handleClick} className="post">{message}</div>

export default Post
