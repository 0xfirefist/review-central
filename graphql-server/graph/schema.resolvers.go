package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kalradev/review-central/graphql-server/graph/generated"
	"github.com/kalradev/review-central/graphql-server/graph/model"
	"github.com/kalradev/review-central/internal/auth"
	"github.com/kalradev/review-central/internal/users"
	"github.com/kalradev/review-central/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		FirstName:  *input.FirstName,
		MiddleName: *input.MiddleName,
		LastName:   *input.LastName,
		Email:      *input.Email,
		Password:   *input.Password,
		Number:     *input.Number,
	}
	user.Create()
	return jwt.GenerateToken(user.Email)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{
		Email:    input.Email,
		Password: input.Password,
	}

	valid := user.Authenticate()

	if valid {
		return jwt.GenerateToken(user.Email)
	}

	return "", errors.New("Couldnot validate user")
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	email, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) AddReview(ctx context.Context, input *model.ReviewInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	token := auth.ForContext(ctx)

	//validate jwt token
	email, err := jwt.ParseToken(token)
	if err != nil {
		return nil, errors.New("Access Denied")
	}
	// create user and check if user exists in db
	user, err := users.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("Access Denied")
	}
	log.Println(user)

	// if user == nil {
	// 	return nil, errors.New("User not found")
	// }

	return &model.User{
		FirstName:  &user.FirstName,
		MiddleName: &user.MiddleName,
		LastName:   &user.LastName,
		Email:      &user.Email,
		Number:     &user.Number,
	}, nil
	// return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
