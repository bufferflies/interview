package storage

import (
	"bytes"
	"io/ioutil"
	"os"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"
	error2 "git.code.oa.com/geeker/awesome-work/pingcap/error"
)

// Search
func Search(file string, key uint64) ([]byte, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	left, right := int64(0), info.Size()-16
	var middle int64
	var middleValue uint64
	for left <= right {
		middle = ((left>>4 + right>>4) / 2) << 4
		middleValue = common.BytesToUint64(body[middle : middle+8])
		if middleValue > key {
			right = middle - 16
		} else if middleValue < key {
			left = middle + 16
		} else {
			if middle-16 < 0 {
				break
			}
			middlePre := common.BytesToUint64(body[middle-16 : middle-8])
			if middlePre != key {
				break
			} else {
				right = middle - 16
			}
		}
	}
	if middleValue != key {
		return nil, error2.KeyIsNotExistError
	}
	ret := bytes.Buffer{}
	//ret.Write(body[middle : middle+16])
	for ; common.BytesToUint64(body[middle:middle+8]) == key; middle = middle + 16 {
		ret.Write(body[middle : middle+16])
	}

	return ret.Bytes(), nil
}
