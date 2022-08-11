package ports

type PortServices interface {
	AccountFindByID(id string) *interface{}
	AccountList() *interface{}
	AccountCreate(mdl interface{}) error
	AccountUpdate(id string, mdl interface{}) error
	AccountDelete(id string) error
}
