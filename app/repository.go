package app

import (
	"context"
	"time"
)

type ChatRepository interface {
	GetByNodeID(ctx context.Context, ID string) ([]Chat, error)
}

// memChatRepositoryDB is a global DB to store chat data for each nodes
//
// Node ID of 1001 and 1002 act as existing data.
var memChatRepositoryDB = map[string][]Chat{
	"1001": {
		{
			ID:        "1",
			NodeID:    "1001",
			From:      "Spongebob",
			Message:   "Hi all!",
			Timestamp: time.Now().Round(1 * time.Hour),
		},
		{
			ID:        "2",
			NodeID:    "1001",
			From:      "Patrick",
			Message:   "Sry I'm busy",
			Timestamp: time.Now().Round(1 * time.Hour).Add(1 * time.Minute),
		},
	},
	"1002": {
		{
			ID:        "3",
			NodeID:    "1002",
			From:      "Sandy",
			Message:   "Hi, anybody here?",
			Timestamp: time.Now().Round(1 * time.Hour).Add(2 * time.Minute),
		},
		{
			ID:        "4",
			NodeID:    "1002",
			From:      "Larry",
			Message:   "Hi",
			Timestamp: time.Now().Round(1 * time.Hour).Add(3 * time.Minute),
		},
		{
			ID:        "5",
			NodeID:    "1002",
			From:      "Larry",
			Message:   "HII!!! I AM HERE!",
			Timestamp: time.Now().Round(1 * time.Hour).Add(4 * time.Minute),
		},
	},
}

// memChatRepository is an in-memory implementation of chat repository
type memChatRepository struct{}

func NewMemChatRepository() *memChatRepository {
	return &memChatRepository{}
}

func (mc *memChatRepository) GetByNodeID(ctx context.Context, ID string) ([]Chat, error) {
	if ID == "" {
		return []Chat{}, ErrMissingNodeID
	}

	v, ok := memChatRepositoryDB[ID]
	if !ok {
		return []Chat{}, ErrChatNotFound
	}

	return v, nil
}
