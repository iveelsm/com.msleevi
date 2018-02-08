// @flow

import axios from 'axios'

function postContactInformation(endpoint: string, information: Object) {
  console.log('Posting data!')
  axios.post(endpoint, information)
    .then((response) => {
      console.log(response)
    })
    .catch((error) => {
      console.log('Error!')
      console.log(error)
    })
}

export { postContactInformation as default }
