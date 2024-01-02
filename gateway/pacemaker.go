package gateway

import (
	"time"

	"golang.org/x/net/websocket"
)

type Pacemaker struct {
	interval time.Duration
	sequence *int
	ticker time.Ticker
	ws *websocket.Conn
}

type Heartbeat struct {
	D *int `json:"d"`
	Op int `json:"op"`
}

func (pm Pacemaker) heartbeat(reset bool) {
	if reset {
		pm.ticker.Reset(pm.interval)
	}
	websocket.JSON.Send(pm.ws, Heartbeat {
		D: pm.sequence,
		Op: 1,
	})
}

func (pm Pacemaker) startTicking() {
	for {
		<-pm.ticker.C
		pm.heartbeat(false)
	}
}

func (pm Pacemaker) Tick() {
	pm.heartbeat(true)
}

func StartPacemaker(ws *websocket.Conn, interval int) Pacemaker {
	duration := time.Millisecond * time.Duration(interval)
	pm := Pacemaker {
		interval: duration,
		sequence: nil,
		ticker: *time.NewTicker(duration),
		ws: ws,
	}
	go pm.startTicking()
	return pm
}

