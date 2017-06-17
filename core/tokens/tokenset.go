package tokens

// tokenSet API
type TokenSet []byte

func NewTokenSet(rawTS []byte) TokenSet {
	newBytes := make([]byte, len(rawTS))
	copy(newBytes, rawTS)
	return TokenSet(newBytes)
}

func (ts TokenSet) AddAt(index int) TokenSet {
	rawTokens, tokenLen := []byte(ts), len(ts)
	if tokenLen > index {
		rawTokens[index] = 1
	} else {
		rawTokens = append(rawTokens, make([]byte, index-tokenLen+1)...)
		rawTokens[index] = 1
	}
	return TokenSet(rawTokens)
}

func (ts TokenSet) DropUntil(index int) TokenSet {
	rawTokens, tokenLen := []byte(ts), len(ts)
	if tokenLen > index {
		rawTokens = rawTokens[index:]
	} else {
		rawTokens = []byte{}
	}
	return TokenSet(rawTokens)
}

func (ts TokenSet) GetAt(index int) bool {
	if len(ts) > index && ts[index] == 1 {
		return true
	}
	return false
}
