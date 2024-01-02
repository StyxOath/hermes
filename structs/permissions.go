package structs

type Permission int


var PERMISSIONS = struct {
	// Allows creation of instant invites.
	CreateInstantInvite Permission

	// Allows kicking members.
	KickMembers Permission

	// Allows banning members.
	BanMembers Permission

	// Allows all permissions and bypasses channel permission overwrites.
	Administrator Permission

	// Allows management and editing of channels.
	ManageChannels Permission

	// Allows management and editing of the guild.
	ManageGuild Permission

	// Allows for the addition of reactions to messages.
	AddReactions Permission

	// Allows for viewing of audit logs.
	ViewAuditLog Permission

	// Allows for using priority speaker in a voice channel.
	PrioritySpeaker Permission

	// Allows the user to go live.
	Stream Permission

	// Allows guild members to view a channel, which includes reading messages in text channels and joining voice channels.
	ViewChannel Permission

	// Allows for sending messages in a channel and creating threads in a forum (does not allow sending messages in threads).
	SendMessages Permission

	// Allows for sending of /tts messages.
	SendTTSMessages Permission

	// Allows for deletion of other users messages.
	ManageMessages Permission

	// Links sent by users with this permission will be auto-embedded.
	EmbedLinks Permission

	// Allows for uploading images and files.
	AttachFiles Permission

	// Allows for reading of message history.
	ReadMessageHistory Permission

	// Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel.
	MentionEveryone Permission

	// Allows the usage of custom emojis from other servers.
	UseExternalEmojis Permission

	// Allows for viewing guild insights.
	ViewGuildInsights Permission

	// Allows for joining of a voice channel.
	Connect Permission

	// Allows for speaking in a voice channel.
	Speak Permission

	// Allows for muting members in a voice channel.
	MuteMembers Permission

	// Allows for deafening of members in a voice channel.
	DeafenMembers Permission

	// Allows for moving of members between voice channels.
	MoveMembers Permission

	// Allows for using voice-activity-detection in a voice channel.
	UseVAD Permission

	// Allows for modification of own nickname.
	ChangeNickname Permission

	// Allows for modification of other users nicknames.
	ManageNicknames Permission

	// Allows management and editing of roles.
	ManageRoles Permission

	// Allows management and editing of webhooks.
	ManageWebhooks Permission

	// Allows for editing and deleting emojis, stickers, and soundboard sounds created by all users.
	ManageGuildExpressions Permission

	// Allows members to use application commands, including slash commands and context menu commands..
	UseApplicationCommands Permission

	// Allows for requesting to speak in stage channels. (This permission is under active development and may be changed or removed.).
	RequestToSpeak Permission

	// Allows for editing and deleting scheduled events created by all users.
	ManageEvents Permission

	// Allows for deleting and archiving threads, and viewing all private threads.
	ManageThreads Permission

	// Allows for creating public and announcement threads.
	CreatePublicThreads Permission

	// Allows for creating private threads.
	CreatePrivateThreads Permission

	// Allows the usage of custom stickers from other servers.
	UseExternalStickers Permission

	// Allows for sending messages in threads.
	SendMessagesInThreads Permission

	// Allows for using Activities (applications with the EMBEDDED flag) in a voice channel.
	UseEmbeddedActivities Permission

	// Allows for timing out users to prevent them from sending or reacting to messages in chat and threads, and from speaking in voice and stage channels.
	ModerateMembers Permission

	// Allows for viewing role subscription insights.
	ViewCreatorMonetizationAnalytics Permission

	// Allows for using soundboard in a voice channel.
	UseSoundboard Permission

	// Allows for creating emojis, stickers, and soundboard sounds, and editing and deleting those created by the current user. Not yet available to developers, see changelog..
	CreateGuildExpressions Permission

	// Allows for creating scheduled events, and editing and deleting those created by the current user. Not yet available to developers, see changelog..
	CreateEvents Permission

	// Allows the usage of custom soundboard sounds from other servers.
	UseExternalSounds Permission

	// Allows sending voice messages.
	SendVoiceMessages Permission
} {
	CreateInstantInvite: 1 << 0,
	KickMembers: 1 << 1,
	BanMembers: 1 << 2,
	Administrator: 1 << 3,
	ManageChannels: 1 << 4,
	ManageGuild: 1 << 5,
	AddReactions: 1 << 6,
	ViewAuditLog: 1 << 7,
	PrioritySpeaker: 1 << 8,
	Stream: 1 << 9,
	ViewChannel: 1 << 10,
	SendMessages: 1 << 11,
	SendTTSMessages: 1 << 12,
	ManageMessages: 1 << 13,
	EmbedLinks: 1 << 14,
	AttachFiles: 1 << 15,
	ReadMessageHistory: 1 << 16,
	MentionEveryone: 1 << 17,
	UseExternalEmojis: 1 << 18,
	ViewGuildInsights: 1 << 19,
	Connect: 1 << 20,
	Speak: 1 << 21,
	MuteMembers: 1 << 22,
	DeafenMembers: 1 << 23,
	MoveMembers: 1 << 24,
	UseVAD: 1 << 25,
	ChangeNickname: 1 << 26,
	ManageNicknames: 1 << 27,
	ManageRoles: 1 << 28,
	ManageWebhooks: 1 << 29,
	ManageGuildExpressions: 1 << 30,
	UseApplicationCommands: 1 << 31,
	RequestToSpeak: 1 << 32,
	ManageEvents: 1 << 33,
	ManageThreads: 1 << 34,
	CreatePublicThreads: 1 << 35,
	CreatePrivateThreads: 1 << 36,
	UseExternalStickers: 1 << 37,
	SendMessagesInThreads: 1 << 38,
	UseEmbeddedActivities: 1 << 39,
	ModerateMembers: 1 << 40,
	ViewCreatorMonetizationAnalytics: 1 << 31,
	UseSoundboard: 1 << 42,
	CreateGuildExpressions: 1 << 43,
	CreateEvents: 1 << 44,
	UseExternalSounds: 1 << 45,
	SendVoiceMessages: 1 << 46,
}

