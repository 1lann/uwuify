// package uwuify is a mistake
package uwuify

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/OneOfOne/xxhash"
)

var replacer = strings.NewReplacer(
	"r", "w", "R", "W",
	"l", "w", "L", "W",
	"f", "fw", "F", "FW",
	"ck", "cc", "cK", "cC", "Ck", "Cc", "CC", "CC",
	"ff", "ffw", "fF", "fFW", "Ff", "Ffw", "FF", "FFW",
	// "ing", "in", "inG", "iN", "iNg", "iN", "iNG", "iN", "Ing", "In", "InG", "IN", "INg", "IN", "ING", "IN",
	"tha", "da", "thA", "dA", "tHa", "Da", "tHA", "DA", "Tha", "Da", "ThA", "DA", "THa", "Da", "THA", "DA",
	"the", "de", "thE", "dE", "tHe", "De", "tHE", "DE", "The", "De", "ThE", "DE", "THe", "De", "THE", "DE",
	"...", ",,,", "..", ",,", "…", ",,,",
	".", "!!!", "?", "?!!!", "!", "!!!",
)

// Kaomoji is a big list of assorted faces.
var Kaomoji = []string{
	"owo", "uwu", "OwO", "UwU", "uvu", "ovo", "^w^", ">w<", "(* ^ ω ^)", "(´ ∀ ` *)", "٩(◕‿◕｡)۶",
	"☆*:.｡.o(≧▽≦)o.｡.:*☆", "(o^▽^o)", "(⌒▽⌒)☆", "<(￣︶￣)>", "。.:☆*:･'(*⌒―⌒*)))", "ヽ(・∀・)ﾉ", "(´｡• ω •｡`)",
	"(￣ω￣)", "｀;:゛;｀;･(°ε° )", "(o･ω･o)", "(＠＾◡＾)", "ヽ(*・ω・)ﾉ", "(o_ _)ﾉ彡☆", "(^人^)", "(o´▽`o)", "(*´▽`*)",
	"｡ﾟ( ﾟ^∀^ﾟ)ﾟ｡", "( ´ ω ` )", "(((o(*°▽°*)o)))", "(≧◡≦)", "(o´∀`o)", "(´• ω •`)", "(＾▽＾)", "(⌒ω⌒)", "∑d(°∀°d)",
	"╰(▔∀▔)╯", "(─‿‿─)", "(*^‿^*)", "ヽ(o^ ^o)ﾉ", "(✯◡✯)", "(◕‿◕)", "(*≧ω≦*)", "(☆▽☆)", "(⌒‿⌒)", "＼(≧▽≦)／",
	"ヽ(o＾▽＾o)ノ", "☆ ～('▽^人)", "(*°▽°*)", "٩(｡•́‿•̀｡)۶", "(✧ω✧)", "ヽ(*⌒▽⌒*)ﾉ", "(´｡• ᵕ •｡`)", "( ´ ▽ ` )",
	"(￣▽￣)", "╰(*´︶`*)╯", "ヽ(>∀<☆)ノ", "o(≧▽≦)o", "(☆ω☆)", "(っ˘ω˘ς )", "＼(￣▽￣)／", "(*¯︶¯*)", "＼(＾▽＾)／",
	"٩(◕‿◕)۶", "(o˘◡˘o)", "\\(★ω★)/", "\\(^ヮ^)/", "(〃＾▽＾〃)", "(╯✧▽✧)╯", "o(>ω<)o", "o( ❛ᴗ❛ )o", "｡ﾟ(TヮT)ﾟ｡",
	"( ‾́ ◡ ‾́ )", "(ﾉ´ヮ`)ﾉ*: ･ﾟ", "(b ᵔ▽ᵔ)b", "(๑˃ᴗ˂)ﻭ", "(๑˘︶˘๑)", "( ˙꒳​˙ )", "(*꒦ິ꒳꒦ີ)", "°˖✧◝(⁰▿⁰)◜✧˖°", "(´･ᴗ･ ` )",
	"(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧", "(„• ֊ •„)", "(.❛ ᴗ ❛.)", "(⁀ᗢ⁀)", "(￢‿￢ )", "(¬‿¬ )", "(*￣▽￣)b", "( ˙▿˙ )", "(¯▿¯)",
	"(ﾉ´ з `)ノ", "(♡μ_μ)", "(*^^*)♡", "☆⌒ヽ(*'､^*)chu", "(♡-_-♡)", "(￣ε￣＠)", "ヽ(♡‿♡)ノ", "( ´ ∀ `)ノ～ ♡",
	"(─‿‿─)♡", "(´｡• ᵕ •｡`) ♡", "(*♡∀♡)", "(｡・//ε//・｡)", "(´ ω `♡)", "♡( ◡‿◡ )", "(◕‿◕)♡", "(/▽＼*)｡o○♡",
	"(ღ˘⌣˘ღ)", "(♡°▽°♡)", "♡(｡- ω -)", "♡ ～('▽^人)", "(´• ω •`) ♡", "(´ ε ` )♡", "(´｡• ω •｡`) ♡",
	"( ´ ▽ ` ).｡ｏ♡", "╰(*´︶`*)╯♡", "(*˘︶˘*).｡.:*♡", "(♡˙︶˙♡)", "♡＼(￣▽￣)／♡", "(≧◡≦) ♡", "(⌒▽⌒)♡",
	"(*¯ ³¯*)♡", "(っ˘з(˘⌣˘ ) ♡", "♡ (˘▽˘>ԅ( ˘⌣˘)", "( ˘⌣˘)♡(˘⌣˘ )", "(/^-^(^ ^*)/ ♡", "٩(♡ε♡)۶", "σ(≧ε≦σ) ♡",
	"♡ (⇀ 3 ↼)", "♡ (￣З￣)", "(❤ω❤)", "(˘∀˘)/(μ‿μ) ❤", "❤ (ɔˆз(ˆ⌣ˆc)", "(´♡‿♡`)", "(°◡°♡)", "Σ>―(〃°ω°〃)♡→",
	"(´,,•ω•,,)♡", "(*・ω・)ﾉ", "(￣▽￣)ノ", "(°▽°)/", "( ´ ∀ ` )ﾉ", "(^-^*)/", "(＠´ー`)ﾉﾞ", "(´• ω •`)ﾉ",
	"( ° ∀ ° )ﾉﾞ", "ヾ(*'▽'*)", "＼(⌒▽⌒)", "ヾ(☆▽☆)", "( ´ ▽ ` )ﾉ", "(^０^)ノ", "~ヾ(・ω・)", "(・∀・)ノ", "ヾ(・ω・*)",
	"(*°ｰ°)ﾉ", "(・_・)ノ", "(o´ω`o)ﾉ", "( ´ ▽ ` )/", "(￣ω￣)/", "( ´ ω ` )ノﾞ", "(⌒ω⌒)ﾉ", "(o^ ^o)/", "(≧▽≦)/",
	"(✧∀✧)/", "(o´▽`o)ﾉ", "(￣▽￣)/", "(づ￣ ³￣)づ", "(つ≧▽≦)つ", "(つ✧ω✧)つ", "(づ ◕‿◕ )づ", "(⊃｡•́‿•̀｡)⊃",
	"(つ . •́ _ʖ •̀ .)つ", "(っಠ‿ಠ)っ", "(づ◡﹏◡)づ", "⊂(´• ω •`⊂)", "⊂(･ω･*⊂)", "⊂(￣▽￣)⊃", "⊂( ´ ▽ ` )⊃", "( ~*-*)~",
	"(^_~)", "( ﾟｏ⌒)", "(^_-)≡☆", "(^ω~)", "(>ω^)", "(~人^)", "(^_-)", "( -_・)", "(^_<)〜☆", "(^人<)〜☆",
	"☆⌒(≧▽​° )", "☆⌒(ゝ。∂)", "(^_<)", "(^_−)☆", "(･ω<)☆", "(^.~)☆", "(^.~)",
}

// Translate a string into its uwuified form.
func Translate(s string) string {
	return TranslateNoKaomoji(s) + " " + DeterministicKaomoji(s)
}

// Translate a string into its uwuified form without adding a kaomoji to the end.
func TranslateNoKaomoji(s string) string {
	return replacer.Replace(s)
}

// DeterministicKaomoji computes the hash of the input string and uses it to return a kaomoji from the Kaomoji list.
func DeterministicKaomoji(s string) string {
	return Kaomoji[xxhash.ChecksumString32(s)%uint32(len(Kaomoji))]
}

// RandomKaomoji returns a random kaomoji from the Kaomoji list.
func RandomKaomoji() string {
	return Kaomoji[rand.Intn(len(Kaomoji))]
}

// Pwint fwowmats using the defwauwt fwowmats fwow its opewands and wwites to standawd output!!! (￢‿￢ ).
func Pwint(a ...interface{}) {
	fmt.Print(Spwint(a...))
}

// Pwintfw fwowmats accowding to a fwowmat specifwiew and wwites to standawd output!!!
func Pwintfw(fowmat string, a ...interface{}) {
	fmt.Print(Spwintf(fowmat, a...))
}

// Pwintwn fwowmats using the defwauwt fwowmats fwow its opewands and wwites to the standawd output!!! (´• ω •`).
func Pwintwn(a ...interface{}) {
	fmt.Print(Spwintwn(a...))
}

// Spwint fwowmats using the defwauwt fwowmats fwow its opewands and wetuwns the wesuwting stwing!!!
// Spaces awe added between opewands when neithew is a stwing!!! (♡°▽°♡).
func Spwint(a ...interface{}) string {
	return Translate(fmt.Sprint(a...))
}

// Spwintfw fwowmats accowding to a fwowmat specifwiew and wetuwns the wesuwting stwing!!! ヾ(・ω・*).
func Spwintf(fowmat string, a ...interface{}) string {
	return Translate(fmt.Sprintf(fowmat, a...))
}

// Spwintwn fwowmats using the defwauwt fwowmats fwow its opewands and wetuwns the wesuwting stwing!!!
// Spaces awe awways added between opewands and a newwine is appended!!! (＠´ー`)ﾉﾞ.
func Spwintwn(a ...interface{}) string {
	return Translate(fmt.Sprintln(a...))
}
