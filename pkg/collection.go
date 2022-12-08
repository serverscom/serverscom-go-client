package serverscom

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Collection is an interface for interfacing with the collection
type Collection[K any] interface {
	IsClean() bool

	HasPreviousPage() bool
	HasNextPage() bool
	HasFirstPage() bool
	HasLastPage() bool

	NextPage(ctx context.Context) ([]K, error)
	PreviousPage(ctx context.Context) ([]K, error)
	FirstPage(ctx context.Context) ([]K, error)
	LastPage(ctx context.Context) ([]K, error)

	Collect(ctx context.Context) ([]K, error)
	List(ctx context.Context) ([]K, error)

	SetPage(page int) Collection[K]
	SetPerPage(perPage int) Collection[K]
	SetParam(name, value string) Collection[K]

	Refresh(ctx context.Context) error
}

// CollectionHandler handles operations around collection
type CollectionHandler[K any] struct {
	client *Client

	path string

	params map[string]string

	clean bool

	rels       map[string]string
	collection []K
}

// NewCollection produces a new CollectionHandler and represents this as an interface of Collection
func NewCollection[K any](client *Client, path string) *CollectionHandler[K] {
	return &CollectionHandler[K]{
		client: client,

		path: path,

		params:     make(map[string]string),
		rels:       make(map[string]string),
		clean:      true,
		collection: make([]K, 0),
	}
}

// IsClean returns a bool value where true is means, this collection not used yet and doesn't contain any state.
func (col *CollectionHandler[K]) IsClean() bool {
	return col.clean
}

// HasPreviousPage returns a bool value where truth is means collection has a previous page.
//
// In case when IsClean returns true, this method will return false, which means no request(s)
// were made and collection doesn't have metadata to know about pagination.
//
// First metadata will come with the first called methods such: NextPage, PreviousPage, LastPage, FirstPage, List, Refresh, Collect.
func (col *CollectionHandler[K]) HasPreviousPage() bool {
	return col.hasRel("prev")
}

// HasNextPage returns a bool value where truth is means collection has a next page.
//
// In case when IsClean returns true, this method will return false, which means no request(s)
// were made and collection doesn't have metadata to know about pagination.
//
// First metadata will come with the first called methods such: NextPage, PreviousPage, LastPage, FirstPage, List, Refresh, Collect.
func (col *CollectionHandler[K]) HasNextPage() bool {
	return col.hasRel("next")
}

// HasFirstPage returns a bool value where truth is means collection has a first page.
//
// In case when IsClean returns true, this method will return false, which means no request(s)
// were made and collection doesn't have metadata to know about pagination.
//
// First metadata will come with the first called methods such: NextPage, PreviousPage, LastPage, FirstPage, List, Refresh, Collect.
func (col *CollectionHandler[K]) HasFirstPage() bool {
	return col.hasRel("first")
}

// HasLastPage returns a bool value where truth is means collection has a last page.
//
// In case when IsClean returns true, this method will return false, which means no request(s)
// were made and collection doesn't have metadata to know about pagination.
//
// First metadata will come with the first called methods such: NextPage, PreviousPage, LastPage, FirstPage, List, Refresh, Collect.
func (col *CollectionHandler[K]) HasLastPage() bool {
	return col.hasRel("last")
}

// NextPage navigates to the next page returns a []Network, produces an error, when a
// collection has no next page.
//
// Before using this method please ensure IsClean returns false and HasNextPage returns true.
// You can force to load pagination metadata by calling Refresh or List methods.
func (col *CollectionHandler[K]) NextPage(ctx context.Context) ([]K, error) {
	return col.navigate(ctx, "next")
}

// PreviousPage navigates to the previous page returns a []Network, produces an error, when a
// collection has no previous page.
//
// Before using this method please ensure IsClean returns false and HasPreviousPage returns true.
// You can force to load pagination metadata by calling Refresh or List methods.
func (col *CollectionHandler[K]) PreviousPage(ctx context.Context) ([]K, error) {
	return col.navigate(ctx, "prev")
}

// FirstPage navigates to the first page returns a []Network, produces an error, when a
// collection has no first page.
//
// Before using this method please ensure IsClean returns false and HasFirstPage returns true.
// You can force to load pagination metadata by calling Refresh or List methods.
func (col *CollectionHandler[K]) FirstPage(ctx context.Context) ([]K, error) {
	return col.navigate(ctx, "first")
}

// LastPage navigates to the last page returns a []Network, produces an error, when a
// collection has no last page.
//
// Before using this method please ensure IsClean returns false and HasLastPage returns true.
// You can force to load pagination metadata by calling Refresh or List methods.
func (col *CollectionHandler[K]) LastPage(ctx context.Context) ([]K, error) {
	return col.navigate(ctx, "last")
}

// Collect navigates by pages until the last page is reached will be reached and returns accumulated data between pages.
//
// This method uses NextPage.
func (col *CollectionHandler[K]) Collect(ctx context.Context) ([]K, error) {
	var accumulatedCollectionElements []K

	currentCollectionElements, err := col.List(ctx)

	if err != nil {
		return nil, err
	}

	accumulatedCollectionElements = append(accumulatedCollectionElements, currentCollectionElements...)

	for col.HasNextPage() {
		nextCollectionElements, err := col.NextPage(ctx)

		if err != nil {
			return nil, err
		}

		accumulatedCollectionElements = append(accumulatedCollectionElements, nextCollectionElements...)
	}

	return accumulatedCollectionElements, nil
}

// List returns a []BandwidthOption limited by pagination.
//
// This method performs request only once when IsClean returns false, also this request doesn't
// perform request when such methods were called before: NextPage, PreviousPage, LastPage, FirstPage, Refresh, Collect.
//
// In the case when previously called method is Collect, this method returns data from the last page.
func (col *CollectionHandler[K]) List(ctx context.Context) ([]K, error) {
	if col.IsClean() {
		if err := col.Refresh(ctx); err != nil {
			return nil, err
		}
	}

	return col.collection, nil
}

// SetPage sets current page param.
func (col *CollectionHandler[K]) SetPage(page int) Collection[K] {
	var currentPage string

	if page > 1 {
		currentPage = strconv.Itoa(page)
	} else {
		currentPage = ""
	}

	col.applyParam("page", currentPage)

	return col
}

// SetPerPage sets current per page param.
func (col *CollectionHandler[K]) SetPerPage(perPage int) Collection[K] {
	var currentPerPage string

	if perPage > 0 {
		currentPerPage = strconv.Itoa(perPage)
	} else {
		currentPerPage = ""
	}

	col.applyParam("per_page", currentPerPage)

	return col
}

// SetParam sets param.
func (col *CollectionHandler[K]) SetParam(name, value string) Collection[K] {
	col.applyParam(name, value)

	return col
}

// Refresh performs the request and then updates accumulated data limited by pagination.
//
// After calling this method accumulated data can be extracted by List method.
func (col *CollectionHandler[K]) Refresh(ctx context.Context) error {
	if err := col.fireHTTPRequest(ctx); err != nil {
		return err
	}

	return nil
}

func (col *CollectionHandler[K]) fireHTTPRequest(ctx context.Context) error {
	var accumulatedCollectionElements []K

	initialURL := col.client.buildURL(col.path)
	url := col.client.applyParams(
		initialURL,
		col.params,
	)

	response, body, err := col.client.buildAndExecRequestWithResponse(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &accumulatedCollectionElements); err != nil {
		return err
	}

	col.clean = false
	col.collection = accumulatedCollectionElements
	col.rels = hyperHeaderParser(response.Header)

	return nil
}

func (col *CollectionHandler[K]) navigate(ctx context.Context, name string) ([]K, error) {
	if col.IsClean() {
		if err := col.Refresh(ctx); err != nil {
			return nil, err
		}
	}

	if err := col.applyRel(name); err != nil {
		return nil, err
	}

	if err := col.Refresh(ctx); err != nil {
		return nil, err
	}

	return col.collection, nil
}

func (col *CollectionHandler[K]) applyParam(name, value string) {
	if value == "" {
		delete(col.params, name)
	} else {
		col.params[name] = value
	}
}

func (col *CollectionHandler[K]) applyRel(name string) error {
	if !col.hasRel(name) {
		return fmt.Errorf("No rel for: %s", name)
	}

	url, err := url.Parse(col.rels[name])

	if err != nil {
		return err
	}

	col.applyParam("page", url.Query().Get("page"))
	col.applyParam("per_page", url.Query().Get("per_page"))

	return nil
}

func (col *CollectionHandler[K]) hasRel(name string) bool {
	if _, ok := col.rels[name]; ok {
		return true
	}

	return false
}
