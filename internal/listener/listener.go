package listener

import (
	"context"
	"log"
	"net"
	"sync"
)

type Listener struct {
	mu      sync.Mutex
	running chan struct{}
	Addr    net.UDPAddr
}

func New(addr net.UDPAddr) (*Listener, error) {
	return &Listener{
		Addr: addr,
	}, nil
}

func (l *Listener) Accept(ctx context.Context) <-chan []byte {
	l.mu.Lock()
	defer l.mu.Unlock()

	publisher := make(chan []byte, 10)
	l.running = make(chan struct{})

	go func() {
		defer l.stopped()

		dataReceiver := make(chan []byte, 10)

		go func() {
			log.Printf("Starting UDP listener %v", l.Addr)

			conn, err := net.ListenUDP("udp", &l.Addr)
			if err != nil {
				panic(err)
			}
			for {
				buf := make([]byte, 2048)
				len, _, err := conn.ReadFromUDP(buf)
				if err != nil {
					panic(err)
				}
				dataReceiver <- buf[:len]
			}
		}()

	RunLoop:
		for {
			select {
			case <-ctx.Done():
				break RunLoop
			case <-l.running:
				break RunLoop
			case data := <-dataReceiver:
				publisher <- data
			}
		}
	}()

	return publisher
}

func (l *Listener) Stop(ctx context.Context) error {
	l.stopped()

	return nil
}

func (l *Listener) stopped() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.running != nil {
		close(l.running)
		l.running = nil
	}
}
