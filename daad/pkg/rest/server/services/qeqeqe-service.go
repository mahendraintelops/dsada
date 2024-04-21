package services

import (
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/daos"
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/models"
)

type QeqeqeService struct {
	qeqeqeDao *daos.QeqeqeDao
}

func NewQeqeqeService() (*QeqeqeService, error) {
	qeqeqeDao, err := daos.NewQeqeqeDao()
	if err != nil {
		return nil, err
	}
	return &QeqeqeService{
		qeqeqeDao: qeqeqeDao,
	}, nil
}

func (qeqeqeService *QeqeqeService) CreateQeqeqe(qeqeqe *models.Qeqeqe) (*models.Qeqeqe, error) {
	return qeqeqeService.qeqeqeDao.CreateQeqeqe(qeqeqe)
}

func (qeqeqeService *QeqeqeService) ListQeqeqes() ([]*models.Qeqeqe, error) {
	return qeqeqeService.qeqeqeDao.ListQeqeqes()
}
