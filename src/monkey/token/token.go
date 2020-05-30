package token

// TokenType トークンのタイプを区別
type TokenType string

// Token トークン型の定義
type Token struct {
	Type    TokenType
	Literal string
}

const (

	// ILLEGAL : トークンや文字が未知の意
	ILLEGAL = "ILLEGAL"
	// EOF : ファイル終端の意
	EOF = "EOF"

	// identifier + literal

	// IDENT : add, foobar, x, y
	IDENT = "IDENT"
	// INT : 整数型
	INT = "INT"

	// operator

	// ASSIGN : 代入
	ASSIGN = "="
	// PLUS : 加算演算子
	PLUS = "+"

	// delimiter

	// COMMA : コンマ
	COMMA = ","
	// SEMICOLON : セミコロン
	SEMICOLON = ";"

	// LPAREN : 括弧開き
	LPAREN = "("
	// RPAREN : 括弧閉じ
	RPAREN = ")"
	// LBRACE : 波括弧開き
	LBRACE = "{"
	// RBRACE : 波括弧閉じ
	RBRACE = "}"

	// keyword

	// FUNCTION : 関数
	FUNCTION = "FUNCTION"
	// LET : 変数宣言
	LET = "LET"
)
