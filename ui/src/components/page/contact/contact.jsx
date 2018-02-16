// @flow

import React from 'react'

class ContactPage extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      name: '',
      email: '',
      message: '',
    }
    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
  }

  handleChange(event: Object) {
    this.setState({
      [event.target.name]: event.target.value,
    })
  }

  handleSubmit() {
    this.state = {
      name: '',
      email: '',
      message: '',
    }
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label htmlFor="name">
          Name:
          <input name="name" type="text" value={this.state.name} onChange={this.handleChange} />
        </label>
        <label htmlFor="email">
          Email:
          <input name="email" type="text" value={this.state.email} onChange={this.handleChange} />
        </label>
        <label htmlFor="message">
          Message:
          <input name="message" type="text" value={this.state.message} onChange={this.handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    )
  }
}

export { ContactPage as default }
