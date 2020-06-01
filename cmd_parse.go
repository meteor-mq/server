package server

// HelpCmd help command
type HelpCmd struct {
	// This root command has no default action, so print the help.
}

// command param
type argsCmd struct {
	// config file path
	cfgPath string
	// Whether to run in the background
	isBackground bool
}

// GetHelpCmd get help command root
func GetHelpCmd() *HelpCmd {
	return &HelpCmd{}
}
