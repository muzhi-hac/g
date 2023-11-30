package geeEngine

import (
	"Gee/geeorm/log"
	"Gee/geeorm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}
func (e *Engine) Close() {
	err := e.db.Close()
	if err != nil {
		log.Error(e)
		return
	}
	log.Info("Close database success")

}
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
