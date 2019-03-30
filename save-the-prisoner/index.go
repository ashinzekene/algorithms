package algorithms

import "fmt"

func main() {
	fmt.Printf("Expected %d got %d \n", 122129406, saveThePrisoner(352926151, 380324688, 94730870))
	fmt.Printf("Expected %d got %d \n", 23525398, saveThePrisoner(94431605, 679262176, 5284458))
	fmt.Printf("Expected %d got %d \n", 72975907, saveThePrisoner(208526924, 756265725, 150817879))
	fmt.Printf("Expected %d got %d \n", 405956028, saveThePrisoner(962975336, 972576181, 396355184))
	fmt.Printf("Expected %d got %d \n", 265162707, saveThePrisoner(464237185, 937820284, 255816794))
	fmt.Printf("Expected %d got %d \n", 91803604, saveThePrisoner(649320641, 742902564, 647542323))
	fmt.Printf("Expected %d got %d \n", 82636723, saveThePrisoner(174512033, 711706897, 68977959))
	fmt.Printf("Expected %d got %d \n", 9258153, saveThePrisoner(107283902, 156676511, 67149447))
	fmt.Printf("Expected %d got %d \n", 81152217, saveThePrisoner(104513201, 298399273, 96292548))
	fmt.Printf("Expected %d got %d \n", 31978708, saveThePrisoner(113378824, 274011418, 98103763))
	fmt.Printf("Expected %d got %d \n", 269539705, saveThePrisoner(934815799, 991959468, 212396037))
	fmt.Printf("Expected %d got %d \n", 18445097, saveThePrisoner(88325121, 435452998, 24617705))
	fmt.Printf("Expected %d got %d \n", 115810626, saveThePrisoner(984873585, 997634055, 103050157))
	fmt.Printf("Expected %d got %d \n", 117586280, saveThePrisoner(344218387, 715364875, 90658180))
	fmt.Printf("Expected %d got %d \n", 216062299, saveThePrisoner(556065259, 615709967, 156417592))
	fmt.Printf("Expected %d got %d \n", 55859471, saveThePrisoner(57109959, 451440582, 4188603))
	fmt.Printf("Expected %d got %d \n", 110226121, saveThePrisoner(353972922, 573651462, 244520504))
	fmt.Printf("Expected %d got %d \n", 674567416, saveThePrisoner(946486979, 973168361, 647886035))
	fmt.Printf("Expected %d got %d \n", 49690850, saveThePrisoner(368127406, 680428368, 105517295))
	fmt.Printf("Expected %d got %d \n", 200235842, saveThePrisoner(884634320, 965112925, 119757238))
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	no_of_sweets := m % n
	return s + no_of_sweets - 1
}
