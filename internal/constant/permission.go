package constant

const (
	// define group.
	GroupDefault   = "DEFAULT"
	GroupSuperUser = "SUPER_USER"

	// define permission.
	PermissionFullAccess = "FULL_ACCESS"
	PermissionAllowGuest = "GUEST_FULL_ACCESS"

	PermissionGroupAll    = "GROUP_ALL"
	PermissionGroupCreate = "GROUP_CREATE"
	PermissionGroupRead   = "GROUP_READ"
	PermissionGroupUpdate = "GROUP_UPDATE"
	PermissionGroupDelete = "GROUP_DELETE"

	PermissionPermissionAll    = "PERMISSION_ALL"
	PermissionPermissionCreate = "PERMISSION_CREATE"
	PermissionPermissionRead   = "PERMISSION_READ"
	PermissionPermissionUpdate = "PERMISSION_Update"
	PermissionPermissionDelete = "PERMISSION_DELETE"

	PermissionGroupPermissionAll    = "GROUP_PERMISSION_ALL"
	PermissionGroupPermissionRead   = "GROUP_PERMISSION_READ"
	PermissionGroupPermissionCreate = "GROUP_PERMISSION_CREATE"
	PermissionGroupPermissionDelete = "GROUP_PERMISSION_DELETE"

	PermissionUserGroupAll    = "USER_GROUP_ALL"
	PermissionUserGroupRead   = "USER_GROUP_READ"
	PermissionUserGroupCreate = "USER_GROUP_CREATE"
	PermissionUserGroupDelete = "USER_GROUP_DELETE"
)

var (
	SeedPermissions = []string{
		PermissionFullAccess,
		PermissionAllowGuest,
		PermissionGroupAll,
		PermissionGroupCreate,
		PermissionGroupRead,
		PermissionGroupUpdate,
		PermissionGroupDelete,
		PermissionPermissionAll,
		PermissionPermissionCreate,
		PermissionPermissionRead,
		PermissionPermissionUpdate,
		PermissionPermissionDelete,
		PermissionGroupPermissionAll,
		PermissionGroupPermissionRead,
		PermissionGroupPermissionCreate,
		PermissionGroupPermissionDelete,
		PermissionUserGroupAll,
		PermissionUserGroupRead,
		PermissionUserGroupCreate,
		PermissionUserGroupDelete,
	}
	SeedGroups = []string{
		GroupDefault,
		GroupSuperUser,
	}
	SeedGroupPermissios = map[string][]string{
		GroupDefault: {
			PermissionGroupRead,
			PermissionPermissionRead,
			PermissionGroupPermissionRead,
			PermissionUserGroupRead,
		},
		GroupSuperUser: {
			PermissionFullAccess,
			PermissionAllowGuest,
			PermissionGroupAll,
			PermissionGroupCreate,
			PermissionGroupRead,
			PermissionGroupUpdate,
			PermissionGroupDelete,
			PermissionPermissionAll,
			PermissionPermissionCreate,
			PermissionPermissionRead,
			PermissionPermissionUpdate,
			PermissionPermissionDelete,
			PermissionGroupPermissionAll,
			PermissionGroupPermissionRead,
			PermissionGroupPermissionCreate,
			PermissionGroupPermissionDelete,
			PermissionUserGroupAll,
			PermissionUserGroupRead,
			PermissionUserGroupCreate,
			PermissionUserGroupDelete,
		},
	}
)
