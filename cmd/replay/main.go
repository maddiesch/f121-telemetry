package main

import (
	"database/sql"
	"flag"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/maddiesch/telemetry/internal/packet"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dbPath string
	var sessionID int64
	var rate int64

	flag.StringVar(&dbPath, "path", ":memory:", "the path to the database")
	flag.Int64Var(&sessionID, "sid", 0, "the session id")
	flag.Int64Var(&rate, "rate", 20, "event rate in hertz")

	flag.Parse()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := runQuery(db, rate, sessionID); err != nil {
		log.Fatal(err)
	}
}

func runQuery(db *sql.DB, rate, sid int64) error {
	tickRate := time.Duration(float64(time.Duration(1)*time.Second) / float64(rate))

	ticker := time.NewTicker(tickRate)
	defer ticker.Stop()

	var currentSessionTime float64

	var maxSessionTime float64

	if err := db.QueryRow("SELECT MAX(sessionTime) FROM session_history WHERE sid = ?;", sid).Scan(&maxSessionTime); err != nil {
		return err
	}

	spew.Dump(maxSessionTime)

	for {
		<-ticker.C

		nextSessionTime := currentSessionTime + (float64(tickRate) / float64(time.Second))

		var count uint64
		rows, err := db.Query("SELECT packet, packetType FROM session_history WHERE sid = ? AND sessionTime >= ? AND sessionTime < ? ORDER BY sessionTime ASC, _rowid_ ASC;", sid, currentSessionTime, nextSessionTime)
		if err != nil {
			return err
		}

		for rows.Next() {
			count += 1
			var data []byte
			var packetType uint8

			if err := rows.Scan(&data, &packetType); err != nil {
				return err
			}

			_, err := packet.Decode(data)

			if err != nil {
				log.Printf("Failed to decode packet (%d): %v", packetType, err)
			}
		}

		if count == 0 {
		}

		log.Printf("Processed %d packets", count)

		currentSessionTime = nextSessionTime

		if currentSessionTime > maxSessionTime {
			break
		}
	}

	return nil
}
