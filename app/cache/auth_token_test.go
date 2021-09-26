package cache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-chat/connect"
	"go-chat/testutil"
	"testing"
)

func TestAuthToken_SetBlackList(t *testing.T) {
	client := testutil.TestRedisClient()
	s := &AuthToken{
		Redis: &connect.Redis{Client: client},
	}

	ctx := context.Background()

	err := s.SetBlackList(ctx, "test-token", 10)
	assert.NoError(t, err)
	assert.NoError(t, s.IsExistBlackList(ctx, "test-token"))
}
