import React, { Component } from "react";
import decode from "jwt-decode";
import TodoApp from "./Todo";

// Component for user login (main component for now)
class LoginControl extends Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.handleFormSubmit = this.handleFormSubmit.bind(this);
    // this.handleLoginClick = this.handleLoginClick.bind(this);
    this.handleLogoutClick = this.handleLogoutClick.bind(this);
    this.Auth = new AuthService();
    this.state = { isLoggedIn: false };
  }

  // handleLoginClick() {
  //   this.setState({ isLoggedIn: true });
  // }

  handleLogoutClick() {
    this.Auth.logout();
    this.setState({ isLoggedIn: false });
  }

  handleFormSubmit(e) {
    e.preventDefault();

    this.Auth.login(this.state.email, this.state.password)
      .then(res => {
        //this.props.history.replace("/");
        this.setState({ isLoggedIn: true });
      })
      .catch(err => {
        alert(err);
      });
  }

  handleChange(e) {
    this.setState({
      [e.target.name]: e.target.value
    });
  }

  componentWillMount() {
    if (this.Auth.loggedIn()) this.setState({ isLoggedIn: true });
    // if (this.Auth.loggedIn()) this.props.history.replace("/");
  }

  render() {
    const isLoggedIn = this.state.isLoggedIn;
    let button;

    if (isLoggedIn) {
      button = <LogoutButton onClick={this.handleLogoutClick} />;
    }
    // else {
    //   button = <LoginButton onClick={this.handleLoginClick} />;
    // }

    return (
      <div>
        {button}
        <Greeting
          isLoggedIn={isLoggedIn}
          handleChange={this.handleChange}
          handleFormSubmit={this.handleFormSubmit}
        />
      </div>
    );
  }
}

// Service component for user login
class AuthService {
  // Initializing important variables
  constructor(domain) {
    this.domain = domain || "http://localhost:8000/api/users"; // API server domain
    this.fetch = this.fetch.bind(this); // React binding stuff
    this.login = this.login.bind(this);
    // this.getProfile = this.getProfile.bind(this);
  }

  login(email, password) {
    // Get a token from api server using the fetch api
    return this.fetch(`${this.domain}/login`, {
      method: "POST",
      body: JSON.stringify({
        email,
        password
      })
    }).then(res => {
      this.setToken(res.token); // Setting the token in localStorage
      return Promise.resolve(res);
    });
  }

  loggedIn() {
    // Checks if there is a saved token and it's still valid
    const token = this.getToken(); // GEtting token from localstorage
    return !!token && !this.isTokenExpired(token); // handwaiving here
  }

  isTokenExpired(token) {
    try {
      const decoded = decode(token);
      if (decoded.exp < Date.now() / 1000) {
        // Checking if token is expired. N
        return true;
      } else return false;
    } catch (err) {
      return false;
    }
  }

  setToken(idToken) {
    // Saves user token to localStorage
    localStorage.setItem("id_token", idToken);
  }

  getToken() {
    // Retrieves the user token from localStorage
    return localStorage.getItem("id_token");
  }

  logout() {
    // Clear user token and profile data from localStorage
    localStorage.removeItem("id_token");
  }

  // getProfile() {
  //   // Using jwt-decode npm package to decode the token
  //   return decode(this.getToken());
  // }

  fetch(url, options) {
    // performs api calls sending the required authentication headers
    const headers = {
      Accept: "application/json",
      "Content-Type": "application/json"
    };

    // Setting Authorization header
    // Authorization: Bearer xxxxxxx.xxxxxxxx.xxxxxx
    if (this.loggedIn()) {
      headers["Authorization"] = "Bearer " + this.getToken();
    }

    return fetch(url, {
      headers,
      ...options
    })
      .then(this._checkStatus)
      .then(response => response.json());
  }

  _checkStatus(response) {
    // raises an error in case response status is not a success
    if (response.status >= 200 && response.status < 300) {
      // Success status lies between 200 to 300
      return response;
    } else {
      var error = new Error(response.statusText);
      error.response = response;
      throw error;
    }
  }
}

// Logged user greeting - returns Todos container
function UserGreeting(props) {
  return (
    <div>
      <h1>Welcome back!</h1>
      {/* todo: add user name */}
      <TodoApp />
    </div>
  );
}

// Anonymous user greeting - returns login form
function GuestGreeting(props) {
  return (
    <div className="center">
      <div className="card">
        <h1>Login</h1>
        <form>
          <input
            className="form-item"
            placeholder="Email goes here..."
            name="email"
            type="text"
            onChange={props.handleChange}
          />
          <input
            className="form-item"
            placeholder="Password goes here..."
            name="password"
            type="password"
            onChange={props.handleChange}
          />
          <input
            className="form-submit"
            value="SUBMIT"
            type="submit"
            onClick={props.handleFormSubmit}
          />
        </form>
      </div>
    </div>

    // todo: user registration
  );
}

// Greetings for two user roles (logged and anonymous)
function Greeting(props) {
  const isLoggedIn = props.isLoggedIn;
  if (isLoggedIn) {
    return <UserGreeting />;
  }
  return (
    <GuestGreeting
      handleChange={props.handleChange}
      handleFormSubmit={props.handleFormSubmit}
    />
  );
}

// Link for login (reserved, not used)
// function LoginButton(props) {
//   return (
//     <a href="#" onClick={props.onClick}>
//       Login
//     </a>
//   );
// }

// Link for user logout
function LogoutButton(props) {
  return <button onClick={props.onClick}>Logout</button>;
}

export default LoginControl;
