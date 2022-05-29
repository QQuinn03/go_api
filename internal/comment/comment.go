package comment

import (
	"context"
	"fmt"
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type CommentStore interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

type Service struct {
	commentStore CommentStore
}

func NewService(store CommentStore) *Service {
	return &Service{commentStore: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment...")
	var cmt Comment
	cmt, err := s.commentStore.GetComment(ctx, id)
	if err != nil {
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("creating a comment...")
	cmt, err := s.commentStore.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("deleting a comment...")
	err := s.commentStore.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println("failing to delete the comment", err)
		return err
	}
	return nil
}
func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	fmt.Println("updating a comment...")
	cmt, err := s.commentStore.UpdateComment(ctx, id, cmt)
	if err != nil {
		return Comment{}, err
	}
	return cmt, nil
}
