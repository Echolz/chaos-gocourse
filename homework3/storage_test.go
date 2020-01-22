package homework3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryStorage_AdminInitialValue(t *testing.T) {
	t.Run("test initial admin creation", func(t *testing.T) {
		storage := NewStorage()

		adminUser, err := storage.GetUser(0)
		assert.Equal(t, err, nil)
		assert.Equal(t, adminUser.Role, AdminRole)

	})
}

func TestInMemoryStorage_CreateUser(t *testing.T) {
	tests := []struct {
		name       string
		storage    Storage
		user       UserRequest
		shouldFail bool
	}{
		{
			name:    "create valid user",
			storage: NewStorage(),
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
			storage: NewStorage(),
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
			err := tc.storage.CreateUser(tc.user)

			if tc.shouldFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestInMemoryStorage_Create_Delete_Update(t *testing.T) {
	storage := NewStorage()

	firstReq := UserRequest{
		Username:  "echolz",
		Password:  "1234qwer",
		Email:     "echolz@abv.bg",
		FirstName: "echolz",
		LastName:  "echolz",
		UserRole:  "user",
	}
	err := storage.CreateUser(firstReq)
	assert.Nil(t, err)

	firstUser, err := storage.GetUser(1)
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
	err = storage.UpdateUser(1, updatedUser)
	assert.Nil(t, err)

	firstUpdatedUser, err := storage.GetUser(1)
	assert.Nil(t, err)
	assert.Equal(t, updatedUser.LastName, firstUpdatedUser.FirstName)

	err = storage.DeleteUser(1)
	assert.Nil(t, err)

	_, err = storage.GetUser(1)
	assert.NotNil(t, err)
}
