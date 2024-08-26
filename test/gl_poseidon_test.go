package test

import (
	"fmt"
	pgoldilocks "github.com/OpenAssetStandards/poseidon-goldilocks-go"
	"testing"
)

func TestGlPoseidon(t *testing.T) {
	var inputs []uint64

	//[{7544075957751693767} {11609535206204781183} {6112371860259757379} {12359656023993954792}]
	for i := 0; i < 1296; i++ {
		inputs = append(inputs, 0)
	}

	hashOut, err := pgoldilocks.HashNoPadU64Array(inputs)

	if err != nil {
		fmt.Println(fmt.Errorf("hash fail %v", err))
	} else {
		fmt.Printf("hash out: %d %d %d %d", hashOut[0], hashOut[1], hashOut[2], hashOut[3])
	}

}
