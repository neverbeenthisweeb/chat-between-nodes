package app_test

import (
	"chatnodes/app"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMemChatRepository_GetByNodeID(t *testing.T) {
	cases := []struct {
		name   string
		ctx    context.Context
		nodeID string
		chats  []app.Chat
		err    error
	}{
		{
			name:   "Missing node ID",
			ctx:    context.Background(),
			nodeID: "",
			chats:  []app.Chat{},
			err:    app.ErrMissingNodeID,
		},
		{
			name:   "Chat not found",
			ctx:    context.Background(),
			nodeID: "NOT_EXISTING_NODE_ID",
			chats:  []app.Chat{},
			err:    app.ErrChatNotFound,
		},
		{
			name:   "OK",
			ctx:    context.Background(),
			nodeID: "1001",
			chats: []app.Chat{
				{
					ID:        "1",
					NodeID:    "1001",
					From:      "Spongebob",
					Message:   "Hi all!",
					Timestamp: time.Now().Round(time.Hour),
				},
				{
					ID:        "2",
					NodeID:    "1001",
					From:      "Patrick",
					Message:   "Sry I'm busy",
					Timestamp: time.Now().Round(time.Hour).Add(1 * time.Minute),
				},
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		repo := app.NewMemChatRepository()
		chats, err := repo.GetByNodeID(tc.ctx, tc.nodeID)
		require.Equal(t, tc.chats, chats)
		require.Equal(t, tc.err, err)
	}
}
