package structs

type EmbedFooter struct {
	embed *MessageEmbed
	IconUrl *string `json:"icon_url,omitempty"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
	Text string `json:"text"`
}

func (footer EmbedFooter) Embed() MessageEmbed {
	if footer.embed == nil {
		footer.embed = &MessageEmbed{}
		footer.embed.Footer = &footer
	}

	return *footer.embed
}

func (footer EmbedFooter) RemoveIcon() EmbedFooter {
	footer.IconUrl = nil
	return footer
}

func (footer EmbedFooter) SetIcon(url string) EmbedFooter {
	footer.IconUrl = &url
	return footer
}

func (footer EmbedFooter) SetText(text string) EmbedFooter {
	footer.Text = text
	return footer
}

