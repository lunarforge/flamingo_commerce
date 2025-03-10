// +build integration

package graphql_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/lunarforge/flamingo_commerce/test/integrationtest"
	"github.com/lunarforge/flamingo_commerce/test/integrationtest/projecttest/helper"
)

func Test_CommerceCustomerStatus(t *testing.T) {
	t.Parallel()
	baseURL := "http://" + FlamingoURL
	type expected struct {
		status bool
		userID string
	}
	tests := []struct {
		name         string
		performLogin bool
		expected     expected
	}{
		{
			name:         "not logged in",
			performLogin: false,
			expected: expected{
				status: false,
				userID: "",
			},
		},
		{
			name:         "logged in",
			performLogin: true,
			expected: expected{
				status: true,
				userID: "username",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := integrationtest.NewHTTPExpect(t, baseURL)
			if tt.performLogin {
				loginTestCustomer(t, e)
			}

			resp := helper.GraphQlRequest(t, e, loadGraphQL(t, "customer_status", nil)).Expect()
			resp.Status(http.StatusOK)
			getValue(resp, "Commerce_Customer_Status", "isLoggedIn").Boolean().Equal(tt.expected.status)
			getValue(resp, "Commerce_Customer_Status", "userID").String().Equal(tt.expected.userID)
		})
	}
}

func Test_CommerceCustomer(t *testing.T) {
	t.Parallel()
	baseURL := "http://" + FlamingoURL

	tests := []struct {
		name         string
		performLogin bool
		validator    func(t *testing.T, response *httpexpect.Response)
	}{
		{
			name:         "not logged in",
			performLogin: false,
			validator: func(t *testing.T, response *httpexpect.Response) {
				response.JSON().Object().Value("data").Object().Value("Commerce_Customer").Null()
			},
		},
		{
			name:         "logged in",
			performLogin: true,
			validator: func(t *testing.T, response *httpexpect.Response) {
				getValue(response, "Commerce_Customer", "id").Equal("username")
				personalData := getValue(response, "Commerce_Customer", "personalData").Object()
				personalData.Value("firstName").Equal("Flamingo")
				personalData.Value("lastName").Equal("Commerce")
				personalData.Value("birthday").Equal("2019-04-02")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := integrationtest.NewHTTPExpect(t, baseURL)
			if tt.performLogin {
				loginTestCustomer(t, e)
			}

			resp := helper.GraphQlRequest(t, e, loadGraphQL(t, "customer", nil)).Expect()
			resp.Status(http.StatusOK)
			tt.validator(t, resp)
		})
	}
}
