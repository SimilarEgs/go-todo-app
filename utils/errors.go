package utils

import "errors"

var (
	// will be trown if client tries to delete a non-existing row
	ErrRowCntDel = errors.New("[Error] no rows was deleted, expected single row affected")
	// will be trown if client tries to update a non-existing row
	ErrRowCntUp = errors.New("[Error] no rows was updated, expected single row affected")
	// will be thrown if update payload would be empty
	ErrEmptyPayload = errors.New("[Error] update payload is empty")
)

var (
	// will be thrown to inform client about empty received rows
	ErrRowCntGet = errors.New("[Info] no rows was recived")
)
