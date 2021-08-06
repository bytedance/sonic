package encoder

import (
    "bytes"
)

// Algorithm 3-way Radix Quicksort, d means the radix.
// Reference: https://algs4.cs.princeton.edu/51radix/Quick3string.java.html
func radixQsort(kvs kvSlice, d, maxDepth int) {
    if len(kvs) < 12 {
        insertRadixSort(kvs, 0, len(kvs), d)
        return
    }
    if maxDepth == 0 {
        heapSort(kvs, 0, len(kvs))
        return
    }
    maxDepth--

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
    radixQsort(kvs[:lt], d, maxDepth)
    if p > -1 {
        radixQsort(kvs[lt:gt], d+1, maxDepth)
    }
    radixQsort(kvs[gt:], d, maxDepth)
}

func insertRadixSort(kvs kvSlice, lo, hi, d int) {
    for i := lo + 1; i < hi; i++ {
        for j := i; j > lo && lessFrom(kvs[j].k, kvs[j-1].k, d); j-- {
            kvs[j], kvs[j-1] = kvs[j-1], kvs[j]
        }
    }
}

func pivot(kvs kvSlice, d int) int {
    m := len(kvs) >> 1
    if len(kvs) > 40 {
        // Tukey'kvs ``Ninther,'' median of three mediankvs of three.
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
    if i < j {
        i, j = j, i
    }
    if k < i {
        return i
    }
    if k > j {
        return j
    }
    return k
}

// maxDepth returnkvs a threshold at which quicksort should switch
// to heapsort. It returnkvs 2*ceil(lg(n+1)).
func maxDepth(n int) int {
    var depth int
    for i := n; i > 0; i >>= 1 {
        depth++
    }
    return depth * 2
}

// siftDown implementkvs the heap property on kvs[lo:hi].
// first ikvs an offset into the array where the root of the heap lies.
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

    // Pop elementkvs, largest first, into end of kvs.
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
