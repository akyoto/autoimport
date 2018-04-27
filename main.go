package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

func main() {
	start := time.Now()

	flag.Parse()
	files := flag.Args()

	packagesByName := getWorkspacePackages()
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		path := file

		go func() {
			processFile(path)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(len(packagesByName), "packages", time.Since(start))
}

func processFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()
	code, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return parse(code, path)
}

func parse(src []byte, fileName string) error {
	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                            // positions are relative to fset
	file := fset.AddFile(fileName, fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil, scanner.ScanComments)

	tokens := []token.Token{}
	literals := []string{}
	identifiers := []string{}

	variableStacks := [][]string{
		nil, // Top-level variables
	}

	isVariable := func(literal string) bool {
		for _, stack := range variableStacks {
			for _, varName := range stack {
				if varName == literal {
					return true
				}
			}
		}

		return false
	}

	packagesUsed := map[string]bool{}
	var funcParameters []string

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		_, tok, literal := s.Scan()

		if tok == token.EOF {
			break
		}

		tokens = append(tokens, tok)
		literals = append(literals, literal)

		switch tok {
		case token.IDENT:
			identifiers = append(identifiers, literal)
		case token.PERIOD:
			lastIdentifier := identifiers[len(identifiers)-1]

			if !isVariable(lastIdentifier) {
				packagesUsed[lastIdentifier] = true
			}

		case token.DEFINE:
			// fmt.Println(fileName, "DEFINE", lastLiteral)
			index := len(tokens) - 2

			for index >= 0 {
				varName := literals[index]

				if varName != "_" {
					variableStacks[len(variableStacks)-1] = append(variableStacks[len(variableStacks)-1], varName)
				}

				index--

				if tokens[index] != token.COMMA {
					break
				}

				index--
			}

		case token.VAR:
			// fmt.Println(fileName, literal, "VARIABLE")

			_, _, varName := s.Scan()

			if varName != "_" {
				variableStacks[len(variableStacks)-1] = append(variableStacks[len(variableStacks)-1], varName)
			}

		case token.LBRACE:
			variableStacks = append(variableStacks, funcParameters)
			funcParameters = nil

		case token.RBRACE:
			variableStacks = variableStacks[:len(variableStacks)-1]

		case token.FUNC:
			level := 0
			parsedParameterName := false

			for {
				_, tok, literal := s.Scan()

				if tok == token.EOF {
					break
				}

				if tok == token.LPAREN {
					level++
					continue
				} else if tok == token.RPAREN {
					level--

					if level == 0 {
						break
					}

					continue
				}

				if level == 1 {
					if !parsedParameterName && tok.IsLiteral() {
						funcParameters = append(funcParameters, literal)
						parsedParameterName = true
						continue
					}

					if tok == token.COMMA {
						parsedParameterName = false
						continue
					}
				}
			}
		}

		// fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	color.Green(fileName)
	for packageName := range packagesUsed {
		fmt.Println(packageName)
	}

	return nil
}
