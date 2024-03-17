package response_test

import (
	"fmt"
	"github.com/ncostamagna/go_lib_response/response"
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
		res      *response.Response[SuccessData, response.SuccessResponse[SuccessData]]
		emptyObj bool
		data     SuccessData
		message  string
		code     int
	}{
		{
			testName: "OK response",
			res:      response.OK("my message", SuccessData{"123", "my name", 443}, nil),
			emptyObj: false,
			data:     SuccessData{"123", "my name", 443},
			message:  "my message",
			code:     200,
		},
	}

	for _, tt := range testMatrix {
		t.Run(tt.testName, func(t *testing.T) {

			if tt.res.Body.Message != tt.message {
				t.Errorf("message '%s' isn't '%s'", tt.res.Body.Message, tt.message)
			}

			if tt.res.Body.Status != tt.code {
				t.Errorf("code '%d' isn't '%d'", tt.res.Body.Status, tt.code)
			}

			if tt.emptyObj {
				if tt.res.Body.Data != nil {
					t.Error("data must be nil")
				}
			} else {
				if tt.res.Body.Data.ID != tt.data.ID || tt.res.Body.Data.Name != tt.data.Name || tt.res.Body.Data.Num != tt.data.Num {
					t.Errorf("object: '%v' isn't '%v'", tt.res.Body.Data, tt.data)
				}
			}

		})
	}

}
