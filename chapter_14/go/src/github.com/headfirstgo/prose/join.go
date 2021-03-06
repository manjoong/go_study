package prose
import "strings"

func JoinWithCommas(phrases []string) string{
	if len(phrases) == 0{  // 쉽게 테이블 추가 가능
		return ""
	}else if len(phrases) == 1{
		return phrases[0]
	} else if len(phrases) == 2{
		return phrases[0] + " and " + phrases[1]
	} else{
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		// result += " and "
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}

