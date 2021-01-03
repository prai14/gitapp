module gitApp

go 1.15

replace gitapp/oauth => ./oauth
replace gitapp/user => ./user

require (
	gitapp/oauth v1.1.0
    gitapp/user v1.1.0
	github.com/spf13/viper v1.7.1 // indirect
)
