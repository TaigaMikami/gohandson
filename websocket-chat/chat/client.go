package main

import (
	"github.com/gorilla/websocket"
	"time"
)

// clientはチャットを行っている１人のユーザーを表す
type client struct {
	socket *websocket.Conn // socket is the web socket for this client.
	send chan *message
	room *room // room is the room this client is chatting in.
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Sent_at = msg.When.Format("01-02-03-04")
			msg.Name = c.userData["name"].(string)
			if avatarURL, ok := c.userData["avatar_url"]; ok {
				msg.AvatarURL, _ = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
