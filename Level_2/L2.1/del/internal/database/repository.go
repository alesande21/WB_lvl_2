package database

import "fmt"

type Repository struct {
	repo string
}

func InitRepo(conn *ConnDb) *Repository {
	fmt.Println(conn.conn)
	return &Repository{repo: "Репозиторий инициализирован..."}
}

func (r Repository) Get() string {
	return r.repo
}
