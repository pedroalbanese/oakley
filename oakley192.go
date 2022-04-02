// Parameters for the Oakley192 Elliptic curve
package oakley192

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var oakley192 *elliptic.CurveParams

func initOakley192() {
	oakley192 = new(elliptic.CurveParams)
	oakley192.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFFFFFFFFFFFF", 16)
	oakley192.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFF99DEF836146BC9B1B4D22831", 16)
	oakley192.B, _ = new(big.Int).SetString("64210519E59C80E70FA7E9AB72243049FEB8DEECC146B9B1", 16)
	oakley192.Gx, _ = new(big.Int).SetString("188DA80EB03090F67CBF20EB43A18800F4FF0AFD82FF1012", 16)
	oakley192.Gy, _ = new(big.Int).SetString("07192B95FFC8DA78631011ED6B24CDD573F977A11E794811", 16)
	oakley192.BitSize = 192
}

func Oakley192() elliptic.Curve {
	initonce.Do(initOakley192)
	return oakley192
}
