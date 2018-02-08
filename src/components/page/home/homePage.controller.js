// @flow

import axios from 'axios'

// eslint-disable-next-line
const homePage = () => null

function getHomePageData(endpoint: string) {
  return axios.get(endpoint)
    .then((response) => {
      console.log(response)
    })
    .catch((error) => {
      console.log('Error')
      console.log(error)
    })
}

export {
  homePage,
  getHomePageData,
}
