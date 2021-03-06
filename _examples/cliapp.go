package main

import (
	"github.com/gookit/cliapp"
	"github.com/gookit/cliapp/builtin"
	"github.com/gookit/cliapp/_examples/cmd"
	// "github.com/gookit/cliapp/builtin/filewatcher"
	// "github.com/gookit/cliapp/builtin/reverseproxy"
	"runtime"
)

// run:
// go run ./_examples/cliapp.go
// go build ./_examples/cliapp.go && ./cliapp
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cliapp.NewApp(func(app *cliapp.Application) {
		app.Version = "1.0.6"
		app.Description = "this is my cli application"
		app.Hooks[cliapp.EvtInit] = func(a *cliapp.Application, data interface{}) {
			// do something...
			// fmt.Println("init app")
		}
		// app.SetVerbose(cliapp.VerbDebug)
		// app.DefaultCommand("example")
	})

	app.Add(cmd.ExampleCommand())
	app.Add(cmd.EnvInfoCommand())
	app.Add(cmd.GitCommand())
	app.Add(cmd.ColorCommand(), cmd.EmojiDemoCmd())
	app.Add(cmd.ShowDemoCommand(), cmd.ProgressDemoCmd(), cmd.SpinnerDemoCmd(), cmd.InteractDemoCommand())
	app.Add(builtin.GenEmojiMapCommand())

	// app.Add(filewatcher.FileWatcher(nil))
	// app.Add(reverseproxy.ReverseProxyCommand())

	app.Add(&cliapp.Command{
		Name:    "test",
		Aliases: []string{"ts"},
		UseFor:  "this is a description <info>message</> for command {$cmd}",
		Func: func(cmd *cliapp.Command, args []string) int {
			cliapp.Print("hello, in the test command\n")
			return 0
		},
	})

	app.Add(builtin.GenAutoCompleteScript())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run()
}
