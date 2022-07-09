package interface_demo

import (
	"bytes"
	"io"
	"os"
	"time"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

//接口内嵌
type ReadWriter interface {
	Read(p []byte) (n int , err error)
	Write(p []byte) (n int, err error)
}

// 一个类型如果拥有一个接口需要的所有方法 name这个类型就实现了这个接口
var w io.Writer
var rwc io.ReadWriter

func main() {
	// 表达一个类型属于某个接口 只要这个类型实现了某个接口
	w = os.Stdout
	w = new(bytes.Buffer)
	w = rwc
	//rwc = w lacks Read Interface achieve
	//w = time.Second
}