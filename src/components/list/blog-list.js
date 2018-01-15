// @ flow

import React from 'react'

var BlogList = React.createClass({
    render: function () {
        return (
            <ul>
                <li>{this.props.list}</li>
            </ul>
        );
    }
})

export default BlogList
