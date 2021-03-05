package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/kalradev/review-central/contracts"
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
	err := user.Create()
	if err != nil {
		return "", err
	}
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
	// get user
	cookie := auth.ForContext(ctx)

	//validate jwt token
	email, err := jwt.ParseToken(cookie)
	if err != nil {
		return "", errors.New("Access Denied")
	}
	// create user and check if user exists in db
	user, err := users.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("Access Denied")
	}

	// Where your local node is running on localhost:5001
	sh := shell.NewShell("https://ipfs.infura.io:5001")
	jsonifiedReview, err := json.Marshal(input)
	if err != nil {
		log.Printf("error: %s", err)
		return "", err
	}
	hash, err := sh.Add(bytes.NewReader(jsonifiedReview))
	if err != nil {
		log.Printf("error: %s", err)
		return "", err
	}

	err = contracts.Put(input.Token, []string{hash})
	if err != nil {
		log.Printf("error: %s", err)
		return "", err
	}
	err = user.AddToken(input.Token)

	// store token, hash on blockchain
	// and store token, user in database
	return fmt.Sprintf("https://ipfs.infura.io/ipfs/%s", hash), nil
}

func (r *queryResolver) Reviews(ctx context.Context, input *model.UserInfo) ([]*model.Review, error) {
	var reviews []*model.Review
	var myClient = &http.Client{Timeout: 10 * time.Second}

	if input != nil && input.CurrentUser != nil && *input.CurrentUser == false {
		tokens, err := contracts.GetTokens()
		if err != nil {
			return nil, errors.New("Coudn't get review tokens")
		}

		for _, token := range tokens {
			// get hash of every token
			hashes, err := contracts.Get(token)
			if err != nil {
				return nil, errors.New("Coudn't get ipfs hash")
			}
			for _, hash := range hashes {
				r, err := myClient.Get(fmt.Sprintf("https://ipfs.infura.io/ipfs/%s", hash))
				if err != nil {
					log.Println("error getting response from ipfs")
				}
				defer r.Body.Close()

				var rev *model.Review
				json.NewDecoder(r.Body).Decode(rev)
				reviews = append(reviews, rev)
			}
		}

	} else {
		// get user
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

		tokens, err := user.GetTokens()
		if err != nil {
			return nil, errors.New("Coudn't get tokens")
		}
		log.Println(tokens)

		for _, token := range tokens {
			// get hash of every token
			hashes, err := contracts.Get(token)
			log.Println(hashes)
			if err != nil {
				return nil, errors.New("Coudn't get ipfs hash")
			}
			for _, hash := range hashes {
				r, err := myClient.Get(fmt.Sprintf("https://ipfs.infura.io/ipfs/%s", hash))
				if err != nil {
					log.Println("error getting response from ipfs")
				}
				defer r.Body.Close()

				var rev model.Review
				json.NewDecoder(r.Body).Decode(&rev)
				reviews = append(reviews, &rev)
			}
		}

	}

	return reviews, nil
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
