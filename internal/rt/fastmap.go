//go:build !goexperiment.swisstable && go1.16 && !go1.23
// +build !goexperiment.swisstable,go1.16,!go1.23

// 	Most Codes are copied from Go1.20.4. Modified parts pls see comments with #MODIFIED

package rt

import (
	"unsafe"

	"github.com/bytedance/sonic/internal/envs"
)

var EnbaleFastMap bool = envs.UseFastMap

type bucketArray struct {
	ptr 	unsafe.Pointer
	len    int
}

// TODO: add a tools to generate map info.
/*
	Continuous Memory layout:
	headers: | hmap1 | hmap2 | ... 
	buckets: | btk1  | btk2  | ...
*/
type MapPool struct {
	headers		[]Hmap
	hdrIndex 	int

	buckets     bucketArray
	bktIndex 	int
	typ         *GoMapType
}

func NewMapPool(t *GoMapType, mapHint int, kvHint int) MapPool {
	var h []Hmap
	var b bucketArray
	if mapHint != 0 {
		h = make([]Hmap, mapHint, mapHint)
	}

	btkHint := kvHint / 6

	// XXX: should not including the empty map?
	// make sure all map will have at least one bucket
	if btkHint < mapHint {
		btkHint = mapHint
	}
	if kvHint != 0 {
		b = bucketArray{ptr: newarray(t.Bucket, btkHint), len: btkHint}
	}

	return MapPool{
		headers: h,
		hdrIndex: 0,
		buckets:  b,
		bktIndex: 0,
		typ: t,
	}
}

func (self *MapPool) GetMap(hint int) *Hmap {
	mp := makemap(self.typ, hint, self)
	return mp
}

func (self *MapPool) Remain() (hdrs int, btks int) {
	hdrs = len(self.headers) - self.hdrIndex
	btks = self.buckets.len - self.bktIndex
	return
}

func (self *MapPool) allocMapHeader() *Hmap {
	h := &(self.headers)[self.hdrIndex]
	self.hdrIndex++
	return h
}

func (self *MapPool) allocBucketArray(b int) unsafe.Pointer {
	if self.bktIndex + b - 1 >= self.buckets.len {
		return newarray(self.typ.Bucket, int(b))
	}
	btk := PtrAdd(self.buckets.ptr, uintptr(self.bktIndex) * uintptr(self.typ.BucketSize))
	self.bktIndex += b
	return btk
}

const (
	// FIXME: hardcode here.
	heapAddrBits = 48
	ptrSize = 8


	_64bit = 1 << (^uintptr(0) >> 63) / 2
	maxAlloc = (1 << heapAddrBits) - (1-_64bit)*1

	// Maximum number of key/elem pairs a bucket can hold.
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

// A header for a Go map. 
type Hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	Buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	Oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra 	*mapextra // optional fields
}

// // A header for a Go map.
// type hmap struct {
// 	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
// 	// Make sure this stays in sync with the compiler's definition.
// 	count     int // # live cells == size of map.  Must be first (used by len() builtin)
// 	flags     uint8
// 	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
// 	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
// 	hash0     uint32 // hash seed

// 	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
// 	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
// 	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

// 	extra *mapextra // optional fields
// }

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// If both key and elem do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and elem do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}


// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt elems.
	// NOTE: packing all the keys together and then all the elems together makes the
	// code a bit more complicated than alternating key/elem/key/elem/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}

//go:linkname MulUintptr  runtime/internal/math.MulUintptr
func MulUintptr(a, b uintptr) (uintptr, bool) 

//go:linkname fastrand  runtime.fastrand
func fastrand() uint32

//go:linkname overLoadFactor  runtime.overLoadFactor
func overLoadFactor(count int, B uint8) bool

//go:linkname newarray runtime.newarray
func newarray(typ *GoType, n int) unsafe.Pointer 

// makemap implements Go map creation for make(map[k]v, hint).
// If the compiler has determined that the map or the first bucket
// can be created on the stack, h and/or bucket may be non-nil.
// If h != nil, the map can be created directly in h.
// If h.buckets != nil, bucket pointed to can be used as the first bucket.
func makemap(t *GoMapType, hint int, pool *MapPool) *Hmap {
	// //println(" makemap  is ", hint)
	mem, overflow := MulUintptr(uintptr(hint), t.Bucket.Size)
	if overflow || mem > maxAlloc {
		hint = 0
	}

	// initialize Hmap
	// if h == nil {
	// 	h = new(hmap)
	// }
	// #MODIFIED: use pool at fist, if not then fallback
	h := pool.allocMapHeader()
	if h == nil {
		h = new(Hmap)
	}
	h.hash0 = fastrand()

	// Find the size parameter B which will hold the requested # of elements.
	// For hint < 0 overLoadFactor returns false since hint < bucketCnt.
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}
	h.B = B

	// allocate initial hash table
	// if B == 0, the buckets field is allocated lazily later (in mapassign)
	// If hint is large zeroing this memory could take a while.
	// #MODIFIED: we always initialized the map
	// if h.B != 0 {
		var nextOverflow *bmap
		h.Buckets, nextOverflow = makeBucketArray(t, h.B, pool)
		if nextOverflow != nil {
			h.extra = new(mapextra)
			h.extra.nextOverflow = nextOverflow
		}
	// }

	return h
}

// used for tests
func calcaulateB(hint int) int {
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}

	b := B
	base := bucketShift(b)
	nbuckets := base

	if b >= 4 {
		nbuckets += bucketShift(b - 4)
		sz := MapType(MapEfaceType).Bucket.Size * nbuckets
		up := roundupsize(sz)
		if up != sz {
			nbuckets = up /  MapType(MapEfaceType).Bucket.Size
		}
	}

	return int(nbuckets)
}

// bucketShift returns 1<<b, optimized for code generation.
func bucketShift(b uint8) uintptr {
	// Masking the shift amount allows overflow checks to be elided.
	return uintptr(1) << (b & (ptrSize*8 - 1))
}

func (b *bmap) setoverflow(t *GoMapType, ovf *bmap) {
	*(**bmap)(add(unsafe.Pointer(b), uintptr(t.BucketSize)- ptrSize)) = ovf
}

// makeBucketArray initializes a backing array for map buckets.
// 1<<b is the minimum number of buckets to allocate.
// dirtyalloc should either be nil or a bucket array previously
// allocated by makeBucketArray with the same t and b parameters.
// If dirtyalloc is nil a new backing array will be alloced and
// otherwise dirtyalloc will be cleared and reused as backing array.
func makeBucketArray(t *GoMapType, b uint8, pool *MapPool) (buckets unsafe.Pointer, nextOverflow *bmap) {
	base := bucketShift(b)
	nbuckets := base
	// For small b, overflow buckets are unlikely.
	// Avoid the overhead of the calculation.
	if b >= 4 {
		// Add on the estimated number of overflow buckets
		// required to insert the median number of elements
		// used with this value of b.
		nbuckets += bucketShift(b - 4)
		sz := t.Bucket.Size * nbuckets
		up := roundupsize(sz)
		if up != sz {
			nbuckets = up / t.Bucket.Size
		}
	}


	// if dirtyalloc == nil {
	// 	buckets = newarray(t.Bucket, int(nbuckets))
	// } else {
	// 	// dirtyalloc was previously generated by
	// 	// the above newarray(t.bucket, int(nbuckets))
	// 	// but may not be empty.
	// 	buckets = dirtyalloc
	// 	size := t.Bucket.Size * nbuckets
	// 	if t.Bucket.PtrData != 0 {
	// 		MemclrHasPointers(buckets, size)
	// 	} else {
	// 		MemclrNoHeapPointers(buckets, size)
	// 	}
	// }

	// #MODIFIED: use pool at first, if not fallback it
	buckets = pool.allocBucketArray(int(nbuckets))

	if base != nbuckets {
		// We preallocated some overflow buckets.
		// To keep the overhead of tracking these overflow buckets to a minimum,
		// we use the convention that if a preallocated overflow bucket's overflow
		// pointer is nil, then there are more available by bumping the pointer.
		// We need a safe non-nil pointer for the last overflow bucket; just use buckets.
		nextOverflow = (*bmap)(add(buckets, base*uintptr(t.BucketSize)))
		last := (*bmap)(add(buckets, (nbuckets-1)*uintptr(t.BucketSize)))
		// //println("last ptr is ", last, " ovcerflow ptr is ", nextOverflow)
		last.setoverflow(t, (*bmap)(buckets))
	}
	return buckets, nextOverflow
}
