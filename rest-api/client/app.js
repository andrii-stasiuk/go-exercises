var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

// Component for user login (main component for now)
var LoginControl = function (_React$Component) {
  _inherits(LoginControl, _React$Component);

  function LoginControl(props) {
    _classCallCheck(this, LoginControl);

    var _this = _possibleConstructorReturn(this, (LoginControl.__proto__ || Object.getPrototypeOf(LoginControl)).call(this, props));

    _this.handleChange = _this.handleChange.bind(_this);
    _this.handleFormSubmit = _this.handleFormSubmit.bind(_this);
    // this.handleLoginClick = this.handleLoginClick.bind(this);
    _this.handleLogoutClick = _this.handleLogoutClick.bind(_this);
    _this.Auth = new AuthService();
    _this.state = { isLoggedIn: false };
    return _this;
  }

  // handleLoginClick() {
  //   this.setState({ isLoggedIn: true });
  // }

  _createClass(LoginControl, [{
    key: "handleLogoutClick",
    value: function handleLogoutClick() {
      this.Auth.logout();
      this.setState({ isLoggedIn: false });
    }
  }, {
    key: "handleFormSubmit",
    value: function handleFormSubmit(e) {
      var _this2 = this;

      e.preventDefault();

      this.Auth.login(this.state.email, this.state.password).then(function (res) {
        //this.props.history.replace("/");
        _this2.setState({ isLoggedIn: true });
      }).catch(function (err) {
        alert(err);
      });
    }
  }, {
    key: "handleChange",
    value: function handleChange(e) {
      this.setState(_defineProperty({}, e.target.name, e.target.value));
    }
  }, {
    key: "componentWillMount",
    value: function componentWillMount() {
      if (this.Auth.loggedIn()) this.setState({ isLoggedIn: true });
      // if (this.Auth.loggedIn()) this.props.history.replace("/");
    }
  }, {
    key: "render",
    value: function render() {
      var isLoggedIn = this.state.isLoggedIn;
      var button = void 0;

      if (isLoggedIn) {
        button = React.createElement(LogoutButton, { onClick: this.handleLogoutClick });
      }
      // else {
      //   button = <LoginButton onClick={this.handleLoginClick} />;
      // }

      return React.createElement(
        "div",
        null,
        button,
        React.createElement(Greeting, {
          isLoggedIn: isLoggedIn,
          handleChange: this.handleChange,
          handleFormSubmit: this.handleFormSubmit
        })
      );
    }
  }]);

  return LoginControl;
}(React.Component);

// Service component for user login


var AuthService = function () {
  // Initializing important variables
  function AuthService(domain) {
    _classCallCheck(this, AuthService);

    this.domain = domain || "http://localhost:8000/api/users"; // API server domain
    this.fetch = this.fetch.bind(this); // React binding stuff
    this.login = this.login.bind(this);
    // this.getProfile = this.getProfile.bind(this);
  }

  _createClass(AuthService, [{
    key: "login",
    value: function login(email, password) {
      var _this3 = this;

      // Get a token from api server using the fetch api
      return this.fetch(this.domain + "/login", {
        method: "POST",
        body: JSON.stringify({
          email: email,
          password: password
        })
      }).then(function (res) {
        _this3.setToken(res.token); // Setting the token in localStorage
        return Promise.resolve(res);
      });
    }
  }, {
    key: "loggedIn",
    value: function loggedIn() {
      // Checks if there is a saved token and it's still valid
      var token = this.getToken(); // GEtting token from localstorage
      return !!token && !this.isTokenExpired(token); // handwaiving here
    }
  }, {
    key: "isTokenExpired",
    value: function isTokenExpired(token) {
      try {
        var decoded = decode(token);
        if (decoded.exp < Date.now() / 1000) {
          // Checking if token is expired. N
          return true;
        } else return false;
      } catch (err) {
        return false;
      }
    }
  }, {
    key: "setToken",
    value: function setToken(idToken) {
      // Saves user token to localStorage
      localStorage.setItem("id_token", idToken);
    }
  }, {
    key: "getToken",
    value: function getToken() {
      // Retrieves the user token from localStorage
      return localStorage.getItem("id_token");
    }
  }, {
    key: "logout",
    value: function logout() {
      // Clear user token and profile data from localStorage
      localStorage.removeItem("id_token");
    }

    // getProfile() {
    //   // Using jwt-decode npm package to decode the token
    //   return decode(this.getToken());
    // }

  }, {
    key: "fetch",
    value: function (_fetch) {
      function fetch(_x, _x2) {
        return _fetch.apply(this, arguments);
      }

      fetch.toString = function () {
        return _fetch.toString();
      };

      return fetch;
    }(function (url, options) {
      // performs api calls sending the required authentication headers
      var headers = {
        Accept: "application/json",
        "Content-Type": "application/json"
      };

      // Setting Authorization header
      // Authorization: Bearer xxxxxxx.xxxxxxxx.xxxxxx
      if (this.loggedIn()) {
        headers["Authorization"] = "Bearer " + this.getToken();
      }

      return fetch(url, Object.assign({
        headers: headers
      }, options)).then(this._checkStatus).then(function (response) {
        return response.json();
      });
    })
  }, {
    key: "_checkStatus",
    value: function _checkStatus(response) {
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
  }]);

  return AuthService;
}();

// Logged user greeting - returns Todos container


function UserGreeting(props) {
  return React.createElement(
    "div",
    null,
    React.createElement(
      "h1",
      null,
      "Welcome back!"
    ),
    React.createElement(TodoApp, null)
  );
}

// Anonymous user greeting - returns login form
function GuestGreeting(props) {
  return React.createElement(
    "div",
    { className: "center" },
    React.createElement(
      "div",
      { className: "card" },
      React.createElement(
        "h1",
        null,
        "Login"
      ),
      React.createElement(
        "form",
        null,
        React.createElement("input", {
          className: "form-item",
          placeholder: "Email goes here...",
          name: "email",
          type: "text",
          onChange: props.handleChange
        }),
        React.createElement("input", {
          className: "form-item",
          placeholder: "Password goes here...",
          name: "password",
          type: "password",
          onChange: props.handleChange
        }),
        React.createElement("input", {
          className: "form-submit",
          value: "SUBMIT",
          type: "submit",
          onClick: props.handleFormSubmit
        })
      )
    )
  )

  // todo: user registration
  ;
}

// Greetings for two user roles (logged and anonymous)
function Greeting(props) {
  var isLoggedIn = props.isLoggedIn;
  if (isLoggedIn) {
    return React.createElement(UserGreeting, null);
  }
  return React.createElement(GuestGreeting, {
    handleChange: props.handleChange,
    handleFormSubmit: props.handleFormSubmit
  });
}

// Link for login (reserved, not used)
function LoginButton(props) {
  return React.createElement(
    "a",
    { href: "#", onClick: props.onClick },
    "Login"
  );
}

// Link for user logout
function LogoutButton(props) {
  return React.createElement(
    "a",
    { href: "#", onClick: props.onClick },
    "Logout"
  );
}

//-------------------------------------------------------------------------------------------------------

// Todos page title
var Title = function Title(_ref) {
  var todoCount = _ref.todoCount;

  return React.createElement(
    "div",
    null,
    React.createElement(
      "div",
      null,
      React.createElement(
        "h1",
        null,
        "To-do's count: ",
        todoCount ? todoCount.length : 0
      )
    )
  );
};

// Form for adding todo

var TodoForm = function (_React$Component2) {
  _inherits(TodoForm, _React$Component2);

  function TodoForm(props) {
    _classCallCheck(this, TodoForm);

    var _this4 = _possibleConstructorReturn(this, (TodoForm.__proto__ || Object.getPrototypeOf(TodoForm)).call(this, props));

    _this4.handleChange = function (e) {
      _this4.setState(_defineProperty({}, e.target.name, e.target.value));
    };

    _this4.onSubmit = function (e) {
      e.preventDefault();
      var form = {
        name: _this4.state.name,
        description: _this4.state.description,
        state: _this4.state.state
      };
      _this4.props.addTodo(form);
      _this4.setState({
        name: "",
        description: "",
        state: "1"
      });
    };

    _this4.state = {
      name: "",
      description: "",
      state: "1"
    };
    return _this4;
  }

  _createClass(TodoForm, [{
    key: "render",
    value: function render() {
      var _this5 = this;

      return React.createElement(
        "div",
        null,
        React.createElement(
          "form",
          null,
          React.createElement(
            "label",
            null,
            "Name:",
            React.createElement("input", {
              className: "form-control col-md-12",
              name: "name",
              value: this.state.name,
              onChange: function onChange(e) {
                return _this5.handleChange(e);
              }
            })
          ),
          React.createElement(
            "label",
            null,
            "Description:",
            React.createElement("input", {
              className: "form-control col-md-12",
              name: "description",
              value: this.state.description,
              onChange: function onChange(e) {
                return _this5.handleChange(e);
              }
            })
          ),
          React.createElement(
            "select",
            {
              name: "state",
              value: this.state.state,
              onChange: function onChange(e) {
                return _this5.handleChange(e);
              }
            },
            React.createElement(
              "option",
              { value: "1" },
              "created"
            ),
            React.createElement(
              "option",
              { value: "2" },
              "wait"
            ),
            React.createElement(
              "option",
              { value: "3" },
              "canceled"
            ),
            React.createElement(
              "option",
              { value: "4" },
              "blocked"
            ),
            React.createElement(
              "option",
              { value: "5" },
              "in process/doing"
            ),
            React.createElement(
              "option",
              { value: "6" },
              "review"
            ),
            React.createElement(
              "option",
              { value: "7" },
              "done"
            ),
            React.createElement(
              "option",
              { value: "8" },
              "archived"
            )
          ),
          React.createElement(
            "button",
            { onClick: function onClick(e) {
                return _this5.onSubmit(e);
              } },
            "Submit"
          )
        )
      );
    }
  }]);

  return TodoForm;
}(React.Component);

// Changes static text label to input


var EditableLabel = function (_React$Component3) {
  _inherits(EditableLabel, _React$Component3);

  function EditableLabel(props) {
    _classCallCheck(this, EditableLabel);

    var _this6 = _possibleConstructorReturn(this, (EditableLabel.__proto__ || Object.getPrototypeOf(EditableLabel)).call(this, props));

    _this6.state = {
      id: props.todo.id,
      name: props.todo.name,
      description: props.todo.description,
      type: props.type,
      editing: false
    };
    _this6.initEditor();
    _this6.edit = _this6.edit.bind(_this6);
    _this6.save = _this6.save.bind(_this6);
    return _this6;
  }

  _createClass(EditableLabel, [{
    key: "initEditor",
    value: function initEditor() {
      var _this7 = this;

      this.editor = React.createElement("input", {
        type: "text",
        name: this.state.type,
        defaultValue: this.state[this.state.type],
        onKeyPress: function onKeyPress(event) {
          var key = event.which || event.keyCode;
          if (key === 13) {
            //enter key
            _this7.save(event);
          }
        },
        autoFocus: true
      });
    }
  }, {
    key: "edit",
    value: function edit(e) {
      var _setState2;

      this.setState((_setState2 = {}, _defineProperty(_setState2, e.target.name, e.target.value), _defineProperty(_setState2, "editing", true), _setState2));
    }
  }, {
    key: "save",
    value: function save(e) {
      var _setState3;

      var form = {
        id: this.state.id,
        name: e.target.name,
        value: e.target.value
      };
      this.props.update(form);
      this.setState((_setState3 = {}, _defineProperty(_setState3, e.target.name, e.target.value), _defineProperty(_setState3, "editing", false), _setState3));
    }
  }, {
    key: "componentDidUpdate",
    value: function componentDidUpdate() {
      this.initEditor();
    }
  }, {
    key: "render",
    value: function render() {
      return this.state.editing ? this.editor : React.createElement(
        "p",
        { name: this.state.type, onClick: this.edit },
        this.state[this.state.type]
      );
    }
  }]);

  return EditableLabel;
}(React.Component);

// Each Todo component


var Todo = function Todo(_ref2) {
  var todo = _ref2.todo,
      remove = _ref2.remove,
      update = _ref2.update;

  return React.createElement(
    "div",
    null,
    React.createElement(EditableLabel, { todo: todo, type: "name", update: update }),
    React.createElement(EditableLabel, { todo: todo, type: "description", update: update }),
    React.createElement(
      "button",
      {
        onClick: function onClick() {
          remove(todo.id);
        }
      },
      "Delete"
    )
  );
};

// All todo components or nothing
var TodoList = function TodoList(_ref3) {
  var todos = _ref3.todos,
      remove = _ref3.remove,
      update = _ref3.update;

  // Map through the todos
  var todoNode = todos ? todos.map(function (todo) {
    return React.createElement(Todo, { todo: todo, key: todo.id, remove: remove, update: update });
  }) : "There is no to-do's";
  return React.createElement(
    "div",
    { className: "list-group", style: { marginTop: "30px" } },
    todoNode
  );
};

// Container component for Todos

var TodoApp = function (_React$Component4) {
  _inherits(TodoApp, _React$Component4);

  function TodoApp(props) {
    _classCallCheck(this, TodoApp);

    // Set initial state
    var _this8 = _possibleConstructorReturn(this, (TodoApp.__proto__ || Object.getPrototypeOf(TodoApp)).call(this, props));
    // Pass props to parent class


    _this8.state = {
      data: []
    };
    _this8.apiUrl = "http://127.0.0.1:8000/api/todos";
    return _this8;
  }

  // Lifecycle method


  _createClass(TodoApp, [{
    key: "componentDidMount",
    value: function componentDidMount() {
      var _this9 = this;

      // Make HTTP request with Axios
      axios.get(this.apiUrl).then(function (res) {
        // Set state with result
        _this9.setState({ data: res.data || [] });
      });
    }

    // Add todo handler

  }, {
    key: "addTodo",
    value: function addTodo(val) {
      var _this10 = this;

      // Update data
      axios.post(this.apiUrl, val).then(function (res) {
        _this10.state.data.push(res.data);
        _this10.setState({ data: _this10.state.data });
      });
    }

    // Handle update

  }, {
    key: "updateTodo",
    value: function updateTodo(val) {
      var form = _defineProperty({}, val.name, val.value);
      axios.patch(this.apiUrl + "/" + val.id, form);
    }

    // Handle remove

  }, {
    key: "removeTodo",
    value: function removeTodo(id) {
      var _this11 = this;

      // Filter all todos except the one to be removed
      var remainder = this.state.data.filter(function (todo) {
        if (todo.id !== id) return todo;
      });
      // Update state with filter
      axios.delete(this.apiUrl + "/" + id).then(function (res) {
        _this11.setState({ data: remainder });
      });
    }
  }, {
    key: "render",
    value: function render() {
      return React.createElement(
        "div",
        null,
        React.createElement(Title, { todoCount: this.state.data }),
        React.createElement(TodoForm, { addTodo: this.addTodo.bind(this) }),
        React.createElement(TodoList, {
          todos: this.state.data,
          remove: this.removeTodo.bind(this),
          update: this.updateTodo.bind(this)
        })
      );
    }
  }]);

  return TodoApp;
}(React.Component);

ReactDOM.render(React.createElement(LoginControl, null), document.getElementById("root"));