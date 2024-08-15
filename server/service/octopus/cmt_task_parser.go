package octopus

import (
	"regexp"
	"strings"
)

var CmtTaskParsers = map[string]CmtTaskParser{
	"xhs": XHS{},
}

type CmtTaskParser interface {
	HandleComment(cmtFrom string, comment string) string
}

type XHS struct {
}

func (xhs XHS) HandleComment(cmtFrom string, comment string) string {
	switch cmtFrom {
	case "find":
		parts := strings.Fields(comment) // 使用 Fields 自动去掉多余空格，并将字符串分割为单词
		if len(parts) < 4 {
			return comment
		}

		// 确定要截取的索引位置
		var cutIndex int
		timePattern := `^\d{1,2}:\d{2}$`
		timeSegment := parts[len(parts)-3]
		match, _ := regexp.MatchString(timePattern, timeSegment)

		if match {
			cutIndex = len(parts) - 4
		} else {
			cutIndex = len(parts) - 3
		}

		// 截取并返回结果，同时去除首尾空格
		return strings.TrimSpace(strings.Join(parts[:cutIndex], " "))

	default:
		// 直接返回原始 comment 并去除首尾空格
		return strings.TrimSpace(comment)
	}
}
