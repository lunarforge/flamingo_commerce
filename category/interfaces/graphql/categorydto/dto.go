package categorydto

import (
	"github.com/lunarforge/flamingo_commerce/category/domain"
	"github.com/lunarforge/flamingo_commerce/product/interfaces/graphql"
)

// CategorySearchResult represents category search result
type CategorySearchResult struct {
	ProductSearchResult *graphql.SearchResultDTO
	Category            domain.Category
}
