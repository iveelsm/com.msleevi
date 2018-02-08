// @flow

import axios from 'axios'

function getProjects(endpoint: string) {
  return axios.get(endpoint)
    .then((response) => {
      console.log(response)
    })
    .catch((error) => {
      console.log('Error')
      console.log(error)
    })
}

function postProject(endpoint: string, project: Object) {
  return axios.post(endpoint, project)
    .then((response) => {
      console.log(response)
    })
    .catch((error) => {
      console.log('Error')
      console.log(error)
    })
}

export {
  getProjects,
  postProject,
}
