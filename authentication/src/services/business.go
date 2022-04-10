package services

import (
	"context"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/services/base"
	"streamer/utils/jwt"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Cache    cache.Cache
	Signer   jwt.Jwt
	Database database.Connection
}

func (h *Handler) Handle(ctx context.Context, input any, service base.BaseBusiness) (response *any, err error) {
	service.SetContext(ctx)

	select {
	case <-ctx.Done():
		return
	default:
		switch service := service.(type) {
		case base.BaseTransactionBusiness:
			return h.handleTransactionService(input, service)
		default:
			return h.handleBaseService(input, service)
		}
	}
}

func (h *Handler) handleBaseService(input any, service base.BaseBusiness) (response *any, err error) {
	return service.Execute(input)

}

func (h *Handler) handleTransactionService(input any, service base.BaseTransactionBusiness) (response *any, err error) {
	tr, err := h.Database.StartTransaction()
	defer func() {
		if err != nil {
			if err := tr.Rollback(); err != nil {
				logrus.WithFields(logrus.Fields{"input": input}).
					Error("fail to rollback transaction")
			}
		} else {
			if err := tr.Commit(); err != nil {
				logrus.WithFields(logrus.Fields{"input": input}).
					Error("fail to commit transaction")
			}
		}
	}()
	if err != nil {
		panic(err)
	}

	service.SetTransaction(tr)
	return service.Execute(input)
}
