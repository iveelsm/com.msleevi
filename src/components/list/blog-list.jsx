// @ flow

import React from 'react'

class BlogList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            posts : [],
            numberOfComments : []
        };
    }

    componentWillMount() {
        console.log("Include fetchPosts shit here")
    }

    render() {
        return (
            <ul>
                <li>{this.props.list}</li>
            </ul>
        );
    }
}

export default BlogList
