// Code generated by goctl. DO NOT EDIT.
package types

type NilReq struct {
}

type NilReply struct {
}

type HealthReply struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	Username  string `json:"username"`
	UUID      string `json:"uuid"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updetedAt"`
}

type FileUploadReply struct {
	FileStoreName string `json:"fileStoreName"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	FileSize      int64  `json:"fileSize"`
}

type CaptchaReply struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}

type Role struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}

type CreateRoleReq struct {
	Name string `json:"name"`
}

type UpdateRoleReq struct {
	Role
}
