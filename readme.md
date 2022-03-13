# Farify

a simple tools for purify your go code from contamination based on our standard.

## Install

requires [Golang](//golang.org/)
```
go install github.com/akbarfa49/farify@latest
```

### Example for replace Double Quote with Grave Accent
 Example using CLI
```
$ echo 'package test

import (
	"fmt"
	"testing"
)

func TestEsteh(t *testing.T) {
	leaderboardhtml := "edsadhkjasdgjashgfdhasvfdshav"
	leaderboardhtml += "sadsadgsjhabdas" + `"dasdsa` + "dsadsada"
	leaderboardhtml += `<tr style="background: #1C5D79; font-family: 'Calibri'; font-weight: normal; font-size: 14px; line-height: 24px; color: #FFFFFF;">`
	leaderboardhtml += `<td style="text-align: center;">` + leaderboardhtml + `</td>`
	leaderboardhtml += `<td>` + leaderboardhtml + `</td>`
	leaderboardhtml += `<td style="text-align: center;">` + leaderboardhtml + `</td>`
	leaderboardhtml += `</tr>
		`
	fmt.Printf("hello %.2f", 69.123121)
}
' > duar_test.go

$ farify doublequote --file duar_test.go
Done 2 replacement

$ cat duar_test.go

package test

import (
"fmt"
"testing"
)

func TestEsteh(t *testing.T) {
leaderboardhtml := `edsadhkjasdgjashgfdhasvfdshav`
leaderboardhtml += `sadsadgsjhabdas` + `"dasdsa` + `dsadsada`
leaderboardhtml += `<tr style="background: #1C5D79; font-family: Calibri; font-weight: normal; font-size: 14px; line-height: 24px; color: #FFFFFF;">`
leaderboardhtml += `<td style="text-align: center;">` + leaderboardhtml + `</td>`
leaderboardhtml += `<td>` + leaderboardhtml + `</td>`
leaderboardhtml += `<td style="text-align: center;">` + leaderboardhtml + `</td>`
leaderboardhtml += `</tr>
`
fmt.Printf(`hello %.2f`, 69.123121)
}

```

 Example using go:generate
```
//go:generate farify doublequote --file duar.go
```
