package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	conn          net.Conn
	name          string
	channels      map[string]*channel
	activeChannel *channel
}

// method that creates a client connection instance to server
func (server *server) newClient(conn net.Conn) {
	log.Printf("new client has joined: %v", conn.RemoteAddr().String())

	client := client{
		conn:          conn,
		name:          "anonymous",
		channels:      make(map[string]*channel),
		activeChannel: server.channels["general"],
	}

	server.clients[conn.RemoteAddr()] = &client

	server.readInput(&client)
}

// method that reads input from client and sends commands to server
func (server *server) readInput(client *client) {
	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {

		args := strings.Split(scanner.Text(), " ")
		command := strings.TrimSpace(args[0])

		switch command {
		case "/name":
			client.activeChannel.messagesReceiver <- message{
				id:         "name",
				date:       time.Now(),
				fromClient: client,
				args:       args,
			}
		case "/channel":
			client.activeChannel.messagesReceiver <- message{
				id:         "channel",
				date:       time.Now(),
				fromClient: client,
				args:       args,
			}
		case "/msg":
			client.activeChannel.messagesReceiver <- message{
				id:         "channelReceive",
				date:       time.Now(),
				fromClient: client,
				args:       args,
			}
		case "/private":
			client.activeChannel.messagesReceiver <- message{
				id:         "privateReceive",
				date:       time.Now(),
				fromClient: client,
				args:       args,
			}
		case "/quit":
			client.activeChannel.messagesReceiver <- message{
				id:         "quit",
				date:       time.Now(),
				fromClient: client,
				args:       args,
			}
		default:
			server.systemChannel <- systemMessage{
				id:       "error",
				date:     time.Now(),
				toClient: client,
				args:     "Unknown command",
			}
		}
	}
}