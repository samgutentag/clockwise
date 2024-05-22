package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/alecsavvy/clockwise/core/interface/commands"
	"github.com/alecsavvy/clockwise/core/interface/entities"
	"github.com/alecsavvy/clockwise/ports/graph/model"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/google/uuid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	core := r.core

	createUserCmd := commands.NewCommand(commands.USER, commands.CREATE, commands.CreateUser{
		ID:      uuid.NewString(),
		Bio:     input.Bio,
		Handle:  input.Handle,
		Address: input.Address,
	})

	userEntity, err := core.CreateUser(createUserCmd)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		ID:      userEntity.ID,
		Handle:  userEntity.Handle,
		Bio:     userEntity.Bio,
		Address: userEntity.Address,
	}

	return newUser, nil
}

// CreateTrack is the resolver for the createTrack field.
func (r *mutationResolver) CreateTrack(ctx context.Context, input model.NewTrack) (*model.Track, error) {
	core := r.core

	createTrackCmd := commands.NewCommand(commands.TRACK, commands.CREATE, commands.CreateTrack{
		ID:          uuid.NewString(),
		Title:       input.Title,
		Description: input.Description,
		StreamURL:   input.StreamURL,
		Genre:       input.Genre,
		UserID:      input.UserID,
	})

	trackEntity, err := core.CreateTrack(createTrackCmd)
	if err != nil {
		return nil, err
	}

	newTrack := &model.Track{
		ID:          trackEntity.ID,
		Title:       trackEntity.Title,
		Description: trackEntity.Description,
		StreamURL:   trackEntity.StreamURL,
		Genre:       trackEntity.Genre,
		UserID:      trackEntity.UserID,
	}

	return newTrack, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// UpdateTrack is the resolver for the updateTrack field.
func (r *mutationResolver) UpdateTrack(ctx context.Context, input model.UpdateTrack) (*model.Track, error) {
	panic(fmt.Errorf("not implemented: UpdateTrack - updateTrack"))
}

// Follow is the resolver for the follow field.
func (r *mutationResolver) Follow(ctx context.Context, input model.NewFollow) (*model.Follow, error) {
	core := r.core

	createFollowCommand := commands.NewCommand(commands.FOLLOW, commands.CREATE, commands.CreateFollow{
		ID:          uuid.NewString(),
		FollowerID:  input.FollowerID,
		FollowingID: input.FollowingID,
	})

	followEntity, err := core.CreateFollow(createFollowCommand)
	if err != nil {
		return nil, err
	}

	newFollow := &model.Follow{
		ID:          followEntity.ID,
		FollowerID:  followEntity.FollowerID,
		FollowingID: followEntity.FollowingID,
	}

	return newFollow, nil
}

// Repost is the resolver for the repost field.
func (r *mutationResolver) Repost(ctx context.Context, input model.NewRepost) (*model.Repost, error) {
	panic(fmt.Errorf("not implemented: Repost - repost"))
}

// Unfollow is the resolver for the unfollow field.
func (r *mutationResolver) Unfollow(ctx context.Context, followID string) (string, error) {
	panic(fmt.Errorf("not implemented: Unfollow - unfollow"))
}

// Unrepost is the resolver for the unrepost field.
func (r *mutationResolver) Unrepost(ctx context.Context, repostID string) (string, error) {
	panic(fmt.Errorf("not implemented: Unrepost - unrepost"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	core := r.core

	users, err := core.GetUsers()
	if err != nil {
		return nil, err
	}

	userModels := utils.Map(users, func(userEntity *entities.UserEntity) *model.User {
		return &model.User{
			ID:      userEntity.ID,
			Handle:  userEntity.Handle,
			Bio:     userEntity.Bio,
			Address: userEntity.Address,
		}
	})
	return userModels, nil
}

// Tracks is the resolver for the tracks field.
func (r *queryResolver) Tracks(ctx context.Context) ([]*model.Track, error) {
	core := r.core

	tracks, err := core.GetTracks()
	if err != nil {
		return nil, err
	}

	trackModels := utils.Map(tracks, func(trackEntity *entities.TrackEntity) *model.Track {
		return &model.Track{
			ID:          trackEntity.ID,
			Title:       trackEntity.Title,
			Description: trackEntity.Description,
			StreamURL:   trackEntity.StreamURL,
			Genre:       trackEntity.Genre,
			UserID:      trackEntity.UserID,
		}
	})

	return trackModels, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	c := r.core

	user, err := c.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return r.userEntitiesToModels([]*entities.UserEntity{user})[0], nil
}

// Track is the resolver for the track field.
func (r *queryResolver) Track(ctx context.Context, trackID string) (*model.Track, error) {
	panic(fmt.Errorf("not implemented: Track - track"))
}

// Followers is the resolver for the followers field.
func (r *queryResolver) Followers(ctx context.Context, userID string) ([]*model.Follow, error) {
	c := r.core

	followers, err := c.GetUserFollowers(userID)
	if err != nil {
		return nil, err
	}

	return r.followEntitiesToModels(followers), nil
}

// Following is the resolver for the following field.
func (r *queryResolver) Following(ctx context.Context, userID string) ([]*model.Follow, error) {
	c := r.core

	following, err := c.GetUserFollowing(userID)
	if err != nil {
		return nil, err
	}

	return r.followEntitiesToModels(following), nil
}

// Reposts is the resolver for the reposts field.
func (r *queryResolver) Reposts(ctx context.Context, trackID string) ([]*model.Repost, error) {
	panic(fmt.Errorf("not implemented: Reposts - reposts"))
}

// UserEvents is the resolver for the userEvents field.
func (r *subscriptionResolver) UserEvents(ctx context.Context, userID string) (<-chan model.UserEvents, error) {
	ues := make(chan model.UserEvents)
	go func() {
		defer close(ues)
		userEvents := r.core.Pubsub().UserPubsub.Subscribe()
		for {
			select {
			case entity, ok := <-userEvents:
				if !ok {
					return
				}
				user := model.User{
					ID:      entity.ID,
					Handle:  entity.Handle,
					Address: entity.Address,
					Bio:     entity.Bio,
				}
				ues <- user
			case <-ctx.Done():
				return // Exit the goroutine if context is canceled
			}
		}
	}()
	return ues, nil
}

// TrackEvents is the resolver for the trackEvents field.
func (r *subscriptionResolver) TrackEvents(ctx context.Context, trackID string) (<-chan model.TrackEvents, error) {
	tes := make(chan model.TrackEvents)
	go func() {
		defer close(tes)
		trackEvents := r.core.Pubsub().TrackPubsub.Subscribe()
		for {
			select {
			case entity, ok := <-trackEvents:
				if !ok {
					return
				}
				track := model.Track{
					ID:          entity.ID,
					Title:       entity.Title,
					Description: entity.Description,
					Genre:       entity.Genre,
					StreamURL:   entity.StreamURL,
					UserID:      entity.UserID,
				}
				tes <- track
			case <-ctx.Done():
				return // Exit the goroutine if context is canceled
			}
		}
	}()
	return tes, nil
}

// GenreEvents is the resolver for the genreEvents field.
func (r *subscriptionResolver) GenreEvents(ctx context.Context, genre string) (<-chan model.TrackEvents, error) {
	panic(fmt.Errorf("not implemented: GenreEvents - genreEvents"))
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Follows(ctx context.Context) ([]*model.Follow, error) {
	panic(fmt.Errorf("not implemented: Follows - follows"))
}
func (r *subscriptionResolver) NewUser(ctx context.Context) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented: NewUser - newUser"))
}
func (r *subscriptionResolver) NewTrack(ctx context.Context) (<-chan *model.Track, error) {
	panic(fmt.Errorf("not implemented: NewTrack - newTrack"))
}
