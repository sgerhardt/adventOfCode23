package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nodeAAA = &node{
	name: "AAA",
}

var nodeBBB = &node{
	name: "BBB",
}

var nodeCCC = &node{
	name: "CCC",
}

var nodeDDD = &node{
	name: "DDD",
}

var nodeEEE = &node{
	name: "EEE",
}
var nodeGGG = &node{
	name: "GGG",
}

var nodeZZZ = &node{
	name: "ZZZ",
}

func Test_findStepsRequired(t *testing.T) {
	type args struct {
		instructions              []rune
		nodes                     []*node
		setupNodes                func()
		expectedSteps             int
		expectedRawStepIterations int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple case, no revisits",
			args: args{
				instructions: []rune{'R', 'L'},
				nodes:        []*node{nodeAAA, nodeBBB, nodeCCC, nodeZZZ},
				setupNodes: func() {
					nodeAAA.leftNode = nodeBBB
					nodeAAA.rightNode = nodeCCC

					nodeBBB.leftNode = nodeDDD
					nodeBBB.rightNode = nodeEEE

					nodeCCC.leftNode = nodeZZZ
					nodeCCC.rightNode = nodeGGG

					nodeDDD.leftNode = nodeDDD
					nodeDDD.rightNode = nodeDDD

					nodeEEE.leftNode = nodeEEE
					nodeEEE.rightNode = nodeEEE

					nodeGGG.leftNode = nodeGGG
					nodeGGG.rightNode = nodeGGG

					nodeZZZ.leftNode = nodeZZZ
					nodeZZZ.rightNode = nodeZZZ
				},
				expectedSteps:             2,
				expectedRawStepIterations: 2,
			},
		},

		{
			name: "revisits",
			args: args{
				instructions: []rune{'L', 'L', 'R'},
				nodes:        []*node{nodeAAA, nodeBBB, nodeCCC, nodeZZZ},
				setupNodes: func() {
					nodeAAA.leftNode = nodeBBB
					nodeAAA.rightNode = nodeBBB

					nodeBBB.leftNode = nodeAAA
					nodeBBB.rightNode = nodeZZZ
				},
				expectedSteps:             6,
				expectedRawStepIterations: 6, //todo is this right?
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.setupNodes()
			stepCount, rawSteps := findStepsRequired(tt.args.instructions, tt.args.nodes)
			assert.Equal(t, tt.args.expectedSteps, stepCount)
			assert.Equal(t, tt.args.expectedRawStepIterations, rawSteps)
		})
	}
}
