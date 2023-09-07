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

type PageReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	KeyWord  string `json:"keyWord"`
}

type User struct {
	ID        int64  `json:"ID"`
	UUID      string `json:"uuid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Role      int64  `json:"role"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Phone     string `json:"phone"`
}

type UserInfoReply struct {
	User User `json:"user"`
}

type AdminChangeRoleReq struct {
	Username string `json:"username"`
	Role     int64  `json:"role"`
}

type UserListReply struct {
	User     []User `json:"user"`
	Total    int64  `json:"total"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
}

type Role struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}

type RoleIDReq struct {
	ID int64 `json:"ID"`
}

type CreateRoleReq struct {
	Name string `json:"name"`
}

type UpdateRoleReq struct {
	Role
}

type Menu struct {
	ID        int64  `json:"ID"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Sorted    int64  `json:"sorted"`
	Role      int64  `json:"role"`
	Meta      Meta   `json:"meta"`
	Status    int64  `json:"status"`
	Children  []Menu `json:"children"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type MenuListByRoleReply struct {
	List []Menu `json:"list"`
}

type MenuChangeStatusReq struct {
	ID     int64 `json:"ID"`
	PID    int64 `json:"pid"`
	Status int64 `json:"status"`
}

type MenuCreateReq struct {
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Sorted    int64  `json:"sorted"`
	Role      int64  `json:"role"`
	Meta      Meta   `json:"meta"`
	Status    int64  `json:"status"`
}
