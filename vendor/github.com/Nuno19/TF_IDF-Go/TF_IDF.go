//Package tfidf Provides a simple implementation of Term-Frequency Inverse-Document-Frequency in golang
//Can return the points in tf values or tf-idf values
package tfidf

import (
	"fmt"
	"math"
	"sort"
	"strings"

	kmeans "github.com/Nuno19/KMeans-Go"
)

//WordSet - Set of Words
type WordSet []string

//WordCounts - Dictionary of word by count
type WordCounts map[string]int

//FloatMap - Dictionary of word by float for tf or idf
type FloatMap map[string]float64

//TF_IDF - Struct for calculation TF-IDF
type TF_IDF struct {
	//List of words
	SetWord WordSet
	//Count of each word by set
	WordCountList []WordCounts
	//Term-Frequency
	Tf []FloatMap
	//Inverse Document Frequency
	Idf FloatMap
	//Term-Frequency Inverse Document Frequency
	TfIdf []FloatMap
}

//Print - prints the set of words
func (slice WordSet) Print() {
	for _, word := range slice {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}

//Print - prints the float values of the map
func (fmap FloatMap) Print() {
	for k, val := range fmap {
		fmt.Printf("MAP[%s]=%f\n", k, val)
	}
}

//Print - Prints the counts of the words
func (fmap WordCounts) Print() {
	for k, val := range fmap {
		fmt.Printf("[%s]: %d  |  ", k, val)
	}
	fmt.Println()
}

//Exists - checks if a word exists in a set
func (slice WordSet) Exists(word string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}

//ToLower - returns the words in the set in lower case
func (slice WordSet) ToLower() WordSet {
	ret := make(WordSet, len(slice))
	for i, word := range slice {
		ret[i] = strings.ToLower(word)
	}
	return ret
}

//InitCounts - returns a map of WordCounts with the set of words
func InitCounts(set WordSet) WordCounts {
	cout := make(WordCounts)
	for _, word := range set {
		cout[word] = 0
	}
	return cout
}

//SetCount - sets the counts of words for the set
func (tf_idf *TF_IDF) SetCount(corpus WordSet) {
	if len(tf_idf.WordCountList) == 0 {
		tf_idf.WordCountList = []WordCounts{}
	}
	tf_idf.WordCountList = append(tf_idf.WordCountList, InitCounts(tf_idf.SetWord))
	idx := len(tf_idf.WordCountList) - 1
	for _, word := range corpus {
		tf_idf.WordCountList[idx][word]++
	}
}

//SetCountIdx - sets the counts of words for the set in that idx
func (tf_idf *TF_IDF) SetCountIdx(corpus WordSet, idx int) {
	if len(tf_idf.WordCountList) == 0 {
		tf_idf.WordCountList = []WordCounts{}
	}
	if len(tf_idf.WordCountList) == idx {
		tf_idf.WordCountList = append(tf_idf.WordCountList, InitCounts(tf_idf.SetWord))
	} else {
		tf_idf.WordCountList[idx] = InitCounts(tf_idf.SetWord)
	}
	for _, word := range corpus {
		tf_idf.WordCountList[idx][word]++
	}
}

//ComputeTF - computes the term frequency for the word set
func (tf_idf *TF_IDF) ComputeTF(corpus WordSet) {
	idx := len(tf_idf.Tf)
	if len(tf_idf.Tf) == 0 {
		tf_idf.Tf = []FloatMap{}
	}
	tf_idf.Tf = append(tf_idf.Tf, make(FloatMap))
	for key, count := range tf_idf.WordCountList[idx] {
		tf_idf.Tf[idx][key] = float64(count) / float64(len(corpus))
	}

}

//ComputeTFIdx - computes the term frequency for the word set in that idx
func (tf_idf *TF_IDF) ComputeTFIdx(corpus WordSet, idx int) {
	if len(tf_idf.Tf) == 0 {
		tf_idf.Tf = []FloatMap{}
	}
	if len(tf_idf.Tf) == idx {
		tf_idf.Tf = append(tf_idf.Tf, make(FloatMap))
	}
	for key, count := range tf_idf.WordCountList[idx] {
		tf_idf.Tf[idx][key] = float64(count) / float64(len(corpus))
	}

}

//ComputeIDF - computes the inverse document frequency for the list of sets
func (tf_idf *TF_IDF) ComputeIDF() {
	n := len(tf_idf.WordCountList)

	idf := make(FloatMap)
	for _, list := range tf_idf.WordCountList {

		for key, count := range list {
			if count > 0 {
				idf[key]++
			}
		}
	}

	for key, count := range idf {
		idf[key] = math.Log(float64(n) / float64(count))
	}
	tf_idf.Idf = idf
}

//ComputeTFIDF - computes the tf-idf for the each set of words
func (tf_idf *TF_IDF) ComputeTFIDF() {
	tf_idf.TfIdf = make([]FloatMap, len(tf_idf.WordCountList))

	for idx, mapIF := range tf_idf.Tf {
		tf_idf.TfIdf[idx] = make(FloatMap)
		for key, val := range mapIF {
			tf_idf.TfIdf[idx][key] = val * tf_idf.Idf[key]
		}
	}

}

//AddToWordSet - adds the set to the list of sets
func (tf_idf *TF_IDF) AddToWordSet(corpus []WordSet) {
	for _, list := range corpus {
		list = list.ToLower()
		for _, word := range list {
			if !tf_idf.SetWord.Exists(word) {
				tf_idf.SetWord = append(tf_idf.SetWord, word)
			}
		}
	}
}

//SortMap - sorts the set by key values
func (tf_idf *TF_IDF) SortMap() {
	sort.Strings(tf_idf.SetWord)
}

//GetPointByIndexTF - gets tf map of the point in the indes
func (tf_idf *TF_IDF) GetPointByIndexTF(idx int) kmeans.Point {
	coord := make(kmeans.Point, len(tf_idf.SetWord))
	for i, key := range tf_idf.SetWord {
		value := (tf_idf.Tf[idx][key])
		coord[i] = value
	}
	return coord
}

//GetAllPointsTF - gets all the points in the list
func (tf_idf *TF_IDF) GetAllPointsTF() []kmeans.Point {
	pointArr := make([]kmeans.Point, len(tf_idf.Tf))
	for i := range tf_idf.Tf {
		pointArr[i] = tf_idf.GetPointByIndexTF(i)
	}
	return pointArr
}

//GetPointByIndexTFIDF - gets tf map of the point in the indes
func (tf_idf *TF_IDF) GetPointByIndexTFIDF(idx int) kmeans.Point {
	coord := make(kmeans.Point, len(tf_idf.SetWord))
	for i, key := range tf_idf.SetWord {
		value := (tf_idf.Tf[idx][key])
		coord[i] = value
	}
	return coord
}

//GetAllPointsTFIDF - gets all the points in the list
func (tf_idf *TF_IDF) GetAllPointsTFIDF() []kmeans.Point {
	pointArr := make([]kmeans.Point, len(tf_idf.Tf))
	for i := range tf_idf.Tf {
		pointArr[i] = tf_idf.GetPointByIndexTFIDF(i)
	}
	return pointArr
}

//GetIDF - gets all the points in the list
func (tf_idf *TF_IDF) GetIDF() FloatMap {
	return tf_idf.Idf
}
