package handlers

import (
	"context"
	"streamer/helpers/services"
	"streamer/microservices/auth"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"sync"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Cache    cache.Cache
	Database database.Connection
	Auth     auth.IUserPassword
}

var (
	handler *Handler
	once    sync.Once
	lock    sync.RWMutex
)

func Init(h *Handler) {
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
	service services.IBaseBusiness[In, Out],
) (*Out, error) {
	service.SetContext(ctx)

	switch service := service.(type) {
	case services.ITransactionBusiness[In, Out]:
		return handleTransactionService(input, service)
	default:
		return handleBaseService(input, service)
	}
}

func handleBaseService[In, Out any](
	input In,
	service services.IBaseBusiness[In, Out],
) (response *Out, err error) {

	return service.Execute(input)
}

func handleTransactionService[In, Out any](
	input In,
	service services.ITransactionBusiness[In, Out],
) (response *Out, err error) {
	h := GetHandler()

	tr, err := h.Database.StartTransaction()
	if err != nil {
		return nil, err
	}

	defer func() {
		panicked := recover()

		if err != nil || panicked != nil {
			logrus.WithFields(logrus.Fields{
				"panic": panicked,
				"err":   err,
			}).Warn("err or panicked")
			if err := tr.Rollback(); err != nil {
				logrus.WithFields(logrus.Fields{"input": input}).
					Error("fail to rollback transaction")
			}

			return
		}

		if err := tr.Commit(); err != nil {
			logrus.WithFields(logrus.Fields{"input": input}).
				Error("fail to commit transaction")
		}
	}()

	service.SetTransaction(tr)

	response, err = service.Execute(input)
	return response, err
}
