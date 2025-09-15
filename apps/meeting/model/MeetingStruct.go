package model

type RedisMeetingStruct struct {
	MeetingId       string          `json:"MeetingId"`
	MeetingName     string          `json:"MeetingName"`
	MeetingPassword string          `json:"MeetingPassword"`
	Description     string          `json:"Description"`
	HostId          string          `json:"HostId"`
	HostName        string          `json:"HostName"`
	Status          string          `json:"Status"` // 1未开始 2进行中 3已结束
	StartTime       string          `json:"StartTime"`
	EndTime         string          `json:"EndTime"`
	Members         []*SingleMember `json:"Members"`
}

type SingleMember struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

type WsMessage struct {
	MeetingId string          `json:"meetingId"`
	HostId    string          `json:"HostId"`
	Action    ActionType      `json:"action"`
	Members   []*SingleMember `json:"members"`
}
