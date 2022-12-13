package structure

import "gorm.io/gorm"

type GroupChat struct {
	UserID     uint
	RoomChatID uint
}

type RoomChat struct {
	gorm.Model
	Name      string
	DeletedBy uint
}

type Message struct {
	gorm.Model
	RoomChatID   int
	Message      string
	Image        string
	SenderUserId int
	ReadingDate  bool
	DeletedBy    int
}
