package tokens

// tokenSet API
type tokenSet []byte

func newTokenSet(rawTS []byte) tokenSet {
	newBytes := make([]byte, len(rawTS))
	copy(newBytes, rawTS)
	return tokenSet(newBytes)
}

func (ts tokenSet) addAt(index int) tokenSet {
	rawTokens, tokenLen := []byte(ts), len(ts)
	if tokenLen > index {
		rawTokens[index] = 1
	} else {
		rawTokens = append(rawTokens, make([]byte, index-tokenLen+1)...)
		rawTokens[index] = 1
	}
	return tokenSet(rawTokens)
}

func (ts tokenSet) dropUntil(index int) tokenSet {
	rawTokens, tokenLen := []byte(ts), len(ts)
	if tokenLen > index {
		rawTokens = rawTokens[index:]
	} else {
		rawTokens = []byte{}
	}
	return tokenSet(rawTokens)
}

func (ts tokenSet) getAt(index int) bool {
	if len(ts) > index && ts[index] == 1 {
		return true
	}
	return false
}
