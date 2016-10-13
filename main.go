package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rekby/mbr"
)

func main() {

	if len(os.Args) < 2 || len(os.Args[1]) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Opening %s -- %s", os.Args[1], err)
		os.Exit(1)
	}
	defer f.Close()

	Mbr, err := mbr.Read(f)
	if err != nil {
		fmt.Printf("Reading MBR -- %s", err)
		os.Exit(2)
	}

	if err := Mbr.Check(); err != nil {
		fmt.Printf("MBR invalid -- %s", err)
		os.Exit(3)
	}

	var offset uint32
	partitions := Mbr.GetAllPartitions()

	for _, p := range partitions {
		if !p.IsEmpty() {
			fmt.Printf("Partition: start %d, len %d, type: %02x\n", p.GetLBAStart(), p.GetLBALen(), p.GetType())
			fmt.Printf("         : mount at %d, size is %d\n", p.GetLBAStart()*512, p.GetLBALen()*512)
			offset = p.GetLBAStart() * 512
		}
	}

	target := "/mnt/tmp"
	if len(os.Args) >= 3 && len(os.Args[2]) > 0 {
		target = os.Args[2]
	}

	cmd := exec.Command("sudo", "mount", "-o",
		fmt.Sprintf("loop,offset=%d", offset), os.Args[1], target)
	//		fmt.Sprintf("ro,loop,offset=%d", offset), os.Args[1], target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
