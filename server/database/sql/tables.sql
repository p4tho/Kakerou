CREATE TABLE IF NOT EXISTS commands (
    command_uid INTEGER PRIMARY KEY AUTOINCREMENT,
    command_id 	INTEGER NOT NULL,
    command 	TEXT,
    status 		INTEGER NOT NULL
);