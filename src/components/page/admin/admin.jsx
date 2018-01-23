// @flow

import React from 'react'

class AdminPage extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			username: "",
			password: "",
		}
	}

	render() {
		return (
      <form>
  			<h3>Please enter credentials</h3>
  			<label>
  				Username:
  				<input name="username" type="text" value={this.state.username} onChange={this.handleChange} />
  			</label>
  			<label>
  				Password:
  				<input name="password" type="text" value={this.state.password} onChange={this.handleChange} />
  			</label>
		</form>
    )
	}
}

export { AdminPage as default }
