package domain

import (
	"context"
	"errors"
	"time"
)

const (
	StatusDeleteMsg = "событие удалено"
	StatusUpdateMsg = "событие обновлено"
	StatusStoreMsg  = "событие создано"
	StatusGetMsg    = "события получены"
)

type Event struct {
	UID  uint      `json:"event_uid"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type Events []*Event

type EventService interface {
	ByDate(context.Context, *EventDTOGet) (Events, error)
	ByDateWithInterval(context.Context, *EventDTOGet) (Events, error)
	StoreEvent(context.Context, *Event) (int, error)
	UpdateEvent(context.Context, *Event) (int, error)
	DeleteEvent(context.Context, *EventDTODelete) bool
}

type EventRepository interface {
	GetByDate(context.Context, *Event) (Events, error)
	GetByDateWithInterval(context.Context, *EventDTOGet) (Events, error)
	Store(context.Context, *Event) (int, error)
	Update(context.Context, *Event) (int, error)
	Delete(context.Context, *EventDTODelete) bool
}

type EventDTOGet struct {
	Stage   uint8     `json:"stage,omitempty"`
	Date    time.Time `json:"date"`
	MaxDate time.Time
}

type EventDTODelete struct {
	UID int `json:"event_uid,omitempty"`
}

type EventRes struct {
	Msg interface{} `json:"result"`
}

var ErrNotFoundRecord = errors.New("события не найдены")
var ErrNotUpdatedRecord = errors.New("событие не обновлено")
var ErrNotCreatedRecord = errors.New("событие не создано")
var ErrNotDeleteRecord = errors.New("событие не удалено")
