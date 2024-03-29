package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(puzzleInput, "\n")
	root := parseFileSystem(lines)
	assignDirSizes(root)

	// part 1
	dirSizeSum := sumDirSizes(root, 100000)
	fmt.Printf("part 1 dir size: %d\n", dirSizeSum)

	// part 2
	maxFileSystem := 70000000
	neededSpace := 30000000
	availableSpace := maxFileSystem - root.Size
	amountToDelete := neededSpace - availableSpace
	smallestAmount := findSmallestDirToDelete(root, amountToDelete)

	fmt.Printf("part 2 smallest dir size: %d\n", smallestAmount)
}

type File struct {
	Name     string
	Size     int
	Children map[string]*File
	Parent   *File
}

func parseFileSystem(lines []string) *File {
	if lines[0] != "$ cd /" {
		panic("unexpected input start")
	}

	root := &File{
		Name:     "/",
		Children: map[string]*File{},
	}

	lines = lines[1:]

	workingDir := root
	for len(lines) > 0 {
		if strings.HasPrefix(lines[0], "$ cd") {
			lines, workingDir = processChangeDir(lines, workingDir)
		} else if strings.HasPrefix(lines[0], "$ ls") {
			lines = processList(lines, workingDir)
		}

	}

	return root
}

func processChangeDir(lines []string, workingDir *File) ([]string, *File) {
	changeDir := strings.ReplaceAll(lines[0], "$ cd ", "")
	lines = lines[1:]

	if changeDir == "/" {
		root := workingDir
		for root.Parent != nil {
			root = root.Parent
		}
		return lines, root
	} else if changeDir == ".." {
		return lines, workingDir.Parent
	} else {
		return lines, workingDir.Children[changeDir]
	}
}

func processList(lines []string, workingDir *File) []string {
	lines = lines[1:]
	for len(lines) > 0 && string(lines[0][0]) != "$" {
		fileParts := strings.Split(lines[0], " ")
		name := fileParts[1]
		size, _ := strconv.Atoi(fileParts[0])
		var children map[string]*File
		if fileParts[0] == "dir" {
			children = map[string]*File{}
		}

		newFile := &File{
			Name:     fileParts[1],
			Size:     size,
			Children: children,
			Parent:   workingDir,
		}

		workingDir.Children[name] = newFile

		lines = lines[1:]
	}

	return lines
}

func assignDirSizes(f *File) {
	if f == nil {
		return
	}

	for _, file := range f.Children {
		if len(file.Children) > 0 {
			assignDirSizes(file)
		}
		f.Size += file.Size
	}
}

func sumDirSizes(f *File, threshold int) int {
	var visited []*File
	visited = append(visited, f)

	sum := 0
	for len(visited) > 0 {
		f := visited[len(visited)-1]
		visited = visited[:len(visited)-1]
		if len(f.Children) > 0 {
			if f.Size <= threshold {
				sum += f.Size
			}

			for _, child := range f.Children {
				visited = append(visited, child)
			}
		}
	}

	return sum
}

func findSmallestDirToDelete(f *File, amountToDelete int) int {
	var visited []*File
	visited = append(visited, f)

	currentSmallest := math.MaxInt
	for len(visited) > 0 {
		f := visited[len(visited)-1]
		visited = visited[:len(visited)-1]
		if len(f.Children) > 0 {
			if f.Size >= amountToDelete && f.Size < currentSmallest {
				currentSmallest = f.Size
			}

			for _, child := range f.Children {
				visited = append(visited, child)
			}
		}
	}

	return currentSmallest
}

const puzzleInput1 = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

const puzzleInput = `$ cd /
$ ls
150555 bch.lht
276291 ccqfdznj.sqg
dir csmqbhjv
dir czdqfr
dir fpfwfzrt
192660 qnbzgp
142026 rpphgdhp.jfr
dir sqphfslv
38077 tvpzh
$ cd csmqbhjv
$ ls
52822 bch.lht
dir dgj
dir fmmblb
dir hjwwtw
dir mtmhst
dir njsccfms
dir wmjsvq
$ cd dgj
$ ls
266484 bch.lht
dir brwncbh
dir dtdzsqps
216678 gvmdvcs.fmq
225948 mdjrhmhf
dir tvpzh
301487 tvpzh.wbp
2555 whsnd.spb
$ cd brwncbh
$ ls
238441 jwmr.plh
$ cd ..
$ cd dtdzsqps
$ ls
294443 fmmblb
81441 nsljm.fcz
$ cd ..
$ cd tvpzh
$ ls
186671 bch.lht
dir czdqfr
dir fgmz
300898 gvmdvcs.fmq
dir rjnv
dir szcrlzmr
dir tvpzh
dir vbrn
$ cd czdqfr
$ ls
dir czdqfr
dir hrhqhcjg
$ cd czdqfr
$ ls
8551 zmcqmq.zvf
$ cd ..
$ cd hrhqhcjg
$ ls
8255 wfvj.lnd
$ cd ..
$ cd ..
$ cd fgmz
$ ls
145921 wfvj.lnd
$ cd ..
$ cd rjnv
$ ls
305697 bch.lht
dir cdqv
dir czdqfr
288685 nsjnqh.fzq
210447 qlg
34660 rbnlc.gmc
143353 vsjg.njm
$ cd cdqv
$ ls
294789 czdqfr.nvt
$ cd ..
$ cd czdqfr
$ ls
298572 bch.lht
dir fmmblb
5599 tvpzh.pnf
47873 vmm
$ cd fmmblb
$ ls
dir cjtb
dir llg
dir rcb
190831 rwf.rzd
$ cd cjtb
$ ls
306568 fmmblb.hns
$ cd ..
$ cd llg
$ ls
68274 hbj.glq
261526 jplstj
285699 rmq
$ cd ..
$ cd rcb
$ ls
dir lmgrr
$ cd lmgrr
$ ls
306217 bch.lht
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd szcrlzmr
$ ls
dir mdjrhmhf
195828 pnllrfpr.lhl
dir rns
286738 sgzhcwj
$ cd mdjrhmhf
$ ls
dir clmdlmc
dir fmmblb
299416 fmmblb.ftr
dir mnm
22720 zpgvqpbr.glm
$ cd clmdlmc
$ ls
289014 fmmblb.njt
246797 gvmdvcs.fmq
92044 tvpzh.cfj
$ cd ..
$ cd fmmblb
$ ls
81512 czdqfr.ltc
173394 ntcdzdc.spn
$ cd ..
$ cd mnm
$ ls
260096 bch.lht
dir czdqfr
dir rwnqgjmm
$ cd czdqfr
$ ls
269467 nfphbtz
$ cd ..
$ cd rwnqgjmm
$ ls
218872 ldflbsm.rzh
256978 lwhhc
dir pgrtzw
$ cd pgrtzw
$ ls
32086 gvmdvcs.fmq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd rns
$ ls
282353 fcwhvqd.qvb
dir fwjb
dir prdd
dir rtrglmt
$ cd fwjb
$ ls
dir jlhqd
217527 tgvqql
$ cd jlhqd
$ ls
196505 qbrbcrgt.wqj
$ cd ..
$ cd ..
$ cd prdd
$ ls
144295 brwncbh
87560 dsrm
$ cd ..
$ cd rtrglmt
$ ls
219787 tcsrq.wzt
$ cd ..
$ cd ..
$ cd ..
$ cd tvpzh
$ ls
dir swrgwp
$ cd swrgwp
$ ls
dir cqzs
dir wtcsc
$ cd cqzs
$ ls
284674 rlhjp.dlf
$ cd ..
$ cd wtcsc
$ ls
50313 wfvj.lnd
$ cd ..
$ cd ..
$ cd ..
$ cd vbrn
$ ls
11650 czdqfr.ccz
dir fmmblb
145626 whsnd.spb
$ cd fmmblb
$ ls
204558 rdh.rms
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd fmmblb
$ ls
dir bhvpslz
$ cd bhvpslz
$ ls
dir jsmrb
$ cd jsmrb
$ ls
1472 fgctpsl.rgf
258119 whsnd.spb
$ cd ..
$ cd ..
$ cd ..
$ cd hjwwtw
$ ls
291890 fmmblb.zrt
$ cd ..
$ cd mtmhst
$ ls
dir dfqm
dir fmmblb
178539 jtmflfb
dir mdjrhmhf
dir rdzg
dir tvpzh
dir wqlf
$ cd dfqm
$ ls
dir czdqfr
dir fmmblb
dir gngvhs
dir ppjdlbp
dir shs
dir srdrb
$ cd czdqfr
$ ls
85242 gwqzn.ppr
$ cd ..
$ cd fmmblb
$ ls
76800 bqbnwdq.zfr
72865 fmmblb
119311 hfbtwgb.nmn
305342 wfvj.lnd
$ cd ..
$ cd gngvhs
$ ls
dir fmmblb
dir jfptn
dir jlqzq
dir mdjrhmhf
dir scmjf
dir tvpzh
$ cd fmmblb
$ ls
174392 whsnd.spb
$ cd ..
$ cd jfptn
$ ls
dir vjm
$ cd vjm
$ ls
193871 bch.lht
$ cd ..
$ cd ..
$ cd jlqzq
$ ls
dir fmmblb
187410 hcfdppj.hjh
dir mdjrhmhf
dir nhb
107865 qntnnqcp
dir wlqjpsh
$ cd fmmblb
$ ls
dir ccpvdzg
240944 fqsjc.cmc
$ cd ccpvdzg
$ ls
304195 whsnd.spb
$ cd ..
$ cd ..
$ cd mdjrhmhf
$ ls
242269 whsnd.spb
$ cd ..
$ cd nhb
$ ls
83616 mtfq
$ cd ..
$ cd wlqjpsh
$ ls
249472 gdg
$ cd ..
$ cd ..
$ cd mdjrhmhf
$ ls
213465 bbsltzd.fvd
dir tqmwn
$ cd tqmwn
$ ls
54460 bwnztql
262257 gvmdvcs.fmq
$ cd ..
$ cd ..
$ cd scmjf
$ ls
dir fhfg
268890 wqdcbprh
$ cd fhfg
$ ls
234806 whsnd.spb
$ cd ..
$ cd ..
$ cd tvpzh
$ ls
79625 zqbsb.mnq
$ cd ..
$ cd ..
$ cd ppjdlbp
$ ls
249976 zdqlmnll.nps
$ cd ..
$ cd shs
$ ls
107840 bsfcgmrw.dzw
dir fmmblb
307016 gvmdvcs.fmq
185876 jjhdj.vpn
254920 ltmp
dir rrg
312847 whsnd.spb
$ cd fmmblb
$ ls
161417 wppzrz.hjz
$ cd ..
$ cd rrg
$ ls
304639 bch.lht
$ cd ..
$ cd ..
$ cd srdrb
$ ls
dir bgchgtg
238106 hcfdppj.hjh
dir lgz
dir zmlm
$ cd bgchgtg
$ ls
dir hhbmwp
166735 psbpml.jdb
dir rtdwhlsq
171684 zcb.dlr
$ cd hhbmwp
$ ls
105599 cppw.mlb
dir dvzmpfzn
72811 fmmblb.bdd
293591 fnh.fdv
dir fspwz
dir gznwz
259403 tvpzh.bfj
$ cd dvzmpfzn
$ ls
127885 nbb.jgs
126415 whsnd.spb
$ cd ..
$ cd fspwz
$ ls
dir fzpfbcvv
dir jzrqvds
$ cd fzpfbcvv
$ ls
114984 czdqfr.pmn
157782 gzwrcdp.mtz
$ cd ..
$ cd jzrqvds
$ ls
47749 hcfdppj.hjh
$ cd ..
$ cd ..
$ cd gznwz
$ ls
77780 lnbtzj.bhz
$ cd ..
$ cd ..
$ cd rtdwhlsq
$ ls
300900 bbs
dir cchdchd
246695 djbmcmn
195895 fvwm.hrd
67210 hcfdppj.hjh
101017 zhnvhfm.wps
$ cd cchdchd
$ ls
52763 gvmdvcs.fmq
124993 tvpzh.nwc
dir zzsbq
$ cd zzsbq
$ ls
263814 zvhnm.sbv
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd lgz
$ ls
121541 wfvj.lnd
134550 whsnd.spb
$ cd ..
$ cd zmlm
$ ls
91948 cjqdnnh.hzr
250728 czdqfr.zjb
dir jnbhfnn
dir rnmghbht
dir tjhnp
144737 whsnd.spb
$ cd jnbhfnn
$ ls
49839 whsnd.spb
$ cd ..
$ cd rnmghbht
$ ls
69330 hcjspd
dir mfgbbvcc
dir qdh
$ cd mfgbbvcc
$ ls
dir gmgdfss
$ cd gmgdfss
$ ls
266023 hcfdppj.hjh
dir mdjrhmhf
dir ndcmqpc
291530 vqwf.bfc
$ cd mdjrhmhf
$ ls
dir brwncbh
$ cd brwncbh
$ ls
dir fmmblb
$ cd fmmblb
$ ls
dir ppwp
$ cd ppwp
$ ls
299841 mdjrhmhf.bdg
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ndcmqpc
$ ls
71727 mdjrhmhf.njs
218233 nwvfb.qzt
$ cd ..
$ cd ..
$ cd ..
$ cd qdh
$ ls
294005 jgbcmg.hjg
$ cd ..
$ cd ..
$ cd tjhnp
$ ls
dir vhjtqdg
$ cd vhjtqdg
$ ls
28638 gvmdvcs.fmq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd fmmblb
$ ls
111152 cnvzdz.cdb
252709 gvmdvcs.fmq
168368 lpfbzh
dir mdjrhmhf
dir mzrzzhpt
72606 qbbllr.shq
160820 whsnd.spb
$ cd mdjrhmhf
$ ls
32945 fmmblb.jpf
$ cd ..
$ cd mzrzzhpt
$ ls
dir gdsv
130898 hcfdppj.hjh
44339 hvnqt
dir mdjrhmhf
297785 mdwng.pbr
dir wwqmlhgg
281612 zvftqhm.pzl
$ cd gdsv
$ ls
10562 cprgfn
$ cd ..
$ cd mdjrhmhf
$ ls
284882 dpjs
202491 lrtgstpp.grn
3714 wfvj.lnd
$ cd ..
$ cd wwqmlhgg
$ ls
84563 fwvhfrh.cfp
4654 hcfdppj.hjh
248672 whsnd.spb
$ cd ..
$ cd ..
$ cd ..
$ cd mdjrhmhf
$ ls
dir mdjrhmhf
25175 mdjrhmhf.nrg
dir mslwjsp
207567 prrqb.qwj
282436 tvpzh.hjh
$ cd mdjrhmhf
$ ls
43516 jzrjwdd.vss
27326 wfvj.lnd
$ cd ..
$ cd mslwjsp
$ ls
141951 cqvh.zzq
$ cd ..
$ cd ..
$ cd rdzg
$ ls
312895 brwncbh.cpb
$ cd ..
$ cd tvpzh
$ ls
228240 bch.lht
dir fjtwlj
dir jnpgqsb
90201 ldh
271575 rfhh.vzr
69760 whsnd.spb
$ cd fjtwlj
$ ls
46126 rhdr.jgg
$ cd ..
$ cd jnpgqsb
$ ls
dir rllbvnm
$ cd rllbvnm
$ ls
212620 hcfdppj.hjh
$ cd ..
$ cd ..
$ cd ..
$ cd wqlf
$ ls
210864 fmmblb
dir jswfprpl
dir nztbsbq
$ cd jswfprpl
$ ls
dir fmmblb
dir lmwz
dir qvj
$ cd fmmblb
$ ls
201940 mhjlhc.npl
$ cd ..
$ cd lmwz
$ ls
dir brwncbh
dir dhff
288199 flgch
187825 gvmdvcs.fmq
203272 nvfllgvn.cjj
dir tvpzh
dir vrv
53288 zsrz.mrd
$ cd brwncbh
$ ls
174732 whsnd.spb
$ cd ..
$ cd dhff
$ ls
dir lddw
dir mfwnzprw
222134 tvpzh.gfm
$ cd lddw
$ ls
296685 bch.lht
$ cd ..
$ cd mfwnzprw
$ ls
6260 dvrcqzp.pmd
$ cd ..
$ cd ..
$ cd tvpzh
$ ls
dir gmvswdr
dir hfzdqdz
dir hhjwt
228542 htljf.tfl
57417 mdjrhmhf
248369 whsnd.spb
$ cd gmvswdr
$ ls
215582 fnzh.mhs
$ cd ..
$ cd hfzdqdz
$ ls
82080 dcw.vmt
142007 fmmblb.fnq
dir ggqcqcc
36091 hcfdppj.hjh
219562 mqb.qsm
2323 wfvj.lnd
$ cd ggqcqcc
$ ls
4679 whsnd.spb
$ cd ..
$ cd ..
$ cd hhjwt
$ ls
dir lbqpgd
dir tdb
$ cd lbqpgd
$ ls
177175 wfvj.lnd
$ cd ..
$ cd tdb
$ ls
270893 bch.lht
$ cd ..
$ cd ..
$ cd ..
$ cd vrv
$ ls
dir bqmnhdts
171679 czdqfr
$ cd bqmnhdts
$ ls
64972 pjzzs.qqd
$ cd ..
$ cd ..
$ cd ..
$ cd qvj
$ ls
dir bfwtb
dir bpnppqg
271362 gcs.srv
dir gvlqddlb
70000 gzdr.ndb
103231 pwspfm
dir swgfvtf
$ cd bfwtb
$ ls
200600 czdqfr.sbv
dir wmbhgcw
$ cd wmbhgcw
$ ls
310044 zwfrcld.mtb
$ cd ..
$ cd ..
$ cd bpnppqg
$ ls
120203 dggplss.whb
$ cd ..
$ cd gvlqddlb
$ ls
12163 bch.lht
154059 llpw
dir pglr
dir rbfpbcpd
dir rstg
209132 tvpzh.djc
287422 wlhjvsz
$ cd pglr
$ ls
145907 brwncbh.cqp
$ cd ..
$ cd rbfpbcpd
$ ls
292475 brwncbh
dir czdqfr
dir dcnh
dir fmmblb
dir qwrsrdr
$ cd czdqfr
$ ls
66107 tvpzh.wfd
$ cd ..
$ cd dcnh
$ ls
297110 rpzsrws.sft
$ cd ..
$ cd fmmblb
$ ls
25829 wmmhq
$ cd ..
$ cd qwrsrdr
$ ls
199762 fmmblb.qpm
$ cd ..
$ cd ..
$ cd rstg
$ ls
236706 bch.lht
$ cd ..
$ cd ..
$ cd swgfvtf
$ ls
227938 lrbcddd.btw
282992 rlnwrd
$ cd ..
$ cd ..
$ cd ..
$ cd nztbsbq
$ ls
116544 czdqfr
dir fhwmhvdn
68451 gvmdvcs.fmq
247136 hcfdppj.hjh
dir mzvq
dir tvpzh
dir wtllfshw
$ cd fhwmhvdn
$ ls
122155 cmtlhcdw
39927 twbfczfb.lcp
$ cd ..
$ cd mzvq
$ ls
dir lnjtfgh
$ cd lnjtfgh
$ ls
161712 phfn
$ cd ..
$ cd ..
$ cd tvpzh
$ ls
dir jgpr
dir mdjrhmhf
dir rnntmvpr
$ cd jgpr
$ ls
85238 gvmdvcs.fmq
$ cd ..
$ cd mdjrhmhf
$ ls
231007 gzrsgvp
$ cd ..
$ cd rnntmvpr
$ ls
285020 tvpzh
$ cd ..
$ cd ..
$ cd wtllfshw
$ ls
135281 gvmdvcs.fmq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd njsccfms
$ ls
dir zdddqvq
$ cd zdddqvq
$ ls
91520 whmwbpd
$ cd ..
$ cd ..
$ cd wmjsvq
$ ls
dir fmmblb
dir htmb
dir lmzhpth
dir mrjq
142052 wldwl
$ cd fmmblb
$ ls
173331 fmmblb.fbm
$ cd ..
$ cd htmb
$ ls
dir gtm
$ cd gtm
$ ls
7647 hcfdppj.hjh
$ cd ..
$ cd ..
$ cd lmzhpth
$ ls
dir tvpzh
dir vdqg
dir zvhmfm
$ cd tvpzh
$ ls
259073 wfvj.lnd
$ cd ..
$ cd vdqg
$ ls
199919 bch.lht
169430 fmmblb.ttq
231127 hcfdppj.hjh
$ cd ..
$ cd zvhmfm
$ ls
274752 gvmdvcs.fmq
$ cd ..
$ cd ..
$ cd mrjq
$ ls
dir fqcspsv
dir fvvmlbf
252802 gnclv
82538 gvmdvcs.fmq
296461 jvbnhrpd
104809 lbh
46038 tfq.gtv
$ cd fqcspsv
$ ls
dir fjqnf
$ cd fjqnf
$ ls
dir tvpzh
$ cd tvpzh
$ ls
dir lhwdrcwz
$ cd lhwdrcwz
$ ls
129410 lcd.dzd
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd fvvmlbf
$ ls
263142 brwncbh.fgg
232863 cvwv
100409 qphswnlb.vpq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd czdqfr
$ ls
dir czdqfr
307029 lntbmr.brq
dir mdjrhmhf
dir tvpzh
$ cd czdqfr
$ ls
289585 gvmdvcs.fmq
$ cd ..
$ cd mdjrhmhf
$ ls
137110 brwncbh.dbt
$ cd ..
$ cd tvpzh
$ ls
258325 wfvj.lnd
$ cd ..
$ cd ..
$ cd fpfwfzrt
$ ls
221993 hcfdppj.hjh
56206 lphphf.thd
255850 mdjrhmhf.hqt
$ cd ..
$ cd sqphfslv
$ ls
104415 mdjrhmhf.gbr
dir pjtzm
dir spnr
61982 stbwbpzf.hwm
$ cd pjtzm
$ ls
dir brwncbh
178769 jnmgzcqh.mht
dir tvpzh
153193 tvpzh.cdc
dir zstcqndd
$ cd brwncbh
$ ls
117400 brwncbh.mvq
dir fbzvn
3239 fmmblb.hzm
73631 hqtnwbgw
dir npzf
172382 pdgg
$ cd fbzvn
$ ls
152876 bch.lht
293370 czc.jtt
43408 czdqfr
dir fqg
105569 qqmcwjz.bzd
19734 whsnd.spb
$ cd fqg
$ ls
88780 jffmmm.qff
22385 lqrhch.vlm
$ cd ..
$ cd ..
$ cd npzf
$ ls
190792 pdtlbrqt.ghq
$ cd ..
$ cd ..
$ cd tvpzh
$ ls
dir blgjrjt
dir dlzjtlbt
dir jnwm
dir qprbng
305178 tbvg
dir tvpzh
308117 wrbldl.jhp
$ cd blgjrjt
$ ls
125644 bch.lht
104319 ctglz.lmd
131972 hcfdppj.hjh
312663 tbgqwf
$ cd ..
$ cd dlzjtlbt
$ ls
98262 bvfldnd.sls
173244 dlmf.mzh
75842 rlbpgjh.wtb
$ cd ..
$ cd jnwm
$ ls
58963 vjftmflp.dcs
$ cd ..
$ cd qprbng
$ ls
308054 fzbrm.jdt
310679 gnwn.psf
257810 tpsgjjdq.mdw
dir tvpzh
305633 vpbwhqs.cqb
$ cd tvpzh
$ ls
220904 rlpj.rdc
$ cd ..
$ cd ..
$ cd tvpzh
$ ls
dir fmmblb
$ cd fmmblb
$ ls
104020 crhq.pnb
74914 czdqfr
$ cd ..
$ cd ..
$ cd ..
$ cd zstcqndd
$ ls
dir brwncbh
dir jwpbtpr
306359 ljg
dir wbssjjd
$ cd brwncbh
$ ls
114145 wfvj.lnd
$ cd ..
$ cd jwpbtpr
$ ls
94567 bch.lht
dir czdqfr
dir fsqvbv
134187 wnpcqmw.tws
$ cd czdqfr
$ ls
dir mdjrhmhf
$ cd mdjrhmhf
$ ls
122398 mdjrhmhf.zts
dir svhh
101920 whsnd.spb
dir zcsnqjj
$ cd svhh
$ ls
120875 tvpzh.jls
$ cd ..
$ cd zcsnqjj
$ ls
279979 tljnvzbq.sgj
$ cd ..
$ cd ..
$ cd ..
$ cd fsqvbv
$ ls
114278 hwrhzs.jbj
dir tvpzh
$ cd tvpzh
$ ls
17235 gvmdvcs.fmq
$ cd ..
$ cd ..
$ cd ..
$ cd wbssjjd
$ ls
dir mdjrhmhf
$ cd mdjrhmhf
$ ls
11336 fqsp
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd spnr
$ ls
298067 sznwzj`
