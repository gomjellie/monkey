package lexer

type Lexer struct {
	input        string
	position     int  // 입력에서 현재 위치
	readPosition int  // 입력에서 현재 읽는 위치(현재 문자의 다음을 가리킴)
	ch           byte // 현재 조사하고 있는 문자
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
