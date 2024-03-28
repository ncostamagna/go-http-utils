package response_test

import (
	"encoding/json"
	"github.com/ncostamagna/go_http_utils/response"
	"testing"
)

func TestSuccessResponse(t *testing.T) {

	type SuccessData struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Num  int    `json:"num"`
	}

	testMatrix := []struct {
		testName string
		res      response.Response
		body     string
		code     int
	}{
		{
			testName: "OK response",
			res:      response.OK("my message", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"my message","status":200,"data":{"id":"123","name":"my name","num":443}}`,
			code:     200,
		},
		{
			testName: "Created response",
			res:      response.Created("", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"Created request.","status":201,"data":{"id":"123","name":"my name","num":443}}`,
			code:     201,
		},
		{
			testName: "Accepted response",
			res:      response.Accepted("Accepted message", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"Accepted message","status":202,"data":{"id":"123","name":"my name","num":443}}`,
			code:     202,
		},
		{
			testName: "NonAuthoritativeInfo response",
			res:      response.NonAuthoritativeInfo("", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"Non-Authoritative Information request.","status":203,"data":{"id":"123","name":"my name","num":443}}`,
			code:     203,
		},
		{
			testName: "NoContent response",
			res:      response.NoContent("", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"No Content request.","status":204,"data":{"id":"123","name":"my name","num":443}}`,
			code:     204,
		},
		{
			testName: "ResetContent response",
			res:      response.ResetContent("my message", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"my message","status":205,"data":{"id":"123","name":"my name","num":443}}`,
			code:     205,
		},
		{
			testName: "PartialContent response",
			res:      response.PartialContent("", SuccessData{"123", "my name", 443}, nil),
			body:     `{"message":"Partial Content request.","status":206,"data":{"id":"123","name":"my name","num":443}}`,
			code:     206,
		},
	}

	for _, tt := range testMatrix {
		t.Run(tt.testName, func(t *testing.T) {
			v, err := json.Marshal(tt.res)
			if err != nil {
				t.Error(err)
			} else {
				if string(v) != tt.body {
					t.Errorf("got: %s - expect: %s", v, tt.body)
				}
			}

		})
	}

}

func TestErrorResponse(t *testing.T) {

	testMatrix := []struct {
		testName string
		res      response.Response
		body     string
		code     int
	}{
		{
			testName: "OK response",
			res:      response.BadRequest("my error"),
			body:     `{"status":400,"message":"my error"}`,
			code:     400,
		},
		{
			testName: "InternalServerError response",
			res:      response.InternalServerError(""),
			body:     `{"status":500,"message":"We encountered an error while processing your request."}`,
			code:     500,
		},
		{
			testName: "NotFound response",
			res:      response.NotFound(""),
			body:     `{"status":404,"message":"The requested resource was not found."}`,
			code:     404,
		},
		{
			testName: "Unauthorized response",
			res:      response.Unauthorized(""),
			body:     `{"status":401,"message":"You are not authenticated to perform the requested action."}`,
			code:     401,
		},
		{
			testName: "Forbidden response",
			res:      response.Forbidden(""),
			body:     `{"status":403,"message":"You are not authorized to perform the requested action."}`,
			code:     403,
		},
	}

	for _, tt := range testMatrix {
		t.Run(tt.testName, func(t *testing.T) {
			v, err := json.Marshal(tt.res)
			if err != nil {
				t.Error(err)
			} else {
				if string(v) != tt.body {
					t.Errorf("got: %s - expect: %s", v, tt.body)
				}
			}

		})
	}

}
