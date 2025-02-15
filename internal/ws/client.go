package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[string]*Client

type Client struct {
	username string
	connection *websocket.Conn
	manager *Manager
	messageChanel chan Event
}

func NewClient(username string, conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		username: username,
		connection: conn,
		manager: manager,
		messageChanel: make(chan Event),
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		//limpando conexão
		c.manager.removeClient(c.username)
	}()

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling event: %v", err)
			break
		}

		if err := c.manager.RouteEvent(request, c); err != nil {
			log.Printf("error marshalling event: %v", err)
		}

	}
}

func (c *Client) WriteMessages() {
	defer func() {
		//limpando conexão
		c.manager.removeClient(c.username)
	}()

	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case message, ok := <-c.messageChanel:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			event, err := json.Marshal(message)
			if err != nil {
				log.Println("erro ao serializar json")
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, event); err != nil {
				log.Printf("failed to send message: %v", err)
			}
		
		case <-ticker.C:
			log.Println("ping")

			if err := c.connection.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("write msg err ", err)
				return
			}  
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}