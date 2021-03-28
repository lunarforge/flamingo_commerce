package graphql

import (
	"context"
	"github.com/lunarforge/flamingo_commerce/category/domain"
	graphqlDto "github.com/lunarforge/flamingo_commerce/category/interfaces/graphql/categorydto"
	productApplication "github.com/lunarforge/flamingo_commerce/product/application"
	"github.com/lunarforge/flamingo_commerce/product/interfaces/graphql"
	"github.com/lunarforge/flamingo_commerce/search/application"
	searchDomain "github.com/lunarforge/flamingo_commerce/search/domain"
	"github.com/lunarforge/flamingo_commerce/search/interfaces/graphql/searchdto"
)

// CommerceCategoryQueryResolver resolves graphql category queries
type CommerceCategoryQueryResolver struct {
	categoryService domain.CategoryService
	searchService   *productApplication.ProductSearchService
}

// Inject dependencies
func (r *CommerceCategoryQueryResolver) Inject(
	service domain.CategoryService,
	searchService *productApplication.ProductSearchService,
) *CommerceCategoryQueryResolver {
	r.categoryService = service
	r.searchService = searchService
	return r
}

// CommerceCategoryTree returns a Tree with the given activeCategoryCode from categoryService
func (r *CommerceCategoryQueryResolver) CommerceCategoryTree(ctx context.Context, activeCategoryCode string) (domain.Tree, error) {
	return r.categoryService.Tree(ctx, activeCategoryCode)
}

// CommerceCategory returns product search result with the given categoryCode from searchService
func (r *CommerceCategoryQueryResolver) CommerceCategory(
	ctx context.Context,
	categoryCode string,
	request *searchdto.CommerceSearchRequest) (*graphqlDto.CategorySearchResult, error) {
	category, err := r.categoryService.Get(ctx, categoryCode)

	if err != nil {
		return nil, err
	}

	searchRequest := &application.SearchRequest{}
	if request != nil {
		var filters []searchDomain.Filter
		for _, filter := range request.KeyValueFilters {
			filters = append(filters, searchDomain.NewKeyValueFilter(filter.K, filter.V))
		}

		filters = append(filters, domain.NewCategoryFacet(categoryCode))
		searchRequest = &application.SearchRequest{
			AdditionalFilter: filters,
			PageSize:         request.PageSize,
			Page:             request.Page,
			SortBy:           request.SortBy,
			Query:            request.Query,
			PaginationConfig: nil,
		}
	}
	result, err := r.searchService.Find(ctx, searchRequest)

	if err != nil {
		return nil, err
	}

	return &graphqlDto.CategorySearchResult{Category: category, ProductSearchResult: graphql.WrapSearchResult(result)}, nil
}
