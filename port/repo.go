package port

type PortRepo interface {
	Close()
	FindByID(id string, mdl interface{}, table, column string) (interface{}, error)
	List(mdl interface{}, table string) (interface{}, error)
	Create(data interface{}, table, column, value string) (int64, error)
	Update(data interface{}, id, table, column, columnValue string) (int64, error)
	Delete(id, table, column string) (int64, error)
	// Auth(uname string, mdl interface{}, table string) (interface{}, error)
}
