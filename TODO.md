go-type -r <roomname> -u <username>

setup:

1. Create socket connection
2. Client sends room name to server
3. Server either creates new room or sends a message to other players in that room
4. Server generates/sends text to client
5. clients wait for countdown
6. Once countdown finishes, clients can begin to type
7. Each time the client's index changes, send to server. Server will then relay than to all clients in room
8. 


SERVER

client {
    username,
    socket
}

room {
    []clients
    
}

rooms map[string]room

1. Wait in a loop for a socket connection
2. Once connected, run a goroutine to handle that connection
3. 
