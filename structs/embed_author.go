package structs

type EmbedAuthor struct {
	embed *MessageEmbed
	IconUrl *string `json:"icon_url,omitempty"`
	Name string `json:"name"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
	Url *string `json:"url,omitempty"`
}

func (author EmbedAuthor) Embed() MessageEmbed {
	if author.embed == nil {
		author.embed = &MessageEmbed{}
		author.embed.Author = &author
	}

	return *author.embed
}

func (author EmbedAuthor) SetIcon(url string) EmbedAuthor {
	if url == "" {
		author.IconUrl = nil
	} else {
		author.IconUrl = &url
	}
	return author
}

func (author EmbedAuthor) SetName(text string) EmbedAuthor {
	author.Name = text
	return author
}

func (author EmbedAuthor) SetUrl(url string) EmbedAuthor {
	if url == "" {
		author.Url = nil
	} else {
		author.Url = &url
	}
	return author
}

