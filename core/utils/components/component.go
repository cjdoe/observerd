package components

type Component interface {
	// Asynchronously runs the component in a separate goroutine.
	// Returns the channel that reports FATAL errors of the component.
	// Exits on the first fatal error.
	Run() <-chan error

	// Synchronous method.
	// Waits until component fully stops.
	// Returns the control only when the component has been stopped.
	Stop()
}
