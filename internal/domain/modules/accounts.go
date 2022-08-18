package modules

import (
	"REVAMP-PHP-GO/internal/domain/ports"
)

var table string = "accounts"

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

var listaccount []Account

func FindById(id string, p ports.PortRepo) (interface{}, error) {
	return p.FindByID(id, &Account{}, table)
}

func List(p ports.PortRepo) (interface{}, error) {
	return p.List(&listaccount, table)
}

func Create(mdl interface{}, p ports.PortRepo) (int64, error) {
	col := "accountcode, parentaccountid, accountname, currencyid, isdebit, accounttype, isdisabled, requirecostcenter, allowallcostcenters"
	val := ":accountcode, :parentaccountid, :accountname, :currencyid, :isdebit, :accounttype, :isdisabled, :requirecostcenter, :allowallcostcenters"
	id, err := p.Create(mdl, table, col, val)
	return id, err
}

func Update(id string, mdl interface{}, p ports.PortRepo) (int64, error) {
	val := "accountcode = :accountcode, parentaccountid = :parentaccountid, accountname = :accountname, currencyid = :currencyid, isdebit = :isdebit, accounttype = :accounttype, isdisabled = :isdisabled, requirecostcenter =:requirecostcenter, allowallcostcenters = :allowallcostcenters"
	id_, err := p.Update(mdl, id, table, val)
	return id_, err
}

func Delete(id string, p ports.PortRepo) (int64, error) {
	id_, err := p.Delete(id, table)
	return id_, err
}
