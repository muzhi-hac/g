package schema

import (
	"Gee/geeorm/dialect"
	"Gee/geeorm/log"
	"go/ast"
	"reflect"
)

// 此文件用于将结构体转换为表结构
// Field 结构体即用于其结构体的一条信息
type Field struct {
	Name string
	Type string
	Tag  string
}
type Schema struct {
	Model     interface{}
	Name      string
	Fields    []*Field
	FieldName []string
	fieldMap  map[string]*Field
	//其实这里可以通过遍历FIelds进行操作，但使用FieldMap可以提升效率？
}

// func (s *Schema) GetField(name string) *Field {
//
//	return s.fieldMap[name] //就蛮奇怪的
//
// }
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 将任意的对象解析为Schema实例
func Parse(value interface{}, dialect dialect.Dialect) *Schema {
	//reflect.Indirect 用于获取指针或接口包含的实际值。
	//如果传入的 Value 不是指针或接口，reflect.Indirect 返回原始 Value。
	//如果传入的 Value 是指针或接口，reflect.Indirect 返回指针或接口包含的实际值的 Value。

	modelType := reflect.Indirect(reflect.ValueOf(value)).Type()
	schema := &Schema{
		Model:    value,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		F := modelType.Field(i)
		if !F.Anonymous && ast.IsExported(F.Name) {
			field := &Field{
				Name: F.Name,
				Type: dialect.DataTypeof(reflect.Indirect(reflect.New(F.Type))),
			}
			var ok bool
			field.Tag, ok = F.Tag.Lookup("geeorm")
			if !ok {
				log.Info("Fail to get Tag")
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldName = append(schema.FieldName, F.Name)
			schema.fieldMap[F.Name] = field
		}
	}
	return schema
}
