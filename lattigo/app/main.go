package main

import (
	"flag"
	"log"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"

	"app/internal/solution"
	"app/utils"
)

func main() {
	ccFile := flag.String("cc", "", "")
	evkFile := flag.String("key_eval", "", "")
	inputFile := flag.String("input", "", "")
	outputFile := flag.String("output", "", "")
	_ = flag.String("key_public", "", "")

	flag.Parse()

	params := new(hefloat.Parameters)
	if err := utils.Deserialize(params, *ccFile); err != nil {
		log.Fatalf(err.Error())
	}

	evk := new(rlwe.MemEvaluationKeySet)
	if err := utils.Deserialize(evk, *evkFile); err != nil {
		log.Fatalf(err.Error())
	}

	in := new(rlwe.Ciphertext)
	if err := utils.Deserialize(in, *inputFile); err != nil {
		log.Fatalf(err.Error())
	}

	out, err := solution.SolveTestcase(params, evk, in)
	if err != nil {
		log.Fatalf("solution.SolveTestcase: %s", err.Error())
	}

	utils.Serialize(out, *outputFile)
}
