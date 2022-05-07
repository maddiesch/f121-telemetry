package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/maddiesch/telemetry/internal/listener"
	"github.com/maddiesch/telemetry/internal/packet"
	"github.com/maddiesch/telemetry/internal/publisher"
)

// type ui struct {
// 	CarTelemetry telemetry.CarTelemetryData
// 	LapData      telemetry.LapData
// }

// func displayMS(ms uint32) string {
// 	min := ms / 60000
// 	sec := (ms - (min * 60000)) / 1000
// 	rem := (ms - (min * 60000)) % 1000
// 	return fmt.Sprintf("%02d:%02d.%03d", min, sec, rem)
// }

// func (i *ui) print() {
// 	t := i.CarTelemetry
// 	ld := i.LapData
// 	fmt.Print("\033[H\033[2J")
// 	fmt.Print("\033[0;0H")
// 	fmt.Print("F1 - 2021")
// 	fmt.Printf("\n%d KPH %d %d %0.2f %0.2f", t.Speed, t.Gear, t.EngineRPM, t.Throttle, t.Brake)
// 	fmt.Printf("\n%d: %s %s (%0.1f/%0.1f)", ld.LapNumber, displayMS(ld.LastLapTimeMS), displayMS(ld.CurrentLapTimeMS), ld.LapDistance, ld.TotalDistance)
// }

func main() {
	addr := net.UDPAddr{
		Port: 20777,
		IP:   net.ParseIP("0.0.0.0"),
	}

	r := &Recorder{}

	if err := r.Open(); err != nil {
		log.Fatal(err)
	}

	l, err := listener.New(addr)
	if err != nil {
		log.Fatal(err)
	}

	p, err := publisher.New()
	if err != nil {
		log.Fatal(err)
	}

	dataReceiver := l.Accept(context.Background())

	dataPublisher, err := p.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for {
		select {
		case data := <-dataReceiver:
			r.Record(data)

			packet, err := packet.Decode(data)
			if err != nil {
				log.Println(err)
			} else {
				dataPublisher <- packet
			}
		case <-signalChan:
			log.Println("Signal Interrupt")

			if err := l.Stop(context.Background()); err != nil {
				log.Println(err)
			}

			if err := r.Close(); err != nil {
				log.Println(err)
			}

			os.Exit(0)
		}
	}
}
