package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {
	Task   string
	Status string
}

type Db struct {
	pool *pgxpool.Pool
}

func New(user, password, dbName, host string, port int) (*Db, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, dbName)
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return &Db{pool: pool}, nil
}

func (db *Db) InsertItem(ctx context.Context, item Item) error {
	query := `INSERT INTO todo_Items (task, status) VALUES  ($1, $2)`
	_, err := db.pool.Exec(ctx, query, item.Task, item.Status)
	return err
}

func (db *Db) GetAllItems(ctx context.Context) ([]Item, error) {
	query := `SELECT task,status FROM todo_items`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Task, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return items, nil

}
func (db *Db) Close() {
	db.pool.Close()
}
