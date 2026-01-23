package database

import (
	"strings"
	"os"
	"fmt"
)

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
            return fmt.Errorf("Failed to execute statement: %s, error: %w", stmt, err)
        }
    }
    return nil
}