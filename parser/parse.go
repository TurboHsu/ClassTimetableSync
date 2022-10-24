// Package parser parses shit
package parser

import (
	"encoding/xml"
	"regexp"
	"strings"
)

//TODO
/*

 */

func readTable(raw string) {
	//Get rid of all break lines
	raw = strings.ReplaceAll(raw, "\n", "")

	//Get the first table out
	tableRegex := regexp.MustCompile(`(<table).*?(\/table>)`)
	tableRaw := tableRegex.FindString(raw)

	//Parse using xml encoder
	var t XMLTableStruct
	xml.Unmarshal([]byte(tableRaw), &t)

}
