package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/zikster3262/graph-api/graph/generated"
	"github.com/zikster3262/graph-api/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	id := uuid.New()

	video := model.Video{
		ID:     id.String(),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: input.UserID, Name: "user" + input.UserID},
	}

	r.videos = append(r.videos, video)
	return &video, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var w []*model.Video

	for _, v := range r.videos {
		w = append(w, &v)
	}
	return w, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.User, error) {
	var a []*model.User
	for _, v := range r.videos {
		a = append(a, v.Author)
	}
	return a, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
