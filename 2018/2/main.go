package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Part 1")
	fmt.Println(checkSum(oneP1))
	fmt.Println(checkSum(input))

	fmt.Println("Part 2")
	fmt.Println(idCheck(oneP2))
	fmt.Println(idCheck(input))
}

func checkSum(input string) int {
	lines := strings.Split(input, "\n")

	doubleCountTotal := 0
	tripleCountTotal := 0
	for _, line := range lines {
		doubleCount, tripleCount := check(line)
		if doubleCount {
			doubleCountTotal++
		}
		if tripleCount {
			tripleCountTotal++
		}
	}

	return doubleCountTotal * tripleCountTotal
}

// returns double count, triple count
func check(input string) (bool, bool) {
	counts := map[rune]int{}

	for _, val := range input {
		counts[val]++
	}

	doubleCount := 0
	tripleCount := 0
	for _, count := range counts {
		if count == 2 {
			doubleCount++
		} else if count == 3 {
			tripleCount++
		}
	}

	return doubleCount > 0, tripleCount > 0
}

func idCheck(input string) string {
	lines := strings.Split(input, "\n")

	var idOne string
	var idTwo string
	for i := 0; i < len(lines) && idOne == ""; i++ {
		one := lines[i]
		for j := 0; j < len(lines) && idTwo == ""; j++ {
			two := lines[j]
			difference := differenceCount(one, two)
			if difference == 1 {
				idOne = one
				idTwo = two
			}
		}
	}

	output := ""
	for i, val := range idOne {
		if val == rune(idTwo[i]) {
			output += string(val)
		}
	}

	return output
}

func differenceCount(id string, check string) int {
	difference := 0
	for i, letter := range id {
		if letter != rune(check[i]) {
			difference++
		}
	}

	return difference
}

var oneP1 = `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`

var oneP2 = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

var input = `prtkqyluibmtcwqaezjmhgfndx
prtkqylusbsmcwvaezjmhgfndt
prgkqyluibsocwvamzjmhgkndx
prjkqyluibsocwvahzjmhgfnsx
prtkqylcibsocwvzezjohgfndx
prtkqyluiksocwziezjmhgfndx
prikqyluiksocwvaezjmkgfndx
prtkgyluibsocwvwezjehgfndx
prtkqyluiysocwvaezjghxfndx
prtkqwluibsoxwvaezjmhgfhdx
prtkqylgibsocwvabzjmhzfndx
prtknyltibnocwvaezjmhgfndx
prdkqyluibrocwvaezjmhgnndx
prtwqyluibsoctvcezjmhgfndx
mrtkqyluibgocwvakzjmhgfndx
prtkqaouibsocwvaezjmhwfndx
prtkqyluihjocwvaezjmhgfpdx
prtkqyluikfxcwvaezjmhgfndx
prtkqybuixsocwvaczjmhgfndx
pvtkayluibsocwxaezjmhgfndx
grtkqgluibsocdvaezjmhgfndx
prlkqyluibsochvaezjmhgzndx
prtkqylxibsocmvaezjmhgfkdx
prtkqyluibsqctvaezjmpgfndx
putkqyluibsocqvaezjmhgfndw
prtjqyluibsiclvaezjmhgfndx
prtkqylvpvsocwvaezjmhgfndx
prnkqyluibsocwvaezjmhefsdx
prtktyluibsocwvaezjkhgrndx
prtkqyluibcovwvaezjthgfndx
prtkqcluibiocwvaezjmhggndx
prtkqyluihsocwveezjmhgfydx
prtklyluibsocwqaszjmhgfndx
prtkqyluibsocwvaezjmfznndx
prtkjyluijsocwvaeejmhgfndx
prtkqtluibsonwvaexjmhgfndx
prtkqyluinsocwbaezjmjgfndx
prtkqyluibslckvaezjmhgyndx
prtkqyluibsodwlpezjmhgfndx
prtkquluibsfcwvaezjhhgfndx
prtkqyluhbsocweaezsmhgfndx
prrkqyluinsocxvaezjmhgfndx
prtkqyluibsoswvaezjmhgyqdx
prtkqbluibdocwvlezjmhgfndx
prtkqyfuibsocpvaezjmhgfnwx
prtkqlluibsqjwvaezjmhgfndx
prtkqyluibrocwvaehjmjgfndx
prtkqyluibsoowvaezgmhgendx
wrtjqyluibsocwvaezfmhgfndx
prtvqyluhbsocwvaezjmtgfndx
prtkqyllibspcwvaezjmkgfndx
pqtzqyeuibsocwvaezjmhgfndx
prtkqyluibsolpvaezjmegfndx
przkayguibsocwvaezjmhgfndx
prtkqyluidsocwvaezjmyufndx
prtuqyluibsocwvaezjmfgfnkx
prtkqwluibsrcwvaezjchgfndx
prtkqyluibsotwhaozjmhgfndx
erwkqylhibsocwvaezjmhgfndx
prtkqyluibsocwvgezjmkgfedx
prskqyluiesocwvaezjmggfndx
prtkqylmitsocwvaezjmhgfnox
prtkqyluinnocwvaezjmhgfkdx
prtktyluibsokwvaezjmhgfcdx
prtkqyluibsomwvakvjmhgfndx
prtkqyltibloawvaezjmhgfndx
prtkqyluibxocwvaezgmhgqndx
prtkqyluibskcmvaezjmhgfngx
artkqylubbsotwvaezjmhgfndx
prtkqyluibzocwvhezjmhgfnbx
prskqkluibsocwvaezjmhgfjdx
prtkqyluibwocwvaezjkhglndx
prukqyluissocwvzezjmhgfndx
puhkqyluibsocwvaezjmhgfsdx
qrtkqyluibsocwvaeujmhgfndd
prtkqyluibsoctvaezjmagfnda
prtkquluibsocwkaezjmhgfqdx
prtkqyluubswcwvaezjmhvfndx
prfkqyluibsocwvaemrmhgfndx
pmtkqyluibpocwvaezjmhggndx
prtkqvluibiocwvaezjqhgfndx
prtkgypuibsocwvaezcmhgfndx
prtpqyquibsovwvaezjmhgfndx
prtwqyluiasocwvaexjmhgfndx
mrtzqyluibbocwvaezjmhgfndx
prtkqyluibsocwmaegwmhgfndx
prtkqyluibvncwvaqzjmhgfndx
prtkqyluiusocwvaezjmhmfbgx
prtkqyljibvocwvaezjehgfndx
prtkqyloibsopavaezjmhgfndx
prckqyakibsocwvaezjmhgfndx
prtkqyluibsdcwvaezjmngfddx
prekqylupbsocwvaezemhgfndx
hrtkqyluibhocwvaezjmhgfnde
prmkqyluibsocwvaezzfhgfndx
prtkqyluiccfcwvaezjmhgfndx
pdtkqyluxbsocwvaezjmhgendx
prokqyluibsocwvuezjmsgfndx
prtkqyluibsacwvaezjyhgfndv
prtkqmluibsocavaezjmhgfndc
prtkqyluibsocwvmezjmhgtnqx
prtkqytuibiocyvaezjmhgfndx
pktkqyiuibsocwvwezjmhgfndx
grtrqyluibsocwvaezjmhgfbdx
prtkqylsibjocwvaezjmhgfnyx
prtkqyhutbsocwvaexjmhgfndx
prtknyluibsocmvaezumhgfndx
prtkwyluibsocwvahzjmhgpndx
prtkqywuibsolhvaezjmhgfndx
prtkcyluibsoccvaezjthgfndx
prtkqyrdibsocwvaezjbhgfndx
prtkqyhuqbsocwvaezjmhgfxdx
pytkqyluibsocwvagzjmhgfndv
prtkqyliibsocwvaexwmhgfndx
prtkqyluibshcwvaeljphgfndx
prtkqyluibsocwvaerjzhbfndx
prtkqyduibsocwvaezvmhgfnzx
drtkqylhibsocwvaezjmhmfndx
prtkqyluibsocwvaezamfvfndx
brtkqyluqbsocwvaezjmhgpndx
prtkqyiuibsocwvuezjmhgfngx
urtkqyluibsocqvaeljmhgfndx
prtkqyluikaocwvaezjmhgfjdx
prqkqzouibsocwvaezjmhgfndx
prtkqyluibsocxvaezjmhgfnxv
prlkqyluibsoxwvaeijmhgfndx
prthuyluibsocwvaezjmhgfnhx
potkqyluizsocwvaezjmhifndx
fstkqyduibsocwvaezjmhgfndx
prtkqxluibsocwvaezjmhgffdm
prtkqylpibsozwvaezmmhgfndx
prxkqylbibsocwvaezjphgfndx
srtkqyluibsicnvaezjmhgfndx
prtktyluibsocwvaezjvhgfnax
pctkqyluxbsocwvaezwmhgfndx
prtkqylusbsoclvaezsmhgfndx
pwtkqyluibsocrvaezjmggfndx
prtkqyluibswcwraezjmhgfndd
prtkqyluibtocwiaezjmhgfnax
prtuqyluibsocwvajzjmngfndx
pwtkqyluibsocwvaerjmogfndx
petkqexuibsocwvaezjmhgfndx
pztkqyluibsocwvaerqmhgfndx
prtkqyluobsocwvaezjmapfndx
prtkqyluiinocwvaeljmhgfndx
prtkqyluibsoowvxezjmhgfnnx
lrtkqyluibsocwvfezjmhgfndc
prtkqyluibokcwvahzjmhgfndx
prtkqmlufbsocwvaegjmhgfndx
prtkqylribsocwvanzjmhgfnda
prtkqyluibspxwvaezkmhgfndx
prtiqyluibsbcwvaezjmhgfntx
prikqzluinsocwvaezjmhgfndx
prtkqnldibsocwvaezjmhxfndx
prtkqyluixsocsvaezjmhwfndx
hrtkqyluibsocwvaezjhhgfodx
prtkqyluibsrcwvaezjmhpfwdx
prtkqyluibsocwyaezjmhgffdk
prtkqyluidsocwvalmjmhgfndx
prukquluabsocwvaezjmhgfndx
prckqyluinsmcwvaezjmhgfndx
prbkqymuibsocwvaezjmhgfndc
prtkfylaibsocwvaezjmkgfndx
zrtkqyluibsocwvrbzjmhgfndx
crtkqyluibsocwvaejjmkgfndx
prttqyluibsocyvaezymhgfndx
prtkqylugbsocwvaezjxhgfmdx
prtkqyluibsocwdlezjmhgfnbx
prtkqjluibsocwvaozjhhgfndx
prtcjyluibsocwbaezjmhgfndx
rrtkqyluiblocwvaezjmhgundx
prtkkyluibsocwfaezjmhgfnyx
prtkqyuuibsocwvaezjmhgfogx
prtkyyluvbsocwvaezjmhgfnox
prpkqyluibyocwvaezjmhggndx
pdtkqyluibdocwvaezjmhgfndy
prtklysuibsocwvaezjmhgfnwx
prtkqyluabsouwvaekjmhgfndx
phtkqyluibsocwvaezjmhgfnxt
prtkqyxuibsocwvaezjmhpfnqx
prtkqyluibsodwsaezdmhgfndx
prtkbyluibsohwvaezjmhgfndr
xrtkqylhibsocwvtezjmhgfndx
prtkqyluvysocwvaezbmhgfndx
prtkqieuibsocwvaeojmhgfndx
pctkqyluibsocwvanzjmhgfnux
vrtkqyluibsozwvaezjmhgandx
prtkqyluiusocwvaezjmhmfngx
prbkqyluibsockvaxzjmhgfndx
prtkqyluibsonwvaczjmhgfndi
prtkqyluiblocwvaezjmhgfnau
prtkqyluibsocwvafzuchgfndx
prdkqyluiysocwvaezjmhgfnax
prnkqyouibsocwvaezjmhgfndq
mrtkqgluibsocwvpezjmhgfndx
pvtkqyluibsocwvaczjmhgnndx
trtkqwluibsohwvaezjmhgfndx
prmkqyluibsofwvaezjmhgfrdx
prtyqyluibpdcwvaezjmhgfndx
ertkqylulbsocwvaezjmhgfnax
prtkqyluibsacwvaeijmhgfndf
prtkqyluibyocwvapzjmhgpndx
potkqyluibgocwvaezjmhzfndx
prtkqyluibsocwyaezxmhgfnpx
prtkqkjuibsncwvaezjmhgfndx
prtqqyluibsocwlaezjmhgkndx
prtkxyluibnocwvaezjmhgkndx
prtkqyluiosocwvapzjmxgfndx
prtkqylumbsocwvyezimhgfndx
prukqyluibsocwvyezjmhgindx
prtkqylbibstcwvaezjxhgfndx
pctkqyuuibsocwvaezjuhgfndx
vrtkqyluibsocwvaezjmhgfnll
urtkqyluibsopwvaezjphgfndx
prtkceluibsocwvaepjmhgfndx
prwkxyluibsocwvaezjmhgfnzx
prtkqyluitsocwvaezqzhgfndx
prtkqkauibsorwvaezjmhgfndx
prtkqyluibsocwvaezfmftfndx
prtkiybuibsocwvaezjkhgfndx
prtkzyluibsocwgaezjmvgfndx
prtkqyluibsocwvaezjmhgqnxg
prtkqyluimsocwvauzjwhgfndx
prtkqyluibsacwgaezjmhgfndd
pwtkuyluibsccwvaezjmhgfndx
prtkqyluibsoawvaezjmvgfnlx
prtkqyluabsocwwaezjmhgftdx
patkqylnibsocwvaezjmhgfnox
prtkqyluibsocwlaxzkmhgfndx
pbtkqpluibsfcwvaezjmhgfndx
prtkqyluibsoywsaezjmhgxndx
prtkqyluibfocwvaezjyhgfhdx
pltbqylcibsocwvaezjmhgfndx
prtkdyluiisocwvvezjmhgfndx
prtkqkxuibsokwvaezjmhgfndx
prtkqyluibsoawvaezzmhgfndm
petkqyluibsgcwvaezjmhgfndu
prtkqyluibsoyxvaezjmlgfndx
prtkqyluibxocwvaezgmhnfndx
prtkikluibsocwvwezjmhgfndx
prbkqyluibsocwvaezjhhgfnux
prtkqylufbsxcwvaezjmhgfnfx
prtkqyluibsdcdvaezjmhgxndx
potkiyluibsocwvaezjmhkfndx
prtkqyluiosocsvhezjmhgfndx
prtkqyluibsocqbaezomhgfndx
prtihyluibsocwvaeujmhgfndx
prtuquruibsocwvaezjmhgfndx
prtkqyloibsocwvaeztmhifndx
ertuqyluibsocwvaeajmhgfndx`
