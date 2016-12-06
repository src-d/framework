package utils

type Languages map[string]float64

func (l Languages) Add(lang string, size float64) {
	l[lang] = l[lang] + size
}

func (l Languages) ContainsAnyOf(langs ...string) bool {
	for _, needle := range langs {
		if _, ok := l[needle]; ok {
			return true
		}
	}
	return false
}
