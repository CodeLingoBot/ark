package lexer

type TokenType int

const (
	Rune TokenType = iota
	Identifier
	Separator
	Operator
	Number
	Erroneous
	String
	Doccomment
)

var tokenStrings = []string{"rune", "identifier", "separator", "operator", "number", "erroneous", "string", "doccomment"}

func (v TokenType) String() string {
	return tokenStrings[v]
}

type Token struct {
	Type     TokenType
	Contents string
	Where    Span
}

type Position struct {
	Filename string

	Line, Char int
}

type Span struct {
	Filename string

	StartLine, StartChar int
	EndLine, EndChar     int
}

func NewSpan(start, end Position) Span {
	return Span{Filename: start.Filename,
		StartLine: start.Line, StartChar: start.Char,
		EndLine: end.Line, EndChar: end.Char,
	}
}

func NewSpanFromTokens(start, end *Token) Span {
	return Span{Filename: start.Where.Filename,
		StartLine: start.Where.StartLine, StartChar: start.Where.StartChar,
		EndLine: end.Where.EndLine, EndChar: end.Where.EndChar,
	}
}

func (s Span) Start() Position {
	return Position{Filename: s.Filename,
		Line: s.StartLine, Char: s.StartChar}
}

func (s Span) End() Position {
	return Position{Filename: s.Filename,
		Line: s.EndLine, Char: s.EndChar}
}
