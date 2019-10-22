package swagger

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type WordSet []string

type WordCounts map[string]int

type FloatMap map[string]float64

type TF_IDF struct {
	wordSet       WordSet
	wordCountList []WordCounts
	tf            []FloatMap
	idf           FloatMap
	tfIdf         []FloatMap
}

func (slice WordSet) print() {
	for _, word := range slice {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}

func (fmap FloatMap) print() {
	for k, val := range fmap {
		fmt.Printf("MAP[%s]=%f\n", k, val)
	}
}

func (fmap WordCounts) print() {
	for k, val := range fmap {
		fmt.Printf("[%s]: %d  |  ", k, val)
	}
	fmt.Println()
}
func (slice WordSet) exists(word string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}

func (slice WordSet) boolExists(word string) bool {
	sort.Strings(slice)
	i := sort.SearchStrings(slice, word)
	if i < len(slice) && slice[i] == word {
		return true
	}
	return false
}

func (slice WordSet) toLower() WordSet {
	ret := make(WordSet, len(slice))
	for i, word := range slice {
		ret[i] = strings.ToLower(word)
	}
	return ret
}

func initCounts(set WordSet) WordCounts {
	cout := make(WordCounts)
	for _, word := range set {
		cout[word] = 0
	}
	return cout
}

func (tf_idf *TF_IDF) setCount(corpus []WordSet) {
	tf_idf.wordCountList = make([]WordCounts, len(corpus))
	for i, set := range corpus {
		tf_idf.wordCountList[i] = initCounts(tf_idf.wordSet)
		for _, word := range set {
			tf_idf.wordCountList[i][word]++
		}
	}
	fmt.Print()
}

func (tf_idf *TF_IDF) computeTF(corpus WordSet, idx int) {
	if len(tf_idf.tf) == 0 {
		tf_idf.tf = []FloatMap{}
	}
	tf_idf.tf = append(tf_idf.tf, make(FloatMap))
	for key, count := range tf_idf.wordCountList[idx] {
		tf_idf.tf[idx][key] = float64(count) / float64(len(corpus))
	}

}

func (tf_idf *TF_IDF) computeIDF() {
	n := len(tf_idf.wordCountList)

	idf := make(FloatMap)
	for _, list := range tf_idf.wordCountList {

		for key, count := range list {
			if count > 0 {
				idf[key]++
			}
		}
	}

	for key, count := range idf {
		idf[key] = math.Log(float64(n) / float64(count))
	}
	tf_idf.idf = idf
	fmt.Println("")
}

func (tf_idf *TF_IDF) computeTFIDF() {
	tf_idf.tfIdf = make([]FloatMap, len(tf_idf.wordCountList))

	for idx, mapIF := range tf_idf.tf {
		tf_idf.tfIdf[idx] = make(FloatMap)
		for key, val := range mapIF {
			tf_idf.tfIdf[idx][key] = val * tf_idf.idf[key]
		}
	}

}

func (tf_idf *TF_IDF) addToWordSet(corpus []WordSet) {
	for _, list := range corpus {
		list = list.toLower()
		for _, word := range list {
			if !tf_idf.wordSet.exists(word) {
				tf_idf.wordSet = append(tf_idf.wordSet, word)
			}
		}
	}
}

func (tf_idf *TF_IDF) sortMap() {

	sort.Strings(tf_idf.wordSet)
}

func (tf_idf *TF_IDF) getPointByIndex(idx int) Point {
	coord := make(Point, len(tf_idf.wordSet))
	for i, key := range tf_idf.wordSet {
		value := (tf_idf.tf[idx][key])
		coord[i] = value
	}
	return coord
}

func (tf_idf *TF_IDF) getAllPoints() []Point {
	pointArr := make([]Point, len(tf_idf.tf))
	for i := range tf_idf.tf {
		pointArr[i] = tf_idf.getPointByIndex(i)
	}
	return pointArr
}

func (tf_idf *TF_IDF) normalize() {
	mapF := initCounts(tf_idf.wordSet)
	for i := range tf_idf.tfIdf {
		for _, key := range tf_idf.wordSet {
			if tf_idf.tfIdf[i][key] != 0 {
				mapF[key]++
			}
		}
	}
	for k, val := range mapF {
		if val == 0 {
			fmt.Println("ZEROOOO")
			for i := range tf_idf.tfIdf {
				delete(tf_idf.tfIdf[i], k)
			}
		}
	}
}

/*
func main() {
	sum1 := 0.0
	sum2 := 0.0
	sse := make([]float64, 35)
	rand.Seed(time.Now().Unix())
	for I := 0; I < 35; I++ {
		dat, _ := ioutil.ReadFile("final.tsv")
		corpus := string(dat)
		//corpus := "Simple example with Cats and Mouse\nAnother simple example with dogs and cats\nAnother simple example with mouse and cheese\ntf and idf is awesome\nsome androids is there"
		//corpus := "Google and Facebook are strangling the free press to death. Democracy is the loserGoogle an \nYour 60-second guide to security stuff Google touted today at Next '18\nA Guide to Using Android Without Selling Your Soul to Google\nReview: Lenovo’s Google Smart Display is pretty and intelligent\nGoogle Maps user spots mysterious object submerged off the coast of Greece - and no-one knows what it is\nAndroid is better than IOS\nIn information retrieval, tf–idf or TFIDF, short for term frequency–inverse document frequency\nis a numerical statistic that is intended to reflect\nhow important a word is to a document in a collection or corpus.\nIt is often used as a weighting factor in searches of information retrieval\ntext mining, and user modeling. The tf-idf value increases proportionally\nto the number of times a word appears in the document\nand is offset by the frequency of the word in the corpus"
		//fmt.Println(corpus)
		list := strings.Split(corpus, "\n")

		corpusSet := make([]WordSet, len(list))
		for i, l := range list {
			corpusSet[i] = WordSet(strings.Split(l, " ")).toLower()
		}
		//fmt.Println("START")

		tf_idf := TF_IDF{}

		tf_idf.addToWordSet(corpusSet)
		//fmt.Println("COUNT")
		tf_idf.sortMap()
		tf_idf.setCount(corpusSet)
		//fmt.Println("TF")

		for idx, val := range corpusSet {
			tf_idf.computeTF(val, idx)
		}

		//fmt.Println("IDF")
		tf_idf.computeIDF()

		//fmt.Println("TF_IDF")
		tf_idf.computeTFIDF()

		//tf_idf.tfIdf[1].print()

		pointArr := make([]Point, len(tf_idf.tf))
		for i := range tf_idf.tf {
			pointArr[i] = tf_idf.getPoint(i)
		}

		//fmt.Println("FINISHED")

		kmeans := KMeans{k: I + 40, maxIter: 200}

		kmeans.fit(pointArr)
		a := kmeans.labelCount()
		printList(a)
		s := kmeans.computeSSE()
		fmt.Printf("ERROR: %f\n", s)
		sse[I] = s
	}
	fmt.Printf("MEANS: %f,%f", sum1/15.0, sum2/15.0)
}
*/
