package account

import "REVAMP-PHP-GO/internal/domain/ports"

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

func FindById(id string, p ports.PortRepo) interface{} {
	data, err := p.FindByID(id, &Account{}, table)
	if err != nil {
		panic(err)
	}
	return data
}

func List(p ports.PortRepo) interface{} {
	data, err := p.List(&listaccount, table)
	if err != nil {
		panic(err)
	}
	return data
}

func Create(mdl interface{}, p ports.PortRepo) error {
	col := "accountcode, parentaccountid, accountname, currencyid, isdebit, accounttype, isdisabled, requirecostcenter, allowallcostcenters"
	val := ":accountcode, :parentaccountid, :accountname, :currencyid, :isdebit, :accounttype, :isdisabled, :requirecostcenter, :allowallcostcenters"
	_, err := p.Create(mdl, table, col, val)
	if err != nil {
		panic(err)
	}
	return err
}

func Update(id string, mdl interface{}, p ports.PortRepo) error {
	val := "accountcode = :accountcode, parentaccountid = :parentaccountid, accountname = :accountname, currencyid = :currencyid, isdebit = :isdebit, accounttype = :accounttype, isdisabled = :isdisabled, requirecostcenter =:requirecostcenter, allowallcostcenters = :allowallcostcenters"
	_, err := p.Update(mdl, id, table, val)
	if err != nil {
		panic(err)
	}
	return err
}

func Delete(id string, p ports.PortRepo) error {
	_, err := p.Delete(id, table)
	if err != nil {
		panic(err)
	}
	return err
}
