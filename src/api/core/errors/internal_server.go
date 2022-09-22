package errors

type InternalServer struct {
	message string
}

func NewInternalServer(message string) *InternalServer {
	return &InternalServer{
		message: message,
	}
}

func (e *InternalServer) Error() string {
	return e.message
}
