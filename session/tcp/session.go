package tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

// package classes
const (
	Request byte = iota
	Response
	Push
)

// Result response result to client
type Result struct {
	err  int16
	data []byte
}

//CallBack request call back
type CallBack func(*Result)

// IHandler 调度器接口, 实现调度器，完成功能
type IHandler interface {
	OnOpen(s *Session)
	OnClose(s *Session)
	OnReq(s *Session, data []byte, cb CallBack)
	OnPush(s *Session, data []byte)
}

// Session 客户端会话
type Session struct {
	id      int64
	conn    net.Conn
	handler IHandler
	packLen int
}

// Start  new session start
func (s *Session) Start() int64 {
	s.handler.OnOpen(s)

	go s.scan()

	return 0
}

// Close close session
func (s *Session) Close() {
	s.conn.Close()
}

// scan  handler scanner
func (s *Session) scan() {
	defer s.Close()

	scanner := bufio.NewScanner(s.conn)
	scanner.Split(s.splitPacket)

	// scanner
	for scanner.Scan() {
		s.dispatcher(scanner.Bytes())
	}
}

func (s *Session) splitPacket(data []byte, atEOF bool) (advance int, token []byte, err error) {
	dLen := len(data)
	width := 2

	if atEOF {
		return 0, nil, nil // session broken ?
	}

	if s.packLen == 0 {
		if dLen < width {
			return 0, nil, nil
		}

		s.packLen = int(binary.LittleEndian.Uint16(data[:width])) // get packet len
	}

	if dLen >= s.packLen+width {
		s.packLen = 0 // reset
		// received a complete packet
		return s.packLen + width, data[width : width+s.packLen], nil
	}

	//request more data
	return 0, nil, nil
	//	return bufio.ScanWords(data, atEOF) // example
}

func (s *Session) dispatcher(data []byte) {
	//one packet = class[byte] + data[byte ...]
	if len(data) <= 0 {
		return
	}

	reader := bytes.NewBuffer(data)
	ntype, err := reader.ReadByte()
	if err != nil {
		s.Close()
		log.Printf("dispatcher err %v", err.Error())
		return
	}

	switch ntype {
	case Request:
		s.onReq(data[1:])
	case Response:
		s.onRsp(data[1:])
	case Push:
		s.onPush(data[1:])
	default:
		s.Close()
		log.Printf("dispatcher err class %d", ntype)
	}
}

func (s *Session) onReq(data []byte) {
	s.handler.OnReq(s, data, func(r *Result) {
		if r == nil {
			return
		}

		s.conn.Write(packet(Response, r.err, r.data))
	})
}

func (s *Session) onRsp(data []byte) {
	// what can I do ?

}

func (s *Session) onPush(data []byte) {
	s.handler.OnPush(s, data)
}

func packet(nType byte, err int16, data []byte) []byte {
	// packlen + [class + [err + data]]
	rsp := new(bytes.Buffer)
	binary.Write(rsp, binary.LittleEndian, uint16(len(data)+2+1))
	rsp.WriteByte(nType)
	binary.Write(rsp, binary.LittleEndian, err)
	if len(data) > 0 {
		binary.Write(rsp, binary.LittleEndian, data)
	}

	return rsp.Bytes()
}
