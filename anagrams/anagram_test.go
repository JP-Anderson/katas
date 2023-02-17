package anagrams

import (
	"fmt"
	"testing"
)

func TestWordToAnagramKey(t *testing.T) {
	key, ok := wordToAnagramKey("abcdefghijklmnopqrstuvwxyz")
	for _, letterCount := range key.letters {
		if letterCount != 1 {
			t.Errorf("Value not 1")
		}
	}
	if !ok {
		t.Errorf("got false bool from wordToAnagramKey!")
	}
}

func TestWordToAnagramKeyPunctuation(t *testing.T) {
	_, ok := wordToAnagramKey("what!")
	if ok {
		t.Errorf("wordToAnagram key returned true but expected false")
	}
}

func TestAnagrams(t *testing.T) {
	result := AnagramsInFile("test1.txt")
	expected := [][]string{
		{ "sunders", "undress" },
		{ "pinkish", "kinship" },
		{ "cheese" },
		{ "leek" },
	}
	if len(result) != len(expected) {
		t.Errorf("Expected list to be of length %d but was of length %d", len(expected), len(result))
		t.Errorf("RESULT %v", result)
	}
	for i, group := range result {
		if len(group) != len(expected[i]) {
			t.Errorf("Group lengths did not match %v != %v", group, expected[i])
		}
		for j, word := range group {
			if word != expected[i][j] {
				t.Errorf("Words did not match %s != %s", word, expected[i][j])	
			}
		}
	} 
}

func TestAnagramsLargeFile(t *testing.T) {
	result := AnagramsInFile("wordlist.txt")
	fmt.Printf("There are %d anagram groups\n", len(result))
}
