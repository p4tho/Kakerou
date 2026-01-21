package database

/*
Command Operations
*/
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
    rows, err := DB.Query(`SELECT uid, command_id, command, status FROM commands`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var commands []Command

    for rows.Next() {
        var command Command
        err := rows.Scan(&command.Uid, &command.Command_id, &command.Command, &command.Status)
        if err != nil {
            return nil, err
        }
        commands = append(commands, command)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
		printDBInfo("Failed to retrieve command commands:", err.Error())
        return nil, err
    }

	printDBInfo("Retrieved commands!")

    return commands, nil
}

/*
Agent Operations
*/
func InsertAgent(agent *Agent) error {
	sql_stmt := `
		INSERT INTO agents (name)
		VALUES (?)
	`
	_, err := DB.Exec(sql_stmt, agent.Name)
	if err != nil {
		printDBInfo("Failed to insert command:", err.Error())
		return err
	}

	printDBInfo("Agent inserted:", agent.Name)

	return nil
}

func GetAllAgents() ([]Agent, error) {
    rows, err := DB.Query(`SELECT uid, name FROM agents`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var agents []Agent

    for rows.Next() {
        var agent Agent
        err := rows.Scan(&agent.Uid, &agent.Name)
        if err != nil {
            return nil, err
        }
        agents = append(agents, agent)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
		printDBInfo("Failed to retrieve command commands:", err.Error())
        return nil, err
    }

	printDBInfo("Retrieved agents!")

    return agents, nil
}