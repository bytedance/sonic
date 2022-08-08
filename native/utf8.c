/*
 * Copyright (C) 2019 Yaoyuan <ibireme@gmail.com>.
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
 *
 * This file may have been modified by ByteDance authors. All ByteDance
 * Modifications are Copyright 2022 ByteDance Authors.
 */

static inline ssize_t valid_utf8_4byte(uint32_t ubin) {
    /*
     Each unicode code point is encoded as 1 to 4 bytes in UTF-8 encoding,
     we use 4-byte mask and pattern value to validate UTF-8 byte sequence,
     this requires the input data to have 4-byte zero padding.
     ---------------------------------------------------
     1 byte
     unicode range [U+0000, U+007F]
     unicode min   [.......0]
     unicode max   [.1111111]
     bit pattern   [0.......]
     ---------------------------------------------------
     2 byte
     unicode range [U+0080, U+07FF]
     unicode min   [......10 ..000000]
     unicode max   [...11111 ..111111]
     bit require   [...xxxx. ........] (1E 00)
     bit mask      [xxx..... xx......] (E0 C0)
     bit pattern   [110..... 10......] (C0 80)
     // 1101 0100 10110000
     // 0001 1110
     ---------------------------------------------------
     3 byte
     unicode range [U+0800, U+FFFF]
     unicode min   [........ ..100000 ..000000]
     unicode max   [....1111 ..111111 ..111111]
     bit require   [....xxxx ..x..... ........] (0F 20 00)
     bit mask      [xxxx.... xx...... xx......] (F0 C0 C0)
     bit pattern   [1110.... 10...... 10......] (E0 80 80)
     ---------------------------------------------------
     3 byte invalid (reserved for surrogate halves)
     unicode range [U+D800, U+DFFF]
     unicode min   [....1101 ..100000 ..000000]
     unicode max   [....1101 ..111111 ..111111]
     bit mask      [....xxxx ..x..... ........] (0F 20 00)
     bit pattern   [....1101 ..1..... ........] (0D 20 00)
     ---------------------------------------------------
     4 byte
     unicode range [U+10000, U+10FFFF]
     unicode min   [........ ...10000 ..000000 ..000000]
     unicode max   [.....100 ..001111 ..111111 ..111111]
     bit err0      [.....100 ........ ........ ........] (04 00 00 00)
     bit err1      [.....011 ..110000 ........ ........] (03 30 00 00)
     bit require   [.....xxx ..xx.... ........ ........] (07 30 00 00)
     bit mask      [xxxxx... xx...... xx...... xx......] (F8 C0 C0 C0)
     bit pattern   [11110... 10...... 10...... 10......] (F0 80 80 80)
     ---------------------------------------------------
     */
    const uint32_t b1_mask = 0x00000080UL;
    const uint32_t b1_patt = 0x00000000UL;
    const uint32_t b2_mask = 0x0000C0E0UL;
    const uint32_t b2_patt = 0x000080C0UL;
    const uint32_t b2_requ = 0x0000001EUL;
    const uint32_t b3_mask = 0x00C0C0F0UL;
    const uint32_t b3_patt = 0x008080E0UL;
    const uint32_t b3_requ = 0x0000200FUL;
    const uint32_t b3_erro = 0x0000200DUL;
    const uint32_t b4_mask = 0xC0C0C0F8UL;
    const uint32_t b4_patt = 0x808080F0UL;
    const uint32_t b4_requ = 0x00003007UL;
    const uint32_t b4_err0 = 0x00000004UL;
    const uint32_t b4_err1 = 0x00003003UL;

#define is_valid_seq_1(uni) ( \
    ((uni & b1_mask) == b1_patt) \
)

#define is_valid_seq_2(uni) ( \
    ((uni & b2_mask) == b2_patt) && \
    ((uni & b2_requ)) \
)
    
#define is_valid_seq_3(uni) ( \
    ((uni & b3_mask) == b3_patt) && \
    ((tmp = (uni & b3_requ))) && \
    ((tmp != b3_erro)) \
)
    
#define is_valid_seq_4(uni) ( \
    ((uni & b4_mask) == b4_patt) && \
    ((tmp = (uni & b4_requ))) && \
    ((tmp & b4_err0) == 0 || (tmp & b4_err1) == 0) \
)
    uint32_t tmp = 0;
   
    if (is_valid_seq_3(ubin)) return 3;
    if (is_valid_seq_2(ubin)) return 2;
    if (is_valid_seq_4(ubin)) return 4;
    return 0;
}

static inline uint32_t less4byte_to_uint32(const char* sp, size_t nb) {
    if (nb == 1) return *(uint8_t*)sp;
    if (nb == 2) return *(uint16_t*)sp;
    uint32_t hi_1 = (*(uint8_t*)(sp + 2));
    uint32_t lo_2 = *(uint16_t*)(sp);
    return hi_1 << 16 | lo_2;
}