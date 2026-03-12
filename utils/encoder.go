package utils

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lookup map[byte]int

func init() {
	lookup = make(map[byte]int)
	for i := range chars {
		lookup[chars[i]] = i
	}
}

func Base62Encoder(id int64) string {
	if id == 0 {
		return "0"
	}

	encode := []byte{}
	for id > 0 {
		index := id % 62
		encode = append(encode, chars[index])
		id /= 62
	}

	for i, j := 0, len(encode)-1; i < j; i, j = i+1, j-1 {
		encode[i], encode[j] = encode[j], encode[i]
	}

	return string(encode)
}

func Base62Decoder(encode string) int64 {
	id := 0
	for i := range encode {
		id = (id * 62) + lookup[encode[i]]
	}

	return int64(id)
}
