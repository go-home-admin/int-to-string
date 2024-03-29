package int_to_string

import (
	"crypto/md5"
	"encoding/json"
	"math"
	"strconv"
	"strings"
)

var words = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type Config struct {
	Words         []string                     `json:"words"` // 自热排序到字母
	wordsContrary map[string]int               // words相反值
	wordsLen      int                          // 模板长度
	Map           map[string]map[string]string `json:"map"` // 自热排序到字母, 个数=len(Words), 第一位string 为个位数映射
	mapContrary   map[string]map[string]string
}

type Factory struct {
	config Config
	len    int // 生成code的固定长度
}

var defaultConfigString = `{"words":["S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4","H","M","N","R","6","F","L","X","E","V","Z","Q"],"map":{"0":{"0":"E","1":"K","2":"A","3":"7","4":"P","5":"Q","6":"V","7":"B","8":"R","9":"F","A":"S","B":"L","C":"G","D":"T","E":"0","F":"1","G":"C","H":"W","I":"M","J":"N","K":"D","L":"2","M":"X","N":"Y","O":"O","P":"H","Q":"4","R":"Z","S":"8","T":"I","U":"J","V":"5","W":"9","X":"3","Y":"U","Z":"6"},"1":{"0":"J","1":"O","2":"D","3":"9","4":"U","5":"P","6":"Z","7":"A","8":"Q","9":"E","A":"R","B":"K","C":"F","D":"S","E":"3","F":"0","G":"B","H":"V","I":"L","J":"M","K":"C","L":"1","M":"W","N":"X","O":"N","P":"G","Q":"6","R":"Y","S":"7","T":"H","U":"I","V":"4","W":"8","X":"2","Y":"T","Z":"5"},"2":{"0":"Y","1":"3","2":"S","3":"O","4":"9","5":"A","6":"E","7":"T","8":"5","9":"Z","A":"6","B":"4","C":"U","D":"7","E":"I","F":"J","G":"Q","H":"F","I":"0","J":"1","K":"R","L":"G","M":"B","N":"C","O":"2","P":"V","Q":"L","R":"D","S":"P","T":"W","U":"X","V":"M","W":"N","X":"H","Y":"8","Z":"K"},"3":{"0":"O","1":"T","2":"I","3":"E","4":"Z","5":"U","6":"4","7":"F","8":"V","9":"J","A":"W","B":"P","C":"K","D":"X","E":"8","F":"5","G":"G","H":"0","I":"Q","J":"R","K":"H","L":"6","M":"1","N":"2","O":"S","P":"L","Q":"B","R":"3","S":"C","T":"M","U":"N","V":"9","W":"D","X":"7","Y":"Y","Z":"A"},"4":{"0":"B","1":"G","2":"5","3":"1","4":"M","5":"N","6":"R","7":"6","8":"I","9":"C","A":"J","B":"H","C":"7","D":"K","E":"V","F":"W","G":"3","H":"S","I":"D","J":"E","K":"4","L":"T","M":"O","N":"P","O":"F","P":"8","Q":"Y","R":"Q","S":"2","T":"9","U":"A","V":"Z","W":"0","X":"U","Y":"L","Z":"X"},"5":{"0":"9","1":"E","2":"3","3":"0","4":"K","5":"L","6":"P","7":"4","8":"M","9":"A","A":"N","B":"F","C":"B","D":"I","E":"T","F":"U","G":"5","H":"Q","I":"G","J":"H","K":"6","L":"V","M":"R","N":"S","O":"D","P":"C","Q":"X","R":"O","S":"1","T":"7","U":"8","V":"Y","W":"2","X":"W","Y":"J","Z":"Z"},"6":{"0":"E","1":"K","2":"A","3":"7","4":"P","5":"Q","6":"V","7":"B","8":"R","9":"F","A":"S","B":"L","C":"G","D":"T","E":"0","F":"1","G":"C","H":"W","I":"M","J":"N","K":"D","L":"2","M":"X","N":"Y","O":"O","P":"H","Q":"4","R":"Z","S":"8","T":"I","U":"J","V":"5","W":"9","X":"3","Y":"U","Z":"6"},"7":{"0":"P","1":"V","2":"L","3":"I","4":"0","5":"1","6":"6","7":"M","8":"2","9":"Q","A":"3","B":"W","C":"R","D":"4","E":"B","F":"C","G":"N","H":"7","I":"X","J":"Y","K":"O","L":"D","M":"8","N":"9","O":"Z","P":"S","Q":"F","R":"A","S":"J","T":"T","U":"U","V":"G","W":"K","X":"E","Y":"5","Z":"H"},"8":{"0":"A","1":"G","2":"6","3":"3","4":"L","5":"M","6":"R","7":"7","8":"N","9":"B","A":"O","B":"H","C":"C","D":"P","E":"W","F":"X","G":"8","H":"S","I":"I","J":"J","K":"9","L":"Y","M":"T","N":"U","O":"K","P":"D","Q":"0","R":"V","S":"4","T":"E","U":"F","V":"1","W":"5","X":"Z","Y":"Q","Z":"2"},"9":{"0":"4","1":"A","2":"0","3":"X","4":"F","5":"G","6":"L","7":"1","8":"H","9":"5","A":"I","B":"B","C":"6","D":"J","E":"Q","F":"R","G":"2","H":"M","I":"C","J":"D","K":"3","L":"S","M":"N","N":"O","O":"E","P":"7","Q":"U","R":"P","S":"Y","T":"8","U":"9","V":"V","W":"Z","X":"T","Y":"K","Z":"W"},"A":{"0":"O","1":"T","2":"I","3":"E","4":"Z","5":"U","6":"4","7":"F","8":"V","9":"J","A":"W","B":"P","C":"K","D":"X","E":"8","F":"5","G":"G","H":"0","I":"Q","J":"R","K":"H","L":"6","M":"1","N":"2","O":"S","P":"L","Q":"B","R":"3","S":"C","T":"M","U":"N","V":"9","W":"D","X":"7","Y":"Y","Z":"A"},"B":{"0":"U","1":"0","2":"Q","3":"N","4":"5","5":"6","6":"B","7":"R","8":"7","9":"V","A":"8","B":"1","C":"W","D":"9","E":"G","F":"H","G":"S","H":"C","I":"2","J":"3","K":"T","L":"I","M":"D","N":"E","O":"4","P":"X","Q":"K","R":"F","S":"O","T":"Y","U":"Z","V":"L","W":"P","X":"J","Y":"A","Z":"M"},"C":{"0":"O","1":"T","2":"I","3":"E","4":"Z","5":"U","6":"4","7":"F","8":"V","9":"J","A":"W","B":"P","C":"K","D":"X","E":"8","F":"5","G":"G","H":"0","I":"Q","J":"R","K":"H","L":"6","M":"1","N":"2","O":"S","P":"L","Q":"B","R":"3","S":"C","T":"M","U":"N","V":"9","W":"D","X":"7","Y":"Y","Z":"A"},"D":{"0":"M","1":"R","2":"G","3":"C","4":"X","5":"Y","6":"2","7":"H","8":"Z","9":"N","A":"U","B":"S","C":"O","D":"V","E":"6","F":"7","G":"I","H":"3","I":"T","J":"P","K":"F","L":"8","M":"4","N":"0","O":"Q","P":"J","Q":"9","R":"1","S":"D","T":"K","U":"L","V":"A","W":"E","X":"5","Y":"W","Z":"B"},"E":{"0":"U","1":"0","2":"Q","3":"N","4":"5","5":"6","6":"B","7":"R","8":"7","9":"V","A":"8","B":"1","C":"W","D":"9","E":"G","F":"H","G":"S","H":"C","I":"2","J":"3","K":"T","L":"I","M":"D","N":"E","O":"4","P":"X","Q":"K","R":"F","S":"O","T":"Y","U":"Z","V":"L","W":"P","X":"J","Y":"A","Z":"M"},"F":{"0":"U","1":"0","2":"Q","3":"N","4":"5","5":"6","6":"B","7":"R","8":"7","9":"V","A":"8","B":"1","C":"W","D":"9","E":"G","F":"H","G":"S","H":"C","I":"2","J":"3","K":"T","L":"I","M":"D","N":"E","O":"4","P":"X","Q":"K","R":"F","S":"O","T":"Y","U":"Z","V":"L","W":"P","X":"J","Y":"A","Z":"M"},"G":{"0":"0","1":"6","2":"W","3":"T","4":"B","5":"C","6":"H","7":"X","8":"D","9":"1","A":"E","B":"7","C":"2","D":"F","E":"M","F":"N","G":"Y","H":"I","I":"8","J":"9","K":"Z","L":"O","M":"J","N":"K","O":"A","P":"3","Q":"Q","R":"L","S":"U","T":"4","U":"5","V":"R","W":"V","X":"P","Y":"G","Z":"S"},"H":{"0":"X","1":"2","2":"R","3":"N","4":"8","5":"9","6":"D","7":"S","8":"A","9":"Y","A":"5","B":"3","C":"Z","D":"6","E":"H","F":"I","G":"T","H":"E","I":"4","J":"0","K":"Q","L":"J","M":"F","N":"B","O":"1","P":"U","Q":"K","R":"C","S":"O","T":"V","U":"W","V":"L","W":"P","X":"G","Y":"7","Z":"M"},"I":{"0":"C","1":"H","2":"6","3":"3","4":"N","5":"O","6":"S","7":"7","8":"P","9":"D","A":"Q","B":"I","C":"E","D":"L","E":"W","F":"X","G":"8","H":"T","I":"J","J":"K","K":"9","L":"Y","M":"U","N":"V","O":"G","P":"F","Q":"0","R":"R","S":"4","T":"A","U":"B","V":"1","W":"5","X":"Z","Y":"M","Z":"2"},"J":{"0":"Y","1":"3","2":"S","3":"O","4":"9","5":"A","6":"E","7":"T","8":"5","9":"Z","A":"6","B":"4","C":"U","D":"7","E":"I","F":"J","G":"Q","H":"F","I":"0","J":"1","K":"R","L":"G","M":"B","N":"C","O":"2","P":"V","Q":"L","R":"D","S":"P","T":"W","U":"X","V":"M","W":"N","X":"H","Y":"8","Z":"K"},"K":{"0":"N","1":"S","2":"H","3":"D","4":"Y","5":"Z","6":"3","7":"I","8":"U","9":"O","A":"V","B":"T","C":"J","D":"W","E":"7","F":"8","G":"F","H":"4","I":"P","J":"Q","K":"G","L":"5","M":"0","N":"1","O":"R","P":"K","Q":"A","R":"2","S":"E","T":"L","U":"M","V":"B","W":"C","X":"6","Y":"X","Z":"9"},"L":{"0":"8","1":"D","2":"3","3":"0","4":"J","5":"K","6":"O","7":"4","8":"L","9":"9","A":"M","B":"E","C":"A","D":"N","E":"T","F":"U","G":"5","H":"P","I":"F","J":"G","K":"6","L":"V","M":"Q","N":"R","O":"H","P":"B","Q":"X","R":"S","S":"1","T":"C","U":"7","V":"Y","W":"2","X":"W","Y":"I","Z":"Z"},"M":{"0":"E","1":"J","2":"8","3":"4","4":"P","5":"Q","6":"U","7":"9","8":"L","9":"F","A":"M","B":"K","C":"A","D":"N","E":"Y","F":"Z","G":"6","H":"V","I":"G","J":"H","K":"7","L":"W","M":"R","N":"S","O":"I","P":"B","Q":"1","R":"T","S":"5","T":"C","U":"D","V":"2","W":"3","X":"X","Y":"O","Z":"0"},"N":{"0":"W","1":"1","2":"Q","3":"N","4":"7","5":"8","6":"C","7":"R","8":"9","9":"X","A":"A","B":"2","C":"Y","D":"5","E":"G","F":"H","G":"S","H":"D","I":"3","J":"4","K":"T","L":"I","M":"E","N":"F","O":"0","P":"Z","Q":"K","R":"B","S":"O","T":"U","U":"V","V":"L","W":"P","X":"J","Y":"6","Z":"M"},"O":{"0":"O","1":"T","2":"I","3":"E","4":"Z","5":"U","6":"4","7":"F","8":"V","9":"J","A":"W","B":"P","C":"K","D":"X","E":"8","F":"5","G":"G","H":"0","I":"Q","J":"R","K":"H","L":"6","M":"1","N":"2","O":"S","P":"L","Q":"B","R":"3","S":"C","T":"M","U":"N","V":"9","W":"D","X":"7","Y":"Y","Z":"A"},"P":{"0":"U","1":"Z","2":"O","3":"K","4":"5","5":"0","6":"A","7":"L","8":"1","9":"P","A":"2","B":"V","C":"Q","D":"3","E":"E","F":"B","G":"M","H":"6","I":"W","J":"X","K":"N","L":"C","M":"7","N":"8","O":"Y","P":"R","Q":"H","R":"9","S":"I","T":"S","U":"T","V":"F","W":"J","X":"D","Y":"4","Z":"G"},"Q":{"0":"H","1":"M","2":"B","3":"7","4":"S","5":"T","6":"X","7":"C","8":"U","9":"I","A":"P","B":"N","C":"J","D":"Q","E":"1","F":"2","G":"D","H":"Y","I":"O","J":"K","K":"A","L":"3","M":"Z","N":"V","O":"L","P":"E","Q":"4","R":"W","S":"8","T":"F","U":"G","V":"5","W":"9","X":"0","Y":"R","Z":"6"},"R":{"0":"T","1":"Y","2":"N","3":"J","4":"4","5":"5","6":"9","7":"O","8":"0","9":"U","A":"1","B":"Z","C":"P","D":"2","E":"D","F":"E","G":"L","H":"A","I":"V","J":"W","K":"M","L":"B","M":"6","N":"7","O":"X","P":"Q","Q":"G","R":"8","S":"K","T":"R","U":"S","V":"H","W":"I","X":"C","Y":"3","Z":"F"},"S":{"0":"I","1":"N","2":"C","3":"8","4":"T","5":"U","6":"Y","7":"D","8":"P","9":"J","A":"Q","B":"O","C":"E","D":"R","E":"2","F":"3","G":"A","H":"Z","I":"K","J":"L","K":"B","L":"0","M":"V","N":"W","O":"M","P":"F","Q":"5","R":"X","S":"9","T":"G","U":"H","V":"6","W":"7","X":"1","Y":"S","Z":"4"},"T":{"0":"U","1":"0","2":"Q","3":"N","4":"5","5":"6","6":"B","7":"R","8":"7","9":"V","A":"8","B":"1","C":"W","D":"9","E":"G","F":"H","G":"S","H":"C","I":"2","J":"3","K":"T","L":"I","M":"D","N":"E","O":"4","P":"X","Q":"K","R":"F","S":"O","T":"Y","U":"Z","V":"L","W":"P","X":"J","Y":"A","Z":"M"},"U":{"0":"C","1":"H","2":"6","3":"2","4":"N","5":"I","6":"S","7":"3","8":"J","9":"7","A":"K","B":"D","C":"8","D":"L","E":"W","F":"T","G":"4","H":"O","I":"E","J":"F","K":"5","L":"U","M":"P","N":"Q","O":"G","P":"9","Q":"Z","R":"R","S":"0","T":"A","U":"B","V":"X","W":"1","X":"V","Y":"M","Z":"Y"},"V":{"0":"1","1":"6","2":"W","3":"T","4":"C","5":"D","6":"H","7":"X","8":"E","9":"2","A":"F","B":"7","C":"3","D":"G","E":"M","F":"N","G":"Y","H":"I","I":"8","J":"9","K":"Z","L":"O","M":"J","N":"K","O":"A","P":"4","Q":"Q","R":"L","S":"U","T":"5","U":"0","V":"R","W":"V","X":"P","Y":"B","Z":"S"},"W":{"0":"C","1":"H","2":"6","3":"3","4":"N","5":"O","6":"S","7":"7","8":"P","9":"D","A":"Q","B":"I","C":"E","D":"L","E":"W","F":"X","G":"8","H":"T","I":"J","J":"K","K":"9","L":"Y","M":"U","N":"V","O":"G","P":"F","Q":"0","R":"R","S":"4","T":"A","U":"B","V":"1","W":"5","X":"Z","Y":"M","Z":"2"},"X":{"0":"K","1":"P","2":"F","3":"C","4":"V","5":"W","6":"0","7":"G","8":"X","9":"L","A":"Y","B":"Q","C":"M","D":"Z","E":"5","F":"6","G":"H","H":"1","I":"R","J":"S","K":"I","L":"7","M":"2","N":"3","O":"T","P":"N","Q":"9","R":"4","S":"D","T":"O","U":"J","V":"A","W":"E","X":"8","Y":"U","Z":"B"},"Y":{"0":"E","1":"K","2":"A","3":"7","4":"P","5":"Q","6":"V","7":"B","8":"R","9":"F","A":"S","B":"L","C":"G","D":"T","E":"0","F":"1","G":"C","H":"W","I":"M","J":"N","K":"D","L":"2","M":"X","N":"Y","O":"O","P":"H","Q":"4","R":"Z","S":"8","T":"I","U":"J","V":"5","W":"9","X":"3","Y":"U","Z":"6"},"Z":{"0":"4","1":"A","2":"0","3":"X","4":"F","5":"G","6":"L","7":"1","8":"H","9":"5","A":"I","B":"B","C":"6","D":"J","E":"Q","F":"R","G":"2","H":"M","I":"C","J":"D","K":"3","L":"S","M":"N","N":"O","O":"E","P":"7","Q":"U","R":"P","S":"Y","T":"8","U":"9","V":"V","W":"Z","X":"T","Y":"K","Z":"W"}},"map_sort":{"0":["E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R"],"1":["F","L","X","E","V","Z","Q","S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4","H","M","N","R","6"],"2":["I","J","O","1","B","8","A","D","Y","4","5","M","N","R","6","H","L","X","E","F","Z","Q","V","W","3","S","G","K","2","7","C","P","T","U","0","9"],"3":["H","M","N","R","6","F","L","X","E","V","Z","Q","S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4"],"4":["W","3","S","G","K","2","7","C","P","T","U","0","9","I","J","O","1","B","8","A","D","Y","4","5","M","N","R","6","H","L","X","E","F","Z","Q","V"],"5":["3","S","W","2","7","G","K","T","U","0","9","C","P","O","1","B","I","J","D","Y","4","5","8","A","R","6","H","M","N","E","F","L","X","Q","V","Z"],"6":["E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R"],"7":["4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U","1","B","I","J","O"],"8":["Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X"],"9":["2","7","G","K","0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W"],"A":["H","M","N","R","6","F","L","X","E","V","Z","Q","S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4"],"B":["1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U"],"C":["H","M","N","R","6","F","L","X","E","V","Z","Q","S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4"],"D":["N","R","6","H","M","X","E","F","L","Q","V","Z","3","S","W","K","2","7","G","P","T","U","0","9","C","J","O","1","B","I","A","D","Y","4","5","8"],"E":["1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U"],"F":["1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U"],"G":["0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K"],"H":["J","O","1","B","I","A","D","Y","4","5","8","N","R","6","H","M","X","E","F","L","Q","V","Z","3","S","W","K","2","7","G","P","T","U","0","9","C"],"I":["Q","V","Z","3","S","W","2","7","G","K","T","U","0","9","C","P","O","1","B","I","J","D","Y","4","5","8","A","R","6","H","M","N","E","F","L","X"],"J":["I","J","O","1","B","8","A","D","Y","4","5","M","N","R","6","H","L","X","E","F","Z","Q","V","W","3","S","G","K","2","7","C","P","T","U","0","9"],"K":["M","N","R","6","H","L","X","E","F","Z","Q","V","W","3","S","G","K","2","7","C","P","T","U","0","9","I","J","O","1","B","8","A","D","Y","4","5"],"L":["3","S","W","2","7","G","K","U","0","9","C","P","T","1","B","I","J","O","Y","4","5","8","A","D","6","H","M","N","R","E","F","L","X","Q","V","Z"],"M":["Z","Q","V","W","3","S","G","K","2","7","C","P","T","U","0","9","I","J","O","1","B","8","A","D","Y","4","5","M","N","R","6","H","L","X","E","F"],"N":["O","1","B","I","J","D","Y","4","5","8","A","R","6","H","M","N","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","T","U","0","9","C","P"],"O":["H","M","N","R","6","F","L","X","E","V","Z","Q","S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4"],"P":["5","8","A","D","Y","4","H","M","N","R","6","F","L","X","E","V","Z","Q","S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1"],"Q":["X","E","F","L","Q","V","Z","3","S","W","K","2","7","G","P","T","U","0","9","C","J","O","1","B","I","A","D","Y","4","5","8","N","R","6","H","M"],"R":["8","A","D","Y","4","5","M","N","R","6","H","L","X","E","F","Z","Q","V","W","3","S","G","K","2","7","C","P","T","U","0","9","I","J","O","1","B"],"S":["L","X","E","F","Z","Q","V","W","3","S","G","K","2","7","C","P","T","U","0","9","I","J","O","1","B","8","A","D","Y","4","5","M","N","R","6","H"],"T":["1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U"],"U":["S","W","3","7","G","K","2","9","C","P","T","U","0","B","I","J","O","1","5","8","A","D","Y","4","H","M","N","R","6","F","L","X","E","V","Z","Q"],"V":["U","0","9","C","P","T","1","B","I","J","O","Y","4","5","8","A","D","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K"],"W":["Q","V","Z","3","S","W","2","7","G","K","T","U","0","9","C","P","O","1","B","I","J","D","Y","4","5","8","A","R","6","H","M","N","E","F","L","X"],"X":["6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","U","0","9","C","P","T","1","B","I","J","O","Y","4","5","8","A","D"],"Y":["E","F","L","X","Q","V","Z","3","S","W","2","7","G","K","0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R"],"Z":["2","7","G","K","0","9","C","P","T","U","1","B","I","J","O","4","5","8","A","D","Y","6","H","M","N","R","E","F","L","X","Q","V","Z","3","S","W"]}}`

// New
// l 限制长度
// config 由Factory{}.BuildConfig() 生成
// New(6).ToCode(123456789)
// New(6).ToId("FDSFDSG")
func New(l int, configs ...string) Factory {
	if len(configs) == 0 {
		configs = append(configs, defaultConfigString)
	}
	config := Config{}
	err := json.Unmarshal([]byte(configs[0]), &config)
	if err != nil {
		panic("id转换的模版配置不对")
	}
	config.wordsLen = len(config.Words)
	config.wordsContrary = make(map[string]int)
	for i, word := range config.Words {
		config.wordsContrary[word] = i
	}
	config.mapContrary = make(map[string]map[string]string)
	for u, m := range config.Map {
		config.mapContrary[u] = make(map[string]string)
		for u2, u3 := range m {
			config.mapContrary[u][u3] = u2
		}
	}

	return Factory{
		config: config,
		len:    l,
	}
}

func (f Factory) BuildConfig() string {
	m := make(map[string]int)
	for i, word := range words {
		m[word] = i
	}
	config := Config{
		Words: make([]string, len(words)),
		Map:   make(map[string]map[string]string),
	}
	i := 0
	for u, _ := range m {
		config.Words[i] = u
		i++
	}
	for _, word := range words {
		config.Map[word] = make(map[string]string)
		i := 0
		for u, _ := range m {
			config.Map[word][u] = words[i]
			i++
		}
	}
	s, _ := json.Marshal(config)
	return string(s)
}

// ToCode number id
// 1, 转 N -1 进制(number1);
// 2, number1取个位(k1), 选取映射模版(M1), M1模版中最后一个(Z)在number1是用不上的
// 3, number1 根据M1 进行映射替换, k1单独使用words替换， 得到code_temp
// 4, code_temp 如果长度不够，拼接(Z), 还不够拼接number1自身再截取， 得到code
func (f Factory) ToCode(id int) string {
	number1 := decimalToAny(id, f.config.wordsLen-1)                  // 整个数字进行进制转换, 把数字内部字符映射
	k2 := anyToDecimal(number1[len(number1)-1:], f.config.wordsLen-1) // 进制转换后, 最后一个转为 int
	// 一位数字进制转换
	k1 := f.config.Words[k2]
	m1 := f.config.Map[k1]

	// 除了最后一位, 其他数字直接按模板换算
	got := ""
	for _, i2 := range strings.ToUpper(number1[:len(number1)-1]) {
		got += m1[string(i2)]
	}
	// 最后一位拼接, 它保存有长度信息
	got += strings.ToUpper(k1)
	if len(got) < f.len {
		// 长度不足, 补足
		z := f.Z(k1)
		got = z + got
		m := md5.New().Sum([]byte(got))
		got = Reverse(strings.ReplaceAll(strings.ToUpper(string(m)), z, "")) + got
		got = got[len(got)-f.len:]
		got = strings.ToUpper(got)
	}
	return got
}

// Z 最后一个字符不使用的
func (f Factory) Z(k1 string) string {
	m1 := f.config.Map[k1]
	z := m1["Z"]

	return z
}

// ToId string code
// 1, 截取个位Z1, 根据wordsContrary得到K1, 也同时得到M1
// 2, M1模版中最后一个(Z), 剪切 code, 得到code_temp
// 3, code_temp根据M1反向映射得到number1
// 4, number1 转0进制得到id
func (f Factory) ToId(code string) int {
	if code == "" {
		return 0
	}
	k1 := strings.ToUpper(code[len(code)-1:])
	m1 := f.config.mapContrary[k1]
	z := f.Z(k1)
	index := strings.Index(code, z)
	code = code[index+1:]
	got := ""
	if code != "" {
		for _, i2 := range code[:len(code)-1] {
			s, ok := m1[string(i2)]
			if !ok {
				//panic("模版缺少了映射, " + code)
				return 0
			}
			got += s
		}
	}

	k2 := f.config.wordsContrary[k1]
	if k2 == 0 {
		got += "0"
	} else {
		got += decimalToAny(k2, f.config.wordsLen-1)
	}
	got = strings.ToLower(got)
	return anyToDecimal(got, f.config.wordsLen-1)
}

var tenToAny = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// map根据value找key
func finder(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 10进制转任意进制
func decimalToAny(num, n int) string {
	newNumStr := ""
	var remainder int
	var remainderString string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainderString = tenToAny[remainder]
		} else {
			remainderString = strconv.Itoa(remainder)
		}
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}

// 任意进制转10进制
func anyToDecimal(num string, n int) int {
	var newNum float64
	newNum = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(finder(value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(newNum)
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
