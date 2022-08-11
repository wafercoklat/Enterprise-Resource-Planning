package ports

type PortRepo interface {
	Close()
	FindByID(id string, mdl interface{}, tbl string) (interface{}, error)
	List(mdl interface{}, tbl string) (interface{}, error)
	Create(mdl interface{}, tbl, col, val string) (int64, error)
	Update(mdl interface{}, id, tbl, val string) (int64, error)
	Delete(id, tbl string) (int64, error)
	// Up() error
	// Drop() error
}
