// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_lspace = []byte{
	// .p2align 4, 0x90
	// _lspace
	0x55, // pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000001 movq         %rsp, %rbp
	0x48, 0x39, 0xd6, //0x00000004 cmpq         %rdx, %rsi
	0x0f, 0x84, 0x4e, 0x00, 0x00, 0x00, //0x00000007 je           LBB0_1
	0x4c, 0x8d, 0x04, 0x37, //0x0000000d leaq         (%rdi,%rsi), %r8
	0x48, 0x8d, 0x44, 0x3a, 0x01, //0x00000011 leaq         $1(%rdx,%rdi), %rax
	0x48, 0x29, 0xf2, //0x00000016 subq         %rsi, %rdx
	0x48, 0xbe, 0x00, 0x26, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, //0x00000019 movabsq      $4294977024, %rsi
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000023 .p2align 4, 0x90
	//0x00000030 LBB0_3
	0x0f, 0xbe, 0x48, 0xff, //0x00000030 movsbl       $-1(%rax), %ecx
	0x83, 0xf9, 0x20, //0x00000034 cmpl         $32, %ecx
	0x0f, 0x87, 0x2c, 0x00, 0x00, 0x00, //0x00000037 ja           LBB0_5
	0x48, 0x0f, 0xa3, 0xce, //0x0000003d btq          %rcx, %rsi
	0x0f, 0x83, 0x22, 0x00, 0x00, 0x00, //0x00000041 jae          LBB0_5
	0x48, 0xff, 0xc0, //0x00000047 incq         %rax
	0x48, 0xff, 0xc2, //0x0000004a incq         %rdx
	0x0f, 0x85, 0xdd, 0xff, 0xff, 0xff, //0x0000004d jne          LBB0_3
	0x49, 0x29, 0xf8, //0x00000053 subq         %rdi, %r8
	0x4c, 0x89, 0xc0, //0x00000056 movq         %r8, %rax
	0x5d, //0x00000059 popq         %rbp
	0xc3, //0x0000005a retq         
	//0x0000005b LBB0_1
	0x48, 0x01, 0xfa, //0x0000005b addq         %rdi, %rdx
	0x49, 0x89, 0xd0, //0x0000005e movq         %rdx, %r8
	0x49, 0x29, 0xf8, //0x00000061 subq         %rdi, %r8
	0x4c, 0x89, 0xc0, //0x00000064 movq         %r8, %rax
	0x5d, //0x00000067 popq         %rbp
	0xc3, //0x00000068 retq         
	//0x00000069 LBB0_5
	0x48, 0xf7, 0xd7, //0x00000069 notq         %rdi
	0x48, 0x01, 0xf8, //0x0000006c addq         %rdi, %rax
	0x5d, //0x0000006f popq         %rbp
	0xc3, //0x00000070 retq         
}
 
