// @flow

import React from 'react'
import { Link, NavLink } from 'react-router-dom'
import { APP_NAME } from '../constants/config'
import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
  NOT_FOUND_PAGE_ROUTE,
} from '../routes/routes'


class NavigationBar extends React.Component {
    render() {
        console.log("Rendering Nav")
        return (
            <nav className="navbar navbar-toggleable-md navbar-inverse fixed-top bg-inverse">
              <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target=".js-navbar-collapse">
                <span className="navbar-toggler-icon" />
              </button>
              <Link to={HOME_PAGE_ROUTE} className="navbar-brand">{APP_NAME}</Link>
              <div className="js-navbar-collapse collapse navbar-collapse">
                <ul className="navbar-nav mr-auto">
                  {[
                    { route: HOME_PAGE_ROUTE, label: 'Home' },
                    { route: BLOG_PAGE_ROUTE, label: 'Blog'},
                    { route: NOT_FOUND_PAGE_ROUTE, label: '404 Demo' },
                  ].map(link => (
                    <li className="nav-item" key={link.route}>
                      <NavLink to={link.route} className="nav-link" activeStyle={{ color: 'white' }} exact onClick={this.props.handleNavLinkClick}>
                          {link.label}
                      </NavLink>
                    </li>
                  ))}
                </ul>
              </div>
            </nav>
        )
    }
}

export { NavigationBar as default }
