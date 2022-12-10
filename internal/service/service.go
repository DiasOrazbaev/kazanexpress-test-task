package service

import (
	"context"
	"github.com/DiasOrazbaev/kazanexpress-test-task/internal/batch"
	"time"
)

// Service может обрабатывать элементы пачками
type Service interface {
	GetLimits() (n uint64, t time.Duration)               // лимит - кол-во элементов в период времени
	Process(ctx context.Context, batch batch.Batch) error // обработать пачку
}
