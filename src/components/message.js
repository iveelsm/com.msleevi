// @ flow

import React from 'react'

var Message = React.createClass({
    render: function () {
        return(
            <p>
                {this.props.message}
            </p>
        );
    }
})

export default Message
