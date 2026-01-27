package database

import (
    "log"
    "strings"
    "strconv"
)

/*
Command Operations
*/
func InsertCommand(c *Command) error {
	sql_stmt := `
		INSERT INTO commands (command_id, command)
		VALUES (?, ?)
	`
	_, err := DB.Exec(sql_stmt, c.Command_id, c.Command)
	if err != nil {
		log.Printf("Failed to insert command: %v\n", err.Error())
		return err
	}

	log.Printf("Command inserted: %v\n", c.Command)

	return nil
}

func GetAllCommands() ([]Command, error) {
    rows, err := DB.Query(`SELECT uid, command_id, command FROM commands`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Initialize emoty array to send
    commands := []Command{}

    // Add each row of commands table
    for rows.Next() {
        var command Command
        err := rows.Scan(&command.Uid, &command.Command_id, &command.Command)
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

    return commands, nil
}

/*
Agent Operations
*/
func InsertAgent(agent *Agent) (int64, error) {
	sql_stmt := `
		INSERT INTO agents (name, executed_commands_list)
		VALUES (?, ?)
	`
	result, err := DB.Exec(sql_stmt, agent.Name, "")
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

func GetAgentCommandsList(agent_uid int) ([]int, error) {
    var executed_commands_list_str string
    var executed_commands_list_arr []string
    var err error

    // Get agent's command list to update
    err = DB.QueryRow(`
        SELECT executed_commands_list 
        FROM agents WHERE 
        uid = ?`, 
        agent_uid,
    ).Scan(&executed_commands_list_str)
    if err != nil {
        return nil, err
    }
    if executed_commands_list_str == "" { // Return empty list it the database string was empty
        return []int{}, nil
    }
    executed_commands_list_arr = strings.Split(executed_commands_list_str, ",")

    // Convert list from string to int
    executed_commands_list_arr_int := make([]int, len(executed_commands_list_arr))
    for idx, s := range executed_commands_list_arr {
        n, convert_err := strconv.Atoi(s)
        if convert_err != nil {
            return nil, convert_err
        }

        executed_commands_list_arr_int[idx] = n
    }

    return executed_commands_list_arr_int, err
}

func AddToAgentExecutedCommandsList(agent_uid int, command_uid int) error {
    var executed_commands_list_str string
    var executed_commands_list_arr []string
    var err error

    // Get agent's command list to update
    err = DB.QueryRow(`
        SELECT executed_commands_list 
        FROM agents WHERE 
        uid = ?`, 
        agent_uid,
    ).Scan(&executed_commands_list_str)
    if err != nil {
        return err
    }

    // Add uid to command list
    if executed_commands_list_str == "" {
        executed_commands_list_arr = []string{}
    } else {
        executed_commands_list_arr = strings.Split(executed_commands_list_str, ",")
    }
    executed_commands_list_arr = append(
        executed_commands_list_arr,
        strconv.Itoa(command_uid),
    )
    executed_commands_list_str = strings.Join(executed_commands_list_arr, ",")

    // Update database with new list
    _, err = DB.Exec(`
        UPDATE agents
        SET executed_commands_list = ?
        WHERE uid = ?`,
        executed_commands_list_str,
        agent_uid,
    )
    if err != nil {
        return err
    }

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
		log.Printf("Failed to retrieve agents: %v\n", err.Error())
        return nil, err
    }

	log.Println("Retrieved agents!")

    return agents, nil
}