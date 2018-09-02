package request

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `field.go`, `field.util.go` or `field_test.go` to
	implement your API calls, helper functions and tests. The file `field.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

var _ = chi.URLParam
var _ = types.JSONText{}

// Field list request parameters
type FieldList struct {
}

func NewFieldList() *FieldList {
	return &FieldList{}
}

func (f *FieldList) Fill(r *http.Request) error {
	var err error
	err = json.NewDecoder(r.Body).Decode(f)
	switch {
	case err == io.EOF:
		err = nil
	case err != nil:
		err = errors.Wrap(err, "error parsing http request body")
	}

	r.ParseForm()
	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	return err
}

var _ RequestFiller = NewFieldList()

// Field type request parameters
type FieldType struct {
	ID string
}

func NewFieldType() *FieldType {
	return &FieldType{}
}

func (f *FieldType) Fill(r *http.Request) error {
	var err error
	err = json.NewDecoder(r.Body).Decode(f)
	switch {
	case err == io.EOF:
		err = nil
	case err != nil:
		err = errors.Wrap(err, "error parsing http request body")
	}

	r.ParseForm()
	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	f.ID = chi.URLParam(r, "id")

	return err
}

var _ RequestFiller = NewFieldType()