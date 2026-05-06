package middlewares

import (
	"fmt"

	"github.com/alfasya/imgo/utils"
)

func VerifyAuth(t string) error {
	err := utils.VerifyJWT(t)
	if err != nil {
		fmt.Println("Error verifying token: %v", err)
		return err
	}

	return nil
}
