// @ flow

import React from 'react'

class Message extends React.Component {
    render: function () {
        return(
            <p>
                {this.props.message}
            </p>
        );
    }
}

export default Message
