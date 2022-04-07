package event

import (
	"context"
	"fmt"
	"github.com/wubba-com/L2/develop/dev11/internal/app/domain"
	ntpTime "github.com/wubba-com/L2/develop/dev11/pkg/ntp_time"
	"time"
)

func NewService(repository domain.EventRepository, timeout time.Duration) domain.EventService {
	return &service{r: repository, timeout: timeout}
}

type service struct {
	r       domain.EventRepository
	timeout time.Duration
}

func (s *service) ByDate(ctx context.Context, input *domain.EventDTOGet) (domain.Events, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	input.Date = input.Date.Add(time.Duration(input.Stage) * (23 * time.Hour)).Add(59 * time.Minute).Add(59 * time.Second)
	t, err := ntpTime.TimeWithTimeZone(input.Date)
	if err != nil {
		return nil, err
	}
	fmt.Println(t)
	event := &domain.Event{Date: t}

	events, err := s.r.GetByDate(ctx, event)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) ByDateWithInterval(ctx context.Context, input *domain.EventDTOGet) (domain.Events, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	input.MaxDate = input.Date.Add(time.Duration(input.Stage) * (23 * time.Hour))
	t, err := ntpTime.TimeWithTimeZone(input.MaxDate)
	if err != nil {
		return nil, err
	}

	input.MaxDate = t
	events, err := s.r.GetByDateWithInterval(ctx, input)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) StoreEvent(ctx context.Context, event *domain.Event) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	event.Date = event.Date.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	uid, err := s.r.Store(ctx, event)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (s *service) UpdateEvent(ctx context.Context, event *domain.Event) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	uid, err := s.r.Update(ctx, event)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (s *service) DeleteEvent(ctx context.Context, input *domain.EventDTODelete) bool {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.r.Delete(ctx, input)
}
