package main

type Tsp float64
type TBs float64

type ML float64

// func tspToMl(tsp Tsp) ML {
// 	return ML(tsp * 4.92)
// }

// func tbsToMl(tbs TBs) ML {
// 	return ML(tbs * 14.79)
// }

func (tsp Tsp) ToML() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToML() ML {
	return ML(tbs * 14.79)
}

func printMeDefined() {
	// ml1 := ML(Tsp(3) * 4.92)
	// fmt.Printf("3 tsps = %.2f ML\n", ml1)
	// ml2 := ML(TBs(3) * 14.79)

	// fmt.Printf("3 tsps = %.2ff ML\n", ml2)

	// fmt.Println("testing functions on ML")
	// fmt.Println("tspToMl of 3 tsp: ", tspToMl(3))
	// fmt.Println("tbsToMl of 3 tbs: ", tbsToMl(3))
	// fmt.Println("Sum of 2 user defined types")
	// fmt.Println("2 tsp + 3 tbs = ", Tsp(2), TBs(3))
	// tsp := Tsp(2)
	// fmt.Println("2 tsp = ", tspToMl(tsp))

}
