package postgres

import (
	"bookmaker/internal/sport"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

var _ sport.Sports = &SportStorage{}

// RobotStorage ...
type SportStorage struct {
	statementStorage

	upsertStmt *sql.Stmt
	findStmt   *sql.Stmt
}

// NewRobotStorage ...
func NewSportStorage(db *DB) (*SportStorage, error) {
	s := &SportStorage{statementStorage: newStatementsStorage(db)}

	stmts := []stmt{
		{Query: upsertSportQuery, Dst: &s.upsertStmt},
		{Query: findSportQuery, Dst: &s.findStmt},
	}

	if err := s.initStatements(stmts); err != nil {
		return nil, errors.Wrap(err, "can't init statements")
	}

	return s, nil
}

const sportFields = "sport, coefficient"

const upsertSportQuery = "INSERT INTO public.sports (" + sportFields + ") VALUES ($1, $2) ON CONFLICT (sport) DO UPDATE SET sport = excluded.sport, coefficient = excluded.coefficient"

// Create or Update
func (s *SportStorage) Upsert(sport *sport.Sport) error {
	if _, err := s.upsertStmt.Exec(&sport.Sport, &sport.Coefficient); err != nil {
		msg := fmt.Sprintf("can not exec query with sport %v", sport.Sport)
		return errors.WithMessage(err, msg)
	}

	return nil
}

const findSportQuery = "SELECT " + sportFields + " FROM public.sports WHERE sport=ANY($1::text[]) ORDER BY sport"

// Find sport lines
func (s *SportStorage) Find(name []string) ([]*sport.Sport, error) {
	var sports []*sport.Sport

	rows, err := s.findStmt.Query(pq.Array(name))
	if err != nil {
		msg := fmt.Sprintf("can't scan sports %v", name)
		return nil, errors.WithMessage(err, msg)
	}

	defer rows.Close()

	for rows.Next() {
		var sp sport.Sport

		if err := scanSport(rows, &sp); err != nil {
			msg := fmt.Sprintf("can't scan sports %v", name)
			return nil, errors.WithMessage(err, msg)
		}

		sports = append(sports, &sp)
	}

	return sports, nil
}

func scanSport(scanner sqlScanner, s *sport.Sport) error {
	return scanner.Scan(&s.Sport, &s.Coefficient)
}
