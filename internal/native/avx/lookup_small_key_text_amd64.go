// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package avx

var _text_lookup_small_key = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, // QUAD $0x4040404040404040; QUAD $0x4040404040404040  // .space 16, '@@@@@@@@@@@@@@@@'
	//0x00000010 LCPI0_1
	0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, 0x5b, //0x00000010 QUAD $0x5b5b5b5b5b5b5b5b; QUAD $0x5b5b5b5b5b5b5b5b  // .space 16, '[[[[[[[[[[[[[[[['
	//0x00000020 LCPI0_2
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, //0x00000020 QUAD $0x0101010101010101; QUAD $0x0101010101010101  // .space 16, '\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01'
	//0x00000030 .p2align 4, 0x90
	//0x00000030 _lookup_small_key
	0x55, //0x00000030 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000031 movq         %rsp, %rbp
	0x41, 0x57, //0x00000034 pushq        %r15
	0x41, 0x56, //0x00000036 pushq        %r14
	0x41, 0x55, //0x00000038 pushq        %r13
	0x41, 0x54, //0x0000003a pushq        %r12
	0x53, //0x0000003c pushq        %rbx
	0x48, 0x83, 0xec, 0x10, //0x0000003d subq         $16, %rsp
	0x4c, 0x8b, 0x57, 0x08, //0x00000041 movq         $8(%rdi), %r10
	0x4c, 0x8b, 0x1e, //0x00000045 movq         (%rsi), %r11
	0x45, 0x0f, 0xb6, 0xc2, //0x00000048 movzbl       %r10b, %r8d
	0x4b, 0x8d, 0x0c, 0x80, //0x0000004c leaq         (%r8,%r8,4), %rcx
	0x45, 0x0f, 0xb6, 0x0c, 0x0b, //0x00000050 movzbl       (%r11,%rcx), %r9d
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000055 movq         $-1, %rax
	0x45, 0x85, 0xc9, //0x0000005c testl        %r9d, %r9d
	0x0f, 0x84, 0x20, 0x03, 0x00, 0x00, //0x0000005f je           LBB0_39
	0x48, 0x89, 0x55, 0xc8, //0x00000065 movq         %rdx, $-56(%rbp)
	0x4c, 0x8b, 0x3f, //0x00000069 movq         (%rdi), %r15
	0x41, 0x8b, 0x44, 0x0b, 0x01, //0x0000006c movl         $1(%r11,%rcx), %eax
	0x48, 0x89, 0x45, 0xd0, //0x00000071 movq         %rax, $-48(%rbp)
	0x8d, 0xb0, 0xa5, 0x00, 0x00, 0x00, //0x00000075 leal         $165(%rax), %esi
	0x4c, 0x01, 0xde, //0x0000007b addq         %r11, %rsi
	0x41, 0x0f, 0xb6, 0xca, //0x0000007e movzbl       %r10b, %ecx
	0x41, 0x83, 0xf8, 0x09, //0x00000082 cmpl         $9, %r8d
	0x0f, 0x83, 0xd0, 0x00, 0x00, 0x00, //0x00000086 jae          LBB0_2
	0x45, 0x8a, 0x27, //0x0000008c movb         (%r15), %r12b
	0x45, 0x8d, 0x68, 0x01, //0x0000008f leal         $1(%r8), %r13d
	0x44, 0x89, 0xcb, //0x00000093 movl         %r9d, %ebx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000096 .p2align 4, 0x90
	//0x000000a0 LBB0_7
	0x44, 0x38, 0x26, //0x000000a0 cmpb         %r12b, (%rsi)
	0x0f, 0x85, 0x97, 0x00, 0x00, 0x00, //0x000000a3 jne          LBB0_8
	0x44, 0x0f, 0xb6, 0x76, 0x01, //0x000000a9 movzbl       $1(%rsi), %r14d
	0xbf, 0x01, 0x00, 0x00, 0x00, //0x000000ae movl         $1, %edi
	0x45, 0x3a, 0x77, 0x01, //0x000000b3 cmpb         $1(%r15), %r14b
	0x0f, 0x85, 0x85, 0x00, 0x00, 0x00, //0x000000b7 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x02, //0x000000bd movzbl       $2(%rsi), %edx
	0xbf, 0x02, 0x00, 0x00, 0x00, //0x000000c1 movl         $2, %edi
	0x41, 0x3a, 0x57, 0x02, //0x000000c6 cmpb         $2(%r15), %dl
	0x0f, 0x85, 0x72, 0x00, 0x00, 0x00, //0x000000ca jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x03, //0x000000d0 movzbl       $3(%rsi), %edx
	0xbf, 0x03, 0x00, 0x00, 0x00, //0x000000d4 movl         $3, %edi
	0x41, 0x3a, 0x57, 0x03, //0x000000d9 cmpb         $3(%r15), %dl
	0x0f, 0x85, 0x5f, 0x00, 0x00, 0x00, //0x000000dd jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x04, //0x000000e3 movzbl       $4(%rsi), %edx
	0xbf, 0x04, 0x00, 0x00, 0x00, //0x000000e7 movl         $4, %edi
	0x41, 0x3a, 0x57, 0x04, //0x000000ec cmpb         $4(%r15), %dl
	0x0f, 0x85, 0x4c, 0x00, 0x00, 0x00, //0x000000f0 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x05, //0x000000f6 movzbl       $5(%rsi), %edx
	0xbf, 0x05, 0x00, 0x00, 0x00, //0x000000fa movl         $5, %edi
	0x41, 0x3a, 0x57, 0x05, //0x000000ff cmpb         $5(%r15), %dl
	0x0f, 0x85, 0x39, 0x00, 0x00, 0x00, //0x00000103 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x06, //0x00000109 movzbl       $6(%rsi), %edx
	0xbf, 0x06, 0x00, 0x00, 0x00, //0x0000010d movl         $6, %edi
	0x41, 0x3a, 0x57, 0x06, //0x00000112 cmpb         $6(%r15), %dl
	0x0f, 0x85, 0x26, 0x00, 0x00, 0x00, //0x00000116 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x07, //0x0000011c movzbl       $7(%rsi), %edx
	0x31, 0xff, //0x00000120 xorl         %edi, %edi
	0x41, 0x3a, 0x57, 0x07, //0x00000122 cmpb         $7(%r15), %dl
	0x40, 0x0f, 0x94, 0xc7, //0x00000126 sete         %dil
	0x48, 0x83, 0xc7, 0x07, //0x0000012a addq         $7, %rdi
	0xe9, 0x0f, 0x00, 0x00, 0x00, //0x0000012e jmp          LBB0_16
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000133 .p2align 4, 0x90
	//0x00000140 LBB0_8
	0x31, 0xff, //0x00000140 xorl         %edi, %edi
	//0x00000142 LBB0_16
	0x48, 0x39, 0xcf, //0x00000142 cmpq         %rcx, %rdi
	0x0f, 0x83, 0x91, 0x01, 0x00, 0x00, //0x00000145 jae          LBB0_17
	0x4c, 0x01, 0xee, //0x0000014b addq         %r13, %rsi
	0x83, 0xc3, 0xff, //0x0000014e addl         $-1, %ebx
	0x0f, 0x85, 0x49, 0xff, 0xff, 0xff, //0x00000151 jne          LBB0_7
	0xe9, 0x51, 0x00, 0x00, 0x00, //0x00000157 jmp          LBB0_20
	//0x0000015c LBB0_2
	0xc4, 0xc1, 0x7a, 0x6f, 0x07, //0x0000015c vmovdqu      (%r15), %xmm0
	0xc4, 0xc1, 0x7a, 0x6f, 0x4f, 0x10, //0x00000161 vmovdqu      $16(%r15), %xmm1
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x00000167 movq         $-1, %rdi
	0x48, 0xd3, 0xe7, //0x0000016e shlq         %cl, %rdi
	0x45, 0x8d, 0x60, 0x01, //0x00000171 leal         $1(%r8), %r12d
	0x44, 0x89, 0xcb, //0x00000175 movl         %r9d, %ebx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000178 .p2align 4, 0x90
	//0x00000180 LBB0_3
	0xc5, 0xf9, 0x74, 0x16, //0x00000180 vpcmpeqb     (%rsi), %xmm0, %xmm2
	0xc5, 0xf9, 0xd7, 0xd2, //0x00000184 vpmovmskb    %xmm2, %edx
	0xc5, 0xf1, 0x74, 0x56, 0x10, //0x00000188 vpcmpeqb     $16(%rsi), %xmm1, %xmm2
	0xc5, 0xf9, 0xd7, 0xc2, //0x0000018d vpmovmskb    %xmm2, %eax
	0xc1, 0xe0, 0x10, //0x00000191 shll         $16, %eax
	0x09, 0xd0, //0x00000194 orl          %edx, %eax
	0x09, 0xf8, //0x00000196 orl          %edi, %eax
	0x83, 0xf8, 0xff, //0x00000198 cmpl         $-1, %eax
	0x0f, 0x84, 0x4a, 0x01, 0x00, 0x00, //0x0000019b je           LBB0_4
	0x4c, 0x01, 0xe6, //0x000001a1 addq         %r12, %rsi
	0x83, 0xc3, 0xff, //0x000001a4 addl         $-1, %ebx
	0x0f, 0x85, 0xd3, 0xff, 0xff, 0xff, //0x000001a7 jne          LBB0_3
	//0x000001ad LBB0_20
	0x48, 0x8b, 0x45, 0xd0, //0x000001ad movq         $-48(%rbp), %rax
	0x48, 0x03, 0x45, 0xc8, //0x000001b1 addq         $-56(%rbp), %rax
	0x49, 0x01, 0xc3, //0x000001b5 addq         %rax, %r11
	0xc5, 0xfa, 0x6f, 0x0d, 0x40, 0xfe, 0xff, 0xff, //0x000001b8 vmovdqu      $-448(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	0xc4, 0xc1, 0x7a, 0x6f, 0x07, //0x000001c0 vmovdqu      (%r15), %xmm0
	0xc5, 0xf9, 0x64, 0xd9, //0x000001c5 vpcmpgtb     %xmm1, %xmm0, %xmm3
	0xc5, 0xfa, 0x6f, 0x15, 0x3f, 0xfe, 0xff, 0xff, //0x000001c9 vmovdqu      $-449(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	0xc5, 0xe9, 0x64, 0xe0, //0x000001d1 vpcmpgtb     %xmm0, %xmm2, %xmm4
	0xc5, 0xd9, 0xdb, 0xe3, //0x000001d5 vpand        %xmm3, %xmm4, %xmm4
	0xc5, 0xfa, 0x6f, 0x1d, 0x3f, 0xfe, 0xff, 0xff, //0x000001d9 vmovdqu      $-449(%rip), %xmm3  /* LCPI0_2+0(%rip) */
	0xc5, 0xd9, 0xdb, 0xe3, //0x000001e1 vpand        %xmm3, %xmm4, %xmm4
	0xc5, 0xd9, 0x71, 0xf4, 0x05, //0x000001e5 vpsllw       $5, %xmm4, %xmm4
	0xc5, 0xd9, 0xfc, 0xc0, //0x000001ea vpaddb       %xmm0, %xmm4, %xmm0
	0x41, 0x0f, 0xb6, 0xca, //0x000001ee movzbl       %r10b, %ecx
	0x41, 0x83, 0xf8, 0x09, //0x000001f2 cmpl         $9, %r8d
	0x0f, 0x83, 0xfa, 0x00, 0x00, 0x00, //0x000001f6 jae          LBB0_21
	0xc4, 0xe3, 0x79, 0x14, 0xc2, 0x01, //0x000001fc vpextrb      $1, %xmm0, %edx
	0xc4, 0xc3, 0x79, 0x14, 0xc4, 0x02, //0x00000202 vpextrb      $2, %xmm0, %r12d
	0xc4, 0xc3, 0x79, 0x14, 0xc5, 0x03, //0x00000208 vpextrb      $3, %xmm0, %r13d
	0xc4, 0xc3, 0x79, 0x14, 0xc7, 0x04, //0x0000020e vpextrb      $4, %xmm0, %r15d
	0xc4, 0xc3, 0x79, 0x14, 0xc2, 0x05, //0x00000214 vpextrb      $5, %xmm0, %r10d
	0xc4, 0xc3, 0x79, 0x14, 0xc6, 0x06, //0x0000021a vpextrb      $6, %xmm0, %r14d
	0xc5, 0xf9, 0x7e, 0xc3, //0x00000220 vmovd        %xmm0, %ebx
	0xc4, 0xe3, 0x79, 0x14, 0xc0, 0x07, //0x00000224 vpextrb      $7, %xmm0, %eax
	0x41, 0x83, 0xc0, 0x01, //0x0000022a addl         $1, %r8d
	0x41, 0x83, 0xf9, 0x02, //0x0000022e cmpl         $2, %r9d
	0xbf, 0x01, 0x00, 0x00, 0x00, //0x00000232 movl         $1, %edi
	0x41, 0x0f, 0x43, 0xf9, //0x00000237 cmovael      %r9d, %edi
	0x90, 0x90, 0x90, 0x90, 0x90, //0x0000023b .p2align 4, 0x90
	//0x00000240 LBB0_25
	0x41, 0x38, 0x1b, //0x00000240 cmpb         %bl, (%r11)
	0x0f, 0x85, 0x77, 0x00, 0x00, 0x00, //0x00000243 jne          LBB0_26
	0xbe, 0x01, 0x00, 0x00, 0x00, //0x00000249 movl         $1, %esi
	0x41, 0x38, 0x53, 0x01, //0x0000024e cmpb         %dl, $1(%r11)
	0x0f, 0x85, 0x6a, 0x00, 0x00, 0x00, //0x00000252 jne          LBB0_34
	0xbe, 0x02, 0x00, 0x00, 0x00, //0x00000258 movl         $2, %esi
	0x45, 0x38, 0x63, 0x02, //0x0000025d cmpb         %r12b, $2(%r11)
	0x0f, 0x85, 0x5b, 0x00, 0x00, 0x00, //0x00000261 jne          LBB0_34
	0xbe, 0x03, 0x00, 0x00, 0x00, //0x00000267 movl         $3, %esi
	0x45, 0x38, 0x6b, 0x03, //0x0000026c cmpb         %r13b, $3(%r11)
	0x0f, 0x85, 0x4c, 0x00, 0x00, 0x00, //0x00000270 jne          LBB0_34
	0xbe, 0x04, 0x00, 0x00, 0x00, //0x00000276 movl         $4, %esi
	0x45, 0x38, 0x7b, 0x04, //0x0000027b cmpb         %r15b, $4(%r11)
	0x0f, 0x85, 0x3d, 0x00, 0x00, 0x00, //0x0000027f jne          LBB0_34
	0xbe, 0x05, 0x00, 0x00, 0x00, //0x00000285 movl         $5, %esi
	0x45, 0x38, 0x53, 0x05, //0x0000028a cmpb         %r10b, $5(%r11)
	0x0f, 0x85, 0x2e, 0x00, 0x00, 0x00, //0x0000028e jne          LBB0_34
	0xbe, 0x06, 0x00, 0x00, 0x00, //0x00000294 movl         $6, %esi
	0x45, 0x38, 0x73, 0x06, //0x00000299 cmpb         %r14b, $6(%r11)
	0x0f, 0x85, 0x1f, 0x00, 0x00, 0x00, //0x0000029d jne          LBB0_34
	0x31, 0xf6, //0x000002a3 xorl         %esi, %esi
	0x41, 0x38, 0x43, 0x07, //0x000002a5 cmpb         %al, $7(%r11)
	0x40, 0x0f, 0x94, 0xc6, //0x000002a9 sete         %sil
	0x48, 0x83, 0xc6, 0x07, //0x000002ad addq         $7, %rsi
	0xe9, 0x0c, 0x00, 0x00, 0x00, //0x000002b1 jmp          LBB0_34
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000002b6 .p2align 4, 0x90
	//0x000002c0 LBB0_26
	0x31, 0xf6, //0x000002c0 xorl         %esi, %esi
	//0x000002c2 LBB0_34
	0x48, 0x39, 0xce, //0x000002c2 cmpq         %rcx, %rsi
	0x0f, 0x83, 0xa0, 0x00, 0x00, 0x00, //0x000002c5 jae          LBB0_35
	0x4d, 0x01, 0xc3, //0x000002cb addq         %r8, %r11
	0x83, 0xc7, 0xff, //0x000002ce addl         $-1, %edi
	0x0f, 0x85, 0x69, 0xff, 0xff, 0xff, //0x000002d1 jne          LBB0_25
	0xe9, 0x83, 0x00, 0x00, 0x00, //0x000002d7 jmp          LBB0_38
	//0x000002dc LBB0_17
	0x4c, 0x01, 0xee, //0x000002dc addq         %r13, %rsi
	0x48, 0x83, 0xc6, 0xff, //0x000002df addq         $-1, %rsi
	0x0f, 0xb6, 0x06, //0x000002e3 movzbl       (%rsi), %eax
	0xe9, 0x9a, 0x00, 0x00, 0x00, //0x000002e6 jmp          LBB0_39
	//0x000002eb LBB0_4
	0x48, 0x01, 0xce, //0x000002eb addq         %rcx, %rsi
	0x0f, 0xb6, 0x06, //0x000002ee movzbl       (%rsi), %eax
	0xe9, 0x8f, 0x00, 0x00, 0x00, //0x000002f1 jmp          LBB0_39
	//0x000002f6 LBB0_21
	0xc4, 0xc1, 0x7a, 0x6f, 0x67, 0x10, //0x000002f6 vmovdqu      $16(%r15), %xmm4
	0xc5, 0xd9, 0x64, 0xc9, //0x000002fc vpcmpgtb     %xmm1, %xmm4, %xmm1
	0xc5, 0xe9, 0x64, 0xd4, //0x00000300 vpcmpgtb     %xmm4, %xmm2, %xmm2
	0xc5, 0xe9, 0xdb, 0xc9, //0x00000304 vpand        %xmm1, %xmm2, %xmm1
	0xc5, 0xf1, 0xdb, 0xcb, //0x00000308 vpand        %xmm3, %xmm1, %xmm1
	0xc5, 0xf1, 0x71, 0xf1, 0x05, //0x0000030c vpsllw       $5, %xmm1, %xmm1
	0xc5, 0xf1, 0xfc, 0xcc, //0x00000311 vpaddb       %xmm4, %xmm1, %xmm1
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000315 movq         $-1, %rax
	0x48, 0xd3, 0xe0, //0x0000031c shlq         %cl, %rax
	0x41, 0x83, 0xc0, 0x01, //0x0000031f addl         $1, %r8d
	0x41, 0x83, 0xf9, 0x02, //0x00000323 cmpl         $2, %r9d
	0xba, 0x01, 0x00, 0x00, 0x00, //0x00000327 movl         $1, %edx
	0x41, 0x0f, 0x43, 0xd1, //0x0000032c cmovael      %r9d, %edx
	//0x00000330 .p2align 4, 0x90
	//0x00000330 LBB0_22
	0xc4, 0xc1, 0x79, 0x74, 0x13, //0x00000330 vpcmpeqb     (%r11), %xmm0, %xmm2
	0xc5, 0xf9, 0xd7, 0xf2, //0x00000335 vpmovmskb    %xmm2, %esi
	0xc4, 0xc1, 0x71, 0x74, 0x53, 0x10, //0x00000339 vpcmpeqb     $16(%r11), %xmm1, %xmm2
	0xc5, 0xf9, 0xd7, 0xfa, //0x0000033f vpmovmskb    %xmm2, %edi
	0xc1, 0xe7, 0x10, //0x00000343 shll         $16, %edi
	0x09, 0xf7, //0x00000346 orl          %esi, %edi
	0x09, 0xc7, //0x00000348 orl          %eax, %edi
	0x83, 0xff, 0xff, //0x0000034a cmpl         $-1, %edi
	0x0f, 0x84, 0x28, 0x00, 0x00, 0x00, //0x0000034d je           LBB0_23
	0x4d, 0x01, 0xc3, //0x00000353 addq         %r8, %r11
	0x83, 0xc2, 0xff, //0x00000356 addl         $-1, %edx
	0x0f, 0x85, 0xd1, 0xff, 0xff, 0xff, //0x00000359 jne          LBB0_22
	//0x0000035f LBB0_38
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x0000035f movq         $-1, %rax
	0xe9, 0x1a, 0x00, 0x00, 0x00, //0x00000366 jmp          LBB0_39
	//0x0000036b LBB0_35
	0x4b, 0x8d, 0x34, 0x18, //0x0000036b leaq         (%r8,%r11), %rsi
	0x48, 0x83, 0xc6, 0xff, //0x0000036f addq         $-1, %rsi
	0x0f, 0xb6, 0x06, //0x00000373 movzbl       (%rsi), %eax
	0xe9, 0x0a, 0x00, 0x00, 0x00, //0x00000376 jmp          LBB0_39
	//0x0000037b LBB0_23
	0x49, 0x01, 0xcb, //0x0000037b addq         %rcx, %r11
	0x4c, 0x89, 0xde, //0x0000037e movq         %r11, %rsi
	0x41, 0x0f, 0xb6, 0x03, //0x00000381 movzbl       (%r11), %eax
	//0x00000385 LBB0_39
	0x48, 0x83, 0xc4, 0x10, //0x00000385 addq         $16, %rsp
	0x5b, //0x00000389 popq         %rbx
	0x41, 0x5c, //0x0000038a popq         %r12
	0x41, 0x5d, //0x0000038c popq         %r13
	0x41, 0x5e, //0x0000038e popq         %r14
	0x41, 0x5f, //0x00000390 popq         %r15
	0x5d, //0x00000392 popq         %rbp
	0xc3, //0x00000393 retq         
}
 