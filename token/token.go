package token

const (
	EOF = iota - 1
	EQ
	PLUS
	MINUS
	NEWLINE
	NUMBER
	IDENT
	STRING
	LABEL
	GOTO
	PRINT
	FN
	ENDFN
	WHILE
	ENDWHILE
	IF
	ELSEIF
	ELSE
	ENDIF
	LET
	NOTEQ
	ASTERISQ
	DOUBLEEQ
	LESSEQ
	ABOVEEQ
	ABOVE
	LESS
	LOGICALAND
	LOGICALOR
	LOGICALXOR
	LOGICALRIGHTBITSHIFT
	LOGICALLEFTBITSHIFT
	COMPAND
	COMPOR
	POWEROP
	INCR
	RPARENT
	LPARENT
	SEPARATOR
	ARROWINITFUNC
	HOOKBEG
	HOOKEND
	BRACEBEG
	BRACEEND
)

var Keyword = map[string]int{
	"let":      LET,
	"while":    WHILE,
	"endwhile": ENDWHILE,
	"if":       IF,
	"endif":    ENDIF,
	"goto":     GOTO,
	"fn":       FN,
	"endfn":    ENDFN,
	"print":    PRINT,
}

type Token_t struct {
	Value   string
	Kind    int
	KindStr string
}

func NewToken(value string, token int, tokenstr string) Token_t {
	return Token_t{
		Value:   value,
		Kind:    token,
		KindStr: tokenstr,
	}
}
