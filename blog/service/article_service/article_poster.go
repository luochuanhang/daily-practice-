package article_service

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"

	"lianxi/blog/pkg/file"
	"lianxi/blog/pkg/qrcode"
	"lianxi/blog/pkg/setting"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string, article *Article, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
		Qr:         qr,
	}
}

//获取海报标志
func GetPosterFlag() string {
	return "poster"
}

//检查是否存在
func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}

	return true
}

//打开合并后的文件
func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	//打开文件
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubTitle string
	X1       int
	Y1       int
	Size1    float64
}

//画海报
func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	//保存地址
	fontSource := setting.AppSetting.RuntimeRootPath + setting.AppSetting.FontSavePath + fontName
	//读取文件
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}
	//ParseFont只是从freetype/truetype包中调用Parse函数。这里提供了它，
	//这样导入这个包的代码就不需要同时包含freetype/truetype包。
	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}
	//NewContext创建一个新的Context。
	fc := freetype.NewContext()
	//SetDPI设置屏幕分辨率，单位是点每英寸。
	fc.SetDPI(72)
	//SetFont设置用于绘制文本的字体。
	fc.SetFont(trueTypeFont)
	//SetFontSize以点为单位设置字体大小(如“12点字体”)。
	fc.SetFontSize(d.Size0)
	//Bounds返回At可以返回非零颜色的域。边界不一定包含点(0,0)
	//设置用于绘图的剪辑矩形。
	fc.SetClip(d.JPG.Bounds())
	//SetSrc设置绘制操作的源图像。这是典型的图像，统一。
	fc.SetDst(d.JPG)
	//SetSrc设置绘制操作的源图像。这是典型的图像，统一。
	fc.SetSrc(image.Black)
	//Pt将以像素为单位的坐标对转换为固定坐标。Point26_6坐标对测量固定。Int26_6单位。
	pt := freetype.Pt(d.X0, d.Y0)
	/*
		DrawString在p处绘制s，并返回按文本范围前进的p。文本的放置是为了使第一个字符s的em方形的
		左边缘与基线在p处相交。大多数受影响的像素将在该点的上方和右侧，但有些可能在该点的下方或左侧。
		例如，用斜体绘制以'J'开头的字符串可能会影响该点下方和左侧的像素。P是固定的。Point26_6，
		因此可以表示亚像素位置。ningld Ti+le净吨
	*/
	_, err = fc.DrawString(d.Title, pt)
	if err != nil {
		return err
	}

	fc.SetFontSize(d.Size1)
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1))
	if err != nil {
		return err
	}
	//Encode使用给定的选项以JPEG 4:2:0基线格式将图像m写到w。如果传入nil *Options，则使用默认参数。
	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}

	return nil
}

//生成文章海报二维码
func (a *ArticlePosterBg) Generate() (string, string, error) {
	//文件地址
	fullPath := qrcode.GetQrCodeFullPath()
	//生成二维码
	fileName, path, err := a.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}

	if !a.CheckMergedImage(path) {
		//打开文件
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		//关闭连接
		defer mergedF.Close()
		//打开文章海报
		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()
		//打开二维码
		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()
		//获取文章海报图像
		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		//获取二维码图像
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}
		//newgba返回一个带有给定边界的新的RGBA图像。
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
		//绘制调用带有nil掩码的DrawMask。
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)

		err = a.DrawPoster(&DrawText{
			JPG:    jpg,
			Merged: mergedF,

			Title: "Golang Gin 系列文章",
			X0:    80,
			Y0:    160,
			Size0: 42,

			SubTitle: "---煎鱼",
			X1:       320,
			Y1:       220,
			Size1:    36,
		}, "msyhbd.ttc")

		if err != nil {
			return "", "", err
		}
	}

	return fileName, path, nil
}
