package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(input string) (results [][]string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		parts := strings.Split(l, "\t")
		if len(parts) != 5 {
			// incomplete line
			continue
		}
		if "Mnemonic" == parts[0] {
			// header
			continue
		}
		results = append(results, parts)
	}
	return results
}

func main() {

	var input string
	var output string
	flag.StringVar(&input, "input", "bytecode/bytecodes.tsv", "input tsv file")
	flag.StringVar(&output, "output", "bytecode/parser.go", "output go file")
	flag.Parse()

	out, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fmt.Fprint(out, fileHeader)
	fmt.Fprintln(out)

	byteCodeLines := readInput(input)

	sortSlice := func(i, j int) bool {
		return byteCodeLines[i][1] < byteCodeLines[j][1]
	}
	sort.SliceStable(byteCodeLines, sortSlice)

	fmt.Fprintln(out, "func buildOpCodeFunctionMap() map[byte]Reader {")
	fmt.Fprintln(out, "\tm := make(map[byte]Reader)")
	for _, parts := range byteCodeLines {
		byteCodeHex := strings.ToUpper(parts[1])
		fmt.Fprintf(out, "\tm[0x%s] = r%s\n", byteCodeHex, byteCodeHex)
	}
	fmt.Fprintln(out, "\treturn m")
	fmt.Fprintln(out, "}")

	for _, parts := range byteCodeLines {

		args, index := determineOpCodeType(parts[2])

		upperHex := strings.ToUpper(parts[1])

		if args == 0 {
			fmt.Fprintf(out, "func r%s(c *Context) (*ByteCode, error) { ", upperHex)
			fmt.Fprintf(out, "return Simple(%q, c) }\n", parts[0])
			continue
		}

		if args > 0 {
			fmt.Fprintf(out, "func r%s(c *Context) (*ByteCode, error) { ", upperHex)
			fmt.Fprintf(out, "return WithArgs(%q, c, %t, %d) }\n", parts[0], index, args)
			continue
		}

		fmt.Fprintf(out, "func r%s(c *Context) (*ByteCode, error) { ", upperHex)

		switch parts[0]{
		case "tableswitch":
			fmt.Fprintf(out, "return TableSwitch(%q, c.p, c) ", "tableswitch")
		case "lookupswitch":
			fmt.Fprintf(out, "return LookupSwitch(%q, c.p, c) ", "lookupswitch")
		case "wide":
			fmt.Fprintf(out, "return Wide(%q, c.p, c) ", "wide")
		default:
			panic("unhandled opcode " + parts[0])
		}

		fmt.Fprintf(out, "}\n")
	}

}

// parse out Other bytes [count]: [operand labels]
func determineOpCodeType(args string) (otherBytes int, index bool) {
	if args != "" {
		ob := strings.SplitN(args, ":", 2)
		if strings.Contains(ob[0], "/") || strings.HasSuffix(ob[0], "+") {
			// tableswitch, lookupswitch, wide
			return -1, false
		}
		i, err := strconv.Atoi(ob[0])
		if err != nil {
			panic(err)
		}
		return i, strings.HasPrefix(strings.TrimSpace(ob[1]), "indexbyte1")
	}

	return 0, false
}

var fileHeader = `package bytecode

`
