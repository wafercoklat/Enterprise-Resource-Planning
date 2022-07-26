// Business logic for Account Module
package models

type Account struct {
	Id                  int    `json:"id,omitempty"`
	AccountCode         string `json:"account_code,omitempty"`
	ParentAccountID     int    `json:"parent_account_id,omitempty"`
	AccountName         string `json:"account_name,omitempty"`
	CurrencyID          string `json:"currency_id,omitempty"`
	IsDebit             int    `json:"is_debit,omitempty"`
	AccountType         int    `json:"account_type,omitempty"`
	LevelAcc            int    `json:"level_acc,omitempty"`
	IsDisabled          int    `json:"is_disabled,omitempty"`
	RequireCostCenter   int    `json:"require_cost_center,omitempty"`
	AllowAllCostCenters int    `json:"allow_all_cost_centers,omitempty"`
}
