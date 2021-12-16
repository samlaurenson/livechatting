import React, {Component} from "react";
import "./MessageBox.scss";


class MessageBox extends Component {
    render() {
        return(
            <div className="messageBoxContainer">
                <input onKeyDown={this.props.send}></input>
            </div>
        );
    }
}

export default MessageBox;