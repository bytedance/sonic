#define XXH_STATIC_LINKING_ONLY   /* access advanced declarations */
#define XXH_IMPLEMENTATION   /* access definitions */

#include <stdio.h>
#include <inttypes.h>
#include "../xxhash.h"
#include "../native.h"
#include "../hashmap_get.c"

// Set hashmap.
void field_hashmap_set(FieldHashMap *fmap, const GoString* key, int64_t id){
    XXH64_hash_t seed   =    123456789;
    uint64_t hash       =    XXH64(key->buf, key->len, seed);
    uint64_t index      =    hash % fmap->N;

    while (fmap->bucket[index].hash != 0) {
        index = (index + 1) % fmap->N;
    }

    fmap->bucket[index].name = (*key);
    fmap->bucket[index].hash = hash;
    fmap->bucket[index].id   = id;
    return;
}

int main() {
    FieldHashMap map;
    map.N      = 10000;
    map.bucket = (FieldEntry*)malloc(sizeof(FieldEntry) * map.N);
    memset(map.bucket, 0, sizeof(sizeof(FieldEntry) * map.N));

    GoString key1;
    GoString key2;
    GoString key3;

    int64_t value1;
    int64_t value2;
    int64_t value3;

    key1.buf   = "Hello";
    key1.len   = 5;
    key2.buf   = "World";
    key2.len   = 5;
    key3.buf   = "!!!!";
    key3.len   = 4;

    field_hashmap_set(&map, &key1, 1);
    field_hashmap_set(&map, &key2, 2);
    value1 = field_hashmap_get(&map, &key1);
    value2 = field_hashmap_get(&map, &key2);
    printf("The value1 is: %" PRId64 "\n", value1);
    printf("The value2 is: %" PRId64 "\n", value2);
    
    value3 = field_hashmap_get(&map, &key3);
    printf("The value3 is: %" PRId64 "\n", value3);
    field_hashmap_set(&map, &key3, 3);
    value3 = field_hashmap_get(&map, &key3);
    printf("The value3 is: %" PRId64 "\n", value3);
    free(map.bucket);
}
