var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

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

var TodoForm = function (_React$Component) {
  _inherits(TodoForm, _React$Component);

  function TodoForm(props) {
    _classCallCheck(this, TodoForm);

    var _this = _possibleConstructorReturn(this, (TodoForm.__proto__ || Object.getPrototypeOf(TodoForm)).call(this, props));

    _this.handleChange = function (e) {
      _this.setState(_defineProperty({}, e.target.name, e.target.value));
    };

    _this.onSubmit = function (e) {
      e.preventDefault();
      var form = {
        name: _this.state.name,
        description: _this.state.description,
        state: _this.state.state
      };
      _this.props.addTodo(form);
      _this.setState({
        name: "",
        description: "",
        state: "1"
      });
    };

    _this.state = {
      name: "",
      description: "",
      state: "1"
    };
    return _this;
  }

  _createClass(TodoForm, [{
    key: "render",
    value: function render() {
      var _this2 = this;

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
                return _this2.handleChange(e);
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
                return _this2.handleChange(e);
              }
            })
          ),
          React.createElement(
            "select",
            {
              name: "state",
              value: this.state.state,
              onChange: function onChange(e) {
                return _this2.handleChange(e);
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
                return _this2.onSubmit(e);
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


var EditableLabel = function (_React$Component2) {
  _inherits(EditableLabel, _React$Component2);

  function EditableLabel(props) {
    _classCallCheck(this, EditableLabel);

    var _this3 = _possibleConstructorReturn(this, (EditableLabel.__proto__ || Object.getPrototypeOf(EditableLabel)).call(this, props));

    _this3.state = {
      id: props.todo.id,
      name: props.todo.name,
      description: props.todo.description,
      type: props.type,
      editing: false
    };
    _this3.initEditor();
    _this3.edit = _this3.edit.bind(_this3);
    _this3.save = _this3.save.bind(_this3);
    return _this3;
  }

  _createClass(EditableLabel, [{
    key: "initEditor",
    value: function initEditor() {
      var _this4 = this;

      this.editor = React.createElement("input", {
        type: "text",
        name: this.state.type,
        defaultValue: this.state[this.state.type],
        onKeyPress: function onKeyPress(event) {
          var key = event.which || event.keyCode;
          if (key === 13) {
            //enter key
            _this4.save(event);
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

var TodoApp = function (_React$Component3) {
  _inherits(TodoApp, _React$Component3);

  function TodoApp(props) {
    _classCallCheck(this, TodoApp);

    // Set initial state
    var _this5 = _possibleConstructorReturn(this, (TodoApp.__proto__ || Object.getPrototypeOf(TodoApp)).call(this, props));
    // Pass props to parent class


    _this5.state = {
      data: []
    };
    _this5.apiUrl = "http://127.0.0.1:8000/api/todos";
    return _this5;
  }

  // Lifecycle method


  _createClass(TodoApp, [{
    key: "componentDidMount",
    value: function componentDidMount() {
      var _this6 = this;

      // Make HTTP request with Axios
      axios.get(this.apiUrl).then(function (res) {
        // Set state with result
        _this6.setState({ data: res.data || [] });
      });
    }

    // Add todo handler

  }, {
    key: "addTodo",
    value: function addTodo(val) {
      var _this7 = this;

      // Update data
      axios.post(this.apiUrl, val).then(function (res) {
        _this7.state.data.push(res.data);
        _this7.setState({ data: _this7.state.data });
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
      var _this8 = this;

      // Filter all todos except the one to be removed
      var remainder = this.state.data.filter(function (todo) {
        if (todo.id !== id) return todo;
      });
      // Update state with filter
      axios.delete(this.apiUrl + "/" + id).then(function (res) {
        _this8.setState({ data: remainder });
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

ReactDOM.render(React.createElement(TodoApp, null), document.getElementById("root"));