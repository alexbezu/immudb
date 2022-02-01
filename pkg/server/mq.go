package server

import (
	"context"

	"github.com/codenotary/immudb/pkg/api/schema"
)

func (s *ImmuServer) MQput(ctx context.Context, req *schema.MQputRequest) (*schema.MQputReply, error) {
	s.mux.Lock()
	if s.mq == nil {
		s.mq = make(map[string]chan []byte)
	}
	_, ok := s.mq[req.Qname]
	if !ok {
		s.mq[req.Qname] = make(chan []byte, 64)
	}
	s.mux.Unlock()
	s.mq[req.Qname] <- req.Value
	return &schema.MQputReply{Value: []byte("Ok")}, nil
}

func (s *ImmuServer) MQpop(ctx context.Context, req *schema.MQpopRequest) (*schema.MQpopReply, error) {
	s.mux.Lock()
	if s.mq == nil {
		s.mq = make(map[string]chan []byte)
	}
	_, ok := s.mq[req.Qname]
	if !ok {
		s.mq[req.Qname] = make(chan []byte, 64)
	}
	s.mux.Unlock()
	reply := schema.MQpopReply{Value: <-s.mq[req.Qname]}
	return &reply, nil
}
