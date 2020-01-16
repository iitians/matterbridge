// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OnenoteSectionRequestBuilder is request builder for OnenoteSection
type OnenoteSectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteSectionRequest
func (b *OnenoteSectionRequestBuilder) Request() *OnenoteSectionRequest {
	return &OnenoteSectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteSectionRequest is request for OnenoteSection
type OnenoteSectionRequest struct{ BaseRequest }

// Get performs GET request for OnenoteSection
func (r *OnenoteSectionRequest) Get(ctx context.Context) (resObj *OnenoteSection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenoteSection
func (r *OnenoteSectionRequest) Update(ctx context.Context, reqObj *OnenoteSection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenoteSection
func (r *OnenoteSectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Pages returns request builder for OnenotePage collection
func (b *OnenoteSectionRequestBuilder) Pages() *OnenoteSectionPagesCollectionRequestBuilder {
	bb := &OnenoteSectionPagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pages"
	return bb
}

// OnenoteSectionPagesCollectionRequestBuilder is request builder for OnenotePage collection
type OnenoteSectionPagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenotePage collection
func (b *OnenoteSectionPagesCollectionRequestBuilder) Request() *OnenoteSectionPagesCollectionRequest {
	return &OnenoteSectionPagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenotePage item
func (b *OnenoteSectionPagesCollectionRequestBuilder) ID(id string) *OnenotePageRequestBuilder {
	bb := &OnenotePageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteSectionPagesCollectionRequest is request for OnenotePage collection
type OnenoteSectionPagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenotePage, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnenotePage
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnenotePage
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Get(ctx context.Context) ([]OnenotePage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Add(ctx context.Context, reqObj *OnenotePage) (resObj *OnenotePage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ParentNotebook is navigation property
func (b *OnenoteSectionRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSectionGroup is navigation property
func (b *OnenoteSectionRequestBuilder) ParentSectionGroup() *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSectionGroup"
	return bb
}