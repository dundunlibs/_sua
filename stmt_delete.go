package sua

func (db *DB) Delete(table string) Delete {
	return Delete{db: db}
}

type Delete struct {
	db    *DB
	table string
	Where
}

func (i Delete) SQL() (string, []any) {
	return "", nil
	// keys := []string{}
	// params := []string{}
	// values := []any{}

	// idx := 1
	// for key, value := range i.data {
	// 	keys = append(keys, key)
	// 	params = append(params, fmt.Sprintf("$%d", idx))
	// 	values = append(values, value)
	// 	idx++
	// }

	// return fmt.Sprintf(
	// 	"INSERT INTO %q (%s) VALUES (%s)",
	// 	i.table,
	// 	strings.Join(keys, ", "),
	// 	strings.Join(params, ", "),
	// ), values
}
