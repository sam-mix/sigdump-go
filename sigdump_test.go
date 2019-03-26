package sigdump

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDumpGoroutine(t *testing.T) {
	f := new(bytes.Buffer)

	dumpGoroutine(f)

	actual := string(f.Bytes())
	expected := "sigdump-go.TestDumpGoroutine" // self test case name

	if !strings.Contains(actual, expected) {
		t.Fatalf("bad: not contains %s", expected)
		t.Fatalf("actual: %s", actual)
	}
}

func TestOpenDumpPath(t *testing.T) {
	// open Stdout
	actual := openDumpPath("-")
	expected := os.Stdout

	if actual != expected {
		t.Fatalf("bad: %+v", actual)
	}

	if "-" != DumpToStdout {
		t.Fatalf("bad: DumpToStdout not equals to \"-\"")
	}

	// open Stderr
	actual = openDumpPath("+")
	expected = os.Stderr

	if actual != expected {
		t.Fatalf("bad: %+v", actual)
	}

	if "+" != DumpToStderr {
		t.Fatalf("bad: DumpToStderr not equals to \"+\"")
	}

	// open file in /tmp dir
	f := openDumpPath("")
	f.Close()

	path := fmt.Sprintf("/tmp/sigdump-%d.log", os.Getpid())
	_, err := os.Stat(path)
	if err != nil {
		t.Fatalf("bad: faild to open dump file - %s", path)
	}

	// open specified file
	path = fmt.Sprintf("/tmp/sigdump-test-%d.log", os.Getpid())
	f = openDumpPath(path)
	f.Close()

	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("bad: faild to open dump file - %s", path)
	}
	os.Remove(path)

	// open Stderr if file not exists
	actual = openDumpPath("/tmp/not_found_dir/not_found_file")
	expected = os.Stderr

	if actual != expected {
		t.Fatalf("bad: %+v", actual)
	}
}
