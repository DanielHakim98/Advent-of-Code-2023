/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DanielHakim98/aoc/day1"
	"github.com/spf13/cobra"
)

// dayOneCmd represents the dayOne command
var dayOneCmd = &cobra.Command{
	Use:   "dayOne",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("No filename is passed in")
		}
		d := day1.Day1{}
		display := "sum of calibration values: "
		switch part {
		case 1:
			fmt.Println(display, d.PartOne(args[0], day1.GetInput))
		case 2:
			fmt.Println(display, d.PartTwo(args[0], day1.GetInput))
		default:
			log.Fatal("Invalid 'part' flag")
		}

	},
}

func init() {
	rootCmd.AddCommand(dayOneCmd)
	dayOneCmd.Flags().IntVarP(&part, "part", "p", 0, "part to run")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayOneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayOneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
