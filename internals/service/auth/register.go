package auth

import (
	"context"
	"expenset/internals/storer/auth"
	"expenset/pkg/utils"
	"fmt"
	"net/http"
)

type Register interface {
	Registration(ctx context.Context, req RegistrationRequest) RegistrationResponse
}

type register struct {
	authWriter auth.Writer
	authReader auth.Reader
}

type RegistrationRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type RegistrationResponse struct {
	ID    string               `json:"id"`
	Email string               `json:"email"`
	Error *utils.ErrorResponse `json:"error"`
}

func (r *register) Registration(ctx context.Context, req RegistrationRequest) RegistrationResponse {
	// check if id already there
	existed, err := r.authReader.GetByID(ctx, req.ID)
	if err == nil {
		return RegistrationResponse{
			ID:    existed.ID,
			Email: existed.Email,
			Error: nil,
		}
	}
	// saved data to db
	data := auth.Account{
		ID:    req.ID,
		Email: req.Email,
	}
	res, err := r.authWriter.Save(ctx, data)
	fmt.Println("Makanan")
	fmt.Println(req)
	fmt.Println(existed)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		return RegistrationResponse{
			ID:    "",
			Email: "",
			Error: &utils.ErrorResponse{
				Code:  http.StatusInternalServerError,
				Error: err,
			},
		}
	}

	return RegistrationResponse{
		ID:    res.ID,
		Email: res.Email,
		Error: nil,
	}
}

func NewRegister(
	authWriter auth.Writer,
	authReader auth.Reader,
) Register {
	return &register{authWriter, authReader}
}
