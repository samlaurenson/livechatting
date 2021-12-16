import React, {Component} from 'react';
import "./LiveButton.scss";

class LiveButton extends Component {
    render() {
        return (
            <div className="buttonContainer">
                <button onClick={this.props.click}>Click me</button>
                <span id="clickCount">{this.props.count}</span>
            </div>
        );
    }
}

export default LiveButton;