package homework3

import "fmt"

type Role int

const (
	AdminRoleString   = "admin"
	UserRoleString    = "user"
	invalidRoleString = "invalid"

	ByFirstNameString = "firstName"
	ByRoleString      = "role"
)

var byDefault = func(u1, u2 User) bool {
	return u1.ID < u2.ID
}

var byFirstName = func(u1, u2 User) bool {
	return u1.FirstName < u2.FirstName
}

var byRole = func(u1, u2 User) bool {
	return u1.Role < u2.Role
}

func stringToSorterFunc(fieldName string) by {
	switch fieldName {
	case ByFirstNameString:
		return byFirstName
	case ByRoleString:
		return byRole
	}
	return byDefault
}

func (r Role) String() string {
	switch r {
	case AdminRole:
		return AdminRoleString
	case UserRole:
		return UserRoleString
	}
	return invalidRoleString
}

func stringToRole(role string) Role {
	switch role {
	case AdminRoleString:
		return AdminRole
	case UserRoleString:
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
	UserRole  string
}

type by func(user1, user2 User) bool

type userSorter struct {
	users []User
	by    by
}

func (u userSorter) Len() int {
	return len(u.users)
}

func (u userSorter) Less(i, j int) bool {
	return u.by(u.users[i], u.users[j])
}

func (u userSorter) Swap(i, j int) {
	u.users[i], u.users[j] = u.users[j], u.users[i]
}

func (u User) String() string {
	return fmt.Sprintf("id: %d, firstname: %s, enabled: %t, expired: %t, role: %s", u.ID, u.FirstName, u.Enabled, u.Expired, u.Role)
}
