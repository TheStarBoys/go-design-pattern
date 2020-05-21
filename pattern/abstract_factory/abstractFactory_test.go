package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	engineer := ComputerEngineer{}

	// now in Intel CPU, pins =  1156
	// now in GAMainboard, cpuHoles =  1156
	engineer.MakeComputer(new(Schema1))

	// now in AMD CPU, pins =  939
	// now in MSIMainboard, cpuHoles =  939
	engineer.MakeComputer(new(Schema2))
}