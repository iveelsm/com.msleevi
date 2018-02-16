// @flow

import axios from 'axios'
import logger from '../../../logging/logger'

function getBlogPosts(endpoint: string) {
  return axios.get(endpoint)
    .then((response) => {
      logger.info(response)
      logger.info('Format data here?')
    })
    .catch((error) => {
      logger.error(error)
      logger.error('Error fetching blog post data')
    })
}

function postBlogPost(post: ?Object, endpoint: string) {
  return endpoint
}

export {
  getBlogPosts,
  postBlogPost,
}
