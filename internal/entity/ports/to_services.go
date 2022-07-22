// It will be out to client (gate of service)
package ports

type portService interface {
	AddData(string) error
	CreateTransaction() error
	DeleteTransaction() error
	UpdateTransaction() error
	ViewAllTransaction() error
	VieyByIdTransaction() error
}
