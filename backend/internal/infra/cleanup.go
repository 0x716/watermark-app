package infra

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// RegisterCleanup 註冊釋放資源的監聽
func RegisterCleanup() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-c
		log.Printf("Recevied signal %v, cleaning up...", sig)

		if DB != nil {
			if err := DB.Close(); err != nil {
				log.Fatalf("DB close failed: %v", err)
			}
		}

		log.Println("Cleanup completed.")
		os.Exit(0)
	}()
}
