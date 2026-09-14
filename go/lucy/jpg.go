package lucy

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
)

// RadarJPG / ScatterJPG / BarsJPG encode the same rasters as PNG, as JPEG.
func RadarJPG(title string, series []RadarSeries, quality int) ([]byte, error) {
	return rasterJPG(func() (*image.RGBA, error) {
		b, err := RadarPNG(title, series)
		if err != nil {
			return nil, err
		}
		return decodePNG(b)
	}, quality)
}

func ScatterJPG(title, xLabel, yLabel string, pts []ScatterPoint, quality int) ([]byte, error) {
	return rasterJPG(func() (*image.RGBA, error) {
		b, err := ScatterPNG(title, xLabel, yLabel, pts)
		if err != nil {
			return nil, err
		}
		return decodePNG(b)
	}, quality)
}

func BarsJPG(title string, board LPD, max, quality int) ([]byte, error) {
	return rasterJPG(func() (*image.RGBA, error) {
		b, err := BarsPNG(title, board, max)
		if err != nil {
			return nil, err
		}
		return decodePNG(b)
	}, quality)
}

func decodePNG(b []byte) (*image.RGBA, error) {
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba, nil
	}
	bnds := img.Bounds()
	out := image.NewRGBA(bnds)
	for y := bnds.Min.Y; y < bnds.Max.Y; y++ {
		for x := bnds.Min.X; x < bnds.Max.X; x++ {
			out.Set(x, y, img.At(x, y))
		}
	}
	return out, nil
}

func rasterJPG(mk func() (*image.RGBA, error), quality int) ([]byte, error) {
	if quality <= 0 {
		quality = 85
	}
	img, err := mk()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
