package apperrors

import "errors"

var ErrDeleteNotFound error = errors.New("failed delete quote: not found")
var ErrDeleteBadID error = errors.New("failed delete: ID is not a number")
