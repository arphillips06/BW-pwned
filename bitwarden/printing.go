package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/models"
	"fmt"
	"strings"
)

func printSummary(breached, safe int) {
	total := breached + safe

	fmt.Println("================================")
	fmt.Printf("%sScan Summary%s\n", helper.Bold, helper.Reset)
	fmt.Println("================================")
	fmt.Printf("%sSAFE:%s      %s%d%s\n", helper.Bold, helper.Reset, helper.Green, safe, helper.Reset)
	fmt.Printf("%sBREACHED:%s  %s%d%s\n", helper.Bold, helper.Reset, helper.Red, breached, helper.Reset)
	fmt.Printf("%sTOTAL:%s     %s%d%s\n", helper.Bold, helper.Reset, helper.Cyan, total, helper.Reset)
	fmt.Println("===============================")
}

func printTable(results []models.Result, breachedCount, safeCount int) {
	maxURI := 0
	maxUser := 0
	maxBreaches := 0

	for _, r := range results {
		if !r.Pwned {
			continue
		}
		maxURI = max(maxURI, len(r.URI))
		maxUser = max(maxUser, len(r.Username))
		maxBreaches = max(maxBreaches, len(fmt.Sprintf("%d", r.PwnedCount)))
	}

	statusWidth := max(len("STATUS"), len("BREACHED")) + 2
	uriWidth := max(len("URI"), maxURI) + 2
	userWidth := max(len("USERNAME"), maxUser) + 2
	breachWidth := max(len("BREACHES"), maxBreaches) + 2

	header :=
		helper.PadANSI(helper.Bold+"STATUS"+helper.Reset, statusWidth) +
			helper.PadANSI(helper.Bold+"URI"+helper.Reset, uriWidth) +
			helper.PadANSI(helper.Bold+"USERNAME"+helper.Reset, userWidth) +
			helper.PadANSI(helper.Bold+"BREACHES"+helper.Reset, breachWidth)

	fmt.Println(header)

	sepLen := statusWidth + uriWidth + userWidth + breachWidth
	fmt.Println(strings.Repeat("-", sepLen))

	for _, r := range results {
		if !r.Pwned {
			continue
		}

		status := helper.Red + "BREACHED" + helper.Reset
		breaches := helper.Yellow + fmt.Sprintf("%d", r.PwnedCount) + helper.Reset

		fmt.Print(
			helper.PadANSI(status, statusWidth) +
				helper.PadANSI(r.URI, uriWidth) +
				helper.PadANSI(r.Username, userWidth) +
				helper.PadANSI(breaches, breachWidth),
		)
		fmt.Println()
	}

	fmt.Println()
	printSummary(breachedCount, safeCount)
}

func printResults(results []models.Result) {
	fmt.Println("Show passwords in output? (y/n): ")
	var choice string
	fmt.Scanln(&choice)
	showPw := strings.ToLower(choice) == "y"
	fmt.Print("Show table view (y/n): ")
	var table string
	fmt.Scanln(&table)
	tableView := strings.ToLower(table) == "y"
	breachedCount := 0
	safeCount := 0
	for _, r := range results {
		if r.Pwned {
			breachedCount++
		} else {
			safeCount++
		}
	}
	if tableView {
		printTable(results, breachedCount, safeCount)
		return
	}
	for _, r := range results {
		if !r.Pwned {
			continue
		}
		pw := r.Password
		if !showPw {
			pw = "********"
		}
		fmt.Printf("%sBREACHED%s\n", helper.Red, helper.Reset)
		fmt.Printf("%sAccount:%s   %s\n", helper.Bold, helper.Reset, r.URI)
		fmt.Printf("%sUsername:%s  %s\n", helper.Bold, helper.Reset, r.Username)
		fmt.Printf("%sPassword:%s  %s\n", helper.Bold, helper.Reset, pw)
		fmt.Printf("%sBreaches:%s  %s%d%s\n", helper.Bold, helper.Reset, helper.Yellow, r.PwnedCount, helper.Reset)
		fmt.Println()
	}

	printSummary(breachedCount, safeCount)
}
