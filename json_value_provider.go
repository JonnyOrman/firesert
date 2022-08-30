package firesert

type JsonValueProvider struct {
	key               string
	configurationJson map[string]interface{}
}

func NewJsonValueProvider(key string) *JsonValueProvider {
	this := new(JsonValueProvider)

	this.key = key

	return this
}

func (this JsonValueProvider) Get() (string, bool) {
	jsonValue := this.configurationJson[this.key]

	var value string

	if jsonValue != nil {
		value = jsonValue.(string)
	}

	return value, len(value) > 0
}
