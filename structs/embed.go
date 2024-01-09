package structs

import "encoding/json"

type MessageEmbed struct {
	// Author information.
	Author *EmbedAuthor `json:"author,omitempty"`

	// Color code of the embed.
	Color int32 `json:"color,omitempty"`

	// Description of embed.
	Description *string `json:"description,omitempty"`

	// Fields information, max of 25.
	Fields []EmbedField `json:"fields,omitempty"`

	// Footer information.
	Footer *EmbedFooter `json:"footer,omitempty"`

	// Image information.
	Image *EmbedImage `json:"image,omitempty"`

	// Provider information.
	Provider *EmbedProvider `json:"provider,omitempty"`

	// Thumbnail information.
	Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`

	// Timestamp of embed content.
	Timestamp *string `json:"timestamp,omitempty"`

	// Title of embed.
	Title *string `json:"title,omitempty"`

	// Type of embed (always "rich" for webhook embeds).
	Type *string `json:"type,omitempty"`

	// Url of embed.
	Url *string `json:"url,omitempty"`

	// Video information.
	Video *EmbedVideo `json:"video,omitempty"`
}

func (embed MessageEmbed) GetAuthor() EmbedAuthor {
	if embed.Author == nil {
		embed.Author = &EmbedAuthor{}
		embed.Author.embed = &embed
	}
	return *embed.Author
}

func (embed MessageEmbed) GetFooter() EmbedFooter {
	if embed.Footer == nil {
		embed.Footer = &EmbedFooter{}
		embed.Footer.embed = &embed
	}

	return *embed.Footer
}

func (embed MessageEmbed) RemoveAuthor() MessageEmbed {
	embed.Author = nil
	return embed
}

func (embed MessageEmbed) RemoveFields() MessageEmbed {
	embed.Fields = nil
	return embed
}

func (embed MessageEmbed) SetColor(color int32) MessageEmbed {
	embed.Color = color
	return embed
}

func (embed MessageEmbed) SetDescription(text string) MessageEmbed {
	embed.Description = &text
	return embed
}

func (embed MessageEmbed) SetImage(url string) MessageEmbed {
	embed.Image = &EmbedImage{
		Url: url,
	}
	return embed
}

func (embed MessageEmbed) SetThumbnail(url string) MessageEmbed {
	embed.Thumbnail = &EmbedThumbnail{
		Url: url,
	}
	return embed
}

func (embed MessageEmbed) JsonString() string {
	bytes, _ := json.Marshal(embed)
	return string(bytes)
}

