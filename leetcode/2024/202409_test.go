package _024

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_maxNumOfMarkedIndices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40},
			},
			wantRes: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxNumOfMarkedIndices(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("maxNumOfMarkedIndices() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_maximumRobots(t *testing.T) {
	type args struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				chargeTimes:  []int{3, 6, 1, 3, 4},
				runningCosts: []int{2, 1, 3, 4, 5},
				budget:       25,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				chargeTimes:  []int{11, 12, 19},
				runningCosts: []int{10, 8, 7},
				budget:       19,
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				chargeTimes:  []int{11, 12, 74, 67, 37, 87, 42, 34, 18, 90, 36, 28, 34, 20},
				runningCosts: []int{18, 98, 2, 84, 7, 57, 54, 65, 59, 91, 7, 23, 94, 20},
				budget:       937,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRobots(tt.args.chargeTimes, tt.args.runningCosts, tt.args.budget); got != tt.want {
				t.Errorf("maximumRobots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_latestTimeCatchTheBus(t *testing.T) {
	type args struct {
		buses      []int
		passengers []int
		capacity   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				buses:      []int{10, 20},
				passengers: []int{2, 17, 18, 19},
				capacity:   2,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := latestTimeCatchTheBus(tt.args.buses, tt.args.passengers, tt.args.capacity); got != tt.want {
				t.Errorf("latestTimeCatchTheBus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				values: []int{8, 1, 5, 2, 6},
			},
			want: 11,
		},
		{
			name: "case 2",
			args: args{
				values: []int{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_takeCharacters(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "aabaaaacaabc",
				k: 2,
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				s: "a",
				k: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	s := Constructor(18)
	fmt.Println(s.Reserve())
	fmt.Println(s.Reserve())

}

func TestOfficialMaxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OfficialMaxScoreSightseeingPair(tt.args.values); got != tt.want {
				t.Errorf("OfficialMaxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_busyStudent(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		queryTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := busyStudent(tt.args.startTime, tt.args.endTime, tt.args.queryTime); got != tt.want {
				t.Errorf("busyStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearDigits(tt.args.s); got != tt.want {
				t.Errorf("clearDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWays(tt.args.nums); got != tt.want {
				t.Errorf("countWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distanceBetweenBusStops(t *testing.T) {
	type args struct {
		distance    []int
		start       int
		destination int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetweenBusStops(tt.args.distance, tt.args.start, tt.args.destination); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distinctNames(t *testing.T) {
	type args struct {
		ideas []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctNames(tt.args.ideas); got != tt.want {
				t.Errorf("distinctNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dp(t *testing.T) {
	type args struct {
		s          string
		k          int
		cnts       map[byte]int
		steps      int
		leftIndex  int
		rightIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp(tt.args.s, tt.args.k, tt.args.cnts, tt.args.steps, tt.args.leftIndex, tt.args.rightIndex); got != tt.want {
				t.Errorf("dp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_edgeScore(t *testing.T) {
	type args struct {
		edges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edgeScore(tt.args.edges); got != tt.want {
				t.Errorf("edgeScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_latestTimeCatchTheBus1(t *testing.T) {
	type args struct {
		buses      []int
		passengers []int
		capacity   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := latestTimeCatchTheBus(tt.args.buses, tt.args.passengers, tt.args.capacity); got != tt.want {
				t.Errorf("latestTimeCatchTheBus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestContinuousSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestContinuousSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestContinuousSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxNumOfMarkedIndices1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxNumOfMarkedIndices(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("maxNumOfMarkedIndices() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_maxScoreSightseeingPair1(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.args.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxStrength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxStrength(tt.args.nums); got != tt.want {
				t.Errorf("maxStrength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximizeWin(t *testing.T) {
	type args struct {
		prizePositions []int
		k              int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeWin(tt.args.prizePositions, tt.args.k); got != tt.want {
				t.Errorf("maximizeWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumRobots1(t *testing.T) {
	type args struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumRobots(tt.args.chargeTimes, tt.args.runningCosts, tt.args.budget); gotAns != tt.wantAns {
				t.Errorf("maximumRobots() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_maximumSubsequenceCount(t *testing.T) {
	type args struct {
		text    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubsequenceCount(tt.args.text, tt.args.pattern); got != tt.want {
				t.Errorf("maximumSubsequenceCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeNodes2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfPoints(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPoints(tt.args.nums); got != tt.want {
				t.Errorf("numberOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStars(tt.args.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_takeCharacters1(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_takeCharacters2(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeRequiredToBuy1(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
