package main

import (
	"database/sql"
	"encoding/binary"
	"log"
	"math"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Recorder struct {
	mu sync.Mutex
	db *sql.DB
}

func (r *Recorder) Record(raw []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	packetID := raw[5]
	sessionID := int64(binary.LittleEndian.Uint64(raw[6:14]))
	sessionTime := math.Float32frombits(binary.LittleEndian.Uint32(raw[14:20]))

	if packetID == 8 {
		log.Printf("FINAL CLASSIFICATION: %d", sessionID)
	}

	_, err := r.db.Exec("INSERT INTO session_history (sid, packetType, sessionTime, packet) VALUES (?, ?, ?, ?)", sessionID, packetID, sessionTime, raw)

	return err
}

func (r *Recorder) Open() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	r.db = db

	if _, err := r.db.Exec("CREATE TABLE IF NOT EXISTS session_history (sid INTEGER NOT NULL, packetType INTEGER NOT NULL, sessionTime REAL NOT NULL, packet BLOB);"); err != nil {
		return err
	}

	if _, err := r.db.Exec("CREATE INDEX IF NOT EXISTS session_time ON session_history (sid, sessionTime, packetType);"); err != nil {
		return err
	}

	if _, err := r.db.Exec("DELETE FROM session_history WHERE sid = ?;", 0); err != nil {
		return err
	}

	if _, err := r.db.Exec("VACUUM"); err != nil {
		return err
	}

	return nil
}

func (r *Recorder) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.db.Close()
}
