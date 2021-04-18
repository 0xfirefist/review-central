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
	cookie, err := auth.ForContext(ctx)
	if err != nil {
		return "", errors.New("Cookie not found, access denied")
	}

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

	// token should not be used previously
	valid, err := user.IsTokenValid(input.Token)
	if valid == true {
		return "", errors.New("cannot reuse a token add offset review instead")
	}

	// token should be verify by seller

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

func (r *mutationResolver) AddOffset(ctx context.Context, input *model.ReviewInput) (string, error) {
	// get user
	cookie, err := auth.ForContext(ctx)
	if err != nil {
		return "", errors.New("Cookie not found, access denied")
	}

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

	// check if user has the token to which it want to add offset review
	valid, err := user.IsTokenValid(input.Token)
	if err != nil || valid == false {
		return "", errors.New("Couldn't verify token")
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

	return fmt.Sprintf("https://ipfs.infura.io/ipfs/%s", hash), nil
}

func (r *queryResolver) GetReviews(ctx context.Context, input *model.UserInfo) ([]*model.AssociatedReview, error) {
	var associatedReviews []*model.AssociatedReview
	var myClient = &http.Client{Timeout: 10 * time.Second}
	var tokens []string
	var err error

	// general reviews
	if input != nil && input.CurrentUser != nil && *input.CurrentUser == false {
		tokens, err = contracts.GetTokens()
		if err != nil {
			return nil, errors.New("Couldn't get review tokens")
		}
	} else {
		// get user
		cookie, err := auth.ForContext(ctx)
		if err != nil {
			return nil, errors.New("Cookie not found, access denied")
		}

		//validate jwt token
		email, err := jwt.ParseToken(cookie)
		if err != nil {
			return nil, errors.New("Access Denied")
		}
		// create user and check if user exists in db
		user, err := users.GetUserByEmail(email)
		if err != nil {
			return nil, errors.New("Access Denied")
		}

		tokens, err = user.GetTokens()
		if err != nil {
			return nil, errors.New("Coudn't get tokens")
		}
	}

	for _, token := range tokens {
		var reviews []*model.Review

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
			// log.Println(hash)
			var rev model.Review
			json.NewDecoder(r.Body).Decode(&rev)
			reviews = append(reviews, &rev)
		}

		// get product details
		productDetails,err := users.GetProduct(token)
		if err!=nil {
			associatedReviews = append(associatedReviews, &model.AssociatedReview{Token: &token,
				Reviews: reviews,
			}) 
		} else {
			associatedReviews = append(associatedReviews, &model.AssociatedReview{Token: &token,
				Reviews: reviews,
				Product: &model.Product{
					Name:         productDetails[0],
					Manufacturer: productDetails[1],
					Model:        productDetails[2],
					Vendor:       productDetails[3],
				},
			})
		}
	}

	return associatedReviews, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	cookie, err := auth.ForContext(ctx)
	if err != nil {
		return nil, errors.New("Cookie not found, access denied")
	}

	//validate jwt token
	email, err := jwt.ParseToken(cookie)
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
