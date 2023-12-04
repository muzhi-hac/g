package geeEngine

import (
	"Gee/geeorm/dialect"
	"Gee/geeorm/log"
	"Gee/geeorm/session"
	"database/sql"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
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
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
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
	return session.New(e.db, e.dialect)
}
