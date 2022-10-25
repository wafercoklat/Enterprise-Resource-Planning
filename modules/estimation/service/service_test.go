package estimation_test

import (
	db "STACK-ERP/database"
	domain "STACK-ERP/modules/estimation/domain"
	"STACK-ERP/port"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	database_ port.PortRepo
	ID        int    = 1
	column    string = "EstimationID"
)

func TestMain(m *testing.M) {
	// .Env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Database
	DSN := os.Getenv("ADM_USER") + ":" + os.Getenv("ADM_PASS") + "@tcp(" + os.Getenv("IP_WHITELIST") + ":" + os.Getenv("PORT") + ")/" + os.Getenv("DATABASE") + "?parseTime=true"
	database_, err = db.New(os.Getenv("DIALECT"), DSN, 10, 10)

	if err != nil {
		fmt.Printf("%s", err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestFindByID_Success(t *testing.T) {
	info := "Test Success FindByID Method"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			_header, err := database_.FindByID(strconv.Itoa(ID), &domain.EstimationHeader{}, domain.Tbl_estimation["header"], column)
			So(err, ShouldBeNil)
			So(_header.(*domain.EstimationHeader).EstimationID, ShouldEqual, ID)
		})
	})
}
