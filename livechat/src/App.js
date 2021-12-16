import logo from './logo.svg';
import './App.css';
import React, { Component }  from 'react';
import { connect, sendMessage } from './api';
import Header from "./Header.jsx";
import ChatHistory from './ChatHistory';
import MessageBox from './MessageBox';
import LiveButton from './LiveButton';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
      buttonCount: 0
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("new message");
      if(JSON.parse(msg.data).type === 2) {
        this.setState((prevState) => ({
          buttonCount: parseInt(JSON.parse(msg.data).body)
        }));
      } else {
        this.setState(prevState => ({
          chatHistory: [...this.state.chatHistory, msg]
        }));
        console.log(this.state);
      }
      
    });
  }

  send(e) {
    // console.log("Sent: Hi :)");
    // sendMessage("Hi :)");
    if(e.keyCode === 13) { //if key pressed was enter key
      sendMessage(e.target.value);
      e.target.value = "";
    }
  }

  click(e) {
    sendMessage("btnpress");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <LiveButton click={this.click} count={this.state.buttonCount}/>
        <ChatHistory chatHistory={this.state.chatHistory} />
        <MessageBox send={this.send} />
      </div>
    );
  }
}

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

export default App;
