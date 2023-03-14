package something

import "testing"

func BenchmarkMakeSomething(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MakeSomething(1000)
	}
}
