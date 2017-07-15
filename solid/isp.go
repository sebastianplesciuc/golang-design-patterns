package main

import (
	"fmt"
)

/*
		The Interface Segregation Principle (ISP)

		The interface-segregation principle (ISP) states that no client should be forced to depend on
	methods it does not use. ISP splits interfaces that are very large into smaller and more specific ones so that
	clients will only have to know about the methods that are of interest to them. Such shrunken interfaces are also
	called role interfaces. ISP is intended to keep a system decoupled and thus easier to refactor, change,
	and redeploy. ISP is one of the five SOLID principles of object-oriented design, similar to
	the High Cohesion Principle of GRASP.


	Ref: https://en.wikipedia.org/wiki/Interface_segregation_principle
*/

type Document string

// !!! Wrong way
type IMachine interface {
	Print([]Document)
	Scan() []Document
	Fax([]Document)
}

type MultiFunctional struct {
}

func (mf *MultiFunctional) Print(docs []Document) {
	fmt.Println("Printing...")
}
func (mf *MultiFunctional) Scan() []Document {
	fmt.Println("Scanning...")
	return []Document{}
}

func (mf *MultiFunctional) Fax(docs []Document) {
	fmt.Println("Faxing...")
}

func NewMachine() IMachine {
	return &MultiFunctional{}
}

// !!! Right way. There are phones who don't have cameras
type IPrinter interface {
	Print([]Document)
}

type IScanner interface {
	Scan() []Document
}

type Printer struct{}

func (p *Printer) Print(docs []Document) {
	fmt.Println("Printing...")
}

type Scanner struct{}

func (s *Scanner) Scan() []Document {
	fmt.Println("Scanning...")
	return []Document{}
}

// Maybe a more fortunate name?
type IScannerAndPrinter interface {
	IPrinter
	IScanner
}

type ScannerAndPrinter struct{}

func (sp *ScannerAndPrinter) Print(docs []Document) {
	fmt.Println("Printing...")
}

func (sp *ScannerAndPrinter) Scan() []Document {
	fmt.Println("Scanning...")
	return []Document{}
}

func NewScannerAndPrinter() IScannerAndPrinter {
	return &ScannerAndPrinter{}
}

func main() {
	fmt.Println("The wrong way...")
	machine := NewMachine()
	machine.Print([]Document{})
	machine.Scan()
	machine.Fax([]Document{})

	fmt.Println()
	fmt.Println("The right way...")

	scannerAndPrinter := NewScannerAndPrinter()
	scannerAndPrinter.Print([]Document{})
	scannerAndPrinter.Scan()
}
