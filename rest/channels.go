package rest

import (
	"fmt"

	"github.com/StyxOath/hermes/structs"
)

func (rest REST) GetChannel(id string) (*structs.Channel, error) {
	route := fmt.Sprintf("/channels/%s", id)

	var channel structs.Channel
	return GetJSON[structs.Channel](rest, route, &channel)
}

