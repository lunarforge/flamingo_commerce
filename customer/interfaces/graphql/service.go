package graphql

import (
	"flamingo.me/graphql"

	"github.com/lunarforge/flamingo_commerce/customer/domain"
	"github.com/lunarforge/flamingo_commerce/customer/interfaces/graphql/dtocustomer"
)

//go:generate go run github.com/go-bindata/go-bindata/v3/go-bindata -nometadata -o fs.go -pkg graphql schema.graphql

// Service is the Graphql-Service of this module
type Service struct{}

var _ graphql.Service = new(Service)

// Schema returns graphql schema of this module
func (*Service) Schema() []byte {
	return MustAsset("schema.graphql")
}

// Types configures the GraphQL to Go resolvers
func (*Service) Types(types *graphql.Types) {
	types.Map("Commerce_Customer_Status_Result", dtocustomer.CustomerStatusResult{})
	types.Map("Commerce_Customer_Result", dtocustomer.CustomerResult{})
	types.Map("Commerce_Customer_PersonData", domain.PersonData{})
	types.Map("Commerce_Customer_Address", domain.Address{})
	types.GoField("Commerce_Customer_Address", "streetNumber", "StreetNr")
	types.Resolve("Query", "Commerce_Customer_Status", CustomerResolver{}, "CommerceCustomerStatus")
	types.Resolve("Query", "Commerce_Customer", CustomerResolver{}, "CommerceCustomer")
}
