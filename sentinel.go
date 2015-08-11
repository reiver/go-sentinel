package sentinel


var (
	globalHandler = newSentinel()
)


type Sentinel interface {
	Error(err error) error
}


type internalHandler struct{}


func newSentinel() Sentinel {
	handler := internalHandler{}

	return &handler
}
