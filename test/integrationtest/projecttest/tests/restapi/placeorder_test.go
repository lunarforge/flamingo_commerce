// +build integration

package restapi_test

import (
	"testing"

	"github.com/lunarforge/flamingo_commerce/checkout/application/placeorder"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/test/integrationtest"
)

func Test_Checkout_SimplePlaceOrderProcess(t *testing.T) {

	e := integrationtest.NewHTTPExpect(t, "http://"+FlamingoURL)
	// add something to the cart
	response := e.POST("/api/v1/cart/delivery/delivery/item").WithQuery("deliveryCode", "delivery").WithQuery("marketplaceCode", "fake_simple").Expect()
	response.Status(200).JSON().Object().Value("Success").Boolean().Equal(true)

	// add billing
	response = e.PUT("/api/v1/cart/billing").WithFormField("firstname", "Max").WithFormField("lastname", "Mustermann").WithFormField("email", "test@test.de").Expect()
	response.Status(200).JSON().Object().Value("Success").Boolean().Equal(true)

	// add shipping
	response = e.PUT("/api/v1/cart/delivery/delivery/").WithFormField("deliveryAddress.firstname", "Max").WithFormField("deliveryAddress.lastname", "Mustermann").WithFormField("deliveryAddress.email", "test@test.de").Expect()
	response.Status(200).JSON().Object().Value("Success").Boolean().Equal(true)

	// add payment selection
	response = e.PUT("/api/v1/cart/payment-selection").WithQuery("gateway", "fake_payment_gateway").WithQuery("method", "payment_waiting_for_customer").Expect()
	response.Status(200).JSON().Object().Value("Success").Boolean().Equal(true)

	// start place order
	response = e.PUT("/api/v1/checkout/placeorder").WithQuery("returnURL", "http://www.example.org").Expect()
	response.Status(201).JSON().Object().Value("UUID").String().NotEmpty()
	uuid := response.Status(201).JSON().Object().Value("UUID").String().Raw()

	// get last place order context
	response = e.GET("/api/v1/checkout/placeorder").WithQuery("returnURL", "http://www.example.org").Expect()
	response.Status(200).JSON().Object().Value("UUID").String().Equal(uuid)

	// cancel place order
	response = e.POST("/api/v1/checkout/placeorder/cancel").Expect()
	response.Status(200).Body().Equal("true\n")

	// get last place order context
	response = e.GET("/api/v1/checkout/placeorder").WithQuery("returnURL", "http://www.example.org").Expect()
	response.Status(200).JSON().Object().Value("State").String().Equal("Failed")

	// clear last place order context
	response = e.DELETE("/api/v1/checkout/placeorder").WithQuery("returnURL", "http://www.example.org").Expect()
	response.Status(200).Body().Equal("true\n")

	// get last place order context
	response = e.GET("/api/v1/checkout/placeorder").WithQuery("returnURL", "http://www.example.org").Expect()
	response.Status(500).JSON().Object().Value("Message").String().Equal(placeorder.ErrNoPlaceOrderProcess.Error())

	// set payment selection to a working one
	response = e.PUT("/api/v1/cart/payment-selection").WithQuery("gateway", "fake_payment_gateway").WithQuery("method", "payment_completed").Expect()
	response.Status(200).JSON().Object().Value("Success").Boolean().Equal(true)

	// start place order again
	response = e.PUT("/api/v1/checkout/placeorder").WithQuery("returnURL", "http://www.example.org").Expect()
	response.Status(201).JSON().Object().Value("UUID").String().NotEmpty()

	// refresh place order
	response = e.POST("/api/v1/checkout/placeorder/refresh").Expect()
	response.Status(200).JSON().Object().Value("State").String().NotEmpty()
	response = e.POST("/api/v1/checkout/placeorder/refresh-blocking").Expect()
	response.Status(200).JSON().Object().Value("State").String().NotEmpty()

	// get last place order context
	response = e.GET("/api/v1/checkout/placeorder").Expect()
	response.Status(200).JSON().Object().Value("FailedReason").String().Equal("")
	response.Status(200).JSON().Object().Value("State").String().Equal(states.Success{}.Name())
}
