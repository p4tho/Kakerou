package database

/* Command Operations */
func InsertCommand(c *Command) error {
	sql_stmt := `
		INSERT INTO commands (command_id, command, status)
		VALUES (?, ?, ?)
	`
	_, err := DB.Exec(sql_stmt, c.Command_id, c.Command, c.Status)
	if err != nil {
		printDBInfo("Failed to insert command:", err.Error())
		return err
	}

	printDBInfo("Command inserted:", c.Command)

	return nil
}

func GetAllCommands() ([]Command, error) {
    rows, err := DB.Query(`SELECT command_uid, command_id, command, status FROM commands`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var commands []Command

    for rows.Next() {
        var c Command
        err := rows.Scan(&c.Command_uid, &c.Command_id, &c.Command, &c.Status)
        if err != nil {
            return nil, err
        }
        commands = append(commands, c)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
		printDBInfo("Failed to retrieve command commands:", err.Error())
        return nil, err
    }

	printDBInfo("Retrieved commands!")

    return commands, nil
}