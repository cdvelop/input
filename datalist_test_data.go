package input

func (d datalist) GoodTestData() (out []string) {
	for k := range d.Data.SourceData() {
		out = append(out, k)
	}
	return
}

func (d datalist) WrongTestData() (out []string) {
	for _, wd := range wrong_data {
		if _, exist := d.Data.SourceData()[wd]; !exist {
			out = append(out, wd)
		}
	}
	return
}
