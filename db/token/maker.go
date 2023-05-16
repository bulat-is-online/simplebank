package token

import "time"

// An interface to manage tokens
type Maker interface {
	//Takes username and duration and creates a token
	CreateToken(username string, duration time.Duration) (string, error)
	//check if token valid and send back payload stored inside of the body of token
	VerifyToken(token string) (*Payload, error)
}
