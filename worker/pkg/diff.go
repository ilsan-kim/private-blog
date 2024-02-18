package pkg

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
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
	GetTime() time.Time
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
	fileName    string
	createdTime time.Time
}

func NewFileNameDIffItem(fileName string) FileNameDiffItem {
	res := FileNameDiffItem{fileName: fileName}
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return res
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return res
	}

	res.createdTime = fileInfo.ModTime()
	return res
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

func (f FileNameDiffItem) GetTime() time.Time {
	return f.createdTime
}

type FileContentDiffItem struct {
	filePath    string
	data        string
	createdTime time.Time
}

func NewFileContentDiffItem(filePath string) (FileContentDiffItem, error) {
	item := FileContentDiffItem{filePath: filePath}
	file, err := os.Open(filePath)
	if err != nil {
		return FileContentDiffItem{}, err
	}
	defer file.Close()

	hash, err := item.hash(file)
	if err != nil {
		return item, err
	}

	item.createdTime, err = item.getModTime(file)
	if err != nil {
		log.Println("can not get createdTime for ", file)
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

func (i FileContentDiffItem) GetTime() time.Time {
	return i.createdTime
}

func (i FileContentDiffItem) hash(f *os.File) (string, error) {
	modTime, err := i.getModTimeStr(f)
	if err != nil {
		return "", err
	}

	hash, err := i.getHashedDataWith(f, []byte(modTime))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (i FileContentDiffItem) getModTime(f *os.File) (time.Time, error) {
	info, err := f.Stat()
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), err
}

func (i FileContentDiffItem) getModTimeStr(f *os.File) (string, error) {
	modTime, err := i.getModTime(f)
	if err != nil {
		return "", err
	}

	modTimeInt := modTime.Round(time.Second).Unix()
	return strconv.FormatInt(modTimeInt, 10), nil
}

func (i FileContentDiffItem) getHashedDataWith(f *os.File, sumWith []byte) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", hasher.Sum(sumWith))
	return hash, nil
}

type FileDiffHandler struct {
	DiffMode string
	DirPath  string
	prev     []DiffItem
}

func (h FileDiffHandler) getDiffItem(filePath string) (DiffItem, error) {
	path := fmt.Sprintf("%s/%s", h.DirPath, filePath)

	// TODO: diffMode 의 상수화
	diffItem, err := NewDiffItem(h.DiffMode, path)
	return diffItem, err
}

func (h FileDiffHandler) listDirectory() ([]os.DirEntry, error) {
	entries, err := os.ReadDir(h.DirPath)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (h FileDiffHandler) getItemsAsDiffItem() ([]DiffItem, error) {
	res := make([]DiffItem, 0)

	files, err := h.listDirectory()
	if err != nil {
		log.Printf("can not list files in directory %s\n", h.DirPath)
		return nil, err
	}

	for _, file := range files {
		postDiffItem, err := h.getDiffItem(file.Name())
		if err != nil {
			log.Printf("can not initialize FileContentDiffItem for %s\n", file.Name())
			continue
		}

		res = append(res, postDiffItem)
	}

	return res, nil
}

func (h FileDiffHandler) find(prev, now []DiffItem) []DiffResult {
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

func (h FileDiffHandler) HandleDiff(handlerFunc func(diffs []DiffResult) error) (FileDiffHandler, error) {
	var err error
	infos, err := h.getItemsAsDiffItem()
	log.Printf("prev > %d, now > %d\n", len(h.prev), len(infos))
	if err != nil {
		return h, err
	}

	diffs := h.find(h.prev, infos)
	if len(diffs) > 0 {
		err = handlerFunc(diffs)
		if err != nil {
			return h, err
		}
	}

	h.prev, err = h.getItemsAsDiffItem()
	return h, err
}
