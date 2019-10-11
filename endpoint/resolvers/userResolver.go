package resolvers

import (
	"context"
	"fmt"

	"hello/models/users"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserResolverHandler struct {
	SessionID string
}

func UserResolver(ctx context.Context) *UserResolverHandler {
	return &UserResolverHandler{
		SessionID: ctx.Value("sessionID").(string),
	}
}

func (handler *UserResolverHandler) User() users.User {
	user, err := users.GetOneByID(handler.SessionID)

	if err == users.ErrUserNotFound {
		return users.User{}
	}

	return *user
}

/*
 *
 * Mutations
 *
 */

// SignInMutationResponse is the return type of SignInMutation
type SignInMutationResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func (handler *UserResolverHandler) SignIn(email string, password string) (SignInMutationResponse, error) {
	user, err := users.GetOneByEmail(email)
	if err == users.ErrUserNotFound {
		return SignInMutationResponse{
			Ok:    false,
			Error: "WRONG_LOGIN_OR_PASSWORD",
		}, nil
	}

	if user.Password == password {
		return SignInMutationResponse{
			Ok: true,
		}, nil
	}

	return SignInMutationResponse{
		Ok:    false,
		Error: "WRONG_LOGIN_OR_PASSWORD",
	}, nil
}

// SignUpMutationResponse is the return type of SignUpMutation
type SignUpMutationResponse struct {
	Ok    bool       `json:"ok"`
	Error string     `json:"string"`
	User  users.User `json:"user"`
}

func (handler *UserResolverHandler) SignUp(email string, password string) (SignUpMutationResponse, error) {
	user, err := users.GetOneByEmail(email)
	if err == users.ErrUserNotFound {
		user = users.CreateOne(users.NewUser{
			Email:    email,
			Password: password,
		})
		return SignUpMutationResponse{
			Ok:   true,
			User: *user,
		}, nil
	}
	fmt.Println(user)

	return SignUpMutationResponse{
		Ok:    false,
		Error: "USER_ALREADY_EXISTS",
	}, nil
}
