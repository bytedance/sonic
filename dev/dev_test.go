package dev_test

import (
	"encoding/json"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/dev"
	"github.com/bytedance/sonic/dev/internal"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func testUnmarshalI(t *testing.T, data string) {
	var got, expect interface{}
	yerr := dev.UnmarshalString(data, &got)
	gerr := json.Unmarshal([]byte(data), &expect)

	assert.Nil(t, yerr)
	assert.Nil(t, gerr)
	assert.Equal(t, expect, got)
}

func TestUnmarshalI(t *testing.T) {
	testUnmarshalI(t, "{}")
	testUnmarshalI(t, `"abc"`)
	testUnmarshalI(t, `{"a": "\\true", "\\b": {"a": [1, 2, 3] }}`)
	testUnmarshalI(t, METADATA)
	testUnmarshalI(t, Twitter_JSON)
}

func BenchmarkUnmarshalIface(b *testing.B) {
	b.Run("Twitter_Sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sval interface{}
			_ = sonic.UnmarshalString(Twitter_JSON, &sval)
		}
	})

	b.Run("Twitter_dev", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sval interface{}
			_ = dev.UnmarshalString(Twitter_JSON, &sval)
		}
	})

	b.Run("Twitter_Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cdom, err := internal.Parse(Twitter_JSON, 0)
			if err != nil {
				cdom.Delete()
			}
		}
	})

	b.Run("Twitter_Std", func(b *testing.B) {
		buf := []byte(Twitter_JSON)
		for i := 0; i < b.N; i++ {
			var val interface{}
			_ = json.Unmarshal(buf, &val)
		}
	})

}

const METADATA = `{
    "max_id": 250126199840518145,
    "since_id": 24012619984051000,
    "refresh_url": "?since_id=250126199840518145&q=%23freebandnames&result_type=mixed&include_entities=1",
    "next_results": "?max_id=249279667666817023&q=%23freebandnames&count=4&include_entities=1&result_type=mixed",
    "count": 4,
	"bool": true,
    "completed_in": 0.035,
    "since_id_str": "24012619984051000",
    "query": "%23freebandnames",
    "max_id_str": "250126199840518145",
	"indics": [
		20, 34, 1, 2, 3, 4, 5
	]
  }`

type searchMetadata struct {
	MaxID       int64    `json:"max_id"`
	SinceID     ***int64 `json:"since_id"`
	RefreshURL  string   `json:"refresh_url"`
	NextResults string   `json:"next_results"`
	Count       int      `json:"count"`
	Bool        bool     `json:"bool"`
	CompletedIn float64  `json:"completed_in"`
	SinceIDStr  string   `json:"since_id_str"`
	Query       string   `json:"query"`
	MaxIDStr    *string  `json:"max_id_str"`
	Indics      []int64  `json:"indics"`
}

var fieldMap = map[string]int{
	"max_id":       10,
	"since_id":     1,
	"refresh_url":  2,
	"next_results": 3,
	"count":        4,
	"bool":         5,
	"completed_in": 6,
	"since_id_str": 7,
	"query":        8,
	"max_id_str":   9,
	"indics":       11,
}

func unmarshalSearchMetadata(data string) (searchMetadata, error) {
	var ret searchMetadata
	ctx, err := internal.NewContext(data, 0)
	if err != nil {
		return ret, nil
	}

	node := ctx.Dom.Root()
	obj, err := node.AsObj()
	if err != nil {
		return ret, err
	}

	size := obj.Len()
	next := obj.Children()
	for i := 0; i < size; i++ {
		key := internal.NewNode(next).String(&ctx)
		val := internal.NewNode(internal.PtrOffset(next, 1))
		idx := fieldMap[key]
		switch idx {
		case 10:
			{
				ret.MaxID, err = val.AsI64()
				if err != nil {
					return ret, err
				}
			}
		case 1:
			id, err := val.AsI64()
			id2 := &id
			id3 := &id2
			if err != nil {
				return ret, err
			}
			ret.SinceID = &id3
		case 2:
			ret.RefreshURL, err = val.AsStr(&ctx)
			if err != nil {
				return ret, err
			}
		case 3:
			ret.NextResults, err = val.AsStr(&ctx)
			if err != nil {
				return ret, err
			}
		case 4:
			count, err := val.AsI64()
			ret.Count = int(count)
			if err != nil {
				return ret, err
			}
		case 5:
			ret.Bool, err = val.AsBool()
			if err != nil {
				return ret, err
			}
		case 6:
			ret.CompletedIn, err = val.AsF64()
			if err != nil {
				return ret, err
			}
		case 7:
			ret.SinceIDStr, err = val.AsStr(&ctx)
			if err != nil {
				return ret, err
			}
		case 8:
			ret.Query, err = val.AsStr(&ctx)
			if err != nil {
				return ret, err
			}
		case 9:
			s, err := val.AsStr(&ctx)
			if err != nil {
				return ret, err
			}
			ret.MaxIDStr = &s
		case 11:
			err := val.AsSliceI64(&ctx, unsafe.Pointer(&ret.Indics))
			if err != nil {
				return ret, err
			}
		default:
			// unknow fields
		}
		next = internal.PtrOffset(next, 2)
	}

	ctx.Delete()
	return ret, nil
}

func TestUnmarshalC_MetaData(t *testing.T) {
	var got, expect, xgot searchMetadata
	got, err := unmarshalSearchMetadata(METADATA)
	gerr := json.Unmarshal([]byte(METADATA), &expect)

	xerr := dev.UnmarshalString(METADATA, &xgot)

	spew.Dump(got, expect, xgot)
	assert.Nil(t, err)
	assert.Nil(t, xerr)
	assert.Nil(t, gerr)
	assert.Equal(t, xgot, expect)
	assert.Equal(t, got, expect)
}

func TestUnmarshalC_Twitter(t *testing.T) {
	var got, expect TwitterStruct
	gerr := json.Unmarshal([]byte(Twitter_JSON), &expect)
	xerr := dev.UnmarshalString(Twitter_JSON, &got)

	assert.Nil(t, gerr)
	assert.Nil(t, xerr)
	assert.Equal(t, expect, got)
}

func BenchmarkUnmarshalConcret(b *testing.B) {

	b.Run("MetaData_CodeGen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var got searchMetadata
			got, _ = unmarshalSearchMetadata(METADATA)
			_ = got
		}
	})

	b.Run("MetaData_dev", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var got searchMetadata
			_ = dev.UnmarshalString(METADATA, &got)
		}
	})

	b.Run("MetaData_Sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var got searchMetadata
			_ = sonic.UnmarshalString(METADATA, &got)
		}
	})

	b.Run("MetaData_Std", func(b *testing.B) {
		buf := []byte(METADATA)
		for i := 0; i < b.N; i++ {
			var got searchMetadata
			_ = json.Unmarshal(buf, &got)
		}
	})

	b.Run("Twitter_dev", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var got TwitterStruct
			_ = dev.UnmarshalString(Twitter_JSON, &got)
		}
	})

	b.Run("Twitter_Sonic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var got TwitterStruct
			_ = sonic.UnmarshalString(Twitter_JSON, &got)
		}
	})

	b.Run("Twitter_Std", func(b *testing.B) {
		buf := []byte(Twitter_JSON)
		for i := 0; i < b.N; i++ {
			var got TwitterStruct
			_ = json.Unmarshal(buf, &got)
		}
	})
}
