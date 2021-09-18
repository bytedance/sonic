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

import (
    . `github.com/bytedance/sonic`
    `testing`
    `math`
    `encoding/json`

    `github.com/bytedance/sonic/decoder`
    `github.com/stretchr/testify/require`
)


func TestNegZeroInIEEE754(t *testing.T) {
    var sonicobj, stdobj float64
    sonicerr := Unmarshal([]byte("-0.0"), &sonicobj)
    stderr := json.Unmarshal([]byte("-0.0"), &stdobj)
    if sonicerr != nil && stderr == nil {
        println(sonicerr.(decoder.SyntaxError).Description())
        require.NoError(t, sonicerr)
    }
    require.Equal(t, math.Float64bits(sonicobj), math.Float64bits(stdobj))

    sonicout, sonicerr2 := Marshal(&stdobj)
    stdout, stderr2 := json.Marshal(&stdobj)
    if sonicerr2 != nil && stderr2 == nil {
        println(sonicerr2)
        require.NoError(t, sonicerr2)
    }
    require.Equal(t, sonicout, stdout)
}