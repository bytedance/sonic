#include "scanning.h"

long get_by_path(const GoString *src, long *p, const GoSlice *path, StateMachine* sm) {
    GoIface *ps = (GoIface*)(path->buf);
    GoIface *pe = (GoIface*)(path->buf) + path->len;
    char c = 0;
    int64_t index;
    long found;

query:
    /* to be safer for invalid json, use slower skip for the demanded fields */
    if (ps == pe) {
        if (sm == NULL) {
            return skip_one_fast_1(src, p);
        }
        return skip_one_1(src, p, sm, 0);
    }

    /* match type: should query key in object, query index in array */
    c = advance_ns(src, p);
    if (is_str(ps)) {
        if (c != '{') {
            goto err_inval;
        }
        goto skip_in_obj;
    } else if (is_int(ps)) {
        if (c != '[') {
            goto err_inval;
        }

        index = get_int(ps);
        if (index < 0) {
            goto err_path;
        }

        goto skip_in_arr;
    } else {
        goto err_path;
    }

skip_in_obj:
    c = advance_ns(src, p);
    if (c == '}') {
        goto not_found;
    }
    if (c != '"') {
        goto err_inval;
    }

    /* parse the object key */
    found = match_key(src, p, get_str(ps));
    if (found < 0) {
        return found; // parse string errors
    }

    /* value should after : */
    c = advance_ns(src, p);
    if (c != ':') {
        goto err_inval;
    }
    if (found) {
        ps++;
        goto query;
    }

    /* skip the unknown fields */
    skip_one_fast_1(src, p);
    c = advance_ns(src, p);
    if (c == '}') {
        goto not_found;
    }
    if (c != ',') {
        goto err_inval;
    }
    goto skip_in_obj;

skip_in_arr:
    /* check empty array */
    c = advance_ns(src, p);
    if (c == ']') {
        goto not_found;
    }
    *p -= 1;

    /* skip array elem one by one */
    while (index-- > 0) {
        skip_one_fast_1(src, p);
        c = advance_ns(src, p);
        if (c == ']') {
            goto not_found;
        }
        if (c != ',') {
            goto err_inval;
        }
    }
    ps++;
    goto query;

not_found:
    *p -= 1; // backward error position
    return -ERR_NOT_FOUND;
err_inval:
    *p -= 1;
    return -ERR_INVAL;
err_path:
    *p -= 1;
    return -ERR_UNSUPPORT_TYPE;
}