package weekly293

import (
	"fmt"
	"testing"
)

func TestConstructorX(t *testing.T) {
	c := ConstructorCountIntervals()
	c.Add(0, 10)
	fmt.Println(c.Count())
	c.Add(90, 100)
	fmt.Println(c.Count())
	c.Add(60, 80)
	t.Error(c.Count())
	c.Add(70, 90)
	t.Error(c.Count())
	//fmt.Println(c.arr, c.Count())
	//
	//c.Add(30, 50)
	//fmt.Println(c.arr, c.Count())

}

/*func TestConstructor1(t *testing.T) {
	c := Constructor()
	c.Add(1, 3)
	fmt.Println(c.arr, c.Count())
	c.Add(3, 6)
	fmt.Println(c.arr, c.Count())
	c.Add(4, 8)

	fmt.Println(c.arr, c.Count())
	c.Add(10, 12)
	fmt.Println(c.arr, c.Count())

	c.Add(7, 8)
	fmt.Println(c.arr, c.Count())

	c.Add(2, 6)
	fmt.Println(c.arr, c.Count())
}

func TestConstructor2(t *testing.T) {
	cmd := []string{"count", "count", "add", "count", "count", "add", "count", "count", "count", "count", "count", "add", "count", "count", "add", "count", "add", "add", "count", "count", "add", "add", "add", "count", "add", "add", "count", "add", "add", "add", "add", "add", "add", "count", "count", "add", "add", "add", "add", "add", "count", "count", "add", "count", "add", "count", "count", "add", "add", "add", "count", "count", "add", "count", "count", "add", "count", "add", "count", "add", "count", "count", "count", "count", "count", "count", "count", "count", "count", "count", "count", "add", "add", "count", "add", "add", "count", "count", "count", "add", "count", "add", "count", "add", "add", "add", "count", "count", "count", "add", "count", "add", "count", "count", "add", "add", "add", "count", "count", "count", "count", "count", "add", "count", "add", "add", "add", "add", "count", "add", "add", "count", "add", "count", "count", "add", "add", "count", "add", "add", "count", "add", "add", "add", "count", "count", "count", "add", "add", "add", "count"}
	nums := [][]int{{}, {}, {571, 770}, {}, {}, {920, 996}, {}, {}, {}, {}, {}, {65, 512}, {}, {}, {959, 959}, {}, {313, 330}, {473, 928}, {}, {}, {75, 561}, {107, 835}, {852, 918}, {}, {12, 774}, {534, 597}, {}, {743, 776}, {456, 556}, {727, 750}, {403, 954}, {342, 803}, {299, 807}, {}, {}, {905, 951}, {365, 686}, {211, 646}, {185, 216}, {826, 910}, {}, {}, {470, 501}, {}, {775, 959}, {}, {}, {343, 647}, {618, 743}, {203, 208}, {}, {}, {924, 962}, {}, {}, {307, 976}, {}, {465, 831}, {}, {411, 598}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {351, 936}, {209, 323}, {}, {69, 194}, {419, 794}, {}, {}, {}, {745, 962}, {}, {613, 690}, {}, {442, 520}, {259, 500}, {39, 272}, {}, {}, {}, {160, 853}, {}, {451, 519}, {}, {}, {363, 909}, {6, 770}, {819, 950}, {}, {}, {}, {}, {}, {865, 891}, {}, {630, 984}, {863, 966}, {385, 497}, {882, 885}, {}, {906, 930}, {705, 717}, {}, {654, 856}, {}, {}, {60, 693}, {666, 740}, {}, {721, 899}, {347, 438}, {}, {811, 994}, {67, 81}, {737, 898}, {}, {}, {}, {7, 664}, {104, 796}, {487, 866}, {}}
	valueStr := "0,0,null,200,200,null,277,277,277,277,277,null,725,725,null,725,null,null,932,932,null,null,null,932,null,null,985,null,null,null,null,null,null,985,985,null,null,null,null,null,985,985,null,985,null,985,985,null,null,null,985,985,null,985,985,null,985,null,985,null,985,985,985,985,985,985,985,985,985,985,985,null,null,985,null,null,985,985,985,null,985,null,985,null,null,null,985,985,985,null,985,null,985,985,null,null,null,991,991,991,991,991,null,991,null,null,null,null,991,null,null,991,null,991,991,null,null,991,null,null,991,null,null,null,991,991,991,null,null,null,991"
	values := strings.Split(valueStr, ",")
	c := Constructor()
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case "count":
			v, err := strconv.Atoi(values[i])
			if err != nil {
				t.Fatal(err)
			}
			if c.Count() != v {
				t.Log(c.arr)
				t.Fatal(c.Count(), v)
			} else {
				t.Log(c.arr, c.Count())
			}
		case "add":
			c.Add(nums[i][0], nums[i][1])
		}
	}
}

//{{},{},{457, 717},{918, 927},{},{660, 675},{},{},{885, 905},{},{},{323, 416},{774, 808},{},{671, 787},{133, 963},{374, 823},{},{494, 556},{},{},{564, 898},{},{514, 530},{},{702, 950},{282, 695},{324, 943},{},{893, 975},{337, 869},{674, 981},{},{847, 966},{12, 544},{706, 862},{},{},{},{},{},{},{633, 964},{17, 86},{209, 315},{},{},{974, 995},{},{},{106, 817},{873, 886},{},{},{},{255, 422},{},{},{71, 927},{852, 889},{980, 994},{},{},{477, 649},{},{},{76, 426},{508, 510},{},{},{345, 611},{816, 980},{},{},{848, 924},{855, 914},{621, 740},{176, 801},{258, 451},{229, 540},{119, 482},{801, 917},{},{},{308, 877},{},{},{927, 961},{546, 881},{563, 867},{64, 284},{},{},{289, 540},{},{},{},{},{851, 898},{},{},{330, 400},{},{724, 826},{},{316, 578},{},{95, 973},{},{469, 971},{568, 806},{775, 953},{556, 958},{109, 488},{142, 305},{},{},{},{686, 778}}
//{null, 0, null, null, 271,null, 271, 271, null, 292,292, null, null, 421, null,null, null, 1217, null, 1217,1217, null, 1217, null, 1217,null, null, null, 1217, null,null, null, 1235, null, null,null, 1356, 1356, 1356, 1356,1356, 1356, null, null, null,1356, 1356, null, 1370, 1370,null, null, 1370, 1370, 1370,null, 1370, 1370, null, null,null, 1370, 1370, null, 1370,1370, null, null, 1370, 1370,null, null, 1370, 1370, null,null, null, null, null, null,null, null, 1370, 1370, null,1370, 1370, null, null, null,null, 1370, 1370, null, 1370,1370, 1370, 1370, null, 1370,1370, null, 1370, null, 1370,null, 1370, null, 1370, null,null, null, null, null, null,1370, 1370, 1370, null}
func TestConstructor3(t *testing.T) {
	cmd := []string{"count", "add", "add", "count", "add", "count", "count", "add", "count", "count", "add", "add", "count", "add", "add", "add", "count", "add", "count", "count", "add", "count", "add", "count", "add", "add", "add", "count", "add", "add", "add", "count", "add", "add", "add", "count", "count", "count", "count", "count", "count", "add", "add", "add", "count", "count", "add", "count", "count", "add", "add", "count", "count", "count", "add", "count", "count", "add", "add", "add", "count", "count", "add", "count", "count", "add", "add", "count", "count", "add", "add", "count", "count", "add", "add", "add", "add", "add", "add", "add", "add", "count", "count", "add", "count", "count", "add", "add", "add", "add", "count", "count", "add", "count", "count", "count", "count", "add", "count", "count", "add", "count", "add", "count", "add", "count", "add", "count", "add", "add", "add", "add", "add", "add", "count", "count", "count", "add"}
	nums := [][]int{{}, {457, 717}, {918, 927}, {}, {660, 675}, {}, {}, {885, 905}, {}, {}, {323, 416}, {774, 808}, {}, {671, 787}, {133, 963}, {374, 823}, {}, {494, 556}, {}, {}, {564, 898}, {}, {514, 530}, {}, {702, 950}, {282, 695}, {324, 943}, {}, {893, 975}, {337, 869}, {674, 981}, {}, {847, 966}, {12, 544}, {706, 862}, {}, {}, {}, {}, {}, {}, {633, 964}, {17, 86}, {209, 315}, {}, {}, {974, 995}, {}, {}, {106, 817}, {873, 886}, {}, {}, {}, {255, 422}, {}, {}, {71, 927}, {852, 889}, {980, 994}, {}, {}, {477, 649}, {}, {}, {76, 426}, {508, 510}, {}, {}, {345, 611}, {816, 980}, {}, {}, {848, 924}, {855, 914}, {621, 740}, {176, 801}, {258, 451}, {229, 540}, {119, 482}, {801, 917}, {}, {}, {308, 877}, {}, {}, {927, 961}, {546, 881}, {563, 867}, {64, 284}, {}, {}, {289, 540}, {}, {}, {}, {}, {851, 898}, {}, {}, {330, 400}, {}, {724, 826}, {}, {316, 578}, {}, {95, 973}, {}, {469, 971}, {568, 806}, {775, 953}, {556, 958}, {109, 488}, {142, 305}, {}, {}, {}, {686, 778}}
	valueStr := "0,null,null,271,null,271,271,null,292,292,null,null,421,null,null,null,831,null,831,831,null,831,null,831,null,null,null,831,null,null,null,849,null,null,null,970,970,970,970,970,970,null,null,null,970,970,null,984,984,null,null,984,984,984,null,984,984,null,null,null,984,984,null,984,984,null,null,984,984,null,null,984,984,null,null,null,null,null,null,null,null,984,984,null,984,984,null,null,null,null,984,984,null,984,984,984,984,null,984,984,null,984,null,984,null,984,null,984,null,null,null,null,null,null,984,984,984,null"
	values := strings.Split(valueStr, ",")
	c := Constructor()
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case "count":
			v, err := strconv.Atoi(values[i])
			if err != nil {
				t.Fatal(err)
			}
			if c.Count() != v {
				t.Fatalf(" %v Count = %v, want %v", c.arr, c.Count(), v)
			} else {
				t.Logf(" %v Count = %v, want %v", c.arr, c.Count(), v)
			}
		case "add":
			c.Add(nums[i][0], nums[i][1])
			t.Log(nums[i][0], nums[i][1], c.arr, c.Count())
		}
	}
}

func TestConstructor4(t *testing.T) {
	cmd := []string{"count", "add", "add", "count", "add", "count", "count", "add", "count", "count", "add", "add", "count", "add", "add", "add", "count", "add", "count", "count", "add", "count", "add", "count", "add", "add", "add", "count", "add", "add", "add", "count", "add", "add", "add", "count", "count", "count", "count", "count", "count", "add", "add", "add", "count", "count", "add", "count", "count", "add", "add", "count", "count", "count", "add", "count", "count", "add", "add", "add", "count", "count", "add", "count", "count", "add", "add", "count", "count", "add", "add", "count", "count", "add", "add", "add", "add", "add", "add", "add", "add", "count", "count", "add", "count", "count", "add", "add", "add", "add", "count", "count", "add", "count", "count", "count", "count", "add", "count", "count", "add", "count", "add", "count", "add", "count", "add", "count", "add", "add", "add", "add", "add", "add", "count", "count", "count", "add"}
	nums := [][]int{{}, {457, 717}, {918, 927}, {}, {660, 675}, {}, {}, {885, 905}, {}, {}, {323, 416}, {774, 808}, {}, {671, 787}, {133, 963}, {374, 823}, {}, {494, 556}, {}, {}, {564, 898}, {}, {514, 530}, {}, {702, 950}, {282, 695}, {324, 943}, {}, {893, 975}, {337, 869}, {674, 981}, {}, {847, 966}, {12, 544}, {706, 862}, {}, {}, {}, {}, {}, {}, {633, 964}, {17, 86}, {209, 315}, {}, {}, {974, 995}, {}, {}, {106, 817}, {873, 886}, {}, {}, {}, {255, 422}, {}, {}, {71, 927}, {852, 889}, {980, 994}, {}, {}, {477, 649}, {}, {}, {76, 426}, {508, 510}, {}, {}, {345, 611}, {816, 980}, {}, {}, {848, 924}, {855, 914}, {621, 740}, {176, 801}, {258, 451}, {229, 540}, {119, 482}, {801, 917}, {}, {}, {308, 877}, {}, {}, {927, 961}, {546, 881}, {563, 867}, {64, 284}, {}, {}, {289, 540}, {}, {}, {}, {}, {851, 898}, {}, {}, {330, 400}, {}, {724, 826}, {}, {316, 578}, {}, {95, 973}, {}, {469, 971}, {568, 806}, {775, 953}, {556, 958}, {109, 488}, {142, 305}, {}, {}, {}, {686, 778}}
	valueStr := "0,null,null,271,null,271,271,null,292,292,null,null,421,null,null,null,831,null,831,831,null,831,null,831,null,null,null,831,null,null,null,849,null,null,null,970,970,970,970,970,970,null,null,null,970,970,null,984,984,null,null,984,984,984,null,984,984,null,null,null,984,984,null,984,984,null,null,984,984,null,null,984,984,null,null,null,null,null,null,null,null,984,984,null,984,984,null,null,null,null,984,984,null,984,984,984,984,null,984,984,null,984,null,984,null,984,null,984,null,null,null,null,null,null,984,984,984,null"
	values := strings.Split(valueStr, ",")
	c := Constructor()
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case "count":
			v, err := strconv.Atoi(values[i])
			if err != nil {
				t.Fatal(err)
			}
			if c.Count() != v {
				t.Fatalf(" %v Count = %v, want %v", c.arr, c.Count(), v)
			}
		case "add":
			c.Add(nums[i][0], nums[i][1])
			t.Log(nums[i][0], nums[i][1], c.arr, c.Count())
		}
	}
}

func TestConstructor5(t *testing.T) {
	c := Constructor()
	for i := 1000000; i < 900001; i = i - 2 {
		c.Add(i-1, i)
	}
	c.Add(900000, 1000000)
	for i := 900001; i < 1000000; i = i + 2 {
		fmt.Print(c.Count(), " ")
	}
}

//[]
//
//[null,0,0,null,200,200,null,277,277,277,277,277,null,725,725,null,726,null,null,1200,1200,null,null,null,1598,null,null,1661,null,null,null,null,null,null,1687,1687,null,null,null,null,null,1687,1687,null,1687,null,1692,1692,null,null,null,1692,1692,null,1695,1695,null,1709,null,1709,null,1709,1709,1709,1709,1709,1709,1709,1709,1709,1709,1709,null,null,1709,null,null,1709,1709,1709,null,1709,null,1709,null,null,null,1709,1709,1709,null,1709,null,1709,1709,null,null,null,1715,1715,1715,1715,1715,null,1715,null,null,null,null,1723,null,null,1723,null,1723,1723,null,null,1723,null,null,1723,null,null,null,1733,1733,1733,null,null,null,1733]
// [null,0,0,null,200,200,null,277,277,277,277,277,null,725,725,null,725,null,null,932,932,null,null,null,932,null,null,985,null,null,null,null,null,null,985,985,null,null,null,null,null,985,985,null,985,null,985,985,null,null,null,985,985,null,985,985,null,985,null,985,null,985,985,985,985,985,985,985,985,985,985,985,null,null,985,null,null,985,985,985,null,985,null,985,null,null,null,985,985,985,null,985,null,985,985,null,null,null,991,991,991,991,991,null,991,null,null,null,null,991,null,null,991,null,991,991,null,null,991,null,null,991,null,null,null,991,991,991,null,null,null,991]

func TestCountIntervals_updateData(t *testing.T) {
	c := Constructor()
	c.arr = []Item{{457, 717}, {918, 927}}
	leftVal, end := 0, 1
	c.updateData(leftVal, end, Item{885, 905})
	fmt.Println(c.arr)
}
*/
