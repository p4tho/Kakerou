# server
This is a daemon that handles communication between the attacker an implants, and tracks commands and their results in a database.

## Go Files
`main.go`: Entry point for server that initializes C2, database, parses flags, and loads config.json

## Packages
`/c2`: Server communication wiring that manages the listeners depending on what method is chosen

`/database`: Responsible for defining database used to track commands and their results

`/handlers`: Defines event handlers for attacker and implant requests

