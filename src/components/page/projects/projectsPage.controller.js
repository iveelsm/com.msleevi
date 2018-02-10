// @flow

import axios from 'axios'
import logger from '../../../logging/logger'


function getProjects(endpoint: string) {
  return axios.get(endpoint)
    .then((response) => {
      logger.info(response)
    })
    .catch((error) => {
      logger.error('Error')
      logger.error(error)
    })
}

function postProject(endpoint: string, project: Object) {
  return axios.post(endpoint, project)
    .then((response) => {
      logger.info(response)
    })
    .catch((error) => {
      logger.error('Error')
      logger.error(error)
    })
}

export {
  getProjects,
  postProject,
}
