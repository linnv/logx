package logx

import (
	"os"
	"path"
)

func checkFileAvailable(filepath string) (*os.File, error) {
	fd, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		if os.IsPermission(err) {
			err = os.Chmod(filepath, 0666)
			if err != nil {
				return nil, err
			}
			return checkFileAvailable(filepath)
		}
		return nil, err
	}
	return fd, nil
}

func checkDirAvailable(filepath string) error {
	_, err := os.Stat(filepath)
	if err == nil {
	} else if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(filepath), 0766)
		if err != nil {
			return err
		}
	} else if os.IsPermission(err) {
		err = os.Chmod(path.Dir(filepath), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
