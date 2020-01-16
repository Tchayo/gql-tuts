package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Message struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	ShortCode    string    `gorm:"size:15;not null" json:"shortcode"`
	Number       string    `gorm:"size:20;not null" json:"status"`
	Message      string    `json:"message"`
	Scheduled    bool      `gorm:"default:false" json:"scheduled"`
	ScheduleTime time.Time `json:"schedule_time"`
	Status       string    `gorm:"size:100;null" json:"status"`
	Author       Author    `json:"author"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveMessage : create new message
func (ms *Message) SaveMessage(db *gorm.DB) (*Message, error) {
	var err error
	err = db.Debug().Model(&Message{}).Create(&ms).Error
	if err != nil {
		return &Message{}, err
	}
	return ms, nil
}

// FindMessageByID : find message by message ID
func (ms *Message) FindMessageByID(db *gorm.DB, sid uint32) (*Message, error) {
	var err error
	err = db.Debug().Model(&Message{}).Where("id = ?", sid).Take(&ms).Error
	if err != nil {
		return &Message{}, err
	}
	return ms, nil
}

func (ms *Message) FindAllMessages(db *gorm.DB) (*[]Message, error) {
	var err error
	var messages []Message
	err = db.Debug().Model(&Message{}).Limit(100).Find(&messages).Error
	if err != nil {
		return &[]Message{}, err
	}
	return &messages, err
}