#define XXH_STATIC_LINKING_ONLY   /* access advanced declarations */
#define XXH_IMPLEMENTATION   /* access definitions */

#include "xxhash.h"

// Get hashmap, return the matching ID, and if not found, return -1.
int64_t field_hashmap_get(FieldHashMap *fmap, const GoString* key){
    XXH64_hash_t seed   =    123456789;
    size_t len          =    fmap->N;
    uint64_t hash       =    XXH64(key->name, len, seed);
    int64_t  id         =    -1;

    if(fmap.bucket[hash] != NULL){
        id              =    fmap.bucket[hash]->id;
    }
    return id;
}
