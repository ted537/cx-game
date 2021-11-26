package cxmath

type IdxPair struct { First,Second int }

// should probably optimize this at some point
func FindClusters(points []Vec2i, radius int) [][]Vec2i {
	connections := findConnections(points,radius)
	clusters := collapseConnectionsToClusters(points,connections)
	return clusters
}

// flood-fill variant
func findConnections(points []Vec2i, radius int) map[IdxPair]bool {
	lookup := map[Vec2i]bool{}
	for _,point := range points { lookup[point] = true }

	connections := map[IdxPair]bool{}
	for pointIdx,point := range points {
		for neighbourIdx,neighbour := range neighbours(point,radius) {
			exists := lookup[neighbour]
			if exists {
				idxPair := IdxPair { pointIdx, neighbourIdx }
				connections[idxPair] = true
			}
		}
	}
	return connections
}

func neighbours(point Vec2i, radius int) []Vec2i {
	diameter := radius*2+1
	count := diameter*diameter-1
	neighbours := make([]Vec2i, 0, count)
	for x := int(point.X) - radius ; x <= int(point.X) + radius ; x++ {
		for y := int(point.Y) - radius ; y <= int(point.Y) + radius ; y++ {
			if x!=int(point.X) || y!=int(point.Y) {
				neighbours = append(neighbours, Vec2i {int32(x),int32(y)})
			}
		}
	}
	return neighbours
}

// graph traversal
func collapseConnectionsToClusters(
	points []Vec2i, connections map[IdxPair]bool,
) [][]Vec2i {
	// to start, assign a unique cluster ID to each point
	clusterIDs := make([]int,len(points))
	for idx := range clusterIDs { clusterIDs[idx] = idx }

	clusters := map[int]*[]Vec2i{}

	for connection,_ := range connections {
		firstClusterID := clusterIDs[connection.First]
		secondClusterID := clusterIDs[connection.Second]

		// point both existing clusters to (possibly) new slice
		newPoints :=
			append(*clusters[firstClusterID], *clusters[secondClusterID]...)
		clusters[firstClusterID] = &newPoints
		clusters[secondClusterID] = &newPoints
	}

	uniqueSlices := findUniqueSlices(clusters)
	return uniqueSlices
}

func findUniqueSlices(clusters map[int]*[]Vec2i) [][]Vec2i {
	seen := map[*[]Vec2i]bool{}
	slices := [][]Vec2i{}
	for _,cluster := range clusters {
		if !seen[cluster] {
			slices = append(slices, *cluster)
		}
		seen[cluster] = true
	}
	return slices
}
