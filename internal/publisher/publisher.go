package publisher

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/maddiesch/telemetry/internal/packet"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Publisher struct {
	mu sync.Mutex
	s  http.Server
}

func New() (*Publisher, error) {
	return &Publisher{
		s: http.Server{
			Addr: net.JoinHostPort("0.0.0.0", "3021"),
		},
	}, nil
}

func (p *Publisher) Start(ctx context.Context) (chan<- packet.Packet, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	receiver := make(chan packet.Packet, 10)

	mux := http.NewServeMux()
	mux.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		if err != nil {
			log.Println(err)
			return
		}
		defer c.Close(websocket.StatusInternalError, "the sky is falling")
		ctx, cancel := context.WithTimeout(r.Context(), time.Hour*2)
		defer cancel()

		ctx = c.CloseRead(ctx)

		t := time.NewTicker(time.Second * 30)
		defer t.Stop()

		for {
			select {
			case <-ctx.Done():
				c.Close(websocket.StatusNormalClosure, "")
				return
			case p := <-receiver:
				err = wsjson.Write(ctx, c, p)
				if err != nil {
					log.Println(err)
					return
				}
			case <-t.C:
				err = wsjson.Write(ctx, c, "PING")
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	})

	p.s.Handler = mux

	go func() {
		err := p.s.ListenAndServe()
		log.Println(err)
	}()

	go func() {

	}()

	return receiver, nil
}

func (p *Publisher) Close(ctx context.Context) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.s.Shutdown(ctx)
}
