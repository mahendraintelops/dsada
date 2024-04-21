package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/daos/clients/sqls"
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/models"
	"github.com/mahendraintelops/dsada/daad/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type QeqeqeController struct {
	qeqeqeService *services.QeqeqeService
}

func NewQeqeqeController() (*QeqeqeController, error) {
	qeqeqeService, err := services.NewQeqeqeService()
	if err != nil {
		return nil, err
	}
	return &QeqeqeController{
		qeqeqeService: qeqeqeService,
	}, nil
}

func (qeqeqeController *QeqeqeController) CreateQeqeqe(context *gin.Context) {
	// validate input
	var input models.Qeqeqe
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// trigger qeqeqe creation
	qeqeqeCreated, err := qeqeqeController.qeqeqeService.CreateQeqeqe(&input)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, qeqeqeCreated)
}

func (qeqeqeController *QeqeqeController) ListQeqeqes(context *gin.Context) {
	// trigger all qeqeqes fetching
	qeqeqes, err := qeqeqeController.qeqeqeService.ListQeqeqes()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, qeqeqes)
}
