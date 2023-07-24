package provaplugin

import (
	"context"
	//"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/util/sets"
	//"k8s.io/apimachinery/pkg/util/wait"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/klog/v2"
	//"k8s.io/kubernetes/pkg/scheduler"
	framework "k8s.io/kubernetes/pkg/scheduler/framework"
  //"sigs.k8s.io/scheduler-plugins/apis/config"
)

const (
	Name = "ProvaPlugin"
)

type ProvaPlugin struct {}

var _ framework.ScorePlugin = &ProvaPlugin{}


func (p *ProvaPlugin) Name() string {
	return Name
}

func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &ProvaPlugin{}, nil
}

func (p *ProvaPlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeScore := int64(0)

	if nodeName == "node-1" {
		nodeScore = 10
	} else if nodeName == "node-2" {
		nodeScore = 5
	} else if nodeName == "node-3" {
		nodeScore = 2
	}

	return nodeScore, nil
}


func (p *ProvaPlugin) ScoreExtensions() framework.ScoreExtensions {
	return nil
}


/*func (p *ProvaPlugin) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	return nil
}*/
