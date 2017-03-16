package utils

import (
	"net/http"
	"strings"

	. "gopkg.in/check.v1"
)

type ReqSerializerSuite struct{}

var _ = Suite(&ReqSerializerSuite{})

func (s *ReqSerializerSuite) TestSerialize(c *C) {
	request, err := http.NewRequest("GET", "http://google.com", strings.NewReader(""))
	c.Assert(err, IsNil)
	requeststr := SerializeHttpReq(request)
	c.Assert(requeststr, Equals, "GET / HTTP/1.1\r\nHost: google.com\r\n\r\n")
}
