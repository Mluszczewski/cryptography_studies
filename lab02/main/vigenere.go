package main

type vkey string

func generateKey(key string) vkey {
	v := vkey(upperOnly(key))

	return v
}

func (k vkey) encryptCipher(pt string) string {
	ct := upperOnly(pt)
	for i, c := range ct {
		ct[i] = 'A' + (c-'A'+k[i%len(k)]-'A')%26
	}
	return string(ct)
}

func (k vkey) decryptCipher(ct string) string {
	pt := make([]byte, len(ct))
	for i := range pt {
		c := ct[i]
		if c < 'A' || c > 'Z' {
			return ""
		}
		pt[i] = 'A' + (c-k[i%len(k)]+26)%26
	}
	return string(pt)
}

func upperOnly(s string) []byte {
	u := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			u = append(u, c)
		} else if c >= 'a' && c <= 'z' {
			u = append(u, c-32)
		}
	}
	return u
}
