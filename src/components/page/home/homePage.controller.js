// @flow

import 'axios'

// eslint-disable-next-line
export const homePage = () => null

function getHomePageData(endpoint: string) {
	return axios.get(endpoint)
		.then(function(response) {
			console.log(response)
		})
		.catch(function(error) {
			console.log("Error")
			console.log(error)
		}
}
