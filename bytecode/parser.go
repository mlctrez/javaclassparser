package bytecode

import (
	"io"
)

func buildOpCodeFunctionMap() map[byte]ByteCodeReader {
	m := make(map[byte]ByteCodeReader)
	m[0x00] = r00
	m[0x01] = r01
	m[0x02] = r02
	m[0x03] = r03
	m[0x04] = r04
	m[0x05] = r05
	m[0x06] = r06
	m[0x07] = r07
	m[0x08] = r08
	m[0x09] = r09
	m[0x0A] = r0A
	m[0x0B] = r0B
	m[0x0C] = r0C
	m[0x0D] = r0D
	m[0x0E] = r0E
	m[0x0F] = r0F
	m[0x10] = r10
	m[0x11] = r11
	m[0x12] = r12
	m[0x13] = r13
	m[0x14] = r14
	m[0x15] = r15
	m[0x16] = r16
	m[0x17] = r17
	m[0x18] = r18
	m[0x19] = r19
	m[0x1A] = r1A
	m[0x1B] = r1B
	m[0x1C] = r1C
	m[0x1D] = r1D
	m[0x1E] = r1E
	m[0x1F] = r1F
	m[0x20] = r20
	m[0x21] = r21
	m[0x22] = r22
	m[0x23] = r23
	m[0x24] = r24
	m[0x25] = r25
	m[0x26] = r26
	m[0x27] = r27
	m[0x28] = r28
	m[0x29] = r29
	m[0x2A] = r2A
	m[0x2B] = r2B
	m[0x2C] = r2C
	m[0x2D] = r2D
	m[0x2E] = r2E
	m[0x2F] = r2F
	m[0x30] = r30
	m[0x31] = r31
	m[0x32] = r32
	m[0x33] = r33
	m[0x34] = r34
	m[0x35] = r35
	m[0x36] = r36
	m[0x37] = r37
	m[0x38] = r38
	m[0x39] = r39
	m[0x3A] = r3A
	m[0x3B] = r3B
	m[0x3C] = r3C
	m[0x3D] = r3D
	m[0x3E] = r3E
	m[0x3F] = r3F
	m[0x40] = r40
	m[0x41] = r41
	m[0x42] = r42
	m[0x43] = r43
	m[0x44] = r44
	m[0x45] = r45
	m[0x46] = r46
	m[0x47] = r47
	m[0x48] = r48
	m[0x49] = r49
	m[0x4A] = r4A
	m[0x4B] = r4B
	m[0x4C] = r4C
	m[0x4D] = r4D
	m[0x4E] = r4E
	m[0x4F] = r4F
	m[0x50] = r50
	m[0x51] = r51
	m[0x52] = r52
	m[0x53] = r53
	m[0x54] = r54
	m[0x55] = r55
	m[0x56] = r56
	m[0x57] = r57
	m[0x58] = r58
	m[0x59] = r59
	m[0x5A] = r5A
	m[0x5B] = r5B
	m[0x5C] = r5C
	m[0x5D] = r5D
	m[0x5E] = r5E
	m[0x5F] = r5F
	m[0x60] = r60
	m[0x61] = r61
	m[0x62] = r62
	m[0x63] = r63
	m[0x64] = r64
	m[0x65] = r65
	m[0x66] = r66
	m[0x67] = r67
	m[0x68] = r68
	m[0x69] = r69
	m[0x6A] = r6A
	m[0x6B] = r6B
	m[0x6C] = r6C
	m[0x6D] = r6D
	m[0x6E] = r6E
	m[0x6F] = r6F
	m[0x70] = r70
	m[0x71] = r71
	m[0x72] = r72
	m[0x73] = r73
	m[0x74] = r74
	m[0x75] = r75
	m[0x76] = r76
	m[0x77] = r77
	m[0x78] = r78
	m[0x79] = r79
	m[0x7A] = r7A
	m[0x7B] = r7B
	m[0x7C] = r7C
	m[0x7D] = r7D
	m[0x7E] = r7E
	m[0x7F] = r7F
	m[0x80] = r80
	m[0x81] = r81
	m[0x82] = r82
	m[0x83] = r83
	m[0x84] = r84
	m[0x85] = r85
	m[0x86] = r86
	m[0x87] = r87
	m[0x88] = r88
	m[0x89] = r89
	m[0x8A] = r8A
	m[0x8B] = r8B
	m[0x8C] = r8C
	m[0x8D] = r8D
	m[0x8E] = r8E
	m[0x8F] = r8F
	m[0x90] = r90
	m[0x91] = r91
	m[0x92] = r92
	m[0x93] = r93
	m[0x94] = r94
	m[0x95] = r95
	m[0x96] = r96
	m[0x97] = r97
	m[0x98] = r98
	m[0x99] = r99
	m[0x9A] = r9A
	m[0x9B] = r9B
	m[0x9C] = r9C
	m[0x9D] = r9D
	m[0x9E] = r9E
	m[0x9F] = r9F
	m[0xA0] = rA0
	m[0xA1] = rA1
	m[0xA2] = rA2
	m[0xA3] = rA3
	m[0xA4] = rA4
	m[0xA5] = rA5
	m[0xA6] = rA6
	m[0xA7] = rA7
	m[0xA8] = rA8
	m[0xA9] = rA9
	m[0xAA] = rAA
	m[0xAB] = rAB
	m[0xAC] = rAC
	m[0xAD] = rAD
	m[0xAE] = rAE
	m[0xAF] = rAF
	m[0xB0] = rB0
	m[0xB1] = rB1
	m[0xB2] = rB2
	m[0xB3] = rB3
	m[0xB4] = rB4
	m[0xB5] = rB5
	m[0xB6] = rB6
	m[0xB7] = rB7
	m[0xB8] = rB8
	m[0xB9] = rB9
	m[0xBA] = rBA
	m[0xBB] = rBB
	m[0xBC] = rBC
	m[0xBD] = rBD
	m[0xBE] = rBE
	m[0xBF] = rBF
	m[0xC0] = rC0
	m[0xC1] = rC1
	m[0xC2] = rC2
	m[0xC3] = rC3
	m[0xC4] = rC4
	m[0xC5] = rC5
	m[0xC6] = rC6
	m[0xC7] = rC7
	m[0xC8] = rC8
	m[0xC9] = rC9
	m[0xCA] = rCA
	m[0xFE] = rFE
	m[0xFF] = rFF
	return m
}
func r00(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("nop", p) }
func r01(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aconst_null", p) }
func r02(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_m1", p) }
func r03(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_0", p) }
func r04(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_1", p) }
func r05(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_2", p) }
func r06(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_3", p) }
func r07(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_4", p) }
func r08(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_5", p) }
func r09(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lconst_0", p) }
func r0A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lconst_1", p) }
func r0B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fconst_0", p) }
func r0C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fconst_1", p) }
func r0D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fconst_2", p) }
func r0E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dconst_0", p) }
func r0F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dconst_1", p) }
func r10(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("bipush", p, r, false, 1) }
func r11(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("sipush", p, r, false, 2) }
func r12(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ldc", p, r, false, 1) }
func r13(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ldc_w", p, r, true, 2) }
func r14(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ldc2_w", p, r, true, 2) }
func r15(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("iload", p, r, false, 1) }
func r16(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("lload", p, r, false, 1) }
func r17(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("fload", p, r, false, 1) }
func r18(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("dload", p, r, false, 1) }
func r19(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("aload", p, r, false, 1) }
func r1A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_0", p) }
func r1B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_1", p) }
func r1C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_2", p) }
func r1D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_3", p) }
func r1E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_0", p) }
func r1F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_1", p) }
func r20(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_2", p) }
func r21(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_3", p) }
func r22(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_0", p) }
func r23(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_1", p) }
func r24(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_2", p) }
func r25(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_3", p) }
func r26(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_0", p) }
func r27(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_1", p) }
func r28(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_2", p) }
func r29(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_3", p) }
func r2A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_0", p) }
func r2B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_1", p) }
func r2C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_2", p) }
func r2D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_3", p) }
func r2E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iaload", p) }
func r2F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("laload", p) }
func r30(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("faload", p) }
func r31(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("daload", p) }
func r32(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aaload", p) }
func r33(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("baload", p) }
func r34(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("caload", p) }
func r35(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("saload", p) }
func r36(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("istore", p, r, false, 1) }
func r37(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("lstore", p, r, false, 1) }
func r38(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("fstore", p, r, false, 1) }
func r39(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("dstore", p, r, false, 1) }
func r3A(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("astore", p, r, false, 1) }
func r3B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_0", p) }
func r3C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_1", p) }
func r3D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_2", p) }
func r3E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_3", p) }
func r3F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_0", p) }
func r40(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_1", p) }
func r41(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_2", p) }
func r42(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_3", p) }
func r43(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_0", p) }
func r44(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_1", p) }
func r45(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_2", p) }
func r46(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_3", p) }
func r47(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_0", p) }
func r48(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_1", p) }
func r49(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_2", p) }
func r4A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_3", p) }
func r4B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_0", p) }
func r4C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_1", p) }
func r4D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_2", p) }
func r4E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_3", p) }
func r4F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iastore", p) }
func r50(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lastore", p) }
func r51(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fastore", p) }
func r52(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dastore", p) }
func r53(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aastore", p) }
func r54(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("bastore", p) }
func r55(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("castore", p) }
func r56(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("sastore", p) }
func r57(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("pop", p) }
func r58(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("pop2", p) }
func r59(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup", p) }
func r5A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup_x1", p) }
func r5B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup_x2", p) }
func r5C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup2", p) }
func r5D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup2_x1", p) }
func r5E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup2_x2", p) }
func r5F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("swap", p) }
func r60(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iadd", p) }
func r61(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ladd", p) }
func r62(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fadd", p) }
func r63(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dadd", p) }
func r64(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("isub", p) }
func r65(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lsub", p) }
func r66(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fsub", p) }
func r67(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dsub", p) }
func r68(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("imul", p) }
func r69(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lmul", p) }
func r6A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fmul", p) }
func r6B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dmul", p) }
func r6C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("idiv", p) }
func r6D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ldiv", p) }
func r6E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fdiv", p) }
func r6F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ddiv", p) }
func r70(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("irem", p) }
func r71(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lrem", p) }
func r72(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("frem", p) }
func r73(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("drem", p) }
func r74(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ineg", p) }
func r75(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lneg", p) }
func r76(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fneg", p) }
func r77(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dneg", p) }
func r78(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ishl", p) }
func r79(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lshl", p) }
func r7A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ishr", p) }
func r7B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lshr", p) }
func r7C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iushr", p) }
func r7D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lushr", p) }
func r7E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iand", p) }
func r7F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("land", p) }
func r80(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ior", p) }
func r81(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lor", p) }
func r82(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ixor", p) }
func r83(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lxor", p) }
func r84(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("iinc", p, r, false, 2) }
func r85(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2l", p) }
func r86(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2f", p) }
func r87(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2d", p) }
func r88(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("l2i", p) }
func r89(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("l2f", p) }
func r8A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("l2d", p) }
func r8B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("f2i", p) }
func r8C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("f2l", p) }
func r8D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("f2d", p) }
func r8E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("d2i", p) }
func r8F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("d2l", p) }
func r90(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("d2f", p) }
func r91(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2b", p) }
func r92(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2c", p) }
func r93(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2s", p) }
func r94(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lcmp", p) }
func r95(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fcmpl", p) }
func r96(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fcmpg", p) }
func r97(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dcmpl", p) }
func r98(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dcmpg", p) }
func r99(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifeq", p, r, false, 2) }
func r9A(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifne", p, r, false, 2) }
func r9B(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("iflt", p, r, false, 2) }
func r9C(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifge", p, r, false, 2) }
func r9D(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifgt", p, r, false, 2) }
func r9E(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifle", p, r, false, 2) }
func r9F(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpeq", p, r, false, 2) }
func rA0(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpne", p, r, false, 2) }
func rA1(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmplt", p, r, false, 2) }
func rA2(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpge", p, r, false, 2) }
func rA3(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpgt", p, r, false, 2) }
func rA4(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmple", p, r, false, 2) }
func rA5(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_acmpeq", p, r, false, 2) }
func rA6(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_acmpne", p, r, false, 2) }
func rA7(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("goto", p, r, false, 2) }
func rA8(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("jsr", p, r, false, 2) }
func rA9(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ret", p, r, false, 1) }
func rAA(p *uint32, r io.Reader) (*ByteCode, error) { return TableSwitch("tableswitch", p, r) }
func rAB(p *uint32, r io.Reader) (*ByteCode, error) { return LookupSwitch("lookupswitch", p, r) }
func rAC(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ireturn", p) }
func rAD(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lreturn", p) }
func rAE(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("freturn", p) }
func rAF(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dreturn", p) }
func rB0(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("areturn", p) }
func rB1(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("return", p) }
func rB2(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("getstatic", p, r, true, 2) }
func rB3(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("putstatic", p, r, true, 2) }
func rB4(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("getfield", p, r, true, 2) }
func rB5(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("putfield", p, r, true, 2) }
func rB6(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokevirtual", p, r, true, 2) }
func rB7(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokespecial", p, r, true, 2) }
func rB8(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokestatic", p, r, true, 2) }
func rB9(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokeinterface", p, r, true, 4) }
func rBA(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokedynamic", p, r, true, 4) }
func rBB(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("new", p, r, true, 2) }
func rBC(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("newarray", p, r, false, 1) }
func rBD(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("anewarray", p, r, true, 2) }
func rBE(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("arraylength", p) }
func rBF(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("athrow", p) }
func rC0(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("checkcast", p, r, true, 2) }
func rC1(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("instanceof", p, r, true, 2) }
func rC2(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("monitorenter", p) }
func rC3(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("monitorexit", p) }
func rC4(p *uint32, r io.Reader) (*ByteCode, error) { return Wide("wide", p, r) }
func rC5(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("multianewarray", p, r, true, 3) }
func rC6(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifnull", p, r, false, 2) }
func rC7(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifnonnull", p, r, false, 2) }
func rC8(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("goto_w", p, r, false, 4) }
func rC9(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("jsr_w", p, r, false, 4) }
func rCA(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("breakpoint", p) }
func rFE(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("impdep1", p) }
func rFF(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("impdep2", p) }
