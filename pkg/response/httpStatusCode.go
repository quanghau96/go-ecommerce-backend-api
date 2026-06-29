package response

const (
	ErrCodeSuccess      = 20001 // sucess
	ErrCodeParamInvalid = 20003 // email is invalid

)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
}
