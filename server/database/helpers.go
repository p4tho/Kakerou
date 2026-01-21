package database

import (
	"strings"
	"os"
	"fmt"
)

func printDBInfo(messages ...any) {
	fmt.Printf("[*DATABASE] - ")
	for _, msg := range messages {
		fmt.Print(msg, " ")
	}
	fmt.Println()
}

func execSQLFileSplit(filepath string) error {
    content, err := os.ReadFile(filepath)
    if err != nil {
        return err
    }

    statements := strings.Split(string(content), ";")

    for _, stmt := range statements {
        stmt = strings.TrimSpace(stmt)
        if stmt == "" {
            continue
        }
        _, err := DB.Exec(stmt)
        if err != nil {
            return fmt.Errorf("failed to execute statement: %s, error: %w", stmt, err)
        }
    }
    return nil
}