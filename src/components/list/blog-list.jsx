// @ flow

import React from 'react'

class BlogList extends React.Component {
    render() {
        return (
            <ul>
                <li>{this.props.list}</li>
            </ul>
        );
    }
}

export default BlogList
