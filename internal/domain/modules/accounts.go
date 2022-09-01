package modules

import (
	"REVAMP-PHP-GO/internal/domain/ports"
)

var tbl_account string = "accounts"
var listaccount []Account

type Account struct {
	Id                  int    `json:"id,omitempty"`
	AccountCode         string `json:"accountcode,omitempty"`
	ParentAccountID     int    `json:"parentaccountid,omitempty"`
	AccountName         string `json:"accountname,omitempty"`
	CurrencyID          string `json:"currencyid,omitempty"`
	IsDebit             int    `json:"isdebit,omitempty"`
	AccountType         int    `json:"accounttype,omitempty"`
	IsDisabled          int    `json:"isdisabled,omitempty"`
	RequireCostCenter   int    `json:"requirecostcenter,omitempty"`
	AllowAllCostCenters int    `json:"allowallcostcenters,omitempty"`
}

func FindById(id string, p ports.PortRepo) (interface{}, error) {
	return p.FindByID(id, &Account{}, tbl_account)
}

func List(p ports.PortRepo) (interface{}, error) {
	return p.List(&listaccount, tbl_account)
}

func Create(mdl interface{}, p ports.PortRepo) (int64, error) {
	col := "accountcode, parentaccountid, accountname, currencyid, isdebit, accounttype, isdisabled, requirecostcenter, allowallcostcenters"
	val := ":accountcode, :parentaccountid, :accountname, :currencyid, :isdebit, :accounttype, :isdisabled, :requirecostcenter, :allowallcostcenters"
	id, err := p.Create(mdl, tbl_account, col, val)
	return id, err
}

func Update(id string, mdl interface{}, p ports.PortRepo) (int64, error) {
	val := "accountcode = :accountcode, parentaccountid = :parentaccountid, accountname = :accountname, currencyid = :currencyid, isdebit = :isdebit, accounttype = :accounttype, isdisabled = :isdisabled, requirecostcenter =:requirecostcenter, allowallcostcenters = :allowallcostcenters"
	id_, err := p.Update(mdl, id, tbl_account, val)
	return id_, err
}

func Delete(id string, p ports.PortRepo) (int64, error) {
	id_, err := p.Delete(id, tbl_account)
	return id_, err
}
