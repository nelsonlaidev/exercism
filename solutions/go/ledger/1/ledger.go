package ledger

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if !validLocale(locale) {
		return "", errors.New("unsupported locale")
	}
	if !validCurrency(currency) {
		return "", errors.New("unsupported currency")
	}

	entriesCopy := append([]Entry(nil), entries...)
	sort.SliceStable(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	var b strings.Builder
	b.WriteString(header(locale))
	for _, e := range entriesCopy {
		date, err := formatDate(e.Date, locale)
		if err != nil {
			return "", err
		}
		amount, err := formatAmount(e.Change, currency, locale)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "%-10s | %-25s | %13s\n", date, formatDescription(e.Description), amount)
	}
	return b.String(), nil
}

func validLocale(locale string) bool {
	return locale == "en-US" || locale == "nl-NL"
}

func validCurrency(currency string) bool {
	return currency == "USD" || currency == "EUR"
}

func header(locale string) string {
	dateCol, descCol, changeCol := "Date", "Description", "Change"
	if locale == "nl-NL" {
		dateCol, descCol, changeCol = "Datum", "Omschrijving", "Verandering"
	}
	return fmt.Sprintf("%-10s | %-25s | %-13s\n", dateCol, descCol, changeCol)
}

func formatDate(date, locale string) (string, error) {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return "", errors.New("invalid date")
	}
	year, month, day := date[0:4], date[5:7], date[8:10]
	if month < "01" || month > "12" || day < "01" || day > "31" {
		return "", errors.New("invalid date")
	}
	if locale == "nl-NL" {
		return fmt.Sprintf("%s-%s-%s", day, month, year), nil
	}
	return fmt.Sprintf("%s/%s/%s", month, day, year), nil
}

func formatDescription(desc string) string {
	const width = 25
	runes := []rune(desc)
	if len(runes) > width {
		return string(runes[:22]) + "..."
	}
	return desc + strings.Repeat(" ", width-len(runes))
}

func formatAmount(change int, currency, locale string) (string, error) {
	symbol, ok := map[string]string{"USD": "$", "EUR": "€"}[currency]
	if !ok {
		return "", errors.New("unsupported currency")
	}

	negative := change < 0
	if negative {
		change = -change
	}

	thousandsSep, decimalSep := ",", "."
	switch locale {
	case "nl-NL":
		thousandsSep, decimalSep = ".", ","
	}

	amount := groupDigits(strconv.Itoa(change/100), thousandsSep) +
		decimalSep + fmt.Sprintf("%02d", change%100)

	if locale == "nl-NL" {
		if negative {
			return symbol + " -" + amount + " ", nil
		}
		return symbol + " " + amount + " ", nil
	}
	if negative {
		return "(" + symbol + amount + ")", nil
	}
	return symbol + amount + " ", nil
}

func groupDigits(s, sep string) string {
	var groups []string
	for len(s) > 3 {
		groups = append(groups, s[len(s)-3:])
		s = s[:len(s)-3]
	}
	groups = append(groups, s)
	slices.Reverse(groups)
	return strings.Join(groups, sep)
}
