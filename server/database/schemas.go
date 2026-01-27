package database

/* Database Schemas */
type Command struct {
	Uid 			int		`json:"uid"`
	Command_id 		int		`json:"command_id"`
	Command 		string	`json:"command"`
}

type Agent struct {
	Uid 			int		`json:"uid"`
	Name			string	`json:"name"`
	Executed_commands_list 	string
}