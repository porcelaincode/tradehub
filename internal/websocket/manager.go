// internal/websocket/manager.go
package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
    Conn *websocket.Conn
}

type Manager struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
    mu         sync.Mutex
}

var manager = Manager{
    clients:    make(map[*Client]bool),
    broadcast:  make(chan []byte),
    register:   make(chan *Client),
    unregister: make(chan *Client),
}

func GetManager() *Manager {
    return &manager
}

func (m *Manager) Start() {
    for {
        select {
        case client := <-m.register:
            m.mu.Lock()
            m.clients[client] = true
            m.mu.Unlock()
        case client := <-m.unregister:
            m.mu.Lock()
            if _, ok := m.clients[client]; ok {
                delete(m.clients, client)
                client.Conn.Close()
            }
            m.mu.Unlock()
        case message := <-m.broadcast:
            m.mu.Lock()
            for client := range m.clients {
                err := client.Conn.WriteMessage(websocket.TextMessage, message)
                if err != nil {
                    client.Conn.Close()
                    delete(m.clients, client)
                }
            }
            m.mu.Unlock()
        }
    }
}

func (m *Manager) RegisterClient(conn *websocket.Conn) {
    client := &Client{Conn: conn}
    m.register <- client
}

func (m *Manager) UnregisterClient(client *Client) {
    m.unregister <- client
}

func (m *Manager) Broadcast(message []byte) {
    m.broadcast <- message
}
