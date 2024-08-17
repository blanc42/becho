package response

type GenericResponse struct {
	Data    *any    `json:"data"`
	Error   *string `json:"error"`
	Message *string `json:"message"`
}
