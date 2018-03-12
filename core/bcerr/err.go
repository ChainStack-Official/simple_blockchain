package bcerr

import (
	"os"
	"errors"
	"fmt"
)

var errLang = ""
const(
	// 新增block的错误
	NewBlockIndexWrongErr = iota
	NewBlockPrevHashWrongErr
	NewBlockHashWrongErr

	// block的错误
	BlockTimestampIsBlankErr
)

func init() {
	errLang = os.Getenv("lang")
}

// 根据错误类型获取，第一个参数传错误类型，第二个参数传语言
func GetError(args ...interface{}) (err error) {
	var errCode int
	var tranOk bool
	var lang string
	if len(args) == 0 {
		return errors.New("unknown err type")
	} else if errCode, tranOk = args[0].(int); !tranOk {
		return errors.New("invalid err type")
	}
	// 如果传来语言，则返回对应语言的错误
	if len(args) > 1 {
		lang, _ = args[1].(string)
	} else {
		// 采用默认的语言
		lang = errLang
	}

	tmpErrs := allErr[errCode]
	if tmpErrs == nil {
		return fmt.Errorf("no err for code: %v", errCode)
	}
	// 根据语言环境返回错误
	switch errLang {
	case "en":
		err = errors.New(allErr[errCode][1])
	default:
		err = errors.New(allErr[errCode][0])
	}
	return fmt.Errorf("no err for lang: %v", lang)
}

// 所有配置了的错误
var allErr = map[int][]string {
	NewBlockIndexWrongErr: { "新区块的编号不正确", "block index wrong" },
	NewBlockPrevHashWrongErr: { "新区块的前PrevHash不正确", "block prev hash wrong" },
	NewBlockHashWrongErr: { "新区块的Hash不正确", "block hash wrong" },
	BlockTimestampIsBlankErr: { "区块时间戳不能为空", "block timestamp is empty" },
	//999999: { "", "" },
}


//var allErr = map[int][]error {
//	NewBlockIndexWrongErr: { errors.New("新区块的编号不正确"), errors.New("block index wrong") },
//	NewBlockPrevHashWrongErr: { errors.New("新区块的前PrevHash不正确"), errors.New("block prev hash wrong") },
//	BlockTimestampIsBlankErr: {  },
//}