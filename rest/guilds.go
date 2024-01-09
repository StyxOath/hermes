package rest

import "github.com/StyxOath/hermes/structs"

func (rest REST) GetGuild(id string) (*structs.Guild, error) {
	route := "/guilds/" + id
	var guild structs.Guild
	
	return GetJSON[structs.Guild](rest, route, &guild)
}

func (rest REST) GetGuildChannels(id string) (*[]structs.Channel, error) {
	route := "/guilds/" + id + "/channels"
	var channels []structs.Channel

	return GetJSON[[]structs.Channel](rest, route, &channels)
}

