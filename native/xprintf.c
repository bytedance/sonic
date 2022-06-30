#include <stdint.h>
#include <stddef.h>

static void __attribute__((naked)) write_syscall(const char *s, size_t n) {
    asm volatile (
        "movq %rsi, %rdx"           "\n"
        "movq %rdi, %rsi"           "\n"
        "movq $1, %rdi"             "\n"
        "movq $0x02000004, %rax"    "\n"
        "syscall"                   "\n"
        "retq"                      "\n"
    );
}

static void printch(char ch) {
    write_syscall(&ch, 1);
}

static void printstr(const char *s) {
    size_t n = 0;
    const char *p = s;
    while (*p++) n++;
    write_syscall(s, n);
}

static void printstrn(const char *s, size_t n) {
    write_syscall(s, n);
}

static void printint(int64_t v) {
    char neg = 0;
    char buf[32] = {};
    char *p = &buf[31];
    if (v == 0) {
        printch('0');
        return;
    }
    if (v < 0) {
        v = -v;
        neg = 1;
    }
    while (v) {
        *--p = (v % 10) + '0';
        v /= 10;
    }
    if (neg) {
        *--p = '-';
    }
    printstr(p);
}

static void printhex(uintptr_t v) {
    if (v == 0) {
        printch('0');
        return;
    }
    char buf[32] = {};
    char *p = &buf[31];
    static const char tab[] = "0123456789abcdef";
    while (v) {
        *--p = tab[v & 0x0f];
        v >>= 4;
    }
    printstr(p);
}

void xprintf(const char *fmt, ...) {
    __builtin_va_list va;
    char buf[256] = {};
    char *p = buf;
    __builtin_va_start(va, fmt);
    for (;;) {
        if (*fmt == 0) {
            break;
        }
        if (*fmt != '%') {
            *p++ = *fmt++;
            continue;
        }
        *p = 0;
        p = buf;
        fmt++;
        printstr(buf);
        switch (*fmt++) {
            case '%': {
                printch('%');
                break;
            }
            case 's': {
                printstr(__builtin_va_arg(va, const char *));
                break;
            }
            case 'd': {
                printint(__builtin_va_arg(va, int64_t));
                break;
            }
            case 'x': {
                printhex(__builtin_va_arg(va, uintptr_t));
                break;
            }
        }
    }
    __builtin_va_end(va);
    if (p != buf) {
        *p = 0;
        printstr(buf);
    }
}
