package errorsx

type BadRequestError struct {
	Err error `json:"-"` // The original error
}

// NewBadRequestError create new error based on internal error.
func NewBadRequestError(err error) error {
	return BadRequestError{err}
}

// Implements the error.Error interface
func (err BadRequestError) Error() string {
	return err.Err.Error()
}

// Implements the errors.Unwrap interface
func (err BadRequestError) Unwrap() error {
	return err.Err // Returns inner error
}

type UnauthorizedError struct {
	Err error `json:"-"` // The original error
}

// NewUnauthorizedError create new error based on internal error.
func NewUnauthorizedError(err error) error {
	return UnauthorizedError{err}
}

// Implements the error.Error interface
func (err UnauthorizedError) Error() string {
	return err.Err.Error()
}

// Implements the errors.Unwrap interface
func (err UnauthorizedError) Unwrap() error {
	return err.Err // Returns inner error
}

type NotFoundError struct {
	Err error `json:"-"` // The original error
}

// NewNotFoundError create new error based on internal error.
func NewNotFoundError(err error) error {
	return NotFoundError{err}
}

// Implements the error.Error interface
func (err NotFoundError) Error() string {
	return err.Err.Error()
}

// Implements the errors.Unwrap interface
func (err NotFoundError) Unwrap() error {
	return err.Err // Returns inner error
}

// Internal server error.
type InternalError struct {
	Err error `json:"-"` // The original error
}

// NewInternalError create new error based on internal error.
func NewInternalError(err error) error {
	return InternalError{err}
}

// Implements the error.Error interface
func (err InternalError) Error() string {
	return err.Err.Error()
}

// Implements the errors.Unwrap interface
func (err InternalError) Unwrap() error {
	return err.Err // Returns inner error
}

// Conflict error.
type ConflictError struct {
	Err error `json:"-"` // The original error
}

// NewConflictError create new error based on internal error.
func NewConflictError(err error) error {
	return ConflictError{err}
}

// Implements the error.Error interface
func (err ConflictError) Error() string {
	return err.Err.Error()
}

// Implements the errors.Unwrap interface
func (err ConflictError) Unwrap() error {
	return err.Err // Returns inner error
}
