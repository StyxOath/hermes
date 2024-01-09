package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type REST struct {
	base string
	token string
	version int8
}

func NewRest(token string) REST {
	return REST {
		base: "https://discord.com/api/v10",
		token: token,
		version: 10,
	}
}

func (rest REST) setHeaders(header http.Header) {
	header.Set("Authorization", "Bot " + rest.token)
	header.Set("User-Agent", "hermes-go (https://bitomic.dev, 0.0.1)")
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")
}

func (rest REST) Get(route string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", rest.base + route, nil)
	rest.setHeaders(req.Header)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, errors.New(res.Status)
	}

	return res.Body, nil
}

func GetJSON[T any](rest REST, route string, result *T) (*T, error) {
	body, err := rest.Get(route)

	if err != nil {
		return nil, err
	}

	defer body.Close()

	err = json.NewDecoder(body).Decode(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func PostJSON[T any](rest REST, route string, body T) (io.ReadCloser, error) {
	encoded, err := json.Marshal(body)
	log.Println(string(encoded))

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", rest.base + route, bytes.NewBuffer(encoded))
	rest.setHeaders(req.Header)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, errors.New(res.Status)
	}

	return res.Body, nil
}

