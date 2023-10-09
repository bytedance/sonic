// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f32toa = 31696
    _entry__f64toa = 176
    _entry__format_significand = 36016
    _entry__format_integer = 3376
    _entry__get_by_path = 27536
    _entry__fsm_exec = 19776
    _entry__advance_string = 16128
    _entry__advance_string_default = 37520
    _entry__do_skip_number = 22592
    _entry__skip_one_fast = 24096
    _entry__unescape = 38432
    _entry__unhex16_is = 9280
    _entry__html_escape = 9472
    _entry__i64toa = 3808
    _entry__u64toa = 4048
    _entry__lspace = 16
    _entry__quote = 5392
    _entry__skip_array = 19728
    _entry__skip_number = 23680
    _entry__skip_object = 22224
    _entry__skip_one = 23856
    _entry__unquote = 7104
    _entry__validate_one = 23920
    _entry__validate_utf8 = 30432
    _entry__validate_utf8_fast = 31104
    _entry__value = 14192
    _entry__vnumber = 17408
    _entry__atof_eisel_lemire64 = 10752
    _entry__atof_native = 13168
    _entry__decimal_to_f64 = 11232
    _entry__left_shift = 36496
    _entry__right_shift = 37040
    _entry__vsigned = 19008
    _entry__vstring = 15952
    _entry__vunsigned = 19360
)

const (
    _stack__f32toa = 56
    _stack__f64toa = 80
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__get_by_path = 296
    _stack__fsm_exec = 176
    _stack__advance_string = 72
    _stack__advance_string_default = 56
    _stack__do_skip_number = 32
    _stack__skip_one_fast = 176
    _stack__unescape = 80
    _stack__unhex16_is = 8
    _stack__html_escape = 64
    _stack__i64toa = 16
    _stack__u64toa = 8
    _stack__lspace = 8
    _stack__quote = 80
    _stack__skip_array = 184
    _stack__skip_number = 88
    _stack__skip_object = 184
    _stack__skip_one = 184
    _stack__unquote = 128
    _stack__validate_one = 184
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 24
    _stack__value = 360
    _stack__vnumber = 272
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 168
    _stack__decimal_to_f64 = 80
    _stack__left_shift = 24
    _stack__right_shift = 16
    _stack__vsigned = 16
    _stack__vstring = 128
    _stack__vunsigned = 24
)

const (
    _size__f32toa = 3664
    _size__f64toa = 3200
    _size__format_significand = 480
    _size__format_integer = 432
    _size__get_by_path = 2896
    _size__fsm_exec = 1900
    _size__advance_string = 1232
    _size__advance_string_default = 912
    _size__do_skip_number = 876
    _size__skip_one_fast = 2924
    _size__unescape = 704
    _size__unhex16_is = 128
    _size__html_escape = 1280
    _size__i64toa = 240
    _size__u64toa = 1296
    _size__lspace = 112
    _size__quote = 1696
    _size__skip_array = 48
    _size__skip_number = 160
    _size__skip_object = 48
    _size__skip_one = 48
    _size__unquote = 2176
    _size__validate_one = 48
    _size__validate_utf8 = 672
    _size__validate_utf8_fast = 544
    _size__value = 1252
    _size__vnumber = 1600
    _size__atof_eisel_lemire64 = 384
    _size__atof_native = 1024
    _size__decimal_to_f64 = 1936
    _size__left_shift = 544
    _size__right_shift = 448
    _size__vsigned = 352
    _size__vstring = 128
    _size__vunsigned = 352
)

var (
    _pcsp__f32toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {3620, 56},
        {3624, 48},
        {3625, 40},
        {3627, 32},
        {3629, 24},
        {3631, 16},
        {3633, 8},
        {3637, 0},
        {3663, 56},
    }
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {3126, 56},
        {3130, 48},
        {3131, 40},
        {3133, 32},
        {3135, 24},
        {3137, 16},
        {3139, 8},
        {3143, 0},
        {3198, 56},
    }
    _pcsp__format_significand = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {468, 24},
        {469, 16},
        {471, 8},
        {473, 0},
    }
    _pcsp__format_integer = [][2]uint32{
        {1, 0},
        {4, 8},
        {412, 16},
        {413, 8},
        {414, 0},
        {423, 16},
        {424, 8},
        {426, 0},
    }
    _pcsp__get_by_path = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2811, 120},
        {2815, 48},
        {2816, 40},
        {2818, 32},
        {2820, 24},
        {2822, 16},
        {2824, 8},
        {2825, 0},
        {2890, 120},
    }
    _pcsp__fsm_exec = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1527, 104},
        {1531, 48},
        {1532, 40},
        {1534, 32},
        {1536, 24},
        {1538, 16},
        {1540, 8},
        {1541, 0},
        {1900, 104},
    }
    _pcsp__advance_string = [][2]uint32{
        {14, 0},
        {18, 8},
        {20, 16},
        {22, 24},
        {24, 32},
        {26, 40},
        {27, 48},
        {512, 72},
        {516, 48},
        {517, 40},
        {519, 32},
        {521, 24},
        {523, 16},
        {525, 8},
        {526, 0},
        {1220, 72},
    }
    _pcsp__advance_string_default = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {381, 56},
        {385, 48},
        {386, 40},
        {388, 32},
        {390, 24},
        {392, 16},
        {394, 8},
        {395, 0},
        {907, 56},
    }
    _pcsp__do_skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {738, 32},
        {739, 24},
        {741, 16},
        {743, 8},
        {744, 0},
        {876, 32},
    }
    _pcsp__skip_one_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {593, 176},
        {594, 168},
        {596, 160},
        {598, 152},
        {600, 144},
        {602, 136},
        {606, 128},
        {2924, 176},
    }
    _pcsp__unescape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {249, 72},
        {253, 48},
        {254, 40},
        {256, 32},
        {258, 24},
        {260, 16},
        {262, 8},
        {263, 0},
        {703, 72},
    }
    _pcsp__unhex16_is = [][2]uint32{
        {1, 0},
        {32, 8},
        {33, 0},
        {59, 8},
        {60, 0},
        {95, 8},
        {96, 0},
        {120, 8},
        {122, 0},
    }
    _pcsp__html_escape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1262, 64},
        {1266, 48},
        {1267, 40},
        {1269, 32},
        {1271, 24},
        {1273, 16},
        {1275, 8},
        {1277, 0},
    }
    _pcsp__i64toa = [][2]uint32{
        {1, 0},
        {151, 8},
        {152, 0},
        {185, 8},
        {186, 0},
        {200, 8},
        {201, 0},
        {225, 8},
        {226, 0},
        {231, 8},
        {237, 0},
    }
    _pcsp__u64toa = [][2]uint32{
        {13, 0},
        {142, 8},
        {143, 0},
        {155, 8},
        {208, 0},
        {413, 8},
        {414, 0},
        {434, 8},
        {512, 0},
        {770, 8},
        {848, 0},
        {1294, 8},
        {1296, 0},
    }
    _pcsp__lspace = [][2]uint32{
        {1, 0},
        {89, 8},
        {91, 0},
    }
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1641, 80},
        {1645, 48},
        {1646, 40},
        {1648, 32},
        {1650, 24},
        {1652, 16},
        {1654, 8},
        {1655, 0},
        {1690, 80},
    }
    _pcsp__skip_array = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {107, 56},
        {111, 48},
        {112, 40},
        {114, 32},
        {116, 24},
        {118, 16},
        {120, 8},
        {121, 0},
        {145, 56},
    }
    _pcsp__skip_object = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__skip_one = [][2]uint32{
        {1, 0},
        {28, 8},
        {34, 0},
    }
    _pcsp__unquote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1604, 120},
        {1608, 48},
        {1609, 40},
        {1611, 32},
        {1613, 24},
        {1615, 16},
        {1617, 8},
        {1618, 0},
        {2176, 120},
    }
    _pcsp__validate_one = [][2]uint32{
        {1, 0},
        {33, 8},
        {39, 0},
    }
    _pcsp__validate_utf8 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {623, 48},
        {627, 40},
        {628, 32},
        {630, 24},
        {632, 16},
        {634, 8},
        {635, 0},
        {666, 48},
    }
    _pcsp__validate_utf8_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {5, 16},
        {247, 24},
        {251, 16},
        {252, 8},
        {253, 0},
        {527, 24},
        {531, 16},
        {532, 8},
        {534, 0},
    }
    _pcsp__value = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {317, 88},
        {321, 48},
        {322, 40},
        {324, 32},
        {326, 24},
        {328, 16},
        {330, 8},
        {331, 0},
        {1252, 88},
    }
    _pcsp__vnumber = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {147, 104},
        {151, 48},
        {152, 40},
        {154, 32},
        {156, 24},
        {158, 16},
        {160, 8},
        {161, 0},
        {1590, 104},
    }
    _pcsp__atof_eisel_lemire64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {302, 32},
        {303, 24},
        {305, 16},
        {307, 8},
        {308, 0},
        {372, 32},
    }
    _pcsp__atof_native = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {997, 88},
        {1001, 48},
        {1002, 40},
        {1004, 32},
        {1006, 24},
        {1008, 16},
        {1010, 8},
        {1012, 0},
    }
    _pcsp__decimal_to_f64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1902, 56},
        {1906, 48},
        {1907, 40},
        {1909, 32},
        {1911, 24},
        {1913, 16},
        {1915, 8},
        {1919, 0},
        {1931, 56},
    }
    _pcsp__left_shift = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {401, 24},
        {402, 16},
        {404, 8},
        {405, 0},
        {413, 24},
        {414, 16},
        {416, 8},
        {417, 0},
        {538, 24},
    }
    _pcsp__right_shift = [][2]uint32{
        {1, 0},
        {4, 8},
        {419, 16},
        {420, 8},
        {421, 0},
        {429, 16},
        {430, 8},
        {431, 0},
        {439, 16},
        {440, 8},
        {442, 0},
    }
    _pcsp__vsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {112, 16},
        {113, 8},
        {114, 0},
        {125, 16},
        {126, 8},
        {127, 0},
        {276, 16},
        {277, 8},
        {278, 0},
        {282, 16},
        {283, 8},
        {284, 0},
        {322, 16},
        {323, 8},
        {324, 0},
        {335, 16},
        {336, 8},
        {338, 0},
    }
    _pcsp__vstring = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {105, 56},
        {109, 40},
        {110, 32},
        {112, 24},
        {114, 16},
        {116, 8},
        {118, 0},
    }
    _pcsp__vunsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {74, 24},
        {75, 16},
        {77, 8},
        {78, 0},
        {89, 24},
        {90, 16},
        {92, 8},
        {93, 0},
        {116, 24},
        {117, 16},
        {119, 8},
        {120, 0},
        {279, 24},
        {280, 16},
        {282, 8},
        {283, 0},
        {321, 24},
        {322, 16},
        {324, 8},
        {325, 0},
        {332, 24},
        {333, 16},
        {335, 8},
        {337, 0},
    }
)

var Funcs = []loader.CFunc{
    {"__native_entry__", 0, 67, 0, nil},
    {"_f32toa", _entry__f32toa, _size__f32toa, _stack__f32toa, _pcsp__f32toa},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
    {"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
    {"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
    {"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
    {"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
    {"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
    {"_advance_string_default", _entry__advance_string_default, _size__advance_string_default, _stack__advance_string_default, _pcsp__advance_string_default},
    {"_do_skip_number", _entry__do_skip_number, _size__do_skip_number, _stack__do_skip_number, _pcsp__do_skip_number},
    {"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
    {"_unescape", _entry__unescape, _size__unescape, _stack__unescape, _pcsp__unescape},
    {"_unhex16_is", _entry__unhex16_is, _size__unhex16_is, _stack__unhex16_is, _pcsp__unhex16_is},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
    {"_lspace", _entry__lspace, _size__lspace, _stack__lspace, _pcsp__lspace},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
    {"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
    {"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
    {"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
    {"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
    {"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
    {"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
    {"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
    {"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
    {"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
    {"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
    {"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
    {"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}
