// Parameters for the Oakley192/256 Elliptic curve
package oakley

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var oakley192 *elliptic.CurveParams
var oakley256 *elliptic.CurveParams

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

func initOakley256() {
	oakley256 = new(elliptic.CurveParams)
	oakley256.P, _ = new(big.Int).SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF", 16)
	oakley256.N, _ = new(big.Int).SetString("FFFFFFFF00000000FFFFFFFFFFFFFFFFBCE6FAADA7179E84F3B9CAC2FC632551", 16)
	oakley256.B, _ = new(big.Int).SetString("5AC635D8AA3A93E7B3EBBD55769886BC651D06B0CC53B0F63BCE3C3E27D2604B", 16)
	oakley256.Gx, _ = new(big.Int).SetString("6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296", 16)
	oakley256.Gy, _ = new(big.Int).SetString("4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5", 16)
	oakley256.BitSize = 256
}

func Oakley256() elliptic.Curve {
	initonce.Do(initOakley256)
	return oakley256
}
