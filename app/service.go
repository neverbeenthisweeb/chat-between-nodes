package app

import "context"

type ChatService interface {
	GetByNodeID(ctx context.Context, ID string) ([]Chat, error)
}

type chatServiceImpl struct {
	repo ChatRepository
}

func NewChatServiceImpl(repo ChatRepository) *chatServiceImpl {
	return &chatServiceImpl{repo}
}

func (cu *chatServiceImpl) GetByNodeID(ctx context.Context, ID string) ([]Chat, error) {
	chats, err := cu.repo.GetByNodeID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return chats, nil
}
