package main

import "fmt"

/*
		The Dependency Inversion Principle (DIP)

		In object-oriented design, the dependency inversion principle refers to a specific form
	of decoupling software modules. When following this principle, the conventional dependency relationships
	established from high-level, policy-setting modules to low-level, dependency modules are reversed,
	thus rendering high-level modules independent of the low-level module implementation details.

		The principle states:
		- High-level modules should not depend on low-level modules. Both should depend on abstractions.
		- Abstractions should not depend on details. Details should depend on abstractions.

		This design principle inverts the way some people may think about object-oriented programming,
	dictating that both high- and low-level objects must depend on the same abstraction.


	Ref: https://en.wikipedia.org/wiki/Dependency_inversion_principle
	Ref: http://www.oodesign.com/dependency-inversion-principle.html
*/

// !!! Wrong
type RegularWorker struct {
}

func (w *RegularWorker) Work() {
	fmt.Println("Working...")
}

type SpecialWorker struct {
}

func (w *SpecialWorker) Work() {
	fmt.Println("Especially working...")
}

type SpecificManager struct {
	regularWorkers []*RegularWorker
	specialWorkers []*SpecialWorker
}

func (m *SpecificManager) DelegateWork() {
	for _, w := range m.regularWorkers {
		w.Work()
	}

	for _, s := range m.specialWorkers {
		s.Work()
	}
}

func (m *SpecificManager) AddRegularWorker(w *RegularWorker) {
	m.regularWorkers = append(m.regularWorkers, w)
}

func (m *SpecificManager) AddSpecialWorker(w *SpecialWorker) {
	m.specialWorkers = append(m.specialWorkers, w)
}

// !!! The right way
type IWorker interface {
	Work()
}

type Manager struct {
	workers []IWorker
}

func (m *Manager) DelegateWork() {
	for _, w := range m.workers {
		w.Work()
	}
}

func (m *Manager) AddWorker(w IWorker) {
	m.workers = append(m.workers, w)
}

func main() {
	fmt.Println("The wrong way, no abstractions")

	specificManager := &SpecificManager{}
	specificManager.AddRegularWorker(&RegularWorker{})
	specificManager.AddSpecialWorker(&SpecialWorker{})
	specificManager.DelegateWork()

	fmt.Println()
	fmt.Println("The right way")

	manager := &Manager{}
	manager.AddWorker(&RegularWorker{})
	manager.AddWorker(&SpecialWorker{})
	manager.DelegateWork()
}
