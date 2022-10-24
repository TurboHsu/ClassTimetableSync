package parser

type Table struct {
	Subject string
	Time    string
	Address string
}

type XMLTableStruct struct {
	Table struct {
		TR struct {
			td []string `xml:"td"`
		} `xml:"tr"`
	} `xml:"table"`
}
