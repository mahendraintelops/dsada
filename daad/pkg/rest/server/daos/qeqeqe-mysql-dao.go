package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type QeqeqeDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateQeqeqes(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS qeqeqes(
		ID int NOT NULL AUTO_INCREMENT,
        
		Qdqdq VARCHAR(100) NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewQeqeqeDao() (*QeqeqeDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateQeqeqes(sqlClient)
	if err != nil {
		return nil, err
	}
	return &QeqeqeDao{
		sqlClient,
	}, nil
}

func (qeqeqeDao *QeqeqeDao) CreateQeqeqe(m *models.Qeqeqe) (*models.Qeqeqe, error) {
	insertQuery := "INSERT INTO qeqeqes(Qdqdq) values(?)"
	res, err := qeqeqeDao.sqlClient.DB.Exec(insertQuery, m.Qdqdq)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return nil, sqls.ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id
	log.Debugf("qeqeqe created")
	return m, nil
}

func (qeqeqeDao *QeqeqeDao) ListQeqeqes() ([]*models.Qeqeqe, error) {
	selectQuery := "SELECT * FROM qeqeqes"
	rows, err := qeqeqeDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var qeqeqes []*models.Qeqeqe
	for rows.Next() {
		m := models.Qeqeqe{}
		if err = rows.Scan(&m.Id, &m.Qdqdq); err != nil {
			return nil, err
		}
		qeqeqes = append(qeqeqes, &m)
	}
	if qeqeqes == nil {
		qeqeqes = []*models.Qeqeqe{}
	}
	log.Debugf("qeqeqe listed")
	return qeqeqes, nil
}
