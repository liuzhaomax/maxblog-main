package business

import (
	"errors"
	"github.com/liuzhaomax/maxblog-main/src/api_article/schema"
	"time"
)

func validateArticleFields(article *schema.ArticleReq) error {
	yes, err := IsZeroValueSimpleType(article.Id)
	if err != nil {
		return err
	}
	if yes {
		return errors.New("零值校验失败: id")
	}
	yes, err = IsZeroValueSimpleType(article.Title)
	if err != nil {
		return err
	}
	if yes {
		return errors.New("零值校验失败: title")
	}
	yes, err = IsZeroValueSimpleType(article.Author)
	if err != nil {
		return err
	}
	if yes {
		return errors.New("零值校验失败: author")
	}
	return nil
}

func IsZeroValueSimpleType(value any) (bool, error) {
	if str, ok := value.(string); ok {
		return str == "", nil
	}
	switch number := value.(type) {
	case int, int64, int32, int16, int8:
		return number == 0, nil
	case uint, uint64, uint32, uint16, uint8:
		return number == 0, nil
	case float32, float64:
		return number == 0, nil
	}
	return false, errors.New("零值校验无法识别输入的格式")
}

func IsValidTimeFormat(str string) bool {
	_, err := time.Parse("2006-01-02 15:04:05.999", str)
	return err == nil
}
