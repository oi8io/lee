package weekly

import (
	"container/heap"
)

type FoodRatings struct {
	cuisinesMap map[string]*hp
	foodMap     map[string]*pair2
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	cuisinesMap := make(map[string]*hp)
	foodMap := make(map[string]*pair2)
	for i, food := range foods {
		cuisine := cuisines[i]
		foodMap[food] = &pair2{cuisine: cuisine, rating: ratings[i]}
		if _, has := cuisinesMap[cuisine]; !has {
			cuisinesMap[cuisine] = &hp{}
		}
		cuisinesMap[cuisine].push(&pair{food: food, rating: ratings[i]})
	}
	return FoodRatings{foodMap: foodMap, cuisinesMap: cuisinesMap}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.foodMap[food].rating = newRating
	cuisine := this.foodMap[food].cuisine
	this.cuisinesMap[cuisine].push(&pair{food: food, rating: newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	for {
		res := this.cuisinesMap[cuisine].top()
		food := res.food
		rating := res.rating
		if this.foodMap[food].rating == rating {
			return food
		}
		this.cuisinesMap[cuisine].pop()
	}
}

type pair struct {
	food   string
	rating int
}

type pair2 struct {
	cuisine string
	rating  int
}

type hp []*pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].rating > h[j].rating || (h[i].rating == h[j].rating && h[i].food < h[j].food)
}
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(*pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v *pair)         { heap.Push(h, v) }
func (h *hp) pop() *pair           { return heap.Pop(h).(*pair) }
func (h *hp) top() *pair           { a := *h; return a[0] }

/**

 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
