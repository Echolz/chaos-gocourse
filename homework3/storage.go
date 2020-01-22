package homework3

import (
	"fmt"
	"github.com/pkg/errors"
)

type Storage interface {
	SortBy(fieldName string) []User
	CreateUser(userRequest UserRequest) error
	GetUser(userId int) (User, error)
	UpdateUser(userId int, request UserRequest) error
	DeleteUser(userId int) error
}

var UID int

type inMemoryStorage struct {
	storageMap map[int]User
}

func NewStorage() Storage {
	storage := &inMemoryStorage{storageMap: make(map[int]User)}

	err := storage.CreateUser(UserRequest{
		Username:  "admin",
		Password:  "admin",
		Email:     "admin@admin.com",
		FirstName: "admin",
		LastName:  "admin",
		UserRole:  "admin",
	})

	if err != nil {
		panic(errors.Wrap(err, "problem with default admin user"))
	}

	return storage
}

func (m inMemoryStorage) SortBy(fieldName string) []User {
	panic("implement me")
}

func (m *inMemoryStorage) CreateUser(userRequest UserRequest) error {
	newUser, err := createUserFromRequest(userRequest, UID)
	if err != nil {
		return errors.Wrap(err, "invalid user data:")
	}

	m.storageMap[UID] = newUser
	UID++

	return nil
}

func (m inMemoryStorage) GetUser(userId int) (User, error) {
	user, ok := m.storageMap[userId]
	if !ok {
		return User{}, fmt.Errorf("user with id: %d not found", userId)
	}

	return user, nil
}

func (m inMemoryStorage) UpdateUser(userId int, request UserRequest) error {
	_, ok := m.storageMap[userId]
	if !ok {
		return fmt.Errorf("user with id: %d not found", userId)
	}

	updatedUser, err := createUserFromRequest(request, userId)
	if err != nil {
		return errors.Wrap(err, "invalid user data:")
	}

	m.storageMap[userId] = updatedUser
	return nil
}

func (m inMemoryStorage) DeleteUser(userId int) error {
	if _, ok := m.storageMap[userId]; !ok {
		return fmt.Errorf("user with id: %d not found", userId)
	}

	delete(m.storageMap, userId)

	return nil
}

func createUserFromRequest(request UserRequest, userId int) (User, error) {
	newUser := User{
		ID:        userId,
		Username:  request.Username,
		Password:  request.Password,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Enabled:   true,
		Expired:   false,
		Role:      stringToRole(request.UserRole),
	}

	if newUser.Role.String() == invalidRoleString {
		return User{}, errors.New("invalid user role")
	}

	return newUser, nil
}