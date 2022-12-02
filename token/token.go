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
)

type Token_t struct {
	Value string
	Kind  int
}

func NewToken(value string, token int) Token_t {
	return Token_t{
		Value: value,
		Kind:  token,
	}
}
