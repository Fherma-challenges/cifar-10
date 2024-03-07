// Package main is a template encrypted arithmetic with floating point values, with a set of example parameters, key generation, encoding, encryption, decryption and decoding.
package main

import (
	"cifar-10/serialization"
	"flag"
	"fmt"
	"math"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

func main() {

	flag.Parse()

	var err error
	var params hefloat.Parameters

	// Example of 128-bit secure parameters enabling depth-7 circuits.
	// LogN:14, LogQP: 431.
	if params, err = hefloat.NewParametersFromLiteral(
		hefloat.ParametersLiteral{
			LogN:            14,                                    // log2(ring degree)
			LogQ:            []int{55, 45, 45, 45, 45, 45, 45, 45}, // log2(primes Q) (ciphertext modulus)
			LogP:            []int{61},                             // log2(primes P) (auxiliary modulus)
			LogDefaultScale: 45,                                    // log2(scale)
		}); err != nil {
		panic(err)
	}

	// Naive in memory EvaluationKeySet
	if err := serialization.SaveParameters(&params); err != nil {
		panic(err)
	}

	paramsTest := &hefloat.Parameters{}

	if err := serialization.LoadParameters(paramsTest); err != nil {
		panic(err)
	}

	params = *paramsTest

	// Key Generator
	kgen := rlwe.NewKeyGenerator(params)

	// Secret Key
	sk := kgen.GenSecretKeyNew()

	rlk := kgen.GenRelinearizationKeyNew(sk)
	gks := kgen.GenGaloisKeysNew([]uint64{5, 32767}, sk)

	// Naive in memory EvaluationKeySet
	evk := rlwe.NewMemEvaluationKeySet(rlk, gks...)
	if err := serialization.SaveMemEvaluationKeySet(evk); err != nil {
		panic(err)
	}

	evkTest := &rlwe.MemEvaluationKeySet{}

	if err := serialization.LoadMemEvaluationKeySet(evkTest); err != nil {
		panic(err)
	}

	evk = evkTest

	// Encoder
	ecd := hefloat.NewEncoder(params)

	// Encryptor
	enc := rlwe.NewEncryptor(params, sk)

	// Decryptor
	dec := rlwe.NewDecryptor(params, sk)

	// Evaluator

	// Any object implementing [rlwe.EvaluationKeySet] will be accepted
	eval := hefloat.NewEvaluator(params, evk)

	// Vector of plaintext values
	slots := params.MaxSlots()
	values := make([]complex128, slots)

	// Populates the vector of plaintext values on the unit circle
	angle := 2 * 3.141592653589793 / float64(params.NthRoot())
	for i := range values {
		values[i] = complex(math.Cos(angle*float64(i)), math.Sin(angle*float64(i)))
	}

	// Allocates a plaintext at the max level.
	// Default rlwe.MetaData:
	// - IsBatched = true (slots encoding)
	// - Scale = params.DefaultScale()
	pt := hefloat.NewPlaintext(params, params.MaxLevel())

	// Encodes the vector of plaintext values
	if err = ecd.Encode(values, pt); err != nil {
		panic(err)
	}

	// Encrypts the vector of plaintext values
	var ct *rlwe.Ciphertext
	if ct, err = enc.EncryptNew(pt); err != nil {
		panic(err)
	}

	if err := serialization.SaveInput(ct); err != nil {
		panic(err)
	}

	ctTest := &rlwe.Ciphertext{}

	if err := serialization.LoadInput(ctTest); err != nil {
		panic(err)
	}

	ct = ctTest

	// Dummy encrypted circuit
	if err = eval.MulRelin(ct, ct, ct); err != nil {
		panic(err)
	}

	if err = eval.Rotate(ct, 1, ct); err != nil {
		panic(err)
	}

	if err = eval.Conjugate(ct, ct); err != nil {
		panic(err)
	}

	if err = eval.Rescale(ct, ct); err != nil {
		panic(err)
	}

	if err := serialization.SaveOutput(ct); err != nil {
		panic(err)
	}

	ctTest = &rlwe.Ciphertext{}

	if err := serialization.LoadOutput(ctTest); err != nil {
		panic(err)
	}

	ct = ctTest

	// Dummy plaintext circuit
	want := make([]complex128, params.MaxSlots())
	copy(want, values)
	for i := range want {
		x := values[(i+1)&(slots-1)]
		y := x * x
		want[i] = complex(real(y), -imag(y))
	}

	PrintPrecisionStats(params, ct, want, ecd, dec)

	serialization.CleanFolder("./data/")
}

// PrintPrecisionStats decrypts, decodes and prints the precision stats of a ciphertext.
func PrintPrecisionStats(params hefloat.Parameters, ct *rlwe.Ciphertext, want []complex128, ecd *hefloat.Encoder, dec *rlwe.Decryptor) {

	var err error

	// Decrypts the vector of plaintext values
	pt := dec.DecryptNew(ct)

	// Decodes the plaintext
	have := make([]complex128, params.MaxSlots())
	if err = ecd.Decode(pt, have); err != nil {
		panic(err)
	}

	// Pretty prints some values
	fmt.Printf("Have: ")
	for i := 0; i < 4; i++ {
		fmt.Printf("%17.15f ", have[i])
	}
	fmt.Printf("...\n")

	fmt.Printf("Want: ")
	for i := 0; i < 4; i++ {
		fmt.Printf("%17.15f ", want[i])
	}
	fmt.Printf("...\n")

	// Pretty prints the precision stats
	fmt.Println(hefloat.GetPrecisionStats(params, ecd, dec, have, want, 0, false).String())
}
