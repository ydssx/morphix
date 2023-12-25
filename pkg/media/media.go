package media

type Media interface {
	Upload(file *File) error
	Download(file *File) error
	Delete(file *File) error
	GetUrl(file *File) (string, error)
}

type File struct {
	Name    string `json:"name"`
	Ext     string `json:"ext"`
	Size    int64  `json:"size"`
	Md5     string `json:"md5"`
	Path    string `json:"path"`
	Url     string `json:"url"`
	IsImage bool   `json:"is_image"`
}

type MediaManager struct {
}

func NewMediaManager() *MediaManager {
	return &MediaManager{}
}

func (m *MediaManager) Upload(file *File) error {
	return nil
}

func (m *MediaManager) Download(file *File) error {
	return nil
}

func (m *MediaManager) Delete(file *File) error {
	return nil
}

func (m *MediaManager) GetUrl(file *File) (string, error) {
	return "", nil
}
