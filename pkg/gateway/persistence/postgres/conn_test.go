package postgres

import "testing"

func TestOpen(t *testing.T) {
	conn := NewConnection("127.0.0.1", 5432, "docker", "docker", "routine")
	db, err := conn.Open()
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
