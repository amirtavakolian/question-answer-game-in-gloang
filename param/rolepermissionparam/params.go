package rolepermissionparam

type StorePermissionParam struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type StoreRoleParam struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type AssignPermissionToRoleParam struct {
	Role_id       int `json:"role_id"`
	Permission_id int `json:"permission_id"`
}
