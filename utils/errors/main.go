package errors

const (
	MissingKey          = "missing key"
	PrefixNotFound      = "prefix %s not found in %s method"
	MethodNotAllowed    = "method not allowed"
	CacheRecordNotFound = "record not found"
	UnknownPanic        = "unknown panic"
	TaskMsg             = "task will be run after %v seconds"
	RoutineError        = "run routine recovered - err: %v"
	ServerListen        = "could listen on %s"
	ServerNotListen     = "could not listen on %s: %v"
	ServerShutDown      = "server is shutting down"
	ServerNotShutdown   = "could not shutdown on %s: %v"
	NotFound            = "not found"
)
