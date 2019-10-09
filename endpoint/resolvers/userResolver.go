package resolvers

import "fmt"

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserResolverHandler struct {
}

type SaveUserMutationResponse struct {
	Ok bool `json:"ok"`
}

func (handler *UserResolverHandler) GetOne(ID string) (User, error) {
	return User{
		ID:    ID,
		Email: "hubert.bettan@gmail.com",
	}, nil
}

func (handler *UserResolverHandler) SaveOne(ID string, email string) (SaveUserMutationResponse, error) {
	fmt.Println(ID, email)
	return SaveUserMutationResponse{
		Ok: true,
	}, nil
}

func UserResolver() *UserResolverHandler {
	return &UserResolverHandler{}
}
