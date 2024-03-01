package Application

import (
	"github.com/bykovme/gotrans"
)

// Ok Response
func (req Request) Ok(body interface{}) {

	req.Response(200, buildResponse(body, gotrans.T("ok"), 200, nil))
}

// Created Response
func (req Request) Created(body interface{}) {
	req.Response(201, buildResponse(body, gotrans.T("created"), 201, nil))
}

// NotAuth Response
func (req Request) NotAuth() {
	req.Response(401, buildResponse(nil, gotrans.T("you_are_not_auth"), 401, nil))
}

func (req Request) BadRequest(err interface{}) {
	req.Response(422, buildResponse(nil, gotrans.T("something_wrong"), 422, err))
}

func (req Request) UserNotFound() {
	req.Response(404, buildResponse(nil, gotrans.T("we_not_found_this_user_in_our_system"), 404, nil))
}

func buildResponse(payload interface{}, message string, code int, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["payload"] = payload
	response["message"] = message
	response["code"] = code
	response["errors"] = err
	return response
}
