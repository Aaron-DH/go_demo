package app

import (
	//"fmt"
	"github.com/spf13/pflag"
)

type KubeletFlags struct {
    KubeConfig          string
    BootstrapKubeconfig string

    // Insert a probability of random errors during calls to the master.
    ChaosChance float64

    // enableServer enables the Kubelet's server
    EnableServer bool
}


func NewKubeletFlags() *KubeletFlags {
	return &KubeletFlags{
		KubeConfig: "/etc/kubernetes/config",
		EnableServer: true,
	}
}

func (f *KubeletFlags) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&f.KubeConfig, "kubeconfig", f.KubeConfig, "Path to a kubeconfig file, specifying how to connect to the API server. P    roviding --kubeconfig enables API server mode, omitting --kubeconfig enables standalone mode.")

	fs.StringVar(&f.BootstrapKubeconfig, "bootstrap-kubeconfig", f.BootstrapKubeconfig, "Path to a kubeconfig file that will be used t    o get client certificate for kubelet. ")

	fs.Float64Var(&f.ChaosChance, "chaos-chance", f.ChaosChance, "If > 0.0, introduce random client errors and latency. Intended for t    esting.")
	fs.BoolVar(&f.EnableServer, "enable-server", f.EnableServer, "Enable the Kubelet's server")
}


