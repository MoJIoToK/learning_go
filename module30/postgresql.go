package main

import "context"

// запрос на вставку данных
_, err := db.Exec(ctx, `
	INSERT INTO users (name)
	VALUES ($1);
	`,
u.name,
)
if err != nil {
return err
}


...

// запрос на выборку данных
rows, err := db.Query(ctx, `
		SELECT * FROM users ORDER BY id;
	`)
if err != nil {
return nil, err
}

// addBooks добавляет в БД массив книг одной транзакцией.
func addBooks(ctx context.Context, db *pgxpool.Pool, books []book) error {
	// начало транзакции
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	batch := new(pgx.Batch)
	// добавление заданий в пакет
	for _, book := range books {
		batch.Queue(`INSERT INTO books(title, year) VALUES ($1, $2)`, book.Title, book.Year)
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(ctx)
}
