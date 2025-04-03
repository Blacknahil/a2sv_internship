package domain_test

import (
	"clean_task_manager_api_tested/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserMarshalJSON(t *testing.T) {
	// Arrange: Create a sample user
	objectID := primitive.NewObjectID()
	user := domain.User{
		ID:    objectID,
		Email: "test@example.com",
	}

	// Act: Marshal the user to JSON
	jsonData, err := user.MarshalJSON()

	// Assert: Check for no errors
	assert.NoError(t, err)

	// Assert: Check the JSON output
	expectedJSON := `{"id":"` + objectID.Hex() + `","email":"test@example.com"}`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}

func TestUserMarshalJSON_EmptyID(t *testing.T) {
	// Arrange: Create a user with an empty ID
	user := domain.User{
		Email: "test@example.com",
	}

	// Act: Marshal the user to JSON
	jsonData, err := user.MarshalJSON()

	// Assert: Check for no errors
	assert.NoError(t, err)

	// Assert: Check the JSON output
	expectedJSON := `{"id":"000000000000000000000000","email":"test@example.com"}`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}
