package structs

type ChannelType int8

var CHANNEL_TYPES = struct {
	// A text channel within a server.
	GuildText ChannelType

	// A direct message between users.
	Dm ChannelType

	// A voice channel within a server.
	GuildVoice ChannelType

	// A direct message between multiple users.
	GroupDm ChannelType

	// An organizational category that contains up to ChannelType
	GuildCategory ChannelType

	// A channel that users can follow and crosspost into their own server (formerly news channels).
	GuildAnnouncement ChannelType

	// A temporary sub-channel within a GUILDAnnouncement channel.
	AnnouncementThread ChannelType

	// A temporary sub-channel within a GUILDText or GUILDForum channel.
	PublicThread ChannelType

	// A temporary sub-channel within a GUILDText channel that is only viewable by those invited and those with the MANAGEThreads permission.
	PrivateThread ChannelType

	// A voice channel for hosting events with an audience.
	GuildStageVoice ChannelType

	// The channel in a hub containing the listed servers.
	GuildDirectory ChannelType

	// Channel that can only contain threads.
	GuildForum ChannelType

	// Channel that can only contain threads, similar to GUILDForum channels.
	GuildMedia ChannelType
} {
	GuildText: 0,
	Dm: 1,
	GuildVoice: 2,
	GroupDm: 3,
	GuildCategory: 4,
	GuildAnnouncement: 5,
	AnnouncementThread: 10,
	PublicThread: 11,
	PrivateThread: 12,
	GuildStageVoice: 13,
	GuildDirectory: 14,
	GuildForum: 15,
	GuildMedia: 16,
}

