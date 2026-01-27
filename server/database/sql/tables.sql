CREATE TABLE IF NOT EXISTS commands (
    uid         INTEGER PRIMARY KEY AUTOINCREMENT,
    command_id 	INTEGER NOT NULL,
    command 	TEXT
);

CREATE TABLE IF NOT EXISTS agents (
    uid         INTEGER PRIMARY KEY AUTOINCREMENT,
    name 		TEXT NOT NULL UNIQUE,
    executed_commands_list TEXT NOT NULL
);