package sentineldriver


import (
	"testing"

	"errors"
	"math/rand"
	"time"
)


func TestError(t *testing.T) {

	// Initialize.
	randomness := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	// Tests.
	numTests := 1 + randomness.Intn(50)
	for testNumber:=0; testNumber<numTests; testNumber++ {

		registry := newInternalRegistry()


		driver := newTestDriver()

		numDrivers := 1 + randomness.Intn(12)

		for i:=0; i<numDrivers; i++ {
			registry.Register(driver)
		}


		if actual := driver.NumErrors; 0 != actual {
			t.Errorf("For test #%d, expected no errors, but actually was %d", testNumber, actual)
		}

		err1 := errors.New("The error")
		err2 := registry.Error(err1)

		if err1 != err2 {
			t.Errorf("For test #%d, expected the errors to be exactly the same but weren't.", testNumber)
		}

		if expected, actual := numDrivers, driver.NumErrors; expected != actual {
			t.Errorf("For test #%d, expected %d errors, but actually was %d", testNumber, expected, actual)
		}
	}
}
