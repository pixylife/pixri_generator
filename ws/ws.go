package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var Messages = make(map[int64] Message)


var (
	upgrader = websocket.Upgrader{}
)

type Message struct {
	UUID string
	Message string
}

type ClientManager struct {
	Clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

var Manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.register:
			manager.Clients[conn] = true
		case conn := <-manager.unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.send)
				delete(manager.Clients, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.Clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}
func WsPage(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	val := c.ParamValues()


	client := &Client{id:val[0], socket: ws, send: make(chan []byte)}

	Manager.register <- client

	for {
		// Write

		for x := range Manager.Clients{
			for  i,message := range Messages {
				if x.id == message.UUID {
					err := x.socket.WriteMessage(websocket.TextMessage, []byte(message.Message))
					delete(Messages, i)
					if err != nil {
						c.Logger().Error(err)
					}
				}
			}
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}

}

