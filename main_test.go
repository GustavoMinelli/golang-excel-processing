package main

import (
	"os"
	"runtime/pprof"
	"testing"
)

func TestMainExecution(b *testing.T) {

	memFile, err := os.Create("memory.pprof")
	checkError(err, b)
	defer memFile.Close()

	cpuFile, err := os.Create("cpu.pprof")
	checkError(err, b)
	defer cpuFile.Close()

	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	main()

	pprof.WriteHeapProfile(memFile)

}

func checkError(err error, b *testing.T) {
	if err != nil {
		b.Fatal(err)
	}
}
