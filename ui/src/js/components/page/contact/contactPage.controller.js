// @flow

import axios from 'axios'
// import logger from '../../../logging/logger'


function postContactInformation(endpoint: string, information: Object) {
  // logger.info('Posting data!')
  axios.post(endpoint, information)
    .then((response) => {
      // logger.info(response)
    })
    .catch((error) => {
      // logger.error('Error!')
      // logger.error(error)
    })
}

export { postContactInformation as default }
