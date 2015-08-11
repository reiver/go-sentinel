package sentineldriver


import (
	"testing"

	"math/rand"
	"time"
)


func TestRegister(t *testing.T) {

	// Initialize.
	randomness := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	// Tests.
	numTests := 1 + randomness.Intn(50)
	for testNumber:=0; testNumber<numTests; testNumber++ {

		registry := newInternalRegistry()

		if actual := len(registry.drivers); 0 != actual {
			t.Errorf("For test #%d, expected nothing to be registered, but actually was %d registered.", testNumber, actual)
		}


		driver := newTestDriver()

		numDrivers := 1 + randomness.Intn(12)

		for i:=0; i<numDrivers; i++ {
			registry.Register(driver)
		}

		if actual, expected := len(registry.drivers), numDrivers; expected != actual {
			t.Errorf("For test #%d, expected there to be %d registered, but actually was %d registered.", testNumber, expected, actual)
		}
	}
}
