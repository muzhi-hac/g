package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

// 通过dialect进行处理，将其转化为可以该数据库可以使用的数据类型
// 在对于某个数据库而言实现该接口的方法即可实现其转化
type Dialect interface {
	DataTypeof(value reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]

	return
}
