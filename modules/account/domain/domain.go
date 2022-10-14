package account

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

var (
	Tbl_account    string = "ac_accounts"
	Acccolumn      string = "accountcode, parentaccountid, accountname, currencyid, isdebit, accounttype, isdisabled, requirecostcenter, allowallcostcenters"
	Accvalue       string = ":accountcode, :parentaccountid, :accountname, :currencyid, :isdebit, :accounttype, :isdisabled, :requirecostcenter, :allowallcostcenters"
	Acccolumnvalue string = "accountcode = :accountcode, parentaccountid = :parentaccountid, accountname = :accountname, currencyid = :currencyid, isdebit = :isdebit, accounttype = :accounttype, isdisabled = :isdisabled, requirecostcenter =:requirecostcenter, allowallcostcenters = :allowallcostcenters"
	Acclist        []Account
)

type Services interface {
	FindByID(id string) *Account
	List() *[]Account
	Create(mdl Account) error
	Update(id string, mdl Account) error
	Delete(id string) error
}
