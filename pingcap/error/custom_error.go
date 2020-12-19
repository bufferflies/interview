package error

import "errors"

var (
	FileNotFindError   = errors.New("file not find")
	KeyIsNotExistError = errors.New("key is not exist")
)
