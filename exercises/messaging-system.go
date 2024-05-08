package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Agent struct {
	id int							// Unique identifier for Agent in network 
	network []int					// IDs for Agents connected to it
	channel chan Message			// Agent Message Channel
	network_channels []chan Message	// All channels from a given network
}

type Message struct {
	recipient int	// Recipient Agent ID
	sender int		// Sender Agent ID
	content string
}

func (a *Agent) sendMessage() {
	for {
		recipient := a.network[rand.Intn(len(a.network))]
		message := Message{
			sender: a.id, 
			recipient: recipient, 
			content: "Greetings from Agent to Agent",
		}

		a.network_channels[recipient] <- message
		fmt.Printf("[AGENT %v] Message sent to Agent %v\n", a.id, message.recipient)
		time.Sleep(time.Second * 2)
	}
	
}

func (a *Agent) readMessage() {
	for {
		message := <- a.channel
		fmt.Printf("[AGENT %v] Message received from Agent %v\n", a.id, message.sender)
		time.Sleep(time.Second * 3)
	}
}

func (a *Agent) start() {
	fmt.Printf("[AGENT %v] Up and Running!\n", a.id)
	go a.sendMessage()
	go a.readMessage()

}

func main() {
	messages_agent0 := make(chan Message, 3)
	messages_agent1 := make(chan Message, 4)
	messages_agent2 := make(chan Message, 5)
	messages_agent3 := make(chan Message, 2)
	messages_agent4 := make(chan Message, 1)
	done := make(chan struct{})

	var channels = []chan Message {
		messages_agent0, 
		messages_agent1, 
		messages_agent2, 
		messages_agent3, 
		messages_agent4,
	}

	agent0 := Agent{id: 0, channel: messages_agent0, network_channels: channels, network: []int {1,3}} 
	agent1 := Agent{id: 1, channel: messages_agent1, network_channels: channels, network: []int {0, 4}}
	agent2 := Agent{id: 2, channel: messages_agent2, network_channels: channels, network: []int {0, 1, 3, 4}}
	agent3 := Agent{id: 3, channel: messages_agent3, network_channels: channels, network: []int {2, 1, 4}}
	agent4 := Agent{id: 4, channel: messages_agent4, network_channels: channels, network: []int {0}}

	go agent0.start()
	go agent1.start()
	go agent2.start()
	go agent3.start()
	go agent4.start()

	<- done
}