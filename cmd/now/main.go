package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Precision string

const (
	Millis Precision = "millis"
	Secs   Precision = "secs"
)

var rootCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current Unix timestamp",
	Run:   runNow,
}

var precision Precision

func init() {
	rootCmd.Flags().VarP(&precisionFlag{&precision}, "precision", "p", "Precision level (millis or secs)")
	precision = Secs
}

func runNow(cmd *cobra.Command, args []string) {
	now := time.Now()
	switch precision {
	case Millis:
		fmt.Println(now.UnixMilli())
	case Secs:
		fmt.Println(now.Unix())
	}
}

type precisionFlag struct {
	value *Precision
}

func (p *precisionFlag) String() string {
	return string(*p.value)
}

func (p *precisionFlag) Set(v string) error {
	switch v {
	case "millis":
		*p.value = Millis
	case "secs":
		*p.value = Secs
	default:
		return fmt.Errorf("invalid precision: %s", v)
	}
	return nil
}

func (p *precisionFlag) Type() string {
	return "precision"
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}