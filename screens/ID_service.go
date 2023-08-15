package screens

type idServies struct {
	messageID int
	chatID    int64
}

func (ids *idServies) GetMessageID() int {
	return ids.messageID
}

func (ids *idServies) setMessageID(messageID int) {
	ids.messageID = messageID
}

func (ids *idServies) GetChatID() int64 {
	return ids.chatID
}

func (ids *idServies) setChatID(chatID int64) {
	ids.chatID = chatID
}
