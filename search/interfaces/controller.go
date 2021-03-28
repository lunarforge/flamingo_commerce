package interfaces

import (
	"context"
	"net/url"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/lunarforge/flamingo_commerce/search/application"
	"github.com/lunarforge/flamingo_commerce/search/domain"
	"github.com/lunarforge/flamingo_commerce/search/utils"
)

type (
	// ViewController demonstrates a search view controller
	ViewController struct {
		responder             *web.Responder
		searchService         *application.SearchService
		paginationInfoFactory *utils.PaginationInfoFactory
	}

	viewData struct {
		SearchMeta     domain.SearchMeta
		SearchResult   map[string]*application.SearchResult
		PaginationInfo utils.PaginationInfo
	}
)

// Inject dependencies
func (vc *ViewController) Inject(responder *web.Responder,
	paginationInfoFactory *utils.PaginationInfoFactory,
	searchService *application.SearchService,
) *ViewController {
	vc.responder = responder
	vc.paginationInfoFactory = paginationInfoFactory
	vc.searchService = searchService

	return vc
}

//http://localhost:3210/en/search/products?category=flat-screen_tvs&q=knapsack
//registry.Route("/search/:type", `search.search(type, *)`)
//registry.Route("/search", `search.search`)

// Get godoc
// Get Response for Search matching :type param based on query q
// @Summary Searches for requested items
// @Tags  Search Products Category
// @Produce json
// @Success 200 {object} APIResult{product=domain.SimpleProduct}
// @Failure 500 {object} APIResult
// @Failure 404 {object} APIResult
// @Param type path string true "search selection type"
// @Router /search/{type} [get]
func (vc *ViewController) Get(c context.Context, r *web.Request) web.Result {
	query, _ := r.Query1("q")

	vd := viewData{}

	searchRequest := application.SearchRequest{
		Query: query,
	}
	for k, v := range r.QueryAll() {
		switch k {
		case "q":
			continue
		case "page":
			page, _ := strconv.ParseInt(v[0], 10, 64)
			searchRequest.Page = int(page)
		case "sort":
			searchRequest.SortBy = v[0]
		default:
			searchRequest.SetAdditionalFilter(domain.NewKeyValueFilter(k, v))
		}
	}

	if typ, ok := r.Params["type"]; ok {
		//Search for a specific type of documents:
		searchResult, err := vc.searchService.FindBy(c, typ, searchRequest)
		if err != nil {
			if re, ok := err.(*domain.RedirectError); ok {
				u, _ := url.Parse(re.To)
				return vc.responder.URLRedirect(u).Permanent()
			}
			return vc.responder.ServerError(err)
		}
		vd.SearchMeta = searchResult.SearchMeta
		vd.SearchMeta.Query = query
		vd.SearchResult = map[string]*application.SearchResult{typ: searchResult}
		vd.PaginationInfo = vc.buildPagination(searchResult, searchRequest.PageSize, r.Request().URL)
		return vc.responder.Render("search/"+typ, vd)
	}

	//Search for all types of documents
	searchResult, err := vc.searchService.Find(c, searchRequest)
	if err != nil {
		if re, ok := err.(*domain.RedirectError); ok {
			u, _ := url.Parse(re.To)
			return vc.responder.URLRedirect(u).Permanent()
		}
		return vc.responder.ServerError(err)
	}
	vd.SearchMeta.Query = query
	vd.SearchMeta.OriginalQuery = query
	vd.SearchResult = searchResult

	return vc.responder.Render("search/search", vd)
}

func (vc *ViewController) buildPagination(searchResult *application.SearchResult, pageSize int, url *url.URL) utils.PaginationInfo {
	return vc.paginationInfoFactory.Build(
		searchResult.SearchMeta.Page,
		searchResult.SearchMeta.NumResults,
		pageSize,
		searchResult.SearchMeta.NumPages,
		url,
	)
}
