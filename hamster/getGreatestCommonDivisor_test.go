package hamster

import (
	"fmt"
	"testing"
)

func Test_getGreatestCommonDivisor(t *testing.T) {
	fmt.Println(getGreatestCommonDivisor(20, 5))
	fmt.Println(getGreatestCommonDivisor(100, 80))
	fmt.Println(getGreatestCommonDivisor(84, 9))

	fmt.Println(getGreatestCommonDivisorV2(20, 5))
	fmt.Println(getGreatestCommonDivisorV2(100, 80))
	fmt.Println(getGreatestCommonDivisorV2(84, 9))

	fmt.Println(getGreatestCommonDivisorV3(20, 5))
	fmt.Println(getGreatestCommonDivisorV3(100, 80))
	fmt.Println(getGreatestCommonDivisorV3(84, 9))
}
