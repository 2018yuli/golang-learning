package pipeline

import (
	"bufio"
	"net"
)

func NetworkSink(addr string, in <-chan int) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	// defer listen.Close()

	go func() {
		defer listen.Close()
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer, in)
	}()

}

func NetworkSource(addr string) <-chan int {
	out := make(chan int)

	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(conn)
		r := ReaderSource(reader, -1)
		for v := range r {
			out <- v
		}

		// 不要忘记关闭资源
		close(out)
	}()

	return out
}
