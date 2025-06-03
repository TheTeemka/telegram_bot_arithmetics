package arithmetic

type Parser struct {
	*lexer
}

func NewParser(text string) *Parser {
	return &Parser{
		lexer: NewLexer(text),
	}
}

func (p *Parser) Parse() (*Node, error) {
	return p.expr()
}

func (p *Parser) expr() (*Node, error) { // +, -
	node, err := p.term()
	if err != nil {
		return nil, err
	}

	for kind, x := p.nextToken(); kind == C_Plus || kind == C_Minus; kind, x = p.nextToken() {
		p.consumeToken()

		tmp := newNode(kind, x)
		tmp.Left = node
		node = tmp

		node.Right, err = p.term()
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}
func (p *Parser) term() (*Node, error) { // *, /
	node, err := p.factor()
	if err != nil {
		return nil, err
	}

	for kind, x := p.nextToken(); kind == C_Divide || kind == C_Multiply; kind, x = p.nextToken() {
		p.consumeToken()

		tmp := newNode(kind, x)
		tmp.Left = node
		node = tmp

		node.Right, err = p.factor()
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}
func (p *Parser) factor() (*Node, error) {
	kind, x := p.readToken()

	switch kind {
	case C_Minus, C_Plus: //unary, or numbers
		factor, err := p.factor()
		if err != nil {
			return nil, err
		}
		return unaryNode(kind, factor), nil
	case C_Num:
		return newNode(kind, x), nil
	case C_LeftBracket:
		expr, err := p.expr()
		if err != nil {
			return nil, err
		}

		kind, _ = p.readToken()
		if kind != C_RightBracket {
			return nil, NewarError("Expected Right Bracket", p.cursor, p.text)
		}
		return expr, nil
	default:
		return nil, NewarError("Expected Left Bracket or Number", p.cursor, p.text)
	}
}
