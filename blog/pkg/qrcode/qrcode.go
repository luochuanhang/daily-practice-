package qrcode

import (
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"lianxi/blog/pkg/file"
	"lianxi/blog/pkg/setting"
	"lianxi/blog/pkg/util"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

// NewQrCode初始化实例
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

// GetQrCodePath获取保存路径
func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}

// GetQrCodeFullPath获取完整保存路径
func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}

// GetQrCodeFullUrl获取完整访问路径
func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

// GetQrCodeFileName获取gr文件名
func GetQrCodeFileName(value string) string {
	//对名字进行MD5编码
	return util.EncodeMD5(value)
}

// GetQrCodeExt获取qr文件后缀
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// Encode生成二维码
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) {
		//Encode返回具有给定内容、纠错级别的QR条形码，并使用给定的编码
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}
		//缩放返回带有给定宽度和高度的调整大小的条形码。
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()
		//Encode使用给定的选项以JPEG 4:2:0基线格式将图像m写到w。如果传入nil *Options，则使用默认参数。
		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
