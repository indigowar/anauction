package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/indigowar/anauction/domain/service"
)

func checkDuplicationError(err error) *service.DuplicationError {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		if pgError.Code == "32505" {
			return &service.DuplicationError{
				Object: pgError.TableName,
				Field:  pgError.ColumnName,
			}
		}
	}
	return nil
}

func checkForeignKeyViolationError(err error) (bool, string) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23503" {
			return true, pgErr.ColumnName
		}
	}
	return false, ""
}
