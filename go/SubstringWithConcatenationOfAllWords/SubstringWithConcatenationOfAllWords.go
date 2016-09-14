/*
30. Substring with Concatenation of All Words

You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).
*/
package main

import (
	"fmt"
)

func wordsCombination(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	if len(words) == 1 {
		return []string{words[0]}
	}
	combination := []string{}
	for i := 0; i < len(words); i++ {
		new_words := make([]string, 0, len(words)-1)
		new_words = append(new_words, words[:i]...)
		new_words = append(new_words, words[i+1:]...)
		for _, v := range wordsCombination(new_words) {
			combination = append(combination, words[i]+v)
		}
	}
	return combination
}

func wordsCombinationIndexs(num int) [][]int {
	if num <= 0 {
		return [][]int{}
	}
	if num == 1 {
		return [][]int{{0}}
	}
	solution := [][]int{}
	for i := 0; i < num; i++ {
		pres := wordsCombinationIndexs(num - 1)
		for _, digits := range pres {
			for k := 0; k < len(digits); k++ {
				if digits[k] >= i {
					digits[k]++
				}
			}
			solution = append(solution, append(digits, i))
		}
	}
	return solution
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	wordLen := len(words[0])
	wordsLen := len(words)
	solution := []int{}
strRange:
	for i := 0; i < (len(s) - (wordLen * wordsLen) + 1); i++ {
		wordsUnuse := map[string]int{}
		for _, w := range words {
			wordsUnuse[w]++
		}
		wn := 0
	wordListRange:
		for ; wn < wordsLen; wn++ {
			for w, _ := range wordsUnuse {
				j := 0
				for ; j < wordLen && w[j] == s[i+j+wordLen*wn]; j++ {

				}
				if j == wordLen {
					wordsUnuse[w]--
					if wordsUnuse[w] == 0 {
						delete(wordsUnuse, w)
					}
					continue wordListRange
				}
			}
			continue strRange
		}
		solution = append(solution, i)
	}
	return solution
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("aaa", []string{"a", "a"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("aaaaaaaa", []string{"aa", "aa", "aa"}))
	fmt.Println(findSubstring("pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel",
		[]string{"dhvf", "sind", "ffsl", "yekr", "zwzq", "kpeo", "cila", "tfty", "modg", "ztjg", "ybty", "heqg", "cpwo", "gdcj", "lnle", "sefg", "vimw", "bxcb"}))

}
