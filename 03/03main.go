package main

import (
	"fmt"
	"regexp"
)

func main() {
	data := `()why()how()^*<~^where(){mul(986,194)!']^!mul(659,964)<-[>,(select():when()mul(102,589)who()]#when()^do()$when()+*mul(433,264)!%where() mul(676,814){who()mul(791,825)?mul(69,953)}mul(348,556)##how()what()who(){[>?how(151,194)mul(346,158)mul(698,833)>(:mul(249,39)!{<{//~when(){mul(686,621)what()mul(516,43)*who()mul(326,716)select()do()^#$/&mul(692,750)when()&[?->{mul(218,8)&&mulwho() $where()}{}#mul(820,744)/mul(230,927)^ :where(351,105)~/(;mul(900,172)$]{?$#where()mul(652,481)<}don't()*where()select(200,394)select()#who()mul(650,602):}when(552,753)-when()how()'mul(349,401),&how(){'mul(989,165)%$* ,,&mul(906,532)mul(701,543)when():?#what()~+how()mul(113,743)[$;select()!how()mul(68who()mul(153,928)%()^[from()mul(555,552)@select(522,482)mul(295,900)where()/mul(189,189)+'&mul(490,144)++/<mul(252,195)&<<mul(274,206)where()};: why()do()[mul(995,775)~[%&*mul(119,442)(why()what()&)from()%why()where()mul(278,384),mul(2?>&from():why()+*do()how()<?mul(635,799)];>[~when(612,413)-when()?do()}who()mul@/why() {?-%%mul(126,29)@+from();mul(548,673)mul(51,970)--}what()~,what()(mul(261,780)@;'why()>;what()mul(843,665)when();@mul(575,431)]:;>{#mul(123,103);]where()<why()?do()/]<?how(484,347)!-,&mul(513,533)mul(427,500)!:mul(649,325)$'~}'((don't()}!select()+%mul(166,962)how()#why()@/$[>where() mul(331,417)?+select()}when()mul(730,923what()who(514,796)where(716,633);when());why(122,722)(mul(923,885)where()''#'where(645,705)}#select()!don't()where(),,;!}&^mul(609-<why()!@)how()!^ mul(367,932)]@how()**when();;do() @$why()[*what()mul(365,75)?<} what()~~!mul(392,384)-$>'(mul(245,580):mul(898,814)how()+mul(999,143)),select()mul(251,507)(%<:who()why()how()/%mul(192,864)
*,/~):^[mul(802,939~;!:@mul(656,831)}]?~mul(92,389)/:who()&why()'mul(698,92):who()]/< @@@mul(956,581)who(252,379),#&,%mul(315,398)mul(518,43)@, &/[* mul(344,100)how();mul(826,831)<when()(>when()select()++'mul(175,965) [ mul(786,297)from()how(520,671):*?]&%]mul(647,834)why()when()select()/mul(585,934)!@<&how()!/mul(965,795)!,@/how())(#{mul(701,685) mul(51,440)/)mul(382,190)@why()how()>>#+;%{mul(808,654)#*mul(128,793)%:->mul(397,504)];$]from()mulselect(533,968))>why();^mul(661,310)mul(415,957)}?&%#@(mul(261,4)*;mul(292,885)!mul(639,791)from()how()when()mul(451,10)mul(853,231)+~+']>mul(543,437)mul(218,441) *mul(447,927)why()from()mul(178,473)how()<*why()[&how()how(),mul(139,541)>~why(644,46) who()what()+~(mul(888,667)^from()&^when()^when()-where():mul(585,185)>from())]+mul(766,787)mul(16,587)%@mul(475,198){<who()select() #-{}]mul(961,363)~%mul(851,371)-(who():$what():>^mul(353,741)&~]$<;]:mul(529,533)where()~mul(980,696)how()when()/*>when()#^don't()![]#[*%-mul(315,512/when(57,948)?why()mul(107,632)%>&who()[mul(554,91):from()!who()?how()why()%#when(313,134)don't()mul(694,300)/~(,+who(382,968)}from()mul(69,889)')w`

	// Compile the regex pattern
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	
	// Find all matches
	matches := re.FindAllString(data, -1)

	for _, match := range matches {
		fmt.Println(match)
	}

	// Here you would continue with further processing of the matched data ...
}