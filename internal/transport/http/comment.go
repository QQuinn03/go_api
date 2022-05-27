package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/QQuinn03/api_toy/internal/comment"
	"github.com/go-playground/validator"
)

type CommentService interface {
	GetComment(context.Context, string) (comment.Comment, error)
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	//GetComment(context.Context, string) (comment.Comment, error)
	//GetComment(context.Context, string) (comment.Comment, error)
}
type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertPostCommentToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmt)
	if err != nil {
		http.Error(w, "not a validate comment", http.StatusBadRequest)
		return
	}

	convertedComment := convertPostCommentToComment(cmt)
	postedComment, err := h.Service.PostComment(r.Context(), convertedComment)
	if err != nil {
		log.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	}

}
