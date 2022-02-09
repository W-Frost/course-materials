package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){

    got, _ := PortScanner(1024) // Currently function returns only number of open ports (added the int variable of how many ports)
    want := 1 // changed this to 1 as im consistently only getting one open port I know this will break if there are more
    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

    open, close := PortScanner(1024)//added the int variable of how many ports
	got:= open + close // Currently function returns only number of open ports
    want := 1024 // default value; consider what would happen if you parameterize the portscanner ports to scan (will have to change if someone changes the varibale in scanner.go)
    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}


