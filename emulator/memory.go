package emulator

// Chip-8 memory
type Memory struct {
	// The Chip 8 has 4K memory in total, which we can emulated as:
	mem []byte

	/*
		The graphics of the Chip 8 are black and white and the screen has a total of 2048 pixels (64 x 32).
		This can easily be implemented using an array that hold the pixel state (1 or 0):
	*/
	frameBuffer []byte

	/*
		Chip 8 has a HEX based keypad (0x0-0xF), the array used to store the current state of the key.
	*/
	key []byte
}

func (m *Memory) Get(address uint16) byte {
	return m.mem[address]
}

func newMemory() *Memory {
	return &Memory{
		mem:         make([]byte, 4096),
		frameBuffer: make([]byte, 64*32),
		key:         make([]byte, 16),
	}
}
