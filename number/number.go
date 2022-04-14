/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 15:22:14
 */
package number

import (
	"github.com/jhunters/goassist/conv"
	"github.com/jhunters/goassist/generic"
)

type Number interface {
	generic.Integer
	String() string
	BinaryString() string
	OctalString() string
	HexString() string
}

type Integer int

// String
func (i Integer) String() string {
	return i.OctalString()
}

// ToBinaryString returns a string representation of the integer argument as an
// unsigned integer in base 2.
func (i Integer) BinaryString() string {
	return conv.FormatInt(i, conv.Binary)
}

// ToOctalString returns a string representation of the integer argument as an
// unsigned integer in base 10.
func (i Integer) OctalString() string {
	return conv.FormatInt(i, conv.Octal)
}

// ToHexString returns a string representation of the integer argument as an
// unsigned integer in base 16.
func (i Integer) HexString() string {
	return conv.FormatInt(i, conv.Hex)
}
