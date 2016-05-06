package server

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/shakdwipeea/shadowfax/domain"
	"net/http"
	"strconv"
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
	router.POST("/business/tags", env.handleAddBusinessTags)
	router.POST("/business", env.handleAddBusiness)
	//todo see how to do pagination here
	router.GET("/business", env.handleGetBusiness)
	router.GET("/tags/:businessID", env.HandleGetTags)
}

//handleAddBusiness Route handler for adding business
func (e *Env) handleAddBusiness(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	decoder := json.NewDecoder(r.Body)

	var req domain.Business
	if err = decoder.Decode(&req); err != nil {
		SendErrorResponse(w, err.Error())
		return
	}

	if req.ID, err = domain.AddBusiness(e.Db, &req); err != nil {
		SendErrorResponse(w, err.Error())
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
		SendErrorResponse(w, err.Error())
		return
	}

	SendResponse(w, domain.GetAllBusinessHTTPResponse{
		Err: false,
		Msg: "Business retreived",
		Business: businessArr,
	})
}

//handleGetTags GET /tags/:businessId Get all tags of a business
func (e *Env) HandleGetTags(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	businessID, err := strconv.Atoi(p.ByName("businessID"))
	if err != nil {
		SendErrorResponse(w, err.Error())
	}

	bID := int64(businessID)

	bTags, err := domain.GetTagsOfBusiness(e.Db, &bID)
	if err != nil {
		SendErrorResponse(w, err.Error())
		return
	}

	SendResponse(w, domain.BusinessTagDetailsHTTPResponse{
		Err: false,
		Msg: "tags retreived",
		BusinessTags: *bTags,
	})
}

//handleAddBusinessTags POST /business/tags  Add tag to a business
func (e *Env) handleAddBusinessTags(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req domain.BusinessTags
	if err := ParseRequestBody(w, r, &req); err != nil {
		SendErrorResponse(w, err.Error())
		return
	}

	if err := domain.AddBusinessTag(e.Db, &req); err != nil {
		SendErrorResponse(w, err.Error())
		return
	}

	SendResponse(w, domain.BusinessTagHTTPResponse{
		Err: false,
		Msg: "Added tag",
		BusinessTags: req,
	})
}
