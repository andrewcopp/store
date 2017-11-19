package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	config *Config
	db     *sql.DB
}

func NewPostgres(config *Config) *Postgres {
	return &Postgres{
		config: config,
	}
}

func (p *Postgres) Connect() error {

	host := p.config.Host
	port := p.config.Port
	user := p.config.User
	pass := p.config.Password
	dbname := p.config.DBName

	tmpl := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	info := fmt.Sprintf(tmpl, host, port, user, pass, dbname)

	var err error
	p.db, err = sql.Open("postgres", info)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Ping() error {
	return p.db.Ping()
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) Find(id int) (map[string]interface{}, error) {
	row := p.db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
	var trial []byte
	row.Scan(
		&trial,
		&trial,
	)
	os.Stdout.Write(trial)
	fmt.Println(trial)
	return map[string]interface{}{"title": "Test"}, nil
}

func (p *Postgres) Search(off int, lim int) ([]map[string]interface{}, error) {
	rows, err := p.db.Query("SELECT title FROM posts")
	if err != nil {
		return []map[string]interface{}{}, err
	}

	result := []map[string]interface{}{}
	for rows.Next() {
		var title string
		rows.Scan(
			&title,
		)
		result = append(result, map[string]interface{}{"title": title})
	}
	return result, nil
}

func (p *Postgres) Create(obj map[string]interface{}) error {
	return nil
}

func (p *Postgres) Update(id int, obj map[string]interface{}) error {
	return nil
}

func (p *Postgres) Delete(id int) error {
	return nil
}
