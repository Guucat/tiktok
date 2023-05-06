package comment_proto

import "google.golang.org/protobuf/proto"

func (c *CommentActionRequest) Encode() (b []byte, e error) {
	b, e = proto.Marshal(c)
	l := int64(len(b))
	c.Len = &l
	return
}

func (c *CommentActionRequest) Length() int {
	return int(*c.Len)
}
