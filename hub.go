package main

import "log"

type Message struct {
	ID   string
	data []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	send chan Message
	//broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	connections map[string]*Client
}

func newHub() *Hub {
	return &Hub{
		send:       make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		connections: make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {

		case client := <-h.register:
			log.Println("Hey new client: " + client.ID)
			//h.clients[client] = true
			h.connections[client.ID] = client

		case client := <-h.unregister:
			log.Println("Hey unregistering client: " + client.ID)
			if _, ok := h.connections[client.ID]; ok {
				log.Println("Hey unregistering client found: " + client.ID)
				delete(h.connections, client.ID)
				log.Println("Hey unregistering after found: " + client.ID)
				close(client.send)
				log.Println("Hey unregistering after w: " + client.ID)
			}

		case Message := <-h.send:
			log.Println("Message")
			if client, ok := h.connections[Message.ID]; ok {
				select {
				case client.send <- Message.data:
				default:
					close(client.send)
					delete(h.connections, client.ID)
				}
			}
			//for client := range h.clients {
			//	select {
			//	case client.send <- message:
			//	default:
			//		close(client.send)
			//		delete(h.clients, client)
			//	}
			//}
		}
	}
}
