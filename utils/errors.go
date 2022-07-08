package utils

import "errors"

var (
	// will be trow if client tries to delete a non-existing row
	ErrRowCnt = errors.New("[Error] no rows was deleted, expected single row affected")
)
