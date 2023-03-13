package main

import (
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
	lerr    error
}

func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

func parse(line string, p *parser, r *result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	// Parse the fields
	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
	}

	r.domain = fields[0]

	// Sum the total visits per domain
	var err error
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}
}

func update(p *parser, r *result) {
	if p.lerr != nil {
		return
	}

	//domain, visits := r.domain, r.visits

	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	p.total += r.visits
	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
}

func printScannerErr(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}

func printParsedLog(p *parser) {
	sort.Strings(p.domains)
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))
	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

func err(p *parser) error {
	return p.lerr
}
