package error_handler

import "log"

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (eh *ErrorHandler) HandleError(errorMessage string) error {
	log.Printf("Error: %s", errorMessage)
	// You can log the error or take other appropriate actions
	return fmt.Errorf("Error: %s", errorMessage)
}