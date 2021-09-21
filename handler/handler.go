package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glavanan/mailgun/dbconnector"
	"github.com/glavanan/mailgun/service"
	log "github.com/sirupsen/logrus"
)

//Delivered a message for a domain
func Delivered(c *gin.Context) {
	db, err := dbconnector.GetDB()
	if err != nil {
		log.WithField("GetDB", "Couldn't get a connection").Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = service.Delivered(c.Param("domain"), db)
	if err != nil {
		log.WithField("Delivered", "exec event failed").Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(204, nil)
	return
}

//Bounced a message for a domain
func Bounced(c *gin.Context) {
	db, err := dbconnector.GetDB()
	if err != nil {
		log.WithField("GetDB", "Couldn't get a connection").Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = service.Bounced(c.Param("domain"), db)
	if err != nil {
		log.WithField("Bounced", "exec event failed").Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(204, nil)
	return
}

//GetCatchAll a message for a domain
func GetCatchAll(c *gin.Context) {
	db, err := dbconnector.GetDB()
	if err != nil {
		log.WithField("GetDB", "Couldn't get a connection").Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	status, err := service.GetCathAllStatus(c.Param("domain"), db)
	if err != nil {
		log.WithField("Bounced", "exec event failed").Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.String(200, status)
	return
}
