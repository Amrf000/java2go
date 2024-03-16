package main

import (
	"bufio"
	"fmt"
	"java2go/grammar"
	"java2go/parser"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//const pgm = "package main;\npublic class foo\n{\n\tpublic static void main(String[] args) {\n\t\tSystem.out.println(\"hello\");\n\t}\n}\n"

func TestLexer(t *testing.T) {
	//parser.JulyDebug = 9

	//l := grammar.NewLexer(bufio.NewReader(strings.NewReader(pgm)), false)

	//lval := &grammar.JulySymType{}

	//var num int
	//for {
	//	rtn := l.Lex(lval)
	//	if rtn == 0 {
	//		break
	//	}
	//
	//	num++
	//
	//	var tname string
	//	if rtn >= 57346 {
	//		tname = grammar.JulyToknames[rtn-57346]
	//	} else {
	//		tname = fmt.Sprintf("tok#%d (ch '%c')", rtn, byte(rtn))
	//	}
	//
	//	fmt.Printf("#%d: rtn %s\n\t%v\n", num, tname, lval)
	//
	//}
	//target := "test/test/HelloWorld.java"
	//target := "test/test2/Comments.java"
	//target := "test/test3/SwitchCase.java"
	//target := "test/test4/JavaLoop.java"
	//target := "test/test5/Penguin.java"
	//target := "test/test6/Inter.java"
	target := filepath.ToSlash("E:\\dev21\\java2go\\test\\JDK-1.0.2\\src\\java\\lang\\Boolean.java")
	/*destXml := strings.Replace(target, ".java", ".xml", 1)
	destJson := strings.Replace(target, ".java", ".json", 1)
	out, err := exec.Command(
		"java",
		"-jar", "tools/java2go1.jar",
		"-i", target,
		"-o", destXml,
	).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	f1, err := os.Open(destXml)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	// Convert
	js, err := xj.Convert(f1)
	if err != nil {
		panic(err)
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, js.Bytes(), "", "\t")
	if err != nil {
		panic(err)
	}
	os.WriteFile(destJson, js.Bytes(), 0644)

	var cls parser.AstClass
	err = json.Unmarshal(js.Bytes(), &cls)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", cls)*/
	//cls.Root.PackageDeclaration.Name

	//buf, err := io.ReadAll(f1)
	//if err != nil {
	//	panic(err)
	//}
	//var p interface{}
	//if err := xml.Unmarshal(buf, &p); err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%v\n", p)
	f, err := os.Open(target)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	l := grammar.NewLexer(bufio.NewReader(f), false)
	prtn := grammar.JulyParse(l)
	if prtn != 0 {
		fmt.Fprintf(os.Stderr, "parser returned %d\n", prtn)
	}
	if l.JavaProgram() == nil {
		fmt.Println("No code found")
	} else {
		//gp := analyze(l.JavaProgram(), target, &parser.Config{}, parser.StandardRules, true,
		//	true)
		if true {
			fmt.Println(sep + " CONVERT " + sep)
			log.Printf("/** Convert %s **/\n", target)
		}
		gp := parser.NewGoProgram(strings.Replace(target, ".java", ".go", -1), &parser.Config{}, true)
		gp.Analyze(l.JavaProgram())

		for _, rule := range parser.StandardRules {
			gp.RunTransform(rule, gp, nil, nil)
		}

		//fmt.Println(sep + " GODUMP " + sep)
		//gp.WriteString(os.Stdout)
		//fmt.Println()
		//fmt.Println(sep + " PARSE TREE " + sep)
		//gp.DumpTree()
		//fmt.Println(sep + " GO " + sep)
		//gp.Dump(os.Stdout)
		if err := gp.Write(""); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot write %v: %v\n", gp.Name(),
				err)
		}
	}
	l.Close()
}

//func TestBoth(t *testing.T) {
//	//parser.JulyDebug = 9
//
//	l := grammar.NewLexer(bufio.NewReader(strings.NewReader(pgm)), false)
//
//	grammar.JulyParse(l)
//	//if err != nil {
//	//	fmt.Printf("Error in \"%v\": %v\n", input, err)
//	//} else {
//	//	fmt.Printf("\"%v\" -> %v<%T>\n", input, st.(java.Val), st)
//	//}
//	//fmt.Printf("input \"%v\" -> %d\n", pgm, sym)
//}
