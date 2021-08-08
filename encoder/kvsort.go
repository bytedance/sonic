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
 
package encoder

import (
    "bytes"
)

// Algorithm 3-way Radix Quicksort, d means the radix.
// Reference: https://algs4.cs.princeton.edu/51radix/Quick3string.java.html
func radixQsort(kvs kvSlice, d, maxDepth int) {
    if maxDepth == 0 {
        heapSort(kvs, 0, len(kvs))
        return
    }
    maxDepth--
    for len(kvs) > 11 {
        p := pivot(kvs, d)
        lt, i, gt := 0, 0, len(kvs)
        for i < gt {
            c := byteAt(kvs[i].k, d)
            if c < p {
                swap(kvs, lt, i)
                i++
                lt++
            } else if c > p {
                gt--
                swap(kvs, i, gt)
            } else {
                i++
            }
        }

        // kvs[0:lt] < v = kvs[lt:gt] < kvs[gt:len(kvs)]
        // Native:
        // radixQsort(kvs[:lt], d, maxDepth)
        // if p > -1 {
        //     radixQsort(kvs[lt:gt], d+1, maxDepth)
        // }
        // radixQsort(kvs[gt:], d, maxDepth)
        // Optimize: make recursive calls only for the smaller parts.
        // Reference: https://www.geeksforgeeks.org/quicksort-tail-call-optimization-reducing-worst-case-space-log-n/
        if p == -1 {
            if lt > len(kvs) - gt {
                radixQsort(kvs[gt:], d, maxDepth)
                kvs = kvs[:lt]
            } else {
                radixQsort(kvs[:lt], d, maxDepth)
                kvs = kvs[gt:]
            }
        } else {
            ml := maxThree(lt, gt-lt, len(kvs)-gt)
            if ml == lt {
                radixQsort(kvs[lt:gt], d+1, maxDepth)
                radixQsort(kvs[gt:], d, maxDepth)
                kvs = kvs[:lt]
            } else if ml == gt-lt {
                radixQsort(kvs[:lt], d, maxDepth)
                radixQsort(kvs[gt:], d, maxDepth)
                kvs = kvs[lt:gt]
                d += 1
            } else {
                radixQsort(kvs[:lt], d, maxDepth)
                radixQsort(kvs[lt:gt], d+1, maxDepth)
                kvs = kvs[gt:] 
            }
        }
    }
    insertRadixSort(kvs, d)
}

func insertRadixSort(kvs kvSlice, d int) {
    for i := 1; i < len(kvs); i++ {
        for j := i; j > 0 && lessFrom(kvs[j].k, kvs[j-1].k, d); j-- {
            kvs[j], kvs[j-1] = kvs[j-1], kvs[j]
        }
    }
}

func pivot(kvs kvSlice, d int) int {
    m := len(kvs) >> 1
    if len(kvs) > 40 {
        // Tukey's ``Ninther,'' median of three mediankvs of three.
        t := len(kvs) / 8
        return medianThree(
            medianThree(byteAt(kvs[0].k, d), byteAt(kvs[t].k, d), byteAt(kvs[2*t].k, d)),
            medianThree(byteAt(kvs[m].k, d), byteAt(kvs[m-t].k, d), byteAt(kvs[m+t].k, d)),
            medianThree(byteAt(kvs[len(kvs)-1].k, d),
                byteAt(kvs[len(kvs)-1-t].k, d),
                byteAt(kvs[len(kvs)-1-2*t].k, d)))
    }
    return medianThree(byteAt(kvs[0].k, d), byteAt(kvs[m].k, d), byteAt(kvs[len(kvs)-1].k, d))
}

func medianThree(i, j, k int) int {
    if i > j {
        i, j = j, i
    } // i < j
    if k < i {
        return i
    }
    if k > j {
        return j
    }
    return k
}

func maxThree(i, j, k int) int {
    max := i
    if max < j {
        max = j
    }
    if max < k {
        max = k
    }
    return max
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returnkvs 2*ceil(lg(n+1)).
func maxDepth(n int) int {
    var depth int
    for i := n; i > 0; i >>= 1 {
        depth++
    }
    return depth * 2
}

// siftDown implements the heap property on kvs[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown(kvs kvSlice, lo, hi, first int) {
    root := lo
    for {
        child := 2*root + 1
        if child >= hi {
            break
        }
        if child+1 < hi && bytes.Compare(kvs[first+child].k, kvs[first+child+1].k) == -1 {
            child++
        }
        if bytes.Compare(kvs[first+root].k, kvs[first+child].k) != -1 {
            return
        }
        swap(kvs, first+root, first+child)
        root = child
    }
}

func heapSort(kvs kvSlice, a, b int) {
    first := a
    lo := 0
    hi := b - a

    // Build heap with greatest element at top.
    for i := (hi - 1) / 2; i >= 0; i-- {
        siftDown(kvs, i, hi, first)
    }

    // Pop elements, largest first, into end of kvs.
    for i := hi - 1; i >= 0; i-- {
        swap(kvs, first, first+i)
        siftDown(kvs, lo, i, first)
    }
}

func swap(kvs kvSlice, a, b int) {
    kvs[a], kvs[b] = kvs[b], kvs[a]
}

// Compare two strings from the pos d.
func lessFrom(a, b []byte, d int) bool {
    l := len(a)
    if l > len(b) {
        l = len(b)
    }
    // skip the last byte '"'
    l -= 1
    for i := d; i < l; i++ {
        if a[i] == b[i] {
            continue
        }
        return a[i] < b[i]
    }
    return len(a) < len(b)
}

func byteAt(b []byte, p int) int {
    if p < len(b) {
        return int(b[p])
    }
    return -1
}
