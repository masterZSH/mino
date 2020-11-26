package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime/debug"
	"strings"

	"github.com/masterZSH/mino"

	"github.com/google/uuid"
)

// Key type
type Key struct {
	Pass []byte `json:"pass"`
	Salt []byte `json:"salt"`
}

var localPort int
var targetURL string

const maxSize = 2 << 10

var k Key

var targetConn net.Conn

func init() {
	flag.IntVar(&localPort, "p", 6666, "local port")
	flag.StringVar(&targetURL, "t", "", "target url")
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", getListenAddr())
	handleError(err)
	end := make(chan bool)
	if targetURL != "" {
		targetConn, err := net.Dial("tcp", targetURL)
		handleError(err)
		go handleConn(targetConn)
		send(targetConn)
	} else {
		targetConn, err := l.Accept()
		handleError(err)
		k = createKey()
		h := make([]byte, 8)
		copy(h, "key")
		var buf bytes.Buffer
		buf.Write(k.Pass)
		buf.WriteString("|")
		buf.Write(k.Salt)
		content := append(h, buf.Bytes()...)
		targetConn.Write(content)
		go handleConn(targetConn)
		send(targetConn)
	}
	<-end
}

func getListenAddr() string {
	return fmt.Sprintf(":%d", localPort)
}

func handleError(err error) {
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
}

func handleConn(conn net.Conn) {
	fmt.Printf("get conn from :%s\n", conn.RemoteAddr())
	for {
		buff := make([]byte, 0xffff)
		n, err := conn.Read(buff)
		if err != nil {
			return
		}
		if n > maxSize {
			err = fmt.Errorf("too many bytes: %d", maxSize)
			handleError(err)
		}
		header := buff[:8]
		headerStr := strings.TrimRight(string(header), string(byte(0)))
		if headerStr == "key" {
			arr := strings.Split(string(buff[8:n]), "|")
			k = Key{
				Pass: []byte(arr[0]),
				Salt: []byte(arr[1]),
			}
			fmt.Print(k)
			continue
		}
		fmt.Printf("receive: %s\n", buff[8:n])
		minoKey, err := mino.NewKey(k.Pass, k.Salt)
		handleError(err)
		plaintext, err := minoKey.Decrypt(buff[8:n])
		handleError(err)
		fmt.Printf("decrypt messageg: %s\n", plaintext)
	}
}

func send(conn net.Conn) {
	var msg []byte
	fmt.Scanln(&msg)
	if len(k.Pass) == 0 {
		return
	}
	minoKey, err := mino.NewKey(k.Pass, k.Salt)
	handleError(err)
	fmt.Printf("send msg:%s\n", msg)
	ciphertext, err := minoKey.Encrypt(msg)
	handleError(err)
	h := make([]byte, 8)
	copy(h, "msg")
	content := append(h, ciphertext...)
	conn.Write(content)
}

func createKey() Key {
	return Key{
		Pass: []byte(uuid.New().String()),
		Salt: []byte("test"),
	}
}
