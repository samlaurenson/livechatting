import React, {Component} from "react";
import "./UserList.scss";

class UserList extends Component {
    render() {
        const activeusers = this.props.users.map(usr => <div>{usr}</div>);
        return(
            <div className="userList">
                <h3>User List</h3>
                {activeusers}
            </div>
        );
    }
}

export default UserList;