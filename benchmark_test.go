package chunkreader_test

import (
	"io"
	"os"
	"testing"

	"github.com/andyollylarkin/chunkreader"
)

func BenchmarkChunkReader(b *testing.B) {
	f, err := os.Open("./500mb_file.bin")
	if err != nil {
		b.Fatal(err)
	}
	outF, err := os.Create("./dst_file.bin")
	if err != nil {
		b.Fatal(err)
	}

	defer os.Remove("./dst_file.bin")

	for i := 0; i < b.N; i++ {
		cr := chunkreader.NewChunkReader(f, 5<<20)
		_, err := io.Copy(outF, cr)
		if err != nil {
			b.Fatal(err)
		}
	}
}
