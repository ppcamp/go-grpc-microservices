package base

import (
	"context"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/services/base"
	"streamer/utils/jwt"
	"sync"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Cache    cache.Cache
	Signer   jwt.Jwt
	Database database.Connection
}

var (
	handler *Handler
	once    sync.Once
	lock    sync.RWMutex
)

func InitHandler(h *Handler) {
	once.Do(func() {
		handler = h
	})
}

func GetHandler() *Handler {
	lock.RLock()
	defer lock.RUnlock()

	if handler == nil {
		logrus.Warn("you need to initialize the handler before")
	}

	return handler
}

// Handle is responsible to identify the service type and assign the correct
// workflow to it. Furthermore, it's also responsible to commit/rollback
// transactions when needed and set another responses parameters if needed.
func Handle[In, Out any](
	ctx context.Context,
	input In,
	service base.BaseBusiness[In, Out],
) (response *Out, err error) {
	service.SetContext(ctx)

	select {
	case <-ctx.Done():
		return
	default:
		switch service := service.(type) {
		case base.BaseTransactionBusiness[In, Out]:
			return handleTransactionService(input, service)
		default:
			return handleBaseService(input, service)
		}
	}
}

func handleBaseService[In, Out any](
	input In,
	service base.BaseBusiness[In, Out],
) (response *Out, err error) {
	return service.Execute(input)
}

func handleTransactionService[In, Out any](
	input In,
	service base.BaseTransactionBusiness[In, Out],
) (response *Out, err error) {
	h := GetHandler()

	tr, err := h.Database.StartTransaction()
	defer func() {
		panicked := recover()

		if err != nil || panicked != nil {
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
		return nil, err
	}

	service.SetTransaction(tr)
	return service.Execute(input)
}
