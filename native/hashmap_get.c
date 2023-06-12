#define XXH_STATIC_LINKING_ONLY   /* access advanced declarations */
#define XXH_IMPLEMENTATION   /* access definitions */

#include "xxhash.h"

// Get hashmap, 返回匹配的id，如果没有找到，就返回 -1。
int64_t field_hashmap_get(FieldHashMap *fmap, const GoString* key){
    XXH64_hash_t seed   =    123456789;
    size_t len          =    fmap->N;
    uint64_t hash       =    XXH64(key->name, len, seed);

    if(fmap.bucket[hash] != NULL){
        return fmap.bucket[hash]->id;
    }
}




