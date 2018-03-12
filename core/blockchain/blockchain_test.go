package blockchain

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type S struct{}
var _ = Suite(&S{})

func (s *S)TestNewBlockchain(c *C) {
	bc := NewBlockchain()
	c.Assert(len(bc.blocks), Equals, 1)
}

