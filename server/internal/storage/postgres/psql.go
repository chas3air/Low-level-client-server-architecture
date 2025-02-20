package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"server/internal/domain/models"
	"server/pkg/lib/logger/sl"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	TableName string
	DB        *sql.DB
	log       *slog.Logger
}

// Не подключается к базе данных
func New(host string, user string, password string, port int, dbname string, tablename string, log *slog.Logger) (*PostgresDB, error) {
	const op = "storage.postgres.New"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		//return nil, fmt.Errorf("%s: %w", op, err)
		panic(err)
	}

	time.Sleep(100 * time.Millisecond)
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PostgresDB{
		TableName: tablename,
		DB:        db,
		log:       log,
	}, nil
}

func (p *PostgresDB) Stop() error {
	return p.DB.Close()
}

func (p *PostgresDB) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "storage.postgres.getUsers"
	log := p.log.With(
		slog.String("op", op),
	)

	rows, err := p.DB.QueryContext(ctx,
		"SELECT * FROM "+p.TableName+";",
	)
	if err != nil {
		log.Warn("%s: %w", op, err)

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var users_from_db []models.User = make([]models.User, 0, 5)
	var user models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick)
		if err != nil {
			log.Warn("error scanning row:", sl.Err(err))

			continue
		}

		users_from_db = append(users_from_db, user)
	}

	return users_from_db, nil
}

func (p *PostgresDB) GetUserById(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "storage.postgres.getUserById"
	log := p.log.With(
		slog.String("op", op),
	)

	row := p.DB.QueryRowContext(ctx,
		"SELECT * FROM "+p.TableName+" WHERE id=$1", uid,
	)

	var user models.User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick)
	if err != nil {
		log.Warn("error scanning row:", sl.Err(err))

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (p *PostgresDB) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "storage.postgres.getUserByEmail"
	log := p.log.With(
		slog.String("op", op),
	)

	row := p.DB.QueryRowContext(ctx,
		"SELECT * FROM "+p.TableName+" WHERE email=$1", email,
	)

	var user models.User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick)
	if err != nil {
		log.Warn("error scanning row:", sl.Err(err))

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (p *PostgresDB) Insert(ctx context.Context, user models.User) error {
	const op = "storage.postgres.getUserById"
	log := p.log.With(
		slog.String("op", op),
	)

	result, err := p.DB.ExecContext(ctx,
		"INSERT INTO "+p.TableName+
			" VALUES($1, $2, $3, $4, $5)",
		user.Id, user.Email, user.Password, user.Role, user.Nick,
	)
	if err != nil {
		log.Warn("error inserting record to table:", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Warn("error insering record to table:", sl.Err(err))

		return fmt.Errorf("%s: %w", op, errors.New("rowsAffected is zero"))
	}

	return nil
}

func (p *PostgresDB) Update(ctx context.Context, uid uuid.UUID, user models.User) error {
	const op = "storage.postgres.getUserById"
	log := p.log.With(
		slog.String("op", op),
	)

	result, err := p.DB.ExecContext(ctx,
		"UPDATE "+p.TableName+
			"SET email=$1 password=$2 role=$3 nick=$4",
		"WHERE id=$5",
		user.Email, user.Password, user.Role, user.Nick, user.Id,
	)
	if err != nil {
		log.Warn("error updating record:", sl.Err(err))

		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Warn("error updating record:", sl.Err(err))

		return fmt.Errorf("%s: %w", op, errors.New("rowsAffected is zero"))
	}

	return nil
}

func (p *PostgresDB) Delete(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "storage.postgres.getUserById"
	log := p.log.With(
		slog.String("op", op),
	)

	var user models.User
	row := p.DB.QueryRowContext(ctx,
		"SELECT * FROM "+p.TableName+
			" WHERE id=$1", uid,
	)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick)
	if err != nil {
		log.Warn("error retrieving record by id before deleting:", sl.Err(err))

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	result, err := p.DB.ExecContext(ctx,
		"DELETE FROM "+p.TableName+
			" WHERE id=$1",
		uid,
	)
	if err != nil {
		log.Warn("error deleting record:", sl.Err(err))

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Warn("error deleting record:", sl.Err(err))

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
