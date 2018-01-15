// @flow

import React from 'react'


var Post = React.createClass({
    render: function() {
        return (
            <div onClick={this.props.handleClick} className="post">
                {this.props.message}
            </div>
        );
    }
})

export default Post
