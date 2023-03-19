package repository

import (
	storage "_entryTask/internal/storage"
	"_entryTask/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mr-linch/go-tg"
	"github.com/rs/zerolog"
)

type DB struct {
	pool   *pgxpool.Pool
	logger zerolog.Logger
}

func (db *DB) CreateUser(ctx context.Context, user storage.User) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbCreateUser] Unable to acquire a db conn")
		return err
	}
	defer conn.Release()
	q := `INSERT INTO users (u_id, username, firstname, lastname) VALUES ($1, $2, $3, $4)`
	_, err = conn.Exec(ctx, q,
		user.Id,
		user.Username,
		user.FirstName,
		user.LastName)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbCreateUser] Unable to create user")
		return err
	}
	db.logger.Info().Msg("[dbCreateUser] User created")
	return nil
}

func (db *DB) AddNewRequest(ctx context.Context, request storage.RequestRecord) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbAddNewRequest] Unable to acquire a db conn")
		return err
	}
	defer conn.Release()
	qCheck := `SELECT EXISTS(SELECT u_id FROM users WHERE u_id = ($1))`
	var found bool
	err = conn.QueryRow(ctx, qCheck, request.User.Id).Scan(&found)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbAddNewRequest]  Unable to found user")
	}
	if found == false {
		if err := db.CreateUser(ctx, request.User); err != nil {
			db.logger.Warn().Err(err).Msg("[dbAddNewRequest] Unable to found && create user")
		}
	}
	q := `INSERT INTO requests(u_id, r_type, r_time, r_args) VALUES ($1, $2, $3, $4)`
	_, err = conn.Exec(ctx, q, request.User.Id, request.RequestType,
		request.RequestTime, request.RequestArgs)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbAddNewRequest] Request was not added to db")
	} else {
		db.logger.Info().Msg("[dbAddNewRequest] request added to db")
	}
	return err
}

func (db *DB) DeleteUserStatsById(ctx context.Context, userId tg.UserID) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbDeleteUserStats] Unable to acquire a db conn")
		return err
	}
	defer conn.Release()
	qCheck := `SELECT EXISTS(SELECT u_id FROM users WHERE u_id = ($1))`
	var found bool
	err = conn.QueryRow(ctx, qCheck, userId).Scan(&found)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbDeleteUserStats] Unable to found user")
		return err
	}
	q := `DELETE FROM requests WHERE u_id = ($1)`
	q2 := `DELETE FROM users WHERE u_id = ($1)`
	if _, err := conn.Exec(ctx, q, userId); err != nil {
		db.logger.Warn().Err(err).Msg("[dbDeleteUserStats] Unable to delete user from requests")
		return err
	}
	if _, err := conn.Exec(ctx, q2, userId); err != nil {
		db.logger.Warn().Err(err).Msg("[dbDeleteUserStats] Unable to delete user from users")
		return err
	}
	db.logger.Info().Msg("[dbDeleteUserStats] UserData deleted successfully")
	return nil
}

func (db *DB) FindAllRequestsById(ctx context.Context, userId tg.UserID) ([]string, error) {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbFindAllUserStats] Unable to acquire a db conn")
		return nil, err
	}
	defer conn.Release()
	qCheck := `SELECT EXISTS(SELECT u_id FROM users WHERE u_id = ($1))`
	var found bool
	err = conn.QueryRow(ctx, qCheck, userId).Scan(&found)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbFindAllUserStats] Unable to found user")
	}
	var records []string
	if found {
		db.logger.Info().Msg("[dbFindAllUserStats] User found")
		q := `SELECT r.r_id, r.r_type, r.r_time, COALESCE(NULLIF(r.r_args, ''), 'no args')
			FROM users u JOIN requests r ON u.u_id = r.u_id WHERE u.u_id =($1)`
		rows, err := conn.Query(ctx, q, userId)
		if err != nil {
			db.logger.Warn().Err(err)
		} else {
			for rows.Next() {
				var record storage.RequestRecord
				err := rows.Scan(&record.Id, &record.RequestType, &record.RequestTime, &record.RequestArgs)
				if err != nil {
					db.logger.Warn().Err(err)
				}
				records = append(records, fmt.Sprintf("Type: %s |Args: %s |Time: %s\n",
					record.RequestType, record.RequestArgs, record.RequestTime.Local()))
			}
		}
	} else {
		db.logger.Warn().Msg("[dbFindAllUserStats] User not found :(")
	}
	conn.Release()
	return records, nil
}

func NewStorage(pool *pgxpool.Pool) storage.Storage {
	return &DB{pool: pool, logger: logger.GetLogger()}
}
