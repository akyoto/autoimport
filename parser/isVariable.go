package parser

func isVariable(variableStacks [][]string, literal string) bool {
	for _, stack := range variableStacks {
		for _, varName := range stack {
			if varName == literal {
				return true
			}
		}
	}

	return false
}
