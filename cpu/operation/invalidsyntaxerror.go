package operation

import "fmt"

type InvalidSyntaxError struct {
	Line string
}

func (err InvalidSyntaxError) Error() string {
	return fmt.Sprintf("invalid syntax in line: %v", err.Line)
}
