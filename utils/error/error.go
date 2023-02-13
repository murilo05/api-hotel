package error

import "gitlab.engdb.com.br/apigin/domain/entities"

//createError - returns error struct
func CreateError(errorCode int, errorMessage string) entities.Error {
	e := entities.Error{}
	e.Code = errorCode
	e.ErrorMessage = errorMessage
	return e
}
