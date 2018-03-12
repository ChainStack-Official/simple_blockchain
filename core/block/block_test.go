package block

import (
	"testing"
	"time"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type S struct{}
var _ = Suite(&S{})


func (s *S) TestNewBlock(c *C) {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2018-03-01 00:00:00")
	b := Block{
		Index: 1,
		Timestamp: testTime.Unix(),
		Msg: "哈哈",
	}
	c.Assert(b.HashForThisBlock(), Equals, "156c32040f92b2e3239a9943fdde31f61e5c5abec785cc6fd69ab6e10275f109")
}

func (s *S) TestGetContentForHash(c *C)  {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2018-03-01 00:00:00")
	b := Block{
		Index: 1,
		Timestamp: testTime.Unix(),
		Msg: "hello",
	}
	c.Assert(b.HashForThisBlock(), Equals, "8ecbb210c1aa7ca83ff59fb53e8a8bed3f210043e53e7a9ae88ba72881cb4a0b")
}
