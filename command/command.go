package command

type Command struct {
	Id                string          `json:"id"`
	Type              int             `json:"type"` // default value is 1
	ApplicationId     string          `json:"application_id"`
	GuildId           string          `json:"guild_id,omitempty"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Options           []CommandOption `json:"options,omitempty"`            // required must be before optional
	DefaultPermission bool            `json:"default_permission,omitempty"` // default is true
	Version           string          `json:"version"`
}

type CommandOption struct {
	Type         int                   `json:"type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Required     bool                  `json:"required,omitempty"` // default true
	Choices      []CommandOptionChoice `json:"choices,omitempty"`
	Options      []CommandOption       `json:"options,omitempty"`
	ChannelTypes []int                 `json:"channel_types,omitempty"`
}

type CommandOptionChoice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
