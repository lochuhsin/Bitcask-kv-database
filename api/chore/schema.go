package chore

type rootResponseSchema struct {
	Message string `json:"message"`
}

type healthzResponseSchema struct {
	Message string `json:"message"`
}
