package event

import (
	"context"
	"fmt"
	"github.com/wubba-com/L2/develop/dev11/internal/app/domain"
	"github.com/wubba-com/L2/develop/dev11/pkg/client/builder"
	"github.com/wubba-com/L2/develop/dev11/pkg/client/pg"
)

func NewRepository(db pg.Client, queryBuilder build.SQLQueryBuilder) domain.EventRepository {
	return &repository{db: db, builder: queryBuilder}
}

const (
	table = "events"
)

type repository struct {
	builder build.SQLQueryBuilder
	db      pg.Client
}

func (r repository) GetByDate(ctx context.Context, event *domain.Event) (domain.Events, error) {
	query := r.builder.Select(table, []string{"event_uid", "name", "date"}).Where("date", "=", "$1").Get() //

	rows, err := r.db.Query(ctx, query, &event.Date)
	if err != nil {
		fmt.Println(err)
		return nil, domain.ErrNotFoundRecord
	}

	events := domain.Events{}
	for rows.Next() {
		e := &domain.Event{}
		if err := rows.Scan(&e.UID, &e.Name, &e.Date); err != nil {
			fmt.Println(err)
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func (r repository) GetByDateWithInterval(ctx context.Context, input *domain.EventDTOGet) (domain.Events, error) {
	query := r.builder.
		Select(table, []string{"event_uid", "name", "date"}).
		WhereAnd("date", ">=", "$1", "date", "<=", "$2").
		OrderBy("date", "asc").
		Get()

	rows, err := r.db.Query(ctx, query, &input.Date, &input.MaxDate)
	if err != nil {
		fmt.Println(err)
		return nil, domain.ErrNotFoundRecord
	}

	events := domain.Events{}
	for rows.Next() {
		e := &domain.Event{}
		if err := rows.Scan(&e.UID, &e.Name, &e.Date); err != nil {
			fmt.Println(err)
			return nil, domain.ErrNotFoundRecord
		}

		events = append(events, e)
	}

	return events, nil
}

func (r repository) Store(ctx context.Context, event *domain.Event) (int, error) {
	var uid int
	fields := []string{"name", "date"}
	query := r.builder.Insert(table, fields).Get()
	query += fmt.Sprintf(" RETURNING event_uid")

	err := r.db.QueryRow(ctx, query, &event.Name, &event.Date).Scan(&uid)
	if err != nil {
		return 0, domain.ErrNotCreatedRecord
	}
	return uid, nil
}

func (r repository) Update(ctx context.Context, event *domain.Event) (int, error) {
	var uid int
	query := fmt.Sprintf("UPDATE %s SET name = $1, date = $2 WHERE event_uid = $3 RETURNING event_uid;", table)
	err := r.db.QueryRow(ctx, query, &event.Name, &event.Date, &event.UID).Scan(&uid)
	if err != nil {
		return 0, domain.ErrNotUpdatedRecord
	}

	return uid, nil
}

func (r repository) Delete(ctx context.Context, input *domain.EventDTODelete) bool {
	query := r.builder.Delete(table, "event_uid").Returning("event_uid").Get()
	var uid int

	err := r.db.QueryRow(ctx, query, &input.UID).Scan(&uid)
	if err != nil {
		return false
	}

	return uid == input.UID
}
