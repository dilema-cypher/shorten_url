package errors

import "errors"

var (
	ErrInvalidURL     = errors.New("invalid url: url is required")
	ErrInvalidURLType = errors.New("invalid url: url must be a valid string")
	ErrDatabaseError  = errors.New("database error")
	ErrRedisError     = errors.New("redis error")
)

type InvalidURLError struct {
	Message string
}

func (e InvalidURLError) Error() string {
	return e.Message
}

type DatabaseError struct {
	Message string
	Err     error
}

func (e DatabaseError) Error() string {
	return e.Message
}

func (e DatabaseError) Unwrap() error {
	return e.Err
}

type RedisError struct {
	Message string
	Err     error
}

func (e RedisError) Error() string {
	return e.Message
}

func (e RedisError) Unwrap() error {
	return e.Err
}