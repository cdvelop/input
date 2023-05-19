package input

func (i ip) GoodTestData(table_name, field_name string, random bool) (out []string) {

	out = []string{
		"120.1.3.206",
		"195.145.149.184",
		"179.183.230.16",
		"253.70.9.26",
		"215.35.117.51",
		"212.149.243.253",
		"126.158.214.250",
		"49.122.253.195",
		"53.218.195.25",
		"190.116.115.117",
		"115.186.149.240",
		"163.95.226.221",
	}

	return
}

func (i ip) WrongTestData() (out []string) {
	out = []string{
		"0.0.0.0",
		"192.168.1.1.8",
	}
	out = append(out, wrong_data...)
	return
}
