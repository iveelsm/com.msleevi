// @flow

import React from 'react'

var Button = React.createClass({
    render: function() {
        return (
            <button onClick={this.props.handleClick}
                className="btn btn-primary"
                type="button">
                {this.props.label}
            </button>
        );
    }
})

export default Button
