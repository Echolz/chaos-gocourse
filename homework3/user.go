package homework3

import "fmt"

type Role int

const (
	adminRoleString   = "admin"
	userRoleString    = "user"
	invalidRoleString = "invalid"
)

func (r Role) String() string {
	switch r {
	case AdminRole:
		return adminRoleString
	case UserRole:
		return userRoleString
	}
	return invalidRoleString
}

func stringToRole(role string) Role {
	switch role {
	case adminRoleString:
		return AdminRole
	case userRoleString:
		return UserRole
	}
	return InvalidRole
}

const (
	AdminRole Role = iota
	UserRole
	InvalidRole
)

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Enabled   bool
	Expired   bool
	Role      Role
}

type UserRequest struct {
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	UserRole      string
}

func (u User) String() string {
	return fmt.Sprintf("id: %d, firstname: %s, enabled: %t, expired: %t, role: %s", u.ID, u.FirstName, u.Enabled, u.Expired, u.Role)
}
