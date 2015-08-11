package sentineldriver



type internalRegistry struct {
	drivers []Driver
}


func newInternalRegistry() *internalRegistry {

	drivers := make([]Driver, 0, 2)

	registry := internalRegistry{
		drivers:drivers,
	}

	return &registry
}
