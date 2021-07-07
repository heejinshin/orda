package errors_test

import (
	"github.com/orda-io/orda/pkg/errors"
	"github.com/orda-io/orda/pkg/log"
	"testing"
)

func TestOrdaError(t *testing.T) {
	_ = call1()
}

func call1() errors.OrdaError {
	return call2()
}

func call2() errors.OrdaError {
	return errors.DatatypeNoOp.New(log.Logger, "This is a generated sample error")
}
