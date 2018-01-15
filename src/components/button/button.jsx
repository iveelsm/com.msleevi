// @flow

import React from 'react'

class Button extends React.Component {
    render() {
        return (
            <button onClick={this.props.handleClick}
                className="btn btn-primary"
                type="button">
                {this.props.label}
            </button>
        );
    }
}

export default Button
