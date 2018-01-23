// @flow

import 'axios'

function postContactInformation(endpoint: string, information: Object) {
	console.log("Posting data!")
	axios.post(endpoint, information)
		.then(function(response) {
			console.log(response)
		})
		.catch(function(error) {
			console.log("Error!")
			console.log(error)
		})
}

export { postContactInformation as default }
