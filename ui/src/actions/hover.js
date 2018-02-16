import React from 'react'

class HoverButton extends React.Component {
  constructor(props) {
    super(props)
    this.state.hover = false
  }

  mouseOver() {
    this.setState({ hover: true })
  }

  mouseOut() {
    this.setState({ hover: false })
  }

  render() {
    let label = 'foo'
    if (this.state.hover) {
      label = 'bar'
    }
    return React.createElement(
      'button',
      { onMouseOver: this.mouseOver, onMouseOut: this.mouseOut },
      label,
    )
  }
}

export default HoverButton
