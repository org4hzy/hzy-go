package tcp

import (
	"log"
	"net"
)

// Server 监听服务对象
type Server struct {
	sessions map[int64]*Session
	hander   IHandler
	sid      int64
}

// NewServer new TCP server
func NewServer(hand *IHandler) *Server {
	return &Server{
		sessions: make(map[int64]*Session),
		hander:   *hand,
		sid:      0,
	}
}

// Start server start
func (s *Server) Start(addr string) {
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Panicf("Fatal err: %v", err)
	} else {
		log.Printf("server running on %s", addr)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("socket accept err %v", err)
			continue
		}

		s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	sess := &Session{
		id:      s.newID(),
		conn:    conn,
		handler: s.hander,
		packLen: 0,
	}

	s.sessions[sess.id] = sess
	sess.Start()
	log.Printf("open a new session %d", sess.id)
}

func (s *Server) newID() int64 {
	s.sid++ // need check max sids
	return s.sid
}
