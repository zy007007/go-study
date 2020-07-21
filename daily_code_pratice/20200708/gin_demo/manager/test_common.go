package manager

import (
	"gin_demo/manager/model"
)

func AddCommon(name string) (int64, error) {
	common := &model.Common{Name:name}

	test, err := common.Insert()
	if err != nil {
		return 0, nil
	}
	return test, err
}
