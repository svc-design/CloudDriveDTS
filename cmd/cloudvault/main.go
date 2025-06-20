package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "login":
		loginCmd(os.Args[2:])
	case "sync":
		syncCmd(os.Args[2:])
	case "daemon":
		daemonCmd(os.Args[2:])
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: cloudvault <command> [options]")
	fmt.Println("Commands: login, sync, daemon")
}

func loginCmd(args []string) {
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	region := fs.String("region", "global", "Region for the provider")
	fs.Parse(args)

	if fs.NArg() < 1 {
		fmt.Println("Usage: cloudvault login <provider> [--region region]")
		os.Exit(1)
	}
	provider := fs.Arg(0)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	passBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	password := strings.TrimSpace(string(passBytes))

	fmt.Print("One-time code: ")
	otpBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	otp := strings.TrimSpace(string(otpBytes))

	cred := Credentials{Username: username, Password: password, OTP: otp}
	base := filepath.Join(os.Getenv("HOME"), ".cloudvault", provider, *region)
	os.MkdirAll(base, 0o700)
	path := filepath.Join(base, "credentials.json")
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to save credentials: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(cred); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write credentials: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Credentials saved to %s\n", path)
}

func syncCmd(args []string) {
	fs := flag.NewFlagSet("sync", flag.ExitOnError)
	src := fs.String("src", "", "Source path")
	dst := fs.String("dst", "", "Destination path")
	fs.Parse(args)

	if *src == "" || *dst == "" {
		fmt.Println("Usage: cloudvault sync --src <path> --dst <path>")
		os.Exit(1)
	}
	fmt.Printf("Simulating sync from %s to %s\n", *src, *dst)
}

func daemonCmd(args []string) {
	fs := flag.NewFlagSet("daemon", flag.ExitOnError)
	config := fs.String("config", "", "Config file")
	fs.Parse(args)

	if *config == "" {
		fmt.Println("Usage: cloudvault daemon --config <file>")
		os.Exit(1)
	}
	fmt.Printf("Starting daemon with config %s\n", *config)
}
