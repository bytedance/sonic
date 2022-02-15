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
    `math/rand`
    `runtime`
    `sync`
    `testing`
    `time`

    `github.com/bytedance/sonic`
)


type GlobalConfig []Conf

type Conf struct {
    A  string             `json:"A"`
    B  SubConf            `json:"B"`
    C []string            `json:"C"`
}

type SubConf struct {
    Slice        []int64 `json:"Slice"`
    Map          map[int64]bool `json:"-"`
}

func IntSlide2Map(l []int64) map[int64]bool {
    m := make(map[int64]bool)
    for _, item := range l {
        m[item] = true
    }
    return m
}

func Reload(t *testing.T, rawData string) (tmp GlobalConfig) {
    buf := []byte(rawData)
    runtime.GC()
    // t.Logf("got bytes %x\n", unsafe.Pointer(&buf[0]))
    // runtime.SetFinalizer(&buf[0], func(x *byte){
        // t.Logf("&byte %x got free\n", x)
    // })
    err := sonic.Unmarshal(buf, &tmp) // better use sonic.UnmarshalString()!
    if err != nil {
        t.Fatalf("failed to unmarshal json, raw data: %v, err: %v", rawData, err)
    }
    runtime.GC()
    // t.Log("unmarshal done")
    for index, conf := range tmp {
        tmp[index].B.Map = IntSlide2Map(conf.B.Slice)
    }
    runtime.GC()
    // t.Log("calc done")
    return
}

func TestIssue186(t *testing.T) {
    t.Parallel()
    var data = `[{"A":"xxx","B":{"Slice":[111]}},{"A":"yyy","B":{"Slice":[222]},"C":["extra"]},{"A":"zzz","B":{"Slice":[333]},"C":["extra"]},{"A":"zzz","B":{"Slice":[333]},"C":["extra"]},{"A":"zzz","B":{"Slice":[1111111111,2222222222,3333333333,44444444444,55555555555]},"C":["extra","aaaaaaaaaaaa","bbbbbbbbbbbbb","ccccccccccccc","ddddddddddddd"]}]`
    // var obj interface{}
    for k:=0; k<100; k++ {
        wg := sync.WaitGroup{}
        for i:=0; i<1000; i++ {
            wg.Add(1)
            go func(){
                defer wg.Done()
                time.Sleep(time.Duration(rand.Intn(100)+1000))
                tmp := Reload(t, data)
                runtime.GC()
                _ = tmp[0].A
                runtime.GC()
                // obj = tmp
            }()
        }
        runtime.GC()
        // t.Log(obj)
        wg.Wait()
    }
}