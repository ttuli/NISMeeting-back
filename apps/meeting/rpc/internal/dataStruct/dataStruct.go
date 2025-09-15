package datastruct

type MeetingInfo struct {
	MeetingId       string `json:"meetingId"`
	MeetingPassword string `json:"meetingPassword"`
	MeetingName     string `json:"meetingName"`
	Description     string `json:"description"`
	HostId          string `json:"hostId"`
	HostName        string `json:"hostName"`
	StartTime       int64  `json:"startTime"`
}
