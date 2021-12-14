package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"sort"
)

//Need to change to Rune for unicode + figure out ranges for unicode characters
//Likely need a isnotsymbol function instead due to ranges
//Check if byte is an alpha
func isAlphaChar(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isPunctuation(b byte) bool {
	return b == '.' || b == ',' || b == '?' || b == ';' || b == '!' || b == ':' || b == '\'' || b == '"' ||b == ']'||b == '['||b == '{'||b == '}'
}

//Need to simplify
//Function to clean up words (removing non-alpha characters
func parseText(word string) []string {
	
	var temp = []byte(word)
	var text []byte
	var result []string
	
	isWord := false //Verify if iterating over a word, else do not append characters
	
	//Loop through word, parsing characters, creating an string array
	for k := range temp {
		
		if !isPunctuation(temp[k]) {
			
			if isWord == false {
				isWord = true
			}
			
			text = append(text, temp[k])
			
		} else {
			
			if isWord == true {
				
				//In case of contractions (shouldn't, can't, etc.)
				if temp[k] == '\'' && k+1 < len(temp) {
				
					//Ensure next character is an aplha
					if isAlphaChar(temp[k+1]) { 
						text = append(text, temp[k])
						continue
					}
				}
				
				isWord = false
				result = append(result, string(text))
				
				//reset byte array, until word is found
				text = []byte{}
			}
		}
	}
	
	//append last string object, if word is found 
	if len(text) > 0 {
		result = append(result, string(text))
	}
	
	return result
}


//Function to build phrase object
func updatePhraseBuffer(text []string, phrase []string) []string{
	
	for k := range text {
		phrase = append(phrase,text[k])
	}
	return phrase
}

//Function generate phrase of N depth of words
func buildPhrase(phrase []string, depth int) string {

	if depth == 1 {
		return phrase[0]
	}
	return phrase[0] + " " + buildPhrase(phrase[1:], depth-1)
}

//Function to trim buffer
func reduceSlice(s []string, depth int) []string {

	if len(s) == depth {
		return s[1:]
	} else {
		return reduceSlice(s[1:], depth)
	}
}


//Function to update Stats map of parsed data
func updateMap (phrase []string, table map[string]int, depth int) map[string]int {
	
	var text string;
	
	for k := range phrase {
	
		text = strings.ToLower(buildPhrase(phrase[k:], depth))
		table[text] += 1 
		
		if len(phrase[k:]) == depth {
			break
		}
	}
	return table
}

func reverseMap(m map[string]int) map[int][]string {
	n := make(map[int][]string, len(m))
	for k, v := range m {
		if len(n[v]) == 0 {
			n[v] = []string{k}
		} else {
			n[v] = append(n[v], k)
		}
	}
	return n
}

func printMap(m map[int][]string, limit int) {
	var keys []int
	var reverse []int

	for key := range m {
		keys = append(keys, key)
	}

	reverse = reverseSort(keys)
	
	l := limit
	
	for _, key := range reverse {
		
		if limit == 0 {
			break;
		}
		
		//Print map data. 
		if len(m[key]) > 1 {
			
			for n := range m[key] {
				if l == 0 {
					break;
				}
				fmt.Println("Phrase :", m[key][n], " Count :", key)
				l--
			}
		} else {
			fmt.Println("Phrase :", m[key][0], " - Count :", key)
			l--
		}
	}
	//get object count 
	//fmt.Println(l)
}

//Need to implement desc mrege sort for Performance
//Reverse sort the data to print 
func reverseSort(keys []int) []int {
	reverse := make([]int, len(keys))
	sort.Ints(keys)

	n := len(keys)
	for n > 0 {
		reverse[len(keys)-n] = keys[n-1]
		n--
	}

	return reverse
}

func main() {
	
	//Get Arguments
	args := os.Args[1:]
	
	var buffer []string
	var stats = map[string]int{}
	var depth = 3
	var statSize = 100
	
	stdinObj := os.Stdin
	stdinStat, err := stdinObj.Stat()
	if err != nil {
		panic(err)
	}
	
	//Check if Size of stdin has data, or ensure there is data piped from the terminal (per documenation)
	if stdinStat.Size() > 0||stdinStat.Mode()&os.ModeCharDevice == 0 {
	
		//Using scanner, which breaks text into space delimited text blocks
		scanner := bufio.NewScanner(stdinObj)
		scanner.Split(bufio.ScanWords)
		
		//repeative code, should put into function
		for scanner.Scan() {
		
			var word = parseText(scanner.Text())
			buffer = updatePhraseBuffer(word, buffer)
			
			//Check buffer
			//fmt.Println(buffer)
			if len(buffer) >= depth{
				stats = updateMap(buffer, stats,depth)
				buffer = reduceSlice(buffer,depth)
			}
			
			//Verifying buffer count
			//fmt.Println(buffer)
		}
	}
	
	//Assumption: merge text of all files
	for k := range args{
			
		f, err := os.Open(args[k])
			
		if (err != nil) {
			panic(err)
		}
				
		//Using scanner, which breaks text into space delimited text blocks
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		
		//repeative code, should put into function		
		for scanner.Scan() {
			
			//Get 
			var word = parseText(scanner.Text())
			buffer = updatePhraseBuffer(word, buffer)
			
			//fmt.Println(buffer)
			if len(buffer) >= depth{
				stats = updateMap(buffer, stats,depth)
				buffer = reduceSlice(buffer,depth)
			}
			
			//Verifying buffer count
			//fmt.Println(buffer)
		}
		f.Close()
	}
	
	var reverseStats = reverseMap(stats)
	printMap(reverseStats,statSize)
}