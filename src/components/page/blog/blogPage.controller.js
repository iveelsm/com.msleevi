// @flow

import axios from 'axios'

function getBlogPosts(endpoint: string) {
  return axios.get(endpoint)
    .then((response) => {
      console.log(response)
      console.log('Format data here?')
    })
    .catch((error) => {
      console.log(error)
      console.log('Error fetching blog post data')
    })
}

function postBlogPost(post: ?Object, endpoint: string) {
  return endpoint
}

export {
  getBlogPosts,
  postBlogPost,
}
