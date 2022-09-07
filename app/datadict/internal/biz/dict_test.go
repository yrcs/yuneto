package biz_test

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/yrcs/yuneto/app/datadict/internal/biz"
	"github.com/yrcs/yuneto/app/datadict/internal/data"
	"github.com/yrcs/yuneto/app/datadict/internal/pkg/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db       *gorm.DB
	mysqlDSN = "root:123456@tcp(192.168.23.111:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	tm biz.Transaction
	u  *biz.DictUsecase
)

func suiteSetUp(suiteName string) func() {
	log.Printf("\tsetUp fixture for suite %s\n", suiteName)
	dat, _, err := data.NewData(db, nil)
	if err != nil {
		log.Fatalf("Failed to new data, but got error %v", err)
	}
	r := data.NewDictRepo(dat, nil)
	tm = data.NewTransaction(dat)
	u = biz.NewDictUsecase(r, tm, nil)
	return func() {
		log.Printf("\ttearDown fixture for suite %s\n", suiteName)
	}
}

func TestFind(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	t.Run("testFindChildren", func(t *testing.T) {
		var (
			err   error
			id    = uint64(1)
			dicts []biz.E
		)
		if dicts, err = u.FindChildren(context.TODO(), id); err != nil {
			t.Errorf("Failed to find dicts, but got error %v", err)
		} else if len(dicts) == 0 {
			t.Errorf("Want %q, but got FindChildren(%v,%d) = empty slice", "[]biz.E", context.TODO(), id)
		}
	})
}

func pkgSetUp(pkgName string) func() {
	log.Printf("package SetUp fixture for %s\n", pkgName)
	var err error
	var sqlDB *sql.DB
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlDSN, // DSN (data source name)
		DefaultStringSize:         256,      // Default length of a string type field
		DisableDatetimePrecision:  true,     // Disable datetime precision, not supported before MySQL 5.6
		SkipInitializeWithVersion: false,    // Automatic configuration based on current MySQL version
	}), &gorm.Config{
		Logger:                                   gl.Default.LogMode(gl.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true, NoLowerCase: true},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
	})
	if err != nil {
		log.Fatalf("Failed to connect database, but got error %v", err)
	} else {
		sqlDB, err = db.DB()
		if err != nil {
			log.Fatalf("Failed to connect database, but got error %v", err)
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Fatalf("Failed to ping sqlDB, but got error %v", err)
		}

		runMigrations()
	}
	return func() {
		log.Printf("package TearDown fixture for %s\n", pkgName)
		if err = sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close the database, but got error %v", err)
		}
	}
}

func TestMain(m *testing.M) {
	defer pkgSetUp("package biz_test")()
	m.Run()
}

func runMigrations() {
	var err error
	allModels := []any{&po.Dict{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	// if err = db.Migrator().DropTable(allModels...); err != nil {
	// 	log.Fatalf("Failed to drop table, but got error %v\n", err)
	// }

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").AutoMigrate(allModels...); err != nil {
		log.Fatalf("Failed to auto migrate, but got error %v\n", err)
	}

	for _, m := range allModels {
		if !db.Migrator().HasTable(m) {
			log.Fatalf("Failed to create table for %#v\n", m)
		}
	}
}
