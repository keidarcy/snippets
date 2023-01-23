package server

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Activity struct {
	ID          uint64 `json:"id"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

type Activities struct {
	// activities []Activity
	db *sql.DB
}

var ErrIDNotFound = fmt.Errorf("ID nt found")

const file string = "activities.db"

const create string = `
CREATE TABLE IF NOT EXISTS activities (
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT
);
`

func NewActivities() (*Activities, error) {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return &Activities{db}, nil
}

func (c *Activities) Insert(a Activity) (int, error) {
	res, err := c.db.Exec("INSERT INTO activities VALUES(NULL, ?, ?);", a.Time, a.Description)
	if err != nil {
		return 0, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	log.Printf("Getting %d", id)

	row := c.db.QueryRow("SELECT id, time, description FROM activities WHERE id=?;", id)

	activity := Activity{}
	var err error

	if err = row.Scan(&activity.ID, &activity.Time, &activity.Description); err == sql.ErrNoRows {
		log.Printf("Id not Found")
		return activity, ErrIDNotFound
	}

	fmt.Printf("activity: %+v\n", activity)

	return activity, nil
}

func (c *Activities) List(offset int) ([]Activity, error) {
	log.Printf("Getting list from offset %d\n", offset)

	rows, err := c.db.Query("SELECT * FROM activities WHERE ID > ? ORDER BY id DESC LIMIT 100;", offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	activities := []Activity{}

	for rows.Next() {
		activity := Activity{}
		err := rows.Scan(&activity.ID, &activity.Time, &activity.Description)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}
