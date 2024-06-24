// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Literary_Snowlands/internal/model"
	"context"
)

type (
	ISearch interface {
		SearchBooks(ctx context.Context, query model.SearchBooksQuery) (*model.SearchBooksInfo, error)
	}
)

var (
	localSearch ISearch
)

func Search() ISearch {
	if localSearch == nil {
		panic("implement not found for interface ISearch, forgot register?")
	}
	return localSearch
}

func RegisterSearch(i ISearch) {
	localSearch = i
}
