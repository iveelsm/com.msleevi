// @flow

import React from 'react'
import { slide as Menu } from 'react-burger-menu'
import { Link, NavLink } from 'react-router-dom'
import { APP_NAME } from '../constants/config'
import {
  HOME_PAGE_ROUTE,
  BLOG_PAGE_ROUTE,
  CONTACT_PAGE_ROUTE,
  PROJECTS_PAGE_ROUTE,
  CV_PAGE_ROUTE,
} from '../server/routes/routes'


class NavigationBar extends React.Component {
  componentWillMount() {
    this.state = {
      userStatus: '',
    }
  }

  // render() {
  //   return (
  //     <nav className="navbar navbar-toggleable-md navbar-inverse fixed-top bg-inverse">
  //       <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target=".js-navbar-collapse">
  //         <span className="navbar-toggler-icon" />
  //       </button>
  //       <Link to={HOME_PAGE_ROUTE} className="navbar-brand">{APP_NAME}</Link>
  //       <div className="js-navbar-collapse collapse navbar-collapse">
  //         <ul className="navbar-nav mr-auto">
  //           {[
  //                   { route: HOME_PAGE_ROUTE, label: 'Home' },
  //                   { route: BLOG_PAGE_ROUTE, label: 'Blog' },
  //                   { route: CONTACT_PAGE_ROUTE, label: 'Contact' },
  //                   { route: PROJECTS_PAGE_ROUTE, label: 'Projects' },
  //                   { route: CV_PAGE_ROUTE, label: 'CV' },
  //                 ].map(link => (
  //                   <li className="nav-item" key={link.route}>
  //                     <NavLink to={link.route} className="nav-link" activeStyle={{ color: 'white' }} exact onClick={this.props.handleNavLinkClick}>
  //                       {link.label}
  //                     </NavLink>
  //                   </li>
  //                 ))}
  //         </ul>
  //       </div>
  //     </nav>
  //   )
  // }

  render() {
    return (
      <Menu>
        <a id="home" className="menu-item" href="/">Home</a>
        <a id="about" className="menu-item" href="/about">About</a>
        <a id="contact" className="menu-item" href="/contact">Contact</a>
      </Menu>
    )
  }
}

export { NavigationBar as default }
