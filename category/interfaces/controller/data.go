package controller

import (
	"context"

	"github.com/lunarforge/flamingo_commerce/category/domain"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// Tree controller for category tree retrieval
	Tree struct {
		categoryService domain.CategoryService
	}

	// Entity controller for category entity retrieval
	Entity struct {
		categoryService domain.CategoryService
	}
)

// Inject the Tree controller required dependencies
func (controller *Tree) Inject(service domain.CategoryService) {
	controller.categoryService = service
}

// Data controller for category trees
func (controller *Tree) Data(c context.Context, r *web.Request, params web.RequestParams) interface{} {
	code, _ := params["code"] // no err check, empty code is fine if not set

	treeRoot, _ := controller.categoryService.Tree(c, code)

	return treeRoot
}

// Inject the Entity controller required dependencies
func (controller *Entity) Inject(service domain.CategoryService) {
	controller.categoryService = service
}

// Data controller for category entities
func (controller *Entity) Data(c context.Context, r *web.Request, params web.RequestParams) interface{} {
	code, _ := params["code"] // no err check, empty code is fine if not set

	category, _ := controller.categoryService.Get(c, code)

	return category
}
