package app

import (
	"context"
	"fmt"
	event3 "github.com/wubba-com/L2/develop/dev11/internal/app/event/delivery/http"
	event "github.com/wubba-com/L2/develop/dev11/internal/app/event/repository/postgres"
	event2 "github.com/wubba-com/L2/develop/dev11/internal/app/event/service"
	"github.com/wubba-com/L2/develop/dev11/internal/config"
	build "github.com/wubba-com/L2/develop/dev11/pkg/client/builder"
	"github.com/wubba-com/L2/develop/dev11/pkg/client/pg"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// MaxAttempts maximum number of attempts to connect to the database
	MaxAttempts = 3
)

// var config
var cfg *config.Config

// var query string builder
var builder build.SQLQueryBuilder

// timeout - context with timeout
var timeout = 3 * time.Second

func Run() {
	// init signal SIGINT, SIGQUIT
	gFulShotDown := make(chan os.Signal, 1)
	signal.Notify(gFulShotDown, syscall.SIGINT, syscall.SIGQUIT)

	// init config app
	cfg = config.GetConfig()
	// init http-router
	mux := http.NewServeMux()

	// init sql-query builder
	builder = pg.NewPSGQueryBuilder()
	// init postgres client
	client, err := pg.NewClient(context.Background(), cfg, MaxAttempts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// init repository
	r := event.NewRepository(client, builder)
	// init service
	s := event2.NewService(r, timeout)
	// init handler
	h := event3.NewHandler(s)

	// start event-handlers
	h.Register(mux)

	// start servers
	go func() {
		fmt.Printf("init server: %s:%s\n", cfg.Listen.BindIP, cfg.Listen.Port)
		log.Fatalln(http.ListenAndServe(net.JoinHostPort(cfg.Listen.BindIP, cfg.Listen.Port), mux))
	}()

	// wait SIGINT, SIGQUIT
	select {
	case <-gFulShotDown:
		fmt.Println("exit!")
	}
	fmt.Println("done")
}
