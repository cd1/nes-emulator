package nes

import (
	"bytes"
	"fmt"
	"io"
)

const (
	GameHeaderSize = 16
	TrainerSize    = 512
	PRGBankSize    = 16384
	CHRBankSize    = 8192
	PlayChoiceSize = 8192
)

var NESMagicNumber = []uint8{0x4e, 0x45, 0x53, 0x1a}

type GameHeader []uint8

func (h GameHeader) MagicNumber() []uint8 {
	return h[0:4]
}

// x 16 kiB
func (h GameHeader) PRGBankCount() uint8 {
	return h[4]
}

// x 8 kiB
func (h GameHeader) CHRBankCount() uint8 {
	return h[5]
}

func (h GameHeader) HasTrainer() bool {
	return h[6]&0x04 != 0x00
}

func (h GameHeader) Flag6() uint8 {
	return h[6]
}

func (h GameHeader) HasPlayChoice() bool {
	return h[7]&0x02 != 0x00
}

func (h GameHeader) Flag7() uint8 {
	return h[7]
}

// x 8 kiB
func (h GameHeader) PRGRAMSize() uint8 {
	return h[8]
}

func (h GameHeader) TVSystem() uint8 {
	return h[9]
}

func (h GameHeader) TVSystem2() uint8 {
	return h[10]
}

func (h GameHeader) UnusedPadding() []uint8 {
	return h[11:16]
}

type Game struct {
	Header            GameHeader
	Trainer           []uint8
	PRG               []uint8
	CHR               []uint8
	PlayChoiceINSTROM []uint8
	PlayChoicePROM    []uint8
}

func LoadGame(data io.Reader) (*Game, error) {
	var headerBuf bytes.Buffer

	if _, err := io.CopyN(&headerBuf, data, GameHeaderSize); err != nil {
		return nil, err
	}

	gameHeader := GameHeader(headerBuf.Bytes())

	if mn := gameHeader.MagicNumber(); !bytes.Equal(mn, NESMagicNumber) {
		return nil, fmt.Errorf("invalid magic number: %v", mn)
	}

	var trainerBuf bytes.Buffer

	if gameHeader.HasTrainer() {
		if _, err := io.CopyN(&trainerBuf, data, TrainerSize); err != nil {
			return nil, err
		}
	}

	var prgBuf bytes.Buffer

	if _, err := io.CopyN(&prgBuf, data, int64(gameHeader.PRGBankCount())*PRGBankSize); err != nil {
		return nil, err
	}

	var chrBuf bytes.Buffer

	if _, err := io.CopyN(&chrBuf, data, int64(gameHeader.CHRBankCount())*CHRBankSize); err != nil {
		return nil, err
	}

	var playChoiceBuf bytes.Buffer

	if gameHeader.HasPlayChoice() {
		if _, err := io.CopyN(&playChoiceBuf, data, PlayChoiceSize); err != nil {
			return nil, err
		}
	}

	return &Game{
		Header:            gameHeader,
		Trainer:           trainerBuf.Bytes(),
		PRG:               prgBuf.Bytes(),
		CHR:               chrBuf.Bytes(),
		PlayChoiceINSTROM: playChoiceBuf.Bytes(),
	}, nil
}
