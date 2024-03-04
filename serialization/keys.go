package serialization

import (
	"flag"
	"fmt"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

var FlagInputPath = flag.String("input", "", "path to the serialized input [rlwe.Ciphertext]")
var FlagOutputPath = flag.String("output", "", "path to the serialized output [rlwe.Ciphertext]")
var FlagCCPath = flag.String("cc", "", "path to the serialized [hefloat.Paramters]")
var FlagKeyEvalPath = flag.String("key_eval", "", "path to the serialized [rlwe.MemEvaluationKeySet]")

func SaveParameters(params *hefloat.Parameters) (err error) {
	if FlagCCPath == nil {
		return fmt.Errorf("invalid path: FlagCCPath is nil")
	}

	return Serialize(params, *FlagCCPath)
}

func LoadParameters(params *hefloat.Parameters) (err error) {

	if FlagCCPath == nil {
		return fmt.Errorf("invalid path: FlagCCPath is nil")
	}

	return Deserialize(params, *FlagCCPath)
}

func SaveMemEvaluationKeySet(evk *rlwe.MemEvaluationKeySet) (err error) {

	if FlagKeyEvalPath == nil {
		return fmt.Errorf("invalid path: FlagKeyEvalPath is nil")
	}

	return Serialize(evk, *FlagKeyEvalPath)
}

func LoadMemEvaluationKeySet(evk *rlwe.MemEvaluationKeySet) (err error) {

	if FlagKeyEvalPath == nil {
		return fmt.Errorf("invalid path: FlagKeyEvalPath is nil")
	}

	return Deserialize(evk, *FlagKeyEvalPath)
}

func SaveInput(input *rlwe.Ciphertext) (err error) {

	if FlagInputPath == nil {
		return fmt.Errorf("invalid path: FlagInputPath is nil")
	}

	return Serialize(input, *FlagInputPath)
}

func LoadInput(input *rlwe.Ciphertext) (err error) {

	if FlagInputPath == nil {
		return fmt.Errorf("invalid path: FlagInputPath is nil")
	}

	return Deserialize(input, *FlagInputPath)
}

func SaveOutput(output *rlwe.Ciphertext) (err error) {

	if FlagOutputPath == nil {
		return fmt.Errorf("invalid path: FlagOutputPath is nil")
	}

	return Serialize(output, *FlagOutputPath)
}

func LoadOutput(output *rlwe.Ciphertext) (err error) {

	if FlagOutputPath == nil {
		return fmt.Errorf("invalid path: FlagOutputPath is nil")
	}

	return Deserialize(output, *FlagOutputPath)
}
