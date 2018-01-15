import React from 'react'

class HoverButton extends React.Component {
  constructor(props) {
      super(props)
      this.state.hover = false
  }

  mouseOver: function () {
      this.setState({hover: true});
  },

  mouseOut: function () {
      this.setState({hover: false});
  },

  render: function() {
      var label = "foo";
      if (this.state.hover) {
          label = "bar";
      }
      return React.createElement(
          "button",
          {onMouseOver: this.mouseOver, onMouseOut: this.mouseOut},
          label
      );
  }
}

export default HoverButton
