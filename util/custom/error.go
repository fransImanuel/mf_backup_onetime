package custom

import "fmt"

func ErrorOperationDB(table string, operation string) error {
	return fmt.Errorf("error on %v %v", operation, table)
}

func ErrorExistDB(table string, column string, value interface{}) error {
	return fmt.Errorf("%v %v already exist in %v", column, value, table)
}

func ErrorNotFoundDB(table string, column string, value interface{}) error {
	return fmt.Errorf("%v %v not found in %v", column, value, table)
}
