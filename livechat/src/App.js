import logo from './logo.svg';
import './App.css';
import React, { Component }  from 'react';
import { connect, sendMessage } from './api';
import Header from "./Header.jsx";
import ChatHistory from './ChatHistory';
import MessageBox from './MessageBox';
import LiveButton from './LiveButton';
import UserList from './UserList';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
      buttonCount: 0,
      activeUsers: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("new message");

      //Adding users to user list when connection established
      if(JSON.parse(msg.data).type === 5) { 
        this.setState((prevState) => ({
          activeUsers: JSON.parse(msg.data).IDlist
        }));
        console.log(this.state.activeUsers);
        return;
      }

      //Removing user from user list when user disconnects
      if(JSON.parse(msg.data).type === 4) {
        this.setState((prevState) => ({
          activeUsers: this.state.activeUsers.filter(function(person) {
            return person !== JSON.parse(msg.data).sender;
          })
        }));
        console.log(this.state.activeUsers);
      }

      //Updating the live button counter when get notified of a button update
      if(JSON.parse(msg.data).type === 2) {
        this.setState((prevState) => ({
          buttonCount: parseInt(JSON.parse(msg.data).body)
        }));
        return;
      }

      //Adding message to chat history
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }));
      console.log(this.state);
      
    });
  }

  send(e) {
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
        <UserList users={this.state.activeUsers}/>
        <ChatHistory chatHistory={this.state.chatHistory} />
        <MessageBox send={this.send} />
      </div>
    );
  }
}

export default App;
