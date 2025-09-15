package model

type ActionType = int

const (
	LEFT ActionType = iota
	JOIN
	CREATE
)

type MeetingMessage struct {
	UserId string             `json:"userId"`
	Action ActionType         `json:"action"`
	Data   *RedisMeetingStruct `json:"data"`
}
