package common

type Token struct {
	Token string
}

func NewToken() *Token {
	return &Token{}
}

func (token *Token) GenerateToken(firstName, lastName string) *Token {
	// dump token
	token.Token = firstName + lastName
	return token
}
