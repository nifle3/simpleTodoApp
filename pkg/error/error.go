package error

import "net/http"

type Http struct {
	StatusCode int
	message    string
}

func (h Http) Error() string {
	return h.message
}

func InternalServer(message string) Http {
	return Http{
		StatusCode: http.StatusInternalServerError,
		message:    message,
	}
}

func OK(message string) Http {
	return Http{
		StatusCode: http.StatusOK,
		message:    message,
	}
}
