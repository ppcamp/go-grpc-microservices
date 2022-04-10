package base

import "streamer/utils/jwt"

type BaseBusiness[Input any, Output any] interface {
	// Execute is the function that will be executed by the controller
	Execute(in Input) (response *Output, err error)

	// SetSession is used after got JWT token
	SetSession(session jwt.Session)
}

type BaseBusinessImpl struct {
	Session jwt.Session
}

func (bc *BaseBusinessImpl) SetSession(session jwt.Session) {
	bc.Session = session
}
