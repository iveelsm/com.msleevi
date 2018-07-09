// @flow

import axios from 'axios'
//import logger from '../../../logging/logger'

function getBlogPosts(endpoint: string) {
  const result = axios.get(endpoint)
    .then((response) => {
      // logger.info(response)
      // logger.info('Format data here?')
      return response.data
    })
    .catch((error) => {
      // logger.error('Error fetching blog post data')
      return ["Errors"]
    })
}

function postBlogPost(post: ?Object, endpoint: string) {
  return endpoint
}

export {
  getBlogPosts,
  postBlogPost,
}
