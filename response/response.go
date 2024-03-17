package response

type Response[D any, B SuccessResponse[D] | ErrorResponse] struct {
	Body B
}
