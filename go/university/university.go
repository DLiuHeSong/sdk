package university

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type University struct {
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type SDK struct {
	data []University
}

// New 创建 SDK 实例，载入数据
func New(jsonFilePath string) (*SDK, error) {
	bytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}
	var data []University
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &SDK{data: data}, nil
}

// ListAll 返回所有学校
func (sdk *SDK) ListAll() []University {
	return sdk.data
}

// SearchByName 按名称模糊查询
func (sdk *SDK) SearchByName(keyword string) []University {
	var results []University
	lowerKey := strings.ToLower(keyword)
	for _, u := range sdk.data {
		if strings.Contains(strings.ToLower(u.Name), lowerKey) {
			results = append(results, u)
		}
	}
	return results
}

// SearchByCity 按城市查询
func (sdk *SDK) SearchByCity(city string) []University {
	var results []University
	lowerCity := strings.ToLower(city)
	for _, u := range sdk.data {
		if strings.ToLower(u.City) == lowerCity {
			results = append(results, u)
		}
	}
	return results
}

// SearchByProvince 按省份查询
func (sdk *SDK) SearchByProvince(province string) []University {
	var results []University
	lowerProvince := strings.ToLower(province)
	for _, u := range sdk.data {
		if strings.ToLower(u.Province) == lowerProvince {
			results = append(results, u)
		}
	}
	return results
}
