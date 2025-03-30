package lvl
import (
	"main/k8s"
	"math/rand"
	"fmt"
	restclient "k8s.io/client-go/rest"
)

func (l *LvlMap) FillLvl(namespace string, kubeConfig *restclient.Config) {
	entitys := k8s.GetDeploy(namespace, kubeConfig)
	i,j := 3, 3
	for _, entity := range entitys.Items {
		l.Grid[i][j].IsPlaced = true
		l.Grid[i][j].Name = entity.Name
		switch  entity.Labels["plant"] {
			case "apple":
				l.Grid[i][j].EntityTextureId = 1
			case "wheat":
				l.Grid[i][j].EntityTextureId = 2
			default: 
				l.Grid[i][j].EntityTextureId = 3
		}
		i += 2 
		if i > (len(l.Grid) - 1)  {
			i = 3
			j += 2
		}
	}
}

func (l *LvlMap) Plant(i, j int, plant, namespace string, kubeConfig *restclient.Config) {
	l.Grid[i][j].IsPlaced = true
	l.Grid[i][j].Name = plant +	"-" + fmt.Sprintf("%d", rand.Int())

	println(l.Grid[i][j].Name)
	switch  plant {
	case "apple":
		l.Grid[i][j].EntityTextureId = 1
	case "wheat":
		l.Grid[i][j].EntityTextureId = 2
	}
	k8s.CreateDeploy(namespace, l.Grid[i][j].Name, plant, kubeConfig)
}

func (l *LvlMap) Break(i, j int, namespace string, kubeConfig *restclient.Config) {
	k8s.DeleteDeploy(namespace, l.Grid[i][j].Name, kubeConfig)
	l.Grid[i][j].EntityTextureId = 0
	l.Grid[i][j].Name = ""
}
