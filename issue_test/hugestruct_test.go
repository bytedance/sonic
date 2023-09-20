/*
 * Copyright 2021 ByteDance Inc.
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

package issue_test

type HugeStruct0 struct {
	Field0 map[string]*int64 `json:"field_0,omitempty"`
	Field1 *int64            `json:"field_1,omitempty"`
	Field2 []*int64          `json:"field_2,omitempty"`
	Field3 map[string]*int64 `json:"field_3,omitempty"`
	Field4 []*int64          `json:"field_4,omitempty"`
}

type HugeStruct1 struct {
	Field0   []*int32                `json:"field_0,omitempty"`
	Field1   []*string               `json:"field_1,omitempty"`
	Field2   []*int64                `json:"field_2,omitempty"`
	Field3   map[string]*int32       `json:"field_3,omitempty"`
	Field4   []*bool                 `json:"field_4,omitempty"`
	Field5   *HugeStruct0            `json:"field_5,omitempty"`
	Field6   map[string]*int32       `json:"field_6,omitempty"`
	Field7   map[string]*bool        `json:"field_7,omitempty"`
	Field8   []*bool                 `json:"field_8,omitempty"`
	Field9   map[string]*HugeStruct0 `json:"field_9,omitempty"`
	Field10  []*string               `json:"field_10,omitempty"`
	Field11  []*bool                 `json:"field_11,omitempty"`
	Field12  []*bool                 `json:"field_12,omitempty"`
	Field13  map[string]*int32       `json:"field_13,omitempty"`
	Field14  map[string]*int32       `json:"field_14,omitempty"`
	Field15  *bool                   `json:"field_15,omitempty"`
	Field16  []*int64                `json:"field_16,omitempty"`
	Field17  []*bool                 `json:"field_17,omitempty"`
	Field18  map[string]*int64       `json:"field_18,omitempty"`
	Field19  []*int64                `json:"field_19,omitempty"`
	Field20  map[string]*string      `json:"field_20,omitempty"`
	Field21  *bool                   `json:"field_21,omitempty"`
	Field22  *HugeStruct0            `json:"field_22,omitempty"`
	Field23  []*string               `json:"field_23,omitempty"`
	Field24  []*int64                `json:"field_24,omitempty"`
	Field25  []*string               `json:"field_25,omitempty"`
	Field26  []*bool                 `json:"field_26,omitempty"`
	Field27  map[string]*int32       `json:"field_27,omitempty"`
	Field28  *HugeStruct0            `json:"field_28,omitempty"`
	Field29  map[string]*int32       `json:"field_29,omitempty"`
	Field30  map[string]*bool        `json:"field_30,omitempty"`
	Field31  map[string]*int32       `json:"field_31,omitempty"`
	Field32  []*HugeStruct0          `json:"field_32,omitempty"`
	Field33  *bool                   `json:"field_33,omitempty"`
	Field34  map[string]*bool        `json:"field_34,omitempty"`
	Field35  map[string]*HugeStruct0 `json:"field_35,omitempty"`
	Field36  *HugeStruct0            `json:"field_36,omitempty"`
	Field37  *string                 `json:"field_37,omitempty"`
	Field38  []*HugeStruct0          `json:"field_38,omitempty"`
	Field39  []*bool                 `json:"field_39,omitempty"`
	Field40  map[string]*string      `json:"field_40,omitempty"`
	Field41  map[string]*int64       `json:"field_41,omitempty"`
	Field42  map[string]*int32       `json:"field_42,omitempty"`
	Field43  *string                 `json:"field_43,omitempty"`
	Field44  map[string]*HugeStruct0 `json:"field_44,omitempty"`
	Field45  map[string]*int32       `json:"field_45,omitempty"`
	Field46  *HugeStruct0            `json:"field_46,omitempty"`
	Field47  *int32                  `json:"field_47,omitempty"`
	Field48  *HugeStruct0            `json:"field_48,omitempty"`
	Field49  *int32                  `json:"field_49,omitempty"`
	Field50  map[string]*string      `json:"field_50,omitempty"`
	Field51  map[string]*bool        `json:"field_51,omitempty"`
	Field52  []*int64                `json:"field_52,omitempty"`
	Field53  map[string]*string      `json:"field_53,omitempty"`
	Field54  []*int32                `json:"field_54,omitempty"`
	Field55  map[string]*int64       `json:"field_55,omitempty"`
	Field56  map[string]*int32       `json:"field_56,omitempty"`
	Field57  map[string]*string      `json:"field_57,omitempty"`
	Field58  map[string]*HugeStruct0 `json:"field_58,omitempty"`
	Field59  []*HugeStruct0          `json:"field_59,omitempty"`
	Field60  map[string]*string      `json:"field_60,omitempty"`
	Field61  map[string]*bool        `json:"field_61,omitempty"`
	Field62  map[string]*int64       `json:"field_62,omitempty"`
	Field63  []*string               `json:"field_63,omitempty"`
	Field64  []*int64                `json:"field_64,omitempty"`
	Field65  map[string]*bool        `json:"field_65,omitempty"`
	Field66  map[string]*HugeStruct0 `json:"field_66,omitempty"`
	Field67  []*int64                `json:"field_67,omitempty"`
	Field68  map[string]*string      `json:"field_68,omitempty"`
	Field69  *HugeStruct0            `json:"field_69,omitempty"`
	Field70  []*bool                 `json:"field_70,omitempty"`
	Field71  map[string]*int64       `json:"field_71,omitempty"`
	Field72  *int32                  `json:"field_72,omitempty"`
	Field73  map[string]*int32       `json:"field_73,omitempty"`
	Field74  *int32                  `json:"field_74,omitempty"`
	Field75  map[string]*int32       `json:"field_75,omitempty"`
	Field76  map[string]*string      `json:"field_76,omitempty"`
	Field77  []*string               `json:"field_77,omitempty"`
	Field78  *string                 `json:"field_78,omitempty"`
	Field79  map[string]*int64       `json:"field_79,omitempty"`
	Field80  []*int64                `json:"field_80,omitempty"`
	Field81  map[string]*bool        `json:"field_81,omitempty"`
	Field82  []*string               `json:"field_82,omitempty"`
	Field83  []*string               `json:"field_83,omitempty"`
	Field84  *bool                   `json:"field_84,omitempty"`
	Field85  []*bool                 `json:"field_85,omitempty"`
	Field86  []*HugeStruct0          `json:"field_86,omitempty"`
	Field87  *HugeStruct0            `json:"field_87,omitempty"`
	Field88  map[string]*HugeStruct0 `json:"field_88,omitempty"`
	Field89  []*int64                `json:"field_89,omitempty"`
	Field90  []*int32                `json:"field_90,omitempty"`
	Field91  *HugeStruct0            `json:"field_91,omitempty"`
	Field92  []*bool                 `json:"field_92,omitempty"`
	Field93  []*string               `json:"field_93,omitempty"`
	Field94  map[string]*int32       `json:"field_94,omitempty"`
	Field95  *int32                  `json:"field_95,omitempty"`
	Field96  *int64                  `json:"field_96,omitempty"`
	Field97  map[string]*bool        `json:"field_97,omitempty"`
	Field98  map[string]*int32       `json:"field_98,omitempty"`
	Field99  []*HugeStruct0          `json:"field_99,omitempty"`
	Field100 *int32                  `json:"field_100,omitempty"`
	Field101 *bool                   `json:"field_101,omitempty"`
	Field102 map[string]*HugeStruct0 `json:"field_102,omitempty"`
	Field103 []*string               `json:"field_103,omitempty"`
	Field104 []*string               `json:"field_104,omitempty"`
	Field105 map[string]*bool        `json:"field_105,omitempty"`
	Field106 []*string               `json:"field_106,omitempty"`
	Field107 []*int64                `json:"field_107,omitempty"`
	Field108 *HugeStruct0            `json:"field_108,omitempty"`
	Field109 *int32                  `json:"field_109,omitempty"`
	Field110 map[string]*HugeStruct0 `json:"field_110,omitempty"`
	Field111 []*string               `json:"field_111,omitempty"`
	Field112 map[string]*HugeStruct0 `json:"field_112,omitempty"`
	Field113 []*bool                 `json:"field_113,omitempty"`
	Field114 []*bool                 `json:"field_114,omitempty"`
	Field115 map[string]*string      `json:"field_115,omitempty"`
	Field116 []*int64                `json:"field_116,omitempty"`
	Field117 []*string               `json:"field_117,omitempty"`
	Field118 map[string]*bool        `json:"field_118,omitempty"`
	Field119 map[string]*string      `json:"field_119,omitempty"`
	Field120 []*HugeStruct0          `json:"field_120,omitempty"`
	Field121 map[string]*HugeStruct0 `json:"field_121,omitempty"`
	Field122 []*bool                 `json:"field_122,omitempty"`
	Field123 *string                 `json:"field_123,omitempty"`
	Field124 []*int64                `json:"field_124,omitempty"`
	Field125 *string                 `json:"field_125,omitempty"`
	Field126 []*string               `json:"field_126,omitempty"`
	Field127 []*string               `json:"field_127,omitempty"`
	Field128 []*int32                `json:"field_128,omitempty"`
	Field129 []*bool                 `json:"field_129,omitempty"`
	Field130 *int32                  `json:"field_130,omitempty"`
	Field131 *int32                  `json:"field_131,omitempty"`
	Field132 []*int32                `json:"field_132,omitempty"`
	Field133 []*int32                `json:"field_133,omitempty"`
	Field134 *int32                  `json:"field_134,omitempty"`
	Field135 []*bool                 `json:"field_135,omitempty"`
	Field136 *bool                   `json:"field_136,omitempty"`
	Field137 []*int32                `json:"field_137,omitempty"`
	Field138 map[string]*int64       `json:"field_138,omitempty"`
	Field139 map[string]*string      `json:"field_139,omitempty"`
	Field140 map[string]*int64       `json:"field_140,omitempty"`
	Field141 map[string]*int64       `json:"field_141,omitempty"`
	Field142 []*int32                `json:"field_142,omitempty"`
	Field143 []*HugeStruct0          `json:"field_143,omitempty"`
	Field144 map[string]*int64       `json:"field_144,omitempty"`
	Field145 []*string               `json:"field_145,omitempty"`
	Field146 map[string]*int64       `json:"field_146,omitempty"`
	Field147 *int32                  `json:"field_147,omitempty"`
	Field148 map[string]*string      `json:"field_148,omitempty"`
	Field149 *int64                  `json:"field_149,omitempty"`
	Field150 map[string]*int64       `json:"field_150,omitempty"`
	Field151 map[string]*int64       `json:"field_151,omitempty"`
	Field152 map[string]*int32       `json:"field_152,omitempty"`
	Field153 []*int32                `json:"field_153,omitempty"`
	Field154 map[string]*HugeStruct0 `json:"field_154,omitempty"`
	Field155 map[string]*string      `json:"field_155,omitempty"`
	Field156 map[string]*int64       `json:"field_156,omitempty"`
	Field157 []*int32                `json:"field_157,omitempty"`
	Field158 []*int32                `json:"field_158,omitempty"`
	Field159 *int32                  `json:"field_159,omitempty"`
	Field160 *HugeStruct0            `json:"field_160,omitempty"`
	Field161 []*bool                 `json:"field_161,omitempty"`
	Field162 []*HugeStruct0          `json:"field_162,omitempty"`
	Field163 []*int32                `json:"field_163,omitempty"`
	Field164 map[string]*string      `json:"field_164,omitempty"`
	Field165 []*bool                 `json:"field_165,omitempty"`
	Field166 *HugeStruct0            `json:"field_166,omitempty"`
	Field167 *int32                  `json:"field_167,omitempty"`
	Field168 []*bool                 `json:"field_168,omitempty"`
	Field169 map[string]*bool        `json:"field_169,omitempty"`
	Field170 map[string]*bool        `json:"field_170,omitempty"`
	Field171 *HugeStruct0            `json:"field_171,omitempty"`
	Field172 map[string]*bool        `json:"field_172,omitempty"`
	Field173 []*bool                 `json:"field_173,omitempty"`
	Field174 map[string]*int64       `json:"field_174,omitempty"`
	Field175 []*HugeStruct0          `json:"field_175,omitempty"`
	Field176 []*int32                `json:"field_176,omitempty"`
	Field177 []*int64                `json:"field_177,omitempty"`
	Field178 map[string]*int64       `json:"field_178,omitempty"`
	Field179 []*int32                `json:"field_179,omitempty"`
	Field180 []*string               `json:"field_180,omitempty"`
	Field181 []*int32                `json:"field_181,omitempty"`
	Field182 map[string]*string      `json:"field_182,omitempty"`
	Field183 []*int64                `json:"field_183,omitempty"`
	Field184 *HugeStruct0            `json:"field_184,omitempty"`
	Field185 []*int32                `json:"field_185,omitempty"`
	Field186 *int32                  `json:"field_186,omitempty"`
	Field187 *HugeStruct0            `json:"field_187,omitempty"`
	Field188 []*HugeStruct0          `json:"field_188,omitempty"`
	Field189 *bool                   `json:"field_189,omitempty"`
	Field190 []*int64                `json:"field_190,omitempty"`
	Field191 map[string]*int32       `json:"field_191,omitempty"`
	Field192 []*HugeStruct0          `json:"field_192,omitempty"`
	Field193 []*HugeStruct0          `json:"field_193,omitempty"`
	Field194 *HugeStruct0            `json:"field_194,omitempty"`
	Field195 []*bool                 `json:"field_195,omitempty"`
	Field196 map[string]*bool        `json:"field_196,omitempty"`
	Field197 []*bool                 `json:"field_197,omitempty"`
	Field198 *string                 `json:"field_198,omitempty"`
	Field199 map[string]*int32       `json:"field_199,omitempty"`
	Field200 map[string]*int64       `json:"field_200,omitempty"`
	Field201 map[string]*string      `json:"field_201,omitempty"`
	Field202 map[string]*HugeStruct0 `json:"field_202,omitempty"`
	Field203 map[string]*int32       `json:"field_203,omitempty"`
	Field204 *bool                   `json:"field_204,omitempty"`
	Field205 map[string]*string      `json:"field_205,omitempty"`
	Field206 []*HugeStruct0          `json:"field_206,omitempty"`
	Field207 []*HugeStruct0          `json:"field_207,omitempty"`
	Field208 *int64                  `json:"field_208,omitempty"`
	Field209 *HugeStruct0            `json:"field_209,omitempty"`
	Field210 map[string]*string      `json:"field_210,omitempty"`
	Field211 map[string]*bool        `json:"field_211,omitempty"`
	Field212 *HugeStruct0            `json:"field_212,omitempty"`
	Field213 *bool                   `json:"field_213,omitempty"`
	Field214 map[string]*bool        `json:"field_214,omitempty"`
	Field215 map[string]*HugeStruct0 `json:"field_215,omitempty"`
	Field216 []*HugeStruct0          `json:"field_216,omitempty"`
	Field217 map[string]*string      `json:"field_217,omitempty"`
	Field218 map[string]*HugeStruct0 `json:"field_218,omitempty"`
	Field219 map[string]*int64       `json:"field_219,omitempty"`
	Field220 *int64                  `json:"field_220,omitempty"`
	Field221 *string                 `json:"field_221,omitempty"`
	Field222 *HugeStruct0            `json:"field_222,omitempty"`
	Field223 []*int64                `json:"field_223,omitempty"`
	Field224 []*bool                 `json:"field_224,omitempty"`
	Field225 []*bool                 `json:"field_225,omitempty"`
	Field226 map[string]*int64       `json:"field_226,omitempty"`
	Field227 map[string]*HugeStruct0 `json:"field_227,omitempty"`
	Field228 []*int64                `json:"field_228,omitempty"`
	Field229 map[string]*bool        `json:"field_229,omitempty"`
	Field230 map[string]*HugeStruct0 `json:"field_230,omitempty"`
	Field231 *int32                  `json:"field_231,omitempty"`
	Field232 *int32                  `json:"field_232,omitempty"`
	Field233 []*string               `json:"field_233,omitempty"`
	Field234 []*HugeStruct0          `json:"field_234,omitempty"`
	Field235 []*string               `json:"field_235,omitempty"`
	Field236 *int32                  `json:"field_236,omitempty"`
	Field237 *string                 `json:"field_237,omitempty"`
	Field238 *HugeStruct0            `json:"field_238,omitempty"`
	Field239 map[string]*HugeStruct0 `json:"field_239,omitempty"`
	Field240 []*HugeStruct0          `json:"field_240,omitempty"`
	Field241 *bool                   `json:"field_241,omitempty"`
	Field242 *int32                  `json:"field_242,omitempty"`
	Field243 map[string]*HugeStruct0 `json:"field_243,omitempty"`
	Field244 map[string]*bool        `json:"field_244,omitempty"`
	Field245 map[string]*HugeStruct0 `json:"field_245,omitempty"`
	Field246 []*int32                `json:"field_246,omitempty"`
	Field247 []*bool                 `json:"field_247,omitempty"`
	Field248 []*string               `json:"field_248,omitempty"`
	Field249 *int64                  `json:"field_249,omitempty"`
	Field250 []*int32                `json:"field_250,omitempty"`
	Field251 *HugeStruct0            `json:"field_251,omitempty"`
	Field252 *bool                   `json:"field_252,omitempty"`
	Field253 map[string]*string      `json:"field_253,omitempty"`
	Field254 map[string]*string      `json:"field_254,omitempty"`
	Field255 []*int32                `json:"field_255,omitempty"`
	Field256 *int32                  `json:"field_256,omitempty"`
	Field257 *string                 `json:"field_257,omitempty"`
	Field258 map[string]*string      `json:"field_258,omitempty"`
	Field259 map[string]*int32       `json:"field_259,omitempty"`
	Field260 []*int64                `json:"field_260,omitempty"`
	Field261 []*int32                `json:"field_261,omitempty"`
	Field262 *HugeStruct0            `json:"field_262,omitempty"`
	Field263 *bool                   `json:"field_263,omitempty"`
	Field264 *int32                  `json:"field_264,omitempty"`
	Field265 map[string]*bool        `json:"field_265,omitempty"`
	Field266 *string                 `json:"field_266,omitempty"`
	Field267 []*int64                `json:"field_267,omitempty"`
	Field268 *string                 `json:"field_268,omitempty"`
	Field269 *int64                  `json:"field_269,omitempty"`
	Field270 map[string]*int64       `json:"field_270,omitempty"`
	Field271 map[string]*int64       `json:"field_271,omitempty"`
	Field272 *HugeStruct0            `json:"field_272,omitempty"`
	Field273 []*string               `json:"field_273,omitempty"`
	Field274 *int32                  `json:"field_274,omitempty"`
	Field275 *HugeStruct0            `json:"field_275,omitempty"`
	Field276 map[string]*bool        `json:"field_276,omitempty"`
	Field277 *HugeStruct0            `json:"field_277,omitempty"`
	Field278 *int64                  `json:"field_278,omitempty"`
	Field279 map[string]*string      `json:"field_279,omitempty"`
	Field280 *string                 `json:"field_280,omitempty"`
	Field281 *int64                  `json:"field_281,omitempty"`
	Field282 *int32                  `json:"field_282,omitempty"`
	Field283 *bool                   `json:"field_283,omitempty"`
	Field284 *HugeStruct0            `json:"field_284,omitempty"`
	Field285 map[string]*int64       `json:"field_285,omitempty"`
	Field286 map[string]*bool        `json:"field_286,omitempty"`
	Field287 map[string]*string      `json:"field_287,omitempty"`
	Field288 *bool                   `json:"field_288,omitempty"`
	Field289 *bool                   `json:"field_289,omitempty"`
	Field290 *int64                  `json:"field_290,omitempty"`
	Field291 []*int64                `json:"field_291,omitempty"`
	Field292 map[string]*string      `json:"field_292,omitempty"`
	Field293 *int32                  `json:"field_293,omitempty"`
	Field294 []*string               `json:"field_294,omitempty"`
	Field295 *bool                   `json:"field_295,omitempty"`
	Field296 []*HugeStruct0          `json:"field_296,omitempty"`
	Field297 *bool                   `json:"field_297,omitempty"`
	Field298 map[string]*int64       `json:"field_298,omitempty"`
	Field299 map[string]*bool        `json:"field_299,omitempty"`
	Field300 *HugeStruct0            `json:"field_300,omitempty"`
	Field301 *bool                   `json:"field_301,omitempty"`
	Field302 []*string               `json:"field_302,omitempty"`
	Field303 []*string               `json:"field_303,omitempty"`
	Field304 map[string]*string      `json:"field_304,omitempty"`
	Field305 *int32                  `json:"field_305,omitempty"`
	Field306 *int32                  `json:"field_306,omitempty"`
	Field307 []*HugeStruct0          `json:"field_307,omitempty"`
	Field308 map[string]*HugeStruct0 `json:"field_308,omitempty"`
	Field309 map[string]*int32       `json:"field_309,omitempty"`
	Field310 []*HugeStruct0          `json:"field_310,omitempty"`
	Field311 *bool                   `json:"field_311,omitempty"`
	Field312 []*bool                 `json:"field_312,omitempty"`
	Field313 *bool                   `json:"field_313,omitempty"`
	Field314 []*HugeStruct0          `json:"field_314,omitempty"`
	Field315 *HugeStruct0            `json:"field_315,omitempty"`
	Field316 *bool                   `json:"field_316,omitempty"`
	Field317 *string                 `json:"field_317,omitempty"`
	Field318 *bool                   `json:"field_318,omitempty"`
	Field319 []*int32                `json:"field_319,omitempty"`
	Field320 *int64                  `json:"field_320,omitempty"`
	Field321 []*HugeStruct0          `json:"field_321,omitempty"`
	Field322 *bool                   `json:"field_322,omitempty"`
	Field323 *int64                  `json:"field_323,omitempty"`
	Field324 []*HugeStruct0          `json:"field_324,omitempty"`
	Field325 *bool                   `json:"field_325,omitempty"`
	Field326 []*int64                `json:"field_326,omitempty"`
	Field327 *bool                   `json:"field_327,omitempty"`
	Field328 *HugeStruct0            `json:"field_328,omitempty"`
	Field329 *HugeStruct0            `json:"field_329,omitempty"`
	Field330 []*HugeStruct0          `json:"field_330,omitempty"`
	Field331 *HugeStruct0            `json:"field_331,omitempty"`
	Field332 []*string               `json:"field_332,omitempty"`
	Field333 *int64                  `json:"field_333,omitempty"`
	Field334 []*HugeStruct0          `json:"field_334,omitempty"`
	Field335 map[string]*HugeStruct0 `json:"field_335,omitempty"`
	Field336 map[string]*bool        `json:"field_336,omitempty"`
	Field337 []*int64                `json:"field_337,omitempty"`
	Field338 map[string]*bool        `json:"field_338,omitempty"`
	Field339 *HugeStruct0            `json:"field_339,omitempty"`
	Field340 map[string]*HugeStruct0 `json:"field_340,omitempty"`
	Field341 []*bool                 `json:"field_341,omitempty"`
	Field342 []*int64                `json:"field_342,omitempty"`
	Field343 []*int32                `json:"field_343,omitempty"`
	Field344 map[string]*bool        `json:"field_344,omitempty"`
	Field345 map[string]*int64       `json:"field_345,omitempty"`
	Field346 *int64                  `json:"field_346,omitempty"`
	Field347 map[string]*bool        `json:"field_347,omitempty"`
	Field348 map[string]*int32       `json:"field_348,omitempty"`
	Field349 []*string               `json:"field_349,omitempty"`
	Field350 map[string]*int32       `json:"field_350,omitempty"`
	Field351 *bool                   `json:"field_351,omitempty"`
	Field352 []*int64                `json:"field_352,omitempty"`
	Field353 []*int64                `json:"field_353,omitempty"`
	Field354 *string                 `json:"field_354,omitempty"`
	Field355 map[string]*int32       `json:"field_355,omitempty"`
	Field356 map[string]*bool        `json:"field_356,omitempty"`
	Field357 []*int32                `json:"field_357,omitempty"`
	Field358 *int64                  `json:"field_358,omitempty"`
	Field359 map[string]*int64       `json:"field_359,omitempty"`
	Field360 *int64                  `json:"field_360,omitempty"`
	Field361 map[string]*int64       `json:"field_361,omitempty"`
	Field362 map[string]*int32       `json:"field_362,omitempty"`
	Field363 []*int64                `json:"field_363,omitempty"`
	Field364 []*bool                 `json:"field_364,omitempty"`
	Field365 *int32                  `json:"field_365,omitempty"`
	Field366 map[string]*string      `json:"field_366,omitempty"`
	Field367 map[string]*bool        `json:"field_367,omitempty"`
	Field368 *int32                  `json:"field_368,omitempty"`
	Field369 *string                 `json:"field_369,omitempty"`
	Field370 *HugeStruct0            `json:"field_370,omitempty"`
	Field371 *HugeStruct0            `json:"field_371,omitempty"`
	Field372 map[string]*HugeStruct0 `json:"field_372,omitempty"`
	Field373 map[string]*bool        `json:"field_373,omitempty"`
}

type HugeStruct2 struct {
	Field0   *bool                   `json:"field_0,omitempty"`
	Field1   map[string]*int64       `json:"field_1,omitempty"`
	Field2   *int32                  `json:"field_2,omitempty"`
	Field3   []*int64                `json:"field_3,omitempty"`
	Field4   map[string]*int32       `json:"field_4,omitempty"`
	Field5   map[string]*int32       `json:"field_5,omitempty"`
	Field6   *bool                   `json:"field_6,omitempty"`
	Field7   map[string]*int32       `json:"field_7,omitempty"`
	Field8   *int64                  `json:"field_8,omitempty"`
	Field9   []*HugeStruct1          `json:"field_9,omitempty"`
	Field10  *int64                  `json:"field_10,omitempty"`
	Field11  map[string]*int64       `json:"field_11,omitempty"`
	Field12  *string                 `json:"field_12,omitempty"`
	Field13  *int64                  `json:"field_13,omitempty"`
	Field14  map[string]*HugeStruct1 `json:"field_14,omitempty"`
	Field15  map[string]*int64       `json:"field_15,omitempty"`
	Field16  map[string]*int32       `json:"field_16,omitempty"`
	Field17  map[string]*int32       `json:"field_17,omitempty"`
	Field18  []*int32                `json:"field_18,omitempty"`
	Field19  *HugeStruct0            `json:"field_19,omitempty"`
	Field20  map[string]*int64       `json:"field_20,omitempty"`
	Field21  *HugeStruct1            `json:"field_21,omitempty"`
	Field22  []*int32                `json:"field_22,omitempty"`
	Field23  map[string]*int64       `json:"field_23,omitempty"`
	Field24  map[string]*int64       `json:"field_24,omitempty"`
	Field25  *int32                  `json:"field_25,omitempty"`
	Field26  map[string]*string      `json:"field_26,omitempty"`
	Field27  []*bool                 `json:"field_27,omitempty"`
	Field28  *int32                  `json:"field_28,omitempty"`
	Field29  []*string               `json:"field_29,omitempty"`
	Field30  []*HugeStruct0          `json:"field_30,omitempty"`
	Field31  []*int64                `json:"field_31,omitempty"`
	Field32  *int64                  `json:"field_32,omitempty"`
	Field33  map[string]*string      `json:"field_33,omitempty"`
	Field34  []*HugeStruct0          `json:"field_34,omitempty"`
	Field35  *bool                   `json:"field_35,omitempty"`
	Field36  *HugeStruct0            `json:"field_36,omitempty"`
	Field37  *string                 `json:"field_37,omitempty"`
	Field38  []*HugeStruct1          `json:"field_38,omitempty"`
	Field39  *int64                  `json:"field_39,omitempty"`
	Field40  map[string]*string      `json:"field_40,omitempty"`
	Field41  *string                 `json:"field_41,omitempty"`
	Field42  *int64                  `json:"field_42,omitempty"`
	Field43  map[string]*int64       `json:"field_43,omitempty"`
	Field44  map[string]*string      `json:"field_44,omitempty"`
	Field45  map[string]*int32       `json:"field_45,omitempty"`
	Field46  *int64                  `json:"field_46,omitempty"`
	Field47  map[string]*int64       `json:"field_47,omitempty"`
	Field48  *int32                  `json:"field_48,omitempty"`
	Field49  []*HugeStruct1          `json:"field_49,omitempty"`
	Field50  *int64                  `json:"field_50,omitempty"`
	Field51  []*int64                `json:"field_51,omitempty"`
	Field52  map[string]*int64       `json:"field_52,omitempty"`
	Field53  *int32                  `json:"field_53,omitempty"`
	Field54  map[string]*bool        `json:"field_54,omitempty"`
	Field55  map[string]*HugeStruct0 `json:"field_55,omitempty"`
	Field56  map[string]*int32       `json:"field_56,omitempty"`
	Field57  map[string]*string      `json:"field_57,omitempty"`
	Field58  []*int64                `json:"field_58,omitempty"`
	Field59  *HugeStruct0            `json:"field_59,omitempty"`
	Field60  []*int64                `json:"field_60,omitempty"`
	Field61  map[string]*int64       `json:"field_61,omitempty"`
	Field62  map[string]*HugeStruct1 `json:"field_62,omitempty"`
	Field63  *HugeStruct0            `json:"field_63,omitempty"`
	Field64  []*int32                `json:"field_64,omitempty"`
	Field65  []*HugeStruct0          `json:"field_65,omitempty"`
	Field66  *int32                  `json:"field_66,omitempty"`
	Field67  []*int64                `json:"field_67,omitempty"`
	Field68  []*bool                 `json:"field_68,omitempty"`
	Field69  *int64                  `json:"field_69,omitempty"`
	Field70  *int64                  `json:"field_70,omitempty"`
	Field71  *int64                  `json:"field_71,omitempty"`
	Field72  map[string]*int32       `json:"field_72,omitempty"`
	Field73  map[string]*int32       `json:"field_73,omitempty"`
	Field74  map[string]*int32       `json:"field_74,omitempty"`
	Field75  map[string]*bool        `json:"field_75,omitempty"`
	Field76  *string                 `json:"field_76,omitempty"`
	Field77  []*int32                `json:"field_77,omitempty"`
	Field78  *int64                  `json:"field_78,omitempty"`
	Field79  *int64                  `json:"field_79,omitempty"`
	Field80  *int64                  `json:"field_80,omitempty"`
	Field81  []*bool                 `json:"field_81,omitempty"`
	Field82  map[string]*int64       `json:"field_82,omitempty"`
	Field83  *int64                  `json:"field_83,omitempty"`
	Field84  *string                 `json:"field_84,omitempty"`
	Field85  map[string]*int32       `json:"field_85,omitempty"`
	Field86  *bool                   `json:"field_86,omitempty"`
	Field87  *HugeStruct1            `json:"field_87,omitempty"`
	Field88  []*int32                `json:"field_88,omitempty"`
	Field89  *int32                  `json:"field_89,omitempty"`
	Field90  []*bool                 `json:"field_90,omitempty"`
	Field91  []*bool                 `json:"field_91,omitempty"`
	Field92  *HugeStruct1            `json:"field_92,omitempty"`
	Field93  *int32                  `json:"field_93,omitempty"`
	Field94  *HugeStruct1            `json:"field_94,omitempty"`
	Field95  map[string]*int32       `json:"field_95,omitempty"`
	Field96  *int64                  `json:"field_96,omitempty"`
	Field97  []*HugeStruct0          `json:"field_97,omitempty"`
	Field98  []*bool                 `json:"field_98,omitempty"`
	Field99  *HugeStruct0            `json:"field_99,omitempty"`
	Field100 []*int32                `json:"field_100,omitempty"`
	Field101 *string                 `json:"field_101,omitempty"`
	Field102 map[string]*bool        `json:"field_102,omitempty"`
	Field103 map[string]*bool        `json:"field_103,omitempty"`
	Field104 []*string               `json:"field_104,omitempty"`
	Field105 map[string]*int32       `json:"field_105,omitempty"`
	Field106 *int64                  `json:"field_106,omitempty"`
	Field107 map[string]*HugeStruct1 `json:"field_107,omitempty"`
	Field108 []*int32                `json:"field_108,omitempty"`
	Field109 []*int64                `json:"field_109,omitempty"`
	Field110 *string                 `json:"field_110,omitempty"`
	Field111 map[string]*bool        `json:"field_111,omitempty"`
	Field112 []*int64                `json:"field_112,omitempty"`
	Field113 *int32                  `json:"field_113,omitempty"`
	Field114 map[string]*HugeStruct0 `json:"field_114,omitempty"`
	Field115 map[string]*int32       `json:"field_115,omitempty"`
	Field116 []*string               `json:"field_116,omitempty"`
	Field117 []*int64                `json:"field_117,omitempty"`
	Field118 []*int32                `json:"field_118,omitempty"`
	Field119 *bool                   `json:"field_119,omitempty"`
	Field120 map[string]*string      `json:"field_120,omitempty"`
	Field121 map[string]*string      `json:"field_121,omitempty"`
	Field122 []*string               `json:"field_122,omitempty"`
	Field123 map[string]*bool        `json:"field_123,omitempty"`
	Field124 map[string]*string      `json:"field_124,omitempty"`
	Field125 map[string]*int32       `json:"field_125,omitempty"`
	Field126 *HugeStruct0            `json:"field_126,omitempty"`
	Field127 *int32                  `json:"field_127,omitempty"`
	Field128 []*int64                `json:"field_128,omitempty"`
	Field129 *HugeStruct1            `json:"field_129,omitempty"`
	Field130 *string                 `json:"field_130,omitempty"`
	Field131 *HugeStruct1            `json:"field_131,omitempty"`
	Field132 []*HugeStruct0          `json:"field_132,omitempty"`
	Field133 map[string]*int64       `json:"field_133,omitempty"`
}

type HugeStruct3 struct {
	Field0   map[string]*int32       `json:"field_0,omitempty"`
	Field1   *int32                  `json:"field_1,omitempty"`
	Field2   map[string]*string      `json:"field_2,omitempty"`
	Field3   []*bool                 `json:"field_3,omitempty"`
	Field4   map[string]*string      `json:"field_4,omitempty"`
	Field5   map[string]*string      `json:"field_5,omitempty"`
	Field6   []*HugeStruct0          `json:"field_6,omitempty"`
	Field7   []*bool                 `json:"field_7,omitempty"`
	Field8   []*int32                `json:"field_8,omitempty"`
	Field9   []*bool                 `json:"field_9,omitempty"`
	Field10  map[string]*int64       `json:"field_10,omitempty"`
	Field11  *HugeStruct1            `json:"field_11,omitempty"`
	Field12  []*bool                 `json:"field_12,omitempty"`
	Field13  []*bool                 `json:"field_13,omitempty"`
	Field14  *int64                  `json:"field_14,omitempty"`
	Field15  *bool                   `json:"field_15,omitempty"`
	Field16  *int32                  `json:"field_16,omitempty"`
	Field17  *HugeStruct0            `json:"field_17,omitempty"`
	Field18  *bool                   `json:"field_18,omitempty"`
	Field19  map[string]*int32       `json:"field_19,omitempty"`
	Field20  map[string]*string      `json:"field_20,omitempty"`
	Field21  map[string]*string      `json:"field_21,omitempty"`
	Field22  *string                 `json:"field_22,omitempty"`
	Field23  []*string               `json:"field_23,omitempty"`
	Field24  []*bool                 `json:"field_24,omitempty"`
	Field25  *int32                  `json:"field_25,omitempty"`
	Field26  []*int64                `json:"field_26,omitempty"`
	Field27  *int32                  `json:"field_27,omitempty"`
	Field28  []*int32                `json:"field_28,omitempty"`
	Field29  []*int64                `json:"field_29,omitempty"`
	Field30  []*bool                 `json:"field_30,omitempty"`
	Field31  map[string]*HugeStruct1 `json:"field_31,omitempty"`
	Field32  []*bool                 `json:"field_32,omitempty"`
	Field33  map[string]*bool        `json:"field_33,omitempty"`
	Field34  []*string               `json:"field_34,omitempty"`
	Field35  []*string               `json:"field_35,omitempty"`
	Field36  []*int32                `json:"field_36,omitempty"`
	Field37  *int32                  `json:"field_37,omitempty"`
	Field38  map[string]*string      `json:"field_38,omitempty"`
	Field39  []*string               `json:"field_39,omitempty"`
	Field40  []*bool                 `json:"field_40,omitempty"`
	Field41  []*bool                 `json:"field_41,omitempty"`
	Field42  map[string]*HugeStruct1 `json:"field_42,omitempty"`
	Field43  *HugeStruct1            `json:"field_43,omitempty"`
	Field44  *bool                   `json:"field_44,omitempty"`
	Field45  []*string               `json:"field_45,omitempty"`
	Field46  map[string]*HugeStruct0 `json:"field_46,omitempty"`
	Field47  map[string]*int64       `json:"field_47,omitempty"`
	Field48  map[string]*HugeStruct2 `json:"field_48,omitempty"`
	Field49  []*bool                 `json:"field_49,omitempty"`
	Field50  []*int64                `json:"field_50,omitempty"`
	Field51  map[string]*bool        `json:"field_51,omitempty"`
	Field52  []*string               `json:"field_52,omitempty"`
	Field53  map[string]*int64       `json:"field_53,omitempty"`
	Field54  map[string]*string      `json:"field_54,omitempty"`
	Field55  map[string]*int64       `json:"field_55,omitempty"`
	Field56  *int64                  `json:"field_56,omitempty"`
	Field57  []*HugeStruct0          `json:"field_57,omitempty"`
	Field58  []*bool                 `json:"field_58,omitempty"`
	Field59  *int64                  `json:"field_59,omitempty"`
	Field60  *int32                  `json:"field_60,omitempty"`
	Field61  map[string]*int32       `json:"field_61,omitempty"`
	Field62  *bool                   `json:"field_62,omitempty"`
	Field63  map[string]*int64       `json:"field_63,omitempty"`
	Field64  map[string]*HugeStruct1 `json:"field_64,omitempty"`
	Field65  []*string               `json:"field_65,omitempty"`
	Field66  []*HugeStruct2          `json:"field_66,omitempty"`
	Field67  map[string]*bool        `json:"field_67,omitempty"`
	Field68  []*bool                 `json:"field_68,omitempty"`
	Field69  map[string]*int64       `json:"field_69,omitempty"`
	Field70  []*int64                `json:"field_70,omitempty"`
	Field71  map[string]*int32       `json:"field_71,omitempty"`
	Field72  []*int64                `json:"field_72,omitempty"`
	Field73  []*int32                `json:"field_73,omitempty"`
	Field74  []*bool                 `json:"field_74,omitempty"`
	Field75  []*int64                `json:"field_75,omitempty"`
	Field76  map[string]*int64       `json:"field_76,omitempty"`
	Field77  *string                 `json:"field_77,omitempty"`
	Field78  *bool                   `json:"field_78,omitempty"`
	Field79  []*string               `json:"field_79,omitempty"`
	Field80  map[string]*bool        `json:"field_80,omitempty"`
	Field81  map[string]*int64       `json:"field_81,omitempty"`
	Field82  []*HugeStruct2          `json:"field_82,omitempty"`
	Field83  map[string]*string      `json:"field_83,omitempty"`
	Field84  *int64                  `json:"field_84,omitempty"`
	Field85  *int64                  `json:"field_85,omitempty"`
	Field86  []*string               `json:"field_86,omitempty"`
	Field87  []*int64                `json:"field_87,omitempty"`
	Field88  []*int64                `json:"field_88,omitempty"`
	Field89  []*HugeStruct1          `json:"field_89,omitempty"`
	Field90  *int32                  `json:"field_90,omitempty"`
	Field91  map[string]*bool        `json:"field_91,omitempty"`
	Field92  *HugeStruct0            `json:"field_92,omitempty"`
	Field93  []*bool                 `json:"field_93,omitempty"`
	Field94  map[string]*string      `json:"field_94,omitempty"`
	Field95  map[string]*int64       `json:"field_95,omitempty"`
	Field96  []*HugeStruct1          `json:"field_96,omitempty"`
	Field97  []*int32                `json:"field_97,omitempty"`
	Field98  []*int64                `json:"field_98,omitempty"`
	Field99  *bool                   `json:"field_99,omitempty"`
	Field100 []*string               `json:"field_100,omitempty"`
	Field101 map[string]*int64       `json:"field_101,omitempty"`
	Field102 map[string]*string      `json:"field_102,omitempty"`
	Field103 []*int32                `json:"field_103,omitempty"`
	Field104 map[string]*string      `json:"field_104,omitempty"`
	Field105 *HugeStruct1            `json:"field_105,omitempty"`
	Field106 []*int32                `json:"field_106,omitempty"`
	Field107 *HugeStruct1            `json:"field_107,omitempty"`
	Field108 []*HugeStruct1          `json:"field_108,omitempty"`
	Field109 []*bool                 `json:"field_109,omitempty"`
	Field110 []*int32                `json:"field_110,omitempty"`
	Field111 map[string]*string      `json:"field_111,omitempty"`
	Field112 map[string]*HugeStruct0 `json:"field_112,omitempty"`
	Field113 map[string]*int32       `json:"field_113,omitempty"`
	Field114 []*bool                 `json:"field_114,omitempty"`
	Field115 []*HugeStruct2          `json:"field_115,omitempty"`
	Field116 map[string]*bool        `json:"field_116,omitempty"`
	Field117 map[string]*string      `json:"field_117,omitempty"`
	Field118 *int32                  `json:"field_118,omitempty"`
	Field119 *int64                  `json:"field_119,omitempty"`
}

type HugeStruct4 struct {
	Field0   *int64                  `json:"field_0,omitempty"`
	Field1   *string                 `json:"field_1,omitempty"`
	Field2   map[string]*int64       `json:"field_2,omitempty"`
	Field3   *HugeStruct3            `json:"field_3,omitempty"`
	Field4   []*string               `json:"field_4,omitempty"`
	Field5   map[string]*string      `json:"field_5,omitempty"`
	Field6   *HugeStruct3            `json:"field_6,omitempty"`
	Field7   map[string]*bool        `json:"field_7,omitempty"`
	Field8   map[string]*bool        `json:"field_8,omitempty"`
	Field9   []*bool                 `json:"field_9,omitempty"`
	Field10  map[string]*string      `json:"field_10,omitempty"`
	Field11  []*string               `json:"field_11,omitempty"`
	Field12  map[string]*int32       `json:"field_12,omitempty"`
	Field13  []*int64                `json:"field_13,omitempty"`
	Field14  map[string]*string      `json:"field_14,omitempty"`
	Field15  *int32                  `json:"field_15,omitempty"`
	Field16  []*int64                `json:"field_16,omitempty"`
	Field17  []*int64                `json:"field_17,omitempty"`
	Field18  map[string]*int64       `json:"field_18,omitempty"`
	Field19  *HugeStruct3            `json:"field_19,omitempty"`
	Field20  map[string]*string      `json:"field_20,omitempty"`
	Field21  []*string               `json:"field_21,omitempty"`
	Field22  []*int64                `json:"field_22,omitempty"`
	Field23  *string                 `json:"field_23,omitempty"`
	Field24  []*int64                `json:"field_24,omitempty"`
	Field25  *HugeStruct2            `json:"field_25,omitempty"`
	Field26  []*bool                 `json:"field_26,omitempty"`
	Field27  []*string               `json:"field_27,omitempty"`
	Field28  *int64                  `json:"field_28,omitempty"`
	Field29  []*bool                 `json:"field_29,omitempty"`
	Field30  map[string]*HugeStruct3 `json:"field_30,omitempty"`
	Field31  []*string               `json:"field_31,omitempty"`
	Field32  []*HugeStruct2          `json:"field_32,omitempty"`
	Field33  *int64                  `json:"field_33,omitempty"`
	Field34  map[string]*int32       `json:"field_34,omitempty"`
	Field35  map[string]*HugeStruct1 `json:"field_35,omitempty"`
	Field36  []*string               `json:"field_36,omitempty"`
	Field37  []*HugeStruct2          `json:"field_37,omitempty"`
	Field38  map[string]*int64       `json:"field_38,omitempty"`
	Field39  *string                 `json:"field_39,omitempty"`
	Field40  *HugeStruct2            `json:"field_40,omitempty"`
	Field41  []*int32                `json:"field_41,omitempty"`
	Field42  []*bool                 `json:"field_42,omitempty"`
	Field43  map[string]*bool        `json:"field_43,omitempty"`
	Field44  *HugeStruct0            `json:"field_44,omitempty"`
	Field45  []*string               `json:"field_45,omitempty"`
	Field46  []*int64                `json:"field_46,omitempty"`
	Field47  []*string               `json:"field_47,omitempty"`
	Field48  []*string               `json:"field_48,omitempty"`
	Field49  map[string]*int64       `json:"field_49,omitempty"`
	Field50  []*HugeStruct2          `json:"field_50,omitempty"`
	Field51  []*string               `json:"field_51,omitempty"`
	Field52  []*int32                `json:"field_52,omitempty"`
	Field53  *HugeStruct1            `json:"field_53,omitempty"`
	Field54  map[string]*int64       `json:"field_54,omitempty"`
	Field55  []*int32                `json:"field_55,omitempty"`
	Field56  *int32                  `json:"field_56,omitempty"`
	Field57  *int32                  `json:"field_57,omitempty"`
	Field58  []*int64                `json:"field_58,omitempty"`
	Field59  *int32                  `json:"field_59,omitempty"`
	Field60  []*HugeStruct0          `json:"field_60,omitempty"`
	Field61  *int64                  `json:"field_61,omitempty"`
	Field62  *HugeStruct3            `json:"field_62,omitempty"`
	Field63  map[string]*int64       `json:"field_63,omitempty"`
	Field64  map[string]*int32       `json:"field_64,omitempty"`
	Field65  []*int32                `json:"field_65,omitempty"`
	Field66  []*HugeStruct1          `json:"field_66,omitempty"`
	Field67  []*HugeStruct2          `json:"field_67,omitempty"`
	Field68  *HugeStruct0            `json:"field_68,omitempty"`
	Field69  *int64                  `json:"field_69,omitempty"`
	Field70  []*int64                `json:"field_70,omitempty"`
	Field71  *int64                  `json:"field_71,omitempty"`
	Field72  map[string]*int32       `json:"field_72,omitempty"`
	Field73  *bool                   `json:"field_73,omitempty"`
	Field74  []*bool                 `json:"field_74,omitempty"`
	Field75  *int32                  `json:"field_75,omitempty"`
	Field76  map[string]*int64       `json:"field_76,omitempty"`
	Field77  map[string]*int32       `json:"field_77,omitempty"`
	Field78  []*int64                `json:"field_78,omitempty"`
	Field79  *int32                  `json:"field_79,omitempty"`
	Field80  map[string]*HugeStruct2 `json:"field_80,omitempty"`
	Field81  map[string]*bool        `json:"field_81,omitempty"`
	Field82  []*HugeStruct0          `json:"field_82,omitempty"`
	Field83  *int32                  `json:"field_83,omitempty"`
	Field84  []*int64                `json:"field_84,omitempty"`
	Field85  map[string]*string      `json:"field_85,omitempty"`
	Field86  *HugeStruct0            `json:"field_86,omitempty"`
	Field87  *bool                   `json:"field_87,omitempty"`
	Field88  map[string]*int64       `json:"field_88,omitempty"`
	Field89  []*string               `json:"field_89,omitempty"`
	Field90  []*bool                 `json:"field_90,omitempty"`
	Field91  map[string]*string      `json:"field_91,omitempty"`
	Field92  *bool                   `json:"field_92,omitempty"`
	Field93  *HugeStruct2            `json:"field_93,omitempty"`
	Field94  map[string]*HugeStruct2 `json:"field_94,omitempty"`
	Field95  []*string               `json:"field_95,omitempty"`
	Field96  []*int32                `json:"field_96,omitempty"`
	Field97  *int32                  `json:"field_97,omitempty"`
	Field98  *string                 `json:"field_98,omitempty"`
	Field99  map[string]*HugeStruct3 `json:"field_99,omitempty"`
	Field100 []*HugeStruct0          `json:"field_100,omitempty"`
	Field101 *int32                  `json:"field_101,omitempty"`
	Field102 *int64                  `json:"field_102,omitempty"`
	Field103 []*HugeStruct3          `json:"field_103,omitempty"`
	Field104 map[string]*HugeStruct3 `json:"field_104,omitempty"`
	Field105 map[string]*int64       `json:"field_105,omitempty"`
	Field106 *bool                   `json:"field_106,omitempty"`
	Field107 []*string               `json:"field_107,omitempty"`
	Field108 []*HugeStruct1          `json:"field_108,omitempty"`
	Field109 *HugeStruct1            `json:"field_109,omitempty"`
	Field110 *int32                  `json:"field_110,omitempty"`
	Field111 *int64                  `json:"field_111,omitempty"`
	Field112 *string                 `json:"field_112,omitempty"`
	Field113 []*int32                `json:"field_113,omitempty"`
	Field114 map[string]*int32       `json:"field_114,omitempty"`
	Field115 *int32                  `json:"field_115,omitempty"`
	Field116 []*int64                `json:"field_116,omitempty"`
	Field117 []*bool                 `json:"field_117,omitempty"`
	Field118 []*bool                 `json:"field_118,omitempty"`
	Field119 *int64                  `json:"field_119,omitempty"`
	Field120 *int32                  `json:"field_120,omitempty"`
	Field121 []*int32                `json:"field_121,omitempty"`
	Field122 map[string]*HugeStruct3 `json:"field_122,omitempty"`
	Field123 []*int64                `json:"field_123,omitempty"`
	Field124 []*string               `json:"field_124,omitempty"`
	Field125 *HugeStruct0            `json:"field_125,omitempty"`
}

type HugeStruct5 struct {
	Field0  *string                 `json:"field_0,omitempty"`
	Field1  map[string]*bool        `json:"field_1,omitempty"`
	Field2  *bool                   `json:"field_2,omitempty"`
	Field3  map[string]*bool        `json:"field_3,omitempty"`
	Field4  *int32                  `json:"field_4,omitempty"`
	Field5  []*bool                 `json:"field_5,omitempty"`
	Field6  []*string               `json:"field_6,omitempty"`
	Field7  *bool                   `json:"field_7,omitempty"`
	Field8  map[string]*HugeStruct0 `json:"field_8,omitempty"`
	Field9  map[string]*HugeStruct0 `json:"field_9,omitempty"`
	Field10 map[string]*int32       `json:"field_10,omitempty"`
	Field11 []*int64                `json:"field_11,omitempty"`
	Field12 *string                 `json:"field_12,omitempty"`
	Field13 map[string]*HugeStruct1 `json:"field_13,omitempty"`
	Field14 *string                 `json:"field_14,omitempty"`
	Field15 *HugeStruct1            `json:"field_15,omitempty"`
	Field16 *bool                   `json:"field_16,omitempty"`
	Field17 map[string]*int32       `json:"field_17,omitempty"`
	Field18 *string                 `json:"field_18,omitempty"`
	Field19 []*HugeStruct3          `json:"field_19,omitempty"`
	Field20 map[string]*int64       `json:"field_20,omitempty"`
	Field21 map[string]*int32       `json:"field_21,omitempty"`
	Field22 *string                 `json:"field_22,omitempty"`
	Field23 map[string]*string      `json:"field_23,omitempty"`
	Field24 map[string]*string      `json:"field_24,omitempty"`
	Field25 *string                 `json:"field_25,omitempty"`
	Field26 *int64                  `json:"field_26,omitempty"`
	Field27 map[string]*int32       `json:"field_27,omitempty"`
	Field28 []*int64                `json:"field_28,omitempty"`
	Field29 []*int32                `json:"field_29,omitempty"`
	Field30 map[string]*HugeStruct1 `json:"field_30,omitempty"`
	Field31 []*bool                 `json:"field_31,omitempty"`
	Field32 *int64                  `json:"field_32,omitempty"`
	Field33 *string                 `json:"field_33,omitempty"`
	Field34 *int64                  `json:"field_34,omitempty"`
	Field35 []*int64                `json:"field_35,omitempty"`
	Field36 map[string]*string      `json:"field_36,omitempty"`
	Field37 *int32                  `json:"field_37,omitempty"`
	Field38 []*int64                `json:"field_38,omitempty"`
	Field39 map[string]*int32       `json:"field_39,omitempty"`
	Field40 map[string]*HugeStruct4 `json:"field_40,omitempty"`
	Field41 []*string               `json:"field_41,omitempty"`
	Field42 *int32                  `json:"field_42,omitempty"`
	Field43 []*bool                 `json:"field_43,omitempty"`
	Field44 []*string               `json:"field_44,omitempty"`
	Field45 *int32                  `json:"field_45,omitempty"`
	Field46 []*HugeStruct2          `json:"field_46,omitempty"`
	Field47 []*HugeStruct4          `json:"field_47,omitempty"`
	Field48 []*bool                 `json:"field_48,omitempty"`
	Field49 *bool                   `json:"field_49,omitempty"`
	Field50 []*string               `json:"field_50,omitempty"`
	Field51 map[string]*string      `json:"field_51,omitempty"`
	Field52 map[string]*bool        `json:"field_52,omitempty"`
	Field53 []*bool                 `json:"field_53,omitempty"`
	Field54 []*string               `json:"field_54,omitempty"`
	Field55 map[string]*HugeStruct0 `json:"field_55,omitempty"`
	Field56 map[string]*int64       `json:"field_56,omitempty"`
	Field57 *bool                   `json:"field_57,omitempty"`
}

type HugeStruct6 struct {
	Field0  map[string]*string `json:"field_0,omitempty"`
	Field1  *int64             `json:"field_1,omitempty"`
	Field2  *HugeStruct4       `json:"field_2,omitempty"`
	Field3  []*string          `json:"field_3,omitempty"`
	Field4  *HugeStruct5       `json:"field_4,omitempty"`
	Field5  *int32             `json:"field_5,omitempty"`
	Field6  []*int32           `json:"field_6,omitempty"`
	Field7  map[string]*int32  `json:"field_7,omitempty"`
	Field8  []*bool            `json:"field_8,omitempty"`
	Field9  *string            `json:"field_9,omitempty"`
	Field10 map[string]*bool   `json:"field_10,omitempty"`
	Field11 *int64             `json:"field_11,omitempty"`
	Field12 map[string]*bool   `json:"field_12,omitempty"`
	Field13 []*HugeStruct5     `json:"field_13,omitempty"`
}

func newIntPtr(i int64) *int64 {
	return &i
}

func GetHugeStruct0() *HugeStruct0 {
	return &HugeStruct0{
		Field0: map[string]*int64{
			"a": nil,
		},
		Field1: nil,
		Field2: []*int64{newIntPtr(1)},
		Field3: map[string]*int64{
			"a": nil,
		},
		Field4: []*int64{newIntPtr(1)},
	}
}

func GetHugeStruct1() *HugeStruct1 {
	return &HugeStruct1{
		Field0: []*int32{},
		Field1: []*string{},
		Field2: []*int64{},
		Field3: map[string]*int32{
			"": nil,
		},
		Field4: []*bool{},
		Field5: GetHugeStruct0(),
		Field6: map[string]*int32{
			"": nil,
		},
		Field7: map[string]*bool{
			"": nil,
		},
		Field8:  []*bool{},
		Field9:  map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field10: []*string{},
		Field11: []*bool{},
		Field12: []*bool{},
		Field13: map[string]*int32{
			"": nil,
		},
		Field14: map[string]*int32{
			"": nil,
		},
		Field15: nil,
		Field16: []*int64{},
		Field17: []*bool{},
		Field18: map[string]*int64{
			"": nil,
		},
		Field19: []*int64{},
		Field20: map[string]*string{
			"": nil,
		},
		Field21: nil,
		Field22: GetHugeStruct0(),
		Field23: []*string{},
		Field24: []*int64{},
		Field25: []*string{},
		Field26: []*bool{},
		Field27: map[string]*int32{
			"": nil,
		},
		Field28: GetHugeStruct0(),
		Field29: map[string]*int32{
			"": nil,
		},
		Field30: map[string]*bool{
			"": nil,
		},
		Field31: map[string]*int32{
			"": nil,
		},
		Field32: []*HugeStruct0{GetHugeStruct0()},
		Field33: nil,
		Field34: map[string]*bool{
			"": nil,
		},
		Field35: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field36: GetHugeStruct0(),
		Field37: nil,
		Field38: []*HugeStruct0{GetHugeStruct0()},
		Field39: []*bool{},
		Field40: map[string]*string{
			"": nil,
		},
		Field41: map[string]*int64{
			"": nil,
		},
		Field42: map[string]*int32{
			"": nil,
		},
		Field43: nil,
		Field44: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field45: map[string]*int32{
			"": nil,
		},
		Field46: GetHugeStruct0(),
		Field47: nil,
		Field48: GetHugeStruct0(),
		Field49: nil,
		Field50: map[string]*string{
			"": nil,
		},
		Field51: map[string]*bool{
			"": nil,
		},
		Field52: []*int64{},
		Field53: map[string]*string{
			"": nil,
		},
		Field54: []*int32{},
		Field55: map[string]*int64{
			"": nil,
		},
		Field56: map[string]*int32{
			"": nil,
		},
		Field57: map[string]*string{
			"": nil,
		},
		Field58: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field59: []*HugeStruct0{GetHugeStruct0()},
		Field60: map[string]*string{
			"": nil,
		},
		Field61: map[string]*bool{
			"": nil,
		},
		Field62: map[string]*int64{
			"": nil,
		},
		Field63: []*string{},
		Field64: []*int64{},
		Field65: map[string]*bool{
			"": nil,
		},
		Field66: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field67: []*int64{},
		Field68: map[string]*string{
			"": nil,
		},
		Field69: GetHugeStruct0(),
		Field70: []*bool{},
		Field71: map[string]*int64{
			"": nil,
		},
		Field72: nil,
		Field73: map[string]*int32{
			"": nil,
		},
		Field74: nil,
		Field75: map[string]*int32{
			"": nil,
		},
		Field76: map[string]*string{
			"": nil,
		},
		Field77: []*string{},
		Field78: nil,
		Field79: map[string]*int64{
			"": nil,
		},
		Field80: []*int64{},
		Field81: map[string]*bool{
			"": nil,
		},
		Field82: []*string{},
		Field83: []*string{},
		Field84: nil,
		Field85: []*bool{},
		Field86: []*HugeStruct0{GetHugeStruct0()},
		Field87: GetHugeStruct0(),
		Field88: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field89: []*int64{},
		Field90: []*int32{},
		Field91: GetHugeStruct0(),
		Field92: []*bool{},
		Field93: []*string{},
		Field94: map[string]*int32{
			"": nil,
		},
		Field95: nil,
		Field96: nil,
		Field97: map[string]*bool{
			"": nil,
		},
		Field98: map[string]*int32{
			"": nil,
		},
		Field99:  []*HugeStruct0{GetHugeStruct0()},
		Field100: nil,
		Field101: nil,
		Field102: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field103: []*string{},
		Field104: []*string{},
		Field105: map[string]*bool{
			"": nil,
		},
		Field106: []*string{},
		Field107: []*int64{},
		Field108: GetHugeStruct0(),
		Field109: nil,
		Field110: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field111: []*string{},
		Field112: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field113: []*bool{},
		Field114: []*bool{},
		Field115: map[string]*string{
			"": nil,
		},
		Field116: []*int64{},
		Field117: []*string{},
		Field118: map[string]*bool{
			"": nil,
		},
		Field119: map[string]*string{
			"": nil,
		},
		Field120: []*HugeStruct0{GetHugeStruct0()},
		Field121: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field122: []*bool{},
		Field123: nil,
		Field124: []*int64{},
		Field125: nil,
		Field126: []*string{},
		Field127: []*string{},
		Field128: []*int32{},
		Field129: []*bool{},
		Field130: nil,
		Field131: nil,
		Field132: []*int32{},
		Field133: []*int32{},
		Field134: nil,
		Field135: []*bool{},
		Field136: nil,
		Field137: []*int32{},
		Field138: map[string]*int64{
			"": nil,
		},
		Field139: map[string]*string{
			"": nil,
		},
		Field140: map[string]*int64{
			"": nil,
		},
		Field141: map[string]*int64{
			"": nil,
		},
		Field142: []*int32{},
		Field143: []*HugeStruct0{GetHugeStruct0()},
		Field144: map[string]*int64{
			"": nil,
		},
		Field145: []*string{},
		Field146: map[string]*int64{
			"": nil,
		},
		Field147: nil,
		Field148: map[string]*string{
			"": nil,
		},
		Field149: nil,
		Field150: map[string]*int64{
			"": nil,
		},
		Field151: map[string]*int64{
			"": nil,
		},
		Field152: map[string]*int32{
			"": nil,
		},
		Field153: []*int32{},
		Field154: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field155: map[string]*string{
			"": nil,
		},
		Field156: map[string]*int64{
			"": nil,
		},
		Field157: []*int32{},
		Field158: []*int32{},
		Field159: nil,
		Field160: GetHugeStruct0(),
		Field161: []*bool{},
		Field162: []*HugeStruct0{GetHugeStruct0()},
		Field163: []*int32{},
		Field164: map[string]*string{
			"": nil,
		},
		Field165: []*bool{},
		Field166: GetHugeStruct0(),
		Field167: nil,
		Field168: []*bool{},
		Field169: map[string]*bool{
			"": nil,
		},
		Field170: map[string]*bool{
			"": nil,
		},
		Field171: GetHugeStruct0(),
		Field172: map[string]*bool{
			"": nil,
		},
		Field173: []*bool{},
		Field174: map[string]*int64{
			"": nil,
		},
		Field175: []*HugeStruct0{GetHugeStruct0()},
		Field176: []*int32{},
		Field177: []*int64{},
		Field178: map[string]*int64{
			"": nil,
		},
		Field179: []*int32{},
		Field180: []*string{},
		Field181: []*int32{},
		Field182: map[string]*string{
			"": nil,
		},
		Field183: []*int64{},
		Field184: GetHugeStruct0(),
		Field185: []*int32{},
		Field186: nil,
		Field187: GetHugeStruct0(),
		Field188: []*HugeStruct0{GetHugeStruct0()},
		Field189: nil,
		Field190: []*int64{},
		Field191: map[string]*int32{
			"": nil,
		},
		Field192: []*HugeStruct0{GetHugeStruct0()},
		Field193: []*HugeStruct0{GetHugeStruct0()},
		Field194: GetHugeStruct0(),
		Field195: []*bool{},
		Field196: map[string]*bool{
			"": nil,
		},
		Field197: []*bool{},
		Field198: nil,
		Field199: map[string]*int32{
			"": nil,
		},
		Field200: map[string]*int64{
			"": nil,
		},
		Field201: map[string]*string{
			"": nil,
		},
		Field202: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field203: map[string]*int32{
			"": nil,
		},
		Field204: nil,
		Field205: map[string]*string{
			"": nil,
		},
		Field206: []*HugeStruct0{GetHugeStruct0()},
		Field207: []*HugeStruct0{GetHugeStruct0()},
		Field208: nil,
		Field209: GetHugeStruct0(),
		Field210: map[string]*string{
			"": nil,
		},
		Field211: map[string]*bool{
			"": nil,
		},
		Field212: GetHugeStruct0(),
		Field213: nil,
		Field214: map[string]*bool{
			"": nil,
		},
		Field215: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field216: []*HugeStruct0{GetHugeStruct0()},
		Field217: map[string]*string{
			"": nil,
		},
		Field218: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field219: map[string]*int64{
			"": nil,
		},
		Field220: nil,
		Field221: nil,
		Field222: GetHugeStruct0(),
		Field223: []*int64{},
		Field224: []*bool{},
		Field225: []*bool{},
		Field226: map[string]*int64{
			"": nil,
		},
		Field227: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field228: []*int64{},
		Field229: map[string]*bool{
			"": nil,
		},
		Field230: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field231: nil,
		Field232: nil,
		Field233: []*string{},
		Field234: []*HugeStruct0{GetHugeStruct0()},
		Field235: []*string{},
		Field236: nil,
		Field237: nil,
		Field238: GetHugeStruct0(),
		Field239: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field240: []*HugeStruct0{GetHugeStruct0()},
		Field241: nil,
		Field242: nil,
		Field243: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field244: map[string]*bool{
			"": nil,
		},
		Field245: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field246: []*int32{},
		Field247: []*bool{},
		Field248: []*string{},
		Field249: nil,
		Field250: []*int32{},
		Field251: GetHugeStruct0(),
		Field252: nil,
		Field253: map[string]*string{
			"": nil,
		},
		Field254: map[string]*string{
			"": nil,
		},
		Field255: []*int32{},
		Field256: nil,
		Field257: nil,
		Field258: map[string]*string{
			"": nil,
		},
		Field259: map[string]*int32{
			"": nil,
		},
		Field260: []*int64{},
		Field261: []*int32{},
		Field262: GetHugeStruct0(),
		Field263: nil,
		Field264: nil,
		Field265: map[string]*bool{
			"": nil,
		},
		Field266: nil,
		Field267: []*int64{},
		Field268: nil,
		Field269: nil,
		Field270: map[string]*int64{
			"": nil,
		},
		Field271: map[string]*int64{
			"": nil,
		},
		Field272: GetHugeStruct0(),
		Field273: []*string{},
		Field274: nil,
		Field275: GetHugeStruct0(),
		Field276: map[string]*bool{
			"": nil,
		},
		Field277: GetHugeStruct0(),
		Field278: nil,
		Field279: map[string]*string{
			"": nil,
		},
		Field280: nil,
		Field281: nil,
		Field282: nil,
		Field283: nil,
		Field284: GetHugeStruct0(),
		Field285: map[string]*int64{
			"": nil,
		},
		Field286: map[string]*bool{
			"": nil,
		},
		Field287: map[string]*string{
			"": nil,
		},
		Field288: nil,
		Field289: nil,
		Field290: nil,
		Field291: []*int64{},
		Field292: map[string]*string{
			"": nil,
		},
		Field293: nil,
		Field294: []*string{},
		Field295: nil,
		Field296: []*HugeStruct0{GetHugeStruct0()},
		Field297: nil,
		Field298: map[string]*int64{
			"": nil,
		},
		Field299: map[string]*bool{
			"": nil,
		},
		Field300: GetHugeStruct0(),
		Field301: nil,
		Field302: []*string{},
		Field303: []*string{},
		Field304: map[string]*string{
			"": nil,
		},
		Field305: nil,
		Field306: nil,
		Field307: []*HugeStruct0{GetHugeStruct0()},
		Field308: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field309: map[string]*int32{
			"": nil,
		},
		Field310: []*HugeStruct0{GetHugeStruct0()},
		Field311: nil,
		Field312: []*bool{},
		Field313: nil,
		Field314: []*HugeStruct0{GetHugeStruct0()},
		Field315: GetHugeStruct0(),
		Field316: nil,
		Field317: nil,
		Field318: nil,
		Field319: []*int32{},
		Field320: nil,
		Field321: []*HugeStruct0{GetHugeStruct0()},
		Field322: nil,
		Field323: nil,
		Field324: []*HugeStruct0{GetHugeStruct0()},
		Field325: nil,
		Field326: []*int64{},
		Field327: nil,
		Field328: GetHugeStruct0(),
		Field329: GetHugeStruct0(),
		Field330: []*HugeStruct0{GetHugeStruct0()},
		Field331: GetHugeStruct0(),
		Field332: []*string{},
		Field333: nil,
		Field334: []*HugeStruct0{GetHugeStruct0()},
		Field335: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field336: map[string]*bool{
			"": nil,
		},
		Field337: []*int64{},
		Field338: map[string]*bool{
			"": nil,
		},
		Field339: GetHugeStruct0(),
		Field340: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field341: []*bool{},
		Field342: []*int64{},
		Field343: []*int32{},
		Field344: map[string]*bool{
			"": nil,
		},
		Field345: map[string]*int64{
			"": nil,
		},
		Field346: nil,
		Field347: map[string]*bool{
			"": nil,
		},
		Field348: map[string]*int32{
			"": nil,
		},
		Field349: []*string{},
		Field350: map[string]*int32{
			"": nil,
		},
		Field351: nil,
		Field352: []*int64{},
		Field353: []*int64{},
		Field354: nil,
		Field355: map[string]*int32{
			"": nil,
		},
		Field356: map[string]*bool{
			"": nil,
		},
		Field357: []*int32{},
		Field358: nil,
		Field359: map[string]*int64{
			"": nil,
		},
		Field360: nil,
		Field361: map[string]*int64{
			"": nil,
		},
		Field362: map[string]*int32{
			"": nil,
		},
		Field363: []*int64{},
		Field364: []*bool{},
		Field365: nil,
		Field366: map[string]*string{
			"": nil,
		},
		Field367: map[string]*bool{
			"": nil,
		},
		Field368: nil,
		Field369: nil,
		Field370: GetHugeStruct0(),
		Field371: GetHugeStruct0(),
		Field372: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field373: map[string]*bool{
			"": nil,
		},
	}
}

func GetHugeStruct2() *HugeStruct2 {
	return &HugeStruct2{
		Field0: nil,
		Field1: map[string]*int64{
			"": nil,
		},
		Field2: nil,
		Field3: []*int64{},
		Field4: map[string]*int32{
			"": nil,
		},
		Field5: map[string]*int32{
			"": nil,
		},
		Field6: nil,
		Field7: map[string]*int32{
			"": nil,
		},
		Field8:  nil,
		Field9:  []*HugeStruct1{},
		Field10: nil,
		Field11: map[string]*int64{
			"": nil,
		},
		Field12: nil,
		Field13: nil,
		Field14: map[string]*HugeStruct1{
			"": {
				Field0: []*int32{},
				Field1: []*string{},
				Field2: []*int64{},
				Field3: map[string]*int32{
					"": nil,
				},
				Field4: []*bool{},
				Field5: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field6: map[string]*int32{
					"": nil,
				},
				Field7: map[string]*bool{
					"": nil,
				},
				Field8: []*bool{},
				Field9: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field10: []*string{},
				Field11: []*bool{},
				Field12: []*bool{},
				Field13: map[string]*int32{
					"": nil,
				},
				Field14: map[string]*int32{
					"": nil,
				},
				Field15: nil,
				Field16: []*int64{},
				Field17: []*bool{},
				Field18: map[string]*int64{
					"": nil,
				},
				Field19: []*int64{},
				Field20: map[string]*string{
					"": nil,
				},
				Field21: nil,
				Field22: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field23: []*string{},
				Field24: []*int64{},
				Field25: []*string{},
				Field26: []*bool{},
				Field27: map[string]*int32{
					"": nil,
				},
				Field28: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field29: map[string]*int32{
					"": nil,
				},
				Field30: map[string]*bool{
					"": nil,
				},
				Field31: map[string]*int32{
					"": nil,
				},
				Field32: []*HugeStruct0{GetHugeStruct0()},
				Field33: nil,
				Field34: map[string]*bool{
					"": nil,
				},
				Field35: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct0{GetHugeStruct0()},
				Field39: []*bool{},
				Field40: map[string]*string{
					"": nil,
				},
				Field41: map[string]*int64{
					"": nil,
				},
				Field42: map[string]*int32{
					"": nil,
				},
				Field43: nil,
				Field44: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field47: nil,
				Field48: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field49: nil,
				Field50: map[string]*string{
					"": nil,
				},
				Field51: map[string]*bool{
					"": nil,
				},
				Field52: []*int64{},
				Field53: map[string]*string{
					"": nil,
				},
				Field54: []*int32{},
				Field55: map[string]*int64{
					"": nil,
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field59: []*HugeStruct0{GetHugeStruct0()},
				Field60: map[string]*string{
					"": nil,
				},
				Field61: map[string]*bool{
					"": nil,
				},
				Field62: map[string]*int64{
					"": nil,
				},
				Field63: []*string{},
				Field64: []*int64{},
				Field65: map[string]*bool{
					"": nil,
				},
				Field66: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field67: []*int64{},
				Field68: map[string]*string{
					"": nil,
				},
				Field69: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field70: []*bool{},
				Field71: map[string]*int64{
					"": nil,
				},
				Field72: nil,
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: nil,
				Field75: map[string]*int32{
					"": nil,
				},
				Field76: map[string]*string{
					"": nil,
				},
				Field77: []*string{},
				Field78: nil,
				Field79: map[string]*int64{
					"": nil,
				},
				Field80: []*int64{},
				Field81: map[string]*bool{
					"": nil,
				},
				Field82: []*string{},
				Field83: []*string{},
				Field84: nil,
				Field85: []*bool{},
				Field86: []*HugeStruct0{GetHugeStruct0()},
				Field87: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field88: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field89: []*int64{},
				Field90: []*int32{},
				Field91: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field92: []*bool{},
				Field93: []*string{},
				Field94: map[string]*int32{
					"": nil,
				},
				Field95: nil,
				Field96: nil,
				Field97: map[string]*bool{
					"": nil,
				},
				Field98: map[string]*int32{
					"": nil,
				},
				Field99:  []*HugeStruct0{GetHugeStruct0()},
				Field100: nil,
				Field101: nil,
				Field102: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field103: []*string{},
				Field104: []*string{},
				Field105: map[string]*bool{
					"": nil,
				},
				Field106: []*string{},
				Field107: []*int64{},
				Field108: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field109: nil,
				Field110: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field111: []*string{},
				Field112: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field113: []*bool{},
				Field114: []*bool{},
				Field115: map[string]*string{
					"": nil,
				},
				Field116: []*int64{},
				Field117: []*string{},
				Field118: map[string]*bool{
					"": nil,
				},
				Field119: map[string]*string{
					"": nil,
				},
				Field120: []*HugeStruct0{GetHugeStruct0()},
				Field121: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field122: []*bool{},
				Field123: nil,
				Field124: []*int64{},
				Field125: nil,
				Field126: []*string{},
				Field127: []*string{},
				Field128: []*int32{},
				Field129: []*bool{},
				Field130: nil,
				Field131: nil,
				Field132: []*int32{},
				Field133: []*int32{},
				Field134: nil,
				Field135: []*bool{},
				Field136: nil,
				Field137: []*int32{},
				Field138: map[string]*int64{
					"": nil,
				},
				Field139: map[string]*string{
					"": nil,
				},
				Field140: map[string]*int64{
					"": nil,
				},
				Field141: map[string]*int64{
					"": nil,
				},
				Field142: []*int32{},
				Field143: []*HugeStruct0{GetHugeStruct0()},
				Field144: map[string]*int64{
					"": nil,
				},
				Field145: []*string{},
				Field146: map[string]*int64{
					"": nil,
				},
				Field147: nil,
				Field148: map[string]*string{
					"": nil,
				},
				Field149: nil,
				Field150: map[string]*int64{
					"": nil,
				},
				Field151: map[string]*int64{
					"": nil,
				},
				Field152: map[string]*int32{
					"": nil,
				},
				Field153: []*int32{},
				Field154: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field155: map[string]*string{
					"": nil,
				},
				Field156: map[string]*int64{
					"": nil,
				},
				Field157: []*int32{},
				Field158: []*int32{},
				Field159: nil,
				Field160: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field161: []*bool{},
				Field162: []*HugeStruct0{GetHugeStruct0()},
				Field163: []*int32{},
				Field164: map[string]*string{
					"": nil,
				},
				Field165: []*bool{},
				Field166: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field167: nil,
				Field168: []*bool{},
				Field169: map[string]*bool{
					"": nil,
				},
				Field170: map[string]*bool{
					"": nil,
				},
				Field171: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field172: map[string]*bool{
					"": nil,
				},
				Field173: []*bool{},
				Field174: map[string]*int64{
					"": nil,
				},
				Field175: []*HugeStruct0{GetHugeStruct0()},
				Field176: []*int32{},
				Field177: []*int64{},
				Field178: map[string]*int64{
					"": nil,
				},
				Field179: []*int32{},
				Field180: []*string{},
				Field181: []*int32{},
				Field182: map[string]*string{
					"": nil,
				},
				Field183: []*int64{},
				Field184: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field185: []*int32{},
				Field186: nil,
				Field187: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field188: []*HugeStruct0{GetHugeStruct0()},
				Field189: nil,
				Field190: []*int64{},
				Field191: map[string]*int32{
					"": nil,
				},
				Field192: []*HugeStruct0{GetHugeStruct0()},
				Field193: []*HugeStruct0{GetHugeStruct0()},
				Field194: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field195: []*bool{},
				Field196: map[string]*bool{
					"": nil,
				},
				Field197: []*bool{},
				Field198: nil,
				Field199: map[string]*int32{
					"": nil,
				},
				Field200: map[string]*int64{
					"": nil,
				},
				Field201: map[string]*string{
					"": nil,
				},
				Field202: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field203: map[string]*int32{
					"": nil,
				},
				Field204: nil,
				Field205: map[string]*string{
					"": nil,
				},
				Field206: []*HugeStruct0{GetHugeStruct0()},
				Field207: []*HugeStruct0{GetHugeStruct0()},
				Field208: nil,
				Field209: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field210: map[string]*string{
					"": nil,
				},
				Field211: map[string]*bool{
					"": nil,
				},
				Field212: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field213: nil,
				Field214: map[string]*bool{
					"": nil,
				},
				Field215: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field216: []*HugeStruct0{GetHugeStruct0()},
				Field217: map[string]*string{
					"": nil,
				},
				Field218: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field219: map[string]*int64{
					"": nil,
				},
				Field220: nil,
				Field221: nil,
				Field222: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field223: []*int64{},
				Field224: []*bool{},
				Field225: []*bool{},
				Field226: map[string]*int64{
					"": nil,
				},
				Field227: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field228: []*int64{},
				Field229: map[string]*bool{
					"": nil,
				},
				Field230: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field231: nil,
				Field232: nil,
				Field233: []*string{},
				Field234: []*HugeStruct0{GetHugeStruct0()},
				Field235: []*string{},
				Field236: nil,
				Field237: nil,
				Field238: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field239: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field240: []*HugeStruct0{GetHugeStruct0()},
				Field241: nil,
				Field242: nil,
				Field243: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field244: map[string]*bool{
					"": nil,
				},
				Field245: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field246: []*int32{},
				Field247: []*bool{},
				Field248: []*string{},
				Field249: nil,
				Field250: []*int32{},
				Field251: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field252: nil,
				Field253: map[string]*string{
					"": nil,
				},
				Field254: map[string]*string{
					"": nil,
				},
				Field255: []*int32{},
				Field256: nil,
				Field257: nil,
				Field258: map[string]*string{
					"": nil,
				},
				Field259: map[string]*int32{
					"": nil,
				},
				Field260: []*int64{},
				Field261: []*int32{},
				Field262: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field263: nil,
				Field264: nil,
				Field265: map[string]*bool{
					"": nil,
				},
				Field266: nil,
				Field267: []*int64{},
				Field268: nil,
				Field269: nil,
				Field270: map[string]*int64{
					"": nil,
				},
				Field271: map[string]*int64{
					"": nil,
				},
				Field272: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field273: []*string{},
				Field274: nil,
				Field275: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field276: map[string]*bool{
					"": nil,
				},
				Field277: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field278: nil,
				Field279: map[string]*string{
					"": nil,
				},
				Field280: nil,
				Field281: nil,
				Field282: nil,
				Field283: nil,
				Field284: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field285: map[string]*int64{
					"": nil,
				},
				Field286: map[string]*bool{
					"": nil,
				},
				Field287: map[string]*string{
					"": nil,
				},
				Field288: nil,
				Field289: nil,
				Field290: nil,
				Field291: []*int64{},
				Field292: map[string]*string{
					"": nil,
				},
				Field293: nil,
				Field294: []*string{},
				Field295: nil,
				Field296: []*HugeStruct0{GetHugeStruct0()},
				Field297: nil,
				Field298: map[string]*int64{
					"": nil,
				},
				Field299: map[string]*bool{
					"": nil,
				},
				Field300: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field301: nil,
				Field302: []*string{},
				Field303: []*string{},
				Field304: map[string]*string{
					"": nil,
				},
				Field305: nil,
				Field306: nil,
				Field307: []*HugeStruct0{GetHugeStruct0()},
				Field308: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field309: map[string]*int32{
					"": nil,
				},
				Field310: []*HugeStruct0{GetHugeStruct0()},
				Field311: nil,
				Field312: []*bool{},
				Field313: nil,
				Field314: []*HugeStruct0{GetHugeStruct0()},
				Field315: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field316: nil,
				Field317: nil,
				Field318: nil,
				Field319: []*int32{},
				Field320: nil,
				Field321: []*HugeStruct0{GetHugeStruct0()},
				Field322: nil,
				Field323: nil,
				Field324: []*HugeStruct0{GetHugeStruct0()},
				Field325: nil,
				Field326: []*int64{},
				Field327: nil,
				Field328: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field329: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field330: []*HugeStruct0{GetHugeStruct0()},
				Field331: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field332: []*string{},
				Field333: nil,
				Field334: []*HugeStruct0{GetHugeStruct0()},
				Field335: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field336: map[string]*bool{
					"": nil,
				},
				Field337: []*int64{},
				Field338: map[string]*bool{
					"": nil,
				},
				Field339: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field340: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field341: []*bool{},
				Field342: []*int64{},
				Field343: []*int32{},
				Field344: map[string]*bool{
					"": nil,
				},
				Field345: map[string]*int64{
					"": nil,
				},
				Field346: nil,
				Field347: map[string]*bool{
					"": nil,
				},
				Field348: map[string]*int32{
					"": nil,
				},
				Field349: []*string{},
				Field350: map[string]*int32{
					"": nil,
				},
				Field351: nil,
				Field352: []*int64{},
				Field353: []*int64{},
				Field354: nil,
				Field355: map[string]*int32{
					"": nil,
				},
				Field356: map[string]*bool{
					"": nil,
				},
				Field357: []*int32{},
				Field358: nil,
				Field359: map[string]*int64{
					"": nil,
				},
				Field360: nil,
				Field361: map[string]*int64{
					"": nil,
				},
				Field362: map[string]*int32{
					"": nil,
				},
				Field363: []*int64{},
				Field364: []*bool{},
				Field365: nil,
				Field366: map[string]*string{
					"": nil,
				},
				Field367: map[string]*bool{
					"": nil,
				},
				Field368: nil,
				Field369: nil,
				Field370: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field371: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field372: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field373: map[string]*bool{
					"": nil,
				},
			},
		},
		Field15: map[string]*int64{
			"": nil,
		},
		Field16: map[string]*int32{
			"": nil,
		},
		Field17: map[string]*int32{
			"": nil,
		},
		Field18: []*int32{},
		Field19: GetHugeStruct0(),
		Field20: map[string]*int64{
			"": nil,
		},
		Field21: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field22: []*int32{},
		Field23: map[string]*int64{
			"": nil,
		},
		Field24: map[string]*int64{
			"": nil,
		},
		Field25: nil,
		Field26: map[string]*string{
			"": nil,
		},
		Field27: []*bool{},
		Field28: nil,
		Field29: []*string{},
		Field30: []*HugeStruct0{GetHugeStruct0()},
		Field31: []*int64{},
		Field32: nil,
		Field33: map[string]*string{
			"": nil,
		},
		Field34: []*HugeStruct0{GetHugeStruct0()},
		Field35: nil,
		Field36: GetHugeStruct0(),
		Field37: nil,
		Field38: []*HugeStruct1{},
		Field39: nil,
		Field40: map[string]*string{
			"": nil,
		},
		Field41: nil,
		Field42: nil,
		Field43: map[string]*int64{
			"": nil,
		},
		Field44: map[string]*string{
			"": nil,
		},
		Field45: map[string]*int32{
			"": nil,
		},
		Field46: nil,
		Field47: map[string]*int64{
			"": nil,
		},
		Field48: nil,
		Field49: []*HugeStruct1{},
		Field50: nil,
		Field51: []*int64{},
		Field52: map[string]*int64{
			"": nil,
		},
		Field53: nil,
		Field54: map[string]*bool{
			"": nil,
		},
		Field55: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field56: map[string]*int32{
			"": nil,
		},
		Field57: map[string]*string{
			"": nil,
		},
		Field58: []*int64{},
		Field59: GetHugeStruct0(),
		Field60: []*int64{},
		Field61: map[string]*int64{
			"": nil,
		},
		Field62: map[string]*HugeStruct1{
			"": {
				Field0: []*int32{},
				Field1: []*string{},
				Field2: []*int64{},
				Field3: map[string]*int32{
					"": nil,
				},
				Field4: []*bool{},
				Field5: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field6: map[string]*int32{
					"": nil,
				},
				Field7: map[string]*bool{
					"": nil,
				},
				Field8: []*bool{},
				Field9: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field10: []*string{},
				Field11: []*bool{},
				Field12: []*bool{},
				Field13: map[string]*int32{
					"": nil,
				},
				Field14: map[string]*int32{
					"": nil,
				},
				Field15: nil,
				Field16: []*int64{},
				Field17: []*bool{},
				Field18: map[string]*int64{
					"": nil,
				},
				Field19: []*int64{},
				Field20: map[string]*string{
					"": nil,
				},
				Field21: nil,
				Field22: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field23: []*string{},
				Field24: []*int64{},
				Field25: []*string{},
				Field26: []*bool{},
				Field27: map[string]*int32{
					"": nil,
				},
				Field28: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field29: map[string]*int32{
					"": nil,
				},
				Field30: map[string]*bool{
					"": nil,
				},
				Field31: map[string]*int32{
					"": nil,
				},
				Field32: []*HugeStruct0{GetHugeStruct0()},
				Field33: nil,
				Field34: map[string]*bool{
					"": nil,
				},
				Field35: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct0{GetHugeStruct0()},
				Field39: []*bool{},
				Field40: map[string]*string{
					"": nil,
				},
				Field41: map[string]*int64{
					"": nil,
				},
				Field42: map[string]*int32{
					"": nil,
				},
				Field43: nil,
				Field44: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field47: nil,
				Field48: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field49: nil,
				Field50: map[string]*string{
					"": nil,
				},
				Field51: map[string]*bool{
					"": nil,
				},
				Field52: []*int64{},
				Field53: map[string]*string{
					"": nil,
				},
				Field54: []*int32{},
				Field55: map[string]*int64{
					"": nil,
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field59: []*HugeStruct0{GetHugeStruct0()},
				Field60: map[string]*string{
					"": nil,
				},
				Field61: map[string]*bool{
					"": nil,
				},
				Field62: map[string]*int64{
					"": nil,
				},
				Field63: []*string{},
				Field64: []*int64{},
				Field65: map[string]*bool{
					"": nil,
				},
				Field66: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field67: []*int64{},
				Field68: map[string]*string{
					"": nil,
				},
				Field69: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field70: []*bool{},
				Field71: map[string]*int64{
					"": nil,
				},
				Field72: nil,
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: nil,
				Field75: map[string]*int32{
					"": nil,
				},
				Field76: map[string]*string{
					"": nil,
				},
				Field77: []*string{},
				Field78: nil,
				Field79: map[string]*int64{
					"": nil,
				},
				Field80: []*int64{},
				Field81: map[string]*bool{
					"": nil,
				},
				Field82: []*string{},
				Field83: []*string{},
				Field84: nil,
				Field85: []*bool{},
				Field86: []*HugeStruct0{GetHugeStruct0()},
				Field87: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field88: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field89: []*int64{},
				Field90: []*int32{},
				Field91: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field92: []*bool{},
				Field93: []*string{},
				Field94: map[string]*int32{
					"": nil,
				},
				Field95: nil,
				Field96: nil,
				Field97: map[string]*bool{
					"": nil,
				},
				Field98: map[string]*int32{
					"": nil,
				},
				Field99:  []*HugeStruct0{GetHugeStruct0()},
				Field100: nil,
				Field101: nil,
				Field102: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field103: []*string{},
				Field104: []*string{},
				Field105: map[string]*bool{
					"": nil,
				},
				Field106: []*string{},
				Field107: []*int64{},
				Field108: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field109: nil,
				Field110: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field111: []*string{},
				Field112: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field113: []*bool{},
				Field114: []*bool{},
				Field115: map[string]*string{
					"": nil,
				},
				Field116: []*int64{},
				Field117: []*string{},
				Field118: map[string]*bool{
					"": nil,
				},
				Field119: map[string]*string{
					"": nil,
				},
				Field120: []*HugeStruct0{GetHugeStruct0()},
				Field121: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field122: []*bool{},
				Field123: nil,
				Field124: []*int64{},
				Field125: nil,
				Field126: []*string{},
				Field127: []*string{},
				Field128: []*int32{},
				Field129: []*bool{},
				Field130: nil,
				Field131: nil,
				Field132: []*int32{},
				Field133: []*int32{},
				Field134: nil,
				Field135: []*bool{},
				Field136: nil,
				Field137: []*int32{},
				Field138: map[string]*int64{
					"": nil,
				},
				Field139: map[string]*string{
					"": nil,
				},
				Field140: map[string]*int64{
					"": nil,
				},
				Field141: map[string]*int64{
					"": nil,
				},
				Field142: []*int32{},
				Field143: []*HugeStruct0{GetHugeStruct0()},
				Field144: map[string]*int64{
					"": nil,
				},
				Field145: []*string{},
				Field146: map[string]*int64{
					"": nil,
				},
				Field147: nil,
				Field148: map[string]*string{
					"": nil,
				},
				Field149: nil,
				Field150: map[string]*int64{
					"": nil,
				},
				Field151: map[string]*int64{
					"": nil,
				},
				Field152: map[string]*int32{
					"": nil,
				},
				Field153: []*int32{},
				Field154: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field155: map[string]*string{
					"": nil,
				},
				Field156: map[string]*int64{
					"": nil,
				},
				Field157: []*int32{},
				Field158: []*int32{},
				Field159: nil,
				Field160: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field161: []*bool{},
				Field162: []*HugeStruct0{GetHugeStruct0()},
				Field163: []*int32{},
				Field164: map[string]*string{
					"": nil,
				},
				Field165: []*bool{},
				Field166: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field167: nil,
				Field168: []*bool{},
				Field169: map[string]*bool{
					"": nil,
				},
				Field170: map[string]*bool{
					"": nil,
				},
				Field171: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field172: map[string]*bool{
					"": nil,
				},
				Field173: []*bool{},
				Field174: map[string]*int64{
					"": nil,
				},
				Field175: []*HugeStruct0{GetHugeStruct0()},
				Field176: []*int32{},
				Field177: []*int64{},
				Field178: map[string]*int64{
					"": nil,
				},
				Field179: []*int32{},
				Field180: []*string{},
				Field181: []*int32{},
				Field182: map[string]*string{
					"": nil,
				},
				Field183: []*int64{},
				Field184: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field185: []*int32{},
				Field186: nil,
				Field187: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field188: []*HugeStruct0{GetHugeStruct0()},
				Field189: nil,
				Field190: []*int64{},
				Field191: map[string]*int32{
					"": nil,
				},
				Field192: []*HugeStruct0{GetHugeStruct0()},
				Field193: []*HugeStruct0{GetHugeStruct0()},
				Field194: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field195: []*bool{},
				Field196: map[string]*bool{
					"": nil,
				},
				Field197: []*bool{},
				Field198: nil,
				Field199: map[string]*int32{
					"": nil,
				},
				Field200: map[string]*int64{
					"": nil,
				},
				Field201: map[string]*string{
					"": nil,
				},
				Field202: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field203: map[string]*int32{
					"": nil,
				},
				Field204: nil,
				Field205: map[string]*string{
					"": nil,
				},
				Field206: []*HugeStruct0{GetHugeStruct0()},
				Field207: []*HugeStruct0{GetHugeStruct0()},
				Field208: nil,
				Field209: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field210: map[string]*string{
					"": nil,
				},
				Field211: map[string]*bool{
					"": nil,
				},
				Field212: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field213: nil,
				Field214: map[string]*bool{
					"": nil,
				},
				Field215: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field216: []*HugeStruct0{GetHugeStruct0()},
				Field217: map[string]*string{
					"": nil,
				},
				Field218: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field219: map[string]*int64{
					"": nil,
				},
				Field220: nil,
				Field221: nil,
				Field222: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field223: []*int64{},
				Field224: []*bool{},
				Field225: []*bool{},
				Field226: map[string]*int64{
					"": nil,
				},
				Field227: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field228: []*int64{},
				Field229: map[string]*bool{
					"": nil,
				},
				Field230: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field231: nil,
				Field232: nil,
				Field233: []*string{},
				Field234: []*HugeStruct0{GetHugeStruct0()},
				Field235: []*string{},
				Field236: nil,
				Field237: nil,
				Field238: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field239: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field240: []*HugeStruct0{GetHugeStruct0()},
				Field241: nil,
				Field242: nil,
				Field243: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field244: map[string]*bool{
					"": nil,
				},
				Field245: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field246: []*int32{},
				Field247: []*bool{},
				Field248: []*string{},
				Field249: nil,
				Field250: []*int32{},
				Field251: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field252: nil,
				Field253: map[string]*string{
					"": nil,
				},
				Field254: map[string]*string{
					"": nil,
				},
				Field255: []*int32{},
				Field256: nil,
				Field257: nil,
				Field258: map[string]*string{
					"": nil,
				},
				Field259: map[string]*int32{
					"": nil,
				},
				Field260: []*int64{},
				Field261: []*int32{},
				Field262: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field263: nil,
				Field264: nil,
				Field265: map[string]*bool{
					"": nil,
				},
				Field266: nil,
				Field267: []*int64{},
				Field268: nil,
				Field269: nil,
				Field270: map[string]*int64{
					"": nil,
				},
				Field271: map[string]*int64{
					"": nil,
				},
				Field272: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field273: []*string{},
				Field274: nil,
				Field275: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field276: map[string]*bool{
					"": nil,
				},
				Field277: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field278: nil,
				Field279: map[string]*string{
					"": nil,
				},
				Field280: nil,
				Field281: nil,
				Field282: nil,
				Field283: nil,
				Field284: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field285: map[string]*int64{
					"": nil,
				},
				Field286: map[string]*bool{
					"": nil,
				},
				Field287: map[string]*string{
					"": nil,
				},
				Field288: nil,
				Field289: nil,
				Field290: nil,
				Field291: []*int64{},
				Field292: map[string]*string{
					"": nil,
				},
				Field293: nil,
				Field294: []*string{},
				Field295: nil,
				Field296: []*HugeStruct0{GetHugeStruct0()},
				Field297: nil,
				Field298: map[string]*int64{
					"": nil,
				},
				Field299: map[string]*bool{
					"": nil,
				},
				Field300: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field301: nil,
				Field302: []*string{},
				Field303: []*string{},
				Field304: map[string]*string{
					"": nil,
				},
				Field305: nil,
				Field306: nil,
				Field307: []*HugeStruct0{GetHugeStruct0()},
				Field308: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field309: map[string]*int32{
					"": nil,
				},
				Field310: []*HugeStruct0{GetHugeStruct0()},
				Field311: nil,
				Field312: []*bool{},
				Field313: nil,
				Field314: []*HugeStruct0{GetHugeStruct0()},
				Field315: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field316: nil,
				Field317: nil,
				Field318: nil,
				Field319: []*int32{},
				Field320: nil,
				Field321: []*HugeStruct0{GetHugeStruct0()},
				Field322: nil,
				Field323: nil,
				Field324: []*HugeStruct0{GetHugeStruct0()},
				Field325: nil,
				Field326: []*int64{},
				Field327: nil,
				Field328: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field329: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field330: []*HugeStruct0{GetHugeStruct0()},
				Field331: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field332: []*string{},
				Field333: nil,
				Field334: []*HugeStruct0{GetHugeStruct0()},
				Field335: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field336: map[string]*bool{
					"": nil,
				},
				Field337: []*int64{},
				Field338: map[string]*bool{
					"": nil,
				},
				Field339: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field340: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field341: []*bool{},
				Field342: []*int64{},
				Field343: []*int32{},
				Field344: map[string]*bool{
					"": nil,
				},
				Field345: map[string]*int64{
					"": nil,
				},
				Field346: nil,
				Field347: map[string]*bool{
					"": nil,
				},
				Field348: map[string]*int32{
					"": nil,
				},
				Field349: []*string{},
				Field350: map[string]*int32{
					"": nil,
				},
				Field351: nil,
				Field352: []*int64{},
				Field353: []*int64{},
				Field354: nil,
				Field355: map[string]*int32{
					"": nil,
				},
				Field356: map[string]*bool{
					"": nil,
				},
				Field357: []*int32{},
				Field358: nil,
				Field359: map[string]*int64{
					"": nil,
				},
				Field360: nil,
				Field361: map[string]*int64{
					"": nil,
				},
				Field362: map[string]*int32{
					"": nil,
				},
				Field363: []*int64{},
				Field364: []*bool{},
				Field365: nil,
				Field366: map[string]*string{
					"": nil,
				},
				Field367: map[string]*bool{
					"": nil,
				},
				Field368: nil,
				Field369: nil,
				Field370: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field371: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field372: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field373: map[string]*bool{
					"": nil,
				},
			},
		},
		Field63: GetHugeStruct0(),
		Field64: []*int32{},
		Field65: []*HugeStruct0{GetHugeStruct0()},
		Field66: nil,
		Field67: []*int64{},
		Field68: []*bool{},
		Field69: nil,
		Field70: nil,
		Field71: nil,
		Field72: map[string]*int32{
			"": nil,
		},
		Field73: map[string]*int32{
			"": nil,
		},
		Field74: map[string]*int32{
			"": nil,
		},
		Field75: map[string]*bool{
			"": nil,
		},
		Field76: nil,
		Field77: []*int32{},
		Field78: nil,
		Field79: nil,
		Field80: nil,
		Field81: []*bool{},
		Field82: map[string]*int64{
			"": nil,
		},
		Field83: nil,
		Field84: nil,
		Field85: map[string]*int32{
			"": nil,
		},
		Field86: nil,
		Field87: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field88: []*int32{},
		Field89: nil,
		Field90: []*bool{},
		Field91: []*bool{},
		Field92: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field93: nil,
		Field94: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field95: map[string]*int32{
			"": nil,
		},
		Field96:  nil,
		Field97:  []*HugeStruct0{GetHugeStruct0()},
		Field98:  []*bool{},
		Field99:  GetHugeStruct0(),
		Field100: []*int32{},
		Field101: nil,
		Field102: map[string]*bool{
			"": nil,
		},
		Field103: map[string]*bool{
			"": nil,
		},
		Field104: []*string{},
		Field105: map[string]*int32{
			"": nil,
		},
		Field106: nil,
		Field107: map[string]*HugeStruct1{
			"": {
				Field0: []*int32{},
				Field1: []*string{},
				Field2: []*int64{},
				Field3: map[string]*int32{
					"": nil,
				},
				Field4: []*bool{},
				Field5: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field6: map[string]*int32{
					"": nil,
				},
				Field7: map[string]*bool{
					"": nil,
				},
				Field8: []*bool{},
				Field9: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field10: []*string{},
				Field11: []*bool{},
				Field12: []*bool{},
				Field13: map[string]*int32{
					"": nil,
				},
				Field14: map[string]*int32{
					"": nil,
				},
				Field15: nil,
				Field16: []*int64{},
				Field17: []*bool{},
				Field18: map[string]*int64{
					"": nil,
				},
				Field19: []*int64{},
				Field20: map[string]*string{
					"": nil,
				},
				Field21: nil,
				Field22: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field23: []*string{},
				Field24: []*int64{},
				Field25: []*string{},
				Field26: []*bool{},
				Field27: map[string]*int32{
					"": nil,
				},
				Field28: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field29: map[string]*int32{
					"": nil,
				},
				Field30: map[string]*bool{
					"": nil,
				},
				Field31: map[string]*int32{
					"": nil,
				},
				Field32: []*HugeStruct0{GetHugeStruct0()},
				Field33: nil,
				Field34: map[string]*bool{
					"": nil,
				},
				Field35: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct0{GetHugeStruct0()},
				Field39: []*bool{},
				Field40: map[string]*string{
					"": nil,
				},
				Field41: map[string]*int64{
					"": nil,
				},
				Field42: map[string]*int32{
					"": nil,
				},
				Field43: nil,
				Field44: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field47: nil,
				Field48: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field49: nil,
				Field50: map[string]*string{
					"": nil,
				},
				Field51: map[string]*bool{
					"": nil,
				},
				Field52: []*int64{},
				Field53: map[string]*string{
					"": nil,
				},
				Field54: []*int32{},
				Field55: map[string]*int64{
					"": nil,
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field59: []*HugeStruct0{GetHugeStruct0()},
				Field60: map[string]*string{
					"": nil,
				},
				Field61: map[string]*bool{
					"": nil,
				},
				Field62: map[string]*int64{
					"": nil,
				},
				Field63: []*string{},
				Field64: []*int64{},
				Field65: map[string]*bool{
					"": nil,
				},
				Field66: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field67: []*int64{},
				Field68: map[string]*string{
					"": nil,
				},
				Field69: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field70: []*bool{},
				Field71: map[string]*int64{
					"": nil,
				},
				Field72: nil,
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: nil,
				Field75: map[string]*int32{
					"": nil,
				},
				Field76: map[string]*string{
					"": nil,
				},
				Field77: []*string{},
				Field78: nil,
				Field79: map[string]*int64{
					"": nil,
				},
				Field80: []*int64{},
				Field81: map[string]*bool{
					"": nil,
				},
				Field82: []*string{},
				Field83: []*string{},
				Field84: nil,
				Field85: []*bool{},
				Field86: []*HugeStruct0{GetHugeStruct0()},
				Field87: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field88: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field89: []*int64{},
				Field90: []*int32{},
				Field91: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field92: []*bool{},
				Field93: []*string{},
				Field94: map[string]*int32{
					"": nil,
				},
				Field95: nil,
				Field96: nil,
				Field97: map[string]*bool{
					"": nil,
				},
				Field98: map[string]*int32{
					"": nil,
				},
				Field99:  []*HugeStruct0{GetHugeStruct0()},
				Field100: nil,
				Field101: nil,
				Field102: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field103: []*string{},
				Field104: []*string{},
				Field105: map[string]*bool{
					"": nil,
				},
				Field106: []*string{},
				Field107: []*int64{},
				Field108: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field109: nil,
				Field110: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field111: []*string{},
				Field112: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field113: []*bool{},
				Field114: []*bool{},
				Field115: map[string]*string{
					"": nil,
				},
				Field116: []*int64{},
				Field117: []*string{},
				Field118: map[string]*bool{
					"": nil,
				},
				Field119: map[string]*string{
					"": nil,
				},
				Field120: []*HugeStruct0{GetHugeStruct0()},
				Field121: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field122: []*bool{},
				Field123: nil,
				Field124: []*int64{},
				Field125: nil,
				Field126: []*string{},
				Field127: []*string{},
				Field128: []*int32{},
				Field129: []*bool{},
				Field130: nil,
				Field131: nil,
				Field132: []*int32{},
				Field133: []*int32{},
				Field134: nil,
				Field135: []*bool{},
				Field136: nil,
				Field137: []*int32{},
				Field138: map[string]*int64{
					"": nil,
				},
				Field139: map[string]*string{
					"": nil,
				},
				Field140: map[string]*int64{
					"": nil,
				},
				Field141: map[string]*int64{
					"": nil,
				},
				Field142: []*int32{},
				Field143: []*HugeStruct0{GetHugeStruct0()},
				Field144: map[string]*int64{
					"": nil,
				},
				Field145: []*string{},
				Field146: map[string]*int64{
					"": nil,
				},
				Field147: nil,
				Field148: map[string]*string{
					"": nil,
				},
				Field149: nil,
				Field150: map[string]*int64{
					"": nil,
				},
				Field151: map[string]*int64{
					"": nil,
				},
				Field152: map[string]*int32{
					"": nil,
				},
				Field153: []*int32{},
				Field154: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field155: map[string]*string{
					"": nil,
				},
				Field156: map[string]*int64{
					"": nil,
				},
				Field157: []*int32{},
				Field158: []*int32{},
				Field159: nil,
				Field160: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field161: []*bool{},
				Field162: []*HugeStruct0{GetHugeStruct0()},
				Field163: []*int32{},
				Field164: map[string]*string{
					"": nil,
				},
				Field165: []*bool{},
				Field166: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field167: nil,
				Field168: []*bool{},
				Field169: map[string]*bool{
					"": nil,
				},
				Field170: map[string]*bool{
					"": nil,
				},
				Field171: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field172: map[string]*bool{
					"": nil,
				},
				Field173: []*bool{},
				Field174: map[string]*int64{
					"": nil,
				},
				Field175: []*HugeStruct0{GetHugeStruct0()},
				Field176: []*int32{},
				Field177: []*int64{},
				Field178: map[string]*int64{
					"": nil,
				},
				Field179: []*int32{},
				Field180: []*string{},
				Field181: []*int32{},
				Field182: map[string]*string{
					"": nil,
				},
				Field183: []*int64{},
				Field184: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field185: []*int32{},
				Field186: nil,
				Field187: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field188: []*HugeStruct0{GetHugeStruct0()},
				Field189: nil,
				Field190: []*int64{},
				Field191: map[string]*int32{
					"": nil,
				},
				Field192: []*HugeStruct0{GetHugeStruct0()},
				Field193: []*HugeStruct0{GetHugeStruct0()},
				Field194: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field195: []*bool{},
				Field196: map[string]*bool{
					"": nil,
				},
				Field197: []*bool{},
				Field198: nil,
				Field199: map[string]*int32{
					"": nil,
				},
				Field200: map[string]*int64{
					"": nil,
				},
				Field201: map[string]*string{
					"": nil,
				},
				Field202: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field203: map[string]*int32{
					"": nil,
				},
				Field204: nil,
				Field205: map[string]*string{
					"": nil,
				},
				Field206: []*HugeStruct0{GetHugeStruct0()},
				Field207: []*HugeStruct0{GetHugeStruct0()},
				Field208: nil,
				Field209: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field210: map[string]*string{
					"": nil,
				},
				Field211: map[string]*bool{
					"": nil,
				},
				Field212: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field213: nil,
				Field214: map[string]*bool{
					"": nil,
				},
				Field215: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field216: []*HugeStruct0{GetHugeStruct0()},
				Field217: map[string]*string{
					"": nil,
				},
				Field218: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field219: map[string]*int64{
					"": nil,
				},
				Field220: nil,
				Field221: nil,
				Field222: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field223: []*int64{},
				Field224: []*bool{},
				Field225: []*bool{},
				Field226: map[string]*int64{
					"": nil,
				},
				Field227: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field228: []*int64{},
				Field229: map[string]*bool{
					"": nil,
				},
				Field230: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field231: nil,
				Field232: nil,
				Field233: []*string{},
				Field234: []*HugeStruct0{GetHugeStruct0()},
				Field235: []*string{},
				Field236: nil,
				Field237: nil,
				Field238: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field239: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field240: []*HugeStruct0{GetHugeStruct0()},
				Field241: nil,
				Field242: nil,
				Field243: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field244: map[string]*bool{
					"": nil,
				},
				Field245: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field246: []*int32{},
				Field247: []*bool{},
				Field248: []*string{},
				Field249: nil,
				Field250: []*int32{},
				Field251: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field252: nil,
				Field253: map[string]*string{
					"": nil,
				},
				Field254: map[string]*string{
					"": nil,
				},
				Field255: []*int32{},
				Field256: nil,
				Field257: nil,
				Field258: map[string]*string{
					"": nil,
				},
				Field259: map[string]*int32{
					"": nil,
				},
				Field260: []*int64{},
				Field261: []*int32{},
				Field262: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field263: nil,
				Field264: nil,
				Field265: map[string]*bool{
					"": nil,
				},
				Field266: nil,
				Field267: []*int64{},
				Field268: nil,
				Field269: nil,
				Field270: map[string]*int64{
					"": nil,
				},
				Field271: map[string]*int64{
					"": nil,
				},
				Field272: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field273: []*string{},
				Field274: nil,
				Field275: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field276: map[string]*bool{
					"": nil,
				},
				Field277: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field278: nil,
				Field279: map[string]*string{
					"": nil,
				},
				Field280: nil,
				Field281: nil,
				Field282: nil,
				Field283: nil,
				Field284: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field285: map[string]*int64{
					"": nil,
				},
				Field286: map[string]*bool{
					"": nil,
				},
				Field287: map[string]*string{
					"": nil,
				},
				Field288: nil,
				Field289: nil,
				Field290: nil,
				Field291: []*int64{},
				Field292: map[string]*string{
					"": nil,
				},
				Field293: nil,
				Field294: []*string{},
				Field295: nil,
				Field296: []*HugeStruct0{GetHugeStruct0()},
				Field297: nil,
				Field298: map[string]*int64{
					"": nil,
				},
				Field299: map[string]*bool{
					"": nil,
				},
				Field300: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field301: nil,
				Field302: []*string{},
				Field303: []*string{},
				Field304: map[string]*string{
					"": nil,
				},
				Field305: nil,
				Field306: nil,
				Field307: []*HugeStruct0{GetHugeStruct0()},
				Field308: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field309: map[string]*int32{
					"": nil,
				},
				Field310: []*HugeStruct0{GetHugeStruct0()},
				Field311: nil,
				Field312: []*bool{},
				Field313: nil,
				Field314: []*HugeStruct0{GetHugeStruct0()},
				Field315: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field316: nil,
				Field317: nil,
				Field318: nil,
				Field319: []*int32{},
				Field320: nil,
				Field321: []*HugeStruct0{GetHugeStruct0()},
				Field322: nil,
				Field323: nil,
				Field324: []*HugeStruct0{GetHugeStruct0()},
				Field325: nil,
				Field326: []*int64{},
				Field327: nil,
				Field328: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field329: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field330: []*HugeStruct0{GetHugeStruct0()},
				Field331: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field332: []*string{},
				Field333: nil,
				Field334: []*HugeStruct0{GetHugeStruct0()},
				Field335: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field336: map[string]*bool{
					"": nil,
				},
				Field337: []*int64{},
				Field338: map[string]*bool{
					"": nil,
				},
				Field339: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field340: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field341: []*bool{},
				Field342: []*int64{},
				Field343: []*int32{},
				Field344: map[string]*bool{
					"": nil,
				},
				Field345: map[string]*int64{
					"": nil,
				},
				Field346: nil,
				Field347: map[string]*bool{
					"": nil,
				},
				Field348: map[string]*int32{
					"": nil,
				},
				Field349: []*string{},
				Field350: map[string]*int32{
					"": nil,
				},
				Field351: nil,
				Field352: []*int64{},
				Field353: []*int64{},
				Field354: nil,
				Field355: map[string]*int32{
					"": nil,
				},
				Field356: map[string]*bool{
					"": nil,
				},
				Field357: []*int32{},
				Field358: nil,
				Field359: map[string]*int64{
					"": nil,
				},
				Field360: nil,
				Field361: map[string]*int64{
					"": nil,
				},
				Field362: map[string]*int32{
					"": nil,
				},
				Field363: []*int64{},
				Field364: []*bool{},
				Field365: nil,
				Field366: map[string]*string{
					"": nil,
				},
				Field367: map[string]*bool{
					"": nil,
				},
				Field368: nil,
				Field369: nil,
				Field370: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field371: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field372: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field373: map[string]*bool{
					"": nil,
				},
			},
		},
		Field108: []*int32{},
		Field109: []*int64{},
		Field110: nil,
		Field111: map[string]*bool{
			"": nil,
		},
		Field112: []*int64{},
		Field113: nil,
		Field114: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field115: map[string]*int32{
			"": nil,
		},
		Field116: []*string{},
		Field117: []*int64{},
		Field118: []*int32{},
		Field119: nil,
		Field120: map[string]*string{
			"": nil,
		},
		Field121: map[string]*string{
			"": nil,
		},
		Field122: []*string{},
		Field123: map[string]*bool{
			"": nil,
		},
		Field124: map[string]*string{
			"": nil,
		},
		Field125: map[string]*int32{
			"": nil,
		},
		Field126: GetHugeStruct0(),
		Field127: nil,
		Field128: []*int64{},
		Field129: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field130: nil,
		Field131: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field132: []*HugeStruct0{GetHugeStruct0()},
		Field133: map[string]*int64{
			"": nil,
		},
	}
}

func GetHugeStruc3() *HugeStruct3 {
	return &HugeStruct3{
		Field0: map[string]*int32{
			"": nil,
		},
		Field1: nil,
		Field2: map[string]*string{
			"": nil,
		},
		Field3: []*bool{},
		Field4: map[string]*string{
			"": nil,
		},
		Field5: map[string]*string{
			"": nil,
		},
		Field6: []*HugeStruct0{GetHugeStruct0()},
		Field7: []*bool{},
		Field8: []*int32{},
		Field9: []*bool{},
		Field10: map[string]*int64{
			"": nil,
		},
		Field11: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field12: []*bool{},
		Field13: []*bool{},
		Field14: nil,
		Field15: nil,
		Field16: nil,
		Field17: GetHugeStruct0(),
		Field18: nil,
		Field19: map[string]*int32{
			"": nil,
		},
		Field20: map[string]*string{
			"": nil,
		},
		Field21: map[string]*string{
			"": nil,
		},
		Field22: nil,
		Field23: []*string{},
		Field24: []*bool{},
		Field25: nil,
		Field26: []*int64{},
		Field27: nil,
		Field28: []*int32{},
		Field29: []*int64{},
		Field30: []*bool{},
		Field31: map[string]*HugeStruct1{
			"": {
				Field0: []*int32{},
				Field1: []*string{},
				Field2: []*int64{},
				Field3: map[string]*int32{
					"": nil,
				},
				Field4: []*bool{},
				Field5: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field6: map[string]*int32{
					"": nil,
				},
				Field7: map[string]*bool{
					"": nil,
				},
				Field8: []*bool{},
				Field9: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field10: []*string{},
				Field11: []*bool{},
				Field12: []*bool{},
				Field13: map[string]*int32{
					"": nil,
				},
				Field14: map[string]*int32{
					"": nil,
				},
				Field15: nil,
				Field16: []*int64{},
				Field17: []*bool{},
				Field18: map[string]*int64{
					"": nil,
				},
				Field19: []*int64{},
				Field20: map[string]*string{
					"": nil,
				},
				Field21: nil,
				Field22: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field23: []*string{},
				Field24: []*int64{},
				Field25: []*string{},
				Field26: []*bool{},
				Field27: map[string]*int32{
					"": nil,
				},
				Field28: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field29: map[string]*int32{
					"": nil,
				},
				Field30: map[string]*bool{
					"": nil,
				},
				Field31: map[string]*int32{
					"": nil,
				},
				Field32: []*HugeStruct0{GetHugeStruct0()},
				Field33: nil,
				Field34: map[string]*bool{
					"": nil,
				},
				Field35: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct0{GetHugeStruct0()},
				Field39: []*bool{},
				Field40: map[string]*string{
					"": nil,
				},
				Field41: map[string]*int64{
					"": nil,
				},
				Field42: map[string]*int32{
					"": nil,
				},
				Field43: nil,
				Field44: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field47: nil,
				Field48: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field49: nil,
				Field50: map[string]*string{
					"": nil,
				},
				Field51: map[string]*bool{
					"": nil,
				},
				Field52: []*int64{},
				Field53: map[string]*string{
					"": nil,
				},
				Field54: []*int32{},
				Field55: map[string]*int64{
					"": nil,
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field59: []*HugeStruct0{GetHugeStruct0()},
				Field60: map[string]*string{
					"": nil,
				},
				Field61: map[string]*bool{
					"": nil,
				},
				Field62: map[string]*int64{
					"": nil,
				},
				Field63: []*string{},
				Field64: []*int64{},
				Field65: map[string]*bool{
					"": nil,
				},
				Field66: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field67: []*int64{},
				Field68: map[string]*string{
					"": nil,
				},
				Field69: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field70: []*bool{},
				Field71: map[string]*int64{
					"": nil,
				},
				Field72: nil,
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: nil,
				Field75: map[string]*int32{
					"": nil,
				},
				Field76: map[string]*string{
					"": nil,
				},
				Field77: []*string{},
				Field78: nil,
				Field79: map[string]*int64{
					"": nil,
				},
				Field80: []*int64{},
				Field81: map[string]*bool{
					"": nil,
				},
				Field82: []*string{},
				Field83: []*string{},
				Field84: nil,
				Field85: []*bool{},
				Field86: []*HugeStruct0{GetHugeStruct0()},
				Field87: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field88: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field89: []*int64{},
				Field90: []*int32{},
				Field91: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field92: []*bool{},
				Field93: []*string{},
				Field94: map[string]*int32{
					"": nil,
				},
				Field95: nil,
				Field96: nil,
				Field97: map[string]*bool{
					"": nil,
				},
				Field98: map[string]*int32{
					"": nil,
				},
				Field99:  []*HugeStruct0{GetHugeStruct0()},
				Field100: nil,
				Field101: nil,
				Field102: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field103: []*string{},
				Field104: []*string{},
				Field105: map[string]*bool{
					"": nil,
				},
				Field106: []*string{},
				Field107: []*int64{},
				Field108: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field109: nil,
				Field110: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field111: []*string{},
				Field112: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field113: []*bool{},
				Field114: []*bool{},
				Field115: map[string]*string{
					"": nil,
				},
				Field116: []*int64{},
				Field117: []*string{},
				Field118: map[string]*bool{
					"": nil,
				},
				Field119: map[string]*string{
					"": nil,
				},
				Field120: []*HugeStruct0{GetHugeStruct0()},
				Field121: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field122: []*bool{},
				Field123: nil,
				Field124: []*int64{},
				Field125: nil,
				Field126: []*string{},
				Field127: []*string{},
				Field128: []*int32{},
				Field129: []*bool{},
				Field130: nil,
				Field131: nil,
				Field132: []*int32{},
				Field133: []*int32{},
				Field134: nil,
				Field135: []*bool{},
				Field136: nil,
				Field137: []*int32{},
				Field138: map[string]*int64{
					"": nil,
				},
				Field139: map[string]*string{
					"": nil,
				},
				Field140: map[string]*int64{
					"": nil,
				},
				Field141: map[string]*int64{
					"": nil,
				},
				Field142: []*int32{},
				Field143: []*HugeStruct0{GetHugeStruct0()},
				Field144: map[string]*int64{
					"": nil,
				},
				Field145: []*string{},
				Field146: map[string]*int64{
					"": nil,
				},
				Field147: nil,
				Field148: map[string]*string{
					"": nil,
				},
				Field149: nil,
				Field150: map[string]*int64{
					"": nil,
				},
				Field151: map[string]*int64{
					"": nil,
				},
				Field152: map[string]*int32{
					"": nil,
				},
				Field153: []*int32{},
				Field154: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field155: map[string]*string{
					"": nil,
				},
				Field156: map[string]*int64{
					"": nil,
				},
				Field157: []*int32{},
				Field158: []*int32{},
				Field159: nil,
				Field160: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field161: []*bool{},
				Field162: []*HugeStruct0{GetHugeStruct0()},
				Field163: []*int32{},
				Field164: map[string]*string{
					"": nil,
				},
				Field165: []*bool{},
				Field166: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field167: nil,
				Field168: []*bool{},
				Field169: map[string]*bool{
					"": nil,
				},
				Field170: map[string]*bool{
					"": nil,
				},
				Field171: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field172: map[string]*bool{
					"": nil,
				},
				Field173: []*bool{},
				Field174: map[string]*int64{
					"": nil,
				},
				Field175: []*HugeStruct0{GetHugeStruct0()},
				Field176: []*int32{},
				Field177: []*int64{},
				Field178: map[string]*int64{
					"": nil,
				},
				Field179: []*int32{},
				Field180: []*string{},
				Field181: []*int32{},
				Field182: map[string]*string{
					"": nil,
				},
				Field183: []*int64{},
				Field184: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field185: []*int32{},
				Field186: nil,
				Field187: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field188: []*HugeStruct0{GetHugeStruct0()},
				Field189: nil,
				Field190: []*int64{},
				Field191: map[string]*int32{
					"": nil,
				},
				Field192: []*HugeStruct0{GetHugeStruct0()},
				Field193: []*HugeStruct0{GetHugeStruct0()},
				Field194: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field195: []*bool{},
				Field196: map[string]*bool{
					"": nil,
				},
				Field197: []*bool{},
				Field198: nil,
				Field199: map[string]*int32{
					"": nil,
				},
				Field200: map[string]*int64{
					"": nil,
				},
				Field201: map[string]*string{
					"": nil,
				},
				Field202: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field203: map[string]*int32{
					"": nil,
				},
				Field204: nil,
				Field205: map[string]*string{
					"": nil,
				},
				Field206: []*HugeStruct0{GetHugeStruct0()},
				Field207: []*HugeStruct0{GetHugeStruct0()},
				Field208: nil,
				Field209: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field210: map[string]*string{
					"": nil,
				},
				Field211: map[string]*bool{
					"": nil,
				},
				Field212: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field213: nil,
				Field214: map[string]*bool{
					"": nil,
				},
				Field215: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field216: []*HugeStruct0{GetHugeStruct0()},
				Field217: map[string]*string{
					"": nil,
				},
				Field218: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field219: map[string]*int64{
					"": nil,
				},
				Field220: nil,
				Field221: nil,
				Field222: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field223: []*int64{},
				Field224: []*bool{},
				Field225: []*bool{},
				Field226: map[string]*int64{
					"": nil,
				},
				Field227: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field228: []*int64{},
				Field229: map[string]*bool{
					"": nil,
				},
				Field230: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field231: nil,
				Field232: nil,
				Field233: []*string{},
				Field234: []*HugeStruct0{GetHugeStruct0()},
				Field235: []*string{},
				Field236: nil,
				Field237: nil,
				Field238: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field239: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field240: []*HugeStruct0{GetHugeStruct0()},
				Field241: nil,
				Field242: nil,
				Field243: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field244: map[string]*bool{
					"": nil,
				},
				Field245: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field246: []*int32{},
				Field247: []*bool{},
				Field248: []*string{},
				Field249: nil,
				Field250: []*int32{},
				Field251: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field252: nil,
				Field253: map[string]*string{
					"": nil,
				},
				Field254: map[string]*string{
					"": nil,
				},
				Field255: []*int32{},
				Field256: nil,
				Field257: nil,
				Field258: map[string]*string{
					"": nil,
				},
				Field259: map[string]*int32{
					"": nil,
				},
				Field260: []*int64{},
				Field261: []*int32{},
				Field262: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field263: nil,
				Field264: nil,
				Field265: map[string]*bool{
					"": nil,
				},
				Field266: nil,
				Field267: []*int64{},
				Field268: nil,
				Field269: nil,
				Field270: map[string]*int64{
					"": nil,
				},
				Field271: map[string]*int64{
					"": nil,
				},
				Field272: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field273: []*string{},
				Field274: nil,
				Field275: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field276: map[string]*bool{
					"": nil,
				},
				Field277: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field278: nil,
				Field279: map[string]*string{
					"": nil,
				},
				Field280: nil,
				Field281: nil,
				Field282: nil,
				Field283: nil,
				Field284: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field285: map[string]*int64{
					"": nil,
				},
				Field286: map[string]*bool{
					"": nil,
				},
				Field287: map[string]*string{
					"": nil,
				},
				Field288: nil,
				Field289: nil,
				Field290: nil,
				Field291: []*int64{},
				Field292: map[string]*string{
					"": nil,
				},
				Field293: nil,
				Field294: []*string{},
				Field295: nil,
				Field296: []*HugeStruct0{GetHugeStruct0()},
				Field297: nil,
				Field298: map[string]*int64{
					"": nil,
				},
				Field299: map[string]*bool{
					"": nil,
				},
				Field300: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field301: nil,
				Field302: []*string{},
				Field303: []*string{},
				Field304: map[string]*string{
					"": nil,
				},
				Field305: nil,
				Field306: nil,
				Field307: []*HugeStruct0{GetHugeStruct0()},
				Field308: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field309: map[string]*int32{
					"": nil,
				},
				Field310: []*HugeStruct0{GetHugeStruct0()},
				Field311: nil,
				Field312: []*bool{},
				Field313: nil,
				Field314: []*HugeStruct0{GetHugeStruct0()},
				Field315: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field316: nil,
				Field317: nil,
				Field318: nil,
				Field319: []*int32{},
				Field320: nil,
				Field321: []*HugeStruct0{GetHugeStruct0()},
				Field322: nil,
				Field323: nil,
				Field324: []*HugeStruct0{GetHugeStruct0()},
				Field325: nil,
				Field326: []*int64{},
				Field327: nil,
				Field328: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field329: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field330: []*HugeStruct0{GetHugeStruct0()},
				Field331: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field332: []*string{},
				Field333: nil,
				Field334: []*HugeStruct0{GetHugeStruct0()},
				Field335: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field336: map[string]*bool{
					"": nil,
				},
				Field337: []*int64{},
				Field338: map[string]*bool{
					"": nil,
				},
				Field339: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field340: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field341: []*bool{},
				Field342: []*int64{},
				Field343: []*int32{},
				Field344: map[string]*bool{
					"": nil,
				},
				Field345: map[string]*int64{
					"": nil,
				},
				Field346: nil,
				Field347: map[string]*bool{
					"": nil,
				},
				Field348: map[string]*int32{
					"": nil,
				},
				Field349: []*string{},
				Field350: map[string]*int32{
					"": nil,
				},
				Field351: nil,
				Field352: []*int64{},
				Field353: []*int64{},
				Field354: nil,
				Field355: map[string]*int32{
					"": nil,
				},
				Field356: map[string]*bool{
					"": nil,
				},
				Field357: []*int32{},
				Field358: nil,
				Field359: map[string]*int64{
					"": nil,
				},
				Field360: nil,
				Field361: map[string]*int64{
					"": nil,
				},
				Field362: map[string]*int32{
					"": nil,
				},
				Field363: []*int64{},
				Field364: []*bool{},
				Field365: nil,
				Field366: map[string]*string{
					"": nil,
				},
				Field367: map[string]*bool{
					"": nil,
				},
				Field368: nil,
				Field369: nil,
				Field370: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field371: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field372: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field373: map[string]*bool{
					"": nil,
				},
			},
		},
		Field32: []*bool{},
		Field33: map[string]*bool{
			"": nil,
		},
		Field34: []*string{},
		Field35: []*string{},
		Field36: []*int32{},
		Field37: nil,
		Field38: map[string]*string{
			"": nil,
		},
		Field39: []*string{},
		Field40: []*bool{},
		Field41: []*bool{},
		Field42: map[string]*HugeStruct1{
			"": {
				Field0: []*int32{},
				Field1: []*string{},
				Field2: []*int64{},
				Field3: map[string]*int32{
					"": nil,
				},
				Field4: []*bool{},
				Field5: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field6: map[string]*int32{
					"": nil,
				},
				Field7: map[string]*bool{
					"": nil,
				},
				Field8: []*bool{},
				Field9: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field10: []*string{},
				Field11: []*bool{},
				Field12: []*bool{},
				Field13: map[string]*int32{
					"": nil,
				},
				Field14: map[string]*int32{
					"": nil,
				},
				Field15: nil,
				Field16: []*int64{},
				Field17: []*bool{},
				Field18: map[string]*int64{
					"": nil,
				},
				Field19: []*int64{},
				Field20: map[string]*string{
					"": nil,
				},
				Field21: nil,
				Field22: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field23: []*string{},
				Field24: []*int64{},
				Field25: []*string{},
				Field26: []*bool{},
				Field27: map[string]*int32{
					"": nil,
				},
				Field28: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field29: map[string]*int32{
					"": nil,
				},
				Field30: map[string]*bool{
					"": nil,
				},
				Field31: map[string]*int32{
					"": nil,
				},
				Field32: []*HugeStruct0{GetHugeStruct0()},
				Field33: nil,
				Field34: map[string]*bool{
					"": nil,
				},
				Field35: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct0{GetHugeStruct0()},
				Field39: []*bool{},
				Field40: map[string]*string{
					"": nil,
				},
				Field41: map[string]*int64{
					"": nil,
				},
				Field42: map[string]*int32{
					"": nil,
				},
				Field43: nil,
				Field44: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field47: nil,
				Field48: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field49: nil,
				Field50: map[string]*string{
					"": nil,
				},
				Field51: map[string]*bool{
					"": nil,
				},
				Field52: []*int64{},
				Field53: map[string]*string{
					"": nil,
				},
				Field54: []*int32{},
				Field55: map[string]*int64{
					"": nil,
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field59: []*HugeStruct0{GetHugeStruct0()},
				Field60: map[string]*string{
					"": nil,
				},
				Field61: map[string]*bool{
					"": nil,
				},
				Field62: map[string]*int64{
					"": nil,
				},
				Field63: []*string{},
				Field64: []*int64{},
				Field65: map[string]*bool{
					"": nil,
				},
				Field66: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field67: []*int64{},
				Field68: map[string]*string{
					"": nil,
				},
				Field69: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field70: []*bool{},
				Field71: map[string]*int64{
					"": nil,
				},
				Field72: nil,
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: nil,
				Field75: map[string]*int32{
					"": nil,
				},
				Field76: map[string]*string{
					"": nil,
				},
				Field77: []*string{},
				Field78: nil,
				Field79: map[string]*int64{
					"": nil,
				},
				Field80: []*int64{},
				Field81: map[string]*bool{
					"": nil,
				},
				Field82: []*string{},
				Field83: []*string{},
				Field84: nil,
				Field85: []*bool{},
				Field86: []*HugeStruct0{GetHugeStruct0()},
				Field87: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field88: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field89: []*int64{},
				Field90: []*int32{},
				Field91: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field92: []*bool{},
				Field93: []*string{},
				Field94: map[string]*int32{
					"": nil,
				},
				Field95: nil,
				Field96: nil,
				Field97: map[string]*bool{
					"": nil,
				},
				Field98: map[string]*int32{
					"": nil,
				},
				Field99:  []*HugeStruct0{GetHugeStruct0()},
				Field100: nil,
				Field101: nil,
				Field102: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field103: []*string{},
				Field104: []*string{},
				Field105: map[string]*bool{
					"": nil,
				},
				Field106: []*string{},
				Field107: []*int64{},
				Field108: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field109: nil,
				Field110: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field111: []*string{},
				Field112: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field113: []*bool{},
				Field114: []*bool{},
				Field115: map[string]*string{
					"": nil,
				},
				Field116: []*int64{},
				Field117: []*string{},
				Field118: map[string]*bool{
					"": nil,
				},
				Field119: map[string]*string{
					"": nil,
				},
				Field120: []*HugeStruct0{GetHugeStruct0()},
				Field121: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field122: []*bool{},
				Field123: nil,
				Field124: []*int64{},
				Field125: nil,
				Field126: []*string{},
				Field127: []*string{},
				Field128: []*int32{},
				Field129: []*bool{},
				Field130: nil,
				Field131: nil,
				Field132: []*int32{},
				Field133: []*int32{},
				Field134: nil,
				Field135: []*bool{},
				Field136: nil,
				Field137: []*int32{},
				Field138: map[string]*int64{
					"": nil,
				},
				Field139: map[string]*string{
					"": nil,
				},
				Field140: map[string]*int64{
					"": nil,
				},
				Field141: map[string]*int64{
					"": nil,
				},
				Field142: []*int32{},
				Field143: []*HugeStruct0{GetHugeStruct0()},
				Field144: map[string]*int64{
					"": nil,
				},
				Field145: []*string{},
				Field146: map[string]*int64{
					"": nil,
				},
				Field147: nil,
				Field148: map[string]*string{
					"": nil,
				},
				Field149: nil,
				Field150: map[string]*int64{
					"": nil,
				},
				Field151: map[string]*int64{
					"": nil,
				},
				Field152: map[string]*int32{
					"": nil,
				},
				Field153: []*int32{},
				Field154: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field155: map[string]*string{
					"": nil,
				},
				Field156: map[string]*int64{
					"": nil,
				},
				Field157: []*int32{},
				Field158: []*int32{},
				Field159: nil,
				Field160: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field161: []*bool{},
				Field162: []*HugeStruct0{GetHugeStruct0()},
				Field163: []*int32{},
				Field164: map[string]*string{
					"": nil,
				},
				Field165: []*bool{},
				Field166: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field167: nil,
				Field168: []*bool{},
				Field169: map[string]*bool{
					"": nil,
				},
				Field170: map[string]*bool{
					"": nil,
				},
				Field171: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field172: map[string]*bool{
					"": nil,
				},
				Field173: []*bool{},
				Field174: map[string]*int64{
					"": nil,
				},
				Field175: []*HugeStruct0{GetHugeStruct0()},
				Field176: []*int32{},
				Field177: []*int64{},
				Field178: map[string]*int64{
					"": nil,
				},
				Field179: []*int32{},
				Field180: []*string{},
				Field181: []*int32{},
				Field182: map[string]*string{
					"": nil,
				},
				Field183: []*int64{},
				Field184: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field185: []*int32{},
				Field186: nil,
				Field187: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field188: []*HugeStruct0{GetHugeStruct0()},
				Field189: nil,
				Field190: []*int64{},
				Field191: map[string]*int32{
					"": nil,
				},
				Field192: []*HugeStruct0{GetHugeStruct0()},
				Field193: []*HugeStruct0{GetHugeStruct0()},
				Field194: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field195: []*bool{},
				Field196: map[string]*bool{
					"": nil,
				},
				Field197: []*bool{},
				Field198: nil,
				Field199: map[string]*int32{
					"": nil,
				},
				Field200: map[string]*int64{
					"": nil,
				},
				Field201: map[string]*string{
					"": nil,
				},
				Field202: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field203: map[string]*int32{
					"": nil,
				},
				Field204: nil,
				Field205: map[string]*string{
					"": nil,
				},
				Field206: []*HugeStruct0{GetHugeStruct0()},
				Field207: []*HugeStruct0{GetHugeStruct0()},
				Field208: nil,
				Field209: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field210: map[string]*string{
					"": nil,
				},
				Field211: map[string]*bool{
					"": nil,
				},
				Field212: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field213: nil,
				Field214: map[string]*bool{
					"": nil,
				},
				Field215: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field216: []*HugeStruct0{GetHugeStruct0()},
				Field217: map[string]*string{
					"": nil,
				},
				Field218: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field219: map[string]*int64{
					"": nil,
				},
				Field220: nil,
				Field221: nil,
				Field222: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field223: []*int64{},
				Field224: []*bool{},
				Field225: []*bool{},
				Field226: map[string]*int64{
					"": nil,
				},
				Field227: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field228: []*int64{},
				Field229: map[string]*bool{
					"": nil,
				},
				Field230: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field231: nil,
				Field232: nil,
				Field233: []*string{},
				Field234: []*HugeStruct0{GetHugeStruct0()},
				Field235: []*string{},
				Field236: nil,
				Field237: nil,
				Field238: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field239: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field240: []*HugeStruct0{GetHugeStruct0()},
				Field241: nil,
				Field242: nil,
				Field243: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field244: map[string]*bool{
					"": nil,
				},
				Field245: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field246: []*int32{},
				Field247: []*bool{},
				Field248: []*string{},
				Field249: nil,
				Field250: []*int32{},
				Field251: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field252: nil,
				Field253: map[string]*string{
					"": nil,
				},
				Field254: map[string]*string{
					"": nil,
				},
				Field255: []*int32{},
				Field256: nil,
				Field257: nil,
				Field258: map[string]*string{
					"": nil,
				},
				Field259: map[string]*int32{
					"": nil,
				},
				Field260: []*int64{},
				Field261: []*int32{},
				Field262: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field263: nil,
				Field264: nil,
				Field265: map[string]*bool{
					"": nil,
				},
				Field266: nil,
				Field267: []*int64{},
				Field268: nil,
				Field269: nil,
				Field270: map[string]*int64{
					"": nil,
				},
				Field271: map[string]*int64{
					"": nil,
				},
				Field272: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field273: []*string{},
				Field274: nil,
				Field275: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field276: map[string]*bool{
					"": nil,
				},
				Field277: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field278: nil,
				Field279: map[string]*string{
					"": nil,
				},
				Field280: nil,
				Field281: nil,
				Field282: nil,
				Field283: nil,
				Field284: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field285: map[string]*int64{
					"": nil,
				},
				Field286: map[string]*bool{
					"": nil,
				},
				Field287: map[string]*string{
					"": nil,
				},
				Field288: nil,
				Field289: nil,
				Field290: nil,
				Field291: []*int64{},
				Field292: map[string]*string{
					"": nil,
				},
				Field293: nil,
				Field294: []*string{},
				Field295: nil,
				Field296: []*HugeStruct0{GetHugeStruct0()},
				Field297: nil,
				Field298: map[string]*int64{
					"": nil,
				},
				Field299: map[string]*bool{
					"": nil,
				},
				Field300: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field301: nil,
				Field302: []*string{},
				Field303: []*string{},
				Field304: map[string]*string{
					"": nil,
				},
				Field305: nil,
				Field306: nil,
				Field307: []*HugeStruct0{GetHugeStruct0()},
				Field308: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field309: map[string]*int32{
					"": nil,
				},
				Field310: []*HugeStruct0{GetHugeStruct0()},
				Field311: nil,
				Field312: []*bool{},
				Field313: nil,
				Field314: []*HugeStruct0{GetHugeStruct0()},
				Field315: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field316: nil,
				Field317: nil,
				Field318: nil,
				Field319: []*int32{},
				Field320: nil,
				Field321: []*HugeStruct0{GetHugeStruct0()},
				Field322: nil,
				Field323: nil,
				Field324: []*HugeStruct0{GetHugeStruct0()},
				Field325: nil,
				Field326: []*int64{},
				Field327: nil,
				Field328: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field329: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field330: []*HugeStruct0{GetHugeStruct0()},
				Field331: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field332: []*string{},
				Field333: nil,
				Field334: []*HugeStruct0{GetHugeStruct0()},
				Field335: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field336: map[string]*bool{
					"": nil,
				},
				Field337: []*int64{},
				Field338: map[string]*bool{
					"": nil,
				},
				Field339: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field340: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field341: []*bool{},
				Field342: []*int64{},
				Field343: []*int32{},
				Field344: map[string]*bool{
					"": nil,
				},
				Field345: map[string]*int64{
					"": nil,
				},
				Field346: nil,
				Field347: map[string]*bool{
					"": nil,
				},
				Field348: map[string]*int32{
					"": nil,
				},
				Field349: []*string{},
				Field350: map[string]*int32{
					"": nil,
				},
				Field351: nil,
				Field352: []*int64{},
				Field353: []*int64{},
				Field354: nil,
				Field355: map[string]*int32{
					"": nil,
				},
				Field356: map[string]*bool{
					"": nil,
				},
				Field357: []*int32{},
				Field358: nil,
				Field359: map[string]*int64{
					"": nil,
				},
				Field360: nil,
				Field361: map[string]*int64{
					"": nil,
				},
				Field362: map[string]*int32{
					"": nil,
				},
				Field363: []*int64{},
				Field364: []*bool{},
				Field365: nil,
				Field366: map[string]*string{
					"": nil,
				},
				Field367: map[string]*bool{
					"": nil,
				},
				Field368: nil,
				Field369: nil,
				Field370: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field371: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field372: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field373: map[string]*bool{
					"": nil,
				},
			},
		},
		Field43: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field44: nil,
		Field45: []*string{},
		Field46: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field47: map[string]*int64{
			"": nil,
		},
		Field48: map[string]*HugeStruct2{
			"": {
				Field0: nil,
				Field1: map[string]*int64{
					"": nil,
				},
				Field2: nil,
				Field3: []*int64{},
				Field4: map[string]*int32{
					"": nil,
				},
				Field5: map[string]*int32{
					"": nil,
				},
				Field6: nil,
				Field7: map[string]*int32{
					"": nil,
				},
				Field8:  nil,
				Field9:  []*HugeStruct1{},
				Field10: nil,
				Field11: map[string]*int64{
					"": nil,
				},
				Field12: nil,
				Field13: nil,
				Field14: map[string]*HugeStruct1{
					"": {
						Field0: []*int32{},
						Field1: []*string{},
						Field2: []*int64{},
						Field3: map[string]*int32{
							"": nil,
						},
						Field4: []*bool{},
						Field5: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field6: map[string]*int32{
							"": nil,
						},
						Field7: map[string]*bool{
							"": nil,
						},
						Field8: []*bool{},
						Field9: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field10: []*string{},
						Field11: []*bool{},
						Field12: []*bool{},
						Field13: map[string]*int32{
							"": nil,
						},
						Field14: map[string]*int32{
							"": nil,
						},
						Field15: nil,
						Field16: []*int64{},
						Field17: []*bool{},
						Field18: map[string]*int64{
							"": nil,
						},
						Field19: []*int64{},
						Field20: map[string]*string{
							"": nil,
						},
						Field21: nil,
						Field22: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field23: []*string{},
						Field24: []*int64{},
						Field25: []*string{},
						Field26: []*bool{},
						Field27: map[string]*int32{
							"": nil,
						},
						Field28: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field29: map[string]*int32{
							"": nil,
						},
						Field30: map[string]*bool{
							"": nil,
						},
						Field31: map[string]*int32{
							"": nil,
						},
						Field32: []*HugeStruct0{GetHugeStruct0()},
						Field33: nil,
						Field34: map[string]*bool{
							"": nil,
						},
						Field35: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field36: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field37: nil,
						Field38: []*HugeStruct0{GetHugeStruct0()},
						Field39: []*bool{},
						Field40: map[string]*string{
							"": nil,
						},
						Field41: map[string]*int64{
							"": nil,
						},
						Field42: map[string]*int32{
							"": nil,
						},
						Field43: nil,
						Field44: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field45: map[string]*int32{
							"": nil,
						},
						Field46: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field47: nil,
						Field48: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field49: nil,
						Field50: map[string]*string{
							"": nil,
						},
						Field51: map[string]*bool{
							"": nil,
						},
						Field52: []*int64{},
						Field53: map[string]*string{
							"": nil,
						},
						Field54: []*int32{},
						Field55: map[string]*int64{
							"": nil,
						},
						Field56: map[string]*int32{
							"": nil,
						},
						Field57: map[string]*string{
							"": nil,
						},
						Field58: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field59: []*HugeStruct0{GetHugeStruct0()},
						Field60: map[string]*string{
							"": nil,
						},
						Field61: map[string]*bool{
							"": nil,
						},
						Field62: map[string]*int64{
							"": nil,
						},
						Field63: []*string{},
						Field64: []*int64{},
						Field65: map[string]*bool{
							"": nil,
						},
						Field66: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field67: []*int64{},
						Field68: map[string]*string{
							"": nil,
						},
						Field69: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field70: []*bool{},
						Field71: map[string]*int64{
							"": nil,
						},
						Field72: nil,
						Field73: map[string]*int32{
							"": nil,
						},
						Field74: nil,
						Field75: map[string]*int32{
							"": nil,
						},
						Field76: map[string]*string{
							"": nil,
						},
						Field77: []*string{},
						Field78: nil,
						Field79: map[string]*int64{
							"": nil,
						},
						Field80: []*int64{},
						Field81: map[string]*bool{
							"": nil,
						},
						Field82: []*string{},
						Field83: []*string{},
						Field84: nil,
						Field85: []*bool{},
						Field86: []*HugeStruct0{GetHugeStruct0()},
						Field87: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field88: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field89: []*int64{},
						Field90: []*int32{},
						Field91: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field92: []*bool{},
						Field93: []*string{},
						Field94: map[string]*int32{
							"": nil,
						},
						Field95: nil,
						Field96: nil,
						Field97: map[string]*bool{
							"": nil,
						},
						Field98: map[string]*int32{
							"": nil,
						},
						Field99:  []*HugeStruct0{GetHugeStruct0()},
						Field100: nil,
						Field101: nil,
						Field102: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field103: []*string{},
						Field104: []*string{},
						Field105: map[string]*bool{
							"": nil,
						},
						Field106: []*string{},
						Field107: []*int64{},
						Field108: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field109: nil,
						Field110: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field111: []*string{},
						Field112: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field113: []*bool{},
						Field114: []*bool{},
						Field115: map[string]*string{
							"": nil,
						},
						Field116: []*int64{},
						Field117: []*string{},
						Field118: map[string]*bool{
							"": nil,
						},
						Field119: map[string]*string{
							"": nil,
						},
						Field120: []*HugeStruct0{GetHugeStruct0()},
						Field121: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field122: []*bool{},
						Field123: nil,
						Field124: []*int64{},
						Field125: nil,
						Field126: []*string{},
						Field127: []*string{},
						Field128: []*int32{},
						Field129: []*bool{},
						Field130: nil,
						Field131: nil,
						Field132: []*int32{},
						Field133: []*int32{},
						Field134: nil,
						Field135: []*bool{},
						Field136: nil,
						Field137: []*int32{},
						Field138: map[string]*int64{
							"": nil,
						},
						Field139: map[string]*string{
							"": nil,
						},
						Field140: map[string]*int64{
							"": nil,
						},
						Field141: map[string]*int64{
							"": nil,
						},
						Field142: []*int32{},
						Field143: []*HugeStruct0{GetHugeStruct0()},
						Field144: map[string]*int64{
							"": nil,
						},
						Field145: []*string{},
						Field146: map[string]*int64{
							"": nil,
						},
						Field147: nil,
						Field148: map[string]*string{
							"": nil,
						},
						Field149: nil,
						Field150: map[string]*int64{
							"": nil,
						},
						Field151: map[string]*int64{
							"": nil,
						},
						Field152: map[string]*int32{
							"": nil,
						},
						Field153: []*int32{},
						Field154: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field155: map[string]*string{
							"": nil,
						},
						Field156: map[string]*int64{
							"": nil,
						},
						Field157: []*int32{},
						Field158: []*int32{},
						Field159: nil,
						Field160: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field161: []*bool{},
						Field162: []*HugeStruct0{GetHugeStruct0()},
						Field163: []*int32{},
						Field164: map[string]*string{
							"": nil,
						},
						Field165: []*bool{},
						Field166: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field167: nil,
						Field168: []*bool{},
						Field169: map[string]*bool{
							"": nil,
						},
						Field170: map[string]*bool{
							"": nil,
						},
						Field171: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field172: map[string]*bool{
							"": nil,
						},
						Field173: []*bool{},
						Field174: map[string]*int64{
							"": nil,
						},
						Field175: []*HugeStruct0{GetHugeStruct0()},
						Field176: []*int32{},
						Field177: []*int64{},
						Field178: map[string]*int64{
							"": nil,
						},
						Field179: []*int32{},
						Field180: []*string{},
						Field181: []*int32{},
						Field182: map[string]*string{
							"": nil,
						},
						Field183: []*int64{},
						Field184: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field185: []*int32{},
						Field186: nil,
						Field187: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field188: []*HugeStruct0{GetHugeStruct0()},
						Field189: nil,
						Field190: []*int64{},
						Field191: map[string]*int32{
							"": nil,
						},
						Field192: []*HugeStruct0{GetHugeStruct0()},
						Field193: []*HugeStruct0{GetHugeStruct0()},
						Field194: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field195: []*bool{},
						Field196: map[string]*bool{
							"": nil,
						},
						Field197: []*bool{},
						Field198: nil,
						Field199: map[string]*int32{
							"": nil,
						},
						Field200: map[string]*int64{
							"": nil,
						},
						Field201: map[string]*string{
							"": nil,
						},
						Field202: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field203: map[string]*int32{
							"": nil,
						},
						Field204: nil,
						Field205: map[string]*string{
							"": nil,
						},
						Field206: []*HugeStruct0{GetHugeStruct0()},
						Field207: []*HugeStruct0{GetHugeStruct0()},
						Field208: nil,
						Field209: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field210: map[string]*string{
							"": nil,
						},
						Field211: map[string]*bool{
							"": nil,
						},
						Field212: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field213: nil,
						Field214: map[string]*bool{
							"": nil,
						},
						Field215: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field216: []*HugeStruct0{GetHugeStruct0()},
						Field217: map[string]*string{
							"": nil,
						},
						Field218: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field219: map[string]*int64{
							"": nil,
						},
						Field220: nil,
						Field221: nil,
						Field222: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field223: []*int64{},
						Field224: []*bool{},
						Field225: []*bool{},
						Field226: map[string]*int64{
							"": nil,
						},
						Field227: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field228: []*int64{},
						Field229: map[string]*bool{
							"": nil,
						},
						Field230: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field231: nil,
						Field232: nil,
						Field233: []*string{},
						Field234: []*HugeStruct0{GetHugeStruct0()},
						Field235: []*string{},
						Field236: nil,
						Field237: nil,
						Field238: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field239: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field240: []*HugeStruct0{GetHugeStruct0()},
						Field241: nil,
						Field242: nil,
						Field243: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field244: map[string]*bool{
							"": nil,
						},
						Field245: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field246: []*int32{},
						Field247: []*bool{},
						Field248: []*string{},
						Field249: nil,
						Field250: []*int32{},
						Field251: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field252: nil,
						Field253: map[string]*string{
							"": nil,
						},
						Field254: map[string]*string{
							"": nil,
						},
						Field255: []*int32{},
						Field256: nil,
						Field257: nil,
						Field258: map[string]*string{
							"": nil,
						},
						Field259: map[string]*int32{
							"": nil,
						},
						Field260: []*int64{},
						Field261: []*int32{},
						Field262: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field263: nil,
						Field264: nil,
						Field265: map[string]*bool{
							"": nil,
						},
						Field266: nil,
						Field267: []*int64{},
						Field268: nil,
						Field269: nil,
						Field270: map[string]*int64{
							"": nil,
						},
						Field271: map[string]*int64{
							"": nil,
						},
						Field272: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field273: []*string{},
						Field274: nil,
						Field275: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field276: map[string]*bool{
							"": nil,
						},
						Field277: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field278: nil,
						Field279: map[string]*string{
							"": nil,
						},
						Field280: nil,
						Field281: nil,
						Field282: nil,
						Field283: nil,
						Field284: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field285: map[string]*int64{
							"": nil,
						},
						Field286: map[string]*bool{
							"": nil,
						},
						Field287: map[string]*string{
							"": nil,
						},
						Field288: nil,
						Field289: nil,
						Field290: nil,
						Field291: []*int64{},
						Field292: map[string]*string{
							"": nil,
						},
						Field293: nil,
						Field294: []*string{},
						Field295: nil,
						Field296: []*HugeStruct0{GetHugeStruct0()},
						Field297: nil,
						Field298: map[string]*int64{
							"": nil,
						},
						Field299: map[string]*bool{
							"": nil,
						},
						Field300: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field301: nil,
						Field302: []*string{},
						Field303: []*string{},
						Field304: map[string]*string{
							"": nil,
						},
						Field305: nil,
						Field306: nil,
						Field307: []*HugeStruct0{GetHugeStruct0()},
						Field308: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field309: map[string]*int32{
							"": nil,
						},
						Field310: []*HugeStruct0{GetHugeStruct0()},
						Field311: nil,
						Field312: []*bool{},
						Field313: nil,
						Field314: []*HugeStruct0{GetHugeStruct0()},
						Field315: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field316: nil,
						Field317: nil,
						Field318: nil,
						Field319: []*int32{},
						Field320: nil,
						Field321: []*HugeStruct0{GetHugeStruct0()},
						Field322: nil,
						Field323: nil,
						Field324: []*HugeStruct0{GetHugeStruct0()},
						Field325: nil,
						Field326: []*int64{},
						Field327: nil,
						Field328: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field329: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field330: []*HugeStruct0{GetHugeStruct0()},
						Field331: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field332: []*string{},
						Field333: nil,
						Field334: []*HugeStruct0{GetHugeStruct0()},
						Field335: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field336: map[string]*bool{
							"": nil,
						},
						Field337: []*int64{},
						Field338: map[string]*bool{
							"": nil,
						},
						Field339: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field340: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field341: []*bool{},
						Field342: []*int64{},
						Field343: []*int32{},
						Field344: map[string]*bool{
							"": nil,
						},
						Field345: map[string]*int64{
							"": nil,
						},
						Field346: nil,
						Field347: map[string]*bool{
							"": nil,
						},
						Field348: map[string]*int32{
							"": nil,
						},
						Field349: []*string{},
						Field350: map[string]*int32{
							"": nil,
						},
						Field351: nil,
						Field352: []*int64{},
						Field353: []*int64{},
						Field354: nil,
						Field355: map[string]*int32{
							"": nil,
						},
						Field356: map[string]*bool{
							"": nil,
						},
						Field357: []*int32{},
						Field358: nil,
						Field359: map[string]*int64{
							"": nil,
						},
						Field360: nil,
						Field361: map[string]*int64{
							"": nil,
						},
						Field362: map[string]*int32{
							"": nil,
						},
						Field363: []*int64{},
						Field364: []*bool{},
						Field365: nil,
						Field366: map[string]*string{
							"": nil,
						},
						Field367: map[string]*bool{
							"": nil,
						},
						Field368: nil,
						Field369: nil,
						Field370: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field371: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field372: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field373: map[string]*bool{
							"": nil,
						},
					},
				},
				Field15: map[string]*int64{
					"": nil,
				},
				Field16: map[string]*int32{
					"": nil,
				},
				Field17: map[string]*int32{
					"": nil,
				},
				Field18: []*int32{},
				Field19: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field20: map[string]*int64{
					"": nil,
				},
				Field21: &HugeStruct1{
					Field0: []*int32{},
					Field1: []*string{},
					Field2: []*int64{},
					Field3: map[string]*int32{
						"": nil,
					},
					Field4: []*bool{},
					Field5: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field6: map[string]*int32{
						"": nil,
					},
					Field7: map[string]*bool{
						"": nil,
					},
					Field8: []*bool{},
					Field9: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field10: []*string{},
					Field11: []*bool{},
					Field12: []*bool{},
					Field13: map[string]*int32{
						"": nil,
					},
					Field14: map[string]*int32{
						"": nil,
					},
					Field15: nil,
					Field16: []*int64{},
					Field17: []*bool{},
					Field18: map[string]*int64{
						"": nil,
					},
					Field19: []*int64{},
					Field20: map[string]*string{
						"": nil,
					},
					Field21: nil,
					Field22: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field23: []*string{},
					Field24: []*int64{},
					Field25: []*string{},
					Field26: []*bool{},
					Field27: map[string]*int32{
						"": nil,
					},
					Field28: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field29: map[string]*int32{
						"": nil,
					},
					Field30: map[string]*bool{
						"": nil,
					},
					Field31: map[string]*int32{
						"": nil,
					},
					Field32: []*HugeStruct0{GetHugeStruct0()},
					Field33: nil,
					Field34: map[string]*bool{
						"": nil,
					},
					Field35: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field36: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field37: nil,
					Field38: []*HugeStruct0{GetHugeStruct0()},
					Field39: []*bool{},
					Field40: map[string]*string{
						"": nil,
					},
					Field41: map[string]*int64{
						"": nil,
					},
					Field42: map[string]*int32{
						"": nil,
					},
					Field43: nil,
					Field44: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field45: map[string]*int32{
						"": nil,
					},
					Field46: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field47: nil,
					Field48: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field49: nil,
					Field50: map[string]*string{
						"": nil,
					},
					Field51: map[string]*bool{
						"": nil,
					},
					Field52: []*int64{},
					Field53: map[string]*string{
						"": nil,
					},
					Field54: []*int32{},
					Field55: map[string]*int64{
						"": nil,
					},
					Field56: map[string]*int32{
						"": nil,
					},
					Field57: map[string]*string{
						"": nil,
					},
					Field58: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field59: []*HugeStruct0{GetHugeStruct0()},
					Field60: map[string]*string{
						"": nil,
					},
					Field61: map[string]*bool{
						"": nil,
					},
					Field62: map[string]*int64{
						"": nil,
					},
					Field63: []*string{},
					Field64: []*int64{},
					Field65: map[string]*bool{
						"": nil,
					},
					Field66: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field67: []*int64{},
					Field68: map[string]*string{
						"": nil,
					},
					Field69: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field70: []*bool{},
					Field71: map[string]*int64{
						"": nil,
					},
					Field72: nil,
					Field73: map[string]*int32{
						"": nil,
					},
					Field74: nil,
					Field75: map[string]*int32{
						"": nil,
					},
					Field76: map[string]*string{
						"": nil,
					},
					Field77: []*string{},
					Field78: nil,
					Field79: map[string]*int64{
						"": nil,
					},
					Field80: []*int64{},
					Field81: map[string]*bool{
						"": nil,
					},
					Field82: []*string{},
					Field83: []*string{},
					Field84: nil,
					Field85: []*bool{},
					Field86: []*HugeStruct0{GetHugeStruct0()},
					Field87: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field88: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field89: []*int64{},
					Field90: []*int32{},
					Field91: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field92: []*bool{},
					Field93: []*string{},
					Field94: map[string]*int32{
						"": nil,
					},
					Field95: nil,
					Field96: nil,
					Field97: map[string]*bool{
						"": nil,
					},
					Field98: map[string]*int32{
						"": nil,
					},
					Field99:  []*HugeStruct0{GetHugeStruct0()},
					Field100: nil,
					Field101: nil,
					Field102: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field103: []*string{},
					Field104: []*string{},
					Field105: map[string]*bool{
						"": nil,
					},
					Field106: []*string{},
					Field107: []*int64{},
					Field108: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field109: nil,
					Field110: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field111: []*string{},
					Field112: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field113: []*bool{},
					Field114: []*bool{},
					Field115: map[string]*string{
						"": nil,
					},
					Field116: []*int64{},
					Field117: []*string{},
					Field118: map[string]*bool{
						"": nil,
					},
					Field119: map[string]*string{
						"": nil,
					},
					Field120: []*HugeStruct0{GetHugeStruct0()},
					Field121: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field122: []*bool{},
					Field123: nil,
					Field124: []*int64{},
					Field125: nil,
					Field126: []*string{},
					Field127: []*string{},
					Field128: []*int32{},
					Field129: []*bool{},
					Field130: nil,
					Field131: nil,
					Field132: []*int32{},
					Field133: []*int32{},
					Field134: nil,
					Field135: []*bool{},
					Field136: nil,
					Field137: []*int32{},
					Field138: map[string]*int64{
						"": nil,
					},
					Field139: map[string]*string{
						"": nil,
					},
					Field140: map[string]*int64{
						"": nil,
					},
					Field141: map[string]*int64{
						"": nil,
					},
					Field142: []*int32{},
					Field143: []*HugeStruct0{GetHugeStruct0()},
					Field144: map[string]*int64{
						"": nil,
					},
					Field145: []*string{},
					Field146: map[string]*int64{
						"": nil,
					},
					Field147: nil,
					Field148: map[string]*string{
						"": nil,
					},
					Field149: nil,
					Field150: map[string]*int64{
						"": nil,
					},
					Field151: map[string]*int64{
						"": nil,
					},
					Field152: map[string]*int32{
						"": nil,
					},
					Field153: []*int32{},
					Field154: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field155: map[string]*string{
						"": nil,
					},
					Field156: map[string]*int64{
						"": nil,
					},
					Field157: []*int32{},
					Field158: []*int32{},
					Field159: nil,
					Field160: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field161: []*bool{},
					Field162: []*HugeStruct0{GetHugeStruct0()},
					Field163: []*int32{},
					Field164: map[string]*string{
						"": nil,
					},
					Field165: []*bool{},
					Field166: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field167: nil,
					Field168: []*bool{},
					Field169: map[string]*bool{
						"": nil,
					},
					Field170: map[string]*bool{
						"": nil,
					},
					Field171: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field172: map[string]*bool{
						"": nil,
					},
					Field173: []*bool{},
					Field174: map[string]*int64{
						"": nil,
					},
					Field175: []*HugeStruct0{GetHugeStruct0()},
					Field176: []*int32{},
					Field177: []*int64{},
					Field178: map[string]*int64{
						"": nil,
					},
					Field179: []*int32{},
					Field180: []*string{},
					Field181: []*int32{},
					Field182: map[string]*string{
						"": nil,
					},
					Field183: []*int64{},
					Field184: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field185: []*int32{},
					Field186: nil,
					Field187: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field188: []*HugeStruct0{GetHugeStruct0()},
					Field189: nil,
					Field190: []*int64{},
					Field191: map[string]*int32{
						"": nil,
					},
					Field192: []*HugeStruct0{GetHugeStruct0()},
					Field193: []*HugeStruct0{GetHugeStruct0()},
					Field194: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field195: []*bool{},
					Field196: map[string]*bool{
						"": nil,
					},
					Field197: []*bool{},
					Field198: nil,
					Field199: map[string]*int32{
						"": nil,
					},
					Field200: map[string]*int64{
						"": nil,
					},
					Field201: map[string]*string{
						"": nil,
					},
					Field202: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field203: map[string]*int32{
						"": nil,
					},
					Field204: nil,
					Field205: map[string]*string{
						"": nil,
					},
					Field206: []*HugeStruct0{GetHugeStruct0()},
					Field207: []*HugeStruct0{GetHugeStruct0()},
					Field208: nil,
					Field209: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field210: map[string]*string{
						"": nil,
					},
					Field211: map[string]*bool{
						"": nil,
					},
					Field212: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field213: nil,
					Field214: map[string]*bool{
						"": nil,
					},
					Field215: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field216: []*HugeStruct0{GetHugeStruct0()},
					Field217: map[string]*string{
						"": nil,
					},
					Field218: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field219: map[string]*int64{
						"": nil,
					},
					Field220: nil,
					Field221: nil,
					Field222: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field223: []*int64{},
					Field224: []*bool{},
					Field225: []*bool{},
					Field226: map[string]*int64{
						"": nil,
					},
					Field227: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field228: []*int64{},
					Field229: map[string]*bool{
						"": nil,
					},
					Field230: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field231: nil,
					Field232: nil,
					Field233: []*string{},
					Field234: []*HugeStruct0{GetHugeStruct0()},
					Field235: []*string{},
					Field236: nil,
					Field237: nil,
					Field238: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field239: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field240: []*HugeStruct0{GetHugeStruct0()},
					Field241: nil,
					Field242: nil,
					Field243: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field244: map[string]*bool{
						"": nil,
					},
					Field245: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field246: []*int32{},
					Field247: []*bool{},
					Field248: []*string{},
					Field249: nil,
					Field250: []*int32{},
					Field251: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field252: nil,
					Field253: map[string]*string{
						"": nil,
					},
					Field254: map[string]*string{
						"": nil,
					},
					Field255: []*int32{},
					Field256: nil,
					Field257: nil,
					Field258: map[string]*string{
						"": nil,
					},
					Field259: map[string]*int32{
						"": nil,
					},
					Field260: []*int64{},
					Field261: []*int32{},
					Field262: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field263: nil,
					Field264: nil,
					Field265: map[string]*bool{
						"": nil,
					},
					Field266: nil,
					Field267: []*int64{},
					Field268: nil,
					Field269: nil,
					Field270: map[string]*int64{
						"": nil,
					},
					Field271: map[string]*int64{
						"": nil,
					},
					Field272: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field273: []*string{},
					Field274: nil,
					Field275: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field276: map[string]*bool{
						"": nil,
					},
					Field277: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field278: nil,
					Field279: map[string]*string{
						"": nil,
					},
					Field280: nil,
					Field281: nil,
					Field282: nil,
					Field283: nil,
					Field284: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field285: map[string]*int64{
						"": nil,
					},
					Field286: map[string]*bool{
						"": nil,
					},
					Field287: map[string]*string{
						"": nil,
					},
					Field288: nil,
					Field289: nil,
					Field290: nil,
					Field291: []*int64{},
					Field292: map[string]*string{
						"": nil,
					},
					Field293: nil,
					Field294: []*string{},
					Field295: nil,
					Field296: []*HugeStruct0{GetHugeStruct0()},
					Field297: nil,
					Field298: map[string]*int64{
						"": nil,
					},
					Field299: map[string]*bool{
						"": nil,
					},
					Field300: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field301: nil,
					Field302: []*string{},
					Field303: []*string{},
					Field304: map[string]*string{
						"": nil,
					},
					Field305: nil,
					Field306: nil,
					Field307: []*HugeStruct0{GetHugeStruct0()},
					Field308: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field309: map[string]*int32{
						"": nil,
					},
					Field310: []*HugeStruct0{GetHugeStruct0()},
					Field311: nil,
					Field312: []*bool{},
					Field313: nil,
					Field314: []*HugeStruct0{GetHugeStruct0()},
					Field315: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field316: nil,
					Field317: nil,
					Field318: nil,
					Field319: []*int32{},
					Field320: nil,
					Field321: []*HugeStruct0{GetHugeStruct0()},
					Field322: nil,
					Field323: nil,
					Field324: []*HugeStruct0{GetHugeStruct0()},
					Field325: nil,
					Field326: []*int64{},
					Field327: nil,
					Field328: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field329: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field330: []*HugeStruct0{GetHugeStruct0()},
					Field331: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field332: []*string{},
					Field333: nil,
					Field334: []*HugeStruct0{GetHugeStruct0()},
					Field335: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field336: map[string]*bool{
						"": nil,
					},
					Field337: []*int64{},
					Field338: map[string]*bool{
						"": nil,
					},
					Field339: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field340: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field341: []*bool{},
					Field342: []*int64{},
					Field343: []*int32{},
					Field344: map[string]*bool{
						"": nil,
					},
					Field345: map[string]*int64{
						"": nil,
					},
					Field346: nil,
					Field347: map[string]*bool{
						"": nil,
					},
					Field348: map[string]*int32{
						"": nil,
					},
					Field349: []*string{},
					Field350: map[string]*int32{
						"": nil,
					},
					Field351: nil,
					Field352: []*int64{},
					Field353: []*int64{},
					Field354: nil,
					Field355: map[string]*int32{
						"": nil,
					},
					Field356: map[string]*bool{
						"": nil,
					},
					Field357: []*int32{},
					Field358: nil,
					Field359: map[string]*int64{
						"": nil,
					},
					Field360: nil,
					Field361: map[string]*int64{
						"": nil,
					},
					Field362: map[string]*int32{
						"": nil,
					},
					Field363: []*int64{},
					Field364: []*bool{},
					Field365: nil,
					Field366: map[string]*string{
						"": nil,
					},
					Field367: map[string]*bool{
						"": nil,
					},
					Field368: nil,
					Field369: nil,
					Field370: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field371: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field372: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field373: map[string]*bool{
						"": nil,
					},
				},
				Field22: []*int32{},
				Field23: map[string]*int64{
					"": nil,
				},
				Field24: map[string]*int64{
					"": nil,
				},
				Field25: nil,
				Field26: map[string]*string{
					"": nil,
				},
				Field27: []*bool{},
				Field28: nil,
				Field29: []*string{},
				Field30: []*HugeStruct0{GetHugeStruct0()},
				Field31: []*int64{},
				Field32: nil,
				Field33: map[string]*string{
					"": nil,
				},
				Field34: []*HugeStruct0{GetHugeStruct0()},
				Field35: nil,
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct1{},
				Field39: nil,
				Field40: map[string]*string{
					"": nil,
				},
				Field41: nil,
				Field42: nil,
				Field43: map[string]*int64{
					"": nil,
				},
				Field44: map[string]*string{
					"": nil,
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: nil,
				Field47: map[string]*int64{
					"": nil,
				},
				Field48: nil,
				Field49: []*HugeStruct1{},
				Field50: nil,
				Field51: []*int64{},
				Field52: map[string]*int64{
					"": nil,
				},
				Field53: nil,
				Field54: map[string]*bool{
					"": nil,
				},
				Field55: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: []*int64{},
				Field59: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field60: []*int64{},
				Field61: map[string]*int64{
					"": nil,
				},
				Field62: map[string]*HugeStruct1{
					"": {
						Field0: []*int32{},
						Field1: []*string{},
						Field2: []*int64{},
						Field3: map[string]*int32{
							"": nil,
						},
						Field4: []*bool{},
						Field5: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field6: map[string]*int32{
							"": nil,
						},
						Field7: map[string]*bool{
							"": nil,
						},
						Field8: []*bool{},
						Field9: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field10: []*string{},
						Field11: []*bool{},
						Field12: []*bool{},
						Field13: map[string]*int32{
							"": nil,
						},
						Field14: map[string]*int32{
							"": nil,
						},
						Field15: nil,
						Field16: []*int64{},
						Field17: []*bool{},
						Field18: map[string]*int64{
							"": nil,
						},
						Field19: []*int64{},
						Field20: map[string]*string{
							"": nil,
						},
						Field21: nil,
						Field22: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field23: []*string{},
						Field24: []*int64{},
						Field25: []*string{},
						Field26: []*bool{},
						Field27: map[string]*int32{
							"": nil,
						},
						Field28: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field29: map[string]*int32{
							"": nil,
						},
						Field30: map[string]*bool{
							"": nil,
						},
						Field31: map[string]*int32{
							"": nil,
						},
						Field32: []*HugeStruct0{GetHugeStruct0()},
						Field33: nil,
						Field34: map[string]*bool{
							"": nil,
						},
						Field35: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field36: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field37: nil,
						Field38: []*HugeStruct0{GetHugeStruct0()},
						Field39: []*bool{},
						Field40: map[string]*string{
							"": nil,
						},
						Field41: map[string]*int64{
							"": nil,
						},
						Field42: map[string]*int32{
							"": nil,
						},
						Field43: nil,
						Field44: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field45: map[string]*int32{
							"": nil,
						},
						Field46: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field47: nil,
						Field48: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field49: nil,
						Field50: map[string]*string{
							"": nil,
						},
						Field51: map[string]*bool{
							"": nil,
						},
						Field52: []*int64{},
						Field53: map[string]*string{
							"": nil,
						},
						Field54: []*int32{},
						Field55: map[string]*int64{
							"": nil,
						},
						Field56: map[string]*int32{
							"": nil,
						},
						Field57: map[string]*string{
							"": nil,
						},
						Field58: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field59: []*HugeStruct0{GetHugeStruct0()},
						Field60: map[string]*string{
							"": nil,
						},
						Field61: map[string]*bool{
							"": nil,
						},
						Field62: map[string]*int64{
							"": nil,
						},
						Field63: []*string{},
						Field64: []*int64{},
						Field65: map[string]*bool{
							"": nil,
						},
						Field66: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field67: []*int64{},
						Field68: map[string]*string{
							"": nil,
						},
						Field69: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field70: []*bool{},
						Field71: map[string]*int64{
							"": nil,
						},
						Field72: nil,
						Field73: map[string]*int32{
							"": nil,
						},
						Field74: nil,
						Field75: map[string]*int32{
							"": nil,
						},
						Field76: map[string]*string{
							"": nil,
						},
						Field77: []*string{},
						Field78: nil,
						Field79: map[string]*int64{
							"": nil,
						},
						Field80: []*int64{},
						Field81: map[string]*bool{
							"": nil,
						},
						Field82: []*string{},
						Field83: []*string{},
						Field84: nil,
						Field85: []*bool{},
						Field86: []*HugeStruct0{GetHugeStruct0()},
						Field87: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field88: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field89: []*int64{},
						Field90: []*int32{},
						Field91: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field92: []*bool{},
						Field93: []*string{},
						Field94: map[string]*int32{
							"": nil,
						},
						Field95: nil,
						Field96: nil,
						Field97: map[string]*bool{
							"": nil,
						},
						Field98: map[string]*int32{
							"": nil,
						},
						Field99:  []*HugeStruct0{GetHugeStruct0()},
						Field100: nil,
						Field101: nil,
						Field102: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field103: []*string{},
						Field104: []*string{},
						Field105: map[string]*bool{
							"": nil,
						},
						Field106: []*string{},
						Field107: []*int64{},
						Field108: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field109: nil,
						Field110: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field111: []*string{},
						Field112: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field113: []*bool{},
						Field114: []*bool{},
						Field115: map[string]*string{
							"": nil,
						},
						Field116: []*int64{},
						Field117: []*string{},
						Field118: map[string]*bool{
							"": nil,
						},
						Field119: map[string]*string{
							"": nil,
						},
						Field120: []*HugeStruct0{GetHugeStruct0()},
						Field121: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field122: []*bool{},
						Field123: nil,
						Field124: []*int64{},
						Field125: nil,
						Field126: []*string{},
						Field127: []*string{},
						Field128: []*int32{},
						Field129: []*bool{},
						Field130: nil,
						Field131: nil,
						Field132: []*int32{},
						Field133: []*int32{},
						Field134: nil,
						Field135: []*bool{},
						Field136: nil,
						Field137: []*int32{},
						Field138: map[string]*int64{
							"": nil,
						},
						Field139: map[string]*string{
							"": nil,
						},
						Field140: map[string]*int64{
							"": nil,
						},
						Field141: map[string]*int64{
							"": nil,
						},
						Field142: []*int32{},
						Field143: []*HugeStruct0{GetHugeStruct0()},
						Field144: map[string]*int64{
							"": nil,
						},
						Field145: []*string{},
						Field146: map[string]*int64{
							"": nil,
						},
						Field147: nil,
						Field148: map[string]*string{
							"": nil,
						},
						Field149: nil,
						Field150: map[string]*int64{
							"": nil,
						},
						Field151: map[string]*int64{
							"": nil,
						},
						Field152: map[string]*int32{
							"": nil,
						},
						Field153: []*int32{},
						Field154: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field155: map[string]*string{
							"": nil,
						},
						Field156: map[string]*int64{
							"": nil,
						},
						Field157: []*int32{},
						Field158: []*int32{},
						Field159: nil,
						Field160: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field161: []*bool{},
						Field162: []*HugeStruct0{GetHugeStruct0()},
						Field163: []*int32{},
						Field164: map[string]*string{
							"": nil,
						},
						Field165: []*bool{},
						Field166: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field167: nil,
						Field168: []*bool{},
						Field169: map[string]*bool{
							"": nil,
						},
						Field170: map[string]*bool{
							"": nil,
						},
						Field171: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field172: map[string]*bool{
							"": nil,
						},
						Field173: []*bool{},
						Field174: map[string]*int64{
							"": nil,
						},
						Field175: []*HugeStruct0{GetHugeStruct0()},
						Field176: []*int32{},
						Field177: []*int64{},
						Field178: map[string]*int64{
							"": nil,
						},
						Field179: []*int32{},
						Field180: []*string{},
						Field181: []*int32{},
						Field182: map[string]*string{
							"": nil,
						},
						Field183: []*int64{},
						Field184: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field185: []*int32{},
						Field186: nil,
						Field187: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field188: []*HugeStruct0{GetHugeStruct0()},
						Field189: nil,
						Field190: []*int64{},
						Field191: map[string]*int32{
							"": nil,
						},
						Field192: []*HugeStruct0{GetHugeStruct0()},
						Field193: []*HugeStruct0{GetHugeStruct0()},
						Field194: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field195: []*bool{},
						Field196: map[string]*bool{
							"": nil,
						},
						Field197: []*bool{},
						Field198: nil,
						Field199: map[string]*int32{
							"": nil,
						},
						Field200: map[string]*int64{
							"": nil,
						},
						Field201: map[string]*string{
							"": nil,
						},
						Field202: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field203: map[string]*int32{
							"": nil,
						},
						Field204: nil,
						Field205: map[string]*string{
							"": nil,
						},
						Field206: []*HugeStruct0{GetHugeStruct0()},
						Field207: []*HugeStruct0{GetHugeStruct0()},
						Field208: nil,
						Field209: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field210: map[string]*string{
							"": nil,
						},
						Field211: map[string]*bool{
							"": nil,
						},
						Field212: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field213: nil,
						Field214: map[string]*bool{
							"": nil,
						},
						Field215: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field216: []*HugeStruct0{GetHugeStruct0()},
						Field217: map[string]*string{
							"": nil,
						},
						Field218: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field219: map[string]*int64{
							"": nil,
						},
						Field220: nil,
						Field221: nil,
						Field222: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field223: []*int64{},
						Field224: []*bool{},
						Field225: []*bool{},
						Field226: map[string]*int64{
							"": nil,
						},
						Field227: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field228: []*int64{},
						Field229: map[string]*bool{
							"": nil,
						},
						Field230: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field231: nil,
						Field232: nil,
						Field233: []*string{},
						Field234: []*HugeStruct0{GetHugeStruct0()},
						Field235: []*string{},
						Field236: nil,
						Field237: nil,
						Field238: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field239: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field240: []*HugeStruct0{GetHugeStruct0()},
						Field241: nil,
						Field242: nil,
						Field243: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field244: map[string]*bool{
							"": nil,
						},
						Field245: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field246: []*int32{},
						Field247: []*bool{},
						Field248: []*string{},
						Field249: nil,
						Field250: []*int32{},
						Field251: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field252: nil,
						Field253: map[string]*string{
							"": nil,
						},
						Field254: map[string]*string{
							"": nil,
						},
						Field255: []*int32{},
						Field256: nil,
						Field257: nil,
						Field258: map[string]*string{
							"": nil,
						},
						Field259: map[string]*int32{
							"": nil,
						},
						Field260: []*int64{},
						Field261: []*int32{},
						Field262: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field263: nil,
						Field264: nil,
						Field265: map[string]*bool{
							"": nil,
						},
						Field266: nil,
						Field267: []*int64{},
						Field268: nil,
						Field269: nil,
						Field270: map[string]*int64{
							"": nil,
						},
						Field271: map[string]*int64{
							"": nil,
						},
						Field272: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field273: []*string{},
						Field274: nil,
						Field275: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field276: map[string]*bool{
							"": nil,
						},
						Field277: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field278: nil,
						Field279: map[string]*string{
							"": nil,
						},
						Field280: nil,
						Field281: nil,
						Field282: nil,
						Field283: nil,
						Field284: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field285: map[string]*int64{
							"": nil,
						},
						Field286: map[string]*bool{
							"": nil,
						},
						Field287: map[string]*string{
							"": nil,
						},
						Field288: nil,
						Field289: nil,
						Field290: nil,
						Field291: []*int64{},
						Field292: map[string]*string{
							"": nil,
						},
						Field293: nil,
						Field294: []*string{},
						Field295: nil,
						Field296: []*HugeStruct0{GetHugeStruct0()},
						Field297: nil,
						Field298: map[string]*int64{
							"": nil,
						},
						Field299: map[string]*bool{
							"": nil,
						},
						Field300: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field301: nil,
						Field302: []*string{},
						Field303: []*string{},
						Field304: map[string]*string{
							"": nil,
						},
						Field305: nil,
						Field306: nil,
						Field307: []*HugeStruct0{GetHugeStruct0()},
						Field308: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field309: map[string]*int32{
							"": nil,
						},
						Field310: []*HugeStruct0{GetHugeStruct0()},
						Field311: nil,
						Field312: []*bool{},
						Field313: nil,
						Field314: []*HugeStruct0{GetHugeStruct0()},
						Field315: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field316: nil,
						Field317: nil,
						Field318: nil,
						Field319: []*int32{},
						Field320: nil,
						Field321: []*HugeStruct0{GetHugeStruct0()},
						Field322: nil,
						Field323: nil,
						Field324: []*HugeStruct0{GetHugeStruct0()},
						Field325: nil,
						Field326: []*int64{},
						Field327: nil,
						Field328: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field329: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field330: []*HugeStruct0{GetHugeStruct0()},
						Field331: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field332: []*string{},
						Field333: nil,
						Field334: []*HugeStruct0{GetHugeStruct0()},
						Field335: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field336: map[string]*bool{
							"": nil,
						},
						Field337: []*int64{},
						Field338: map[string]*bool{
							"": nil,
						},
						Field339: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field340: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field341: []*bool{},
						Field342: []*int64{},
						Field343: []*int32{},
						Field344: map[string]*bool{
							"": nil,
						},
						Field345: map[string]*int64{
							"": nil,
						},
						Field346: nil,
						Field347: map[string]*bool{
							"": nil,
						},
						Field348: map[string]*int32{
							"": nil,
						},
						Field349: []*string{},
						Field350: map[string]*int32{
							"": nil,
						},
						Field351: nil,
						Field352: []*int64{},
						Field353: []*int64{},
						Field354: nil,
						Field355: map[string]*int32{
							"": nil,
						},
						Field356: map[string]*bool{
							"": nil,
						},
						Field357: []*int32{},
						Field358: nil,
						Field359: map[string]*int64{
							"": nil,
						},
						Field360: nil,
						Field361: map[string]*int64{
							"": nil,
						},
						Field362: map[string]*int32{
							"": nil,
						},
						Field363: []*int64{},
						Field364: []*bool{},
						Field365: nil,
						Field366: map[string]*string{
							"": nil,
						},
						Field367: map[string]*bool{
							"": nil,
						},
						Field368: nil,
						Field369: nil,
						Field370: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field371: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field372: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field373: map[string]*bool{
							"": nil,
						},
					},
				},
				Field63: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field64: []*int32{},
				Field65: []*HugeStruct0{GetHugeStruct0()},
				Field66: nil,
				Field67: []*int64{},
				Field68: []*bool{},
				Field69: nil,
				Field70: nil,
				Field71: nil,
				Field72: map[string]*int32{
					"": nil,
				},
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: map[string]*int32{
					"": nil,
				},
				Field75: map[string]*bool{
					"": nil,
				},
				Field76: nil,
				Field77: []*int32{},
				Field78: nil,
				Field79: nil,
				Field80: nil,
				Field81: []*bool{},
				Field82: map[string]*int64{
					"": nil,
				},
				Field83: nil,
				Field84: nil,
				Field85: map[string]*int32{
					"": nil,
				},
				Field86: nil,
				Field87: &HugeStruct1{
					Field0: []*int32{},
					Field1: []*string{},
					Field2: []*int64{},
					Field3: map[string]*int32{
						"": nil,
					},
					Field4: []*bool{},
					Field5: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field6: map[string]*int32{
						"": nil,
					},
					Field7: map[string]*bool{
						"": nil,
					},
					Field8: []*bool{},
					Field9: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field10: []*string{},
					Field11: []*bool{},
					Field12: []*bool{},
					Field13: map[string]*int32{
						"": nil,
					},
					Field14: map[string]*int32{
						"": nil,
					},
					Field15: nil,
					Field16: []*int64{},
					Field17: []*bool{},
					Field18: map[string]*int64{
						"": nil,
					},
					Field19: []*int64{},
					Field20: map[string]*string{
						"": nil,
					},
					Field21: nil,
					Field22: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field23: []*string{},
					Field24: []*int64{},
					Field25: []*string{},
					Field26: []*bool{},
					Field27: map[string]*int32{
						"": nil,
					},
					Field28: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field29: map[string]*int32{
						"": nil,
					},
					Field30: map[string]*bool{
						"": nil,
					},
					Field31: map[string]*int32{
						"": nil,
					},
					Field32: []*HugeStruct0{GetHugeStruct0()},
					Field33: nil,
					Field34: map[string]*bool{
						"": nil,
					},
					Field35: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field36: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field37: nil,
					Field38: []*HugeStruct0{GetHugeStruct0()},
					Field39: []*bool{},
					Field40: map[string]*string{
						"": nil,
					},
					Field41: map[string]*int64{
						"": nil,
					},
					Field42: map[string]*int32{
						"": nil,
					},
					Field43: nil,
					Field44: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field45: map[string]*int32{
						"": nil,
					},
					Field46: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field47: nil,
					Field48: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field49: nil,
					Field50: map[string]*string{
						"": nil,
					},
					Field51: map[string]*bool{
						"": nil,
					},
					Field52: []*int64{},
					Field53: map[string]*string{
						"": nil,
					},
					Field54: []*int32{},
					Field55: map[string]*int64{
						"": nil,
					},
					Field56: map[string]*int32{
						"": nil,
					},
					Field57: map[string]*string{
						"": nil,
					},
					Field58: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field59: []*HugeStruct0{GetHugeStruct0()},
					Field60: map[string]*string{
						"": nil,
					},
					Field61: map[string]*bool{
						"": nil,
					},
					Field62: map[string]*int64{
						"": nil,
					},
					Field63: []*string{},
					Field64: []*int64{},
					Field65: map[string]*bool{
						"": nil,
					},
					Field66: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field67: []*int64{},
					Field68: map[string]*string{
						"": nil,
					},
					Field69: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field70: []*bool{},
					Field71: map[string]*int64{
						"": nil,
					},
					Field72: nil,
					Field73: map[string]*int32{
						"": nil,
					},
					Field74: nil,
					Field75: map[string]*int32{
						"": nil,
					},
					Field76: map[string]*string{
						"": nil,
					},
					Field77: []*string{},
					Field78: nil,
					Field79: map[string]*int64{
						"": nil,
					},
					Field80: []*int64{},
					Field81: map[string]*bool{
						"": nil,
					},
					Field82: []*string{},
					Field83: []*string{},
					Field84: nil,
					Field85: []*bool{},
					Field86: []*HugeStruct0{GetHugeStruct0()},
					Field87: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field88: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field89: []*int64{},
					Field90: []*int32{},
					Field91: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field92: []*bool{},
					Field93: []*string{},
					Field94: map[string]*int32{
						"": nil,
					},
					Field95: nil,
					Field96: nil,
					Field97: map[string]*bool{
						"": nil,
					},
					Field98: map[string]*int32{
						"": nil,
					},
					Field99:  []*HugeStruct0{GetHugeStruct0()},
					Field100: nil,
					Field101: nil,
					Field102: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field103: []*string{},
					Field104: []*string{},
					Field105: map[string]*bool{
						"": nil,
					},
					Field106: []*string{},
					Field107: []*int64{},
					Field108: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field109: nil,
					Field110: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field111: []*string{},
					Field112: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field113: []*bool{},
					Field114: []*bool{},
					Field115: map[string]*string{
						"": nil,
					},
					Field116: []*int64{},
					Field117: []*string{},
					Field118: map[string]*bool{
						"": nil,
					},
					Field119: map[string]*string{
						"": nil,
					},
					Field120: []*HugeStruct0{GetHugeStruct0()},
					Field121: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field122: []*bool{},
					Field123: nil,
					Field124: []*int64{},
					Field125: nil,
					Field126: []*string{},
					Field127: []*string{},
					Field128: []*int32{},
					Field129: []*bool{},
					Field130: nil,
					Field131: nil,
					Field132: []*int32{},
					Field133: []*int32{},
					Field134: nil,
					Field135: []*bool{},
					Field136: nil,
					Field137: []*int32{},
					Field138: map[string]*int64{
						"": nil,
					},
					Field139: map[string]*string{
						"": nil,
					},
					Field140: map[string]*int64{
						"": nil,
					},
					Field141: map[string]*int64{
						"": nil,
					},
					Field142: []*int32{},
					Field143: []*HugeStruct0{GetHugeStruct0()},
					Field144: map[string]*int64{
						"": nil,
					},
					Field145: []*string{},
					Field146: map[string]*int64{
						"": nil,
					},
					Field147: nil,
					Field148: map[string]*string{
						"": nil,
					},
					Field149: nil,
					Field150: map[string]*int64{
						"": nil,
					},
					Field151: map[string]*int64{
						"": nil,
					},
					Field152: map[string]*int32{
						"": nil,
					},
					Field153: []*int32{},
					Field154: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field155: map[string]*string{
						"": nil,
					},
					Field156: map[string]*int64{
						"": nil,
					},
					Field157: []*int32{},
					Field158: []*int32{},
					Field159: nil,
					Field160: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field161: []*bool{},
					Field162: []*HugeStruct0{GetHugeStruct0()},
					Field163: []*int32{},
					Field164: map[string]*string{
						"": nil,
					},
					Field165: []*bool{},
					Field166: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field167: nil,
					Field168: []*bool{},
					Field169: map[string]*bool{
						"": nil,
					},
					Field170: map[string]*bool{
						"": nil,
					},
					Field171: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field172: map[string]*bool{
						"": nil,
					},
					Field173: []*bool{},
					Field174: map[string]*int64{
						"": nil,
					},
					Field175: []*HugeStruct0{GetHugeStruct0()},
					Field176: []*int32{},
					Field177: []*int64{},
					Field178: map[string]*int64{
						"": nil,
					},
					Field179: []*int32{},
					Field180: []*string{},
					Field181: []*int32{},
					Field182: map[string]*string{
						"": nil,
					},
					Field183: []*int64{},
					Field184: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field185: []*int32{},
					Field186: nil,
					Field187: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field188: []*HugeStruct0{GetHugeStruct0()},
					Field189: nil,
					Field190: []*int64{},
					Field191: map[string]*int32{
						"": nil,
					},
					Field192: []*HugeStruct0{GetHugeStruct0()},
					Field193: []*HugeStruct0{GetHugeStruct0()},
					Field194: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field195: []*bool{},
					Field196: map[string]*bool{
						"": nil,
					},
					Field197: []*bool{},
					Field198: nil,
					Field199: map[string]*int32{
						"": nil,
					},
					Field200: map[string]*int64{
						"": nil,
					},
					Field201: map[string]*string{
						"": nil,
					},
					Field202: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field203: map[string]*int32{
						"": nil,
					},
					Field204: nil,
					Field205: map[string]*string{
						"": nil,
					},
					Field206: []*HugeStruct0{GetHugeStruct0()},
					Field207: []*HugeStruct0{GetHugeStruct0()},
					Field208: nil,
					Field209: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field210: map[string]*string{
						"": nil,
					},
					Field211: map[string]*bool{
						"": nil,
					},
					Field212: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field213: nil,
					Field214: map[string]*bool{
						"": nil,
					},
					Field215: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field216: []*HugeStruct0{GetHugeStruct0()},
					Field217: map[string]*string{
						"": nil,
					},
					Field218: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field219: map[string]*int64{
						"": nil,
					},
					Field220: nil,
					Field221: nil,
					Field222: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field223: []*int64{},
					Field224: []*bool{},
					Field225: []*bool{},
					Field226: map[string]*int64{
						"": nil,
					},
					Field227: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field228: []*int64{},
					Field229: map[string]*bool{
						"": nil,
					},
					Field230: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field231: nil,
					Field232: nil,
					Field233: []*string{},
					Field234: []*HugeStruct0{GetHugeStruct0()},
					Field235: []*string{},
					Field236: nil,
					Field237: nil,
					Field238: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field239: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field240: []*HugeStruct0{GetHugeStruct0()},
					Field241: nil,
					Field242: nil,
					Field243: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field244: map[string]*bool{
						"": nil,
					},
					Field245: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field246: []*int32{},
					Field247: []*bool{},
					Field248: []*string{},
					Field249: nil,
					Field250: []*int32{},
					Field251: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field252: nil,
					Field253: map[string]*string{
						"": nil,
					},
					Field254: map[string]*string{
						"": nil,
					},
					Field255: []*int32{},
					Field256: nil,
					Field257: nil,
					Field258: map[string]*string{
						"": nil,
					},
					Field259: map[string]*int32{
						"": nil,
					},
					Field260: []*int64{},
					Field261: []*int32{},
					Field262: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field263: nil,
					Field264: nil,
					Field265: map[string]*bool{
						"": nil,
					},
					Field266: nil,
					Field267: []*int64{},
					Field268: nil,
					Field269: nil,
					Field270: map[string]*int64{
						"": nil,
					},
					Field271: map[string]*int64{
						"": nil,
					},
					Field272: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field273: []*string{},
					Field274: nil,
					Field275: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field276: map[string]*bool{
						"": nil,
					},
					Field277: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field278: nil,
					Field279: map[string]*string{
						"": nil,
					},
					Field280: nil,
					Field281: nil,
					Field282: nil,
					Field283: nil,
					Field284: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field285: map[string]*int64{
						"": nil,
					},
					Field286: map[string]*bool{
						"": nil,
					},
					Field287: map[string]*string{
						"": nil,
					},
					Field288: nil,
					Field289: nil,
					Field290: nil,
					Field291: []*int64{},
					Field292: map[string]*string{
						"": nil,
					},
					Field293: nil,
					Field294: []*string{},
					Field295: nil,
					Field296: []*HugeStruct0{GetHugeStruct0()},
					Field297: nil,
					Field298: map[string]*int64{
						"": nil,
					},
					Field299: map[string]*bool{
						"": nil,
					},
					Field300: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field301: nil,
					Field302: []*string{},
					Field303: []*string{},
					Field304: map[string]*string{
						"": nil,
					},
					Field305: nil,
					Field306: nil,
					Field307: []*HugeStruct0{GetHugeStruct0()},
					Field308: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field309: map[string]*int32{
						"": nil,
					},
					Field310: []*HugeStruct0{GetHugeStruct0()},
					Field311: nil,
					Field312: []*bool{},
					Field313: nil,
					Field314: []*HugeStruct0{GetHugeStruct0()},
					Field315: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field316: nil,
					Field317: nil,
					Field318: nil,
					Field319: []*int32{},
					Field320: nil,
					Field321: []*HugeStruct0{GetHugeStruct0()},
					Field322: nil,
					Field323: nil,
					Field324: []*HugeStruct0{GetHugeStruct0()},
					Field325: nil,
					Field326: []*int64{},
					Field327: nil,
					Field328: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field329: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field330: []*HugeStruct0{GetHugeStruct0()},
					Field331: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field332: []*string{},
					Field333: nil,
					Field334: []*HugeStruct0{GetHugeStruct0()},
					Field335: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field336: map[string]*bool{
						"": nil,
					},
					Field337: []*int64{},
					Field338: map[string]*bool{
						"": nil,
					},
					Field339: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field340: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field341: []*bool{},
					Field342: []*int64{},
					Field343: []*int32{},
					Field344: map[string]*bool{
						"": nil,
					},
					Field345: map[string]*int64{
						"": nil,
					},
					Field346: nil,
					Field347: map[string]*bool{
						"": nil,
					},
					Field348: map[string]*int32{
						"": nil,
					},
					Field349: []*string{},
					Field350: map[string]*int32{
						"": nil,
					},
					Field351: nil,
					Field352: []*int64{},
					Field353: []*int64{},
					Field354: nil,
					Field355: map[string]*int32{
						"": nil,
					},
					Field356: map[string]*bool{
						"": nil,
					},
					Field357: []*int32{},
					Field358: nil,
					Field359: map[string]*int64{
						"": nil,
					},
					Field360: nil,
					Field361: map[string]*int64{
						"": nil,
					},
					Field362: map[string]*int32{
						"": nil,
					},
					Field363: []*int64{},
					Field364: []*bool{},
					Field365: nil,
					Field366: map[string]*string{
						"": nil,
					},
					Field367: map[string]*bool{
						"": nil,
					},
					Field368: nil,
					Field369: nil,
					Field370: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field371: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field372: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field373: map[string]*bool{
						"": nil,
					},
				},
				Field88: []*int32{},
				Field89: nil,
				Field90: []*bool{},
				Field91: []*bool{},
				Field92: &HugeStruct1{
					Field0: []*int32{},
					Field1: []*string{},
					Field2: []*int64{},
					Field3: map[string]*int32{
						"": nil,
					},
					Field4: []*bool{},
					Field5: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field6: map[string]*int32{
						"": nil,
					},
					Field7: map[string]*bool{
						"": nil,
					},
					Field8: []*bool{},
					Field9: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field10: []*string{},
					Field11: []*bool{},
					Field12: []*bool{},
					Field13: map[string]*int32{
						"": nil,
					},
					Field14: map[string]*int32{
						"": nil,
					},
					Field15: nil,
					Field16: []*int64{},
					Field17: []*bool{},
					Field18: map[string]*int64{
						"": nil,
					},
					Field19: []*int64{},
					Field20: map[string]*string{
						"": nil,
					},
					Field21: nil,
					Field22: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field23: []*string{},
					Field24: []*int64{},
					Field25: []*string{},
					Field26: []*bool{},
					Field27: map[string]*int32{
						"": nil,
					},
					Field28: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field29: map[string]*int32{
						"": nil,
					},
					Field30: map[string]*bool{
						"": nil,
					},
					Field31: map[string]*int32{
						"": nil,
					},
					Field32: []*HugeStruct0{GetHugeStruct0()},
					Field33: nil,
					Field34: map[string]*bool{
						"": nil,
					},
					Field35: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field36: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field37: nil,
					Field38: []*HugeStruct0{GetHugeStruct0()},
					Field39: []*bool{},
					Field40: map[string]*string{
						"": nil,
					},
					Field41: map[string]*int64{
						"": nil,
					},
					Field42: map[string]*int32{
						"": nil,
					},
					Field43: nil,
					Field44: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field45: map[string]*int32{
						"": nil,
					},
					Field46: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field47: nil,
					Field48: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field49: nil,
					Field50: map[string]*string{
						"": nil,
					},
					Field51: map[string]*bool{
						"": nil,
					},
					Field52: []*int64{},
					Field53: map[string]*string{
						"": nil,
					},
					Field54: []*int32{},
					Field55: map[string]*int64{
						"": nil,
					},
					Field56: map[string]*int32{
						"": nil,
					},
					Field57: map[string]*string{
						"": nil,
					},
					Field58: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field59: []*HugeStruct0{GetHugeStruct0()},
					Field60: map[string]*string{
						"": nil,
					},
					Field61: map[string]*bool{
						"": nil,
					},
					Field62: map[string]*int64{
						"": nil,
					},
					Field63: []*string{},
					Field64: []*int64{},
					Field65: map[string]*bool{
						"": nil,
					},
					Field66: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field67: []*int64{},
					Field68: map[string]*string{
						"": nil,
					},
					Field69: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field70: []*bool{},
					Field71: map[string]*int64{
						"": nil,
					},
					Field72: nil,
					Field73: map[string]*int32{
						"": nil,
					},
					Field74: nil,
					Field75: map[string]*int32{
						"": nil,
					},
					Field76: map[string]*string{
						"": nil,
					},
					Field77: []*string{},
					Field78: nil,
					Field79: map[string]*int64{
						"": nil,
					},
					Field80: []*int64{},
					Field81: map[string]*bool{
						"": nil,
					},
					Field82: []*string{},
					Field83: []*string{},
					Field84: nil,
					Field85: []*bool{},
					Field86: []*HugeStruct0{GetHugeStruct0()},
					Field87: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field88: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field89: []*int64{},
					Field90: []*int32{},
					Field91: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field92: []*bool{},
					Field93: []*string{},
					Field94: map[string]*int32{
						"": nil,
					},
					Field95: nil,
					Field96: nil,
					Field97: map[string]*bool{
						"": nil,
					},
					Field98: map[string]*int32{
						"": nil,
					},
					Field99:  []*HugeStruct0{GetHugeStruct0()},
					Field100: nil,
					Field101: nil,
					Field102: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field103: []*string{},
					Field104: []*string{},
					Field105: map[string]*bool{
						"": nil,
					},
					Field106: []*string{},
					Field107: []*int64{},
					Field108: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field109: nil,
					Field110: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field111: []*string{},
					Field112: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field113: []*bool{},
					Field114: []*bool{},
					Field115: map[string]*string{
						"": nil,
					},
					Field116: []*int64{},
					Field117: []*string{},
					Field118: map[string]*bool{
						"": nil,
					},
					Field119: map[string]*string{
						"": nil,
					},
					Field120: []*HugeStruct0{GetHugeStruct0()},
					Field121: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field122: []*bool{},
					Field123: nil,
					Field124: []*int64{},
					Field125: nil,
					Field126: []*string{},
					Field127: []*string{},
					Field128: []*int32{},
					Field129: []*bool{},
					Field130: nil,
					Field131: nil,
					Field132: []*int32{},
					Field133: []*int32{},
					Field134: nil,
					Field135: []*bool{},
					Field136: nil,
					Field137: []*int32{},
					Field138: map[string]*int64{
						"": nil,
					},
					Field139: map[string]*string{
						"": nil,
					},
					Field140: map[string]*int64{
						"": nil,
					},
					Field141: map[string]*int64{
						"": nil,
					},
					Field142: []*int32{},
					Field143: []*HugeStruct0{GetHugeStruct0()},
					Field144: map[string]*int64{
						"": nil,
					},
					Field145: []*string{},
					Field146: map[string]*int64{
						"": nil,
					},
					Field147: nil,
					Field148: map[string]*string{
						"": nil,
					},
					Field149: nil,
					Field150: map[string]*int64{
						"": nil,
					},
					Field151: map[string]*int64{
						"": nil,
					},
					Field152: map[string]*int32{
						"": nil,
					},
					Field153: []*int32{},
					Field154: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field155: map[string]*string{
						"": nil,
					},
					Field156: map[string]*int64{
						"": nil,
					},
					Field157: []*int32{},
					Field158: []*int32{},
					Field159: nil,
					Field160: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field161: []*bool{},
					Field162: []*HugeStruct0{GetHugeStruct0()},
					Field163: []*int32{},
					Field164: map[string]*string{
						"": nil,
					},
					Field165: []*bool{},
					Field166: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field167: nil,
					Field168: []*bool{},
					Field169: map[string]*bool{
						"": nil,
					},
					Field170: map[string]*bool{
						"": nil,
					},
					Field171: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field172: map[string]*bool{
						"": nil,
					},
					Field173: []*bool{},
					Field174: map[string]*int64{
						"": nil,
					},
					Field175: []*HugeStruct0{GetHugeStruct0()},
					Field176: []*int32{},
					Field177: []*int64{},
					Field178: map[string]*int64{
						"": nil,
					},
					Field179: []*int32{},
					Field180: []*string{},
					Field181: []*int32{},
					Field182: map[string]*string{
						"": nil,
					},
					Field183: []*int64{},
					Field184: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field185: []*int32{},
					Field186: nil,
					Field187: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field188: []*HugeStruct0{GetHugeStruct0()},
					Field189: nil,
					Field190: []*int64{},
					Field191: map[string]*int32{
						"": nil,
					},
					Field192: []*HugeStruct0{GetHugeStruct0()},
					Field193: []*HugeStruct0{GetHugeStruct0()},
					Field194: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field195: []*bool{},
					Field196: map[string]*bool{
						"": nil,
					},
					Field197: []*bool{},
					Field198: nil,
					Field199: map[string]*int32{
						"": nil,
					},
					Field200: map[string]*int64{
						"": nil,
					},
					Field201: map[string]*string{
						"": nil,
					},
					Field202: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field203: map[string]*int32{
						"": nil,
					},
					Field204: nil,
					Field205: map[string]*string{
						"": nil,
					},
					Field206: []*HugeStruct0{GetHugeStruct0()},
					Field207: []*HugeStruct0{GetHugeStruct0()},
					Field208: nil,
					Field209: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field210: map[string]*string{
						"": nil,
					},
					Field211: map[string]*bool{
						"": nil,
					},
					Field212: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field213: nil,
					Field214: map[string]*bool{
						"": nil,
					},
					Field215: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field216: []*HugeStruct0{GetHugeStruct0()},
					Field217: map[string]*string{
						"": nil,
					},
					Field218: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field219: map[string]*int64{
						"": nil,
					},
					Field220: nil,
					Field221: nil,
					Field222: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field223: []*int64{},
					Field224: []*bool{},
					Field225: []*bool{},
					Field226: map[string]*int64{
						"": nil,
					},
					Field227: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field228: []*int64{},
					Field229: map[string]*bool{
						"": nil,
					},
					Field230: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field231: nil,
					Field232: nil,
					Field233: []*string{},
					Field234: []*HugeStruct0{GetHugeStruct0()},
					Field235: []*string{},
					Field236: nil,
					Field237: nil,
					Field238: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field239: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field240: []*HugeStruct0{GetHugeStruct0()},
					Field241: nil,
					Field242: nil,
					Field243: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field244: map[string]*bool{
						"": nil,
					},
					Field245: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field246: []*int32{},
					Field247: []*bool{},
					Field248: []*string{},
					Field249: nil,
					Field250: []*int32{},
					Field251: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field252: nil,
					Field253: map[string]*string{
						"": nil,
					},
					Field254: map[string]*string{
						"": nil,
					},
					Field255: []*int32{},
					Field256: nil,
					Field257: nil,
					Field258: map[string]*string{
						"": nil,
					},
					Field259: map[string]*int32{
						"": nil,
					},
					Field260: []*int64{},
					Field261: []*int32{},
					Field262: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field263: nil,
					Field264: nil,
					Field265: map[string]*bool{
						"": nil,
					},
					Field266: nil,
					Field267: []*int64{},
					Field268: nil,
					Field269: nil,
					Field270: map[string]*int64{
						"": nil,
					},
					Field271: map[string]*int64{
						"": nil,
					},
					Field272: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field273: []*string{},
					Field274: nil,
					Field275: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field276: map[string]*bool{
						"": nil,
					},
					Field277: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field278: nil,
					Field279: map[string]*string{
						"": nil,
					},
					Field280: nil,
					Field281: nil,
					Field282: nil,
					Field283: nil,
					Field284: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field285: map[string]*int64{
						"": nil,
					},
					Field286: map[string]*bool{
						"": nil,
					},
					Field287: map[string]*string{
						"": nil,
					},
					Field288: nil,
					Field289: nil,
					Field290: nil,
					Field291: []*int64{},
					Field292: map[string]*string{
						"": nil,
					},
					Field293: nil,
					Field294: []*string{},
					Field295: nil,
					Field296: []*HugeStruct0{GetHugeStruct0()},
					Field297: nil,
					Field298: map[string]*int64{
						"": nil,
					},
					Field299: map[string]*bool{
						"": nil,
					},
					Field300: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field301: nil,
					Field302: []*string{},
					Field303: []*string{},
					Field304: map[string]*string{
						"": nil,
					},
					Field305: nil,
					Field306: nil,
					Field307: []*HugeStruct0{GetHugeStruct0()},
					Field308: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field309: map[string]*int32{
						"": nil,
					},
					Field310: []*HugeStruct0{GetHugeStruct0()},
					Field311: nil,
					Field312: []*bool{},
					Field313: nil,
					Field314: []*HugeStruct0{GetHugeStruct0()},
					Field315: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field316: nil,
					Field317: nil,
					Field318: nil,
					Field319: []*int32{},
					Field320: nil,
					Field321: []*HugeStruct0{GetHugeStruct0()},
					Field322: nil,
					Field323: nil,
					Field324: []*HugeStruct0{GetHugeStruct0()},
					Field325: nil,
					Field326: []*int64{},
					Field327: nil,
					Field328: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field329: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field330: []*HugeStruct0{GetHugeStruct0()},
					Field331: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field332: []*string{},
					Field333: nil,
					Field334: []*HugeStruct0{GetHugeStruct0()},
					Field335: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field336: map[string]*bool{
						"": nil,
					},
					Field337: []*int64{},
					Field338: map[string]*bool{
						"": nil,
					},
					Field339: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field340: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field341: []*bool{},
					Field342: []*int64{},
					Field343: []*int32{},
					Field344: map[string]*bool{
						"": nil,
					},
					Field345: map[string]*int64{
						"": nil,
					},
					Field346: nil,
					Field347: map[string]*bool{
						"": nil,
					},
					Field348: map[string]*int32{
						"": nil,
					},
					Field349: []*string{},
					Field350: map[string]*int32{
						"": nil,
					},
					Field351: nil,
					Field352: []*int64{},
					Field353: []*int64{},
					Field354: nil,
					Field355: map[string]*int32{
						"": nil,
					},
					Field356: map[string]*bool{
						"": nil,
					},
					Field357: []*int32{},
					Field358: nil,
					Field359: map[string]*int64{
						"": nil,
					},
					Field360: nil,
					Field361: map[string]*int64{
						"": nil,
					},
					Field362: map[string]*int32{
						"": nil,
					},
					Field363: []*int64{},
					Field364: []*bool{},
					Field365: nil,
					Field366: map[string]*string{
						"": nil,
					},
					Field367: map[string]*bool{
						"": nil,
					},
					Field368: nil,
					Field369: nil,
					Field370: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field371: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field372: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field373: map[string]*bool{
						"": nil,
					},
				},
				Field93: nil,
				Field94: &HugeStruct1{
					Field0: []*int32{},
					Field1: []*string{},
					Field2: []*int64{},
					Field3: map[string]*int32{
						"": nil,
					},
					Field4: []*bool{},
					Field5: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field6: map[string]*int32{
						"": nil,
					},
					Field7: map[string]*bool{
						"": nil,
					},
					Field8: []*bool{},
					Field9: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field10: []*string{},
					Field11: []*bool{},
					Field12: []*bool{},
					Field13: map[string]*int32{
						"": nil,
					},
					Field14: map[string]*int32{
						"": nil,
					},
					Field15: nil,
					Field16: []*int64{},
					Field17: []*bool{},
					Field18: map[string]*int64{
						"": nil,
					},
					Field19: []*int64{},
					Field20: map[string]*string{
						"": nil,
					},
					Field21: nil,
					Field22: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field23: []*string{},
					Field24: []*int64{},
					Field25: []*string{},
					Field26: []*bool{},
					Field27: map[string]*int32{
						"": nil,
					},
					Field28: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field29: map[string]*int32{
						"": nil,
					},
					Field30: map[string]*bool{
						"": nil,
					},
					Field31: map[string]*int32{
						"": nil,
					},
					Field32: []*HugeStruct0{GetHugeStruct0()},
					Field33: nil,
					Field34: map[string]*bool{
						"": nil,
					},
					Field35: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field36: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field37: nil,
					Field38: []*HugeStruct0{GetHugeStruct0()},
					Field39: []*bool{},
					Field40: map[string]*string{
						"": nil,
					},
					Field41: map[string]*int64{
						"": nil,
					},
					Field42: map[string]*int32{
						"": nil,
					},
					Field43: nil,
					Field44: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field45: map[string]*int32{
						"": nil,
					},
					Field46: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field47: nil,
					Field48: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field49: nil,
					Field50: map[string]*string{
						"": nil,
					},
					Field51: map[string]*bool{
						"": nil,
					},
					Field52: []*int64{},
					Field53: map[string]*string{
						"": nil,
					},
					Field54: []*int32{},
					Field55: map[string]*int64{
						"": nil,
					},
					Field56: map[string]*int32{
						"": nil,
					},
					Field57: map[string]*string{
						"": nil,
					},
					Field58: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field59: []*HugeStruct0{GetHugeStruct0()},
					Field60: map[string]*string{
						"": nil,
					},
					Field61: map[string]*bool{
						"": nil,
					},
					Field62: map[string]*int64{
						"": nil,
					},
					Field63: []*string{},
					Field64: []*int64{},
					Field65: map[string]*bool{
						"": nil,
					},
					Field66: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field67: []*int64{},
					Field68: map[string]*string{
						"": nil,
					},
					Field69: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field70: []*bool{},
					Field71: map[string]*int64{
						"": nil,
					},
					Field72: nil,
					Field73: map[string]*int32{
						"": nil,
					},
					Field74: nil,
					Field75: map[string]*int32{
						"": nil,
					},
					Field76: map[string]*string{
						"": nil,
					},
					Field77: []*string{},
					Field78: nil,
					Field79: map[string]*int64{
						"": nil,
					},
					Field80: []*int64{},
					Field81: map[string]*bool{
						"": nil,
					},
					Field82: []*string{},
					Field83: []*string{},
					Field84: nil,
					Field85: []*bool{},
					Field86: []*HugeStruct0{GetHugeStruct0()},
					Field87: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field88: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field89: []*int64{},
					Field90: []*int32{},
					Field91: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field92: []*bool{},
					Field93: []*string{},
					Field94: map[string]*int32{
						"": nil,
					},
					Field95: nil,
					Field96: nil,
					Field97: map[string]*bool{
						"": nil,
					},
					Field98: map[string]*int32{
						"": nil,
					},
					Field99:  []*HugeStruct0{GetHugeStruct0()},
					Field100: nil,
					Field101: nil,
					Field102: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field103: []*string{},
					Field104: []*string{},
					Field105: map[string]*bool{
						"": nil,
					},
					Field106: []*string{},
					Field107: []*int64{},
					Field108: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field109: nil,
					Field110: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field111: []*string{},
					Field112: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field113: []*bool{},
					Field114: []*bool{},
					Field115: map[string]*string{
						"": nil,
					},
					Field116: []*int64{},
					Field117: []*string{},
					Field118: map[string]*bool{
						"": nil,
					},
					Field119: map[string]*string{
						"": nil,
					},
					Field120: []*HugeStruct0{GetHugeStruct0()},
					Field121: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field122: []*bool{},
					Field123: nil,
					Field124: []*int64{},
					Field125: nil,
					Field126: []*string{},
					Field127: []*string{},
					Field128: []*int32{},
					Field129: []*bool{},
					Field130: nil,
					Field131: nil,
					Field132: []*int32{},
					Field133: []*int32{},
					Field134: nil,
					Field135: []*bool{},
					Field136: nil,
					Field137: []*int32{},
					Field138: map[string]*int64{
						"": nil,
					},
					Field139: map[string]*string{
						"": nil,
					},
					Field140: map[string]*int64{
						"": nil,
					},
					Field141: map[string]*int64{
						"": nil,
					},
					Field142: []*int32{},
					Field143: []*HugeStruct0{GetHugeStruct0()},
					Field144: map[string]*int64{
						"": nil,
					},
					Field145: []*string{},
					Field146: map[string]*int64{
						"": nil,
					},
					Field147: nil,
					Field148: map[string]*string{
						"": nil,
					},
					Field149: nil,
					Field150: map[string]*int64{
						"": nil,
					},
					Field151: map[string]*int64{
						"": nil,
					},
					Field152: map[string]*int32{
						"": nil,
					},
					Field153: []*int32{},
					Field154: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field155: map[string]*string{
						"": nil,
					},
					Field156: map[string]*int64{
						"": nil,
					},
					Field157: []*int32{},
					Field158: []*int32{},
					Field159: nil,
					Field160: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field161: []*bool{},
					Field162: []*HugeStruct0{GetHugeStruct0()},
					Field163: []*int32{},
					Field164: map[string]*string{
						"": nil,
					},
					Field165: []*bool{},
					Field166: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field167: nil,
					Field168: []*bool{},
					Field169: map[string]*bool{
						"": nil,
					},
					Field170: map[string]*bool{
						"": nil,
					},
					Field171: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field172: map[string]*bool{
						"": nil,
					},
					Field173: []*bool{},
					Field174: map[string]*int64{
						"": nil,
					},
					Field175: []*HugeStruct0{GetHugeStruct0()},
					Field176: []*int32{},
					Field177: []*int64{},
					Field178: map[string]*int64{
						"": nil,
					},
					Field179: []*int32{},
					Field180: []*string{},
					Field181: []*int32{},
					Field182: map[string]*string{
						"": nil,
					},
					Field183: []*int64{},
					Field184: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field185: []*int32{},
					Field186: nil,
					Field187: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field188: []*HugeStruct0{GetHugeStruct0()},
					Field189: nil,
					Field190: []*int64{},
					Field191: map[string]*int32{
						"": nil,
					},
					Field192: []*HugeStruct0{GetHugeStruct0()},
					Field193: []*HugeStruct0{GetHugeStruct0()},
					Field194: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field195: []*bool{},
					Field196: map[string]*bool{
						"": nil,
					},
					Field197: []*bool{},
					Field198: nil,
					Field199: map[string]*int32{
						"": nil,
					},
					Field200: map[string]*int64{
						"": nil,
					},
					Field201: map[string]*string{
						"": nil,
					},
					Field202: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field203: map[string]*int32{
						"": nil,
					},
					Field204: nil,
					Field205: map[string]*string{
						"": nil,
					},
					Field206: []*HugeStruct0{GetHugeStruct0()},
					Field207: []*HugeStruct0{GetHugeStruct0()},
					Field208: nil,
					Field209: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field210: map[string]*string{
						"": nil,
					},
					Field211: map[string]*bool{
						"": nil,
					},
					Field212: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field213: nil,
					Field214: map[string]*bool{
						"": nil,
					},
					Field215: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field216: []*HugeStruct0{GetHugeStruct0()},
					Field217: map[string]*string{
						"": nil,
					},
					Field218: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field219: map[string]*int64{
						"": nil,
					},
					Field220: nil,
					Field221: nil,
					Field222: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field223: []*int64{},
					Field224: []*bool{},
					Field225: []*bool{},
					Field226: map[string]*int64{
						"": nil,
					},
					Field227: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field228: []*int64{},
					Field229: map[string]*bool{
						"": nil,
					},
					Field230: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field231: nil,
					Field232: nil,
					Field233: []*string{},
					Field234: []*HugeStruct0{GetHugeStruct0()},
					Field235: []*string{},
					Field236: nil,
					Field237: nil,
					Field238: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field239: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field240: []*HugeStruct0{GetHugeStruct0()},
					Field241: nil,
					Field242: nil,
					Field243: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field244: map[string]*bool{
						"": nil,
					},
					Field245: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field246: []*int32{},
					Field247: []*bool{},
					Field248: []*string{},
					Field249: nil,
					Field250: []*int32{},
					Field251: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field252: nil,
					Field253: map[string]*string{
						"": nil,
					},
					Field254: map[string]*string{
						"": nil,
					},
					Field255: []*int32{},
					Field256: nil,
					Field257: nil,
					Field258: map[string]*string{
						"": nil,
					},
					Field259: map[string]*int32{
						"": nil,
					},
					Field260: []*int64{},
					Field261: []*int32{},
					Field262: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field263: nil,
					Field264: nil,
					Field265: map[string]*bool{
						"": nil,
					},
					Field266: nil,
					Field267: []*int64{},
					Field268: nil,
					Field269: nil,
					Field270: map[string]*int64{
						"": nil,
					},
					Field271: map[string]*int64{
						"": nil,
					},
					Field272: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field273: []*string{},
					Field274: nil,
					Field275: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field276: map[string]*bool{
						"": nil,
					},
					Field277: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field278: nil,
					Field279: map[string]*string{
						"": nil,
					},
					Field280: nil,
					Field281: nil,
					Field282: nil,
					Field283: nil,
					Field284: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field285: map[string]*int64{
						"": nil,
					},
					Field286: map[string]*bool{
						"": nil,
					},
					Field287: map[string]*string{
						"": nil,
					},
					Field288: nil,
					Field289: nil,
					Field290: nil,
					Field291: []*int64{},
					Field292: map[string]*string{
						"": nil,
					},
					Field293: nil,
					Field294: []*string{},
					Field295: nil,
					Field296: []*HugeStruct0{GetHugeStruct0()},
					Field297: nil,
					Field298: map[string]*int64{
						"": nil,
					},
					Field299: map[string]*bool{
						"": nil,
					},
					Field300: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field301: nil,
					Field302: []*string{},
					Field303: []*string{},
					Field304: map[string]*string{
						"": nil,
					},
					Field305: nil,
					Field306: nil,
					Field307: []*HugeStruct0{GetHugeStruct0()},
					Field308: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field309: map[string]*int32{
						"": nil,
					},
					Field310: []*HugeStruct0{GetHugeStruct0()},
					Field311: nil,
					Field312: []*bool{},
					Field313: nil,
					Field314: []*HugeStruct0{GetHugeStruct0()},
					Field315: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field316: nil,
					Field317: nil,
					Field318: nil,
					Field319: []*int32{},
					Field320: nil,
					Field321: []*HugeStruct0{GetHugeStruct0()},
					Field322: nil,
					Field323: nil,
					Field324: []*HugeStruct0{GetHugeStruct0()},
					Field325: nil,
					Field326: []*int64{},
					Field327: nil,
					Field328: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field329: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field330: []*HugeStruct0{GetHugeStruct0()},
					Field331: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field332: []*string{},
					Field333: nil,
					Field334: []*HugeStruct0{GetHugeStruct0()},
					Field335: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field336: map[string]*bool{
						"": nil,
					},
					Field337: []*int64{},
					Field338: map[string]*bool{
						"": nil,
					},
					Field339: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field340: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field341: []*bool{},
					Field342: []*int64{},
					Field343: []*int32{},
					Field344: map[string]*bool{
						"": nil,
					},
					Field345: map[string]*int64{
						"": nil,
					},
					Field346: nil,
					Field347: map[string]*bool{
						"": nil,
					},
					Field348: map[string]*int32{
						"": nil,
					},
					Field349: []*string{},
					Field350: map[string]*int32{
						"": nil,
					},
					Field351: nil,
					Field352: []*int64{},
					Field353: []*int64{},
					Field354: nil,
					Field355: map[string]*int32{
						"": nil,
					},
					Field356: map[string]*bool{
						"": nil,
					},
					Field357: []*int32{},
					Field358: nil,
					Field359: map[string]*int64{
						"": nil,
					},
					Field360: nil,
					Field361: map[string]*int64{
						"": nil,
					},
					Field362: map[string]*int32{
						"": nil,
					},
					Field363: []*int64{},
					Field364: []*bool{},
					Field365: nil,
					Field366: map[string]*string{
						"": nil,
					},
					Field367: map[string]*bool{
						"": nil,
					},
					Field368: nil,
					Field369: nil,
					Field370: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field371: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field372: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field373: map[string]*bool{
						"": nil,
					},
				},
				Field95: map[string]*int32{
					"": nil,
				},
				Field96: nil,
				Field97: []*HugeStruct0{GetHugeStruct0()},
				Field98: []*bool{},
				Field99: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field100: []*int32{},
				Field101: nil,
				Field102: map[string]*bool{
					"": nil,
				},
				Field103: map[string]*bool{
					"": nil,
				},
				Field104: []*string{},
				Field105: map[string]*int32{
					"": nil,
				},
				Field106: nil,
				Field107: map[string]*HugeStruct1{
					"": {
						Field0: []*int32{},
						Field1: []*string{},
						Field2: []*int64{},
						Field3: map[string]*int32{
							"": nil,
						},
						Field4: []*bool{},
						Field5: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field6: map[string]*int32{
							"": nil,
						},
						Field7: map[string]*bool{
							"": nil,
						},
						Field8: []*bool{},
						Field9: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field10: []*string{},
						Field11: []*bool{},
						Field12: []*bool{},
						Field13: map[string]*int32{
							"": nil,
						},
						Field14: map[string]*int32{
							"": nil,
						},
						Field15: nil,
						Field16: []*int64{},
						Field17: []*bool{},
						Field18: map[string]*int64{
							"": nil,
						},
						Field19: []*int64{},
						Field20: map[string]*string{
							"": nil,
						},
						Field21: nil,
						Field22: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field23: []*string{},
						Field24: []*int64{},
						Field25: []*string{},
						Field26: []*bool{},
						Field27: map[string]*int32{
							"": nil,
						},
						Field28: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field29: map[string]*int32{
							"": nil,
						},
						Field30: map[string]*bool{
							"": nil,
						},
						Field31: map[string]*int32{
							"": nil,
						},
						Field32: []*HugeStruct0{GetHugeStruct0()},
						Field33: nil,
						Field34: map[string]*bool{
							"": nil,
						},
						Field35: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field36: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field37: nil,
						Field38: []*HugeStruct0{GetHugeStruct0()},
						Field39: []*bool{},
						Field40: map[string]*string{
							"": nil,
						},
						Field41: map[string]*int64{
							"": nil,
						},
						Field42: map[string]*int32{
							"": nil,
						},
						Field43: nil,
						Field44: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field45: map[string]*int32{
							"": nil,
						},
						Field46: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field47: nil,
						Field48: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field49: nil,
						Field50: map[string]*string{
							"": nil,
						},
						Field51: map[string]*bool{
							"": nil,
						},
						Field52: []*int64{},
						Field53: map[string]*string{
							"": nil,
						},
						Field54: []*int32{},
						Field55: map[string]*int64{
							"": nil,
						},
						Field56: map[string]*int32{
							"": nil,
						},
						Field57: map[string]*string{
							"": nil,
						},
						Field58: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field59: []*HugeStruct0{GetHugeStruct0()},
						Field60: map[string]*string{
							"": nil,
						},
						Field61: map[string]*bool{
							"": nil,
						},
						Field62: map[string]*int64{
							"": nil,
						},
						Field63: []*string{},
						Field64: []*int64{},
						Field65: map[string]*bool{
							"": nil,
						},
						Field66: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field67: []*int64{},
						Field68: map[string]*string{
							"": nil,
						},
						Field69: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field70: []*bool{},
						Field71: map[string]*int64{
							"": nil,
						},
						Field72: nil,
						Field73: map[string]*int32{
							"": nil,
						},
						Field74: nil,
						Field75: map[string]*int32{
							"": nil,
						},
						Field76: map[string]*string{
							"": nil,
						},
						Field77: []*string{},
						Field78: nil,
						Field79: map[string]*int64{
							"": nil,
						},
						Field80: []*int64{},
						Field81: map[string]*bool{
							"": nil,
						},
						Field82: []*string{},
						Field83: []*string{},
						Field84: nil,
						Field85: []*bool{},
						Field86: []*HugeStruct0{GetHugeStruct0()},
						Field87: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field88: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field89: []*int64{},
						Field90: []*int32{},
						Field91: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field92: []*bool{},
						Field93: []*string{},
						Field94: map[string]*int32{
							"": nil,
						},
						Field95: nil,
						Field96: nil,
						Field97: map[string]*bool{
							"": nil,
						},
						Field98: map[string]*int32{
							"": nil,
						},
						Field99:  []*HugeStruct0{GetHugeStruct0()},
						Field100: nil,
						Field101: nil,
						Field102: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field103: []*string{},
						Field104: []*string{},
						Field105: map[string]*bool{
							"": nil,
						},
						Field106: []*string{},
						Field107: []*int64{},
						Field108: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field109: nil,
						Field110: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field111: []*string{},
						Field112: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field113: []*bool{},
						Field114: []*bool{},
						Field115: map[string]*string{
							"": nil,
						},
						Field116: []*int64{},
						Field117: []*string{},
						Field118: map[string]*bool{
							"": nil,
						},
						Field119: map[string]*string{
							"": nil,
						},
						Field120: []*HugeStruct0{GetHugeStruct0()},
						Field121: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field122: []*bool{},
						Field123: nil,
						Field124: []*int64{},
						Field125: nil,
						Field126: []*string{},
						Field127: []*string{},
						Field128: []*int32{},
						Field129: []*bool{},
						Field130: nil,
						Field131: nil,
						Field132: []*int32{},
						Field133: []*int32{},
						Field134: nil,
						Field135: []*bool{},
						Field136: nil,
						Field137: []*int32{},
						Field138: map[string]*int64{
							"": nil,
						},
						Field139: map[string]*string{
							"": nil,
						},
						Field140: map[string]*int64{
							"": nil,
						},
						Field141: map[string]*int64{
							"": nil,
						},
						Field142: []*int32{},
						Field143: []*HugeStruct0{GetHugeStruct0()},
						Field144: map[string]*int64{
							"": nil,
						},
						Field145: []*string{},
						Field146: map[string]*int64{
							"": nil,
						},
						Field147: nil,
						Field148: map[string]*string{
							"": nil,
						},
						Field149: nil,
						Field150: map[string]*int64{
							"": nil,
						},
						Field151: map[string]*int64{
							"": nil,
						},
						Field152: map[string]*int32{
							"": nil,
						},
						Field153: []*int32{},
						Field154: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field155: map[string]*string{
							"": nil,
						},
						Field156: map[string]*int64{
							"": nil,
						},
						Field157: []*int32{},
						Field158: []*int32{},
						Field159: nil,
						Field160: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field161: []*bool{},
						Field162: []*HugeStruct0{GetHugeStruct0()},
						Field163: []*int32{},
						Field164: map[string]*string{
							"": nil,
						},
						Field165: []*bool{},
						Field166: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field167: nil,
						Field168: []*bool{},
						Field169: map[string]*bool{
							"": nil,
						},
						Field170: map[string]*bool{
							"": nil,
						},
						Field171: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field172: map[string]*bool{
							"": nil,
						},
						Field173: []*bool{},
						Field174: map[string]*int64{
							"": nil,
						},
						Field175: []*HugeStruct0{GetHugeStruct0()},
						Field176: []*int32{},
						Field177: []*int64{},
						Field178: map[string]*int64{
							"": nil,
						},
						Field179: []*int32{},
						Field180: []*string{},
						Field181: []*int32{},
						Field182: map[string]*string{
							"": nil,
						},
						Field183: []*int64{},
						Field184: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field185: []*int32{},
						Field186: nil,
						Field187: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field188: []*HugeStruct0{GetHugeStruct0()},
						Field189: nil,
						Field190: []*int64{},
						Field191: map[string]*int32{
							"": nil,
						},
						Field192: []*HugeStruct0{GetHugeStruct0()},
						Field193: []*HugeStruct0{GetHugeStruct0()},
						Field194: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field195: []*bool{},
						Field196: map[string]*bool{
							"": nil,
						},
						Field197: []*bool{},
						Field198: nil,
						Field199: map[string]*int32{
							"": nil,
						},
						Field200: map[string]*int64{
							"": nil,
						},
						Field201: map[string]*string{
							"": nil,
						},
						Field202: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field203: map[string]*int32{
							"": nil,
						},
						Field204: nil,
						Field205: map[string]*string{
							"": nil,
						},
						Field206: []*HugeStruct0{GetHugeStruct0()},
						Field207: []*HugeStruct0{GetHugeStruct0()},
						Field208: nil,
						Field209: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field210: map[string]*string{
							"": nil,
						},
						Field211: map[string]*bool{
							"": nil,
						},
						Field212: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field213: nil,
						Field214: map[string]*bool{
							"": nil,
						},
						Field215: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field216: []*HugeStruct0{GetHugeStruct0()},
						Field217: map[string]*string{
							"": nil,
						},
						Field218: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field219: map[string]*int64{
							"": nil,
						},
						Field220: nil,
						Field221: nil,
						Field222: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field223: []*int64{},
						Field224: []*bool{},
						Field225: []*bool{},
						Field226: map[string]*int64{
							"": nil,
						},
						Field227: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field228: []*int64{},
						Field229: map[string]*bool{
							"": nil,
						},
						Field230: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field231: nil,
						Field232: nil,
						Field233: []*string{},
						Field234: []*HugeStruct0{GetHugeStruct0()},
						Field235: []*string{},
						Field236: nil,
						Field237: nil,
						Field238: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field239: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field240: []*HugeStruct0{GetHugeStruct0()},
						Field241: nil,
						Field242: nil,
						Field243: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field244: map[string]*bool{
							"": nil,
						},
						Field245: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field246: []*int32{},
						Field247: []*bool{},
						Field248: []*string{},
						Field249: nil,
						Field250: []*int32{},
						Field251: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field252: nil,
						Field253: map[string]*string{
							"": nil,
						},
						Field254: map[string]*string{
							"": nil,
						},
						Field255: []*int32{},
						Field256: nil,
						Field257: nil,
						Field258: map[string]*string{
							"": nil,
						},
						Field259: map[string]*int32{
							"": nil,
						},
						Field260: []*int64{},
						Field261: []*int32{},
						Field262: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field263: nil,
						Field264: nil,
						Field265: map[string]*bool{
							"": nil,
						},
						Field266: nil,
						Field267: []*int64{},
						Field268: nil,
						Field269: nil,
						Field270: map[string]*int64{
							"": nil,
						},
						Field271: map[string]*int64{
							"": nil,
						},
						Field272: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field273: []*string{},
						Field274: nil,
						Field275: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field276: map[string]*bool{
							"": nil,
						},
						Field277: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field278: nil,
						Field279: map[string]*string{
							"": nil,
						},
						Field280: nil,
						Field281: nil,
						Field282: nil,
						Field283: nil,
						Field284: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field285: map[string]*int64{
							"": nil,
						},
						Field286: map[string]*bool{
							"": nil,
						},
						Field287: map[string]*string{
							"": nil,
						},
						Field288: nil,
						Field289: nil,
						Field290: nil,
						Field291: []*int64{},
						Field292: map[string]*string{
							"": nil,
						},
						Field293: nil,
						Field294: []*string{},
						Field295: nil,
						Field296: []*HugeStruct0{GetHugeStruct0()},
						Field297: nil,
						Field298: map[string]*int64{
							"": nil,
						},
						Field299: map[string]*bool{
							"": nil,
						},
						Field300: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field301: nil,
						Field302: []*string{},
						Field303: []*string{},
						Field304: map[string]*string{
							"": nil,
						},
						Field305: nil,
						Field306: nil,
						Field307: []*HugeStruct0{GetHugeStruct0()},
						Field308: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field309: map[string]*int32{
							"": nil,
						},
						Field310: []*HugeStruct0{GetHugeStruct0()},
						Field311: nil,
						Field312: []*bool{},
						Field313: nil,
						Field314: []*HugeStruct0{GetHugeStruct0()},
						Field315: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field316: nil,
						Field317: nil,
						Field318: nil,
						Field319: []*int32{},
						Field320: nil,
						Field321: []*HugeStruct0{GetHugeStruct0()},
						Field322: nil,
						Field323: nil,
						Field324: []*HugeStruct0{GetHugeStruct0()},
						Field325: nil,
						Field326: []*int64{},
						Field327: nil,
						Field328: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field329: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field330: []*HugeStruct0{GetHugeStruct0()},
						Field331: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field332: []*string{},
						Field333: nil,
						Field334: []*HugeStruct0{GetHugeStruct0()},
						Field335: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field336: map[string]*bool{
							"": nil,
						},
						Field337: []*int64{},
						Field338: map[string]*bool{
							"": nil,
						},
						Field339: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field340: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field341: []*bool{},
						Field342: []*int64{},
						Field343: []*int32{},
						Field344: map[string]*bool{
							"": nil,
						},
						Field345: map[string]*int64{
							"": nil,
						},
						Field346: nil,
						Field347: map[string]*bool{
							"": nil,
						},
						Field348: map[string]*int32{
							"": nil,
						},
						Field349: []*string{},
						Field350: map[string]*int32{
							"": nil,
						},
						Field351: nil,
						Field352: []*int64{},
						Field353: []*int64{},
						Field354: nil,
						Field355: map[string]*int32{
							"": nil,
						},
						Field356: map[string]*bool{
							"": nil,
						},
						Field357: []*int32{},
						Field358: nil,
						Field359: map[string]*int64{
							"": nil,
						},
						Field360: nil,
						Field361: map[string]*int64{
							"": nil,
						},
						Field362: map[string]*int32{
							"": nil,
						},
						Field363: []*int64{},
						Field364: []*bool{},
						Field365: nil,
						Field366: map[string]*string{
							"": nil,
						},
						Field367: map[string]*bool{
							"": nil,
						},
						Field368: nil,
						Field369: nil,
						Field370: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field371: &HugeStruct0{
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
						Field372: map[string]*HugeStruct0{
							"": {
								Field0: map[string]*int64{
									"": nil,
								},
								Field1: nil,
								Field2: []*int64{},
								Field3: map[string]*int64{
									"": nil,
								},
								Field4: []*int64{},
							},
						},
						Field373: map[string]*bool{
							"": nil,
						},
					},
				},
				Field108: []*int32{},
				Field109: []*int64{},
				Field110: nil,
				Field111: map[string]*bool{
					"": nil,
				},
				Field112: []*int64{},
				Field113: nil,
				Field114: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field115: map[string]*int32{
					"": nil,
				},
				Field116: []*string{},
				Field117: []*int64{},
				Field118: []*int32{},
				Field119: nil,
				Field120: map[string]*string{
					"": nil,
				},
				Field121: map[string]*string{
					"": nil,
				},
				Field122: []*string{},
				Field123: map[string]*bool{
					"": nil,
				},
				Field124: map[string]*string{
					"": nil,
				},
				Field125: map[string]*int32{
					"": nil,
				},
				Field126: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field127: nil,
				Field128: []*int64{},
				Field129: &HugeStruct1{
					Field0: []*int32{},
					Field1: []*string{},
					Field2: []*int64{},
					Field3: map[string]*int32{
						"": nil,
					},
					Field4: []*bool{},
					Field5: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field6: map[string]*int32{
						"": nil,
					},
					Field7: map[string]*bool{
						"": nil,
					},
					Field8: []*bool{},
					Field9: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field10: []*string{},
					Field11: []*bool{},
					Field12: []*bool{},
					Field13: map[string]*int32{
						"": nil,
					},
					Field14: map[string]*int32{
						"": nil,
					},
					Field15: nil,
					Field16: []*int64{},
					Field17: []*bool{},
					Field18: map[string]*int64{
						"": nil,
					},
					Field19: []*int64{},
					Field20: map[string]*string{
						"": nil,
					},
					Field21: nil,
					Field22: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field23: []*string{},
					Field24: []*int64{},
					Field25: []*string{},
					Field26: []*bool{},
					Field27: map[string]*int32{
						"": nil,
					},
					Field28: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field29: map[string]*int32{
						"": nil,
					},
					Field30: map[string]*bool{
						"": nil,
					},
					Field31: map[string]*int32{
						"": nil,
					},
					Field32: []*HugeStruct0{GetHugeStruct0()},
					Field33: nil,
					Field34: map[string]*bool{
						"": nil,
					},
					Field35: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field36: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field37: nil,
					Field38: []*HugeStruct0{GetHugeStruct0()},
					Field39: []*bool{},
					Field40: map[string]*string{
						"": nil,
					},
					Field41: map[string]*int64{
						"": nil,
					},
					Field42: map[string]*int32{
						"": nil,
					},
					Field43: nil,
					Field44: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field45: map[string]*int32{
						"": nil,
					},
					Field46: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field47: nil,
					Field48: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field49: nil,
					Field50: map[string]*string{
						"": nil,
					},
					Field51: map[string]*bool{
						"": nil,
					},
					Field52: []*int64{},
					Field53: map[string]*string{
						"": nil,
					},
					Field54: []*int32{},
					Field55: map[string]*int64{
						"": nil,
					},
					Field56: map[string]*int32{
						"": nil,
					},
					Field57: map[string]*string{
						"": nil,
					},
					Field58: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field59: []*HugeStruct0{GetHugeStruct0()},
					Field60: map[string]*string{
						"": nil,
					},
					Field61: map[string]*bool{
						"": nil,
					},
					Field62: map[string]*int64{
						"": nil,
					},
					Field63: []*string{},
					Field64: []*int64{},
					Field65: map[string]*bool{
						"": nil,
					},
					Field66: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field67: []*int64{},
					Field68: map[string]*string{
						"": nil,
					},
					Field69: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field70: []*bool{},
					Field71: map[string]*int64{
						"": nil,
					},
					Field72: nil,
					Field73: map[string]*int32{
						"": nil,
					},
					Field74: nil,
					Field75: map[string]*int32{
						"": nil,
					},
					Field76: map[string]*string{
						"": nil,
					},
					Field77: []*string{},
					Field78: nil,
					Field79: map[string]*int64{
						"": nil,
					},
					Field80: []*int64{},
					Field81: map[string]*bool{
						"": nil,
					},
					Field82: []*string{},
					Field83: []*string{},
					Field84: nil,
					Field85: []*bool{},
					Field86: []*HugeStruct0{GetHugeStruct0()},
					Field87: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field88: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field89: []*int64{},
					Field90: []*int32{},
					Field91: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field92: []*bool{},
					Field93: []*string{},
					Field94: map[string]*int32{
						"": nil,
					},
					Field95: nil,
					Field96: nil,
					Field97: map[string]*bool{
						"": nil,
					},
					Field98: map[string]*int32{
						"": nil,
					},
					Field99:  []*HugeStruct0{GetHugeStruct0()},
					Field100: nil,
					Field101: nil,
					Field102: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field103: []*string{},
					Field104: []*string{},
					Field105: map[string]*bool{
						"": nil,
					},
					Field106: []*string{},
					Field107: []*int64{},
					Field108: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field109: nil,
					Field110: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field111: []*string{},
					Field112: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field113: []*bool{},
					Field114: []*bool{},
					Field115: map[string]*string{
						"": nil,
					},
					Field116: []*int64{},
					Field117: []*string{},
					Field118: map[string]*bool{
						"": nil,
					},
					Field119: map[string]*string{
						"": nil,
					},
					Field120: []*HugeStruct0{GetHugeStruct0()},
					Field121: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field122: []*bool{},
					Field123: nil,
					Field124: []*int64{},
					Field125: nil,
					Field126: []*string{},
					Field127: []*string{},
					Field128: []*int32{},
					Field129: []*bool{},
					Field130: nil,
					Field131: nil,
					Field132: []*int32{},
					Field133: []*int32{},
					Field134: nil,
					Field135: []*bool{},
					Field136: nil,
					Field137: []*int32{},
					Field138: map[string]*int64{
						"": nil,
					},
					Field139: map[string]*string{
						"": nil,
					},
					Field140: map[string]*int64{
						"": nil,
					},
					Field141: map[string]*int64{
						"": nil,
					},
					Field142: []*int32{},
					Field143: []*HugeStruct0{GetHugeStruct0()},
					Field144: map[string]*int64{
						"": nil,
					},
					Field145: []*string{},
					Field146: map[string]*int64{
						"": nil,
					},
					Field147: nil,
					Field148: map[string]*string{
						"": nil,
					},
					Field149: nil,
					Field150: map[string]*int64{
						"": nil,
					},
					Field151: map[string]*int64{
						"": nil,
					},
					Field152: map[string]*int32{
						"": nil,
					},
					Field153: []*int32{},
					Field154: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field155: map[string]*string{
						"": nil,
					},
					Field156: map[string]*int64{
						"": nil,
					},
					Field157: []*int32{},
					Field158: []*int32{},
					Field159: nil,
					Field160: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field161: []*bool{},
					Field162: []*HugeStruct0{GetHugeStruct0()},
					Field163: []*int32{},
					Field164: map[string]*string{
						"": nil,
					},
					Field165: []*bool{},
					Field166: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field167: nil,
					Field168: []*bool{},
					Field169: map[string]*bool{
						"": nil,
					},
					Field170: map[string]*bool{
						"": nil,
					},
					Field171: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field172: map[string]*bool{
						"": nil,
					},
					Field173: []*bool{},
					Field174: map[string]*int64{
						"": nil,
					},
					Field175: []*HugeStruct0{GetHugeStruct0()},
					Field176: []*int32{},
					Field177: []*int64{},
					Field178: map[string]*int64{
						"": nil,
					},
					Field179: []*int32{},
					Field180: []*string{},
					Field181: []*int32{},
					Field182: map[string]*string{
						"": nil,
					},
					Field183: []*int64{},
					Field184: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field185: []*int32{},
					Field186: nil,
					Field187: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field188: []*HugeStruct0{GetHugeStruct0()},
					Field189: nil,
					Field190: []*int64{},
					Field191: map[string]*int32{
						"": nil,
					},
					Field192: []*HugeStruct0{GetHugeStruct0()},
					Field193: []*HugeStruct0{GetHugeStruct0()},
					Field194: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field195: []*bool{},
					Field196: map[string]*bool{
						"": nil,
					},
					Field197: []*bool{},
					Field198: nil,
					Field199: map[string]*int32{
						"": nil,
					},
					Field200: map[string]*int64{
						"": nil,
					},
					Field201: map[string]*string{
						"": nil,
					},
					Field202: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field203: map[string]*int32{
						"": nil,
					},
					Field204: nil,
					Field205: map[string]*string{
						"": nil,
					},
					Field206: []*HugeStruct0{GetHugeStruct0()},
					Field207: []*HugeStruct0{GetHugeStruct0()},
					Field208: nil,
					Field209: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field210: map[string]*string{
						"": nil,
					},
					Field211: map[string]*bool{
						"": nil,
					},
					Field212: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field213: nil,
					Field214: map[string]*bool{
						"": nil,
					},
					Field215: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field216: []*HugeStruct0{GetHugeStruct0()},
					Field217: map[string]*string{
						"": nil,
					},
					Field218: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field219: map[string]*int64{
						"": nil,
					},
					Field220: nil,
					Field221: nil,
					Field222: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field223: []*int64{},
					Field224: []*bool{},
					Field225: []*bool{},
					Field226: map[string]*int64{
						"": nil,
					},
					Field227: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field228: []*int64{},
					Field229: map[string]*bool{
						"": nil,
					},
					Field230: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field231: nil,
					Field232: nil,
					Field233: []*string{},
					Field234: []*HugeStruct0{GetHugeStruct0()},
					Field235: []*string{},
					Field236: nil,
					Field237: nil,
					Field238: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field239: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field240: []*HugeStruct0{GetHugeStruct0()},
					Field241: nil,
					Field242: nil,
					Field243: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field244: map[string]*bool{
						"": nil,
					},
					Field245: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field246: []*int32{},
					Field247: []*bool{},
					Field248: []*string{},
					Field249: nil,
					Field250: []*int32{},
					Field251: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field252: nil,
					Field253: map[string]*string{
						"": nil,
					},
					Field254: map[string]*string{
						"": nil,
					},
					Field255: []*int32{},
					Field256: nil,
					Field257: nil,
					Field258: map[string]*string{
						"": nil,
					},
					Field259: map[string]*int32{
						"": nil,
					},
					Field260: []*int64{},
					Field261: []*int32{},
					Field262: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field263: nil,
					Field264: nil,
					Field265: map[string]*bool{
						"": nil,
					},
					Field266: nil,
					Field267: []*int64{},
					Field268: nil,
					Field269: nil,
					Field270: map[string]*int64{
						"": nil,
					},
					Field271: map[string]*int64{
						"": nil,
					},
					Field272: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field273: []*string{},
					Field274: nil,
					Field275: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field276: map[string]*bool{
						"": nil,
					},
					Field277: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field278: nil,
					Field279: map[string]*string{
						"": nil,
					},
					Field280: nil,
					Field281: nil,
					Field282: nil,
					Field283: nil,
					Field284: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field285: map[string]*int64{
						"": nil,
					},
					Field286: map[string]*bool{
						"": nil,
					},
					Field287: map[string]*string{
						"": nil,
					},
					Field288: nil,
					Field289: nil,
					Field290: nil,
					Field291: []*int64{},
					Field292: map[string]*string{
						"": nil,
					},
					Field293: nil,
					Field294: []*string{},
					Field295: nil,
					Field296: []*HugeStruct0{GetHugeStruct0()},
					Field297: nil,
					Field298: map[string]*int64{
						"": nil,
					},
					Field299: map[string]*bool{
						"": nil,
					},
					Field300: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field301: nil,
					Field302: []*string{},
					Field303: []*string{},
					Field304: map[string]*string{
						"": nil,
					},
					Field305: nil,
					Field306: nil,
					Field307: []*HugeStruct0{GetHugeStruct0()},
					Field308: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field309: map[string]*int32{
						"": nil,
					},
					Field310: []*HugeStruct0{GetHugeStruct0()},
					Field311: nil,
					Field312: []*bool{},
					Field313: nil,
					Field314: []*HugeStruct0{GetHugeStruct0()},
					Field315: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field316: nil,
					Field317: nil,
					Field318: nil,
					Field319: []*int32{},
					Field320: nil,
					Field321: []*HugeStruct0{GetHugeStruct0()},
					Field322: nil,
					Field323: nil,
					Field324: []*HugeStruct0{GetHugeStruct0()},
					Field325: nil,
					Field326: []*int64{},
					Field327: nil,
					Field328: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field329: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field330: []*HugeStruct0{GetHugeStruct0()},
					Field331: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field332: []*string{},
					Field333: nil,
					Field334: []*HugeStruct0{GetHugeStruct0()},
					Field335: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field336: map[string]*bool{
						"": nil,
					},
					Field337: []*int64{},
					Field338: map[string]*bool{
						"": nil,
					},
					Field339: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field340: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field341: []*bool{},
					Field342: []*int64{},
					Field343: []*int32{},
					Field344: map[string]*bool{
						"": nil,
					},
					Field345: map[string]*int64{
						"": nil,
					},
					Field346: nil,
					Field347: map[string]*bool{
						"": nil,
					},
					Field348: map[string]*int32{
						"": nil,
					},
					Field349: []*string{},
					Field350: map[string]*int32{
						"": nil,
					},
					Field351: nil,
					Field352: []*int64{},
					Field353: []*int64{},
					Field354: nil,
					Field355: map[string]*int32{
						"": nil,
					},
					Field356: map[string]*bool{
						"": nil,
					},
					Field357: []*int32{},
					Field358: nil,
					Field359: map[string]*int64{
						"": nil,
					},
					Field360: nil,
					Field361: map[string]*int64{
						"": nil,
					},
					Field362: map[string]*int32{
						"": nil,
					},
					Field363: []*int64{},
					Field364: []*bool{},
					Field365: nil,
					Field366: map[string]*string{
						"": nil,
					},
					Field367: map[string]*bool{
						"": nil,
					},
					Field368: nil,
					Field369: nil,
					Field370: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field371: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field372: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field373: map[string]*bool{
						"": nil,
					},
				},
				Field130: nil,
				Field131: &HugeStruct1{
					Field0: []*int32{},
					Field1: []*string{},
					Field2: []*int64{},
					Field3: map[string]*int32{
						"": nil,
					},
					Field4: []*bool{},
					Field5: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field6: map[string]*int32{
						"": nil,
					},
					Field7: map[string]*bool{
						"": nil,
					},
					Field8: []*bool{},
					Field9: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field10: []*string{},
					Field11: []*bool{},
					Field12: []*bool{},
					Field13: map[string]*int32{
						"": nil,
					},
					Field14: map[string]*int32{
						"": nil,
					},
					Field15: nil,
					Field16: []*int64{},
					Field17: []*bool{},
					Field18: map[string]*int64{
						"": nil,
					},
					Field19: []*int64{},
					Field20: map[string]*string{
						"": nil,
					},
					Field21: nil,
					Field22: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field23: []*string{},
					Field24: []*int64{},
					Field25: []*string{},
					Field26: []*bool{},
					Field27: map[string]*int32{
						"": nil,
					},
					Field28: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field29: map[string]*int32{
						"": nil,
					},
					Field30: map[string]*bool{
						"": nil,
					},
					Field31: map[string]*int32{
						"": nil,
					},
					Field32: []*HugeStruct0{GetHugeStruct0()},
					Field33: nil,
					Field34: map[string]*bool{
						"": nil,
					},
					Field35: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field36: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field37: nil,
					Field38: []*HugeStruct0{GetHugeStruct0()},
					Field39: []*bool{},
					Field40: map[string]*string{
						"": nil,
					},
					Field41: map[string]*int64{
						"": nil,
					},
					Field42: map[string]*int32{
						"": nil,
					},
					Field43: nil,
					Field44: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field45: map[string]*int32{
						"": nil,
					},
					Field46: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field47: nil,
					Field48: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field49: nil,
					Field50: map[string]*string{
						"": nil,
					},
					Field51: map[string]*bool{
						"": nil,
					},
					Field52: []*int64{},
					Field53: map[string]*string{
						"": nil,
					},
					Field54: []*int32{},
					Field55: map[string]*int64{
						"": nil,
					},
					Field56: map[string]*int32{
						"": nil,
					},
					Field57: map[string]*string{
						"": nil,
					},
					Field58: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field59: []*HugeStruct0{GetHugeStruct0()},
					Field60: map[string]*string{
						"": nil,
					},
					Field61: map[string]*bool{
						"": nil,
					},
					Field62: map[string]*int64{
						"": nil,
					},
					Field63: []*string{},
					Field64: []*int64{},
					Field65: map[string]*bool{
						"": nil,
					},
					Field66: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field67: []*int64{},
					Field68: map[string]*string{
						"": nil,
					},
					Field69: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field70: []*bool{},
					Field71: map[string]*int64{
						"": nil,
					},
					Field72: nil,
					Field73: map[string]*int32{
						"": nil,
					},
					Field74: nil,
					Field75: map[string]*int32{
						"": nil,
					},
					Field76: map[string]*string{
						"": nil,
					},
					Field77: []*string{},
					Field78: nil,
					Field79: map[string]*int64{
						"": nil,
					},
					Field80: []*int64{},
					Field81: map[string]*bool{
						"": nil,
					},
					Field82: []*string{},
					Field83: []*string{},
					Field84: nil,
					Field85: []*bool{},
					Field86: []*HugeStruct0{GetHugeStruct0()},
					Field87: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field88: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field89: []*int64{},
					Field90: []*int32{},
					Field91: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field92: []*bool{},
					Field93: []*string{},
					Field94: map[string]*int32{
						"": nil,
					},
					Field95: nil,
					Field96: nil,
					Field97: map[string]*bool{
						"": nil,
					},
					Field98: map[string]*int32{
						"": nil,
					},
					Field99:  []*HugeStruct0{GetHugeStruct0()},
					Field100: nil,
					Field101: nil,
					Field102: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field103: []*string{},
					Field104: []*string{},
					Field105: map[string]*bool{
						"": nil,
					},
					Field106: []*string{},
					Field107: []*int64{},
					Field108: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field109: nil,
					Field110: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field111: []*string{},
					Field112: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field113: []*bool{},
					Field114: []*bool{},
					Field115: map[string]*string{
						"": nil,
					},
					Field116: []*int64{},
					Field117: []*string{},
					Field118: map[string]*bool{
						"": nil,
					},
					Field119: map[string]*string{
						"": nil,
					},
					Field120: []*HugeStruct0{GetHugeStruct0()},
					Field121: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field122: []*bool{},
					Field123: nil,
					Field124: []*int64{},
					Field125: nil,
					Field126: []*string{},
					Field127: []*string{},
					Field128: []*int32{},
					Field129: []*bool{},
					Field130: nil,
					Field131: nil,
					Field132: []*int32{},
					Field133: []*int32{},
					Field134: nil,
					Field135: []*bool{},
					Field136: nil,
					Field137: []*int32{},
					Field138: map[string]*int64{
						"": nil,
					},
					Field139: map[string]*string{
						"": nil,
					},
					Field140: map[string]*int64{
						"": nil,
					},
					Field141: map[string]*int64{
						"": nil,
					},
					Field142: []*int32{},
					Field143: []*HugeStruct0{GetHugeStruct0()},
					Field144: map[string]*int64{
						"": nil,
					},
					Field145: []*string{},
					Field146: map[string]*int64{
						"": nil,
					},
					Field147: nil,
					Field148: map[string]*string{
						"": nil,
					},
					Field149: nil,
					Field150: map[string]*int64{
						"": nil,
					},
					Field151: map[string]*int64{
						"": nil,
					},
					Field152: map[string]*int32{
						"": nil,
					},
					Field153: []*int32{},
					Field154: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field155: map[string]*string{
						"": nil,
					},
					Field156: map[string]*int64{
						"": nil,
					},
					Field157: []*int32{},
					Field158: []*int32{},
					Field159: nil,
					Field160: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field161: []*bool{},
					Field162: []*HugeStruct0{GetHugeStruct0()},
					Field163: []*int32{},
					Field164: map[string]*string{
						"": nil,
					},
					Field165: []*bool{},
					Field166: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field167: nil,
					Field168: []*bool{},
					Field169: map[string]*bool{
						"": nil,
					},
					Field170: map[string]*bool{
						"": nil,
					},
					Field171: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field172: map[string]*bool{
						"": nil,
					},
					Field173: []*bool{},
					Field174: map[string]*int64{
						"": nil,
					},
					Field175: []*HugeStruct0{GetHugeStruct0()},
					Field176: []*int32{},
					Field177: []*int64{},
					Field178: map[string]*int64{
						"": nil,
					},
					Field179: []*int32{},
					Field180: []*string{},
					Field181: []*int32{},
					Field182: map[string]*string{
						"": nil,
					},
					Field183: []*int64{},
					Field184: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field185: []*int32{},
					Field186: nil,
					Field187: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field188: []*HugeStruct0{GetHugeStruct0()},
					Field189: nil,
					Field190: []*int64{},
					Field191: map[string]*int32{
						"": nil,
					},
					Field192: []*HugeStruct0{GetHugeStruct0()},
					Field193: []*HugeStruct0{GetHugeStruct0()},
					Field194: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field195: []*bool{},
					Field196: map[string]*bool{
						"": nil,
					},
					Field197: []*bool{},
					Field198: nil,
					Field199: map[string]*int32{
						"": nil,
					},
					Field200: map[string]*int64{
						"": nil,
					},
					Field201: map[string]*string{
						"": nil,
					},
					Field202: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field203: map[string]*int32{
						"": nil,
					},
					Field204: nil,
					Field205: map[string]*string{
						"": nil,
					},
					Field206: []*HugeStruct0{GetHugeStruct0()},
					Field207: []*HugeStruct0{GetHugeStruct0()},
					Field208: nil,
					Field209: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field210: map[string]*string{
						"": nil,
					},
					Field211: map[string]*bool{
						"": nil,
					},
					Field212: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field213: nil,
					Field214: map[string]*bool{
						"": nil,
					},
					Field215: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field216: []*HugeStruct0{GetHugeStruct0()},
					Field217: map[string]*string{
						"": nil,
					},
					Field218: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field219: map[string]*int64{
						"": nil,
					},
					Field220: nil,
					Field221: nil,
					Field222: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field223: []*int64{},
					Field224: []*bool{},
					Field225: []*bool{},
					Field226: map[string]*int64{
						"": nil,
					},
					Field227: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field228: []*int64{},
					Field229: map[string]*bool{
						"": nil,
					},
					Field230: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field231: nil,
					Field232: nil,
					Field233: []*string{},
					Field234: []*HugeStruct0{GetHugeStruct0()},
					Field235: []*string{},
					Field236: nil,
					Field237: nil,
					Field238: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field239: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field240: []*HugeStruct0{GetHugeStruct0()},
					Field241: nil,
					Field242: nil,
					Field243: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field244: map[string]*bool{
						"": nil,
					},
					Field245: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field246: []*int32{},
					Field247: []*bool{},
					Field248: []*string{},
					Field249: nil,
					Field250: []*int32{},
					Field251: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field252: nil,
					Field253: map[string]*string{
						"": nil,
					},
					Field254: map[string]*string{
						"": nil,
					},
					Field255: []*int32{},
					Field256: nil,
					Field257: nil,
					Field258: map[string]*string{
						"": nil,
					},
					Field259: map[string]*int32{
						"": nil,
					},
					Field260: []*int64{},
					Field261: []*int32{},
					Field262: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field263: nil,
					Field264: nil,
					Field265: map[string]*bool{
						"": nil,
					},
					Field266: nil,
					Field267: []*int64{},
					Field268: nil,
					Field269: nil,
					Field270: map[string]*int64{
						"": nil,
					},
					Field271: map[string]*int64{
						"": nil,
					},
					Field272: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field273: []*string{},
					Field274: nil,
					Field275: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field276: map[string]*bool{
						"": nil,
					},
					Field277: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field278: nil,
					Field279: map[string]*string{
						"": nil,
					},
					Field280: nil,
					Field281: nil,
					Field282: nil,
					Field283: nil,
					Field284: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field285: map[string]*int64{
						"": nil,
					},
					Field286: map[string]*bool{
						"": nil,
					},
					Field287: map[string]*string{
						"": nil,
					},
					Field288: nil,
					Field289: nil,
					Field290: nil,
					Field291: []*int64{},
					Field292: map[string]*string{
						"": nil,
					},
					Field293: nil,
					Field294: []*string{},
					Field295: nil,
					Field296: []*HugeStruct0{GetHugeStruct0()},
					Field297: nil,
					Field298: map[string]*int64{
						"": nil,
					},
					Field299: map[string]*bool{
						"": nil,
					},
					Field300: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field301: nil,
					Field302: []*string{},
					Field303: []*string{},
					Field304: map[string]*string{
						"": nil,
					},
					Field305: nil,
					Field306: nil,
					Field307: []*HugeStruct0{GetHugeStruct0()},
					Field308: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field309: map[string]*int32{
						"": nil,
					},
					Field310: []*HugeStruct0{GetHugeStruct0()},
					Field311: nil,
					Field312: []*bool{},
					Field313: nil,
					Field314: []*HugeStruct0{GetHugeStruct0()},
					Field315: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field316: nil,
					Field317: nil,
					Field318: nil,
					Field319: []*int32{},
					Field320: nil,
					Field321: []*HugeStruct0{GetHugeStruct0()},
					Field322: nil,
					Field323: nil,
					Field324: []*HugeStruct0{GetHugeStruct0()},
					Field325: nil,
					Field326: []*int64{},
					Field327: nil,
					Field328: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field329: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field330: []*HugeStruct0{GetHugeStruct0()},
					Field331: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field332: []*string{},
					Field333: nil,
					Field334: []*HugeStruct0{GetHugeStruct0()},
					Field335: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field336: map[string]*bool{
						"": nil,
					},
					Field337: []*int64{},
					Field338: map[string]*bool{
						"": nil,
					},
					Field339: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field340: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field341: []*bool{},
					Field342: []*int64{},
					Field343: []*int32{},
					Field344: map[string]*bool{
						"": nil,
					},
					Field345: map[string]*int64{
						"": nil,
					},
					Field346: nil,
					Field347: map[string]*bool{
						"": nil,
					},
					Field348: map[string]*int32{
						"": nil,
					},
					Field349: []*string{},
					Field350: map[string]*int32{
						"": nil,
					},
					Field351: nil,
					Field352: []*int64{},
					Field353: []*int64{},
					Field354: nil,
					Field355: map[string]*int32{
						"": nil,
					},
					Field356: map[string]*bool{
						"": nil,
					},
					Field357: []*int32{},
					Field358: nil,
					Field359: map[string]*int64{
						"": nil,
					},
					Field360: nil,
					Field361: map[string]*int64{
						"": nil,
					},
					Field362: map[string]*int32{
						"": nil,
					},
					Field363: []*int64{},
					Field364: []*bool{},
					Field365: nil,
					Field366: map[string]*string{
						"": nil,
					},
					Field367: map[string]*bool{
						"": nil,
					},
					Field368: nil,
					Field369: nil,
					Field370: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field371: &HugeStruct0{
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
					Field372: map[string]*HugeStruct0{
						"": {
							Field0: map[string]*int64{
								"": nil,
							},
							Field1: nil,
							Field2: []*int64{},
							Field3: map[string]*int64{
								"": nil,
							},
							Field4: []*int64{},
						},
					},
					Field373: map[string]*bool{
						"": nil,
					},
				},
				Field132: []*HugeStruct0{GetHugeStruct0()},
				Field133: map[string]*int64{
					"": nil,
				},
			},
		},
		Field49: []*bool{},
		Field50: []*int64{},
		Field51: map[string]*bool{
			"": nil,
		},
		Field52: []*string{},
		Field53: map[string]*int64{
			"": nil,
		},
		Field54: map[string]*string{
			"": nil,
		},
		Field55: map[string]*int64{
			"": nil,
		},
		Field56: nil,
		Field57: []*HugeStruct0{GetHugeStruct0()},
		Field58: []*bool{},
		Field59: nil,
		Field60: nil,
		Field61: map[string]*int32{
			"": nil,
		},
		Field62: nil,
		Field63: map[string]*int64{
			"": nil,
		},
		Field64: map[string]*HugeStruct1{
			"": {
				Field0: []*int32{},
				Field1: []*string{},
				Field2: []*int64{},
				Field3: map[string]*int32{
					"": nil,
				},
				Field4: []*bool{},
				Field5: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field6: map[string]*int32{
					"": nil,
				},
				Field7: map[string]*bool{
					"": nil,
				},
				Field8: []*bool{},
				Field9: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field10: []*string{},
				Field11: []*bool{},
				Field12: []*bool{},
				Field13: map[string]*int32{
					"": nil,
				},
				Field14: map[string]*int32{
					"": nil,
				},
				Field15: nil,
				Field16: []*int64{},
				Field17: []*bool{},
				Field18: map[string]*int64{
					"": nil,
				},
				Field19: []*int64{},
				Field20: map[string]*string{
					"": nil,
				},
				Field21: nil,
				Field22: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field23: []*string{},
				Field24: []*int64{},
				Field25: []*string{},
				Field26: []*bool{},
				Field27: map[string]*int32{
					"": nil,
				},
				Field28: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field29: map[string]*int32{
					"": nil,
				},
				Field30: map[string]*bool{
					"": nil,
				},
				Field31: map[string]*int32{
					"": nil,
				},
				Field32: []*HugeStruct0{GetHugeStruct0()},
				Field33: nil,
				Field34: map[string]*bool{
					"": nil,
				},
				Field35: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field36: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field37: nil,
				Field38: []*HugeStruct0{GetHugeStruct0()},
				Field39: []*bool{},
				Field40: map[string]*string{
					"": nil,
				},
				Field41: map[string]*int64{
					"": nil,
				},
				Field42: map[string]*int32{
					"": nil,
				},
				Field43: nil,
				Field44: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field45: map[string]*int32{
					"": nil,
				},
				Field46: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field47: nil,
				Field48: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field49: nil,
				Field50: map[string]*string{
					"": nil,
				},
				Field51: map[string]*bool{
					"": nil,
				},
				Field52: []*int64{},
				Field53: map[string]*string{
					"": nil,
				},
				Field54: []*int32{},
				Field55: map[string]*int64{
					"": nil,
				},
				Field56: map[string]*int32{
					"": nil,
				},
				Field57: map[string]*string{
					"": nil,
				},
				Field58: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field59: []*HugeStruct0{GetHugeStruct0()},
				Field60: map[string]*string{
					"": nil,
				},
				Field61: map[string]*bool{
					"": nil,
				},
				Field62: map[string]*int64{
					"": nil,
				},
				Field63: []*string{},
				Field64: []*int64{},
				Field65: map[string]*bool{
					"": nil,
				},
				Field66: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field67: []*int64{},
				Field68: map[string]*string{
					"": nil,
				},
				Field69: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field70: []*bool{},
				Field71: map[string]*int64{
					"": nil,
				},
				Field72: nil,
				Field73: map[string]*int32{
					"": nil,
				},
				Field74: nil,
				Field75: map[string]*int32{
					"": nil,
				},
				Field76: map[string]*string{
					"": nil,
				},
				Field77: []*string{},
				Field78: nil,
				Field79: map[string]*int64{
					"": nil,
				},
				Field80: []*int64{},
				Field81: map[string]*bool{
					"": nil,
				},
				Field82: []*string{},
				Field83: []*string{},
				Field84: nil,
				Field85: []*bool{},
				Field86: []*HugeStruct0{GetHugeStruct0()},
				Field87: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field88: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field89: []*int64{},
				Field90: []*int32{},
				Field91: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field92: []*bool{},
				Field93: []*string{},
				Field94: map[string]*int32{
					"": nil,
				},
				Field95: nil,
				Field96: nil,
				Field97: map[string]*bool{
					"": nil,
				},
				Field98: map[string]*int32{
					"": nil,
				},
				Field99:  []*HugeStruct0{GetHugeStruct0()},
				Field100: nil,
				Field101: nil,
				Field102: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field103: []*string{},
				Field104: []*string{},
				Field105: map[string]*bool{
					"": nil,
				},
				Field106: []*string{},
				Field107: []*int64{},
				Field108: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field109: nil,
				Field110: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field111: []*string{},
				Field112: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field113: []*bool{},
				Field114: []*bool{},
				Field115: map[string]*string{
					"": nil,
				},
				Field116: []*int64{},
				Field117: []*string{},
				Field118: map[string]*bool{
					"": nil,
				},
				Field119: map[string]*string{
					"": nil,
				},
				Field120: []*HugeStruct0{GetHugeStruct0()},
				Field121: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field122: []*bool{},
				Field123: nil,
				Field124: []*int64{},
				Field125: nil,
				Field126: []*string{},
				Field127: []*string{},
				Field128: []*int32{},
				Field129: []*bool{},
				Field130: nil,
				Field131: nil,
				Field132: []*int32{},
				Field133: []*int32{},
				Field134: nil,
				Field135: []*bool{},
				Field136: nil,
				Field137: []*int32{},
				Field138: map[string]*int64{
					"": nil,
				},
				Field139: map[string]*string{
					"": nil,
				},
				Field140: map[string]*int64{
					"": nil,
				},
				Field141: map[string]*int64{
					"": nil,
				},
				Field142: []*int32{},
				Field143: []*HugeStruct0{GetHugeStruct0()},
				Field144: map[string]*int64{
					"": nil,
				},
				Field145: []*string{},
				Field146: map[string]*int64{
					"": nil,
				},
				Field147: nil,
				Field148: map[string]*string{
					"": nil,
				},
				Field149: nil,
				Field150: map[string]*int64{
					"": nil,
				},
				Field151: map[string]*int64{
					"": nil,
				},
				Field152: map[string]*int32{
					"": nil,
				},
				Field153: []*int32{},
				Field154: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field155: map[string]*string{
					"": nil,
				},
				Field156: map[string]*int64{
					"": nil,
				},
				Field157: []*int32{},
				Field158: []*int32{},
				Field159: nil,
				Field160: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field161: []*bool{},
				Field162: []*HugeStruct0{GetHugeStruct0()},
				Field163: []*int32{},
				Field164: map[string]*string{
					"": nil,
				},
				Field165: []*bool{},
				Field166: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field167: nil,
				Field168: []*bool{},
				Field169: map[string]*bool{
					"": nil,
				},
				Field170: map[string]*bool{
					"": nil,
				},
				Field171: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field172: map[string]*bool{
					"": nil,
				},
				Field173: []*bool{},
				Field174: map[string]*int64{
					"": nil,
				},
				Field175: []*HugeStruct0{GetHugeStruct0()},
				Field176: []*int32{},
				Field177: []*int64{},
				Field178: map[string]*int64{
					"": nil,
				},
				Field179: []*int32{},
				Field180: []*string{},
				Field181: []*int32{},
				Field182: map[string]*string{
					"": nil,
				},
				Field183: []*int64{},
				Field184: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field185: []*int32{},
				Field186: nil,
				Field187: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field188: []*HugeStruct0{GetHugeStruct0()},
				Field189: nil,
				Field190: []*int64{},
				Field191: map[string]*int32{
					"": nil,
				},
				Field192: []*HugeStruct0{GetHugeStruct0()},
				Field193: []*HugeStruct0{GetHugeStruct0()},
				Field194: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field195: []*bool{},
				Field196: map[string]*bool{
					"": nil,
				},
				Field197: []*bool{},
				Field198: nil,
				Field199: map[string]*int32{
					"": nil,
				},
				Field200: map[string]*int64{
					"": nil,
				},
				Field201: map[string]*string{
					"": nil,
				},
				Field202: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field203: map[string]*int32{
					"": nil,
				},
				Field204: nil,
				Field205: map[string]*string{
					"": nil,
				},
				Field206: []*HugeStruct0{GetHugeStruct0()},
				Field207: []*HugeStruct0{GetHugeStruct0()},
				Field208: nil,
				Field209: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field210: map[string]*string{
					"": nil,
				},
				Field211: map[string]*bool{
					"": nil,
				},
				Field212: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field213: nil,
				Field214: map[string]*bool{
					"": nil,
				},
				Field215: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field216: []*HugeStruct0{GetHugeStruct0()},
				Field217: map[string]*string{
					"": nil,
				},
				Field218: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field219: map[string]*int64{
					"": nil,
				},
				Field220: nil,
				Field221: nil,
				Field222: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field223: []*int64{},
				Field224: []*bool{},
				Field225: []*bool{},
				Field226: map[string]*int64{
					"": nil,
				},
				Field227: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field228: []*int64{},
				Field229: map[string]*bool{
					"": nil,
				},
				Field230: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field231: nil,
				Field232: nil,
				Field233: []*string{},
				Field234: []*HugeStruct0{GetHugeStruct0()},
				Field235: []*string{},
				Field236: nil,
				Field237: nil,
				Field238: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field239: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field240: []*HugeStruct0{GetHugeStruct0()},
				Field241: nil,
				Field242: nil,
				Field243: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field244: map[string]*bool{
					"": nil,
				},
				Field245: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field246: []*int32{},
				Field247: []*bool{},
				Field248: []*string{},
				Field249: nil,
				Field250: []*int32{},
				Field251: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field252: nil,
				Field253: map[string]*string{
					"": nil,
				},
				Field254: map[string]*string{
					"": nil,
				},
				Field255: []*int32{},
				Field256: nil,
				Field257: nil,
				Field258: map[string]*string{
					"": nil,
				},
				Field259: map[string]*int32{
					"": nil,
				},
				Field260: []*int64{},
				Field261: []*int32{},
				Field262: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field263: nil,
				Field264: nil,
				Field265: map[string]*bool{
					"": nil,
				},
				Field266: nil,
				Field267: []*int64{},
				Field268: nil,
				Field269: nil,
				Field270: map[string]*int64{
					"": nil,
				},
				Field271: map[string]*int64{
					"": nil,
				},
				Field272: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field273: []*string{},
				Field274: nil,
				Field275: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field276: map[string]*bool{
					"": nil,
				},
				Field277: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field278: nil,
				Field279: map[string]*string{
					"": nil,
				},
				Field280: nil,
				Field281: nil,
				Field282: nil,
				Field283: nil,
				Field284: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field285: map[string]*int64{
					"": nil,
				},
				Field286: map[string]*bool{
					"": nil,
				},
				Field287: map[string]*string{
					"": nil,
				},
				Field288: nil,
				Field289: nil,
				Field290: nil,
				Field291: []*int64{},
				Field292: map[string]*string{
					"": nil,
				},
				Field293: nil,
				Field294: []*string{},
				Field295: nil,
				Field296: []*HugeStruct0{GetHugeStruct0()},
				Field297: nil,
				Field298: map[string]*int64{
					"": nil,
				},
				Field299: map[string]*bool{
					"": nil,
				},
				Field300: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field301: nil,
				Field302: []*string{},
				Field303: []*string{},
				Field304: map[string]*string{
					"": nil,
				},
				Field305: nil,
				Field306: nil,
				Field307: []*HugeStruct0{GetHugeStruct0()},
				Field308: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field309: map[string]*int32{
					"": nil,
				},
				Field310: []*HugeStruct0{GetHugeStruct0()},
				Field311: nil,
				Field312: []*bool{},
				Field313: nil,
				Field314: []*HugeStruct0{GetHugeStruct0()},
				Field315: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field316: nil,
				Field317: nil,
				Field318: nil,
				Field319: []*int32{},
				Field320: nil,
				Field321: []*HugeStruct0{GetHugeStruct0()},
				Field322: nil,
				Field323: nil,
				Field324: []*HugeStruct0{GetHugeStruct0()},
				Field325: nil,
				Field326: []*int64{},
				Field327: nil,
				Field328: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field329: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field330: []*HugeStruct0{GetHugeStruct0()},
				Field331: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field332: []*string{},
				Field333: nil,
				Field334: []*HugeStruct0{GetHugeStruct0()},
				Field335: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field336: map[string]*bool{
					"": nil,
				},
				Field337: []*int64{},
				Field338: map[string]*bool{
					"": nil,
				},
				Field339: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field340: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field341: []*bool{},
				Field342: []*int64{},
				Field343: []*int32{},
				Field344: map[string]*bool{
					"": nil,
				},
				Field345: map[string]*int64{
					"": nil,
				},
				Field346: nil,
				Field347: map[string]*bool{
					"": nil,
				},
				Field348: map[string]*int32{
					"": nil,
				},
				Field349: []*string{},
				Field350: map[string]*int32{
					"": nil,
				},
				Field351: nil,
				Field352: []*int64{},
				Field353: []*int64{},
				Field354: nil,
				Field355: map[string]*int32{
					"": nil,
				},
				Field356: map[string]*bool{
					"": nil,
				},
				Field357: []*int32{},
				Field358: nil,
				Field359: map[string]*int64{
					"": nil,
				},
				Field360: nil,
				Field361: map[string]*int64{
					"": nil,
				},
				Field362: map[string]*int32{
					"": nil,
				},
				Field363: []*int64{},
				Field364: []*bool{},
				Field365: nil,
				Field366: map[string]*string{
					"": nil,
				},
				Field367: map[string]*bool{
					"": nil,
				},
				Field368: nil,
				Field369: nil,
				Field370: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field371: &HugeStruct0{
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
				Field372: map[string]*HugeStruct0{
					"": {
						Field0: map[string]*int64{
							"": nil,
						},
						Field1: nil,
						Field2: []*int64{},
						Field3: map[string]*int64{
							"": nil,
						},
						Field4: []*int64{},
					},
				},
				Field373: map[string]*bool{
					"": nil,
				},
			},
		},
		Field65: []*string{},
		Field66: []*HugeStruct2{},
		Field67: map[string]*bool{
			"": nil,
		},
		Field68: []*bool{},
		Field69: map[string]*int64{
			"": nil,
		},
		Field70: []*int64{},
		Field71: map[string]*int32{
			"": nil,
		},
		Field72: []*int64{},
		Field73: []*int32{},
		Field74: []*bool{},
		Field75: []*int64{},
		Field76: map[string]*int64{
			"": nil,
		},
		Field77: nil,
		Field78: nil,
		Field79: []*string{},
		Field80: map[string]*bool{
			"": nil,
		},
		Field81: map[string]*int64{
			"": nil,
		},
		Field82: []*HugeStruct2{},
		Field83: map[string]*string{
			"": nil,
		},
		Field84: nil,
		Field85: nil,
		Field86: []*string{},
		Field87: []*int64{},
		Field88: []*int64{},
		Field89: []*HugeStruct1{},
		Field90: nil,
		Field91: map[string]*bool{
			"": nil,
		},
		Field92: GetHugeStruct0(),
		Field93: []*bool{},
		Field94: map[string]*string{
			"": nil,
		},
		Field95: map[string]*int64{
			"": nil,
		},
		Field96:  []*HugeStruct1{},
		Field97:  []*int32{},
		Field98:  []*int64{},
		Field99:  nil,
		Field100: []*string{},
		Field101: map[string]*int64{
			"": nil,
		},
		Field102: map[string]*string{
			"": nil,
		},
		Field103: []*int32{},
		Field104: map[string]*string{
			"": nil,
		},
		Field105: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field106: []*int32{},
		Field107: &HugeStruct1{
			Field0: []*int32{},
			Field1: []*string{},
			Field2: []*int64{},
			Field3: map[string]*int32{
				"": nil,
			},
			Field4: []*bool{},
			Field5: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field6: map[string]*int32{
				"": nil,
			},
			Field7: map[string]*bool{
				"": nil,
			},
			Field8: []*bool{},
			Field9: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field10: []*string{},
			Field11: []*bool{},
			Field12: []*bool{},
			Field13: map[string]*int32{
				"": nil,
			},
			Field14: map[string]*int32{
				"": nil,
			},
			Field15: nil,
			Field16: []*int64{},
			Field17: []*bool{},
			Field18: map[string]*int64{
				"": nil,
			},
			Field19: []*int64{},
			Field20: map[string]*string{
				"": nil,
			},
			Field21: nil,
			Field22: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field23: []*string{},
			Field24: []*int64{},
			Field25: []*string{},
			Field26: []*bool{},
			Field27: map[string]*int32{
				"": nil,
			},
			Field28: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field29: map[string]*int32{
				"": nil,
			},
			Field30: map[string]*bool{
				"": nil,
			},
			Field31: map[string]*int32{
				"": nil,
			},
			Field32: []*HugeStruct0{GetHugeStruct0()},
			Field33: nil,
			Field34: map[string]*bool{
				"": nil,
			},
			Field35: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field36: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field37: nil,
			Field38: []*HugeStruct0{GetHugeStruct0()},
			Field39: []*bool{},
			Field40: map[string]*string{
				"": nil,
			},
			Field41: map[string]*int64{
				"": nil,
			},
			Field42: map[string]*int32{
				"": nil,
			},
			Field43: nil,
			Field44: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field45: map[string]*int32{
				"": nil,
			},
			Field46: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field47: nil,
			Field48: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field49: nil,
			Field50: map[string]*string{
				"": nil,
			},
			Field51: map[string]*bool{
				"": nil,
			},
			Field52: []*int64{},
			Field53: map[string]*string{
				"": nil,
			},
			Field54: []*int32{},
			Field55: map[string]*int64{
				"": nil,
			},
			Field56: map[string]*int32{
				"": nil,
			},
			Field57: map[string]*string{
				"": nil,
			},
			Field58: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field59: []*HugeStruct0{GetHugeStruct0()},
			Field60: map[string]*string{
				"": nil,
			},
			Field61: map[string]*bool{
				"": nil,
			},
			Field62: map[string]*int64{
				"": nil,
			},
			Field63: []*string{},
			Field64: []*int64{},
			Field65: map[string]*bool{
				"": nil,
			},
			Field66: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field67: []*int64{},
			Field68: map[string]*string{
				"": nil,
			},
			Field69: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field70: []*bool{},
			Field71: map[string]*int64{
				"": nil,
			},
			Field72: nil,
			Field73: map[string]*int32{
				"": nil,
			},
			Field74: nil,
			Field75: map[string]*int32{
				"": nil,
			},
			Field76: map[string]*string{
				"": nil,
			},
			Field77: []*string{},
			Field78: nil,
			Field79: map[string]*int64{
				"": nil,
			},
			Field80: []*int64{},
			Field81: map[string]*bool{
				"": nil,
			},
			Field82: []*string{},
			Field83: []*string{},
			Field84: nil,
			Field85: []*bool{},
			Field86: []*HugeStruct0{GetHugeStruct0()},
			Field87: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field88: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field89: []*int64{},
			Field90: []*int32{},
			Field91: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field92: []*bool{},
			Field93: []*string{},
			Field94: map[string]*int32{
				"": nil,
			},
			Field95: nil,
			Field96: nil,
			Field97: map[string]*bool{
				"": nil,
			},
			Field98: map[string]*int32{
				"": nil,
			},
			Field99:  []*HugeStruct0{GetHugeStruct0()},
			Field100: nil,
			Field101: nil,
			Field102: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field103: []*string{},
			Field104: []*string{},
			Field105: map[string]*bool{
				"": nil,
			},
			Field106: []*string{},
			Field107: []*int64{},
			Field108: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field109: nil,
			Field110: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field111: []*string{},
			Field112: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field113: []*bool{},
			Field114: []*bool{},
			Field115: map[string]*string{
				"": nil,
			},
			Field116: []*int64{},
			Field117: []*string{},
			Field118: map[string]*bool{
				"": nil,
			},
			Field119: map[string]*string{
				"": nil,
			},
			Field120: []*HugeStruct0{GetHugeStruct0()},
			Field121: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field122: []*bool{},
			Field123: nil,
			Field124: []*int64{},
			Field125: nil,
			Field126: []*string{},
			Field127: []*string{},
			Field128: []*int32{},
			Field129: []*bool{},
			Field130: nil,
			Field131: nil,
			Field132: []*int32{},
			Field133: []*int32{},
			Field134: nil,
			Field135: []*bool{},
			Field136: nil,
			Field137: []*int32{},
			Field138: map[string]*int64{
				"": nil,
			},
			Field139: map[string]*string{
				"": nil,
			},
			Field140: map[string]*int64{
				"": nil,
			},
			Field141: map[string]*int64{
				"": nil,
			},
			Field142: []*int32{},
			Field143: []*HugeStruct0{GetHugeStruct0()},
			Field144: map[string]*int64{
				"": nil,
			},
			Field145: []*string{},
			Field146: map[string]*int64{
				"": nil,
			},
			Field147: nil,
			Field148: map[string]*string{
				"": nil,
			},
			Field149: nil,
			Field150: map[string]*int64{
				"": nil,
			},
			Field151: map[string]*int64{
				"": nil,
			},
			Field152: map[string]*int32{
				"": nil,
			},
			Field153: []*int32{},
			Field154: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field155: map[string]*string{
				"": nil,
			},
			Field156: map[string]*int64{
				"": nil,
			},
			Field157: []*int32{},
			Field158: []*int32{},
			Field159: nil,
			Field160: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field161: []*bool{},
			Field162: []*HugeStruct0{GetHugeStruct0()},
			Field163: []*int32{},
			Field164: map[string]*string{
				"": nil,
			},
			Field165: []*bool{},
			Field166: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field167: nil,
			Field168: []*bool{},
			Field169: map[string]*bool{
				"": nil,
			},
			Field170: map[string]*bool{
				"": nil,
			},
			Field171: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field172: map[string]*bool{
				"": nil,
			},
			Field173: []*bool{},
			Field174: map[string]*int64{
				"": nil,
			},
			Field175: []*HugeStruct0{GetHugeStruct0()},
			Field176: []*int32{},
			Field177: []*int64{},
			Field178: map[string]*int64{
				"": nil,
			},
			Field179: []*int32{},
			Field180: []*string{},
			Field181: []*int32{},
			Field182: map[string]*string{
				"": nil,
			},
			Field183: []*int64{},
			Field184: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field185: []*int32{},
			Field186: nil,
			Field187: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field188: []*HugeStruct0{GetHugeStruct0()},
			Field189: nil,
			Field190: []*int64{},
			Field191: map[string]*int32{
				"": nil,
			},
			Field192: []*HugeStruct0{GetHugeStruct0()},
			Field193: []*HugeStruct0{GetHugeStruct0()},
			Field194: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field195: []*bool{},
			Field196: map[string]*bool{
				"": nil,
			},
			Field197: []*bool{},
			Field198: nil,
			Field199: map[string]*int32{
				"": nil,
			},
			Field200: map[string]*int64{
				"": nil,
			},
			Field201: map[string]*string{
				"": nil,
			},
			Field202: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field203: map[string]*int32{
				"": nil,
			},
			Field204: nil,
			Field205: map[string]*string{
				"": nil,
			},
			Field206: []*HugeStruct0{GetHugeStruct0()},
			Field207: []*HugeStruct0{GetHugeStruct0()},
			Field208: nil,
			Field209: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field210: map[string]*string{
				"": nil,
			},
			Field211: map[string]*bool{
				"": nil,
			},
			Field212: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field213: nil,
			Field214: map[string]*bool{
				"": nil,
			},
			Field215: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field216: []*HugeStruct0{GetHugeStruct0()},
			Field217: map[string]*string{
				"": nil,
			},
			Field218: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field219: map[string]*int64{
				"": nil,
			},
			Field220: nil,
			Field221: nil,
			Field222: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field223: []*int64{},
			Field224: []*bool{},
			Field225: []*bool{},
			Field226: map[string]*int64{
				"": nil,
			},
			Field227: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field228: []*int64{},
			Field229: map[string]*bool{
				"": nil,
			},
			Field230: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field231: nil,
			Field232: nil,
			Field233: []*string{},
			Field234: []*HugeStruct0{GetHugeStruct0()},
			Field235: []*string{},
			Field236: nil,
			Field237: nil,
			Field238: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field239: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field240: []*HugeStruct0{GetHugeStruct0()},
			Field241: nil,
			Field242: nil,
			Field243: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field244: map[string]*bool{
				"": nil,
			},
			Field245: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field246: []*int32{},
			Field247: []*bool{},
			Field248: []*string{},
			Field249: nil,
			Field250: []*int32{},
			Field251: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field252: nil,
			Field253: map[string]*string{
				"": nil,
			},
			Field254: map[string]*string{
				"": nil,
			},
			Field255: []*int32{},
			Field256: nil,
			Field257: nil,
			Field258: map[string]*string{
				"": nil,
			},
			Field259: map[string]*int32{
				"": nil,
			},
			Field260: []*int64{},
			Field261: []*int32{},
			Field262: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field263: nil,
			Field264: nil,
			Field265: map[string]*bool{
				"": nil,
			},
			Field266: nil,
			Field267: []*int64{},
			Field268: nil,
			Field269: nil,
			Field270: map[string]*int64{
				"": nil,
			},
			Field271: map[string]*int64{
				"": nil,
			},
			Field272: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field273: []*string{},
			Field274: nil,
			Field275: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field276: map[string]*bool{
				"": nil,
			},
			Field277: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field278: nil,
			Field279: map[string]*string{
				"": nil,
			},
			Field280: nil,
			Field281: nil,
			Field282: nil,
			Field283: nil,
			Field284: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field285: map[string]*int64{
				"": nil,
			},
			Field286: map[string]*bool{
				"": nil,
			},
			Field287: map[string]*string{
				"": nil,
			},
			Field288: nil,
			Field289: nil,
			Field290: nil,
			Field291: []*int64{},
			Field292: map[string]*string{
				"": nil,
			},
			Field293: nil,
			Field294: []*string{},
			Field295: nil,
			Field296: []*HugeStruct0{GetHugeStruct0()},
			Field297: nil,
			Field298: map[string]*int64{
				"": nil,
			},
			Field299: map[string]*bool{
				"": nil,
			},
			Field300: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field301: nil,
			Field302: []*string{},
			Field303: []*string{},
			Field304: map[string]*string{
				"": nil,
			},
			Field305: nil,
			Field306: nil,
			Field307: []*HugeStruct0{GetHugeStruct0()},
			Field308: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field309: map[string]*int32{
				"": nil,
			},
			Field310: []*HugeStruct0{GetHugeStruct0()},
			Field311: nil,
			Field312: []*bool{},
			Field313: nil,
			Field314: []*HugeStruct0{GetHugeStruct0()},
			Field315: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field316: nil,
			Field317: nil,
			Field318: nil,
			Field319: []*int32{},
			Field320: nil,
			Field321: []*HugeStruct0{GetHugeStruct0()},
			Field322: nil,
			Field323: nil,
			Field324: []*HugeStruct0{GetHugeStruct0()},
			Field325: nil,
			Field326: []*int64{},
			Field327: nil,
			Field328: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field329: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field330: []*HugeStruct0{GetHugeStruct0()},
			Field331: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field332: []*string{},
			Field333: nil,
			Field334: []*HugeStruct0{GetHugeStruct0()},
			Field335: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field336: map[string]*bool{
				"": nil,
			},
			Field337: []*int64{},
			Field338: map[string]*bool{
				"": nil,
			},
			Field339: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field340: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field341: []*bool{},
			Field342: []*int64{},
			Field343: []*int32{},
			Field344: map[string]*bool{
				"": nil,
			},
			Field345: map[string]*int64{
				"": nil,
			},
			Field346: nil,
			Field347: map[string]*bool{
				"": nil,
			},
			Field348: map[string]*int32{
				"": nil,
			},
			Field349: []*string{},
			Field350: map[string]*int32{
				"": nil,
			},
			Field351: nil,
			Field352: []*int64{},
			Field353: []*int64{},
			Field354: nil,
			Field355: map[string]*int32{
				"": nil,
			},
			Field356: map[string]*bool{
				"": nil,
			},
			Field357: []*int32{},
			Field358: nil,
			Field359: map[string]*int64{
				"": nil,
			},
			Field360: nil,
			Field361: map[string]*int64{
				"": nil,
			},
			Field362: map[string]*int32{
				"": nil,
			},
			Field363: []*int64{},
			Field364: []*bool{},
			Field365: nil,
			Field366: map[string]*string{
				"": nil,
			},
			Field367: map[string]*bool{
				"": nil,
			},
			Field368: nil,
			Field369: nil,
			Field370: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field371: &HugeStruct0{
				Field0: map[string]*int64{
					"": nil,
				},
				Field1: nil,
				Field2: []*int64{},
				Field3: map[string]*int64{
					"": nil,
				},
				Field4: []*int64{},
			},
			Field372: map[string]*HugeStruct0{
				"": {
					Field0: map[string]*int64{
						"": nil,
					},
					Field1: nil,
					Field2: []*int64{},
					Field3: map[string]*int64{
						"": nil,
					},
					Field4: []*int64{},
				},
			},
			Field373: map[string]*bool{
				"": nil,
			},
		},
		Field108: []*HugeStruct1{},
		Field109: []*bool{},
		Field110: []*int32{},
		Field111: map[string]*string{
			"": nil,
		},
		Field112: map[string]*HugeStruct0{"a": GetHugeStruct0()},
		Field113: map[string]*int32{
			"": nil,
		},
		Field114: []*bool{},
		Field115: []*HugeStruct2{},
		Field116: map[string]*bool{
			"": nil,
		},
		Field117: map[string]*string{
			"": nil,
		},
		Field118: nil,
		Field119: nil,
	}
}

func GetHugeStruct4() *HugeStruct4 {
	return &HugeStruct4{
		Field3:  GetHugeStruc3(),
		Field25: GetHugeStruct2(),
		Field30: map[string]*HugeStruct3{
			"a": GetHugeStruc3(),
		},
		Field103: []*HugeStruct3{
			GetHugeStruc3(),
		},
	}
}

func GetHugeStruct5() *HugeStruct5 {
	return &HugeStruct5{
		Field40: map[string]*HugeStruct4{
			"a": GetHugeStruct4(),
		},
		Field47: []*HugeStruct4{
			GetHugeStruct4(),
		},
	}
}

func GetHugeStruct6() *HugeStruct6 {
	return &HugeStruct6{
		Field2: GetHugeStruct4(),
		Field4: GetHugeStruct5(),
		Field13: []*HugeStruct5{
			GetHugeStruct5(),
		},
	}
}
