package debug

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Debug command flag names and usage descriptions
const (
	FlagRuntime      = "runtime"
	FlagRuntimeUsage = "Runtime environment type"

	FlagTarget      = "target"
	FlagTargetUsage = "Target container (name or ID)"

	FlagNamespace      = "namespace"
	FlagNamespaceUsage = "Namespace to target (k8s runtime)"

	FlagPod      = "pod"
	FlagPodUsage = "Pod to target (k8s runtime)"

	FlagDebugImage      = "debug-image"
	FlagDebugImageUsage = "Debug image to use for the debug side-car container"

	FlagEntrypoint      = "entrypoint"
	FlagEntrypointUsage = "Custom ENTRYPOINT to use for the debug side-car container."

	FlagCmd      = "cmd"
	FlagCmdUsage = "Custom CMD to use for the debug side-car container (alternatively pass custom CMD params after '--')."

	FlagTerminal      = "terminal"
	FlagTerminalUsage = "Attach interactive terminal to the debug container"

	FlagListDebugImage      = "list-debug-images"
	FlagListDebugImageUsage = "List possible debug images to use for the debug side-car container (use this flag by itself)"

	FlagKubeconfig      = "kubeconfig"
	FlagKubeconfigUsage = "Kubeconfig file location (k8s runtime)"
)

const (
	NicolakaNetshootImage = "nicolaka/netshoot"
	KoolkitsNodeImage     = "lightruncom/koolkits:node"
	KoolkitsPythonImage   = "lightruncom/koolkits:python"
	KoolkitsGolangImage   = "lightruncom/koolkits:golang"
	KoolkitsJVMImage      = "lightruncom/koolkits:jvm"
	DigitaloceanDoksImage = "digitalocean/doks-debug:latest"
	ZinclabsUbuntuImage   = "public.ecr.aws/zinclabs/debug-ubuntu-base:latest"
	BusyboxImage          = "busybox:latest"
	WolfiBaseImage        = "cgr.dev/chainguard/wolfi-base:latest"
)

var Flags = map[string]cli.Flag{
	FlagRuntime: &cli.StringFlag{
		Name:    FlagRuntime,
		Value:   DockerRuntime,
		Usage:   FlagRuntimeUsage,
		EnvVars: []string{"DSLIM_DBG_RT"},
	},
	FlagTarget: &cli.StringFlag{
		Name:    FlagTarget,
		Value:   "",
		Usage:   FlagTargetUsage,
		EnvVars: []string{"DSLIM_DBG_TARGET"},
	},
	FlagNamespace: &cli.StringFlag{
		Name:    FlagNamespace,
		Value:   "default",
		Usage:   FlagNamespaceUsage,
		EnvVars: []string{"DSLIM_DBG_TARGET_NS"},
	},
	FlagPod: &cli.StringFlag{
		Name:    FlagPod,
		Value:   "",
		Usage:   FlagPodUsage,
		EnvVars: []string{"DSLIM_DBG_TARGET_POD"},
	},
	FlagDebugImage: &cli.StringFlag{
		Name:    FlagDebugImage,
		Value:   NicolakaNetshootImage,
		Usage:   FlagDebugImageUsage,
		EnvVars: []string{"DSLIM_DBG_IMAGE"},
	},
	FlagEntrypoint: &cli.StringFlag{
		Name:    FlagEntrypoint,
		Value:   "",
		Usage:   FlagEntrypointUsage,
		EnvVars: []string{"DSLIM_DBG_ENTRYPOINT"},
	},
	FlagCmd: &cli.StringFlag{
		Name:    FlagCmd,
		Value:   "",
		Usage:   FlagCmdUsage,
		EnvVars: []string{"DSLIM_DBG_CMD"},
	},
	FlagTerminal: &cli.BoolFlag{
		Name:    FlagTerminal,
		Value:   true, //attach interactive terminal by default
		Usage:   FlagTerminalUsage,
		EnvVars: []string{"DSLIM_DBG_TERMINAL"},
	},
	FlagListDebugImage: &cli.BoolFlag{
		Name:    FlagListDebugImage,
		Value:   false,
		Usage:   FlagListDebugImageUsage,
		EnvVars: []string{"DSLIM_DBG_LIST_IMAGES"},
	},
	FlagKubeconfig: &cli.StringFlag{
		Name:    FlagKubeconfig,
		Value:   "${HOME}/.kube/config",
		Usage:   FlagKubeconfigUsage,
		EnvVars: []string{"DSLIM_DBG_KUBECONFIG"},
	},
}

func cflag(name string) cli.Flag {
	cf, ok := Flags[name]
	if !ok {
		log.Fatalf("unknown flag='%s'", name)
	}

	return cf
}
