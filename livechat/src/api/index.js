var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (cb) => {
    console.log("Attempting connection...");
    socket.onopen = () => {
        console.log("Connected successfully");
    };
    socket.onmessage = (msg) => {
        console.log("New message! ", msg);
        cb(msg);
    };
    socket.onclose = (event) => {
        console.log("Socket closed connection: ", event);
    };
    socket.onerror = (error) => {
        console.log("Socket error: ", error);
    };
};

let sendMessage = (message) => {
    console.log("sending message: ", message);
    socket.send(message);
};

export { connect, sendMessage };