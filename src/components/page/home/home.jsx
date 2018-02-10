// @flow

import React from 'react'
import Helmet from 'react-helmet'

import { APP_NAME } from '../../../constants/config'

const styles = {
  hoverMe: {
    '&:hover': {
      color: 'red',
    },
  },
  '@media (max-width: 800px)': {
    resizeMe: {
      color: 'red',
    },
  },
  specialButton: {
    composes: ['btn', 'btn-primary'],
    backgroundColor: 'limegreen',
  },
}

class HomePage extends React.Component {
  componentWillMount() {
    this.state = {
      name: { name: 'description', content: 'Michael Sleevi\'s Home Page' },
      properties: { property: 'og:title', content: APP_NAME },
      style: styles,
    }
  }

  render() {
    return (
      <div>
        <Helmet
          meta={[
                  {
                    name: this.state.name.description,
                    content: this.state.name.content,
                  },
                  {
                    property: this.state.properties.property,
                    content: this.state.properties.content,
                  },
                ]}
        />
        <div className="jumbotron">
          <div className="container">
            <h1 className="display-3 mb-4">{APP_NAME}</h1>
          </div>
        </div>
      </div>
    )
  }
}

export { HomePage as default }
