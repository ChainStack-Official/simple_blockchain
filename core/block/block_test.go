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
		Index:     1,
		Timestamp: testTime.Unix(),
		Msg:       "哈哈",
	}
	c.Assert(b.HashForThisBlock(), Equals, "7a655569d7aab2029ca20b72aea1237bb7852fc2a5b4a297e9a49b4fed453b2a")
}

func (s *S) TestGetContentForHash(c *C) {
	testTime, _ := time.Parse("2006-01-02 15:04:05", "2018-03-01 00:00:00")
	b := Block{
		Index:     1,
		Timestamp: testTime.Unix(),
		Msg:       "hello",
	}
	c.Assert(b.HashForThisBlock(), Equals, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
}

func (s *S) TestCutSlice(c *C) {
	arr := []int{1, 2}
	arr = arr[1:]
	c.Assert(arr[0], Equals, 2)
}
