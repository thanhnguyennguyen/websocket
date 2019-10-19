import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
    let socket = new WebSocket("ws://127.0.0.1:8080/ws");
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
        socket.send("Hi From the Client!")
    };
    socket.onmessage = event => {
        console.log("Receive message from socket: ", event.data);
    };
    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
    return (
        <div className="App">
        <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        </header>
        </div>
);
}

export default App;
