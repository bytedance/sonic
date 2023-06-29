#define XXH_STATIC_LINKING_ONLY   /* access advanced declarations */
#define XXH_IMPLEMENTATION   /* access definitions */

#include <string.h>
#include "native.h"
#include "xxhash.h"

// Get hashmap, return the matching ID, and if not found, return -1.
int64_t field_hashmap_get(FieldHashMap *fmap, const GoString* key){
    XXH64_hash_t seed   =    123456789;
    uint64_t hash       =    XXH64(key->buf, key->len, seed);
    uint64_t index      =    hash % fmap->N;

    while (fmap->bucket[index].hash != 0) {
        if(fmap->bucket[index].hash == hash \
        && fmap->bucket[index].name.len == key->len \
        && memcmp(fmap->bucket[index].name.buf, key->buf, key->len) == 0){
            return fmap->bucket[index].id;
        } else {
            index = (index + 1) % fmap->N;
        }
    }
    return -1;
}
