package dictionary

type Dictionary struct {
	dic map[int64][]string
}

func NewDictionary() *Dictionary {
	return &Dictionary{dic: make(map[int64][]string)}
}

func (d *Dictionary) CreateDictionary(str []string) *map[int64][]string {

	for _, word := range str {
		key, lowerCase, err := d.getKeyForWord(word)
		if err != nil {
			continue
		}

	}
	return &d.dic
}

func (d *Dictionary) getKeyForWord(w string) (int64, string, error) {

	return 0, "", nil
}
