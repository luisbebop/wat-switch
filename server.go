package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-mruby"
	"io/ioutil"
	"net"
	"strings"
)

func ExecRubyFile(name string, input string) (string, error) {
	mrb := mruby.NewMrb()
	defer mrb.Close()

	filename := name + ".rb"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	mrb.LoadString(fmt.Sprintf("$INPUT = \"%s\"", input))
	output, err := mrb.LoadString(string(b))
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func echo(s net.Conn, i int) {
	defer s.Close()

	fmt.Printf("%d: %v <-> %v\n", i, s.LocalAddr(), s.RemoteAddr())
	b := bufio.NewReader(s)
	for {
		line, e := b.ReadBytes('\n')
		if e != nil {
			break
		}

		exec := strings.Fields(string(line))
		if len(exec) > 1 {
			fmt.Printf("run ruby file=%s, input=%s\n", exec[0], exec[1])
			ret, err := ExecRubyFile(exec[0], exec[1])
			if err != nil {
				s.Write([]byte(fmt.Sprintf("err mrb=%s\n", err.Error())))
			} else {
				s.Write([]byte(ret + "\n"))
			}
		} else {
			s.Write(line)
		}
	}
	fmt.Printf("%d: closed\n", i)
}

func main() {
	fmt.Printf("Listening port :31415 ...\n")
	l, e := net.Listen("tcp", ":31415")
	for i := 0; e == nil; i++ {
		var s net.Conn
		s, e = l.Accept()
		go echo(s, i)
	}
}
