package nes

const (
	StatusCarry uint8 = 1 << iota
	StatusZero
	StatusInterrupt
	StatusDecimal
	StatusBreak
	StatusUnused
	StatusOverflow
	StatusNegative
)

type CPU struct {
	Accumulator    uint8
	IndexX         uint8
	IndexY         uint8
	StackPointer   uint8
	ProgramCounter uint16
	Status         uint8
}

func (c CPU) GetStatus(statusFlag uint8) bool {
	return (c.Status&statusFlag != 0x00)
}

func (c *CPU) SetStatus(statusFlag uint8, isSet bool) {
	if isSet {
		c.Status |= statusFlag
	} else {
		c.Status &= ^statusFlag
	}
}
