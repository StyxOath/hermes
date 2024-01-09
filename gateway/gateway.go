package gateway

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

type Gateway struct {
	Channels Dispatch
	events chan []byte
	pacemaker *Pacemaker
	resumeUrl string
	session string
	token string
	ws *websocket.Conn
}

type GatewayEvent struct {
	D map[string]interface{} `json:"d"`
	Op int `json:"op"`
	T *string `json:"t"`
	S *int `json:"s"`
}

func NewGateway(token string) (*Gateway, error) {
	ws, err := websocket.Dial("wss://gateway.discord.gg?v=10&encoding=json", "", "https://discord.gg")

	if err != nil {
		return nil, err
	}

	events := make(chan []byte)
	gateway := &Gateway {
		events: events,
		token: token,
		ws: ws,
	}

	go gateway.handle()
	return gateway, nil
}

func (gateway Gateway) Listen() {
	for {
		var message []byte
		err := websocket.Message.Receive(gateway.ws, &message)
		if err != nil {
			panic(err)
		}
		gateway.events <- message
	}
}

func (gateway Gateway) handle() {
	for {
		event := <-gateway.events

		var message GatewayEvent
		json.Unmarshal(event, &message)

		if message.S != nil {
			gateway.pacemaker.sequence = message.S
		}

		switch message.Op {
		case 0:
			if *message.T == "MESSAGE_CREATE" {
				send(&event, gateway.Channels.MessageCreate)
			}
		case 1:
			gateway.pacemaker.Tick()

		case 10:
			interval := int(message.D["heartbeat_interval"].(float64))
			pacemaker := StartPacemaker(gateway.ws, interval)
			gateway.pacemaker = &pacemaker
			gateway.identify()
		}
	}
}

func send[T any](event *[]byte, channel *chan T) *T {
	if channel == nil {
		return nil
	}

	var data DispatchEvent[T]
	json.Unmarshal(*event, &data)
	*channel <- data.Payload

	return &data.Payload
}

type ConnectionProperties struct {
	Browser string `json:"browser"`
	Device string `json:"device"`
	Os string `json:"os"`
}

type IdentifyPayload struct {
	Intents int `json:"intents"`
	Properties ConnectionProperties `json:"properties"`
	Token string `json:"token"`
}

type IdentifyGatewayEvent struct {
	Op int `json:"op"`
	Payload IdentifyPayload `json:"d"`
}

func (gateway Gateway) identify() {
	properties := ConnectionProperties {
		Browser: "hermes-go",
		Device: "hermes-go",
		Os: "linux",
	}
	payload := IdentifyPayload {
		Intents: 512,
		Properties: properties,
		Token: gateway.token,
	}
	event := IdentifyGatewayEvent {
		Payload: payload,
		Op: 2,
	}
	websocket.JSON.Send(gateway.ws, event)
}

