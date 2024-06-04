package solution

import (
	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

func SolveTestcase(params *hefloat.Parameters, evk *rlwe.MemEvaluationKeySet, in *rlwe.Ciphertext) (out *rlwe.Ciphertext, err error) {
	// Put your solution here
	eval := hefloat.NewEvaluator(*params, evk)
	out, err = eval.RotateNew(in, -2)
	return out, err
}
