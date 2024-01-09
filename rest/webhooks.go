package rest

import (
	"encoding/json"
	"fmt"

	"github.com/StyxOath/hermes/structs"
)

type CreateWebhookParams struct {
	Avatar string `json:"avatar,omitempty"`
	Name string `json:"name"`
}

func (rest REST) CreateWebhook(channel string, params CreateWebhookParams) (*structs.Webhook, error) {
	route := fmt.Sprintf("/channels/%s/webhooks", channel)

	body, err := PostJSON[CreateWebhookParams](rest, route, params)

	if err != nil {
		return nil, err
	}

	var webhook structs.Webhook
	json.NewDecoder(body).Decode(&webhook)
	return &webhook, nil
}

func (rest REST) GetChannelWebhooks(channel string) (*[]structs.Webhook, error) {
	route := fmt.Sprintf("/channels/%s/webhooks", channel)

	var result []structs.Webhook
	return GetJSON[[]structs.Webhook](rest, route, &result)
}

func (rest REST) GetWebhook(id string) (*structs.Webhook, error) {
	route := fmt.Sprintf("/webhooks/%s", id)

	var result structs.Webhook
	return GetJSON[structs.Webhook](rest, route, &result)
}

type WebhookPayload struct {
	AvatarUrl string `json:"avatar_url,omitempty"`
	Content string `json:"content,omitempty"`
	Embeds *[]structs.MessageEmbed `json:"embeds,omitempty"`
	Username string `json:"username,omitempty"`
}

func (rest REST) ExecuteWebhook(id string, token string, params WebhookPayload) error {
	route := fmt.Sprintf("/webhooks/%s/%s", id, token)

	_, err := PostJSON[WebhookPayload](rest, route, params)
	return err
}

