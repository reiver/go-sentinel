package sentinel


import (
	"github.com/reiver/go-sentinel/driver"
)


func Error(err error) error {
	return globalHandler.Error(err)
}
func (handler *internalHandler) Error(err error) error {
	return sentineldriver.Error(err)
}
