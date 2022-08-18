package modules_test

import (
	datasources "REVAMP-PHP-GO/internal/datasources/sqlDB"
	account "REVAMP-PHP-GO/internal/domain/modules"
	"REVAMP-PHP-GO/internal/domain/ports"
	"fmt"
	"os"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	datasource ports.PortRepo
	err        error
	temp_id    int64
	mydata     = account.Account{
		AccountCode:         "Testing",
		AccountName:         "Testing",
		CurrencyID:          "",
		IsDebit:             1,
		AccountType:         1,
		IsDisabled:          0,
		RequireCostCenter:   0,
		AllowAllCostCenters: 1,
	}
	mydata_fail = account.Account{
		AccountCode:         "Testing",
		AccountName:         "Hello World",
		CurrencyID:          "MYR",
		IsDebit:             0,
		AccountType:         0,
		IsDisabled:          1,
		RequireCostCenter:   1,
		AllowAllCostCenters: 0,
	}
	mydata_update = account.Account{
		AccountCode:         "Update Testing",
		AccountName:         "Update Testing",
		CurrencyID:          "IDR",
		IsDebit:             0,
		AccountType:         0,
		IsDisabled:          1,
		RequireCostCenter:   1,
		AllowAllCostCenters: 0,
	}

	mydata_update_fail = account.Account{
		AccountCode: "1110",
	}
)

func TestMain(m *testing.M) {
	dialect := "mysql"
	dsn := "root:@tcp(localhost:3306)/enterprise022?parseTime=true"
	datasource, err = datasources.New(dialect, dsn, 10, 10)
	if err != nil {
		fmt.Printf("%s", err)
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestAccount_Create_Success(t *testing.T) {
	info := "Test Success Create Account"

	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			// Create data for testing
			id, err := account.Create(mydata, datasource)

			// Must be nil
			So(err, ShouldBeNil)
			So(strconv.Itoa(int(id)), ShouldNotBeBlank)

			temp_id = id
			fmt.Println(temp_id)
		})
	})
}

func TestAccount_Create_Fail(t *testing.T) {
	info := "Test Fail Create Account"

	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			// Create data for testing
			id, err := account.Create(mydata_fail, datasource)

			// Must be nil
			So(err, ShouldBeError)
			So(int(id), ShouldBeZeroValue)
		})
	})
}

func TestAccount_FindByID_Success(t *testing.T) {
	info := "Test Success Find By ID"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			acc, _ := account.FindById(strconv.Itoa(int(temp_id)), datasource)
			_account, ok := acc.(*account.Account)

			// check if data is exist
			So(ok, ShouldBeTrue)

			// Check if sample is equal from output data
			So(_account.Id, ShouldEqual, int(temp_id))
			So(_account.AccountCode, ShouldEqual, mydata.AccountCode)
			So(_account.AccountName, ShouldEqual, mydata.AccountName)
		})

	})
}

func TestAccount_FindByID_Fail(t *testing.T) {
	info := "Test Fail Find By ID"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			acc, _ := account.FindById(strconv.Itoa(int(temp_id)), datasource)
			_account := acc.(*account.Account)

			// Check if sample is not equal from output data
			So(_account.Id, ShouldNotEqual, "20")
			So(_account.CurrencyID, ShouldNotEqual, mydata_fail.CurrencyID)
			So(_account.AccountName, ShouldNotEqual, mydata_fail.AccountName)
		})

	})
}

func TestAccount_List_Success(t *testing.T) {
	info := "Test Success List Account"

	data := map[string]string{
		"id":                  strconv.Itoa(int(temp_id)),
		"accountcode":         mydata.AccountCode,
		"parentaccountid":     strconv.Itoa(mydata.ParentAccountID),
		"accountname":         mydata.AccountName,
		"currencyid":          mydata.CurrencyID,
		"isdebit":             strconv.Itoa(mydata.IsDebit),
		"accounttype":         strconv.Itoa(mydata.AccountType),
		"isdisabled":          strconv.Itoa(mydata.IsDisabled),
		"requirecostcenter":   strconv.Itoa(mydata.RequireCostCenter),
		"allowallcostcenters": strconv.Itoa(mydata.AllowAllCostCenters),
	}
	var myArr []map[string]string

	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			acc, _ := account.List(datasource)
			_account := acc.(*[]account.Account)

			for _, val := range *_account {
				temp := make(map[string]string)
				temp["id"] = fmt.Sprintf("%v", val.Id)
				temp["accountcode"] = fmt.Sprintf("%v", val.AccountCode)
				temp["parentaccountid"] = fmt.Sprintf("%v", val.ParentAccountID)
				temp["accountname"] = fmt.Sprintf("%v", val.AccountName)
				temp["currencyid"] = fmt.Sprintf("%v", val.CurrencyID)
				temp["isdebit"] = fmt.Sprintf("%v", val.IsDebit)
				temp["accounttype"] = fmt.Sprintf("%v", val.AccountType)
				temp["isdisabled"] = fmt.Sprintf("%v", val.IsDisabled)
				temp["requirecostcenter"] = fmt.Sprintf("%v", val.RequireCostCenter)
				temp["allowallcostcenters"] = fmt.Sprintf("%v", val.AllowAllCostCenters)

				myArr = append(myArr, temp)
			}

			// Check list is exist
			So(myArr, ShouldNotBeEmpty)

			// Check data is exist in list
			So(data, ShouldBeIn, myArr)
			So(myArr, ShouldContain, data)

			// Check key if exist in list
			So(myArr[0], ShouldContainKey, "id")
		})
	})
}

func TestAccount_List_Fail(t *testing.T) {
	info := "Test Fail List Account"

	data := map[string]string{
		"id":                  "x",
		"accountcode":         "15+23",
		"accountname":         "Testing",
		"currencyid":          "",
		"isdebit":             "1",
		"accounttype":         "1",
		"isdisabled":          "0",
		"requirecostcenter":   "0",
		"allowallcostcenters": "1",
	}
	var myArr []map[string]string

	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			acc, _ := account.List(datasource)
			_account := acc.(*[]account.Account)

			for _, val := range *_account {
				temp := make(map[string]string)
				temp["id"] = fmt.Sprintf("%v", val.Id)
				temp["accountcode"] = fmt.Sprintf("%v", val.AccountCode)
				temp["parentaccountid"] = fmt.Sprintf("%v", val.ParentAccountID)
				temp["accountname"] = fmt.Sprintf("%v", val.AccountName)
				temp["currencyid"] = fmt.Sprintf("%v", val.CurrencyID)
				temp["isdebit"] = fmt.Sprintf("%v", val.IsDebit)
				temp["accounttype"] = fmt.Sprintf("%v", val.AccountType)
				temp["isdisabled"] = fmt.Sprintf("%v", val.IsDisabled)
				temp["requirecostcenter"] = fmt.Sprintf("%v", val.RequireCostCenter)
				temp["allowallcostcenters"] = fmt.Sprintf("%v", val.AllowAllCostCenters)

				myArr = append(myArr, temp)
			}

			// Check data is exist in list
			So(data, ShouldNotBeIn, myArr)
			So(myArr, ShouldNotContain, data)

			// Check key if exist in list
			So(myArr[0], ShouldNotContainKey, "test")
		})
	})
}

func TestAccount_Edit_Success(t *testing.T) {
	info := "Test Success Edit Account"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			RowsAffected, err := account.Update(strconv.Itoa(int(temp_id)), mydata_update, datasource)

			// Must be nil
			So(err, ShouldBeNil)
			So(strconv.Itoa(int(RowsAffected)), ShouldNotBeBlank)

			acc, err := account.FindById(strconv.Itoa(int(temp_id)), datasource)
			data_check, ok := acc.(*account.Account)

			// check if data is exist
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)

			// Check if sample is equal from output data
			So(data_check.Id, ShouldEqual, int(temp_id))
			So(data_check.AccountCode, ShouldEqual, mydata_update.AccountCode)
			So(data_check.AccountName, ShouldEqual, mydata_update.AccountName)

		})
	})
}

func TestAccount_Edit_Fail(t *testing.T) {
	info := "Test Fail Edit Account"

	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			RowsAffected, err := account.Update(strconv.Itoa(int(temp_id)), mydata_update_fail, datasource)

			// Must be nil
			So(err, ShouldBeError)
			So(RowsAffected, ShouldBeZeroValue)

			acc, err := account.FindById(strconv.Itoa(int(temp_id)), datasource)
			data_check, ok := acc.(*account.Account)

			// check if data is exist
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)

			// Check if sample is equal from output data
			So(data_check.Id, ShouldEqual, int(temp_id))
			So(data_check.AccountCode, ShouldEqual, mydata_update.AccountCode)
			So(data_check.AccountName, ShouldEqual, mydata_update.AccountName)
		})
	})
}

func TestAccount_Delete_Success(t *testing.T) {
	info := "Test Success Delete Account"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			RowsAffected, err := account.Delete(strconv.Itoa(int(temp_id)), datasource)

			// Must be nil
			So(err, ShouldBeNil)
			So(strconv.Itoa(int(RowsAffected)), ShouldNotBeBlank)
		})
	})
}

func TestAccount_Delete_Fail(t *testing.T) {
	info := "Test Fail Delete Account"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			RowsAffected, _ := account.Delete(strconv.Itoa(int(temp_id)), datasource)

			// Must be nil
			So(int(RowsAffected), ShouldBeZeroValue)
		})
	})
}

// func TestArray(t *testing.T) {
// 	t.Run("Test Success List", func(t *testing.T) {
// 		Convey("Test Success List", t, func() {
// 			// var chickens = []map[string]string{
// 			// 	{"name": "chicken blue", "gender": "male"},
// 			// 	{"name": "chicken red", "gender": "male"},
// 			// 	{"name": "chicken yellow", "gender": "female"},
// 			// }

// 			var chicken2 = map[string]string{
// 				"name": "chicken red", "gender": "male",
// 			}

// 			// So(chicken2, ShouldBeIn, chickens)
// 			So(chicken2, ShouldContainKey, "name")
// 		})
// 	})
// }

// func TestAccount_List_Fail(t *testing.T) {
// 	t.Run("Test Find By ID", func(t *testing.T) {
// 		Convey("Test Find By ID", t, func() {
// 			acc := account.List(datasource)
// 			fmt.Println(acc)
// 			_account := acc.(*account.Account)
// 			So(_account.AccountCode, ShouldNotEqual, "20")
// 			So(_account.AccountName, ShouldNotEqual, "Test")
// 		})

// 	})
// }
