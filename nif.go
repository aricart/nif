package main

import (
	"fmt"
	"net"
	"os"
	"text/tabwriter"
)


func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("PrintAddresses: %+v\n", err.Error()))
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	fmt.Fprintf(w, "%s \t%s \t %s\n", "Type", "Name", "IP")
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("PrintAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				if v.IP.IsLinkLocalMulticast() {
					fmt.Fprintf(w, "%s", "LLM ")
				}
				if v.IP.IsLinkLocalUnicast() {
					fmt.Fprintf(w, "%s", "LLU ")
				}
				if v.IP.IsInterfaceLocalMulticast() {
					fmt.Fprintf(w, "%s", "ILM ")
				}
				if v.IP.IsLoopback() {
					fmt.Fprintf(w, "%s", "L ")
				}
				if v.IP.IsUnspecified() {
					fmt.Fprintf(w, "%s", "U ")
				}
				if v.IP.IsGlobalUnicast() {
					fmt.Fprintf(w, "%s", "GU ")
				}

				fmt.Fprintf(w, "\t%v \t%v\n", i.Name, v.IP.String())
			case *net.IPNet:
				if v.IP.IsLinkLocalMulticast() {
					fmt.Fprintf(w, "%s", "LLM ")
				}
				if v.IP.IsLinkLocalUnicast() {
					fmt.Fprintf(w, "%s", "LLU ")
				}
				if v.IP.IsInterfaceLocalMulticast() {
					fmt.Fprintf(w, "%s", "ILM ")
				}
				if v.IP.IsLoopback() {
					fmt.Fprintf(w, "%s", "L ")
				}
				if v.IP.IsUnspecified() {
					fmt.Fprintf(w, "%s", "U ")
				}
				if v.IP.IsGlobalUnicast() {
					fmt.Fprintf(w, "%s", "GU ")
				}

				fmt.Fprintf(w, "\t%v \t%v\n", i.Name, v.IP.String())
			default:

			}

		}
	}

	_ = w.Flush()
}
