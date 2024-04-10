package main

import(
  "fmt"
  "testing"
  "runtime"
)

//Tests GameboardStable function
func TestGameboardStable(t *testing.T){

  type test struct {
		board Gameboard
		answer bool
	}
  tests := make([]test, 3)
  tests[0].board = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[0].answer = true
  tests[1].board = Gameboard{{0, 3, 0}, {4, 2, 2}}
  tests[1].answer = false
  tests[2].board = Gameboard{{0, 0, 3, 9}, {4, 1, 5, 2}, {6, 7, 3, 1}, {2, 3, 0, 1}}
  tests[2].answer = false

  for i, test := range tests {
		outcome := GameboardStable(test.board)

		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the board is", test.board, "the answer is", test.answer)
		}
	}
}

//Tests InField function
func TestInField(t *testing.T){

  type test struct {
		a,b,c int
		answer bool
	}
  tests := make([]test, 3)
  tests[0].a = 3
  tests[0].b = 0
  tests[0].c = 4
  tests[0].answer = true
  tests[1].a = 5
  tests[1].b = 2
  tests[1].c = 5
  tests[1].answer = false
  tests[2].a = 6
  tests[2].b = 7
  tests[2].c = 10
  tests[2].answer = false

  for i, test := range tests {
		outcome := InField(test.a, test.b, test.c)

		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the index is", test.a, "and the min size is", test.b, "and the max size is", test.c, "the answer is", test.answer)
		}
	}
}

//Tests GameboardsEqual function
func TestGameboardsEqual(t *testing.T){

  type test struct {
		board1, board2 Gameboard
		answer bool
	}
  tests := make([]test, 3)
  tests[0].board1 = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[0].board2 = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[0].answer = true
  tests[1].board1 = Gameboard{{0, 2, 0}, {2, 3, 2}, {0, 2, 0}}
  tests[1].board2 = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[1].answer = false
  tests[2].board1 = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[2].board2 = Gameboard{{0, 2, 0}, {2, 2, 2}}
  tests[2].answer = false

  for i, test := range tests {
		outcome := GameboardsEqual(test.board1, test.board2)

		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the boards are", test.board1, "and", test.board2, "the answer is", test.answer)
		}
	}
}

//Tests SandpileMultiprocs function
func TestSandpileMultiprocs(t *testing.T){

  type test struct {
		board Gameboard
		answer Gameboard
	}
  tests := make([]test, 3)
  tests[0].board = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[0].answer = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[1].board = Gameboard{{0, 2, 0}, {3, 4, 0}, {1, 9, 5}}
  tests[1].answer = Gameboard{{0, 3, 0}, {4, 1, 2}, {2, 7, 2}}
  tests[2].board = Gameboard{{3, 5}, {2, 8}}
  tests[2].answer = Gameboard{{4, 2}, {3, 5}}

  numProcs := runtime.NumCPU()

  for i, test := range tests {
		outcome := SandpileMultiprocs(test.board, numProcs)

		if !GameboardsEqual(outcome, test.answer) {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the board is", test.board, "the next board is", test.answer)
		}
	}
}

//Tests UpdateGameboard function
func TestUpdateGameboard(t *testing.T){

  type test struct {
		board Gameboard
		answer Gameboard
	}
  tests := make([]test, 3)
  tests[0].board = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[0].answer = Gameboard{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
  tests[1].board = Gameboard{{0, 2, 0}, {3, 4, 0}, {1, 9, 5}}
  tests[1].answer = Gameboard{{0, 3, 0}, {4, 1, 2}, {2, 7, 2}}
  tests[2].board = Gameboard{{3, 5}, {2, 8}}
  tests[2].answer = Gameboard{{4, 2}, {3, 5}}

  for i, test := range tests {
		outcome := UpdateGameboard(test.board)

		if !GameboardsEqual(outcome, test.answer) {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the board is", test.board, "the next board is", test.answer)
		}
	}
}

//Tests Topple function
func TestTopple(t *testing.T){

  type test struct {
		row, col, size, startIndex int
		answer []position
	}

  tests := make([]test, 3)
  tests[0].row = 0
  tests[0].col = 0
  tests[0].size = 5
  tests[0].startIndex = 0
  tests[0].answer = []position{[3]int{0,0,-4}, [3]int{1,0,1}, [3]int{0,1,1}}
  tests[1].row = 1
  tests[1].col = 1
  tests[1].size = 3
  tests[1].startIndex = 0
  tests[1].answer = []position{[3]int{1,1,-4}, [3]int{0,1,1}, [3]int{2,1,1}, [3]int{1,0,1}, [3]int{1,2,1}}
  tests[2].row = 2
  tests[2].col = 1
  tests[2].size = 40
  tests[2].startIndex = 16
  tests[2].answer = []position{[3]int{18,1,-4}, [3]int{17,1,1}, [3]int{19,1,1}, [3]int{18,0,1}, [3]int{18,2,1}}

  for i, test := range tests {
		outcome := Topple(test.row, test.col, test.size, test.startIndex)

		if !ToppleSeqEqual(outcome, test.answer) {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the row, col, size, startIndex is", test.row, test.col, test.size, test.startIndex , "the answer is", test.answer)
		}
	}
}

//ToppleSeqEqual checks if two slices of positions are equal or not
//Input: Two slices of position
//Output: Bool true or false
func ToppleSeqEqual(toppleA, toppleB []position)bool{

  var flag bool
  flag = true

  if len(toppleA) != len(toppleB){
    flag = false
    return flag
  }

  for i := range toppleA{
    for j := 0; j < 3; j++{
      if toppleA[i][j] != toppleB[i][j]{
        flag = false
        return flag
      }
    }
  }
  return flag
}

//Tests ToppleSeqEqual function 
func TestToppleSeqEqual(t *testing.T){

  type test struct {
		posA, posB []position
		answer bool
	}
  tests := make([]test, 3)
  tests[0].posA = []position{[3]int{0,0,-4}, [3]int{1,0,1}, [3]int{0,1,1}}
  tests[0].posB = []position{[3]int{0,0,-4}, [3]int{1,0,1}, [3]int{0,1,1}}
  tests[0].answer = true
  tests[1].posA = []position{[3]int{0,0,-4}, [3]int{1,0,1}, [3]int{0,1,1}}
  tests[1].posB = []position{[3]int{0,0,0}, [3]int{1,0,1}, [3]int{0,1,1}}
  tests[1].answer = false
  tests[2].posA = []position{[3]int{0,0,-4}, [3]int{1,0,1}, [3]int{0,1,1}}
  tests[2].posB = []position{[3]int{0,0,-4}, [3]int{1,0,1}}
  tests[2].answer = false

  for i, test := range tests {
		outcome := ToppleSeqEqual(test.posA, test.posB)

		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct answer is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! When the positions are", test.posA, "and", test.posB, "the answer is", test.answer)
		}
	}
}
