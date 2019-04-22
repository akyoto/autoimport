package autoimport

import (
	"bytes"
	"errors"
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/blitzprog/color"
)

func processFile(path string) error {
	// Read file contents
	file, err := os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		return err
	}

	defer file.Close()
	code, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	// Parse contents
	importPaths, err := parse(code, path)

	if err != nil {
		return err
	}

	if len(importPaths) == 0 {
		return nil
	}

	// Find package definition
	packagePos := bytes.Index(code, []byte("package "))

	if packagePos == -1 {
		return errors.New("Package definition missing")
	}

	seekPos := int64(0)

	for i := packagePos; i < len(code); i++ {
		if code[i] == '\n' {
			seekPos = int64(i + 1)
			break
		}
	}

	// Seek to the beginning (after the package line)
	file.Seek(seekPos, 0)

	// importCommand := fmt.Sprintf("\nimport (\n\t\"%s\"\n)\n\n", strings.Join(importPaths, "\"\n\t\""))
	// file.WriteString(importCommand)
	// file.Write(code[seekPos:])
	// file.Sync()
	fmt.Println(importPaths)

	return nil
}

func parse(src []byte, fileName string) ([]string, error) {
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

				if tok == token.IDENT {
					identifiers = append(identifiers, literal)
				}

				if tok == token.PERIOD {
					lastIdentifier := identifiers[len(identifiers)-1]

					if !isVariable(lastIdentifier) {
						packagesUsed[lastIdentifier] = true
					}
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

	// color.Green(fileName)
	importPaths := []string{}

	for packageName := range packagesUsed {
		packages := packagesByName[packageName]

		if len(packages) == 0 {
			color.Red("Can't find a package import for %s", packageName)
			continue
		}

		// for _, pkg := range packages {
		// 	fmt.Println(" - " + pkg.RealPath)
		// }

		correctPackage := findCorrectPackage(packages)
		importPaths = append(importPaths, correctPackage.ImportPath)
	}

	return importPaths, nil
}
