// Component for user login (main component for now)
class LoginControl extends React.Component {
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
function LoginButton(props) {
  return (
    <a href="#" onClick={props.onClick}>
      Login
    </a>
  );
}

// Link for user logout
function LogoutButton(props) {
  return (
    <a href="#" onClick={props.onClick}>
      Logout
    </a>
  );
}

//-------------------------------------------------------------------------------------------------------

// Todos page title
const Title = ({ todoCount }) => {
  return (
    <div>
      <div>
        <h1>To-do&apos;s count: {todoCount ? todoCount.length : 0}</h1>
      </div>
    </div>
  );
};

// Form for adding todo
class TodoForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      description: "",
      state: "1"
    };
  }

  handleChange = e => {
    this.setState({
      [e.target.name]: e.target.value
    });
  };

  onSubmit = e => {
    e.preventDefault();
    const form = {
      name: this.state.name,
      description: this.state.description,
      state: this.state.state
    };
    this.props.addTodo(form);
    this.setState({
      name: "",
      description: "",
      state: "1"
    });
  };

  render() {
    return (
      <div>
        <form>
          <label>
            Name:
            <input
              className="form-control col-md-12"
              name="name"
              value={this.state.name}
              onChange={e => this.handleChange(e)}
            />
          </label>
          <label>
            Description:
            <input
              className="form-control col-md-12"
              name="description"
              value={this.state.description}
              onChange={e => this.handleChange(e)}
            />
          </label>
          <select
            name="state"
            value={this.state.state}
            onChange={e => this.handleChange(e)}
          >
            <option value="1">created</option>
            <option value="2">wait</option>
            <option value="3">canceled</option>
            <option value="4">blocked</option>
            <option value="5">in process/doing</option>
            <option value="6">review</option>
            <option value="7">done</option>
            <option value="8">archived</option>
          </select>
          <button onClick={e => this.onSubmit(e)}>Submit</button>
        </form>
      </div>
    );
  }
}

// Changes static text label to input
class EditableLabel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: props.todo.id,
      name: props.todo.name,
      description: props.todo.description,
      type: props.type,
      editing: false
    };
    this.initEditor();
    this.edit = this.edit.bind(this);
    this.save = this.save.bind(this);
  }

  initEditor() {
    this.editor = (
      <input
        type="text"
        name={this.state.type}
        defaultValue={this.state[this.state.type]}
        onKeyPress={event => {
          const key = event.which || event.keyCode;
          if (key === 13) {
            //enter key
            this.save(event);
          }
        }}
        autoFocus={true}
      />
    );
  }

  edit(e) {
    this.setState({
      [e.target.name]: e.target.value,
      editing: true
    });
  }

  save(e) {
    const form = {
      id: this.state.id,
      name: e.target.name,
      value: e.target.value
    };
    this.props.update(form);
    this.setState({
      [e.target.name]: e.target.value,
      editing: false
    });
  }

  componentDidUpdate() {
    this.initEditor();
  }

  render() {
    return this.state.editing ? (
      this.editor
    ) : (
      <p name={this.state.type} onClick={this.edit}>
        {this.state[this.state.type]}
      </p>
    );
  }
}

// Each Todo component
const Todo = ({ todo, remove, update }) => {
  return (
    <div>
      <EditableLabel todo={todo} type="name" update={update} />
      <EditableLabel todo={todo} type="description" update={update} />
      <button
        onClick={() => {
          remove(todo.id);
        }}
      >
        Delete
      </button>
    </div>
  );
};

// All todo components or nothing
const TodoList = ({ todos, remove, update }) => {
  // Map through the todos
  const todoNode = todos
    ? todos.map(todo => {
        return (
          <Todo todo={todo} key={todo.id} remove={remove} update={update} />
        );
      })
    : "There is no to-do's";
  return (
    <div className="list-group" style={{ marginTop: "30px" }}>
      {todoNode}
    </div>
  );
};

// Container component for Todos
class TodoApp extends React.Component {
  constructor(props) {
    // Pass props to parent class
    super(props);
    // Set initial state
    this.state = {
      data: []
    };
    this.apiUrl = "http://127.0.0.1:8000/api/todos";
  }

  // Lifecycle method
  componentDidMount() {
    // Make HTTP request with Axios
    axios.get(this.apiUrl).then(res => {
      // Set state with result
      this.setState({ data: res.data || [] });
    });
  }

  // Add todo handler
  addTodo(val) {
    // Update data
    axios.post(this.apiUrl, val).then(res => {
      this.state.data.push(res.data);
      this.setState({ data: this.state.data });
    });
  }

  // Handle update
  updateTodo(val) {
    const form = {
      [val.name]: val.value
    };
    axios.patch(this.apiUrl + "/" + val.id, form);
  }

  // Handle remove
  removeTodo(id) {
    // Filter all todos except the one to be removed
    const remainder = this.state.data.filter(todo => {
      if (todo.id !== id) return todo;
    });
    // Update state with filter
    axios.delete(this.apiUrl + "/" + id).then(res => {
      this.setState({ data: remainder });
    });
  }

  render() {
    return (
      <div>
        <Title todoCount={this.state.data} />
        <TodoForm addTodo={this.addTodo.bind(this)} />
        <TodoList
          todos={this.state.data}
          remove={this.removeTodo.bind(this)}
          update={this.updateTodo.bind(this)}
        />
      </div>
    );
  }
}

ReactDOM.render(<LoginControl />, document.getElementById("root"));
