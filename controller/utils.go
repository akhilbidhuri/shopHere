package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/akhilbidhuri/shopHere/models"
	"gorm.io/gorm"
)

func parseJsonFromReq(req *http.Request, model interface{}) error {
	jsonData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, model); err != nil {
		return err
	}
	return nil
}

func validateToken(req *http.Request, db *gorm.DB) error {
	token := req.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		user := models.User{}
		if err := user.GetUserByToken(db, token); err != nil {
			return err
		}
	} else {
		return errors.New("authorization token not passed")
	}
	return nil
}
