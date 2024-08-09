package custom

import "fmt"

func ResponseMessageSuccess(operation string) string {
	return fmt.Sprintf("%v successfully", operation)
}

func ResponseMessageFailed(operation string, err error) string {
	return fmt.Sprintf("fail on %v with error: %v", operation, err.Error())
}
