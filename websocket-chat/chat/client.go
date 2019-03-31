package main

import "github.com/gorilla/websocket"

// clientはチャットを行っている１人のユーザーを表す
type client struct {
	socket *websocket.Conn // socket is the web socket for this client.
	send chan []byte // send is a channel on which messages are sent.
	room *room // room is the room this client is chatting in.
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
				c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
