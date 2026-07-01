package response

const (
	ErrCodeSuccess      = 20001 // sucess
	ErrCodeParamInvalid = 20003 // email is invalid
	ErrInvalidToken     = 30001 // invalid token

)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken:     "invalid token",
}
