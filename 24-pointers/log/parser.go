package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(line string, p *parser, parsed *result) (err error) {
	// Parse the fields
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return err
	}

	parsed.domain = fields[0]

	// Sum the total visits per domain
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
		return err
	}

	return
}

func update(p *parser, parsed *result) {
	domain, visits := parsed.domain, parsed.visits

	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	p.total += visits
	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
}

func printScannerErr(in *bufio.Scanner) {
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}

func printParsedLog(p parser) {
	sort.Strings(p.domains)
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}
