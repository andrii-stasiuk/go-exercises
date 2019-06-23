import React, { Component } from "react";
import LoginControl from "./Login";

class App extends Component {
  render() {
    return (
      <div>
        <p>To-do list application</p>
        <LoginControl />
      </div>
    );
  }
}

export default App;
