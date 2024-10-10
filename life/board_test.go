package life

import (
	"testing"
)

type TestCaseGetNextState struct {
	currentState bool
	score int
	expectedNextState bool
}

func TestGetNextState(t *testing.T) {
	testCases := []TestCaseGetNextState{
		TestCaseGetNextState{true,0,false},
		TestCaseGetNextState{true,1,false},
		TestCaseGetNextState{true,2,true},
		TestCaseGetNextState{true,3,true},
		TestCaseGetNextState{true,4,false},
		TestCaseGetNextState{true,5,false},
		TestCaseGetNextState{true,6,false},
		TestCaseGetNextState{true,7,false},
		TestCaseGetNextState{true,8,false},

		TestCaseGetNextState{false,0,false},
		TestCaseGetNextState{false,1,false},
		TestCaseGetNextState{false,2,false},
		TestCaseGetNextState{false,3,true},
		TestCaseGetNextState{false,4,false},
		TestCaseGetNextState{false,5,false},
		TestCaseGetNextState{false,6,false},
		TestCaseGetNextState{false,7,false},
		TestCaseGetNextState{false,8,false},
	}
	for _,testCase := range testCases {
		nextState := getNextState(testCase.currentState, testCase.score)
		if nextState != testCase.expectedNextState {
			t.Errorf(
				"getNextState(%t, %d) = %t, wanted %t",
				testCase.currentState,
				testCase.score,
				nextState,
				testCase.expectedNextState,
			)
		}
	}
}

type TestCaseAlive struct {
	x int
	y int
	expectedState bool
}

func TestAlive(t *testing.T) {
	board := NewBoard(1,1)
	board.Cells[0][0] = true
	testCases := []TestCaseAlive{
		TestCaseAlive{-1,-1,false},
		TestCaseAlive{0,-1,false},
		TestCaseAlive{1,-1,false},
		TestCaseAlive{-1,0,false},
		TestCaseAlive{0,0,true},
		TestCaseAlive{1,0,false},
		TestCaseAlive{-1,1,false},
		TestCaseAlive{0,1,false},
		TestCaseAlive{1,1,false},
	}
	for _,testCase := range testCases {
		state := board.Alive(testCase.x, testCase.y)
		if state != testCase.expectedState {
			t.Errorf(
				"Alive(%d, %d) = %t, wanted %t",
				testCase.x,
				testCase.y,
				state,
				testCase.expectedState,
			)
		}
	}
}

type TestCell struct {
	x int
	y int
	alive bool
}

type TestCaseNeighborsAlive struct {
	cells []TestCell
	expectedScore int
}

func TestNeighborsAlive(t *testing.T) {
	board := NewBoard(3,3)
	testCases := []TestCaseNeighborsAlive{
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			0,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,true},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,true},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,true},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,true},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,true},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,true},
				TestCell{1,2,false},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,true},
				TestCell{2,2,false},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,false},
				TestCell{1,0,false},
				TestCell{2,0,false},
				TestCell{0,1,false},
				TestCell{1,1,false},
				TestCell{2,1,false},
				TestCell{0,2,false},
				TestCell{1,2,false},
				TestCell{2,2,true},
			},
			1,
		},
		TestCaseNeighborsAlive{
			[]TestCell{
				TestCell{0,0,true},
				TestCell{1,0,true},
				TestCell{2,0,true},
				TestCell{0,1,true},
				TestCell{1,1,true},
				TestCell{2,1,true},
				TestCell{0,2,true},
				TestCell{1,2,true},
				TestCell{2,2,true},
			},
			8,
		},
	}

	for testCaseIndex,testCase := range testCases {
		for _,testCell := range testCase.cells {
			board.Cells[testCell.y][testCell.x] = testCell.alive
		}
		score := board.neighborsAlive(1,1)
		if score != testCase.expectedScore {
			t.Errorf(
				"testCaseIndex: %d neighborAlive(1, 1) = %d, wanted %d",
				testCaseIndex,
				score,
				testCase.expectedScore,
			)
		}
	}
}
