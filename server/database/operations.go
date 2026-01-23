package database

import (
    "log"
)

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
		log.Printf("Failed to insert command: %v\n", err.Error())
		return err
	}

	log.Printf("Command inserted: %v\n", c.Command)

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
		log.Printf("Failed to retrieve commands: %v\n", err.Error())
        return nil, err
    }

	log.Println("Retrieved commands!")

    return commands, nil
}

/*
Agent Operations
*/
func InsertAgent(agent *Agent) (int64, error) {
	sql_stmt := `
		INSERT INTO agents (name)
		VALUES (?)
	`
	result, err := DB.Exec(sql_stmt, agent.Name)
	if err != nil {
		log.Printf("Failed to insert agent: %v\n", err.Error())
		return -1, err
	}

	uid, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to get last insert id: %v\n", err.Error())
		return 0, err
	}

	log.Printf("Agent inserted: %s (UID=%d)\n", agent.Name, uid)

	return uid, nil
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
		log.Printf("Failed to retrieve agents: %v\n", err.Error())
        return nil, err
    }

	log.Println("Retrieved agents!")

    return agents, nil
}