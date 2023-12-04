package session

import (
	"Gee/geeorm/dialect"
	"Gee/geeorm/log"
	"Gee/geeorm/schema"
	"database/sql"
	"strings"
)

// 这个包用于与原生交互的部分
type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db,
		dialect: dialect}

}
func Newsession() {

}
func (s *Session) Raw(raw string, value ...interface{}) *Session {
	s.sql.WriteString(raw)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, value...)
	return s

}
func (s *Session) DB() *sql.DB {
	return s.db
}
func (s *Session) Clear() {
	s.sql.Reset() //清空session中的sql命令，可以重复使用Session
	//使用strings.Builder更有优势其使用了一个[]byte缓冲区
	//err := s.db.Close()
	//if err != nil {
	//return
	//}

}
func (s *Session) Exec() (sql.Result, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	result, err := s.DB().Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)

	return s.db.QueryRow(s.sql.String(), s.sqlVars...)

}
func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)

	rows, err := s.db.Query(s.sql.String(), s.sqlVars...)

	return rows, err
}
