package storage

import (
	"os"

	"k8s.io/klog"
)

func Write(file string, data []byte) error {
	f, err := os.Create(file)
	if err != nil {
		klog.Errorf("create file error ")
		return err
	}
	_, err = f.Write(data)
	return err
}
