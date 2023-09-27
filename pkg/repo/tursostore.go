package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	ev "event-todo/pkg/events"
	"fmt"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite" // Use this if you choose the modernc.org/sqlite driver
)

type EventStoreTurso struct {
	db *sql.DB
}

func NewSQLiteEventStore(databaseURL string) (*EventStoreTurso, error) {
	db, err := sql.Open("libsql", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS events (aggregateID TEXT, event BLOB)`)
	if err != nil {
		return nil, fmt.Errorf("could not create table: %v", err)
	}

	return &EventStoreTurso{db: db}, nil
}

func (es *EventStoreTurso) Save(aggregateID string, event ev.Event) error {
	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not marshal event: %v", err)
	}

	_, err = es.db.Exec(`INSERT INTO events (aggregateID, event) VALUES (?, ?)`, aggregateID, eventData)
	if err != nil {
		return fmt.Errorf("could not insert event: %v", err)
	}

	return nil
}

func (es *EventStoreTurso) Load(aggregateID string) ([]ev.Event, error) {
	rows, err := es.db.Query(`SELECT event FROM events WHERE aggregateID = ?`, aggregateID)
	if err != nil {
		return nil, fmt.Errorf("could not query events: %v", err)
	}
	defer rows.Close()

	var events []ev.Event
	for rows.Next() {
		var eventData []byte
		if err := rows.Scan(&eventData); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}

		var event ev.Event
		if err := json.Unmarshal(eventData, &event); err != nil {
			return nil, fmt.Errorf("could not unmarshal event: %v", err)
		}

		events = append(events, event)
	}

	if len(events) == 0 {
		return nil, errors.New("no events found")
	}

	return events, nil
}
