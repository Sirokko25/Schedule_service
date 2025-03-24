package storage

import (
	"context"
	"time"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"

	"sheduler/internal/models"

)

type DB struct {
	conn *bun.DB
}

type ShedulesStruct struct {
	bun.BaseModel `bun:"table:shedules,alias:sd"`

	SheduleId       int64     `bun:"shedule_id,pk,autoincrement"`
	MedicineName    string    `bun:"medicine_name,notnull"`
	UserId          int64     `bun:"user_id,notnull"`
	Period          int64     `bun:"period,notnull"`
	Duration        string    `bun:"duration,notnull"`
	StartDate       time.Time `bun:"start_date,nullzero,default:now()"`
	EndDate         time.Time `bun:"end_date,notnull"`
	Status          string    `bun:"status,notnull"`
	ReceptionTiming []byte    `bun:"reception_timing,notnull"`
}

type HistoryShedulesStruct struct {
	bun.BaseModel `bun:"table:shedules_history,alias:hsd"`

	SheduleId    int64     `bun:"shedule_id,pk,autoincrement"`
	MedicineName string    `bun:"medicine_name,notnull"`
	UserId       int64     `bun:"user_id,notnull"`
	StartDate    time.Time `bun:"start_date,notnull"`
	EndDate      time.Time `bun:"end_date,notnull"`
}

type StorageInterface interface {
	AppendShedule(models.Shedule) (int64, error)
	FindShedule(string, string) (models.SheduleWithChart, error)
	FindShedules(string) ([]int64, error)
	NextTakings(string) ([]models.Medicine, error)

	PingDB(context.Context) error
	Actualize()
}
