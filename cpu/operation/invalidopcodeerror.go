package operation

import "fmt"

type InvalidOpCodeError struct {
	OpCode uint8
}

func (err InvalidOpCodeError) Error() string {
	return fmt.Sprintf("invalid op code: $%02X", err.OpCode)
}
