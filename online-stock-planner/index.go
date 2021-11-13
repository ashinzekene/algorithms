package algorithms

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: make([][2]int, 0),
	}
}

func (ss *StockSpanner) Next(price int) int {
	span := 1
	if len(ss.stack) == 0 {
		ss.stack = append(ss.stack, [2]int{price, 1})
		return 1
	} else {
		lastItem := ss.stack[len(ss.stack)-1]
		for len(ss.stack) > 0 && lastItem[0] <= price {
			span += lastItem[1]
			ss.stack = ss.stack[:len(ss.stack)-1]
			if len(ss.stack) > 0 {
				lastItem = ss.stack[len(ss.stack)-1]
			}
		}
		ss.stack = append(ss.stack, [2]int{price, span})
	}
	return span
}
