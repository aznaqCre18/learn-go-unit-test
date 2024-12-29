package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// untuk menjalankan perintah test dapat menggunakan command "go test"
// untuk melihat detail function test mana saja yang sudah di running dapat menggunakan command "go test -v"
// untuk menjalankan function test secasra spesifik dapat menjalankan perintah "go test -v -run TestNamaFunction"
// dan jika ingin menjalankan semua function test daritop folder modiule dapat menjalankan perintah "go test ./..."

// untuk handling error ada dari gosudah menyediakan yaitu t.Error() dan t.Fatal(),
// bedanya t.Error() akan mereturn error testing bersamaan dengan log nya tapi proses testing di functiontest tersebut akan tetap dijalankan
// dan untuk t.Fatal() akan mereturn error testing dengan log parameternya tapi proses testing di function test tersebut akan langsung berhenti

// untuk handling error menggunakan assertion dengan package testify, ada dua juga yaitu assert dan require
// bedanya untuk assert sama seperti t.Error() dan untuk require sama seperti t.Fatal()

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Aziz")

	if result != "Hello Aziz" {
		t.Error("Expected 'Hello Aziz', got ", result)
	}

}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Aziz")

	if result != "Hello Aziz" {
		t.Fatal("Result is not Hello Aziz!")
	}
}

// implementasi assertion dan require dari module package testify

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Aziz")
	require.Equal(t, "Hello Aziz", result)

	fmt.Println("Apakah dijalankan ?")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Aziz")
	assert.Equal(t, "Hello Aziz", result)

	fmt.Println("Apakah assert dijalankan ?")
}

// implementasi skip unit testing

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa berjalan di mac device")
	}

	result := HelloWorld("Aziz")
	require.Equal(t, "Hello Aziz", result)
}

// implementasi Before dan After menggunakan func TestMain()
// berfungsi untuk mengatur jalannya sebuah unit testing, jika ada fungsi TestMain dalam package tersebut maka ketika
// akan melakukan unit testing fungsi ini yang akan dijalankan

func TestMain(m *testing.M) {
	fmt.Println("Dijalankan sebelum unit test di eksekusi")

	// ini perintah atau code untuk menjalankan unit test
	m.Run()

	fmt.Println("Dijalankan setelah unit test selesai di eksekusi")
}

// implementasi test subtest

/*
 - Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
   go test -run TestNamaFunction
 - Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
   go test -run TestNamaFunction/NamaSubTest
 - Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
   go test -run /NamaSubTest
*/

func TestHelloWorldSubTest(t *testing.T) {
	t.Run("Aziz", func(t *testing.T) {
		result := HelloWorld("Aziz")
		require.Equal(t, "Hello Aziz", result)
	})

	t.Run("Darko", func(t *testing.T) {
		result := HelloWorld("Darko")
		require.Equal(t, "Hello Darko", result)
	})
}

// implementasi table test
// Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi hasil dari unit test
// Lalu slice tersebut kita iterasi menggunakan sub test

func TestHelloWorldTableTest(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		request  string
	}{
		{
			name:     "HelloWorld(Aziz)",
			expected: "Hello Aziz",
			request:  "Aziz",
		},
		{
			name:     "HelloWorld(Umi)",
			expected: "Hello Umi",
			request:  "Umi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// implementasi unit testing untuk benchmark
/*
 - Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench :
   go test -v -bench=.
 - Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah :
   go test -v -run=NotMathUnitTest -bench=.
 - Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah :
   go test -v -run=NotMathUnitTest -bench=BenchmarkTest
 - Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah :
   go test -v -bench=. ./...
*/

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Aziz")
	}
}

// implementasi benchmark test menggunakan sub benchmark
// Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah :
// go test -v -bench=BenchmarkNama/NamaSub

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Aziz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aziz")

		}
	})

	b.Run("Darko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Darko")
		}
	})
}

// implementasi benchmark test table
func BenchmarkHelloWorldTableTest(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Aziz)",
			request: "Aziz",
		},
		{
			name:    "HelloWorld(Umi)",
			request: "Umi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// PR BELAJAR UNIT TEST MOCK MENGGUNAKAN TESTIFY MOCK
