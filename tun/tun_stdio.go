package tun

import (
	"bufio"
	"io"
	"os"
)

func OpenTunDeviceStdin(name, addr, gw, mask string, dnsServers []string, persist bool) (io.ReadWriteCloser, error) {
	return NewStdioReadWriteCloser(), nil
}

type StdioReadWriteCloser struct {
	io.ReadWriteCloser

	Stdin  *bufio.Reader
	Stdout *bufio.Writer
}

func NewStdioReadWriteCloser() *StdioReadWriteCloser {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	return &StdioReadWriteCloser{Stdin: reader, Stdout: writer}
}

func (rwc *StdioReadWriteCloser) Read(b []byte) (n int, err error) {
	return rwc.Stdin.Read(b)
}

func (rwc *StdioReadWriteCloser) Write(b []byte) (n int, err error) {
	return rwc.Stdout.Write(b)
}

func (rwc *StdioReadWriteCloser) Close() error {
	return os.Stdout.Close()
}
