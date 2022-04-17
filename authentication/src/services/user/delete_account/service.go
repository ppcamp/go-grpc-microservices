package delete_account

import (
	"authentication/helpers/services"
	"authentication/repositories/database"
)

type RecoverPasswordService struct {
	services.TransactionBusiness[database.UserStorage]
}

func NewService() services.ITransactionBusiness[Input, Output] {
	return new(RecoverPasswordService)
}

func (s *RecoverPasswordService) Execute(in Input) (*Output, error) {
	return new(Output), nil
}
