package database

/* Database Schemas */
type Command struct {
	Command_uid int		`json:"command_uid"`
	Command_id 	int		`json:"command_id"`
	Command 	string	`json:"command"`
	Status 		int		`json:"status"`
}