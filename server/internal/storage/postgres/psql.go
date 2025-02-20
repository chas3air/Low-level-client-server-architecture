package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"server/internal/domain/models"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	TableName string
	DB        *sql.DB
	log       *slog.Logger
}

// Создание нового подключения к базе данных
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
		log.Error("Failed to open database connection", slog.String("operation", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	time.Sleep(100 * time.Millisecond)
	err = db.Ping()
	if err != nil {
		log.Warn("Failed to ping database", slog.String("operation", op), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Database connected successfully", slog.String("operation", op))
	return &PostgresDB{
		TableName: tablename,
		DB:        db,
		log:       log,
	}, nil
}

func (p *PostgresDB) Stop() error {
	log := p.log.With(slog.String("operation", "storage.postgres.Stop"))
	if err := p.DB.Close(); err != nil {
		log.Warn("Failed to close database connection", slog.String("error", err.Error()))
		return err
	}
	log.Info("Database connection closed successfully")
	return nil
}

func (p *PostgresDB) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "storage.postgres.GetUsers"
	log := p.log.With(slog.String("op", op))

	rows, err := p.DB.QueryContext(ctx, "SELECT * FROM "+p.TableName+";")
	if err != nil {
		log.Warn("Error querying users", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var users_from_db []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick); err != nil {
			log.Warn("Error scanning user row", slog.String("error", err.Error()))
			continue
		}
		users_from_db = append(users_from_db, user)
	}

	log.Info("Fetched users successfully", slog.Int("count", len(users_from_db)))
	return users_from_db, nil
}

func (p *PostgresDB) GetUserById(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "storage.postgres.GetUserById"
	log := p.log.With(slog.String("op", op))

	row := p.DB.QueryRowContext(ctx, "SELECT * FROM "+p.TableName+" WHERE id=$1", uid)
	var user models.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick); err != nil {
		log.Warn("Error retrieving user by ID", slog.String("userId", uid.String()), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User retrieved successfully", slog.String("userId", uid.String()))
	return user, nil
}

func (p *PostgresDB) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "storage.postgres.GetUserByEmail"
	log := p.log.With(slog.String("op", op))

	row := p.DB.QueryRowContext(ctx, "SELECT * FROM "+p.TableName+" WHERE email=$1", email)
	var user models.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick); err != nil {
		log.Warn("Error retrieving user by email", slog.String("email", email), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User retrieved successfully", slog.String("email", email))
	return user, nil
}

func (p *PostgresDB) Insert(ctx context.Context, user models.User) error {
	const op = "storage.postgres.Insert"
	log := p.log.With(slog.String("op", op))

	result, err := p.DB.ExecContext(ctx,
		"INSERT INTO "+p.TableName+" VALUES($1, $2, $3, $4, $5)",
		user.Id, user.Email, user.Password, user.Role, user.Nick,
	)
	if err != nil {
		log.Warn("Error inserting user", slog.Any("user", user), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Warn("No rows affected during insert operation", slog.Any("user", user))
		return fmt.Errorf("%s: %w", op, errors.New("no rows affected"))
	}

	log.Info("User inserted successfully", slog.Any("user", user))
	return nil
}

func (p *PostgresDB) Update(ctx context.Context, uid uuid.UUID, user models.User) error {
	const op = "storage.postgres.Update"
	log := p.log.With(slog.String("op", op))

	result, err := p.DB.ExecContext(ctx,
		"UPDATE "+p.TableName+" SET email=$1, password=$2, role=$3, nick=$4 WHERE id=$5",
		user.Email, user.Password, user.Role, user.Nick, uid,
	)
	if err != nil {
		log.Warn("Error updating user", slog.String("userId", uid.String()), slog.Any("user", user), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Warn("No rows affected during update operation", slog.String("userId", uid.String()))
		return fmt.Errorf("%s: %w", op, errors.New("no rows affected"))
	}

	log.Info("User updated successfully", slog.String("userId", uid.String()), slog.Any("user", user))
	return nil
}

func (p *PostgresDB) Delete(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "storage.postgres.Delete"
	log := p.log.With(slog.String("op", op))

	var user models.User
	row := p.DB.QueryRowContext(ctx, "SELECT * FROM "+p.TableName+" WHERE id=$1", uid)
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.Nick); err != nil {
		log.Warn("Error retrieving user for deletion", slog.String("userId", uid.String()), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	result, err := p.DB.ExecContext(ctx, "DELETE FROM "+p.TableName+" WHERE id=$1", uid)
	if err != nil {
		log.Warn("Error deleting user", slog.String("userId", uid.String()), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		log.Warn("No rows affected during delete operation", slog.String("userId", uid.String()))
		return models.User{}, fmt.Errorf("%s: %w", op, errors.New("no rows affected"))
	}

	log.Info("User deleted successfully", slog.String("userId", uid.String()), slog.Any("user", user))
	return user, nil
}
