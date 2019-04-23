package parser

import (
	"go/scanner"
	"go/token"
)

// PackageIdentifiers returns all identifiers that are referring to packages.
func PackageIdentifiers(src []byte) map[string]bool {
	packageIdentifiers := map[string]bool{}

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	// Track variables because they look similar to package identifiers.
	variableStacks := [][]string{
		nil, // Top-level variables
	}

	// State
	lastIdentifier := ""
	var lastIdentifiers []string
	inFunctionSignature := false
	inFunctionParameters := false
	var functionParameters []string
	functionName := ""
	parameterName := ""
	parameterType := ""

	// Traverse the input source.
	// Our goal is to find identifiers followed by a period.
	for {
		_, tok, literal := s.Scan()

		if tok == token.EOF {
			break
		}

		switch tok {
		// myVarName
		case token.IDENT:
			lastIdentifier = literal

			// Function signature
			if inFunctionSignature {
				// Function name
				if functionName == "" {
					functionName = literal
					break
				}

				// Variable
				if parameterName == "" {
					parameterName = literal
					functionParameters = append(functionParameters, parameterName)
					break
				}

				// Type
				parameterType += literal
				break
			}

			lastIdentifiers = append(lastIdentifiers, literal)

		// .
		case token.PERIOD:
			lastIdentifiers = nil

			if inFunctionSignature && parameterType != "" {
				parameterType += "."
			}

			if isVariable(variableStacks, lastIdentifier) {
				break
			}

			// color.Yellow("IMPORT %s", lastIdentifier)
			packageIdentifiers[lastIdentifier] = true

		// :=
		case token.DEFINE:
			// color.Magenta("VARS %v", lastIdentifiers)
			stackIndex := len(variableStacks) - 1

			for _, varName := range lastIdentifiers {
				if varName == "_" {
					continue
				}

				variableStacks[stackIndex] = append(variableStacks[stackIndex], varName)
			}

			lastIdentifiers = nil

		// func
		case token.FUNC:
			inFunctionSignature = true

		// (
		case token.LPAREN:
			if inFunctionSignature {
				inFunctionParameters = true
			}

		// )
		case token.RPAREN:
			if inFunctionSignature && inFunctionParameters {
				inFunctionParameters = false
				inFunctionSignature = false
				parameterName = ""
				parameterType = ""
				// color.Cyan(functionName)
				functionName = ""
			}

		// ,
		case token.COMMA:
			if inFunctionSignature {
				parameterName = ""
				parameterType = ""
				break
			}

			lastIdentifiers = append(lastIdentifiers, literal)

		// ;
		case token.SEMICOLON:
			lastIdentifiers = nil

		// {
		case token.LBRACE:
			// Part of type definition
			if inFunctionSignature && parameterType != "" {
				parameterType += literal
				break
			}

			variableStacks = append(variableStacks, functionParameters)
			functionParameters = nil

		// }
		case token.RBRACE:
			// Part of type definition
			if inFunctionSignature && parameterType != "" {
				parameterType += literal
				break
			}

			variableStacks = variableStacks[:len(variableStacks)-1]

		// [
		case token.LBRACK:
			// Part of type definition
			if inFunctionSignature && parameterType != "" {
				parameterType += literal
				break
			}

		// ]
		case token.RBRACK:
			// Part of type definition
			if inFunctionSignature && parameterType != "" {
				parameterType += literal
				break
			}

		// *
		case token.MUL:
			// Part of type definition
			if inFunctionSignature && parameterType != "" {
				parameterType += literal
				break
			}
		}

		// fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, literal)
	}

	return packageIdentifiers
}
