package utils

import "errors"

var (
	// will be trown if client tries to delete a non-existing row
	ErrRowCnt = errors.New("[Error] no rows was deleted, expected single row affected")
	// will be thrown if update payload would be empty
	ErrEmptyPayload = errors.New("[Error] update payload is empty")
)
