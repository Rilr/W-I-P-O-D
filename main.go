package main

import (
	"fmt"

	"github.com/Rilr/W-I-P-O-D/manage"
)

func main() {
	memberID := "171" // Statically set for now
	timeSheets := manage.GetTimeSheets(memberID)

	for _, sheet := range timeSheets {
		fmt.Printf("ID: %d, Member: %s, Year: %d, Period: %d, Status: %s\n",
			sheet.ID, sheet.Member.Name, sheet.Year, sheet.Period, sheet.Status)
	}
}