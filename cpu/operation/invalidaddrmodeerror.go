package operation

import "fmt"

type InvalidAddressModeError struct {
	AddressMode uint8
}

func (err InvalidAddressModeError) Error() string {
	return fmt.Sprintf("invalid address mode: %v", AddressModeString(err.AddressMode))
}
