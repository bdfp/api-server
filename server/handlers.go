package server

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/shakdwipeea/shadowfax/domain"
	"net/http"
)

//HTTPResponse general response
type HTTPResponse struct {
	err     bool
	message string
}

// Env Functionalists required by the entire Application
type Env struct {
	Db *sql.DB
}

// RegisterHandlers adds the route handlers for various calls
func RegisterHandlers(router *httprouter.Router, env Env) {
	router.POST("/business", env.handleAddBusiness)
	router.GET("/business", env.handleGetBusiness)
}

//handleAddBusiness Route handler for adding business
func (e *Env) handleAddBusiness(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	decoder := json.NewDecoder(r.Body)

	var req domain.Business
	if err = decoder.Decode(&req); err != nil {
		SendResponse(w, HTTPErrorResponse{
			Err:    true,
			Reason: err.Error(),
		})
		return
	}

	if req.ID, err = domain.AddBusiness(e.Db, req); err != nil {
		SendResponse(w, HTTPErrorResponse{
			Err:    true,
			Reason: err.Error(),
		})
		return
	}

	SendResponse(w, domain.BusinessHTTPResponse{
		Err:      false,
		Msg:      "Business Added",
		Business: req,
	})

}
 
//handleGetBusiness GET /business Api handler to get all Business 
func (e *Env) handleGetBusiness(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	businessArr, err := domain.GetAllBusiness(e.Db)
	if err != nil {
		SendResponse(w, HTTPErrorResponse{
			Err:    true,
			Reason: err.Error(),
		})
		return
	}

	SendResponse(w, domain.GetAllBusinessHTTPResponse{
		Err: false,
		Msg: "Business retreived",
		Business: businessArr,
	})
}