package doublequote

import (
	"fmt"
	"io"
	"log"
	"os"
)

func PurifyDq(f *os.File) {
	importSkipped := false

	byt, err := io.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}
	openingDq := 0
	gaFound := false
	openingGa := 0
	parenthesis := false
	// isDQ := false
	// `"`
	// "`"
	skip := false
	replacement := 0
	for idx, v := range byt {
		if skip {
			if v == '\n' {
				skip = false
			}
			continue
		}
		if openingDq == 0 && openingGa == 0 && v == '/' {
			if byt[idx+1] == '/' {
				skip = true
			}
		}

		if !importSkipped {
			switch {
			case !parenthesis:
				if v == '(' {
					openingDq = 1
				} else if v == ')' {
					openingDq = 0
					importSkipped = true
				}
			case parenthesis && v == '"':
				importSkipped = true
			case openingDq == 0 && v == '"':
				parenthesis = true
			}
			continue
		}

		switch {
		case openingDq == 0 && v == '`':
			if openingGa > 0 {
				openingGa = 0
			} else {
				openingGa = idx
			}

		case openingGa == 0 && v == '"':
			if openingDq > 0 {
				if !gaFound {
					byt[openingDq] = '`'
					byt[idx] = '`'
					replacement += 2
				}
				gaFound = false
				openingDq = 0

			} else {
				openingDq = idx
			}
		case openingDq > 0 && openingGa == 0:
			if v == '`' || v == byte(92) {
				gaFound = true
			}
		}
	}
	_, err = f.WriteAt(byt, 0)
	if err == nil {
		fmt.Println(`done `, replacement, ` replacement`)
	}
}
