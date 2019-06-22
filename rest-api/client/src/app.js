class LoginControl extends React.Component {
  constructor(props) {
    super(props);
    this.handleLoginClick = this.handleLoginClick.bind(this);
    this.handleLogoutClick = this.handleLogoutClick.bind(this);
    this.state = { isLoggedIn: false };
  }

  handleLoginClick() {
    this.setState({ isLoggedIn: true });
  }

  handleLogoutClick() {
    this.setState({ isLoggedIn: false });
  }

  render() {
    const isLoggedIn = this.state.isLoggedIn;
    let button;

    if (isLoggedIn) {
      button = <LogoutButton onClick={this.handleLogoutClick} />;
    } else {
      button = <LoginButton onClick={this.handleLoginClick} />;
    }

    return (
      <div>
        {button}
        <Greeting isLoggedIn={isLoggedIn} />
      </div>
    );
  }
}

function UserGreeting(props) {
  return (
    <div>
      <h1>Welcome back!</h1>
      <TodoApp />
    </div>
  );
}

function GuestGreeting(props) {
  return <h1>Please sign up.</h1>;
}

function Greeting(props) {
  const isLoggedIn = props.isLoggedIn;
  if (isLoggedIn) {
    return <UserGreeting />;
  }
  return <GuestGreeting />;
}

function LoginButton(props) {
  return (
    <a href="#" onClick={props.onClick}>
      Login
    </a>
  );
}

function LogoutButton(props) {
  return (
    <a href="#" onClick={props.onClick}>
      Logout
    </a>
  );
}

//-------------------------------------------------------------------------------------------------------

// Page title
const Title = ({ todoCount }) => {
  return (
    <div>
      <div>
        <h1>To-do&apos;s count: {todoCount ? todoCount.length : 0}</h1>
      </div>
    </div>
  );
};

// Form for to-do adding
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

// Each Todo
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

// All todos
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

// Container Component
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
