package pkg

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	DiffResultMarkAdded = iota
	DiffResultMarkDeleted
	DiffResultMarkEdited
)

type DiffResult struct {
	Mark int
	Item DiffItem
}

type DiffItem interface {
	EqualTo(compare DiffItem) bool
	GetValue() string
	GetName() string
}

func NewDiffItem(diffMode string, filePath string) (DiffItem, error) {
	var item DiffItem
	var err error
	switch diffMode {
	case "CONTENT":
		item, err = NewFileContentDiffItem(filePath)
		return item, err
	case "FILENAME":
		item = NewFileNameDIffItem(filePath)
		return item, nil
	}
	return item, nil
}

type FileNameDiffItem struct {
	fileName string
}

func NewFileNameDIffItem(fileName string) FileNameDiffItem {
	return FileNameDiffItem{fileName: fileName}
}

func (f FileNameDiffItem) EqualTo(compare DiffItem) bool {
	return f.GetValue() == compare.GetValue()
}

func (f FileNameDiffItem) GetValue() string {
	return f.fileName
}

func (f FileNameDiffItem) GetName() string {
	return f.fileName
}

type FileContentDiffItem struct {
	filePath string
	data     string
}

func NewFileContentDiffItem(filePath string) (FileContentDiffItem, error) {
	item := FileContentDiffItem{filePath: filePath}
	hash, err := item.hash()
	if err != nil {
		return item, err
	}

	item.data = hash
	return item, nil

}

func (i FileContentDiffItem) EqualTo(compare DiffItem) bool {
	return i.GetValue() == compare.GetValue()
}

func (i FileContentDiffItem) GetValue() string {
	return i.data
}

func (i FileContentDiffItem) GetName() string {
	return i.filePath
}

func (i FileContentDiffItem) hash() (string, error) {
	file, err := os.Open(i.filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	modTime, err := i.getModTimeStr(file)
	if err != nil {
		return "", err
	}

	hash, err := i.getHashedDateWith(file, []byte(modTime))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (i FileContentDiffItem) getModTimeStr(f *os.File) (string, error) {
	info, err := f.Stat()
	if err != nil {
		return "", err
	}

	modTime := info.ModTime().Round(time.Second).Unix()
	return strconv.FormatInt(modTime, 10), nil
}

func (i FileContentDiffItem) getHashedDateWith(f *os.File, sumWith []byte) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", hasher.Sum(sumWith))
	return hash, nil
}

type DiffFinder interface {
	Find(prev, now []DiffItem) []DiffResult
}

type DiffHandler struct {
	handleFunc func() error
}

func (h DiffHandler) Find(prev, now []DiffItem) []DiffResult {
	var results []DiffResult

	prevMap := make(map[string]DiffItem)
	for _, item := range prev {
		prevMap[item.GetName()] = item
	}

	// 겹치는것중에 상세 데이터가 같지 않은것 >> 수정됨
	// nowMap 에만 남아있는 아이템 >> 결과적으로 추가됨
	nowMap := make(map[string]struct{})
	for _, item := range now {
		if _, found := prevMap[item.GetName()]; found {
			if !prevMap[item.GetName()].EqualTo(item) {
				results = append(results, DiffResult{Mark: DiffResultMarkEdited, Item: item})
			}
			delete(prevMap, item.GetName())
		} else {
			results = append(results, DiffResult{Mark: DiffResultMarkAdded, Item: item})
		}
		nowMap[item.GetName()] = struct{}{}
	}

	// now 에 있는 아이템을 삭제하고서도 prevMap 에 남아있는 아이템 >> 삭제됨
	for name, item := range prevMap {
		if _, found := nowMap[name]; !found {
			results = append(results, DiffResult{Mark: DiffResultMarkDeleted, Item: item})
		}
	}

	return results
}

func (h DiffHandler) Handle() error {
	return h.handleFunc()
}
