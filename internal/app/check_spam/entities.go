package check_spam

type GetPermissionsReq struct {
	SenderID    uint64 `json:"sender_id"`
	MessageText string `json:"message_text"`
}

type GetPermissionsRes struct {
	Permission string `json:"permission"`
	Reason     string `json:"reason"`
}
