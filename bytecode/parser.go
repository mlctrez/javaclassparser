package bytecode


func buildOpCodeFunctionMap() map[byte]Reader {
	m := make(map[byte]Reader)
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
func r00(c *Context) (*ByteCode, error) { return Simple("nop", c) }
func r01(c *Context) (*ByteCode, error) { return Simple("aconst_null", c) }
func r02(c *Context) (*ByteCode, error) { return Simple("iconst_m1", c) }
func r03(c *Context) (*ByteCode, error) { return Simple("iconst_0", c) }
func r04(c *Context) (*ByteCode, error) { return Simple("iconst_1", c) }
func r05(c *Context) (*ByteCode, error) { return Simple("iconst_2", c) }
func r06(c *Context) (*ByteCode, error) { return Simple("iconst_3", c) }
func r07(c *Context) (*ByteCode, error) { return Simple("iconst_4", c) }
func r08(c *Context) (*ByteCode, error) { return Simple("iconst_5", c) }
func r09(c *Context) (*ByteCode, error) { return Simple("lconst_0", c) }
func r0A(c *Context) (*ByteCode, error) { return Simple("lconst_1", c) }
func r0B(c *Context) (*ByteCode, error) { return Simple("fconst_0", c) }
func r0C(c *Context) (*ByteCode, error) { return Simple("fconst_1", c) }
func r0D(c *Context) (*ByteCode, error) { return Simple("fconst_2", c) }
func r0E(c *Context) (*ByteCode, error) { return Simple("dconst_0", c) }
func r0F(c *Context) (*ByteCode, error) { return Simple("dconst_1", c) }
func r10(c *Context) (*ByteCode, error) { return WithArgs("bipush", c, false, 1) }
func r11(c *Context) (*ByteCode, error) { return WithArgs("sipush", c, false, 2) }
func r12(c *Context) (*ByteCode, error) { return WithArgs("ldc", c, false, 1) }
func r13(c *Context) (*ByteCode, error) { return WithArgs("ldc_w", c, true, 2) }
func r14(c *Context) (*ByteCode, error) { return WithArgs("ldc2_w", c, true, 2) }
func r15(c *Context) (*ByteCode, error) { return WithArgs("iload", c, false, 1) }
func r16(c *Context) (*ByteCode, error) { return WithArgs("lload", c, false, 1) }
func r17(c *Context) (*ByteCode, error) { return WithArgs("fload", c, false, 1) }
func r18(c *Context) (*ByteCode, error) { return WithArgs("dload", c, false, 1) }
func r19(c *Context) (*ByteCode, error) { return WithArgs("aload", c, false, 1) }
func r1A(c *Context) (*ByteCode, error) { return Simple("iload_0", c) }
func r1B(c *Context) (*ByteCode, error) { return Simple("iload_1", c) }
func r1C(c *Context) (*ByteCode, error) { return Simple("iload_2", c) }
func r1D(c *Context) (*ByteCode, error) { return Simple("iload_3", c) }
func r1E(c *Context) (*ByteCode, error) { return Simple("lload_0", c) }
func r1F(c *Context) (*ByteCode, error) { return Simple("lload_1", c) }
func r20(c *Context) (*ByteCode, error) { return Simple("lload_2", c) }
func r21(c *Context) (*ByteCode, error) { return Simple("lload_3", c) }
func r22(c *Context) (*ByteCode, error) { return Simple("fload_0", c) }
func r23(c *Context) (*ByteCode, error) { return Simple("fload_1", c) }
func r24(c *Context) (*ByteCode, error) { return Simple("fload_2", c) }
func r25(c *Context) (*ByteCode, error) { return Simple("fload_3", c) }
func r26(c *Context) (*ByteCode, error) { return Simple("dload_0", c) }
func r27(c *Context) (*ByteCode, error) { return Simple("dload_1", c) }
func r28(c *Context) (*ByteCode, error) { return Simple("dload_2", c) }
func r29(c *Context) (*ByteCode, error) { return Simple("dload_3", c) }
func r2A(c *Context) (*ByteCode, error) { return Simple("aload_0", c) }
func r2B(c *Context) (*ByteCode, error) { return Simple("aload_1", c) }
func r2C(c *Context) (*ByteCode, error) { return Simple("aload_2", c) }
func r2D(c *Context) (*ByteCode, error) { return Simple("aload_3", c) }
func r2E(c *Context) (*ByteCode, error) { return Simple("iaload", c) }
func r2F(c *Context) (*ByteCode, error) { return Simple("laload", c) }
func r30(c *Context) (*ByteCode, error) { return Simple("faload", c) }
func r31(c *Context) (*ByteCode, error) { return Simple("daload", c) }
func r32(c *Context) (*ByteCode, error) { return Simple("aaload", c) }
func r33(c *Context) (*ByteCode, error) { return Simple("baload", c) }
func r34(c *Context) (*ByteCode, error) { return Simple("caload", c) }
func r35(c *Context) (*ByteCode, error) { return Simple("saload", c) }
func r36(c *Context) (*ByteCode, error) { return WithArgs("istore", c, false, 1) }
func r37(c *Context) (*ByteCode, error) { return WithArgs("lstore", c, false, 1) }
func r38(c *Context) (*ByteCode, error) { return WithArgs("fstore", c, false, 1) }
func r39(c *Context) (*ByteCode, error) { return WithArgs("dstore", c, false, 1) }
func r3A(c *Context) (*ByteCode, error) { return WithArgs("astore", c, false, 1) }
func r3B(c *Context) (*ByteCode, error) { return Simple("istore_0", c) }
func r3C(c *Context) (*ByteCode, error) { return Simple("istore_1", c) }
func r3D(c *Context) (*ByteCode, error) { return Simple("istore_2", c) }
func r3E(c *Context) (*ByteCode, error) { return Simple("istore_3", c) }
func r3F(c *Context) (*ByteCode, error) { return Simple("lstore_0", c) }
func r40(c *Context) (*ByteCode, error) { return Simple("lstore_1", c) }
func r41(c *Context) (*ByteCode, error) { return Simple("lstore_2", c) }
func r42(c *Context) (*ByteCode, error) { return Simple("lstore_3", c) }
func r43(c *Context) (*ByteCode, error) { return Simple("fstore_0", c) }
func r44(c *Context) (*ByteCode, error) { return Simple("fstore_1", c) }
func r45(c *Context) (*ByteCode, error) { return Simple("fstore_2", c) }
func r46(c *Context) (*ByteCode, error) { return Simple("fstore_3", c) }
func r47(c *Context) (*ByteCode, error) { return Simple("dstore_0", c) }
func r48(c *Context) (*ByteCode, error) { return Simple("dstore_1", c) }
func r49(c *Context) (*ByteCode, error) { return Simple("dstore_2", c) }
func r4A(c *Context) (*ByteCode, error) { return Simple("dstore_3", c) }
func r4B(c *Context) (*ByteCode, error) { return Simple("astore_0", c) }
func r4C(c *Context) (*ByteCode, error) { return Simple("astore_1", c) }
func r4D(c *Context) (*ByteCode, error) { return Simple("astore_2", c) }
func r4E(c *Context) (*ByteCode, error) { return Simple("astore_3", c) }
func r4F(c *Context) (*ByteCode, error) { return Simple("iastore", c) }
func r50(c *Context) (*ByteCode, error) { return Simple("lastore", c) }
func r51(c *Context) (*ByteCode, error) { return Simple("fastore", c) }
func r52(c *Context) (*ByteCode, error) { return Simple("dastore", c) }
func r53(c *Context) (*ByteCode, error) { return Simple("aastore", c) }
func r54(c *Context) (*ByteCode, error) { return Simple("bastore", c) }
func r55(c *Context) (*ByteCode, error) { return Simple("castore", c) }
func r56(c *Context) (*ByteCode, error) { return Simple("sastore", c) }
func r57(c *Context) (*ByteCode, error) { return Simple("pop", c) }
func r58(c *Context) (*ByteCode, error) { return Simple("pop2", c) }
func r59(c *Context) (*ByteCode, error) { return Simple("dup", c) }
func r5A(c *Context) (*ByteCode, error) { return Simple("dup_x1", c) }
func r5B(c *Context) (*ByteCode, error) { return Simple("dup_x2", c) }
func r5C(c *Context) (*ByteCode, error) { return Simple("dup2", c) }
func r5D(c *Context) (*ByteCode, error) { return Simple("dup2_x1", c) }
func r5E(c *Context) (*ByteCode, error) { return Simple("dup2_x2", c) }
func r5F(c *Context) (*ByteCode, error) { return Simple("swap", c) }
func r60(c *Context) (*ByteCode, error) { return Simple("iadd", c) }
func r61(c *Context) (*ByteCode, error) { return Simple("ladd", c) }
func r62(c *Context) (*ByteCode, error) { return Simple("fadd", c) }
func r63(c *Context) (*ByteCode, error) { return Simple("dadd", c) }
func r64(c *Context) (*ByteCode, error) { return Simple("isub", c) }
func r65(c *Context) (*ByteCode, error) { return Simple("lsub", c) }
func r66(c *Context) (*ByteCode, error) { return Simple("fsub", c) }
func r67(c *Context) (*ByteCode, error) { return Simple("dsub", c) }
func r68(c *Context) (*ByteCode, error) { return Simple("imul", c) }
func r69(c *Context) (*ByteCode, error) { return Simple("lmul", c) }
func r6A(c *Context) (*ByteCode, error) { return Simple("fmul", c) }
func r6B(c *Context) (*ByteCode, error) { return Simple("dmul", c) }
func r6C(c *Context) (*ByteCode, error) { return Simple("idiv", c) }
func r6D(c *Context) (*ByteCode, error) { return Simple("ldiv", c) }
func r6E(c *Context) (*ByteCode, error) { return Simple("fdiv", c) }
func r6F(c *Context) (*ByteCode, error) { return Simple("ddiv", c) }
func r70(c *Context) (*ByteCode, error) { return Simple("irem", c) }
func r71(c *Context) (*ByteCode, error) { return Simple("lrem", c) }
func r72(c *Context) (*ByteCode, error) { return Simple("frem", c) }
func r73(c *Context) (*ByteCode, error) { return Simple("drem", c) }
func r74(c *Context) (*ByteCode, error) { return Simple("ineg", c) }
func r75(c *Context) (*ByteCode, error) { return Simple("lneg", c) }
func r76(c *Context) (*ByteCode, error) { return Simple("fneg", c) }
func r77(c *Context) (*ByteCode, error) { return Simple("dneg", c) }
func r78(c *Context) (*ByteCode, error) { return Simple("ishl", c) }
func r79(c *Context) (*ByteCode, error) { return Simple("lshl", c) }
func r7A(c *Context) (*ByteCode, error) { return Simple("ishr", c) }
func r7B(c *Context) (*ByteCode, error) { return Simple("lshr", c) }
func r7C(c *Context) (*ByteCode, error) { return Simple("iushr", c) }
func r7D(c *Context) (*ByteCode, error) { return Simple("lushr", c) }
func r7E(c *Context) (*ByteCode, error) { return Simple("iand", c) }
func r7F(c *Context) (*ByteCode, error) { return Simple("land", c) }
func r80(c *Context) (*ByteCode, error) { return Simple("ior", c) }
func r81(c *Context) (*ByteCode, error) { return Simple("lor", c) }
func r82(c *Context) (*ByteCode, error) { return Simple("ixor", c) }
func r83(c *Context) (*ByteCode, error) { return Simple("lxor", c) }
func r84(c *Context) (*ByteCode, error) { return WithArgs("iinc", c, false, 2) }
func r85(c *Context) (*ByteCode, error) { return Simple("i2l", c) }
func r86(c *Context) (*ByteCode, error) { return Simple("i2f", c) }
func r87(c *Context) (*ByteCode, error) { return Simple("i2d", c) }
func r88(c *Context) (*ByteCode, error) { return Simple("l2i", c) }
func r89(c *Context) (*ByteCode, error) { return Simple("l2f", c) }
func r8A(c *Context) (*ByteCode, error) { return Simple("l2d", c) }
func r8B(c *Context) (*ByteCode, error) { return Simple("f2i", c) }
func r8C(c *Context) (*ByteCode, error) { return Simple("f2l", c) }
func r8D(c *Context) (*ByteCode, error) { return Simple("f2d", c) }
func r8E(c *Context) (*ByteCode, error) { return Simple("d2i", c) }
func r8F(c *Context) (*ByteCode, error) { return Simple("d2l", c) }
func r90(c *Context) (*ByteCode, error) { return Simple("d2f", c) }
func r91(c *Context) (*ByteCode, error) { return Simple("i2b", c) }
func r92(c *Context) (*ByteCode, error) { return Simple("i2c", c) }
func r93(c *Context) (*ByteCode, error) { return Simple("i2s", c) }
func r94(c *Context) (*ByteCode, error) { return Simple("lcmp", c) }
func r95(c *Context) (*ByteCode, error) { return Simple("fcmpl", c) }
func r96(c *Context) (*ByteCode, error) { return Simple("fcmpg", c) }
func r97(c *Context) (*ByteCode, error) { return Simple("dcmpl", c) }
func r98(c *Context) (*ByteCode, error) { return Simple("dcmpg", c) }
func r99(c *Context) (*ByteCode, error) { return WithArgs("ifeq", c, false, 2) }
func r9A(c *Context) (*ByteCode, error) { return WithArgs("ifne", c, false, 2) }
func r9B(c *Context) (*ByteCode, error) { return WithArgs("iflt", c, false, 2) }
func r9C(c *Context) (*ByteCode, error) { return WithArgs("ifge", c, false, 2) }
func r9D(c *Context) (*ByteCode, error) { return WithArgs("ifgt", c, false, 2) }
func r9E(c *Context) (*ByteCode, error) { return WithArgs("ifle", c, false, 2) }
func r9F(c *Context) (*ByteCode, error) { return WithArgs("if_icmpeq", c, false, 2) }
func rA0(c *Context) (*ByteCode, error) { return WithArgs("if_icmpne", c, false, 2) }
func rA1(c *Context) (*ByteCode, error) { return WithArgs("if_icmplt", c, false, 2) }
func rA2(c *Context) (*ByteCode, error) { return WithArgs("if_icmpge", c, false, 2) }
func rA3(c *Context) (*ByteCode, error) { return WithArgs("if_icmpgt", c, false, 2) }
func rA4(c *Context) (*ByteCode, error) { return WithArgs("if_icmple", c, false, 2) }
func rA5(c *Context) (*ByteCode, error) { return WithArgs("if_acmpeq", c, false, 2) }
func rA6(c *Context) (*ByteCode, error) { return WithArgs("if_acmpne", c, false, 2) }
func rA7(c *Context) (*ByteCode, error) { return WithArgs("goto", c, false, 2) }
func rA8(c *Context) (*ByteCode, error) { return WithArgs("jsr", c, false, 2) }
func rA9(c *Context) (*ByteCode, error) { return WithArgs("ret", c, false, 1) }
func rAA(c *Context) (*ByteCode, error) { return TableSwitch("tableswitch", c.p, c) }
func rAB(c *Context) (*ByteCode, error) { return LookupSwitch("lookupswitch", c.p, c) }
func rAC(c *Context) (*ByteCode, error) { return Simple("ireturn", c) }
func rAD(c *Context) (*ByteCode, error) { return Simple("lreturn", c) }
func rAE(c *Context) (*ByteCode, error) { return Simple("freturn", c) }
func rAF(c *Context) (*ByteCode, error) { return Simple("dreturn", c) }
func rB0(c *Context) (*ByteCode, error) { return Simple("areturn", c) }
func rB1(c *Context) (*ByteCode, error) { return Simple("return", c) }
func rB2(c *Context) (*ByteCode, error) { return WithArgs("getstatic", c, true, 2) }
func rB3(c *Context) (*ByteCode, error) { return WithArgs("putstatic", c, true, 2) }
func rB4(c *Context) (*ByteCode, error) { return WithArgs("getfield", c, true, 2) }
func rB5(c *Context) (*ByteCode, error) { return WithArgs("putfield", c, true, 2) }
func rB6(c *Context) (*ByteCode, error) { return WithArgs("invokevirtual", c, true, 2) }
func rB7(c *Context) (*ByteCode, error) { return WithArgs("invokespecial", c, true, 2) }
func rB8(c *Context) (*ByteCode, error) { return WithArgs("invokestatic", c, true, 2) }
func rB9(c *Context) (*ByteCode, error) { return WithArgs("invokeinterface", c, true, 4) }
func rBA(c *Context) (*ByteCode, error) { return WithArgs("invokedynamic", c, true, 4) }
func rBB(c *Context) (*ByteCode, error) { return WithArgs("new", c, true, 2) }
func rBC(c *Context) (*ByteCode, error) { return WithArgs("newarray", c, false, 1) }
func rBD(c *Context) (*ByteCode, error) { return WithArgs("anewarray", c, true, 2) }
func rBE(c *Context) (*ByteCode, error) { return Simple("arraylength", c) }
func rBF(c *Context) (*ByteCode, error) { return Simple("athrow", c) }
func rC0(c *Context) (*ByteCode, error) { return WithArgs("checkcast", c, true, 2) }
func rC1(c *Context) (*ByteCode, error) { return WithArgs("instanceof", c, true, 2) }
func rC2(c *Context) (*ByteCode, error) { return Simple("monitorenter", c) }
func rC3(c *Context) (*ByteCode, error) { return Simple("monitorexit", c) }
func rC4(c *Context) (*ByteCode, error) { return Wide("wide", c.p, c) }
func rC5(c *Context) (*ByteCode, error) { return WithArgs("multianewarray", c, true, 3) }
func rC6(c *Context) (*ByteCode, error) { return WithArgs("ifnull", c, false, 2) }
func rC7(c *Context) (*ByteCode, error) { return WithArgs("ifnonnull", c, false, 2) }
func rC8(c *Context) (*ByteCode, error) { return WithArgs("goto_w", c, false, 4) }
func rC9(c *Context) (*ByteCode, error) { return WithArgs("jsr_w", c, false, 4) }
func rCA(c *Context) (*ByteCode, error) { return Simple("breakpoint", c) }
func rFE(c *Context) (*ByteCode, error) { return Simple("impdep1", c) }
func rFF(c *Context) (*ByteCode, error) { return Simple("impdep2", c) }
