package formdata

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime/multipart"
)

func MapToBase64(files *[]multipart.FileHeader) ([]string, error) {
	base64Strings := make([]string, len(*files))
	// Обрабатываем каждый файл
	for i, file := range *files {
		// Открываем файл
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		// Читаем содержимое файла
		content, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}

		// Кодируем содержимое файла в base64
		base64String := base64.StdEncoding.EncodeToString(content)

		// Сохраняем строку в слайс
		base64Strings[i] = base64String
	}
	return base64Strings, nil
}

func MapToCarParts(data []string, model interface{}) error {
	var byteSlice [][]byte
	for _, s := range data {
		byteSlice = append(byteSlice, []byte(s))
	}
	byteRes := bytes.Join(byteSlice, []byte{})

	err := json.Unmarshal(byteRes, &model)

	if err != nil {
		return err
	}
	return nil
}
