package zkdec

import (
	"testing"

	"github.com/taurusgroup/cmp-ecdsa/pkg/arith"
	"github.com/taurusgroup/cmp-ecdsa/pkg/curve"
	"github.com/taurusgroup/cmp-ecdsa/pkg/zk/zkcommon"
)

func TestDec(t *testing.T) {
	verifier := zkcommon.Pedersen
	prover := zkcommon.ProverPaillierPublic

	y := arith.Sample(arith.L, false)
	x := curve.NewScalarBigInt(y)

	C, rho := prover.Enc(y, nil)

	proof := NewProof(prover, verifier, C, x, y, rho)
	if !proof.Verify(prover, verifier, C, x) {
		t.Error("failed to verify")
	}
}