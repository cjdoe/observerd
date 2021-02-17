package e

import "errors"

var (
	ErrNoBlocks     = errors.New("no blocks")
	ErrInvalidBlock = errors.New("invalid block")
)
