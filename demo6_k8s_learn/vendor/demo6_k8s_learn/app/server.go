package app

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
    componentKubelet = "kubelet"
)

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
    if strings.Contains(name, "_") {
        return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
    }
    return pflag.NormalizedName(name)
}

func PrintFlags(flags *pflag.FlagSet) {
    flags.VisitAll(func(flag *pflag.Flag) {
       fmt.Println("FLAG: --%s=%q", flag.Name, flag.Value)
    })
}

func NewKubeletCommand() *cobra.Command {
    cleanFlagSet := pflag.NewFlagSet(componentKubelet, pflag.ContinueOnError)
    cleanFlagSet.SetNormalizeFunc(WordSepNormalizeFunc)
	kubeletFlags := NewKubeletFlags()
	kubeletFlags.AddFlags(cleanFlagSet)

    cmd := &cobra.Command{
        Use: componentKubelet,
        Long: `The kubelet`,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
		    // initial flag parse, since we disable cobra's flag parsing
		    if err := cleanFlagSet.Parse(args); err != nil {
		        cmd.Usage()
		        fmt.Println(err)
		    }
		    // check if there are non-flag arguments in the command line
		    cmds := cleanFlagSet.Args()
		    if len(cmds) > 0 {
		        cmd.Usage()
		        fmt.Println("unknown command: %s", cmds[0])
		    }
		    // short-circuit on help
		    help, err := cleanFlagSet.GetBool("help")
		    if err != nil {
		        fmt.Println(`"help" flag is non-bool, programmer error, please correct`)
		    }
		    if help {
		        cmd.Help()
		        return
			}

			PrintFlags(cleanFlagSet)

			for {
				;
			}
		},
	}

	cleanFlagSet.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))

	// ugly, but necessary, because Cobra's default UsageFunc and HelpFunc pollute the flagset with global flags
	const usagefmt = "Usage:\n  %s\n\nFlags:\n%s"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
	    fmt.Fprintf(cmd.OutOrStderr(), usagefmt, cmd.UseLine(), cleanFlagSet.FlagUsagesWrapped(2))
	    return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
	    fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usagefmt, cmd.Long, cmd.UseLine(), cleanFlagSet.FlagUsagesWrapped(2))
	})
	return cmd
}
