package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MeetingEntity struct {
	MeetingId       string `gorm:"primaryKey;column:meeting_id" json:"meetingId"`
	MeetingName     string `gorm:"column:meeting_name" json:"meetingName"`
	MeetingPassword string `gorm:"column:meeting_password" json:"meetingPassword"`
	Description     string `gorm:"column:description" json:"description"`
	HostId          string `gorm:"column:host_id" json:"hostId"`
	HostName        string `gorm:"column:host_name" json:"hostName"`
	StartTime       int64  `gorm:"column:start_time" json:"startTime"`
	EndTime         int64  `gorm:"column:end_time" json:"endTime"`
}

func (MeetingEntity) TableName() string {
	return "meeting_entities"
}

type MeetingHistory struct {
	MeetingId string `gorm:"primaryKey;column:meeting_id;type:varchar(64);not null"`
	UserId    string `gorm:"primaryKey;column:user_id;type:varchar(50);not null"`

	// 关联
	MeetingEntity MeetingEntity `gorm:"foreignKey:MeetingId;references:MeetingId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MeetingHistory) TableName() string {
	return "meeting_history"
}

type BaseMeetingModel interface {
	CreateMeeting(meeting *MeetingEntity) error
	GetMeetingById(userId string) ([]*MeetingHistory, error)
	UpdateMeeting(meeting *MeetingEntity) error
	DeleteMeeting(meetingId string) error
	Transaction(fc func(tx *gorm.DB) error) error
}

type MeetingModel struct {
	WriteDb *gorm.DB
	ReadDb  *gorm.DB
}

func NewMeetingModel(readDsn, writeDsn string) *MeetingModel {
	rdb, err := gorm.Open(mysql.Open(readDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect read database: " + err.Error())
	}
	wdb, err := gorm.Open(mysql.Open(writeDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect read database: " + err.Error())
	}

	return &MeetingModel{
		ReadDb:  rdb,
		WriteDb: wdb,
	}
}

func (m *MeetingModel) CreateMeeting(meeting *MeetingEntity) error {
	return m.WriteDb.Create(meeting).Error
}

func (m *MeetingModel) GetMeetingById(userId string) ([]*MeetingHistory, error) {
	var meeting []*MeetingHistory
	if err := m.ReadDb.Where("user_id = ?", userId).Preload("MeetingEntity").Find(&meeting).Error; err != nil {
		return nil, err
	}
	return meeting, nil
}

func (m *MeetingModel) UpdateStatus(meetingId string, st int) error {
	return m.WriteDb.Model(&MeetingEntity{}).Where("meeting_id = ?", meetingId).Update("status", st).Error
}

func (m *MeetingModel) UpdateMeeting(meeting *MeetingEntity) error {
	return m.WriteDb.Save(meeting).Error
}

func (m *MeetingModel) DeleteMeeting(meetingId string) error {
	return m.WriteDb.Delete(&MeetingEntity{}, "meeting_id = ?", meetingId).Error
}

func (m *MeetingModel) Transaction(fc func(tx *gorm.DB) error) error {
	return m.WriteDb.Transaction(fc)
}
