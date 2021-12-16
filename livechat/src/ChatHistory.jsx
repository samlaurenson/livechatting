import React, {Component} from "react";
import "./ChatHistory.scss";
import Message from "./Message";

class ChatHistory extends Component {
    render() {
        // const messages = this.props.chatHistory.map((msg, index) => (
        //     <p key={index}>{msg.data}</p>
        // ));
        const messages = this.props.chatHistory.map(msg => <Message message={msg.data} />);
        return(
            <div className="chatHistory">
                <h2>History</h2>
                {messages}
            </div>
        );
    }
}

export default ChatHistory;