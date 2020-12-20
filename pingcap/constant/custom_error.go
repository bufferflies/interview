package constant

import "errors"

var (
	FileNotFindError   = errors.New("file not find")
	FileEofErrpr       = errors.New("file read eof")
	KeyIsNotExistError = errors.New("key is not exist")
)
