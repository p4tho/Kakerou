package database

/* Database Schemas */
type Command struct {
	Uid 		int		`json:"uid"`
	Command_id 	int		`json:"command_id"`
	Command 	string	`json:"command"`
	Status 		int		`json:"status"`
}

type Agent struct {
	Uid 		int		`json:"uid"`
	Name		string	`json:"name"`
}