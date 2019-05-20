package postgres

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type Postgres struct {
	session *gorm.DB
}

func NewPostgres(databaseURL string) (*Postgres, error) {
	if strings.TrimSpace(databaseURL) == "" {
		return nil, errors.New("databaseURL is required")
	}

	db, err := gorm.Open("postgres", databaseURL)
	defer db.Close()

	if err != nil {
		return nil, fmt.Errorf("error on Postgres connection: %q", err)
	}

	return &Postgres{
		session: db,
	}, nil

}

func (pg *Postgres) Insert(item interface{}) {
	pg.session.Create(&item)
}

func (pg *Postgres) Update(item interface{}) {
	pg.session.Save(&item)
}

func (pg *Postgres) FindOne(item interface{}) {
	pg.session.Take(&item)
}

func (pg *Postgres) FindAll(items interface{}) {
	pg.session.Find(&items)
}

func (pg *Postgres) RemoveOne(item interface{}) {
	pg.session.Delete(&item)
}

func (pg *Postgres) ExecuteQuery(query string) {
	pg.session.Exec(query)
}
