package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn // socket is the web socket for this client.
	send chan []byte // send is a channel on which messages are sent.
	room *room // room is the room this client is chatting in.
}
