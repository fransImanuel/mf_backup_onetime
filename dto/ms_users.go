package dto


type MSUsersRequestDto struct {
	FilterBaseDto
	UserId []int64 `json:"user_id"`
}
