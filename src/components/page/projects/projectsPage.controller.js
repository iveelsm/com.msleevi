// @flow

import 'axios'

function getProjects(endpoint: string) {
	return axios.get(endpoint)
		.then(function(response) {
			console.log(response)
		})
		.catch(function(error) {
			console.log("Error")
			console.log(error)
		})
}

function postProject(endpoint: string, project: Object) {
	return axios.post(endpoint, project)
		.then(function(response) {
			console.log(response)
		})
		.catch(function(error) {
			console.log("Error")
			console.log(error)
		})
}
