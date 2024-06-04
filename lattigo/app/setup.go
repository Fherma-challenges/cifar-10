package main

import (
	"app/utils"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

func main() {
	ccFile := flag.String("cc", "", "")
	skFile := flag.String("sk", "", "")
	_ = flag.String("key_public", "", "")
	evalFile := flag.String("key_eval", "", "")
	inputFile := flag.String("input", "", "")

	flag.Parse()

	paramsJSON := struct {
		LogN            int   `json:"log_n"`
		LogQ            []int `json:"log_q"`
		LogP            []int `json:"log_p"`
		LogDefaultScale int   `json:"log_default_scale"`
	}{}

	dataJSON, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("os.Open(%s): %s", "config.json", err.Error())
	}

	if err := json.Unmarshal(dataJSON, &paramsJSON); err != nil {
		log.Fatalf(err.Error())
	}

	var params hefloat.Parameters
	if params, err = hefloat.NewParametersFromLiteral(
		hefloat.ParametersLiteral{
			LogN:            paramsJSON.LogN,
			LogQ:            paramsJSON.LogQ,
			LogP:            paramsJSON.LogP,
			LogDefaultScale: paramsJSON.LogDefaultScale,
		}); err != nil {
		log.Fatalf(err.Error())
	}

	kgen := rlwe.NewKeyGenerator(params)

	sk := kgen.GenSecretKeyNew()

	ecd := hefloat.NewEncoder(params)

	enc := rlwe.NewEncryptor(params, sk)

	rlk := kgen.GenRelinearizationKeyNew(sk)

	evk := rlwe.NewMemEvaluationKeySet(rlk)

	/* #nosec G404 */
	r := rand.New(rand.NewSource(0))

	values := make([]float64, params.MaxSlots())
	for i := range values {
		values[i] = float64(i%256) + (2*r.Float64()-1)*1e-5
	}

	pt := hefloat.NewPlaintext(params, params.MaxLevel())

	if err = ecd.Encode(values, pt); err != nil {
		log.Fatalf(err.Error())
	}

	input, err := enc.EncryptNew(pt)

	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(params, *ccFile); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(sk, *skFile); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(evk, *evalFile); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(input, *inputFile); err != nil {
		log.Fatalf(err.Error())
	}
}
