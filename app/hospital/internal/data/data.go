package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/app/hospital/internal/conf"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewTransaction, NewHospitalSettingRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewDB
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "hospital/data/db"))

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.Database.Source, // DSN (data source name)
		DefaultStringSize:         256,                  // Default length of a string type field
		DisableDatetimePrecision:  true,                 // Disable datetime precision, not supported before MySQL 5.6
		SkipInitializeWithVersion: false,                // Automatic configuration based on current MySQL version
	}), &gorm.Config{
		Logger:                                   gl.Default.LogMode(gl.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true, NoLowerCase: true},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
	})

	if err != nil {
		log.Fatalf("Failed opening connection to mysql: %v", err)
	}

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医院设置表'").AutoMigrate(&po.HospitalSetting{}); err != nil {
		log.Fatal(err)
	}

	return db
}

// NewData
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "hospital/data/data"))
	d := &Data{
		db:  db,
		log: l,
	}

	sqlDB, _ := d.db.DB()
	// Set the maximum number of idle connections in the connection pool
	sqlDB.SetMaxIdleConns(10)
	// Set the maximum number of open database connections
	sqlDB.SetMaxOpenConns(100)
	// Set the maximum amount of time a connection can be left idle
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	// Set the maximum duration of time a connection can be reused
	sqlDB.SetConnMaxLifetime(time.Hour)

	cleanup := func() {
		l.Info("message", "closing the data resources")
		if err := sqlDB.Close(); err != nil {
			l.Error(err)
		}
	}
	return d, cleanup, nil
}
