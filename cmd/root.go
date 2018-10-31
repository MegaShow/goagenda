package cmd

import (
	"fmt"
	"github.com/MegaShow/goagenda/config"
	"github.com/MegaShow/goagenda/controller"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var cfgFile string
var verbose bool

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "An meeting management system",
	Long: `Agenda is a meetings management system.
This application is a perfect and essential tool
to be well organized in your work.`,
	Run:               func(cmd *cobra.Command, args []string) { cmd.Usage() },
	PersistentPostRun: controller.CtrlRelease,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLog)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $AGENDA_HOME/.agenda.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display the verbose information")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".agenda" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + string(os.PathSeparator) + ".agenda")
		viper.SetConfigName(".agenda")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		if cfgFile == "" {
			path := home + string(os.PathSeparator) + ".agenda"
			if _, err := os.Stat(path); err != nil {
				err := os.MkdirAll(path, 0777)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
			}
			f, err := os.OpenFile(path+string(os.PathSeparator)+".agenda.yaml", os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			defer f.Close()
			data, err := yaml.Marshal(config.Default())
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			_, err = f.Write(data)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			fmt.Println("\nCreate default config in \"" + path + "\"")
			fmt.Println("Please rerun agenda")
		}
		os.Exit(1)
	}

	arr := strings.Split(viper.ConfigFileUsed(), string(os.PathSeparator))
	dir := strings.Join(arr[:len(arr)-1], string(os.PathSeparator)) + string(os.PathSeparator)
	viper.Set("Log.Path", dir+viper.GetString("Log.Path"))
	viper.Set("Database.Path", dir+viper.GetString("Database.Path"))
}

func initLog() {
	log.SetVerbose(verbose)
}
