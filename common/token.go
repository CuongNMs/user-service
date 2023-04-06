package common

type Token struct {
	Token string
}

func NewToken() *Token {
	return &Token{}
}

func (token *Token) GenerateToken(firstName, lastName string) *Token {
	// fake token
	token.Token = firstName + lastName
	return token
}
