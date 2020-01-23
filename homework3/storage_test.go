package homework3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryStorage_AdminInitialValue(t *testing.T) {
	t.Run("test initial admin creation", func(t *testing.T) {
		storage := newStorage()

		adminUser, err := storage.getUser(0)
		assert.Equal(t, err, nil)
		assert.Equal(t, adminUser.Role, AdminRole)

	})
}

func TestInMemoryStorage_SortBy(t *testing.T) {
	t.Run("test sorting by different fields", func(t *testing.T) {
		initialAdmR := UserRequest{
			Username:  "admin",
			Password:  "admin",
			Email:     "admin@admin.com",
			FirstName: "admin",
			LastName:  "admin",
			UserRole:  "admin",
		}
		storage := newStorageWithAdmin(initialAdmR)

		//ID1
		u1R := UserRequest{
			Username:  "",
			Password:  "",
			Email:     "",
			FirstName: "bba",
			LastName:  "",
			UserRole:  "user",
		}

		//ID2
		u2R := UserRequest{
			Username:  "",
			Password:  "",
			Email:     "",
			FirstName: "aab",
			LastName:  "",
			UserRole:  "admin",
		}

		err := storage.createUser(u1R)
		assert.Nil(t, err)

		err = storage.createUser(u2R)
		assert.Nil(t, err)

		initialAdm, err := createUserFromRequest(initialAdmR, 0)
		assert.Nil(t, err)

		u1, err := createUserFromRequest(u1R, 1)
		assert.Nil(t, err)

		u2, err := createUserFromRequest(u2R, 2)
		assert.Nil(t, err)

		byFirstName := []User{u2, initialAdm, u1}
		assert.Equal(t, storage.sortBy(byFirstNameString), byFirstName)

		byId := []User{initialAdm, u1, u2}
		assert.Equal(t, storage.sortBy(""), byId)

		//sorting is not consistent
		//byRole := []User{initialAdm, u2, u1}
		//assert.Equal(t, storage.sortBy(byRoleString), byRole)

	})
}

func TestInMemoryStorage_CreateUser(t *testing.T) {
	tests := []struct {
		name       string
		storage    storage
		user       UserRequest
		shouldFail bool
	}{
		{
			name:    "create valid user",
			storage: newStorage(),
			user: UserRequest{
				Username:  "newname",
				Password:  "1234qwer",
				Email:     "echolz@abv.bg",
				FirstName: "newname",
				LastName:  "newname",
				UserRole:  "user",
			},
			shouldFail: false,
		},
		{
			name:    "create invalid user",
			storage: newStorage(),
			user: UserRequest{
				Username:  "newname",
				Password:  "1234qwer",
				Email:     "echolz@abv.bg",
				FirstName: "newname",
				LastName:  "newname",
				UserRole:  "invalidrole",
			},
			shouldFail: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.storage.createUser(tc.user)

			if tc.shouldFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestInMemoryStorage_Create_Delete_Update(t *testing.T) {
	storage := newStorage()

	firstReq := UserRequest{
		Username:  "echolz",
		Password:  "1234qwer",
		Email:     "echolz@abv.bg",
		FirstName: "echolz",
		LastName:  "echolz",
		UserRole:  "user",
	}
	err := storage.createUser(firstReq)
	assert.Nil(t, err)

	firstUser, err := storage.getUser(1)
	assert.Nil(t, err)
	assert.Equal(t, firstReq.Email, firstUser.Email)

	updatedUser := UserRequest{
		Username:  "newname",
		Password:  "1234qwer",
		Email:     "echolz@abv.bg",
		FirstName: "newname",
		LastName:  "newname",
		UserRole:  "user",
	}
	err = storage.updateUser(1, updatedUser)
	assert.Nil(t, err)

	firstUpdatedUser, err := storage.getUser(1)
	assert.Nil(t, err)
	assert.Equal(t, updatedUser.LastName, firstUpdatedUser.FirstName)

	err = storage.deleteUser(1)
	assert.Nil(t, err)

	_, err = storage.getUser(1)
	assert.NotNil(t, err)
}
