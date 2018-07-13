package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
//	"strconv"

	. "github.com/skelterjohn/go.matrix"
)

func errcheck(err error) {
	if err != nil {
		fmt.Printf("\nError:  %v occured", err)
	}
}

// Find the dot product between two vectors
func DotProduct(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return float64(0), errors.New("Cannot dot vectors of different length")
	}
	prod := float64(0)
	for i := 0; i < len(a); i++ {
		prod += a[i] * b[i]
	}
	return prod, nil
}

// For cosine similarity. Returns sqrt of sum of squared elements.
func NormSquared(a []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i] * a[i]
	}
	return math.Sqrt(sum)
}

// Cosine Similarity between two vectors
// Returns cos similarity on a scale from 0 to 1.
func CosineSim(a, b []float64) float64 {
	dp, err := DotProduct(a, b)
	errcheck(err)
	a_squared := NormSquared(a)
	b_sqaured := NormSquared(b)
	return dp / (a_squared * b_sqaured)
}

func multArray(a, b [] float64 ) float64 {
	sumMult := 0.0
	for i := 0; i < len(a); i++ {
		sumMult += a[i] * b[i]
	}
	return sumMult
}

func sumSquare (a [] float64) float64 {
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += (a[i] * a[i])
	}
	return sum
}
func PearsonСorrelationСoefficient(a, b [] float64)  float64{
	x := (float64(len(a)) * sumSquare(a) - (sum(a) * sum(a))) * (float64(len(a)) * sumSquare(b) - (sum(b) * sum(b)))
	res := (float64(len(a)) * multArray(a,b) - sum(a) * sum(b)) / math.Sqrt(x)
	return res
}
// defined as A n B / A u B. Used for binary user/product matrices.
func Jaccard(a, b []float64) float64 {
	intersection := float64(0)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			intersection += 1
		}
	}
	union := float64(0)
	for i := 0; i < len(a); i++ {
		if a[i] > 0 || b[i] > 0 {
			union += 1
		}
	}
	return intersection / union
}

func replaceNA(prefs *DenseMatrix) *DenseMatrix {
	arr := prefs.Array()
	for i := 0; i < len(arr); i++ {
		if math.IsNaN(arr[i]) {
			arr[i] = float64(0)
		}
	}
	return MakeDenseMatrix(arr, prefs.Rows(), prefs.Cols())
}

// Gets Recommendations for a user (row index) based on the prefs matrix.
// Uses cosine similarity for rating scale, and jaccard similarity if binary
func GetRecommendations(prefs *DenseMatrix, user int, products []string) ([] Recommendation, error) {
	// make sure user is in the preference matrix
	if user >= prefs.Rows() {
		return nil, errors.New("user index out of range")
	}
	//prefs = replaceNA(prefs)									// убираем неопределенности
	// item ratings
	ratings := make(map[int]float64, 0)
	sims := make(map[int]float64, 0)
	// Get user row from prefs matrix
	user_ratings := prefs.GetRowVector(user).Array()			// строка - товары пользователя
	var cntOfNeighbor int
	for i := 0; i < prefs.Rows(); i++ {
		// don't compare row to itself.
		if i != user {
			// get cosine similarity for other scores.
			other := prefs.GetRowVector(i).Array()				// other - товары другого пользователя
			//cos_sim := CosineSim(user_ratings, other)			// косинусная мера м-ду векторами [0,1]

			cos_sim := PearsonСorrelationСoefficient(user_ratings,other)

			// get product recs for neighbors
			if (cos_sim > 0) {
				for idx, val := range other { // проходим по товарам
					if (user_ratings[idx] == 0 || math.IsNaN(user_ratings[idx])) && val != 0 { // смотрим для некупленных товаров пользователя
						weighted_rating := val * cos_sim // задаем вес товара на значение вес соседа * косинусная мера
						ratings[idx] += weighted_rating
						sims[idx] += cos_sim // в симс лежит косинусная мера
					//	if (ratings[idx] > cos_sim) {
					//		fmt.Println(idx, " ", cos_sim, " ", ratings[idx], " ", sims[idx])
					//	}
					}
				}
				cntOfNeighbor++
			}
		}
	}
//	fmt.Println("Count of neighbors: ", cntOfNeighbor)
	recommendations := calculateWeightedMean(ratings, sims, products)
	return recommendations, nil
}

func sum(x []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
/*
// Gets Recommendations for a user (row index) based on the prefs matrix.
// Uses cosine similarity for rating scale, and jaccard similarity if binary
func GetBinaryRecommendations(prefs *DenseMatrix, user int, products []string) ([]string, []float64, error) {
	// make sure user is in the preference matrix
	if user >= prefs.Rows() {
		return nil, nil, errors.New("user index out of range")
	}
	prefs = replaceNA(prefs)
	// item ratings
	ratings := make(map[float64]string)
	// Get user row from prefs matrix
	user_ratings := prefs.GetRowVector(user).Array()

	for ii := 0; ii < prefs.Cols(); ii++ {
		if user_ratings[ii] == float64(0) {
			num_liked := sum(prefs.GetColVector(ii).Array())
			num_disliked := float64(prefs.Rows()) - num_liked
			jaccard_liked := make([]float64, 0)
			jaccard_disliked := make([]float64, 0)
			for i := 0; i < prefs.Rows(); i++ {
				if i != user {
					other := prefs.GetRowVector(i).Array()
					if other[ii] == float64(0) {
						jaccard_disliked = append(jaccard_disliked, Jaccard(user_ratings, other))
					} else {
						jaccard_liked = append(jaccard_liked, Jaccard(user_ratings, other))
					}
				}
			}
			rating := (sum(jaccard_liked) - sum(jaccard_disliked)) / (num_disliked + num_liked)
			if products != nil {
				ratings[rating] = products[ii]
			} else {
				ratings[rating] = strconv.Itoa(ii)
			}
		}
	}
	prods, scores := sortMap(ratings)
	return prods, scores, nil
}
*/

/*func calculateWeightedMean(ratings, sims map[int]float64, products []string) (recommends []string, values []float64) {
	recommendations := make(map[float64]string, 0)
	for k, v := range ratings {
		mean_product_rating := v / sims[k]			// ratings[k] / sims[k]
		if products != nil {
			recommendations[mean_product_rating] = products[k]
		} else {
			recommendations[mean_product_rating] = strconv.Itoa(k)
		}
	}
	recommends, values = sortMap(recommendations)
	return
}
*/

type Recommendation struct {
	product string
	mpRating float64
}

func calculateWeightedMean(ratings, sims map[int]float64, products []string) ([] Recommendation) {
	//recommendations := make(map[float64]string, 0)
	recommendations := make([] Recommendation, 0)
	for k, v := range ratings {
		mean_product_rating := v / sims[k]			// ratings[k] / sims[k]
		if products != nil {
			//recommendations[mean_product_rating] = products[k]
			recommendations = append(recommendations, Recommendation{
				product: products[k],
				mpRating: mean_product_rating,
			})
		} /*else {
			recommendations = append(recommendations, Recommendation{
				product: strconv.Itoa(k),
				mpRating: mean_product_rating,
			})
		}*/
	}
	//fmt.Println(len(recommendations))
	sort.Slice(recommendations, func (i, j int) bool { return recommendations[i].mpRating > recommendations[j].mpRating})
	/*for i := 0; i < len(recommendations); i++ {
		fmt.Println(recommendations[i].product, "-->", recommendations[i].mpRating)
	}*/
	return recommendations
}

func findRecByVal (recs [] Recommendation, val float64, prods [] string) {
	for i := 0; i < len(recs); i++ {
		if recs[i].mpRating == val {
			//fmt.Print(i)
			prods = append(prods, recs[i].product)
		}
	}
	//fmt.Println(len(prods))
}

/*
// Sorts a map of floats -> strings to get best recommendations. Probably a better way to do this.
func sortMap(recs map[float64]string) ([]string, []float64) {
	vals := make([]float64, 0)
	for k, _ := range recs {
		vals = append(vals, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	prods := make([]string, 0)
	for _, val := range vals {
		prods = append(prods, recs[val])
	}
	return prods, vals
}
*/


// Sorts a map of floats -> strings to get best recommendations. Probably a better way to do this.
func sortMap(recs [] Recommendation) ([]string, []float64) {
	vals := make([]float64, 0)
	for i:= range recs {
		vals = append(vals, recs[i].mpRating)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	prods := make([]string, 0)
	for _, val := range vals {
		//prods = append(prods, recs[val])
		findRecByVal(recs, val, prods)
	}
	//fmt.Println(prods)
	//fmt.Println(len(vals))
	return prods, vals
}


func MakeRatingMatrix(ratings []float64, rows, cols int) *DenseMatrix {
	return MakeDenseMatrix(ratings, rows, cols)
}