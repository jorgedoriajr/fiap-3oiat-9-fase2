package querymapper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type UpdateQueryCommand struct {
	Table      string
	Condition  string
	UpdateData any
}

func GenerateUpdateQuery(updateCommand UpdateQueryCommand) (string, []any) {
	v := reflect.ValueOf(updateCommand.UpdateData)
	t := v.Type()

	var setClauses []string
	var values []any

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		if value != nil {
			setClauses = append(setClauses, field.Name+" = $"+strconv.Itoa(len(values)+1))
			values = append(values, value)
		}
	}

	setClause := strings.Join(setClauses, ", ")
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", updateCommand.Table, setClause, updateCommand.Condition)

	return sql, values
}
