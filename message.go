package tgbotapi

const (
	MentionMessageEntityType       MessageEntityType = "mention"
	HashtagMessageEntityType       MessageEntityType = "hashtag"
	CashtagMessageEntityType       MessageEntityType = "cashtag"
	BotCommandMessageEntityType    MessageEntityType = "bot_command"
	UrlMessageEntityType           MessageEntityType = "url"
	EmailAddressMessageEntityType  MessageEntityType = "email"
	PhoneNumberMessageEntityType   MessageEntityType = "phone_number"
	BoldMessageEntityType          MessageEntityType = "bold"
	ItalicMessageEntityType        MessageEntityType = "italic"
	UnderlineMessageEntityType     MessageEntityType = "underline"
	StrikethroughMessageEntityType MessageEntityType = "strikethrough"
	CodeMessageEntityType          MessageEntityType = "code"
	PreMessageEntityType           MessageEntityType = "pre"
	TextLinkMessageEntityType      MessageEntityType = "text_link"
	TextMentionMessageEntityType   MessageEntityType = "text_mention"
	CustomEmojiMessageEntityType   MessageEntityType = "custom_emoji"
)

type SendMessageOptions struct {
	ChatId          int64            `json:"chat_id"`
	MessageThreadId int              `json:"message_thread_id,omitempty"`
	Text            string           `json:"text"`
	ParseMode       string           `json:"parse_mode,omitempty"`
	Entities        []*MessageEntity `json:"entities,omitempty"`
	DisablePreview  bool             `json:"disable_web_page_preview,omitempty"`
	DisableNotify   bool             `json:"disable_notification,omitempty"`
	ProtectContent  bool             `json:"protect_content,omitempty"`
	ReplyToMessage  int              `json:"reply_to_message_id,omitempty"`
	AllowSending    bool             `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup     Markup           `json:"reply_markup,omitempty"`
}

type MessageEntityType string

type MessageEntity struct {
	Type          MessageEntityType `json:"type"`
	Offset        int               `json:"offset"`
	Length        int               `json:"length"`
	Url           string            `json:"url,omitempty"`
	User          *User             `json:"user,omitempty"`
	Language      string            `json:"language,omitempty"`
	CustomEmojiId string            `json:"custom_emoji_id,omitempty"`
}

type Message struct {
	Id                   int64            `json:"message_id"`
	MessageThreadId      int              `json:"message_thread_id,omitempty"`
	From                 *User            `json:"from,omitempty"`
	FenderChat           *Chat            `json:"sender_chat,omitempty"`
	Fate                 int              `json:"date"`
	Fhat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId int              `json:"forward_from_message_id,omitempty"`
	ForwardSignature     string           `json:"forward_signature,omitempty"`
	ForwardSenderName    string           `json:"forward_sender_name,omitempty"`
	ForwardDate          int              `json:"forward_date,omitempty"`
	IsTopicMessage       bool             `json:"is_topic_message,omitempty"`
	IsAutomaticForward   bool             `json:"is_automatic_forward,omitempty"`
	ReplyToMessage       *Message         `json:"reply_to_message,omitempty"`
	ViaBot               *User            `json:"via_bot,omitempty"`
	EditDate             int              `json:"edit_date,omitempty"`
	HasProtectedContent  bool             `json:"has_protected_content,omitempty"`
	MediaGroupId         string           `json:"media_group_id,omitempty"`
	AuthorSignature      string           `json:"author_signature,omitempty"`
	Text                 string           `json:"text,omitempty"`
	Entities             []*MessageEntity `json:"entities,omitempty"`
	// Animation				 *Animation           `json:"animation,omitempty"`
	Audio    *Audio       `json:"audio,omitempty"`
	Document *Document    `json:"document,omitempty"`
	Photo    []*PhotoSize `json:"photo,omitempty"`
	Sticker  *Sticker     `json:"sticker,omitempty"`
	// Story					 *Story               `json:"story,omitempty"`
	Video *Video `json:"video,omitempty"`
	// VideoNote				 *VideoNote           `json:"video_note,omitempty"`
	Voice           *Voice           `json:"voice,omitempty"`
	Caption         string           `json:"caption,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	HasMediaSpoiler bool             `json:"has_media_spoiler,omitempty"`
	Contact         *Contact         `json:"contact,omitempty"`
	// Dice					 *Dice                `json:"dice,omitempty"`
	// Game					 *Game                `json:"game,omitempty"`
	Poll                  *Poll        `json:"poll,omitempty"`
	Venue                 *Venue       `json:"venue,omitempty"`
	Location              *Location    `json:"location,omitempty"`
	NewChatMembers        []*User      `json:"new_chat_members,omitempty"`
	LeftChatMember        *User        `json:"left_chat_member,omitempty"`
	NewChatTitle          string       `json:"new_chat_title,omitempty"`
	NewChatPhoto          []*PhotoSize `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool         `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool         `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool         `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool         `json:"channel_chat_created,omitempty"`
	// MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId   int      `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId int      `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage     *Message `json:"pinned_message,omitempty"`
	// Invoice					 *Invoice             `json:"invoice,omitempty"`
	// SuccessfulPayment		 *SuccessfulPayment   `json:"successful_payment,omitempty"`
	// UserShared				 *UserShared          `json:"user_shared,omitempty"`
	// ChatShared				 *ChatShared          `json:"chat_shared,omitempty"`
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// WriteAccessAllowed		 *WriteAccessAllowed  `json:"write_access_allowed,omitempty"`
	// PassportData			 *PassportData        `json:"passport_data,omitempty"`
	// ProximityAlertTriggered	 *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	// ForumTopicCreated		 *ForumTopicCreated   `json:"forum_topic_created,omitempty"`
	// ForumTopicEdited		 *ForumTopicEdited    `json:"forum_topic_edited,omitempty"`
	// ForumTopicClosed		 *ForumTopicClosed    `json:"forum_topic_closed,omitempty"`
	// ForumTopicReopened		 *ForumTopicReopened  `json:"forum_topic_reopened,omitempty"`
	// GeneralForumTopicHidden	 *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`
	// GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`
	// VideoChatScheduled		 *VideoChatScheduled  `json:"video_chat_scheduled,omitempty"`
	// VideoChatStarted		 *VideoChatStarted    `json:"video_chat_started,omitempty"`
	// VideoChatEnded			 *VideoChatEnded      `json:"video_chat_ended,omitempty"`
	// VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	// WebAppData				 *WebAppData          `json:"web_app_data,omitempty"`
	// ReplyMarkup				 *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *Message) IsCommand() bool {
	return m.Entities != nil && len(m.Entities) > 0 && m.Entities[0].Type == BotCommandMessageEntityType
}

func (m *Message) Command() string {
	if !m.IsCommand() {
		return ""
	}
	return m.Text[m.Entities[0].Offset+1 : m.Entities[0].Length]
}
