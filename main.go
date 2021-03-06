package main

import (
	"github.com/PremiereGlobal/stim/stim"
	"github.com/PremiereGlobal/stim/stimpacks/aws"
	"github.com/PremiereGlobal/stim/stimpacks/completion"
	"github.com/PremiereGlobal/stim/stimpacks/deploy"
	"github.com/PremiereGlobal/stim/stimpacks/kubernetes"
	"github.com/PremiereGlobal/stim/stimpacks/pagerduty"
	"github.com/PremiereGlobal/stim/stimpacks/slack"
	"github.com/PremiereGlobal/stim/stimpacks/vault"
	"github.com/PremiereGlobal/stim/stimpacks/version"
)

func main() {
	stim := stim.New()
	stim.AddStimpack(aws.New())
	stim.AddStimpack(completion.New())
	stim.AddStimpack(deploy.New())
	stim.AddStimpack(kubernetes.New())
	stim.AddStimpack(pagerduty.New())
	stim.AddStimpack(slack.New())
	stim.AddStimpack(vault.New())
	stim.AddStimpack(version.New())
	stim.Execute()
}
