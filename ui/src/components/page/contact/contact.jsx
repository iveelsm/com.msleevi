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

  handleSubmit(event: Object) {

  }

  render() {
    return (
      <form class="contact" onSubmit={this.handleSubmit}>
        <label htmlFor="name">
          Name:
        </label>
        <input name="name" type="text" value={this.state.name} onChange={this.handleChange} />
        <label htmlFor="email">
          Email:
        </label>
        <input name="email" type="email" value={this.state.email} onChange={this.handleChange} />
        <label htmlFor="message">
          Message:
        </label>
        <textarea name="message" class="message-box"/>
        <input type="submit" value="Submit" onClick={this.handleSubmit} onChange={this.handleChange} />
      </form>
    )
  }
}

export { ContactPage as default }
