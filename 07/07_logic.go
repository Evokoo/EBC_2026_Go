package quest07

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I
// ========================
func I(clouds CloudMap, power int) int {
	DFS(clouds.FindInitialCloud(), power, &clouds)
	return clouds.CalcuateTotalEnergy()
}

// ========================
// PART III
// ========================
func III(clouds CloudMap, power, limit int) int {
	return 0
}

func DFS(id, power int, clouds *CloudMap) {
	cloud := (*clouds)[id]

	if power > 0 {
		cloud.charge += power

		for _, connection := range cloud.connections {
			DFS(connection, power-1, clouds)
		}
	}
}

// ========================
// PARSER
// ========================

type Cloud struct {
	id, energy, charge int
	connections        []int
}

type CloudMap map[int]*Cloud

func (cm CloudMap) FindInitialCloud() int {
	id, maxConnections := 1, len(cm[1].connections)

	for _id, cloud := range cm {
		if count := len(cloud.connections); count > maxConnections {
			id = _id
			maxConnections = count
		}
	}

	return id
}
func (cm CloudMap) CalcuateTotalEnergy() int {
	result := 0

	for _, cloud := range cm {
		result += cloud.charge + cloud.energy
	}

	return result
}

func parseInput(file string) CloudMap {
	data := utils.ReadFile(file)

	cloudMap := make(CloudMap)

	for line := range strings.SplitSeq(data, "\n") {
		cloud := &Cloud{}

		for i, n := range utils.QuickMatch(line, `\d+`) {
			num, _ := strconv.Atoi(n)
			switch i {
			case 0:
				cloud.id = num
			case 1:
				cloud.energy = num
			default:
				cloud.connections = append(cloud.connections, num)
			}
		}

		cloudMap[cloud.id] = cloud
	}
	return cloudMap
}
