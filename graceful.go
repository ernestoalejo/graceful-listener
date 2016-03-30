package graceful

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

// ShutdownCallback is implemented by callbacks of graceful listeners and will be called
// before closing the listener.
type ShutdownCallback func()

// NewGracefulListener creates a new listener that will close itself gracefully after
// Ctrl+C is pressed without leaving clients behind. It will run all callbacks before
// closing the listener.
func NewGracefulListener(address string, callbacks ...ShutdownCallback) (net.Listener, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt

		signal.Stop(interrupt)

		for _, callback := range callbacks {
			callback()
		}

		listener.Close()
	}()

	return listener, nil
}
