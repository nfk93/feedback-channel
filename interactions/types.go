package interactions

import "github.com/bwmarrin/discordgo"

type Interaction struct {
	ID            string             `json:"id"`
	ApplicationId string             `json:"application_id"`
	Type          int                `json:"type"`
	Data          *InteractionData   `json:"data,omitempty"`
	GuildID       string             `json:"guild_id,omitempty"`
	Member        *GuildMember       `json:"member,omitempty"`
	User          *User              `json:"user,omitempty"`
	Token         string             `json:"token"`
	Version       int                `json:"version"`
	Message       *discordgo.Message `json:"message,omitempty"`
	Locale        string             `json:"locale,omitempty"`
	GuildLocale   string             `json:"guild_locale,omitempty"`
}

type InteractionData struct {
	ID            string                  `json:"id"`
	Name          string                  `json:"name"`
	Type          int                     `json:"type"`
	Resolved      *ResolvedData           `json:"resolved,omitempty"`
	Options       []InteractionOptionData `json:"options,omitempty"`
	CustomID      string                  `json:"custom_id,omitempty"`
	ComponentType int                     `json:"component_type,omitempty"`
	Values        []SelectOption          `json:"values,omitempty"`
	TargetID      string                  `json:"target_id,omitempty"`
}

type ResolvedData struct {
	User     map[string]User              `json:"users,omitempty"`
	Members  map[string]GuildMember       `json:"members,omitempty"`
	Roles    map[string]discordgo.Role    `json:"roles,omitempty"`
	Channels map[string]discordgo.Channel `json:"channels,omitempty"`
	Messages map[string]discordgo.Message `json:"messages,omitempty"`
}

type InteractionOptionData struct {
	Name    string                  `json:"name"`
	Type    int                     `json:"type"`
	Value   string                  `json:"value,omitempty"`
	Options []InteractionOptionData `json:"options,omitempty"`
	Focused bool                    `json:"focused,omitempty"`
}

type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Default     bool   `json:"default,omitempty"`
	// Emoji
}

type GuildMember struct {
	User                       *User    `json:"user,omitempty"`
	Nick                       string   `json:"nick,omitempty"`
	Avatar                     string   `json:"avatar,omitempty"`
	Roles                      []string `json:"roles"`
	JoinedAt                   string   `json:"joined_at"`
	PremiumSince               string   `json:"premium_since,omitempty"`
	Deaf                       bool     `json:"deaf"`
	Mute                       bool     `json:"mute"`
	Pending                    bool     `json:"pending,omitempty"`
	Permissions                string   `json:"permissions,omitempty"`
	CommunicationDisabledUntil string   `json:"communication_disabled_until,omitempty"`
}

type User struct {
	// The ID of the user.
	ID string `json:"id"`

	// The user's username.
	Username string `json:"username"`

	// The discriminator of the user (4 numbers after name).
	Discriminator string `json:"discriminator"`

	// The hash of the user's avatar. Use Session.UserAvatar
	// to retrieve the avatar itself.
	Avatar string `json:"avatar"`

	// Whether the user is a bot.
	Bot bool `json:"bot,omitempty"`

	// Whether the user is an Official Discord System user (part of the urgent message system).
	System bool `json:"system,omitempty"`

	// Whether the user has multi-factor authentication enabled.
	MFAEnabled bool `json:"mfa_enabled,omitempty"`

	Banner string `json:"banner,omitempty"`

	AccentColor int `json:"accent_color,omitempty"`

	// The user's chosen language option.
	Locale string `json:"locale,omitempty"`

	// The email of the user. This is only present when
	// the application possesses the email scope for the user.
	Email string `json:"email,omitempty"`

	// Whether the user's email is verified.
	Verified bool `json:"verified,omitempty"`

	// The public flags on a user's account.
	// This is a combination of bit masks; the presence of a certain flag can
	// be checked by performing a bitwise AND between this int and the flag.
	PublicFlags int `json:"public_flags,omitempty"`

	// The type of Nitro subscription on a user's account.
	// Only available when the request is authorized via a Bearer token.
	PremiumType int `json:"premium_type,omitempty"`

	// The flags on a user's account.
	// Only available when the request is authorized via a Bearer token.
	Flags int `json:"flags,omitempty"`
}
