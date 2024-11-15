package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bndrmrtn/go-gale"
	"github.com/bndrmrtn/my-cloud/database/models"
	"github.com/bndrmrtn/my-cloud/utils"
	"github.com/stretchr/testify/assert"
)

func Test_AuthMiddlewareSucceed(t *testing.T) {
	// Setup test database
	mockDB, err := setupTestDB()
	if err != nil {
		t.Fatal("error setting up test database", err)
	}

	sess, err := createTestUserSession(mockDB)
	if err != nil {
		t.Fatal("error creating test user", err)
	}

	// Setup test request
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Add cookie to request to simulate a session
	r.AddCookie(&http.Cookie{
		Name:  "session",
		Value: "session-token",
	})

	// Setup test framework
	app, errHandler := setupGaleServer()
	ctx := app.NewTestContext(rr, r)

	// Set test session ID
	_ = ctx.Session().Set(utils.AuthSessionKey, []byte(sess.ID))

	// Test the middleware
	err = AuthMiddleware(mockDB)(ctx)
	if err != nil {
		if err := errHandler(ctx, err); err != nil {
			t.Fatal("failed to make response", err)
		}
	}

	// Assert the response
	assert.Equal(t, http.StatusOK, rr.Code, "response should be OK")

	switch user := ctx.Get(utils.RequestAuthUserKey).(type) {
	case *models.User:
		assert.Equal(t, user.ID, sess.HasUser.UserID, "user should be the same")
	default:
		t.Fatal("user should be a pointer to models.User")
	}
}

func Test_AuthMiddlewareFail(t *testing.T) {
	// Setup test database
	mockDB, err := setupTestDB()
	if err != nil {
		t.Fatal("error setting up test database", err)
	}

	// Setup test request
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Setup test framework
	app, errHandler := setupGaleServer()
	ctx := app.NewTestContext(rr, r)

	// Test the middleware
	err = AuthMiddleware(mockDB)(ctx)
	if err != nil {
		if err := errHandler(ctx, err); err != nil {
			t.Fatal("failed to make response", err)
		}
	}

	// Assert the response
	assert.Error(t, err, "error should be an error")
	assert.Equal(t, http.StatusUnauthorized, rr.Code, "error code should be 401 unauthorized")
	assert.IsType(t, &gale.Error{}, err, "error should be a gale.Error")
}
