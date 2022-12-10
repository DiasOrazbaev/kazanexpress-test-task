package service

import (
	"context"
	"encoding/json"
	"github.com/DiasOrazbaev/kazanexpress-test-task/internal/batch"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

// Service - defines external service that can process batches of items.
type Service interface {
	GetLimits() (n uint64, t time.Duration)               // limit
	Process(ctx context.Context, batch batch.Batch) error // process batch
}

// ExternalService - implementation of our service
type ExternalService struct {
	MaxItemCount  uint64
	ProcessPeriod time.Duration
	status        atomic.Bool
}

func NewExternalService(maxItemCount uint64, processPeriod time.Duration) *ExternalService {
	return &ExternalService{MaxItemCount: maxItemCount, ProcessPeriod: processPeriod, status: atomic.Bool{}}
}

func (e *ExternalService) GetLimits() (n uint64, t time.Duration) {
	return e.MaxItemCount, e.ProcessPeriod
}

func (e *ExternalService) Process(ctx context.Context, batch batch.Batch) error {
	if e.status.Load() || uint64(len(batch)) > e.MaxItemCount {
		return ErrBlocked
	}
	e.status.Store(true)
	log.Println("[ExternalService] start work with batch sized", len(batch))

	go func(e *ExternalService) {
		select {
		case <-ctx.Done():
			log.Println("[canceled] canceled")
			e.status.Store(false)
		case <-time.After(e.ProcessPeriod):
			log.Printf("[success] succedded processed %d items batch\n", len(batch))
			e.status.Store(false)
		}
	}(e)

	return nil
}

func (e *ExternalService) BatchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b batch.Batch

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("bad body structure"))
		if err != nil {
			return
		}
		return
	}

	if err := e.Process(context.Background(), b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Service started processing your batches, just wait..."))
	if err != nil {
		return
	}
}
