package pause

import (
	"os"
	"os/signal"
	"syscall"
)

func Pause() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
