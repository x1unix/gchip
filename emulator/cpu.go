package emulator

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
	opcode uint8

	/*
	CPU registers: The Chip 8 has 15 8-bit general purpose registers named V0,V1 up to VE.
	The 16th register is used  for the ‘carry flag’.
	Eight bits is one byte so we can use an unsigned char for this purpose:
	 */
	v []uint8

	// Index register
	i uint8

	// Program counter which can have a value from 0x000 to 0xFFF
	pc uint8

	// Stack
	stack []uint8

	// Stack pointer
	stackPointer uint8

	/*
	Interupts and hardware registers.
	The Chip 8 has none, but there are two timer registers that count at 60 Hz.
	When set above zero they will count down to zero.
	 */
	delayTimer byte;
	soundTimer byte;
}

func newCpu() *CPU {
	return &CPU {
		v: make([]uint8, 16),
		stack: make([]uint8, 16),
	}
}
