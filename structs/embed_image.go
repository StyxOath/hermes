package structs

type EmbedImage struct {
	Height *int `json:"height,omitempty"`
	ProxyUrl *string `json:"proxy_url,omitempty"`
	Url string `json:"url"`
	Width *int `json:"width,omitempty"`
}

type EmbedThumbnail EmbedImage

type EmbedVideo EmbedImage

