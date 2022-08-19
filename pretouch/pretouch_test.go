package pretouch

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"

	. "github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/option"
	"github.com/stretchr/testify/require"
)

var TwitterJson = func() string {
	out, err := ioutil.ReadFile("../testdata/twitter.json")
	if err != nil {
		panic(err)
	}
	return string(out)
}()

type TwitterStruct struct {
	Statuses       []Statuses     `json:"statuses"`
	SearchMetadata SearchMetadata `json:"search_metadata"`
}

type Hashtags struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
}

type Entities struct {
	Urls         []interface{} `json:"urls"`
	Hashtags     []Hashtags    `json:"hashtags"`
	UserMentions []interface{} `json:"user_mentions"`
}

type Metadata struct {
	IsoLanguageCode string `json:"iso_language_code"`
	ResultType      string `json:"result_type"`
}

type Urls struct {
	ExpandedURL interface{} `json:"expanded_url"`
	URL         string      `json:"url"`
	Indices     []int       `json:"indices"`
}

type URL struct {
	Urls []Urls `json:"urls"`
}

type Description struct {
	Urls []interface{} `json:"urls"`
}

type UserEntities struct {
	URL         URL         `json:"url"`
	Description Description `json:"description"`
}

type User struct {
	ProfileSidebarFillColor        string       `json:"profile_sidebar_fill_color"`
	ProfileSidebarBorderColor      string       `json:"profile_sidebar_border_color"`
	ProfileBackgroundTile          bool         `json:"profile_background_tile"`
	Name                           string       `json:"name"`
	ProfileImageURL                string       `json:"profile_image_url"`
	CreatedAt                      string       `json:"created_at"`
	Location                       string       `json:"location"`
	FollowRequestSent              interface{}  `json:"follow_request_sent"`
	ProfileLinkColor               string       `json:"profile_link_color"`
	IsTranslator                   bool         `json:"is_translator"`
	IDStr                          string       `json:"id_str"`
	Entities                       UserEntities `json:"entities"`
	DefaultProfile                 bool         `json:"default_profile"`
	ContributorsEnabled            bool         `json:"contributors_enabled"`
	FavouritesCount                int          `json:"favourites_count"`
	URL                            interface{}  `json:"url"`
	ProfileImageURLHTTPS           string       `json:"profile_image_url_https"`
	UtcOffset                      int          `json:"utc_offset"`
	ID                             int          `json:"id"`
	ProfileUseBackgroundImage      bool         `json:"profile_use_background_image"`
	ListedCount                    int          `json:"listed_count"`
	ProfileTextColor               string       `json:"profile_text_color"`
	Lang                           string       `json:"lang"`
	FollowersCount                 int          `json:"followers_count"`
	Protected                      bool         `json:"protected"`
	Notifications                  interface{}  `json:"notifications"`
	ProfileBackgroundImageURLHTTPS string       `json:"profile_background_image_url_https"`
	ProfileBackgroundColor         string       `json:"profile_background_color"`
	Verified                       bool         `json:"verified"`
	GeoEnabled                     bool         `json:"geo_enabled"`
	TimeZone                       string       `json:"time_zone"`
	Description                    string       `json:"description"`
	DefaultProfileImage            bool         `json:"default_profile_image"`
	ProfileBackgroundImageURL      string       `json:"profile_background_image_url"`
	StatusesCount                  int          `json:"statuses_count"`
	FriendsCount                   int          `json:"friends_count"`
	Following                      interface{}  `json:"following"`
	ShowAllInlineMedia             bool         `json:"show_all_inline_media"`
	ScreenName                     string       `json:"screen_name"`
}

type Statuses struct {
	Coordinates          interface{} `json:"coordinates"`
	Favorited            bool        `json:"favorited"`
	Truncated            bool        `json:"truncated"`
	CreatedAt            string      `json:"created_at"`
	IDStr                string      `json:"id_str"`
	Entities             Entities    `json:"entities"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	Contributors         interface{} `json:"contributors"`
	Text                 string      `json:"text"`
	Metadata             Metadata    `json:"metadata"`
	RetweetCount         int         `json:"retweet_count"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	ID                   int64       `json:"id"`
	Geo                  interface{} `json:"geo"`
	Retweeted            bool        `json:"retweeted"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	Place                interface{} `json:"place"`
	User                 User        `json:"user"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	Source               string      `json:"source"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
}

type SearchMetadata struct {
	MaxID       int64   `json:"max_id"`
	SinceID     int64   `json:"since_id"`
	RefreshURL  string  `json:"refresh_url"`
	NextResults string  `json:"next_results"`
	Count       int     `json:"count"`
	CompletedIn float64 `json:"completed_in"`
	SinceIDStr  string  `json:"since_id_str"`
	Query       string  `json:"query"`
	MaxIDStr    string  `json:"max_id_str"`
}
 

func TestPretouchTwitterStruct(t *testing.T) {
	m := new(TwitterStruct)
	s := time.Now()
	println("start decoder pretouch:", s.UnixNano())
	require.Nil(t, decoder.Pretouch(reflect.TypeOf(*m), option.WithCompileMaxInlineDepth(10)))
	e := time.Now()
	println("end decoder pretouch:", e.UnixNano())
	println("elapsed:", e.Sub(s).Milliseconds(), "ms")
	
	s = time.Now()
	println("start decode:", s.UnixNano())
	require.Nil(t, UnmarshalString(TwitterJson, m))
	e = time.Now()
	println("end decode:", e.UnixNano())
	d1 := e.Sub(s).Nanoseconds()
	println("elapsed:", d1, "ns")

	s = time.Now()
	println("start decode:", s.UnixNano())
	require.Nil(t, UnmarshalString(TwitterJson, m))
	e = time.Now()
	println("end decode:", e.UnixNano())
	d2 := e.Sub(s).Nanoseconds()
	println("elapsed:", d2, "ns")
	if d1 > d2 * 10 {
		t.Fatal("decoder pretouch not finish yet")
	}

	s = time.Now()
	println("start decode:", s.UnixNano())
	require.Nil(t, UnmarshalString(TwitterJson, m))
	e = time.Now()
	println("end decode:", e.UnixNano())
	d5 := e.Sub(s).Nanoseconds()
	println("elapsed:", d5, "ns")
	if d2 > d5 * 10 {
		t.Fatal("decoder pretouch not finish yet")
	}

	s = time.Now()
	println("start encoder pretouch:", s.UnixNano())
	require.Nil(t, encoder.Pretouch(reflect.TypeOf(m), option.WithCompileMaxInlineDepth(10)))
	e = time.Now()
	println("end encoder pretouch:", e.UnixNano())
	println("elapsed:", e.Sub(s).Milliseconds(), "ms")

	s = time.Now()
	println("start encode:", s.UnixNano())
	_, err := MarshalString(m)
	require.Nil(t, err)
	e = time.Now()
	println("end encode:", e.UnixNano())
	d3 := e.Sub(s).Nanoseconds()
	println("elapsed:", d3, "ns")
	
	s = time.Now()
	println("start encode:", s.UnixNano())
	_, err = MarshalString(m)
	require.Nil(t, err)
	e = time.Now()
	println("end encode:", e.UnixNano())
	d4 := e.Sub(s).Nanoseconds()
	println("elapsed:", d4, "ns")
	if d3 > d4 * 10 {
		t.Fatal("encoder pretouch not finish yet")
	}

	s = time.Now()
	println("start encode:", s.UnixNano())
	_, err = MarshalString(m)
	require.Nil(t, err)
	e = time.Now()
	println("end encode:", e.UnixNano())
	d6 := e.Sub(s).Nanoseconds()
	println("elapsed:", d6, "ns")
	if d4 > d6 * 10 {
		t.Fatal("encoder pretouch not finish yet")
	}
}

func BenchmarkDecodeTwitterStruct(b *testing.B) {
	m := new(TwitterStruct)
	require.Nil(b, Pretouch(reflect.TypeOf(m), option.WithCompileRecursiveDepth(10)))

	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UnmarshalString(TwitterJson, m)
	}
}

func BenchmarkEncodeTwitterStruct(b *testing.B) {
	m := new(TwitterStruct)
	require.Nil(b, Pretouch(reflect.TypeOf(m), option.WithCompileRecursiveDepth(10)))
	require.Nil(b, UnmarshalString(TwitterJson, m))

	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MarshalString(m)
	}
}