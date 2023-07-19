// Code generated by goctl. DO NOT EDIT.
package types

type NilReq struct {
}

type NilReply struct {
}

type HealthReply struct {
	Message string `json:"message"`
}

type UserInfoReq struct {
	UserID string `json:"userId"`
}

type UserInfoReply struct {
	UserID string `json:"userId"`
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
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}