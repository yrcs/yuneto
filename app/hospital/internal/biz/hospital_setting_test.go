package biz_test

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/app/hospital/internal/data"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/do"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/po"
	"github.com/yrcs/yuneto/pkg/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db       *gorm.DB
	mysqlDSN = "root:123456@tcp(192.168.23.111:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	logger klog.Logger
	tm     biz.Transaction
	hsu    *biz.HospitalSettingUsecase
)

func suiteSetUp(suiteName string) func() {
	log.Printf("\tsetUp fixture for suite %s\n", suiteName)
	logger = klog.With(klog.NewStdLogger(os.Stdout),
		"service.name", "Hospital",
		"service.version", "test",
		"ts", klog.DefaultTimestamp,
		"caller", klog.DefaultCaller,
	)
	dat, _, err := data.NewData(db, logger)
	if err != nil {
		log.Fatalf("Failed to new data, but got error %v", err)
	}
	hsr := data.NewHospitalSettingRepo(dat, logger)
	tm = data.NewTransaction(dat)
	hsu = biz.NewHospitalSettingUsecase(hsr, tm, logger)
	return func() {
		log.Printf("\ttearDown fixture for suite %s\n", suiteName)
	}
}

func TestAdd(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	hs := do.HospitalSetting{
		Id:                 1553670574587252736,
		Name:               "中山大学附属第一医院",
		RegistrationNumber: "350002440102150131",
		ContactPerson:      "肖海鹏",
		ContactMobile:      "16602000001",
		ApiUrl:             "http://127.0.0.1:8000",
	}
	t.Run("testAdd", func(t *testing.T) {
		if _, err := hsu.Add(context.TODO(), &hs); err != nil {
			t.Errorf("Failed to add hospital setting, but got error %v", err)
		} else if hs.CreatedAt.IsZero() {
			t.Fatalf("hs' created at should not zero, %v", hs.CreatedAt)
		}
	})
}

func TestExists(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	t.Run("testNameExists", func(t *testing.T) {
		name := "中山大学附属第一医院"
		if exists, err := hsu.NameExists(context.TODO(), name); err != nil {
			t.Errorf("Failed to check whether name exists, but got error %v", err)
		} else if exists {
			t.Errorf("Want %t, but got NameExists(%v,%q) = %t", false, context.TODO(), name, exists)
		}
	})
	t.Run("testRegistrationNumberExists", func(t *testing.T) {
		registrationNumber := "350002440102150131"
		if exists, err := hsu.RegistrationNumberExists(context.TODO(), registrationNumber); err != nil {
			t.Errorf("Failed to check whether registration number exists, but got error %v", err)
		} else if exists {
			t.Errorf("Want %t, but got RegistrationNumberExists(%v,%q) = %t", false, context.TODO(), registrationNumber, exists)
		}
	})
}

func TestUnique(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	id := uint64(1553670574587252736)
	t.Run("testNameUnique", func(t *testing.T) {
		name := "中山大学附属第一医院"
		if unique, err := hsu.NameUnique(context.TODO(), id, name); err != nil {
			t.Errorf("Failed to check whether name is unique, but got error %v", err)
		} else if !unique {
			t.Errorf("Want %t, but got NameUnique(%v,%d,%q) = %t", true, context.TODO(), id, name, unique)
		}
	})
	t.Run("testRegistrationNumberUnique", func(t *testing.T) {
		registrationNumber := "350002440102150131"
		if unique, err := hsu.RegistrationNumberUnique(context.TODO(), id, registrationNumber); err != nil {
			t.Errorf("Failed to check whether registration number is unique, but got error %v", err)
		} else if !unique {
			t.Errorf("Want %t, but got RegistrationNumberUnique(%v,%d,%q) = %t", true, context.TODO(), id, registrationNumber, unique)
		}
	})
}

func TestFind(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	t.Run("testFind", func(t *testing.T) {
		var (
			err error
			id  = uint64(1553670574587252736)
			hs  *do.HospitalSetting
		)
		tm.InTx(context.TODO(), func(ctx context.Context) error {
			if hs, err = hsu.Find(context.TODO(), id, repo.UPDATE); err != nil {
				return err
			} else if hs == nil {
				t.Errorf("Want %q instance, but got Find(%v,%d) = nil", "*HospitalSetting", context.TODO(), id)
			}
			return nil
		})
		if err != nil {
			t.Errorf("Failed to find hospital setting, but got error %v", err)
		}
	})
	t.Run("testFindByField", func(t *testing.T) {
		var (
			err        error
			fieldName  = "RegistrationNumber"
			fieldValue = "350002440102150131"
			hs         *do.HospitalSetting
		)
		tm.InTx(context.TODO(), func(ctx context.Context) error {
			if hs, err = hsu.FindByField(context.TODO(), fieldName, fieldValue, true, repo.SHARE); err != nil {
				return err
			} else if hs == nil {
				t.Errorf("Want %q instance, but got FindByField(%v,%s,%v,%s) = nil", "*HospitalSetting", context.TODO(), fieldName, fieldValue, repo.SHARE)
			}
			return nil
		})
		if err != nil {
			t.Errorf("Failed to find hospital setting, but got error %v", err)
		}
	})
}

func TestEdit(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	m := make(map[string]any, 2)
	m["Id"] = 1553670574587252736
	m["ContactMobile"] = "16602000088"
	t.Run("testEdit", func(t *testing.T) {
		if do, err := hsu.Edit(context.TODO(), m); err != nil {
			t.Errorf("Failed to edit hospital setting, but got error %v", err)
		} else if do.UpdatedAt.IsZero() {
			t.Fatalf("do's updated at should not zero, %v", do.UpdatedAt)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Cleanup(suiteSetUp(t.Name()))
	id := uint64(1553670574587252736)
	t.Run("testDelete", func(t *testing.T) {
		if err := hsu.Delete(context.TODO(), id); err != nil {
			t.Errorf("Failed to delete hospital setting, but got error %v", err)
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
	allModels := []any{&po.HospitalSetting{}}
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
