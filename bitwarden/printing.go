package bitwarden

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arphillips06/bw-pwned/helper"
	"github.com/arphillips06/bw-pwned/models"
)

func printSummary(breached, safe int) {
	total := breached + safe
	fmt.Println("================================")
	fmt.Printf("%sScan Summary%s\n", helper.Bold, helper.Reset)
	fmt.Println("================================")
	fmt.Printf("%sSAFE:%s      %s%d%s\n",
		helper.Bold, helper.Reset, helper.Green, safe, helper.Reset)
	fmt.Printf("%sBREACHED:%s  %s%d%s\n",
		helper.Bold, helper.Reset, helper.Red, breached, helper.Reset)
	fmt.Printf("%sTOTAL:%s     %s%d%s\n",
		helper.Bold, helper.Reset, helper.Cyan, total, helper.Reset)
	fmt.Println("===============================")
}

func printTable(results []models.Result, breachedCount, safeCount int) {
	maxURI := len("URI")
	maxUser := len("USERNAME")
	maxBreach := len("BREACHES")
	for _, r := range results {
		if !r.Pwned {
			continue
		}
		if len(r.URI) > maxURI {
			maxURI = len(r.URI)
		}
		if len(r.Username) > maxUser {
			maxUser = len(r.Username)
		}
		b := len(strconv.Itoa(int(r.PwnedCount)))
		if b > maxBreach {
			maxBreach = b
		}
	}
	statusWidth := len("BREACHED") + 2
	uriWidth := maxURI + 2
	userWidth := maxUser + 2
	breachWidth := maxBreach + 2
	header := ""
	header += helper.PadANSI(helper.Bold+"STATUS"+helper.Reset, statusWidth)
	header += helper.PadANSI(helper.Bold+"URI"+helper.Reset, uriWidth)
	header += helper.PadANSI(helper.Bold+"USERNAME"+helper.Reset, userWidth)
	header += helper.PadANSI(helper.Bold+"BREACHES"+helper.Reset, breachWidth)
	fmt.Println(header)
	sepLen := statusWidth + uriWidth + userWidth + breachWidth
	fmt.Println(strings.Repeat("-", sepLen))
	for _, r := range results {
		if !r.Pwned {
			continue
		}
		status := helper.Red + "BREACHED" + helper.Reset
		breaches := helper.Yellow + strconv.Itoa(int(r.PwnedCount)) + helper.Reset
		fmt.Println(
			helper.PadANSI(status, statusWidth) +
				helper.PadANSI(r.URI, uriWidth) +
				helper.PadANSI(r.Username, userWidth) +
				helper.PadANSI(breaches, breachWidth),
		)
	}
	fmt.Println()
	printSummary(breachedCount, safeCount)
}

func printResults(results []models.Result) {
	showPw := helper.PromptYesNo("Show passwords in output? (y/n): ")
	tableView := helper.PromptYesNo("Show table view (y/n): ")
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
		fmt.Printf("%sBreaches:%s  %s%d%s\n",
			helper.Bold, helper.Reset, helper.Yellow, r.PwnedCount, helper.Reset)
		fmt.Println()
	}
	printSummary(breachedCount, safeCount)
}
