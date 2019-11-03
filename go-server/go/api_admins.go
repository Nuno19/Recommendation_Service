/*
 * Recommendation API
 *
 * This is a recommendation API using k-means Clustering
 *
 * API version: 1.0.0
 * Contact: capela.nuno@ua.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var kmeans KMeans = KMeans{k: 3, maxIter: 100}
var tf_idf TF_IDF = TF_IDF{}
var lastIdx int
var cList []string
var stopWords = WordSet{"the", "a", "and", "you", "your", "he", "she", "his", "her", "\\N"}
var types []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func treatWord(word string) string {
	//word = porterstemmer.StemString(word)
	var re = regexp.MustCompile(`(tt\d{7}\/\d{4}(-\d)?|(ch|co|ev|nm|tt)\d{7})`)
	word = re.ReplaceAllString(word, "")
	reg := regexp.MustCompile("[^A-Za-zÀ-ÿ]+")
	word = reg.ReplaceAllString(word, "")
	word = strings.ToLower(word)

	return strings.ReplaceAll(word, "\"", "")
}

func LoadItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	corpus := r.FormValue("itemList")

	list := strings.Split(corpus, "\n")

	types = strings.Split(list[0], "\t")
	list = list[1:]

	cList = append(cList, list...)

	corpusSet := make([]WordSet, len(list))
	for i, l := range list {
		l = strings.ReplaceAll(l, "\t", " ")
		words := strings.Split(l, " ")
		set := WordSet{}
		for _, w := range words {
			w = treatWord(w)
			if !stopWords.exists(w) && w != "" && len(w) > 3 {
				set = append(set, w)
			}
		}
		corpusSet[i] = set
	}

	tf_idf.addToWordSet(corpusSet)
	fmt.Printf("WORD  SET LENGTH : %d\n", len(tf_idf.wordSet))
	tf_idf.sortMap()

	for _, val := range corpusSet {
		tf_idf.setCount(val)
		tf_idf.computeTF(val)
	}
	//fmt.Println(corpusSet[0])
	//fmt.Println(tf_idf.wordCountList[0])

	//tf_idf.computeIDF()

	//tf_idf.computeTFIDF()

	fitted := kmeans.fit(tf_idf.getAllPoints())
	if !fitted {
		http.Error(w, "Couldn't fit points in centroids!", http.StatusInternalServerError)
		return
	}
	dir, err := os.Create("files/")
	defer dir.Close()
	dir.Sync()

	s := "files/Items.txt"

	f, err := os.Create(s)
	check(err)

	for i, l := range cList {
		vals := strings.Split(l, "\t")
		m := map[string]string{}

		for j, lab := range types {
			m[lab] = vals[j]
		}

		js, err := json.Marshal(m)
		check(err)

		it := Item{Id: strconv.FormatInt(int64(i), 10), Data: string(js), BelongsTo: kmeans.labels[i]}

		t, err := json.Marshal(it)
		check(err)

		f.Write(t)
		f.WriteString("\n")

		f.Sync()
	}
	defer f.Close()
	lastIdx += len(list)
	sse := kmeans.computeSSE()

	fmt.Printf("LabelCount: %s Error: %f\n", fmt.Sprint(kmeans.labelCount()), sse)
	item := map[string]string{"Info": fmt.Sprintf("Uploaded %d items now there is %d total items! With error: %f", len(list), lastIdx, sse)}

	js, err := json.Marshal(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}

func LoadItemList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	Kerr := math.MaxFloat64
	labels := make([]int, len(kmeans.points))

	for i := 0; i < 10; i++ {
		fitted := kmeans.fit(tf_idf.getAllPoints())
		if !fitted {
			http.Error(w, "Couldn't fit points in centroids!", http.StatusInternalServerError)
			return
		}
		sse := kmeans.computeSSE()
		if sse < Kerr {
			labels = kmeans.labels
			Kerr = sse
			fmt.Printf("LabelCount: %s ERROR: %f\n", fmt.Sprint(labelCount(labels, kmeans.k)), Kerr)
		}
	}
	dir, err := os.Create("files/")
	defer dir.Close()
	dir.Sync()

	s := "files/Items.txt"

	f, err := os.Create(s)
	check(err)

	for i, l := range cList {
		vals := strings.Split(l, "\t")
		m := map[string]string{}

		for j, lab := range types {
			m[lab] = vals[j]
		}

		js, err := json.Marshal(m)
		check(err)

		it := Item{Id: strconv.FormatInt(int64(i), 10), Data: string(js), BelongsTo: labels[i]}

		t, err := json.Marshal(it)
		check(err)

		f.Write(t)
		f.WriteString("\n")

		f.Sync()
	}

	defer f.Close()
	item := map[string]string{"Info": fmt.Sprintf("Recalculated Kmeans with error %f", Kerr)}

	js, err := json.Marshal(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}

func SetClusterCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	data := r.FormValue("k")
	val, err := strconv.Atoi(data)
	if err != nil {
		http.Error(w, "Invalid k value!", http.StatusInternalServerError)
		return
	}
	kmeans.k = val
	kmeans.maxIter = 100
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Value of K updated to "+data)
}