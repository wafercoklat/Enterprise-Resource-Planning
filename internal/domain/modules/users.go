package modules

import "REVAMP-PHP-GO/internal/domain/ports"

type User struct {
	ID         int    `json:"id,omitempty"`
	UID        int    `json:"uid,omitempty"`
	Name       string `json:"name,omitempty"`
	Mail       string `json:"mail,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Location   string `json:"location,omitempty"`
	IsDisabled int    `json:"isdisabled,omitempty"`
	Verified   int    `json:"verified,omitempty"`
}

var tbl_users string = "users"
var listusers []User

func UserFindById(id string, p ports.PortRepo) (interface{}, error) {
	return p.FindByID(id, &Account{}, tbl_users)
}

func UserList(p ports.PortRepo) (interface{}, error) {
	return p.List(&listusers, tbl_users)
}

func UserCreate(mdl interface{}, p ports.PortRepo) (int64, error) {
	col := "uid, name, mail, username, password, location, isdisabled, verified"
	val := ":uid, :name, :mail, :username, :password, :location, :isdisabled, :verified"
	id, err := p.Create(mdl, tbl_users, col, val)
	return id, err
}

func UserUpdate(id string, mdl interface{}, p ports.PortRepo) (int64, error) {
	val := "accountcode = :accountcode, parentaccountid = :parentaccountid, accountname = :accountname, currencyid = :currencyid, isdebit = :isdebit, accounttype = :accounttype, isdisabled = :isdisabled, requirecostcenter =:requirecostcenter, allowallcostcenters = :allowallcostcenters"
	id_, err := p.Update(mdl, id, tbl_users, val)
	return id_, err
}

func UserDelete(id string, p ports.PortRepo) (int64, error) {
	id_, err := p.Delete(id, tbl_users)
	return id_, err
}
