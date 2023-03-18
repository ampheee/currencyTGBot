package repository

import (
	storage "_entryTask/internal/storage"
	"_entryTask/pkg/logger"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/rs/zerolog"
)

type DB struct {
	pool   *pgxpool.Pool
	logger zerolog.Logger
}

func (db *DB) Create(ctx context.Context, update *tgb.Update) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbCreateUser] Unable to acquire a db conn")
		return err
	}
	defer conn.Release()
	q := `INSERT INTO users (u_id, username, firstame, lastname) VALUES ($1, $2, $3, $4)`
	_, err = conn.Exec(ctx, q,
		update.Message.From.ID,
		update.Message.From.Username,
		update.Message.From.FirstName,
		update.Message.From.LastName)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbCreateUser] Unable to create user")
		return err
	}
	db.logger.Info().Msg("[dbCreateUser] User created")
	return nil
}

func (db *DB) FindAllRequestsById(ctx context.Context, userId tg.UserID) (data []string, err error) {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbFindUser] Unable to acquire a db conn")
		return nil, err
	}
	defer conn.Release()
	qCheck := `SELECT EXISTS(SELECT u_id FROM users WHERE u_id = ($1))`
	ct, err := conn.Query(ctx, qCheck, userId)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbFindUser] Unable to found user")
	}
	var found bool
	ct.Scan(&found)
	if !found {
		//q := `SELECT r.r_id, r.r_type, r.r_time, r.r_response
		//	FROM users u JOIN requests r ON u.u_id = r.u_id WHERE u.u_id = ($1)`
		//ct := conn.QueryRow(ctx, q, userId)
		//if err != nil {
		//	db.logger.Warn().Err(err)
		//}
		//var record
		//for ct.Scan() {
		//
		//}
	}
	return nil, nil
}

func (db *DB) FindFirstRequest(ctx context.Context, userId int) (string, error) {
	//TODO Implement me!
	return "", nil
}

func (db *DB) AddNewRequest(ctx context.Context, update *tgb.Update) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbAddNewRequest] Unable to acquire a db conn")
		return err
	}
	defer conn.Release()
	qCheck := `SELECT EXISTS(SELECT u_id FROM users WHERE u_id = ($1))`
	ct, err := conn.Query(ctx, qCheck, update.Message.From.ID)
	if err != nil {
		db.logger.Warn().Err(err).Msg("[dbAddNewRequest]  Unable to found user")
	}
	var found bool
	ct.Scan(&found)
	if !found {
		if err := db.Create(ctx, update); err != nil {
			db.logger.Warn().Err(err).Msg("[dbAddNewRequest] Unable to found && create user")
		}
	}
	q := `INSERT INTO requests(r_id, r_type, ) VALUES ()`
}

func NewStorage(pool *pgxpool.Pool) storage.Storage {
	return &DB{pool: pool, logger: logger.GetLogger()}
}
