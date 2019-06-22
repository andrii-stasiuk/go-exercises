var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var LoginControl = function (_React$Component) {
  _inherits(LoginControl, _React$Component);

  function LoginControl(props) {
    _classCallCheck(this, LoginControl);

    var _this = _possibleConstructorReturn(this, (LoginControl.__proto__ || Object.getPrototypeOf(LoginControl)).call(this, props));

    _this.handleLoginClick = _this.handleLoginClick.bind(_this);
    _this.handleLogoutClick = _this.handleLogoutClick.bind(_this);
    _this.state = { isLoggedIn: false };
    return _this;
  }

  _createClass(LoginControl, [{
    key: "handleLoginClick",
    value: function handleLoginClick() {
      this.setState({ isLoggedIn: true });
    }
  }, {
    key: "handleLogoutClick",
    value: function handleLogoutClick() {
      this.setState({ isLoggedIn: false });
    }
  }, {
    key: "render",
    value: function render() {
      var isLoggedIn = this.state.isLoggedIn;
      var button = void 0;

      if (isLoggedIn) {
        button = React.createElement(LogoutButton, { onClick: this.handleLogoutClick });
      } else {
        button = React.createElement(LoginButton, { onClick: this.handleLoginClick });
      }

      return React.createElement(
        "div",
        null,
        button,
        React.createElement(Greeting, { isLoggedIn: isLoggedIn })
      );
    }
  }]);

  return LoginControl;
}(React.Component);

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

function GuestGreeting(props) {
  return React.createElement(
    "h1",
    null,
    "Please sign up."
  );
}

function Greeting(props) {
  var isLoggedIn = props.isLoggedIn;
  if (isLoggedIn) {
    return React.createElement(UserGreeting, null);
  }
  return React.createElement(GuestGreeting, null);
}

function LoginButton(props) {
  return React.createElement(
    "a",
    { href: "#", onClick: props.onClick },
    "Login"
  );
}

function LogoutButton(props) {
  return React.createElement(
    "a",
    { href: "#", onClick: props.onClick },
    "Logout"
  );
}

//-------------------------------------------------------------------------------------------------------

// Page title
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

// Form for to-do adding

var TodoForm = function (_React$Component2) {
  _inherits(TodoForm, _React$Component2);

  function TodoForm(props) {
    _classCallCheck(this, TodoForm);

    var _this2 = _possibleConstructorReturn(this, (TodoForm.__proto__ || Object.getPrototypeOf(TodoForm)).call(this, props));

    _this2.handleChange = function (e) {
      _this2.setState(_defineProperty({}, e.target.name, e.target.value));
    };

    _this2.onSubmit = function (e) {
      e.preventDefault();
      var form = {
        name: _this2.state.name,
        description: _this2.state.description,
        state: _this2.state.state
      };
      _this2.props.addTodo(form);
      _this2.setState({
        name: "",
        description: "",
        state: "1"
      });
    };

    _this2.state = {
      name: "",
      description: "",
      state: "1"
    };
    return _this2;
  }

  _createClass(TodoForm, [{
    key: "render",
    value: function render() {
      var _this3 = this;

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
                return _this3.handleChange(e);
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
                return _this3.handleChange(e);
              }
            })
          ),
          React.createElement(
            "select",
            {
              name: "state",
              value: this.state.state,
              onChange: function onChange(e) {
                return _this3.handleChange(e);
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
                return _this3.onSubmit(e);
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

    var _this4 = _possibleConstructorReturn(this, (EditableLabel.__proto__ || Object.getPrototypeOf(EditableLabel)).call(this, props));

    _this4.state = {
      id: props.todo.id,
      name: props.todo.name,
      description: props.todo.description,
      type: props.type,
      editing: false
    };
    _this4.initEditor();
    _this4.edit = _this4.edit.bind(_this4);
    _this4.save = _this4.save.bind(_this4);
    return _this4;
  }

  _createClass(EditableLabel, [{
    key: "initEditor",
    value: function initEditor() {
      var _this5 = this;

      this.editor = React.createElement("input", {
        type: "text",
        name: this.state.type,
        defaultValue: this.state[this.state.type],
        onKeyPress: function onKeyPress(event) {
          var key = event.which || event.keyCode;
          if (key === 13) {
            //enter key
            _this5.save(event);
          }
        },
        autoFocus: true
      });
    }
  }, {
    key: "edit",
    value: function edit(e) {
      var _setState;

      this.setState((_setState = {}, _defineProperty(_setState, e.target.name, e.target.value), _defineProperty(_setState, "editing", true), _setState));
    }
  }, {
    key: "save",
    value: function save(e) {
      var _setState2;

      var form = {
        id: this.state.id,
        name: e.target.name,
        value: e.target.value
      };
      this.props.update(form);
      this.setState((_setState2 = {}, _defineProperty(_setState2, e.target.name, e.target.value), _defineProperty(_setState2, "editing", false), _setState2));
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

// Each Todo


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

// All todos
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

// Container Component

var TodoApp = function (_React$Component4) {
  _inherits(TodoApp, _React$Component4);

  function TodoApp(props) {
    _classCallCheck(this, TodoApp);

    // Set initial state
    var _this6 = _possibleConstructorReturn(this, (TodoApp.__proto__ || Object.getPrototypeOf(TodoApp)).call(this, props));
    // Pass props to parent class


    _this6.state = {
      data: []
    };
    _this6.apiUrl = "http://127.0.0.1:8000/api/todos";
    return _this6;
  }

  // Lifecycle method


  _createClass(TodoApp, [{
    key: "componentDidMount",
    value: function componentDidMount() {
      var _this7 = this;

      // Make HTTP request with Axios
      axios.get(this.apiUrl).then(function (res) {
        // Set state with result
        _this7.setState({ data: res.data || [] });
      });
    }

    // Add todo handler

  }, {
    key: "addTodo",
    value: function addTodo(val) {
      var _this8 = this;

      // Update data
      axios.post(this.apiUrl, val).then(function (res) {
        _this8.state.data.push(res.data);
        _this8.setState({ data: _this8.state.data });
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
      var _this9 = this;

      // Filter all todos except the one to be removed
      var remainder = this.state.data.filter(function (todo) {
        if (todo.id !== id) return todo;
      });
      // Update state with filter
      axios.delete(this.apiUrl + "/" + id).then(function (res) {
        _this9.setState({ data: remainder });
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