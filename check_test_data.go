package input

func (c check) GoodTestData() (out []string) {
	for k := range c.Data.SourceData() {
		out = append(out, k)
	}
	return
}

func (c check) WrongTestData() (out []string) {

	for _, wd := range wrong_data {
		if _, exist := c.Data.SourceData()[wd]; !exist {
			out = append(out, wd)
		}
	}

	return
}
