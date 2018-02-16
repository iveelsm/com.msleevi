// @flow

import axios from 'axios'
import logger from '../../../logging/logger'

// eslint-disable-next-line
const homePage = () => null

function getHomePageData(endpoint: string) {
  return axios.get(endpoint)
    .then((response) => {
      logger.info(response)
    })
    .catch((error) => {
      logger.error('Error')
      logger.error(error)
    })
}

export {
  homePage,
  getHomePageData,
}
