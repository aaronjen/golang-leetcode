package main

/*
 * @lc app=leetcode id=638 lang=golang
 *
 * [638] Shopping Offers
 */

// @lc code=start
func shoppingOffersSolve(price []int, special [][]int, needs []int, pos int) int {
	localMin := directBuy(price, needs)
	for i := pos; i < len(special); i++ {
		offer := special[i]
		tmp := []int{}
		for j := 0; j < len(needs); j++ {
			if needs[j] < offer[j] {
				tmp = nil
				break
			}
			tmp = append(tmp, needs[j]-offer[j])
		}

		if tmp != nil {
			newValue := offer[len(offer)-1] + shoppingOffersSolve(price, special, tmp, i)
			if newValue < localMin {
				localMin = newValue
			}
		}
	}
	return localMin
}

func directBuy(price []int, needs []int) int {
	sum := 0
	for i := 0; i < len(price); i++ {
		sum += price[i] * needs[i]
	}

	return sum
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	return shoppingOffersSolve(price, special, needs, 0)
}

// @lc code=end
