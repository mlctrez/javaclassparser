package bytecode

import (
	"io"
)

func buildOpCodeFunctionMap() map[byte]ByteCodeReader {
	m := make(map[byte]ByteCodeReader)
	m[0x00] = readOpCode_00
	m[0x01] = readOpCode_01
	m[0x02] = readOpCode_02
	m[0x03] = readOpCode_03
	m[0x04] = readOpCode_04
	m[0x05] = readOpCode_05
	m[0x06] = readOpCode_06
	m[0x07] = readOpCode_07
	m[0x08] = readOpCode_08
	m[0x09] = readOpCode_09
	m[0x0A] = readOpCode_0A
	m[0x0B] = readOpCode_0B
	m[0x0C] = readOpCode_0C
	m[0x0D] = readOpCode_0D
	m[0x0E] = readOpCode_0E
	m[0x0F] = readOpCode_0F
	m[0x10] = readOpCode_10
	m[0x11] = readOpCode_11
	m[0x12] = readOpCode_12
	m[0x13] = readOpCode_13
	m[0x14] = readOpCode_14
	m[0x15] = readOpCode_15
	m[0x16] = readOpCode_16
	m[0x17] = readOpCode_17
	m[0x18] = readOpCode_18
	m[0x19] = readOpCode_19
	m[0x1A] = readOpCode_1A
	m[0x1B] = readOpCode_1B
	m[0x1C] = readOpCode_1C
	m[0x1D] = readOpCode_1D
	m[0x1E] = readOpCode_1E
	m[0x1F] = readOpCode_1F
	m[0x20] = readOpCode_20
	m[0x21] = readOpCode_21
	m[0x22] = readOpCode_22
	m[0x23] = readOpCode_23
	m[0x24] = readOpCode_24
	m[0x25] = readOpCode_25
	m[0x26] = readOpCode_26
	m[0x27] = readOpCode_27
	m[0x28] = readOpCode_28
	m[0x29] = readOpCode_29
	m[0x2A] = readOpCode_2A
	m[0x2B] = readOpCode_2B
	m[0x2C] = readOpCode_2C
	m[0x2D] = readOpCode_2D
	m[0x2E] = readOpCode_2E
	m[0x2F] = readOpCode_2F
	m[0x30] = readOpCode_30
	m[0x31] = readOpCode_31
	m[0x32] = readOpCode_32
	m[0x33] = readOpCode_33
	m[0x34] = readOpCode_34
	m[0x35] = readOpCode_35
	m[0x36] = readOpCode_36
	m[0x37] = readOpCode_37
	m[0x38] = readOpCode_38
	m[0x39] = readOpCode_39
	m[0x3A] = readOpCode_3A
	m[0x3B] = readOpCode_3B
	m[0x3C] = readOpCode_3C
	m[0x3D] = readOpCode_3D
	m[0x3E] = readOpCode_3E
	m[0x3F] = readOpCode_3F
	m[0x40] = readOpCode_40
	m[0x41] = readOpCode_41
	m[0x42] = readOpCode_42
	m[0x43] = readOpCode_43
	m[0x44] = readOpCode_44
	m[0x45] = readOpCode_45
	m[0x46] = readOpCode_46
	m[0x47] = readOpCode_47
	m[0x48] = readOpCode_48
	m[0x49] = readOpCode_49
	m[0x4A] = readOpCode_4A
	m[0x4B] = readOpCode_4B
	m[0x4C] = readOpCode_4C
	m[0x4D] = readOpCode_4D
	m[0x4E] = readOpCode_4E
	m[0x4F] = readOpCode_4F
	m[0x50] = readOpCode_50
	m[0x51] = readOpCode_51
	m[0x52] = readOpCode_52
	m[0x53] = readOpCode_53
	m[0x54] = readOpCode_54
	m[0x55] = readOpCode_55
	m[0x56] = readOpCode_56
	m[0x57] = readOpCode_57
	m[0x58] = readOpCode_58
	m[0x59] = readOpCode_59
	m[0x5A] = readOpCode_5A
	m[0x5B] = readOpCode_5B
	m[0x5C] = readOpCode_5C
	m[0x5D] = readOpCode_5D
	m[0x5E] = readOpCode_5E
	m[0x5F] = readOpCode_5F
	m[0x60] = readOpCode_60
	m[0x61] = readOpCode_61
	m[0x62] = readOpCode_62
	m[0x63] = readOpCode_63
	m[0x64] = readOpCode_64
	m[0x65] = readOpCode_65
	m[0x66] = readOpCode_66
	m[0x67] = readOpCode_67
	m[0x68] = readOpCode_68
	m[0x69] = readOpCode_69
	m[0x6A] = readOpCode_6A
	m[0x6B] = readOpCode_6B
	m[0x6C] = readOpCode_6C
	m[0x6D] = readOpCode_6D
	m[0x6E] = readOpCode_6E
	m[0x6F] = readOpCode_6F
	m[0x70] = readOpCode_70
	m[0x71] = readOpCode_71
	m[0x72] = readOpCode_72
	m[0x73] = readOpCode_73
	m[0x74] = readOpCode_74
	m[0x75] = readOpCode_75
	m[0x76] = readOpCode_76
	m[0x77] = readOpCode_77
	m[0x78] = readOpCode_78
	m[0x79] = readOpCode_79
	m[0x7A] = readOpCode_7A
	m[0x7B] = readOpCode_7B
	m[0x7C] = readOpCode_7C
	m[0x7D] = readOpCode_7D
	m[0x7E] = readOpCode_7E
	m[0x7F] = readOpCode_7F
	m[0x80] = readOpCode_80
	m[0x81] = readOpCode_81
	m[0x82] = readOpCode_82
	m[0x83] = readOpCode_83
	m[0x84] = readOpCode_84
	m[0x85] = readOpCode_85
	m[0x86] = readOpCode_86
	m[0x87] = readOpCode_87
	m[0x88] = readOpCode_88
	m[0x89] = readOpCode_89
	m[0x8A] = readOpCode_8A
	m[0x8B] = readOpCode_8B
	m[0x8C] = readOpCode_8C
	m[0x8D] = readOpCode_8D
	m[0x8E] = readOpCode_8E
	m[0x8F] = readOpCode_8F
	m[0x90] = readOpCode_90
	m[0x91] = readOpCode_91
	m[0x92] = readOpCode_92
	m[0x93] = readOpCode_93
	m[0x94] = readOpCode_94
	m[0x95] = readOpCode_95
	m[0x96] = readOpCode_96
	m[0x97] = readOpCode_97
	m[0x98] = readOpCode_98
	m[0x99] = readOpCode_99
	m[0x9A] = readOpCode_9A
	m[0x9B] = readOpCode_9B
	m[0x9C] = readOpCode_9C
	m[0x9D] = readOpCode_9D
	m[0x9E] = readOpCode_9E
	m[0x9F] = readOpCode_9F
	m[0xA0] = readOpCode_A0
	m[0xA1] = readOpCode_A1
	m[0xA2] = readOpCode_A2
	m[0xA3] = readOpCode_A3
	m[0xA4] = readOpCode_A4
	m[0xA5] = readOpCode_A5
	m[0xA6] = readOpCode_A6
	m[0xA7] = readOpCode_A7
	m[0xA8] = readOpCode_A8
	m[0xA9] = readOpCode_A9
	m[0xAA] = readOpCode_AA
	m[0xAB] = readOpCode_AB
	m[0xAC] = readOpCode_AC
	m[0xAD] = readOpCode_AD
	m[0xAE] = readOpCode_AE
	m[0xAF] = readOpCode_AF
	m[0xB0] = readOpCode_B0
	m[0xB1] = readOpCode_B1
	m[0xB2] = readOpCode_B2
	m[0xB3] = readOpCode_B3
	m[0xB4] = readOpCode_B4
	m[0xB5] = readOpCode_B5
	m[0xB6] = readOpCode_B6
	m[0xB7] = readOpCode_B7
	m[0xB8] = readOpCode_B8
	m[0xB9] = readOpCode_B9
	m[0xBA] = readOpCode_BA
	m[0xBB] = readOpCode_BB
	m[0xBC] = readOpCode_BC
	m[0xBD] = readOpCode_BD
	m[0xBE] = readOpCode_BE
	m[0xBF] = readOpCode_BF
	m[0xC0] = readOpCode_C0
	m[0xC1] = readOpCode_C1
	m[0xC2] = readOpCode_C2
	m[0xC3] = readOpCode_C3
	m[0xC4] = readOpCode_C4
	m[0xC5] = readOpCode_C5
	m[0xC6] = readOpCode_C6
	m[0xC7] = readOpCode_C7
	m[0xC8] = readOpCode_C8
	m[0xC9] = readOpCode_C9
	m[0xCA] = readOpCode_CA
	m[0xFE] = readOpCode_FE
	m[0xFF] = readOpCode_FF
	return m
}
func readOpCode_00(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("nop", p) }
func readOpCode_01(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aconst_null", p) }
func readOpCode_02(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_m1", p) }
func readOpCode_03(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_0", p) }
func readOpCode_04(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_1", p) }
func readOpCode_05(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_2", p) }
func readOpCode_06(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_3", p) }
func readOpCode_07(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_4", p) }
func readOpCode_08(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iconst_5", p) }
func readOpCode_09(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lconst_0", p) }
func readOpCode_0A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lconst_1", p) }
func readOpCode_0B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fconst_0", p) }
func readOpCode_0C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fconst_1", p) }
func readOpCode_0D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fconst_2", p) }
func readOpCode_0E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dconst_0", p) }
func readOpCode_0F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dconst_1", p) }
func readOpCode_10(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("bipush", p, r, false, 1) }
func readOpCode_11(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("sipush", p, r, false, 2) }
func readOpCode_12(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ldc", p, r, false, 1) }
func readOpCode_13(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ldc_w", p, r, true, 2) }
func readOpCode_14(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ldc2_w", p, r, true, 2) }
func readOpCode_15(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("iload", p, r, false, 1) }
func readOpCode_16(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("lload", p, r, false, 1) }
func readOpCode_17(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("fload", p, r, false, 1) }
func readOpCode_18(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("dload", p, r, false, 1) }
func readOpCode_19(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("aload", p, r, false, 1) }
func readOpCode_1A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_0", p) }
func readOpCode_1B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_1", p) }
func readOpCode_1C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_2", p) }
func readOpCode_1D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iload_3", p) }
func readOpCode_1E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_0", p) }
func readOpCode_1F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_1", p) }
func readOpCode_20(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_2", p) }
func readOpCode_21(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lload_3", p) }
func readOpCode_22(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_0", p) }
func readOpCode_23(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_1", p) }
func readOpCode_24(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_2", p) }
func readOpCode_25(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fload_3", p) }
func readOpCode_26(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_0", p) }
func readOpCode_27(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_1", p) }
func readOpCode_28(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_2", p) }
func readOpCode_29(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dload_3", p) }
func readOpCode_2A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_0", p) }
func readOpCode_2B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_1", p) }
func readOpCode_2C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_2", p) }
func readOpCode_2D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aload_3", p) }
func readOpCode_2E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iaload", p) }
func readOpCode_2F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("laload", p) }
func readOpCode_30(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("faload", p) }
func readOpCode_31(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("daload", p) }
func readOpCode_32(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aaload", p) }
func readOpCode_33(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("baload", p) }
func readOpCode_34(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("caload", p) }
func readOpCode_35(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("saload", p) }
func readOpCode_36(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("istore", p, r, false, 1) }
func readOpCode_37(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("lstore", p, r, false, 1) }
func readOpCode_38(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("fstore", p, r, false, 1) }
func readOpCode_39(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("dstore", p, r, false, 1) }
func readOpCode_3A(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("astore", p, r, false, 1) }
func readOpCode_3B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_0", p) }
func readOpCode_3C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_1", p) }
func readOpCode_3D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_2", p) }
func readOpCode_3E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("istore_3", p) }
func readOpCode_3F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_0", p) }
func readOpCode_40(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_1", p) }
func readOpCode_41(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_2", p) }
func readOpCode_42(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lstore_3", p) }
func readOpCode_43(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_0", p) }
func readOpCode_44(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_1", p) }
func readOpCode_45(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_2", p) }
func readOpCode_46(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fstore_3", p) }
func readOpCode_47(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_0", p) }
func readOpCode_48(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_1", p) }
func readOpCode_49(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_2", p) }
func readOpCode_4A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dstore_3", p) }
func readOpCode_4B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_0", p) }
func readOpCode_4C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_1", p) }
func readOpCode_4D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_2", p) }
func readOpCode_4E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("astore_3", p) }
func readOpCode_4F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iastore", p) }
func readOpCode_50(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lastore", p) }
func readOpCode_51(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fastore", p) }
func readOpCode_52(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dastore", p) }
func readOpCode_53(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("aastore", p) }
func readOpCode_54(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("bastore", p) }
func readOpCode_55(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("castore", p) }
func readOpCode_56(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("sastore", p) }
func readOpCode_57(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("pop", p) }
func readOpCode_58(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("pop2", p) }
func readOpCode_59(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup", p) }
func readOpCode_5A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup_x1", p) }
func readOpCode_5B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup_x2", p) }
func readOpCode_5C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup2", p) }
func readOpCode_5D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup2_x1", p) }
func readOpCode_5E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dup2_x2", p) }
func readOpCode_5F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("swap", p) }
func readOpCode_60(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iadd", p) }
func readOpCode_61(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ladd", p) }
func readOpCode_62(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fadd", p) }
func readOpCode_63(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dadd", p) }
func readOpCode_64(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("isub", p) }
func readOpCode_65(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lsub", p) }
func readOpCode_66(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fsub", p) }
func readOpCode_67(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dsub", p) }
func readOpCode_68(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("imul", p) }
func readOpCode_69(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lmul", p) }
func readOpCode_6A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fmul", p) }
func readOpCode_6B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dmul", p) }
func readOpCode_6C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("idiv", p) }
func readOpCode_6D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ldiv", p) }
func readOpCode_6E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fdiv", p) }
func readOpCode_6F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ddiv", p) }
func readOpCode_70(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("irem", p) }
func readOpCode_71(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lrem", p) }
func readOpCode_72(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("frem", p) }
func readOpCode_73(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("drem", p) }
func readOpCode_74(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ineg", p) }
func readOpCode_75(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lneg", p) }
func readOpCode_76(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fneg", p) }
func readOpCode_77(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dneg", p) }
func readOpCode_78(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ishl", p) }
func readOpCode_79(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lshl", p) }
func readOpCode_7A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ishr", p) }
func readOpCode_7B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lshr", p) }
func readOpCode_7C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iushr", p) }
func readOpCode_7D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lushr", p) }
func readOpCode_7E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("iand", p) }
func readOpCode_7F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("land", p) }
func readOpCode_80(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ior", p) }
func readOpCode_81(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lor", p) }
func readOpCode_82(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ixor", p) }
func readOpCode_83(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lxor", p) }
func readOpCode_84(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("iinc", p, r, false, 2) }
func readOpCode_85(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2l", p) }
func readOpCode_86(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2f", p) }
func readOpCode_87(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2d", p) }
func readOpCode_88(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("l2i", p) }
func readOpCode_89(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("l2f", p) }
func readOpCode_8A(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("l2d", p) }
func readOpCode_8B(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("f2i", p) }
func readOpCode_8C(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("f2l", p) }
func readOpCode_8D(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("f2d", p) }
func readOpCode_8E(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("d2i", p) }
func readOpCode_8F(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("d2l", p) }
func readOpCode_90(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("d2f", p) }
func readOpCode_91(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2b", p) }
func readOpCode_92(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2c", p) }
func readOpCode_93(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("i2s", p) }
func readOpCode_94(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lcmp", p) }
func readOpCode_95(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fcmpl", p) }
func readOpCode_96(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("fcmpg", p) }
func readOpCode_97(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dcmpl", p) }
func readOpCode_98(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dcmpg", p) }
func readOpCode_99(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifeq", p, r, false, 2) }
func readOpCode_9A(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifne", p, r, false, 2) }
func readOpCode_9B(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("iflt", p, r, false, 2) }
func readOpCode_9C(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifge", p, r, false, 2) }
func readOpCode_9D(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifgt", p, r, false, 2) }
func readOpCode_9E(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifle", p, r, false, 2) }
func readOpCode_9F(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpeq", p, r, false, 2) }
func readOpCode_A0(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpne", p, r, false, 2) }
func readOpCode_A1(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmplt", p, r, false, 2) }
func readOpCode_A2(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpge", p, r, false, 2) }
func readOpCode_A3(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmpgt", p, r, false, 2) }
func readOpCode_A4(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_icmple", p, r, false, 2) }
func readOpCode_A5(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_acmpeq", p, r, false, 2) }
func readOpCode_A6(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("if_acmpne", p, r, false, 2) }
func readOpCode_A7(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("goto", p, r, false, 2) }
func readOpCode_A8(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("jsr", p, r, false, 2) }
func readOpCode_A9(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ret", p, r, false, 1) }
func readOpCode_AA(p *uint32, r io.Reader) (*ByteCode, error) { return TableSwitch("tableswitch", p, r) }
func readOpCode_AB(p *uint32, r io.Reader) (*ByteCode, error) { return LookupSwitch("lookupswitch", p, r) }
func readOpCode_AC(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("ireturn", p) }
func readOpCode_AD(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("lreturn", p) }
func readOpCode_AE(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("freturn", p) }
func readOpCode_AF(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("dreturn", p) }
func readOpCode_B0(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("areturn", p) }
func readOpCode_B1(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("return", p) }
func readOpCode_B2(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("getstatic", p, r, true, 2) }
func readOpCode_B3(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("putstatic", p, r, true, 2) }
func readOpCode_B4(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("getfield", p, r, true, 2) }
func readOpCode_B5(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("putfield", p, r, true, 2) }
func readOpCode_B6(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokevirtual", p, r, true, 2) }
func readOpCode_B7(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokespecial", p, r, true, 2) }
func readOpCode_B8(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokestatic", p, r, true, 2) }
func readOpCode_B9(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokeinterface", p, r, true, 4) }
func readOpCode_BA(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("invokedynamic", p, r, true, 4) }
func readOpCode_BB(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("new", p, r, true, 2) }
func readOpCode_BC(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("newarray", p, r, false, 1) }
func readOpCode_BD(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("anewarray", p, r, true, 2) }
func readOpCode_BE(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("arraylength", p) }
func readOpCode_BF(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("athrow", p) }
func readOpCode_C0(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("checkcast", p, r, true, 2) }
func readOpCode_C1(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("instanceof", p, r, true, 2) }
func readOpCode_C2(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("monitorenter", p) }
func readOpCode_C3(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("monitorexit", p) }
func readOpCode_C4(p *uint32, r io.Reader) (*ByteCode, error) { return Wide("wide", p, r) }
func readOpCode_C5(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("multianewarray", p, r, true, 3) }
func readOpCode_C6(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifnull", p, r, false, 2) }
func readOpCode_C7(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("ifnonnull", p, r, false, 2) }
func readOpCode_C8(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("goto_w", p, r, false, 4) }
func readOpCode_C9(p *uint32, r io.Reader) (*ByteCode, error) { return WithArgs("jsr_w", p, r, false, 4) }
func readOpCode_CA(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("breakpoint", p) }
func readOpCode_FE(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("impdep1", p) }
func readOpCode_FF(p *uint32, _ io.Reader) (*ByteCode, error) { return Simple("impdep2", p) }
