// @flow

import React from 'react'


class Post extends React.Component {
    render() {
        return (
            <div onClick={this.props.handleClick} className="post">
                {this.props.message}
            </div>
        );
    }
}

export default Post
