package day05

import "testing"

var sampleTestInput = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

var actualTestInput = []string{
	"seeds: 280775197 7535297 3229061264 27275209 77896732 178275214 2748861189 424413807 3663093536 130341162 613340959 352550713 1532286286 1115055792 1075412586 241030710 3430371306 138606714 412141395 146351614",
	"",
	"seed-to-soil map:",
	"2328878418 2173757269 55676388",
	"1604614185 284259424 18300254",
	"871339571 2660737044 103656521",
	"0 2281891142 53219493",
	"3600371492 4267529956 27437340",
	"1755732868 600174302 280911746",
	"225053611 2335110635 5908609",
	"3686191373 3693094019 157260119",
	"2301308430 2341019244 27569988",
	"2249924082 2817678844 51384348",
	"3843451492 3390401982 116711122",
	"866920588 2368589232 4418983",
	"319632173 0 10243502",
	"2384554806 302559678 74545000",
	"2036644614 2373008215 45489224",
	"480632565 1433058022 386288023",
	"3448897781 4166650940 100879016",
	"2082133838 377104678 167790244",
	"2644445775 2418497439 242239605",
	"1305360177 881086048 292964881",
	"3960162614 3638948449 54145570",
	"385155055 1819346045 95477510",
	"4021970860 3919567265 19798429",
	"3085107497 3850354138 69213127",
	"230962220 195589471 88669953",
	"1261037103 2129434195 44323074",
	"3627808832 3085107497 52096478",
	"329875675 544894922 55279380",
	"1622914439 2229433657 52457485",
	"1675371924 1352697078 80360944",
	"974996092 1914823555 107394862",
	"3549776797 3574405015 50594695",
	"3679905310 3632662386 6286063",
	"1082390954 1174050929 178646149",
	"2459099806 10243502 185345969",
	"160435271 2764393565 53285279",
	"3154320624 3939365694 227285246",
	"53219493 2022218417 107215778",
	"3381605870 3507113104 67291911",
	"4041769289 3137203975 253198007",
	"4014308184 3624999710 7662676",
	"213720550 2875352319 11333061",
	"1598325058 2869063192 6289127",
	"",
	"soil-to-fertilizer map:",
	"3389090999 2596751746 608341779",
	"105086589 871525822 220821758",
	"3187814092 3748571797 34038948",
	"1063982526 822144230 49381592",
	"4071580733 3394766608 36570558",
	"1969786851 1283647913 135669504",
	"2338688657 3782610745 100105665",
	"2105456355 1092347580 1733094",
	"1113364118 0 435632229",
	"325908347 1524404006 161994939",
	"2438794322 4084049871 210917425",
	"4108151291 3471527086 26197923",
	"3997432778 2273224151 74147955",
	"3161148447 2395583294 26665645",
	"621898776 435632229 386512001",
	"4134349214 3882716410 46400134",
	"2649711747 3205093525 86087105",
	"4256353957 2384855243 10728051",
	"487903286 1149652423 133995490",
	"0 1419317417 105086589",
	"4218232485 3356645136 38121472",
	"4180749348 2347372106 37483137",
	"2910301659 3497725009 250846788",
	"2735798852 2422248939 23521261",
	"2273224151 3291180630 65464506",
	"3262042960 3957001832 127048039",
	"1548996347 1686398945 420790504",
	"4267082008 3929116544 27885288",
	"2759320113 2445770200 150981546",
	"1008410777 1094080674 55571749",
	"3221853040 3431337166 40189920",
	"",
	"fertilizer-to-water map:",
	"2007324874 2683611319 23475372",
	"1313988767 1367341468 459297221",
	"1773285988 1938578086 234038886",
	"257353189 1257656713 109684755",
	"221312488 2981571133 36040701",
	"3464645820 3633779898 112824320",
	"1132250794 2637724313 45887006",
	"1178137800 0 135850967",
	"810899989 2295076898 321350805",
	"3706800774 4118877604 99172458",
	"3805973232 4218050062 17186653",
	"3823159885 3746604218 37073582",
	"2305284688 135850967 719035066",
	"3146779680 1826638689 111939397",
	"3024319754 2172616972 122459926",
	"3959767492 3783677800 335199804",
	"3934534290 4235236715 25233202",
	"408129309 854886033 402770680",
	"3258719077 2616427703 21296610",
	"3577470140 3504449264 129330634",
	"367037944 3238924322 41091365",
	"3860233467 4260469917 34497379",
	"3894730846 3464645820 39803444",
	"0 3017611834 221312488",
	"2030800246 2707086691 274484442",
	"",
	"water-to-light map:",
	"3878354467 2554575322 61538919",
	"280412900 1377596407 93760042",
	"684284359 195116588 99127467",
	"4078572373 3977755696 216394923",
	"2619535971 3527291682 385503365",
	"1493593206 1581399685 135608385",
	"3693719871 2668857297 184634596",
	"2554575322 3912795047 64960649",
	"3236793954 2853491893 456925917",
	"817630982 1044576172 333020235",
	"1205196262 294244055 132442996",
	"1739244827 1012243273 32332899",
	"1771577726 0 101384292",
	"0 101384292 4968095",
	"101989723 448809000 178423177",
	"4019116409 4194150619 21593654",
	"4040710063 3489429372 37862310",
	"4968095 627232177 97021628",
	"374172942 724253805 287989468",
	"783411826 106352387 34219156",
	"3005039336 2616114241 52743056",
	"3939893386 4215744273 79223023",
	"1629201591 1471356449 110043236",
	"3057782392 3310417810 179011562",
	"1337639258 1717008070 155953948",
	"1150651217 140571543 54545045",
	"662162410 426687051 22121949",
	"",
	"light-to-temperature map:",
	"2762906378 2204259687 134492279",
	"106349111 681820034 721307054",
	"997408643 115890239 153244527",
	"1512335036 1822350940 381908747",
	"3858505517 4045949291 249018005",
	"1894243783 3569882248 437693111",
	"3820131585 4007575359 38373932",
	"1505707247 3508506612 6627789",
	"2693538653 3439138887 69367725",
	"3029056394 2470409703 186112104",
	"54501260 269134766 51847851",
	"1452421821 3193053968 53285426",
	"850537566 0 115890239",
	"4107523522 3246339394 159303690",
	"4266827212 3164913884 28140084",
	"0 447698702 54501260",
	"3450202466 1452421821 369929119",
	"2897398657 2338751966 131657737",
	"966427805 650839196 30980838",
	"2331936894 3515134401 54747847",
	"2420180544 2656521807 273358109",
	"1206941949 502199962 83686237",
	"1150653170 320982617 56288779",
	"1361055492 585886199 42071596",
	"1290628186 377271396 70427306",
	"2386684741 3405643084 33495803",
	"827656165 627957795 22881401",
	"3215168498 2929879916 235033968",
	"",
	"temperature-to-humidity map:",
	"3512511508 3713321076 246696808",
	"1575177202 3115625465 474937809",
	"4090629460 2000274085 135040194",
	"1141835951 1250511056 29255284",
	"151257068 40127528 247347323",
	"796868343 1046685126 203825930",
	"32926673 475089348 115925678",
	"3091489330 2135314279 421022178",
	"4225669654 1871071039 45515305",
	"4018596644 1533302854 72032816",
	"2105871498 1332898874 176621643",
	"398604391 732156704 170521927",
	"2475375726 4289723131 5244165",
	"4271184959 1509520517 23782337",
	"2480619891 2698929267 337643768",
	"3911645962 1605335670 55526138",
	"3759208316 3590563274 122757802",
	"3058110291 3082246426 33379039",
	"1404278872 1916586344 83687741",
	"15821415 1279766340 17105258",
	"0 902678631 15821415",
	"148852351 943795712 2404717",
	"2966942025 2607761001 91168266",
	"1308993408 3036573035 45673391",
	"1000694273 591015026 141141678",
	"756740815 0 40127528",
	"3881966118 1308993408 23905466",
	"1171091235 918500046 25295666",
	"1539650216 3960017884 35526986",
	"1354666799 4047228473 49612073",
	"3905871584 1865296661 5774378",
	"3967172100 2556336457 51424544",
	"2050115011 1809540174 55756487",
	"1487966613 3995544870 51683603",
	"2818263659 1660861808 148678366",
	"569126318 287474851 187614497",
	"2282493141 4096840546 192882585",
	"1196386901 946200429 100484697",
	"",
	"humidity-to-location map:",
	"1176010827 433228953 11431675",
	"3343304446 3212529205 243926878",
	"3019400776 2809960450 121184456",
	"598353273 717816068 6343029",
	"2345678711 3761800910 44017519",
	"2389696230 4104193148 30873153",
	"334421220 1167819222 181559786",
	"2083857214 2364759797 169609191",
	"169788889 1665299576 164632331",
	"1547379326 837819247 83916771",
	"636928438 1349379008 3985237",
	"1997878235 4242793968 52173328",
	"4036700740 3456456083 70103553",
	"1701551461 724159097 55714896",
	"2050051563 4072991568 31201580",
	"1495464492 0 51914834",
	"1771229473 779873993 57945254",
	"1757266357 444660628 13963116",
	"640913675 1829931907 79566117",
	"604696302 1135587086 32232136",
	"2790976533 2058196402 571198",
	"0 183116767 169788889",
	"581107897 458623744 17245376",
	"1032415123 921736018 143595704",
	"2880695042 3109061793 3719369",
	"3596511671 3547168702 214632208",
	"1989658157 4064771490 8220078",
	"2791547731 1989658157 68538245",
	"1631296097 1065331722 70255364",
	"2253466405 2606294055 92212306",
	"2591363195 3865158152 199613338",
	"2532023472 3805818429 59339723",
	"1253517544 475869120 241946948",
	"3811143879 2058767600 225556861",
	"3334473574 2931144906 8830872",
	"720479792 1353364245 311935331",
	"2420569383 2698506361 111454089",
	"2081253143 2603689984 2604071",
	"3140585232 2939975778 159805668",
	"2860085976 3526559636 20609066",
	"1829174727 352905656 80323297",
	"515981006 117989876 65126891",
	"3587231324 3099781446 9280347",
	"1187442502 51914834 66075042",
	"2953735407 3112781162 65665369",
	"4187239629 4135066301 107727667",
	"3300390900 3178446531 34082674",
	"2884414411 2534368988 69320996",
	"4106804293 2284324461 80435336",
}

func TestFindLowestLocationNumberForSeeds(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{sampleTestInput, 35},
		{actualTestInput, 3374647},
	}
	for _, test := range tests {
		if output := FindLowestLocationNumberForSeeds(test.input); output != test.expected {
			t.Errorf("[TestFindLowestLocationNumberForSeeds] Expected(%d) Received(%d)", test.expected, output)
		}
	}
}

func TestFindLowestLocationNumberForSeedRange(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{sampleTestInput, 46},
		{actualTestInput, 6082852},
	}
	for _, test := range tests {
		if output := FindLowestLocationNumberForSeedRange(test.input); output != test.expected {
			t.Errorf("[TestFindLowestLocationNumberForSeedRange] Expected(%d) Received(%d)", test.expected, output)
		}
	}
}
