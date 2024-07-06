package kvdb

// БД "ключ-значение"
type DB map[string]string

// Конструктор объекта БД.
func New() DB {
	return make(map[string]string)
}

// Получение элемента из БД по ключу.
func (db DB) GET(key string) string {
	if val, ok := db[key]; ok {
		return val
	}
	return ""
}

// Запись в БД значения для ключа.
func (db DB) SET(key string, val string) {
	db[key] = val
}
