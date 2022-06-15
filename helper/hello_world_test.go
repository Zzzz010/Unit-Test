package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestTableTest(t *testing.T) {
	tes := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Raditya)",
			request:  "Raditya",
			expected: "Hi Raditya",
		},
		{
			name:     "HelloWorld(Bagus)",
			request:  "Bagus",
			expected: "Hello Bagus",
		},
	}

	for _, test := range tes {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Raung", func(t *testing.T) {
		hasil := HelloWorld("Raung")
		require.Equal(t, "Hi Raung", hasil, "Harus Hi Raung")
	})
	t.Run("Jayan", func(t *testing.T) {
		hasil := HelloWorld("Jayan")
		require.Equal(t, "Hello Jayan", hasil, "Harus Hello Jayan")
	})

}

func TestMain(t *testing.M) {
	fmt.Println("Sebelum Unit Test Dimulai")

	t.Run() // semua function akan dijanlakan

	fmt.Println("Sesudah Unit Test Dijalankan")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa dijalankan di Windows")
	}

	hasil := HelloWorld("Raul")
	require.Equal(t, "Hello Raung", hasil, "Mestinya 'Hello Raung'")

}

func TestHelloWorldAssertion(t *testing.T) {
	hasil := HelloWorld("Raung")

	assert.Equal(t, "Hello Raung", hasil, "hasilnya harus 'Hello Raung'")

	fmt.Println("Selesai Eksekusi")

}

func TestHelloWorldRequire(t *testing.T) {
	hasil := HelloWorld("Raung")

	require.Equal(t, "Hello Raung", hasil, "hasilnya harus 'Hello Raung'")

	fmt.Println("Tidak Dieksekusi")

}

func TestHelloWorld1(t *testing.T) {
	hasil := HelloWorld("Raung")
	if hasil != "Hello Raung" {
		// unit test failed
		t.Fail()
	}

	fmt.Println("TestHelloWorld1 done")
}

func TestHelloWorldJayan(t *testing.T) {
	hasil := HelloWorld("Jayan")
	if hasil != "Hello Jayan" {
		// unit test failed
		t.FailNow()
	}
}

func TestHelloWorldRaul(t *testing.T) {
	hasil := HelloWorld("Raul")
	if hasil != "Hello Raul" {
		// unit test failed
		t.Error("Output Harus Hello Raul")
	}

	fmt.Println("Tereksekusi")
}

func TestHelloWorldKama(t *testing.T) {
	hasil := HelloWorld("Kama")
	if hasil != "Hello Kama" {
		// unit test failed
		t.Fatal("Output Harus Hello Kama")
	}

	fmt.Println("Tidak tereksekusi")
}
