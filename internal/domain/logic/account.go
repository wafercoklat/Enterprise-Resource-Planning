// Business logic for Account Module
package domain

type Account struct {
	Id                  int    `JSON:"_id,omitempty" bson:"_id"`
	AccountCode         string `JSON:"account_code,omitempty"  bson:"account_code"`
	ParentAccountID     int    `JSON:"parent_account_id,omitempty"  bson:"parent_account_id"`
	AccountName         string `JSON:"account_name,omitempty"  bson:"account_name"`
	CurrencyID          string `JSON:"currency_id,omitempty"  bson:"currency_id"`
	IsDebit             int    `JSON:"is_debit,omitempty"  bson:"is_debit"`
	LevelAcc            int    `JSON:"level_acc,omitempty"  bson:"level_acc"`
	IsDisabled          int    `JSON:"is_disabled,omitempty"  bson:"is_disabled"`
	RequireCostCenter   int    `JSON:"require_cost_center,omitempty"  bson:"require_cost_center"`
	AllowAllCostCenters int    `JSON:"allow_all_cost_centers,omitempty"  bson:"allow_all_cost_centers"`
}

var Collection string

// Here is will be CRUD function
