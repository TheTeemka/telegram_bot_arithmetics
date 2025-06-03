package arithmetic

func SolveExpr(str string) (int, error) {
	parser := NewParser(str)
	root, err := parser.Parse()
	if err != nil {
		return 0, err
	}
	return int(root.Solve()), nil
}
