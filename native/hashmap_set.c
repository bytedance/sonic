#define XXH_STATIC_LINKING_ONLY   /* access advanced declarations */
#define XXH_IMPLEMENTATION   /* access definitions */

#include "xxhash.h"

void field_hashmap_set(FieldHashMap *fmap, const GoString* key, int64_t id){
    XXH64_hash_t seed   =    123456789;
    size_t len          =    fmap->N;
    uint64_t hash       =    XXH64(key->name, len, seed);
    FieldEntry bucket;
    bucket.name         =    key;
    bucket.hash         =    hash;
    bucket.id           =    id;
    fmap.bucket[hash]   =    bucket;
    return;
}
