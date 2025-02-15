package ws

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct{
	clients ClientList
	sync.RWMutex

	handlers map[string]EventHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients: make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessageHandler
}

func SendMessageHandler(event Event, c *Client, m *Manager) error {
	var message SendMessageEvent
	if err := json.Unmarshal(event.Payload, &message); err != nil {
		return errors.New("erro ao desserializar json")
	}
	if client, ok := m.clients[message.To]; ok {
		client.messageChanel <- event
	}

	return nil
}

func (m *Manager) RouteEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c, m); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}
}

func (m *Manager) ServeWS(c *gin.Context) {

	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	username := c.MustGet("username").(string)
	client := NewClient(username, conn, m)
	m.addClient(username, client)

	go client.ReadMessages()
	go client.WriteMessages()
}

func (m *Manager) addClient(username string, client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[username] = client
}

func (m *Manager) removeClient(username string) {
	m.Lock()
	defer m.Unlock()
	
	if _, ok := m.clients[username]; ok {
		m.clients[username].connection.Close()
		delete(m.clients, username)
	}
}