package sentineldriver


func Error(err error) error {
	return globalRegistry.Error(err)
}
func (registry *internalRegistry) Error(err error) error {
	for _,driver := range registry.drivers {
		driver.Error(err)
	}

	return err
}
