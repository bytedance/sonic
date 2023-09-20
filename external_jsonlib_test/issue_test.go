/**
 * Copyright 2023 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package external_jsonlib_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/klauspost/compress/gzip"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Snapshot struct {
   Revision int64             `json:"rev"`
   Sheets   map[string]*Sheet `json:"sheets,omitempty"`
}

type InnerSheet struct {
   RowCount       int32  `json:"rowCount,omitempty"`
   ColumnCount    int32  `json:"columnCount,omitempty"`
   Index          int32  `json:"index"`
   FrozenRowCount int32  `json:"frozenRowCount,omitempty"`
   FrozenColCount int32  `json:"frozenColCount,omitempty"`
   Name           string `json:"name,omitempty"`
   ID             string `json:"id,omitempty"`
   DelTime        string `json:"delTime,omitempty"`
   Hidden         bool   `json:"hidden"`
   GridLineHidden bool   `json:"gridLineHidden"`
   Version        int32  `json:"version"`

   ExtResource *Resource `json:"externalResource,omitempty"`
}

type Resource struct {
   BlockType  string `json:"blockType,omitempty"`
   BlockToken string `json:"blockToken,omitempty"`
   BlockID    string `json:"blockId,omitempty"`
}

type Sheet struct {
   InnerSheet
   Blocks           []*SheetBlock             `json:"blocks,omitempty"`
   Comment          CommentMeta               `json:"comment,omitempty"`
   Image            ImageMeta                 `json:"image,omitempty"`
   FloatImageMap    FloatImageMeta            `json:"floatImageMap,omitempty"`
   Attachment       AttachmentMeta            `json:"attachment,omitempty"`
   Reminders        ReminderMeta              `json:"reminders,omitempty"`
   Spans            SpanModel                 `json:"spans,omitempty"`
   Columns          []*RowColInfo             `json:"columns,omitempty"`
   Rows             []*RowColInfo             `json:"rows,omitempty"`
   RowFilter        *HideRowFilter            `json:"rowFilter,omitempty"`
   ChartMap         map[string]*ChartBlock    `json:"chartMap,omitempty"`
   ConditionFormats []ConditionFormat         `json:"conditionalFormats,omitempty"`
   Mention          map[int32]*Mention        `json:"mention,omitempty"`
   Entities         *Entities                 `json:"entities,omitempty"`
   Views            []*BaseView               `json:"views,omitempty"`
   DataValidations  map[int32]*DataValidation `json:"dataValidations,omitempty"`
   FloatBlocks      map[string]*FloatBlock    `json:"floatBlocks,omitempty"`
   EmbeddedBlocks   map[string]EmbeddedBlock  `json:"embeddedBlocks,omitempty"`
   Perms            Permission                `json:"perms,omitempty"`
}

type Perm struct {
   Info      string `json:"info,omitempty"`
   Range     *Range `json:"range,omitempty"`
   CreatorID string `json:"creatorId,omitempty"` // 创建人ID
}

type Permissions map[string]Perm

type Permission struct {
   Sheet   Permissions `json:"sheet,omitempty"`
   Rows    Permissions `json:"rows,omitempty"`
   Columns Permissions `json:"columns,omitempty"`
}

type EmbeddedBlock struct { // 新增结构
   BlockType    string          `json:"blockType"`  // 类型：透视表 PIVOT_TABLE
   BlockToken   string          `json:"blockToken"` // 内嵌block token
   Position     CellPosition    `json:"position"`
   BlockSetting json.RawMessage `json:"blockSetting,omitempty"`
}

type BaseView struct {
   ID   string `json:"id,omitempty"`
   Name string `json:"name,omitempty"`
}

type View struct {
   BaseView
   RowFilter *RowFilter `json:"rowFilter,omitempty"`
}

type RowFilter struct {
   FilterRules map[int32]*FilterViewCondition `json:"filterRules,omitempty"`
   Range       *Range                         `json:"range,omitempty"`
}

type FilterViewCondition struct {
   ConType     int32       `json:"conType"`
   CompareType int32       `json:"compareType"`
   Contents    interface{} `json:"contents"`
}

type Entities struct {
   StyleMap        map[int32]*Style    `json:"styles,omitempty"`
   DropDownMap     map[int32]*DropDown `json:"dropdowns,omitempty"`
   SegmentStyleMap map[int32]*Style    `json:"segmentStyles,omitempty"`
}

type Style struct {
   Font           string                 `json:"font,omitempty"`
   Color          string                 `json:"foreColor,omitempty"`
   BgColor        string                 `json:"backColor,omitempty"`
   HAlign         interface{}            `json:"hAlign,omitempty"`
   VAlign         interface{}            `json:"vAlign,omitempty"`
   TextDecoration int                    `json:"textDecoration,omitempty"`
   WordWrap       interface{}            `json:"wordWrap,omitempty"`
   Formatter      interface{}            `json:"formatter,omitempty"`
   AutoFormatter  interface{}            `json:"autoFormatter,omitempty"`
   DropdownID     *int32                 `json:"dropdown,omitempty"`
   BorderBottom   map[string]interface{} `json:"borderBottom,omitempty"`
   BorderLeft     map[string]interface{} `json:"borderLeft,omitempty"`
   BorderRight    map[string]interface{} `json:"borderRight,omitempty"`
   BorderTop      map[string]interface{} `json:"borderTop,omitempty"`
   FontInfo       *FontInfo              `json:"fontInfo,omitempty"`
}

type FontInfo struct {
   FontStyle  *string  `json:"fontStyle,omitempty"`  // 'normal' | 'italic'
   FontWeight *float64 `json:"fontWeight,omitempty"` // 400常规 700加粗
   FontSize   *float64 `json:"fontSize,omitempty"`
   FontFamily *string  `json:"fontFamily,omitempty"`
}

type DropDown struct {
   IsEnableOptionColor bool     `json:"isEnableOptionColor,omitempty"`
   IsValid             bool     `json:"isValid"`
   List                []string `json:"list"`
   Color               []string `json:"color,omitempty"`
}

type DataValidation struct {
   Type     string        `json:"type"`
   Formulas []interface{} `json:"formulas"`
}

type Mention struct {
   MentionType int32 `json:"mentionType"`
}

type ConditionFormat struct {
   HasRef   bool            `json:"hasRef"`
   CFId     string          `json:"cfId"`
   RuleType string          `json:"ruleType"`
   Ranges   []RangeV2       `json:"ranges"`
   Style    json.RawMessage `json:"style"`
   Attrs    json.RawMessage `json:"attrs,omitempty"`
}

type ChartBlockSettingEx struct {
   CellPosition
   ChartBlockSetting
}

type DataRange struct {
   SheetID string `json:"sheetId"`
   Range
}

type RangeV2Type int

const (
   RangeV2Normal  RangeV2Type = 0
   RangeV2Row     RangeV2Type = 1
   RangeV2Col     RangeV2Type = 2
   RangeV2All     RangeV2Type = 3
   RangeV2Cell    RangeV2Type = 4
   RangeV2Invalid RangeV2Type = 5
)

type RangeV2 struct {
   Row      *int32      `json:"row,omitempty"`
   RowCount int32       `json:"rowCount,omitempty"`
   Col      *int32      `json:"col,omitempty"`
   ColCount int32       `json:"colCount,omitempty"`
   Type     RangeV2Type `json:"type"`
}

type FloatBlock struct {
   BlockSetting json.RawMessage `json:"blockSetting"`
   BlockToken   string          `json:"blockToken"` // 图表token
   BlockType    string          `json:"blockType"`  // 类型： 图表
}
type ChartBlock struct {
   SheetId      string              `json:"sheetId"`
   ChartId      string              `json:"chartId"`
   BlockSetting ChartBlockSettingEx `json:"blockSetting"`
   GraphSetting ChartGraphSetting   `json:"graphSetting"`
}

type ChartBlockSetting struct {
   CellPositionOffset
   Width     float32   `json:"width"`
   Height    float32   `json:"height"`
   TimeStamp int64     `json:"timestamp"`
   Padding   []float32 `json:"padding"`
   // z轴，用于确定层级前后关系
   ZIndex *int32 `json:"zIndex,omitempty"`
}

type ChartGraphSetting struct {
   IsStack           bool         `json:"isStack"`
   IsTitleHidden     bool         `json:"isTitleHidden"`
   IsDataLabelShown  bool         `json:"isDataLabelShown"`
   IsGridLineHidden  bool         `json:"isGridLineHidden"`
   IsAggregate       bool         `json:"isAggregate"`
   LegendPosition    int32        `json:"legendPosition"`
   Title             string       `json:"title"`
   Type              string       `json:"type"`
   XAxisPos          string       `json:"xAxisPos"`
   CombineDirection  string       `json:"combineDirection"`
   ChartColorType    string       `json:"chartColorType"`
   MainYaxisMinValue string       `json:"mainYaxisMinValue"`
   MainYaxisMaxValue string       `json:"mainYaxisMaxValue"`
   HasHeaders        *bool        `json:"hasHeaders,omitempty"`
   HasLabels         *bool        `json:"hasLabels,omitempty"`
   DataRanges        []*DataRange `json:"dataRanges"`
}

type HideRowFilterItem struct {
   Conditions   []*Condition `json:"conditions,omitempty"`
   Index        int32        `json:"index"`
   FilteredRows []int32      `json:"filteredRows"`
}

type Condition struct {
   ConType     int32       `json:"conType"`
   CompareType int32       `json:"compareType"`
   Expected    interface{} `json:"expected"`
}

type HideRowFilter struct {
   FilterButtonVisibleInfo map[string]bool      `json:"filterButtonVisibleInfo,omitempty"`
   FilterItemMap           []*HideRowFilterItem `json:"filterItemMap,omitempty"`
   FilteredColumns         []int32              `json:"filteredColumns,omitempty"`
   Range                   *Range               `json:"range,omitempty"`
   ShowFilterButton        bool                 `json:"showFilterButton,omitempty"`
   FilteredOutRows         []int32              `json:"filteredOutRows,omitempty"`
}

type RowColInfo struct {
   PageBreak bool     `json:"pageBreak,omitempty"`
   Resizable bool     `json:"resizable,omitempty"`
   Size      float64  `json:"size,omitempty"`
   Visible   *bool    `json:"visible,omitempty"`
   FixedSize *float64 `json:"fixedSize,omitempty"`
   FID       string   `json:"fId,omitempty"`
}

type SpanModel []*Range
type Range struct {
   Row      int32 `json:"row"`
   RowCount int32 `json:"rowCount,omitempty"`
   Col      int32 `json:"col"`
   ColCount int32 `json:"colCount,omitempty"`
}

type Reminder struct {
   NotifyTime     int64         `json:"notifyTime"`
   ExpireTime     int64         `json:"expireTime"`
   NotifyUserIds  []string      `json:"notifyUserIds"`
   NotifyStrategy int32         `json:"notifyStrategy"`
   NotifyText     *string       `json:"notifyText,omitempty"`
   NotifyCell     *ReminderCell `json:"notifyCell,omitempty"`
}

type ReminderCell struct {
   Row int32 `json:"row"`
   Col int32 `json:"col"`
}

type ReminderMeta map[string]*Reminder

type Comment struct {
   ID           string `json:"id"`
   MentionKeyID string `json:"mentionKeyId"`
   IsResolved   bool   `json:"isResolved"`
}

type SheetBlock struct {
   Revision      int64      `json:"revision"`
   BlockID       int64      `json:"blockId"`
   StartRowIndex int32      `json:"startRowIndex"`
   RowsCount     int32      `json:"rowsCount"`
   Size          int32      `json:"size"`
   Token         BlockToken `json:"token"`
   gzipDataTable []byte
}
type CommentMeta map[int32]map[int32][]Comment
type BlockToken string
type FloatImageMeta map[string]*FloatImageInfo
type ImageMeta map[string]map[string]interface{}
type AttachmentMeta map[string]map[string]interface{}

type FloatImageInfo struct {
   SheetId             string `json:"sheetId"`
   FloatImageId        string `json:"floatImageId"`
   FloatImageSettingEx `json:"floatImageSetting"`
}

type FloatImageSettingEx struct {
   CellPosition
   FloatImageSetting
}

type FloatImageSetting struct {
   CellPositionOffset
   // 宽高几何信息
   Width  float32 `json:"width"`
   Height float32 `json:"height"`
   // z轴，用于确定层级前后关系
   ZIndex int32 `json:"zIndex"`
   // 图片信息
   ImageToken string `json:"imageToken"`
}

type CellPosition struct { // 定位单元格
   Row int32 `json:"row"`
   Col int32 `json:"col"`
}

type CellPositionOffset struct { // 相对位置偏移量
   X float32 `json:"x"`
   Y float32 `json:"y"`
}

func TrimPrefix(code int32) int32 {
   return code % 10000
}

type Error struct {
   Code    int32
   Message string
}

func (err Error) Error() string {
   return fmt.Sprintf("[code=%d] %s", err.Code, err.Message)
}

func (err Error) IsCode(code int32) bool {
   if code < 10000 {
      return TrimPrefix(err.Code) == code
   }
   return err.Code == code
}

func NewError(code int32, msg string) error {
   return errors.WithStack(Error{Code: code, Message: msg})
}

func From(code int32, err error) error {
   if err != nil {
      err = NewError(code, err.Error())
   }

   return err
}

func UnGzip(dst io.ReaderFrom, src []byte) error {
   br := bytes.NewReader(src)
   zr, err := gzip.NewReader(br)
   if err != nil {
      return err
   }

   if _, err := dst.ReadFrom(zr); err != nil {
      zr.Close()
      return err
   }

   return zr.Close()
}

func UnGzipBytes(src []byte) ([]byte, error) {
   buf := bytes.Buffer{}
   err := UnGzip(&buf, src)
   return buf.Bytes(), err
}

func Gzip(dst io.Writer, src []byte) error {
   zw := gzip.NewWriter(dst)
   _, err := zw.Write(src)
   if err != nil {
      return err
   }

   return zw.Close()
}

func GzipBytes(src []byte) ([]byte, error) {
   buf := bytes.Buffer{}
   err := Gzip(&buf, src)
   return buf.Bytes(), err
}

const (
   ErrorCode = 1000
)

// 从报错的栈上看是这个函数
// 这个函数上层会有多个 MQ 消费者进行调用，但是每次传递 gzipContentMeta 参数都是通过 rpc 得到的，不会复用。
func UnmarshalSnapshot(gzipContentMeta []byte) (*Snapshot, error) {
   snapshot := new(Snapshot)
   rawSnapshot, err := UnGzipBytes(gzipContentMeta)
   if err != nil {
      return nil, From(ErrorCode, err)
   }
   err = sonic.Unmarshal(rawSnapshot, snapshot)
   if err != nil {
      return nil, From(ErrorCode, err)
   }
   return snapshot, nil
}

func TestUnmarshal(t *testing.T) {
	const N = 100
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i:=0; i<N; i++ {
		go func() {
			defer wg.Done()
			contentMeta, err := os.ReadFile("snapshot.json")
			require.Nil(t, err)
			require.Nil(t, err)
			gzipContentMeta, err := GzipBytes(contentMeta)
			require.Nil(t, err)
			_, err = UnmarshalSnapshot(gzipContentMeta)
			assert.Nil(t, err)
		}()
	}
	wg.Wait()
}