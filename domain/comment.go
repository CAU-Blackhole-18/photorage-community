package domain

import "fmt"

/*
Comment Entity를 표현하기 위한 구조체
*/
type Comment struct {
	AlbumSeq int64
	Content  string
}

// Comment에 Bind(Receiver)를 붙인 Method
func (c *Comment) PrintCommentInfo() {
	fmt.Println("My Comment Info = ", c.AlbumSeq, c.Content)
}

func (c Comment) Run() int32 {
	fmt.Println("Interface = ", c.AlbumSeq, c.Content)
	return int32(c.AlbumSeq)
}

type Animal interface {
	Run() int32
}
