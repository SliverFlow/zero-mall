// Code generated by goctl. DO NOT EDIT.
package types

type NilReq struct {
}

type NilReply struct {
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token string `json:"token"`
	User  User   `json:"user"`
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

type IDReq struct {
	ID int64 `json:"ID"`
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

type UserChangeStatusReq struct {
	ID     int64 `json:"ID"`
	Status int64 `json:"status"`
}

type UserIDReq struct {
	ID int64 `json:"ID"`
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

type MenuUpdateReq struct {
	ID        int64  `json:"ID"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Sorted    int64  `json:"sorted"`
	Role      int64  `json:"role"`
	Meta      Meta   `json:"meta"`
	Status    int64  `json:"status"`
}

type Business struct {
	BusinessID string   `json:"businessId"`
	Name       string   `json:"name"`
	UUID       string   `json:"uuid"`
	Detail     string   `json:"detail"`
	Score      int64    `json:"score"`
	Image      []string `json:"image"`
	Status     int64    `json:"status"`
	CreatedAt  int64    `json:"createdAt"`
	UpdatedAt  int64    `json:"updatedAt"`
}

type BusinessListReply struct {
	List     []Business `json:"list"`
	Total    int64      `json:"total"`
	Page     int64      `json:"page"`
	PageSize int64      `json:"pageSize"`
}

type BusinessChangeStatusReq struct {
	BusinessID string `json:"businessId"`
	Status     int64  `json:"status"`
}

type BusinessIDReq struct {
	BusinessID string `json:"businessId"`
}

type BusinessUpdateReq struct {
	BusinessID string   `json:"businessId"`
	Name       string   `json:"name"`
	UUID       string   `json:"uuid"`
	Detail     string   `json:"detail"`
	Score      int64    `json:"score"`
	Image      []string `json:"image"`
	Status     int64    `json:"status"`
}

type Category struct {
	ID         int64      `json:"ID"`
	CategoryID string     `json:"categoryId"`
	ParentId   string     `json:"parentId"`
	Name       string     `json:"name"`
	Status     int64      `json:"status"`
	Type       int64      `json:"type"`
	Sorted     int64      `json:"sorted"`
	BusinessID string     `json:"businessId"`
	CreatedAt  int64      `json:"createdAt"`
	UpdatedAt  int64      `json:"updatedAt"`
	Children   []Category `json:"children"`
}

type CategoryListAllReply struct {
	List []Category `json:"list"`
}

type CategoryCreateReq struct {
	ParentId string `json:"parentId"`
	Name     string `json:"name"`
	Sorted   int64  `json:"sorted"`
	Status   int64  `json:"status"`
	Type     int64  `json:"type"`
}

type CategoryChangeStatusReq struct {
	CategoryID string `json:"categoryId"`
	Status     int64  `json:"status"`
}

type CategoryIDReq struct {
	CategoryID string `json:"categoryId"`
}

type CategoryUpdateReq struct {
	ParentId   string `json:"parentId"`
	Name       string `json:"name"`
	Sorted     int64  `json:"sorted"`
	Status     int64  `json:"status"`
	Type       int64  `json:"type"`
	CategoryID string `json:"categoryId"`
}

type IDAndPID struct {
	ParentId   string `json:"parentId"`
	CategoryID string `json:"categoryId"`
}

type CategoryBatchDelteReq struct {
	KVS []IDAndPID `json:"kvs"`
}

type Product struct {
	ProductID  string     `json:"productId"`
	BusinessID string     `json:"businessId"`
	Name       string     `json:"name"`
	Subtitle   string     `json:"subtitle"`
	Image      []string   `json:"image"`
	Detail     string     `json:"detail"`
	Price      float64    `json:"price"`
	Stock      int64      `json:"stock"`
	Status     int64      `json:"status"`
	Categories []Category `json:"categories"`
	CreatedAt  int64      `json:"createdAt"`
	UpdatedAt  int64      `json:"updatedAt"`
}

type ProductCreateReq struct {
	Name       string     `json:"name"`
	Subtitle   string     `json:"subtitle"`
	Image      []string   `json:"image"`
	Detail     string     `json:"detail"`
	Price      float64    `json:"price"`
	Stock      int64      `json:"stock"`
	Status     int64      `json:"status"`
	Categories []Category `json:"categories"`
}

type ProductListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	KeyWord  string `json:"keyWord"`
}

type ProductListReply struct {
	Page     int64     `json:"page"`
	PageSize int64     `json:"pageSize"`
	Total    int64     `json:"total"`
	List     []Product `json:"list"`
}

type ProductDeleteReq struct {
	ProductID string `json:"productId"`
}

type ProductChangeStatusReq struct {
	ProductID string `json:"productId"`
	Status    int64  `json:"status"`
}

type ProductFindReq struct {
	ProductID string `json:"productId"`
}

type ProductUpdateReq struct {
	ProductID  string     `json:"productId"`
	BusinessID string     `json:"businessId"`
	Name       string     `json:"name"`
	Subtitle   string     `json:"subtitle"`
	Image      []string   `json:"image"`
	Detail     string     `json:"detail"`
	Price      float64    `json:"price"`
	Stock      int64      `json:"stock"`
	Status     int64      `json:"status"`
	Categories []Category `json:"categories"`
}

type FileTokenReply struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	StsToken        string `json:"stsToken"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
	FilePath        string `json:"filePath"`
}

type FileTokenReq struct {
	Path string `path:"path"`
}
