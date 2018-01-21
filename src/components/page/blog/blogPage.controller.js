// @flow

import 'axios'

function getBlogPosts(endpoint: string) {
	return axios.get(endpoint)
		.then(function(response) {
			console.log(response)
			console.log("Format data here?")
		})
		.catch(function(error) {
			console.log(error)
			console.log("Error fetching blog post data")
		})
}

function postBlogPosts(endpoint: string, posts: Object) {

}

function
