// Code generated by Gojay. DO NOT EDIT.

package query

import (
	"github.com/francoispqt/gojay"
	"github.com/viant/bigquery/internal"
	"google.golang.org/api/bigquery/v2"
	"reflect"
	"strconv"
	"sync/atomic"
	"unsafe"
)

const endSuffixLen = 6

type TableCell bigquery.TableCell
type ErrorProto bigquery.ErrorProto
type TableFieldSchema bigquery.TableFieldSchema
type TableRow bigquery.TableRow
type JobReference bigquery.JobReference
type TableFieldSchemaCategories bigquery.TableFieldSchemaCategories
type TableFieldSchemaPolicyTags bigquery.TableFieldSchemaPolicyTags
type TableSchema bigquery.TableSchema

type ErrorProtosPtr []*ErrorProto

func (s *ErrorProtosPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {

	var value = &ErrorProto{}
	if err := dec.Object(value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s ErrorProtosPtr) IsNil() bool {
	return len(s) == 0
}

type TableCellsPtr []*TableCell

func (s *TableCellsPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = &TableCell{}
	if err := dec.Object(value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s TableCellsPtr) IsNil() bool {
	return len(s) == 0
}

type TableFieldSchemasPtr []*TableFieldSchema

func (s *TableFieldSchemasPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var value = &TableFieldSchema{}
	if err := dec.Object(value); err != nil {
		return err
	}
	*s = append(*s, value)
	return nil
}

func (s TableFieldSchemasPtr) IsNil() bool {
	return len(s) == 0
}

type TableRowsPtr struct {
	session *internal.Session
}

func (s *TableRowsPtr) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	decCursor := cursor(dec)
	region := internal.Region{Begin: decCursor}
	rowsLen := len(s.session.Rows)
	if rowsLen > 0 {
		s.session.Rows[rowsLen-1].End = region.Begin - len(k) - endSuffixLen
	}
	s.session.Rows = append(s.session.Rows, region)
	return nil
}

func (s *TableRowsPtr) NKeys() int {
	return 1
}

func (s *TableRowsPtr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	return dec.Object(s)
}

type Strings []string

func (s *Strings) UnmarshalJSONArray(dec *gojay.Decoder) error {
	v := ""
	err := dec.String(&v)
	*s = append(*s, v)
	return err
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (p *ErrorProto) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "debugInfo":
		return dec.String(&p.DebugInfo)

	case "location":
		return dec.String(&p.Location)

	case "message":
		return dec.String(&p.Message)

	case "reason":
		return dec.String(&p.Reason)

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (p *ErrorProto) NKeys() int { return 4 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (r *JobReference) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "jobId":
		return dec.String(&r.JobId)

	case "location":
		return dec.String(&r.Location)

	case "projectId":
		return dec.String(&r.ProjectId)

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (r *JobReference) NKeys() int { return 3 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (r *Response) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	rowsLen := len(r.session.Rows)
	if rowsLen > 0 {
		if r.session.Rows[rowsLen-1].End == 0 {
			r.session.Rows[rowsLen-1].End = cursor(dec) - len(k) - endSuffixLen
		}
	}
	switch k {
	case "cacheHit":
		return dec.Bool(&r.CacheHit)

	case "errors":
		var aSlice = ErrorProtosPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			r.Errors = *(*[]*bigquery.ErrorProto)(unsafe.Pointer(&aSlice))
		}
		return err

	case "jobComplete":
		return dec.Bool(&r.JobComplete)

	case "jobReference":
		var value = &JobReference{}
		err := dec.Object(value)
		if err == nil {
			r.JobReference = (*bigquery.JobReference)(value)
		}
		return err

	case "kind":
		return dec.String(&r.Kind)

	case "numDmlAffectedRows":
		return decodeInt64(dec, &r.NumDmlAffectedRows)

	case "pageToken":
		return dec.String(&r.PageToken)

	case "rows":

		var rows = TableRowsPtr{
			session: r.session,
		}
		err := dec.Array(&rows)
		return err
	case "schema":
		if r.session.Schema != nil {
			r.Schema = r.session.Schema
			return nil
		}
		var value = &TableSchema{}
		err := dec.Object(value)
		if err == nil {
			r.Schema = (*bigquery.TableSchema)(value)
			r.session.Schema = r.Schema
			err = r.session.Init(r.Schema)
		}
		return err
	case "totalBytesProcessed":
		return decodeInt64(dec, &r.TotalBytesProcessed)
	case "totalRows":
		err := decodeUint64(dec, &r.TotalRows)
		if err != nil {
			return err
		}
		atomic.StoreUint64(&r.session.TotalRows, r.TotalRows)
	}
	return nil
}

func decodeUint64(dec *gojay.Decoder, target *uint64) error {
	value := ""
	err := dec.String(&value)
	if err == nil {
		if val, err := strconv.Atoi(value); err == nil {
			*target = uint64(val)
		}
	}
	return err
}

// NKeys returns the number of keys to unmarshal
func (r *Response) NKeys() int { return 11 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (c *TableCell) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "v":
		var value = gojay.EmbeddedJSON{}
		err := dec.AddEmbeddedJSON(&value)
		if err == nil && len(value) > 0 {
			c.V = gojay.EmbeddedJSON(value)
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (c *TableCell) NKeys() int { return 1 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (s *TableFieldSchema) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "categories":
		var value = &TableFieldSchemaCategories{}
		err := dec.Object(value)
		if err == nil {
			s.Categories = (*bigquery.TableFieldSchemaCategories)(unsafe.Pointer(value))
		}

		return err

	case "description":
		return dec.String(&s.Description)

	case "fields":
		var aSlice = TableFieldSchemasPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			s.Fields = *(*[]*bigquery.TableFieldSchema)(unsafe.Pointer(&aSlice))
		}
		return err

	case "mode":
		return dec.String(&s.Mode)

	case "name":
		return dec.String(&s.Name)

	case "policyTags":
		var value = &TableFieldSchemaPolicyTags{}
		err := dec.Object(value)
		if err == nil {
			s.PolicyTags = (*bigquery.TableFieldSchemaPolicyTags)(value)
		}
		return err

	case "type":
		return dec.String(&s.Type)

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (s *TableFieldSchema) NKeys() int { return 7 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (c *TableFieldSchemaCategories) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "names":
		var aSlice = Strings{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			c.Names = []string(aSlice)
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (c *TableFieldSchemaCategories) NKeys() int { return 1 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (t *TableFieldSchemaPolicyTags) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "names":
		var aSlice = Strings{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			t.Names = []string(aSlice)
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (t *TableFieldSchemaPolicyTags) NKeys() int { return 1 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (r *TableRow) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "f":
		var aSlice = TableCellsPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			r.F = *(*[]*bigquery.TableCell)(unsafe.Pointer(&aSlice))
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (r *TableRow) NKeys() int { return 1 }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (s *TableSchema) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "fields":
		var aSlice = TableFieldSchemasPtr{}
		err := dec.Array(&aSlice)
		if err == nil && len(aSlice) > 0 {
			s.Fields = *(*[]*bigquery.TableFieldSchema)(unsafe.Pointer(&aSlice))
		}
		return err

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (s *TableSchema) NKeys() int { return 1 }

func decodeInt64(dec *gojay.Decoder, target *int64) error {
	value := ""
	err := dec.String(&value)
	if err == nil {
		if val, err := strconv.Atoi(value); err == nil {
			*target = int64(val)
		}
	}
	return err
}

var curOffset uintptr

func init() {
	cur, ok := reflect.TypeOf(gojay.Decoder{}).FieldByName("cursor")
	if !ok {
		panic("failed to get Decoder.cursor field")
	}
	curOffset = cur.Offset
}

func cursor(dec *gojay.Decoder) int {
	return *(*int)(unsafe.Add(unsafe.Pointer(dec), curOffset))
}
