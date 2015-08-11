package sentineldriver


type testDriver struct{
	NumErrors int
}
func (driver *testDriver) Error(err error) {
	driver.NumErrors++
}
func newTestDriver() *testDriver {
	testDriver := testDriver{}

	return &testDriver
}
