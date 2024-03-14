package response

type Response[D any] interface {
	StatusCode() int
	GetBody() ([]byte, error)
	Error() string
	GetData() D
}
