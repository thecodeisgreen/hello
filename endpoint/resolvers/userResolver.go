package resolvers

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserResolverHandler struct {
	sessionID string
}

func UserResolver() *UserResolverHandler {
	return &UserResolverHandler{}
}

func (handler *UserResolverHandler) User() User {
	return User{
		Email: "hubert.bettan@gmail.com",
	}
}

/*
 *
 * Mutations
 *
 */

// SignInMutationResponse is the return type of SignInMutation
type SignInMutationResponse struct {
	Ok bool `json:"ok"`
}

func (handler *UserResolverHandler) SignIn(email string, password string) (SignInMutationResponse, error) {
	return SignInMutationResponse{
		Ok: true,
	}, nil
}

// SignUoMutationResponse is the return type of SignUpMutation
type SignUpMutationResponse struct {
	Ok bool `json:"ok"`
}

func (handler *UserResolverHandler) SignUp(email string, password string) (SignUpMutationResponse, error) {
	return SignUpMutationResponse{
		Ok: true,
	}, nil
}
