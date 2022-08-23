/*
 * Copyright 2022 ByteDance Inc.
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

#ifndef XPRINTF_H
#define XPRINTF_H

#include <sys/types.h>

static void __attribute__((naked)) write_syscall(const char *s, size_t n)
{
    asm volatile(
        "movq %rsi, %rdx"
        "\n"
        "movq %rdi, %rsi"
        "\n"
        "movq $1, %rdi"
        "\n"
        "movq $0x02000004, %rax"
        "\n"
        "syscall"
        "\n"
        "retq"
        "\n");
}

static void printch(const char ch)
{
    write_syscall(&ch, 1);
}

static void printstr(const char *s)
{
    size_t n = 0;
    const char *p = s;
    while (*p++)
        n++;
    write_syscall(s, n);
}

static void printint(int64_t v)
{
    char neg = 0;
    char buf[32] = {};
    char *p = &buf[31];
    uint64_t u;
    if (v < 0) {
        u = ~v + 1;
        neg = 1;
    } else {
        u = v;
    }
    while (u)
    {
        *--p = (u % 10) + '0';
        u /= 10;
    }
    if (neg) {
        *--p = '-';
    }
    printstr(p);
}

static void printuint(uint64_t v)
{
    char neg = 0;
    char buf[32] = {};
    char *p = &buf[31];
    if (v == 0)
    {
        printch('0');
        return;
    }
    while (v)
    {
        *--p = (v % 10) + '0';
        v /= 10;
    }
    printstr(p);
}

static const char tab[] = "0123456789abcdef";

static void printhex(uintptr_t v)
{
    if (v == 0)
    {
        printch('0');
        return;
    }
    char buf[32] = {};
    char *p = &buf[31];

    while (v)
    {
        *--p = tab[v & 0x0f];
        v >>= 4;
    }
    printstr(p);
}

#define MAX_BUF_LEN 100

static void printbytes(GoSlice *s)
{
    printch('[');
    int i = 0;
    if (s->len > MAX_BUF_LEN)
    {
        i = s->len - MAX_BUF_LEN;
    }
    for (; i < s->len; i++)
    {
        printch(tab[((s->buf[i]) & 0xf0) >> 4]);
        printch(tab[(s->buf[i]) & 0x0f]);
        if (i != s->len - 1)
            printch(',');
    }
    printch(']');
}

static void printgostr(GoString *s)
{
    printch('"');
    if (s->len < MAX_BUF_LEN)
    {
        write_syscall(s->buf, s->len);
    }
    else
    {
        write_syscall(&s->buf[s->len - MAX_BUF_LEN], MAX_BUF_LEN);
    }
    printch('"');
}

static void xprintf(const char *fmt, ...)
{
#ifdef DEBUG
    __builtin_va_list va;
    char buf[256] = {};
    char *p = buf;
    __builtin_va_start(va, fmt);
    for (;;)
    {
        if (*fmt == 0)
        {
            break;
        }
        if (*fmt != '%')
        {
            *p++ = *fmt++;
            continue;
        }
        *p = 0;
        p = buf;
        fmt++;
        printstr(buf);
        switch (*fmt++)
        {
        case '%':
        {
            printch('%');
            break;
        }
        case 'g':
        {
            printgostr(__builtin_va_arg(va, GoString *));
            break;
        }
        case 's':
        {
            printstr(__builtin_va_arg(va, const char *));
            break;
        }
        case 'd':
        {
            printint(__builtin_va_arg(va, int64_t));
            break;
        }
        case 'u':
        {
            printuint(__builtin_va_arg(va, uint64_t));
            break;
        }
        case 'f':
        {
            printint(__builtin_va_arg(va, double));
            break;
        }
        case 'c':
        {
            printch((char)(__builtin_va_arg(va, int)));
            break;
        }
        case 'x':
        {
            printhex(__builtin_va_arg(va, uintptr_t));
            break;
        }
        case 'l':
        {
            printbytes(__builtin_va_arg(va, GoSlice *));
            break;
        }
        }
    }
    __builtin_va_end(va);
    if (p != buf)
    {
        *p = 0;
        printstr(buf);
    }
#endif
}

#endif // XPRINTF_H