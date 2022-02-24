package main

func main() {

	// iota는 정수 뿐만 아닌 부동 소수점 타입에도 사용할 수 있고 계산식과 혼합해 사용할 수 있다.
	type ByteSize int64

	const (
		_           = iota             // ignore
		KB ByteSize = 1 << (10 * iota) // 1 << (10 * 1) = 1024
		MB                             // MegaByte 1 << (10 * 2) = 1048576
		GB                             // GigaByte 1 << (10 * 3) = 1073741824
		TB                             // TeraByte 1 << (10 * 4) = 1099511627776
		PB                             // Petabyte 1 << (10 * 5) = 1125899906842624
		EB                             // ExaByte  1 << (10 * 6) = 1152921504606846976
	)

	const (
		DEFAULT_RATE = 5 + 0.3*iota // 5
		GREEN_RATE                  // 5.3
		SILVER_RATE                 // 5.6
		GOLD_RATE                   // 5.9
	)

}
