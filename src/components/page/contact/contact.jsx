// @flow

import React from 'react'
import './contactPage.controller'


class ContactPage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			name: "",
			email: "",
			message: "",
		}
		this.handleChange = this.handleChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}

	handleChange(event: Object) {
	  this.setState({
			[event.target.name] : event.target.value
		});
		console.log(this.state)
	}

	handleSubmit() {
		postContactInformation(this.state)
		this.state = {
			name: "",
			email: "",
			message: "",
		}
	}

	render() {
		console.log("Rendering Contact!")
		return (
			<form onSubmit={this.handleSubmit}>
        <label>
					Name:
          <input name="name" type="text" value={this.state.value} onChange={this.handleChange} />
        </label>
			  <label>
					Email:
					<input name="email" type="text" value={this.state.value} onChange={this.handleChange} />
			  </label>
			  <label>
					Message:
					<input name="message" type="text" value={this.state.value} onChange={this.handleChange} />
			  </label>
	      <input type="submit" value="Submit" />
      </form>
		)
	}
}

export { ContactPage as default }
