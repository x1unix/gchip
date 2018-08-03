package emulator

type Machine struct {
	cpu CPU
	mem Memory
}

func (m *Machine) Test() string {
	return "Works"
}

func CreateMachine() *Machine {
	mem := newMemory()
	return &Machine{
		cpu: *newCpu(mem),
		mem: *mem,
	}
}