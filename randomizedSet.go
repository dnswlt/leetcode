package leetcode

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]
	if !ok {
		this.m[val] = len(this.a)
		this.a = append(this.a, val)
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]
	if ok {
		l := len(this.a)
		if l > 1 {
			this.a[i], this.a[l-1] = this.a[l-1], this.a[i]
			this.m[this.a[i]] = i
		}
		this.a = this.a[:l-1]
		delete(this.m, val)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}
