package emulator

const BYTE_OFFSET = 8

/*
Chip 8 CPU

Based on: http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/
*/
type CPU struct {
	/*
		The Chip 8 has 35 opcodes which are all two bytes long.
		To store the current opcode, we need a data type that allows us to store two bytes.
		An unsigned short has the length of two bytes and therefor fits our needs:

		See: http://en.wikipedia.org/wiki/CHIP-8#Opcode_table
	*/
	opcode uint16

	/*
		CPU registers: The Chip 8 has 15 8-bit general purpose registers named V0,V1 up to VE.
		The 16th register is used  for the ‘carry flag’.
		Eight bits is one byte so we can use an unsigned char for this purpose:
	*/
	v []uint8

	// Index register
	i uint16

	// Program counter which can have a value from 0x000 to 0xFFF
	programCounter uint16

	// Stack
	stack []uint8

	// Stack pointer
	stackPointer uint8

	/*
		Interupts and hardware registers.
		The Chip 8 has none, but there are two timer registers that count at 60 Hz.
		When set above zero they will count down to zero.
	*/
	delayTimer byte
	soundTimer byte

	memory *Memory
}

func (c *CPU) next() {
	c.programCounter += 2
}

func (c *CPU) decodeOpcode() {
	pc := c.programCounter

	oX := uint16(c.memory.Get(pc))
	oY := uint16(c.memory.Get(pc + 1))

	// Opcode size is 2 bytes. (1001 0110)
	//
	// Move first byte (1111) to 4 bits left
	// and merge with the second byte to get full opcode
	opcode := uint16(oX<<BYTE_OFFSET | oY)

	switch opcode & 0xf000 {

	}
	c.next()
}

func newCpu(m *Memory) *CPU {
	return &CPU{
		v:              make([]uint8, 16),
		stack:          make([]uint8, 16),
		programCounter: 0x200,
		delayTimer:     0,
		soundTimer:     0,
		memory:         m,
	}
}
