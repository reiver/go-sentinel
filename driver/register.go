package sentineldriver


func Register(driver Driver) error {
	return globalRegistry.Register(driver)
}
func (registry *internalRegistry) Register(driver Driver) error {

	registry.drivers = append(registry.drivers, driver)

	return nil
}
