package graphql

import (
	"context"
	productApplication "github.com/lunarforge/flamingo_commerce/product/application"
	"github.com/lunarforge/flamingo_commerce/product/domain"
	productDto "github.com/lunarforge/flamingo_commerce/product/interfaces/graphql/product/dto"
	"github.com/lunarforge/flamingo_commerce/search/application"
	searchDomain "github.com/lunarforge/flamingo_commerce/search/domain"
	"github.com/lunarforge/flamingo_commerce/search/interfaces/graphql/searchdto"
)

// CommerceProductQueryResolver resolves graphql product queries
type CommerceProductQueryResolver struct {
	productService domain.ProductService
	searchService  *productApplication.ProductSearchService
}

// Inject dependencies
func (r *CommerceProductQueryResolver) Inject(
	productService domain.ProductService,
	searchService *productApplication.ProductSearchService,
) *CommerceProductQueryResolver {
	r.productService = productService
	r.searchService = searchService
	return r
}

// CommerceProduct returns a product with the given marketplaceCode from productService
func (r *CommerceProductQueryResolver) CommerceProduct(ctx context.Context, marketplaceCode string, variantMarketPlaceCode *string) (productDto.Product, error) {

	product, err := r.productService.Get(ctx, marketplaceCode)

	if err != nil {
		return nil, err
	}

	return productDto.NewGraphqlProductDto(product, variantMarketPlaceCode), nil
}

// CommerceProductSearch returns a search result of products based on the given search request
func (r *CommerceProductQueryResolver) CommerceProductSearch(ctx context.Context, request searchdto.CommerceSearchRequest) (*SearchResultDTO, error) {

	var filters []searchDomain.Filter
	for _, filter := range request.KeyValueFilters {
		filters = append(filters, searchDomain.NewKeyValueFilter(filter.K, filter.V))
	}

	result, err := r.searchService.Find(ctx, &application.SearchRequest{
		AdditionalFilter: filters,
		PageSize:         request.PageSize,
		Page:             request.Page,
		SortBy:           request.SortBy,
		Query:            request.Query,
		PaginationConfig: nil,
	})

	if err != nil {
		return nil, err
	}

	return WrapSearchResult(result), nil
}
