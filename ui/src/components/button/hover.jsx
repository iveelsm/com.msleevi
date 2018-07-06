import React from 'react'
import Button from './button';

class HoverButton extends Button {
  constructor(props) {
    super(props)
    this.state.hover = false
    this.state.hoverLabel = props.hoverLabel
  }

  mouseOver() {
    this.setState({ hover: true })
  }

  mouseOut() {
    this.setState({ hover: false })
  }

  render() {
    let label = this.state.label

    if (this.state.hover) {
      label = this.state.hoverLabel
    }

    return (
      <button onMouseOver={this.mouseOver} onMouseOut={this.mouseOut}>
        {label}
      </button>
    )
  }
}

export { HoverButton as default }
