package controller

type CustomError struct {
	Message string
}

// Error method satisfies the error interface
func (e CustomError) Error() string {
	return e.Message
}
