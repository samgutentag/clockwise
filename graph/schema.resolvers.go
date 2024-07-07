package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/graph/model"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"github.com/google/uuid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	msg := &gen.CreateUser{
		Headers: &gen.Headers{
			Signature:   "sig",
			MessageType: gen.MessageType_MESSAGE_TYPE_CREATE_USER,
		},
		Data: &gen.CreateUser_Data{
			Handle:  input.Handle,
			Address: input.Address,
			Bio:     input.Bio,
		},
	}

	err := core.SendTx(r.logger, r.core.Rpc(), msg)
	if err != nil {
		return nil, err
	}

	resData := msg.GetData()

	user := &model.User{
		Handle:  resData.Handle,
		Address: resData.Address,
		Bio:     resData.Bio,
	}

	return user, nil
}

// CreateTrack is the resolver for the createTrack field.
func (r *mutationResolver) CreateTrack(ctx context.Context, input model.NewTrack) (*model.Track, error) {
	msg := &gen.CreateTrack{
		Headers: &gen.Headers{
			Signature:   "sig",
			MessageType: gen.MessageType_MESSAGE_TYPE_CREATE_TRACK,
		},
		Data: &gen.CreateTrack_Data{
			Id:          uuid.NewString(),
			Title:       input.Title,
			StreamUrl:   input.StreamURL,
			Description: input.Description,
			UserId:      input.UserID,
		},
	}

	err := core.SendTx(r.logger, r.core.Rpc(), msg)
	if err != nil {
		return nil, err
	}

	resData := msg.GetData()

	track := &model.Track{
		ID:          resData.Id,
		Title:       resData.Title,
		StreamURL:   resData.StreamUrl,
		Description: resData.Description,
		UserID:      resData.UserId,
	}

	return track, nil
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, input model.NewFollow) (*model.Follow, error) {
	panic(fmt.Errorf("not implemented: FollowUser - followUser"))
}

// RepostTrack is the resolver for the repostTrack field.
func (r *mutationResolver) RepostTrack(ctx context.Context, input model.NewRepost) (*model.Repost, error) {
	panic(fmt.Errorf("not implemented: RepostTrack - repostTrack"))
}

// UnfollowUser is the resolver for the unfollowUser field.
func (r *mutationResolver) UnfollowUser(ctx context.Context, input model.NewUnfollow) (bool, error) {
	panic(fmt.Errorf("not implemented: UnfollowUser - unfollowUser"))
}

// UnrepostTrack is the resolver for the unrepostTrack field.
func (r *mutationResolver) UnrepostTrack(ctx context.Context, input model.NewUnrepost) (bool, error) {
	panic(fmt.Errorf("not implemented: UnrepostTrack - unrepostTrack"))
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUsers - getUsers"))
}

// GetTracks is the resolver for the getTracks field.
func (r *queryResolver) GetTracks(ctx context.Context) ([]*model.Track, error) {
	panic(fmt.Errorf("not implemented: GetTracks - getTracks"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, handle string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetTrack is the resolver for the getTrack field.
func (r *queryResolver) GetTrack(ctx context.Context, title string) (*model.Track, error) {
	panic(fmt.Errorf("not implemented: GetTrack - getTrack"))
}

// Tracks is the resolver for the tracks field.
func (r *subscriptionResolver) Tracks(ctx context.Context) (<-chan *model.Track, error) {
	panic(fmt.Errorf("not implemented: Tracks - tracks"))
}

// Users is the resolver for the users field.
func (r *subscriptionResolver) Users(ctx context.Context) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Follows is the resolver for the follows field.
func (r *subscriptionResolver) Follows(ctx context.Context) (<-chan *model.Follow, error) {
	panic(fmt.Errorf("not implemented: Follows - follows"))
}

// Reposts is the resolver for the reposts field.
func (r *subscriptionResolver) Reposts(ctx context.Context) (<-chan *model.Repost, error) {
	panic(fmt.Errorf("not implemented: Reposts - reposts"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
