package event

import (
	"fmt"
	"github.com/wubba-com/L2/develop/dev11/internal/app/domain"
	"github.com/wubba-com/L2/develop/dev11/pkg/middleware"
	"github.com/wubba-com/L2/develop/dev11/pkg/res"
	"github.com/wubba-com/L2/develop/dev11/pkg/validation"
	"net/http"
)

const (
	// paths
	eventByDay   = "/events_for_day"
	eventByWeek  = "/events_for_week"
	eventByMonth = "/events_for_month"
	eventCreate  = "/create_event"
	eventUpdate  = "/update_event"
	eventDelete  = "/delete_event"

	// Day daily events
	Day = 1
	// Week weekly events
	Week = 7
	// Month monthly events
	Month = 30

	// pattern date
	layoutDate = "2006-01-02"
	// date on the calendar
	queryDateParam = "date"
)

func NewHandler(service domain.EventService) Handler {
	return &handler{s: service}
}

type UIDEvent int

type handler struct {
	s domain.EventService
}

// Register - регистрирует event-обработчиков
func (h *handler) Register(mux *http.ServeMux) {
	// GET method
	mux.HandleFunc(eventByDay, middleware.Use(h.byDay, middleware.Log))
	mux.HandleFunc(eventByWeek, middleware.Use(h.byWeek, middleware.Log))
	mux.HandleFunc(eventByMonth, middleware.Use(h.byMonth, middleware.Log))

	// POST method
	mux.HandleFunc(eventCreate, middleware.Use(h.create, middleware.Log))
	mux.HandleFunc(eventUpdate, middleware.Use(h.update, middleware.Log))
	mux.HandleFunc(eventDelete, middleware.Use(h.delete, middleware.Log))
}

func (h *handler) byDay(w http.ResponseWriter, r *http.Request) {
	//  Get user date
	date := r.URL.Query().Get(queryDateParam)

	// transform user-date in time.Time
	currentDate, err := validation.IsDatePattern(layoutDate, date)
	if err != nil {
		// send
		res.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// init input data
	input := &domain.EventDTOGet{Date: currentDate, Stage: uint8(Day)}

	// sending in business logic
	events, err := h.s.ByDate(r.Context(), input)
	if err != nil {
		fmt.Println(err)
		// send
		res.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// init event response
	msg := &domain.EventRes{Msg: events}

	// send
	res.JSON(w, http.StatusOK, msg)
}

func (h *handler) byWeek(w http.ResponseWriter, r *http.Request) {
	//  Get user date
	date := r.URL.Query().Get(queryDateParam)

	// transform user-date in time.Time
	currentDate, err := validation.IsDatePattern(layoutDate, date)
	if err != nil {
		res.ERROR(w, http.StatusBadRequest, domain.ErrNotFoundRecord)
		return
	}

	// init input data
	input := &domain.EventDTOGet{Date: currentDate, Stage: uint8(Week)}

	// sending in business logic
	events, err := h.s.ByDateWithInterval(r.Context(), input)
	if err != nil {
		fmt.Println(err)
		res.ERROR(w, http.StatusInternalServerError, domain.ErrNotFoundRecord)
		return
	}

	// init event response
	msg := domain.EventRes{Msg: events}

	// send
	res.JSON(w, http.StatusOK, msg)
}

func (h *handler) byMonth(w http.ResponseWriter, r *http.Request) {
	//  Get user date
	date := r.URL.Query().Get(queryDateParam)

	// transform user-date in time.Time
	currentDate, err := validation.IsDatePattern(layoutDate, date)
	if err != nil {
		res.ERROR(w, http.StatusBadRequest, domain.ErrNotFoundRecord)
		return
	}

	// init input data
	input := &domain.EventDTOGet{Date: currentDate, Stage: uint8(Month)}

	// sending in business logic
	events, err := h.s.ByDateWithInterval(r.Context(), input)
	if err != nil {
		res.ERROR(w, http.StatusServiceUnavailable, err)
		return
	}

	// init event response
	msg := &domain.EventRes{Msg: events}

	// send
	res.JSON(w, http.StatusOK, msg)
}

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	// init domain struct
	e := &domain.Event{}

	// r.Body => domain struct
	err := res.Bind(r, &e)
	if err != nil {
		fmt.Println(err)
		res.ERROR(w, http.StatusInternalServerError, domain.ErrNotCreatedRecord)
		return
	}

	// Name event is empty?
	if err = validation.IsEmpty(e.Name); err != nil {
		fmt.Println(err)
		res.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// store event
	uid, err := h.s.StoreEvent(r.Context(), e)
	if err != nil {
		res.ERROR(w, http.StatusServiceUnavailable, domain.ErrNotCreatedRecord)
		return
	}

	// init event response
	msg := &domain.EventRes{Msg: &uid}

	// send
	res.JSON(w, http.StatusOK, msg)
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	// init domain struct
	e := &domain.Event{}

	// r.Body => domain struct
	err := res.Bind(r, e)
	if err != nil {
		fmt.Println(err)
		res.ERROR(w, http.StatusInternalServerError, domain.ErrNotUpdatedRecord)
		return
	}

	// Name event is empty?
	if err = validation.IsEmpty(e.Name); err != nil {
		fmt.Println(err)
		res.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// updated event
	uid, err := h.s.UpdateEvent(r.Context(), e)
	if err != nil {
		fmt.Println(err)
		// send
		res.ERROR(w, http.StatusServiceUnavailable, domain.ErrNotUpdatedRecord)
		return
	}

	// init event response
	msg := &domain.EventRes{Msg: &uid}

	// send
	res.JSON(w, http.StatusOK, msg)
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	// init domain struct
	input := &domain.EventDTODelete{}

	// r.Body => event-input
	err := res.Bind(r, input)
	if err != nil {
		fmt.Println(err)
	}

	// delete event
	ok := h.s.DeleteEvent(r.Context(), input)
	if ok {
		// init event response
		e := &domain.EventRes{Msg: domain.StatusDeleteMsg}
		//send
		res.JSON(w, http.StatusOK, e)
		return
	}

	// send
	res.ERROR(w, http.StatusServiceUnavailable, domain.ErrNotDeleteRecord)
}
