package cli

import (
	"github.com/spf13/cobra"
	"os"
)

// CompletionCmd adds shell completion support for bash, zsh, fish, and powershell.
var CompletionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion scripts",
	Long: `To load completions:

Bash:
  $ source <(ex-crl completion bash)
  # To load completions for each session, execute once:
  # Linux:
  $ ex-crl completion bash > /etc/bash_completion.d/ex-crl
  # macOS:
  $ ex-crl completion bash > /usr/local/etc/bash_completion.d/ex-crl

Zsh:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  $ ex-crl completion zsh > "${fpath[1]}/_ex-crl"

Fish:
  $ ex-crl completion fish | source
  $ ex-crl completion fish > ~/.config/fish/completions/ex-crl.fish

PowerShell:
  PS> ex-crl completion powershell | Out-String | Invoke-Expression
  # To load completions for every new session, run:
  PS> ex-crl completion powershell > ex-crl.ps1
  # and source this file from your PowerShell profile.
`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			RootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			RootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			RootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			RootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}
