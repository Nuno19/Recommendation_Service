package kmeans

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

//KMeans - Struct for KMeans Algorithm
type KMeans struct {
	//number of clusters
	K int
	//max number of iterations
	MaxIter int
	//Centroids of the clustering
	Centroids []Point
	//Points for clustering
	Points []Point
	//Labels of the corresponding point
	Labels []int
}

//InitCentroids - Inits the Centroids for random
func (kmeans *KMeans) InitCentroids() bool {
	if len(kmeans.Points) <= kmeans.K {
		return false
	}
	kmeans.Centroids = make([]Point, kmeans.K)
	var perm = rand.Perm(len(kmeans.Points))
	for i := 0; i < kmeans.K; i++ {
		kmeans.Centroids[i] = kmeans.Points[perm[i]]
	}
	return true
}

//ComputeSSE - computes Sum of Squared Erros
func (kmeans *KMeans) ComputeSSE() float64 {
	var wg1 sync.WaitGroup
	wg1.Add(len(kmeans.Centroids))
	distances := make(Point, len(kmeans.Centroids))
	for i := range kmeans.Centroids {
		go func(i int) {
			for j, point := range kmeans.Points {
				if kmeans.Labels[j] == i {
					distances[i] += point.Subtract(kmeans.Centroids[i]).Norm()
				}
			}
			defer wg1.Done()
		}(i)
	}
	wg1.Wait()
	distance := 0.0
	for _, dist := range distances {
		distance += math.Pow(dist, 2)
	}
	return distance
}

//ComputeClosestCentroidIdx - returns the index of the closest centroid for that point
func (kmeans *KMeans) ComputeClosestCentroidIdx(point Point) int {

	min := math.MaxFloat64
	minIdx := -1
	for i, cent := range kmeans.Centroids {
		distance := point.PointDist(cent)
		if distance < min {
			min = distance
			minIdx = i
		}
	}
	return minIdx
}

//ComputeLabels - Computes Labels of the Points
func (kmeans *KMeans) ComputeLabels() []int {
	var Labels = make([]int, len(kmeans.Points))
	var wg1 sync.WaitGroup
	wg1.Add(len(kmeans.Points))

	for i := range kmeans.Points {
		go func(i int) {
			min := math.MaxFloat64
			minIdx := -1
			for j, cent := range kmeans.Centroids {
				distance := kmeans.Points[i].PointDist(cent)
				if distance < min {
					min = distance
					minIdx = j
				}
			}

			Labels[i] = minIdx
			defer wg1.Done()
		}(i)
	}
	wg1.Wait()
	return Labels
}

//ComputeCentroids - Recalculates Centroids position
func ComputeCentroids(pointList []Point, distIndex []int, K int) []Point {
	clusters := make([][]Point, K)
	for i, point := range distIndex {
		clusters[point] = append(clusters[point], pointList[i])
	}

	Centroids := make([]Point, K)
	for c, clu := range clusters {
		sums := make([]float64, len(pointList[0]))
		for _, point := range clu {
			for j, val := range point {
				sums[j] += val
			}
		}
		Centroids[c] = make([]float64, len(sums))
		for i, s := range sums {
			Centroids[c][i] = s / float64(len(clu))
		}
	}
	return Centroids
}
func (kmeans *KMeans) update() bool {
	oldCentroids := make([]Point, len(kmeans.Centroids))
	for i := 0; i < kmeans.MaxIter; i++ {
		oldCentroids = kmeans.Centroids
		kmeans.Labels = kmeans.ComputeLabels()

		kmeans.Centroids = ComputeCentroids(kmeans.Points, kmeans.Labels, kmeans.K)

		if equals(kmeans.Centroids, oldCentroids) {

			return true
		}
	}

	return false
}

//Fit - Fits the Points into K Centroids
func (kmeans *KMeans) Fit(pointList []Point) bool {
	if kmeans.K == 0 {
		fmt.Printf("K not defined")
		return false
	}

	kmeans.Points = pointList
	init := kmeans.InitCentroids()
	if !init {
		return false
	}
	found := kmeans.update()
	if found {
		return true
	}
	return false

}

//AddPoint - Add a new Point to the list and recalculate Centroids
func (kmeans *KMeans) AddPoint(point Point) {
	kmeans.Points = append(kmeans.Points, point)
	found := kmeans.update()
	if found {
		fmt.Printf("Centroids Found!\n")
	} else {
		fmt.Printf("Centroids Not Found!\n")
	}
}

//AddPointList - Add a list of Points to the list and recalculate Centroids
func (kmeans *KMeans) AddPointList(Points []Point) {
	for _, p := range Points {
		kmeans.Points = append(kmeans.Points, p)
	}
	found := kmeans.update()
	if found {
		fmt.Printf("Centroids Found!\n")
	} else {
		fmt.Printf("Centroids Not Found!\n")
	}
}

//PointsToString - Returns the String of the Points
func (kmeans *KMeans) PointsToString() string {
	var toReturn string

	for i, point := range kmeans.Points {
		for j, p := range point {
			if j < len(point)-1 {
				toReturn += fmt.Sprintf("%f,", p)
			} else {
				toReturn += fmt.Sprintf("%f", p)
			}
		}
		if i < len(kmeans.Points)-1 {
			toReturn += fmt.Sprintf(";")
		}
	}
	return toReturn
}

//CentroidsToString - returns the string of Centroids
func (kmeans *KMeans) CentroidsToString() string {
	var toReturn string

	for i, cent := range kmeans.Centroids {
		for j, p := range cent {
			if j < len(cent)-1 {
				toReturn += fmt.Sprintf("%f,", p)
			} else {
				toReturn += fmt.Sprintf("%f", p)
			}
		}
		if i < len(kmeans.Centroids)-1 {
			toReturn += fmt.Sprintf(";")
		}
	}
	return toReturn
}

//LabelsToString - Returns the string of Labels
func (kmeans *KMeans) LabelsToString() string {
	var toReturn string

	for i, label := range kmeans.Labels {

		if i < len(kmeans.Labels)-1 {
			toReturn += fmt.Sprintf("%d,", label)
		} else {
			toReturn += fmt.Sprintf("%d", label)
		}
	}
	return toReturn
}

//LabelCount - Returns the number of Points per centroid
func (kmeans *KMeans) LabelCount() []int {
	count := make([]int, len(kmeans.Centroids))
	for _, label := range kmeans.Labels {
		count[label]++
	}
	return count
}

//LabelCount - Returns the number of Points per centroid
func LabelCount(Labels []int, K int) []int {
	count := make([]int, K)
	for _, label := range Labels {
		count[label]++
	}
	return count
}

//GetPointIdxOfCentroid - return indexes of Points of a certain centroid
func (kmeans *KMeans) GetPointIdxOfCentroid(centroidIdx int) []int {
	var idxs []int

	for i, label := range kmeans.Labels {
		if label == centroidIdx {
			idxs = append(idxs, i)
		}
	}
	return idxs
}
