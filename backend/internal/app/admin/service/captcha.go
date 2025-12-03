package service

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

// CaptchaService 验证码服务
type CaptchaService struct {
	store sync.Map // 简单内存存储
}

type captchaItem struct {
	code      string
	expiredAt time.Time
}

// NewCaptchaService 创建验证码服务
func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

// Generate 生成验证码
func (s *CaptchaService) Generate() (id string, b64s string, err error) {
	// 生成随机验证码
	code := s.randomCode(4)
	id = uuid.New().String()

	// 存储验证码，5分钟过期
	s.store.Store(id, captchaItem{
		code:      code,
		expiredAt: time.Now().Add(5 * time.Minute),
	})

	// 生成图片
	imgData := s.generateImage(code)
	b64s = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgData)

	return id, b64s, nil
}

// Verify 验证验证码
func (s *CaptchaService) Verify(id, answer string, clear bool) bool {
	if id == "" || answer == "" {
		return false
	}

	val, ok := s.store.Load(id)
	if !ok {
		return false
	}

	item := val.(captchaItem)
	if time.Now().After(item.expiredAt) {
		s.store.Delete(id)
		return false
	}

	// 不区分大小写
	matched := equalFold(item.code, answer)
	if clear {
		s.store.Delete(id)
	}
	return matched
}

func equalFold(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		c1, c2 := s1[i], s2[i]
		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 32
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 32
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

func (s *CaptchaService) randomCode(length int) string {
	const chars = "0123456789abcdefghjkmnpqrstuvwxyz"
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rnd.Intn(len(chars))]
	}
	return string(result)
}

func (s *CaptchaService) generateImage(code string) []byte {
	width, height := 120, 40
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 背景色
	bgColor := color.RGBA{240, 240, 240, 255}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, bgColor)
		}
	}

	// 简单绘制字符（使用像素点阵）
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	textColor := color.RGBA{uint8(rnd.Intn(100)), uint8(rnd.Intn(100)), uint8(rnd.Intn(100)), 255}

	for i, c := range code {
		s.drawChar(img, int(c), i*25+15, 12, textColor)
	}

	// 添加干扰点
	for i := 0; i < 100; i++ {
		x := rnd.Intn(width)
		y := rnd.Intn(height)
		img.Set(x, y, color.RGBA{uint8(rnd.Intn(255)), uint8(rnd.Intn(255)), uint8(rnd.Intn(255)), 255})
	}

	// 编码为 PNG
	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes()
}

func (s *CaptchaService) drawChar(img *image.RGBA, char int, x, y int, c color.Color) {
	// 简单的5x7像素字体
	fonts := map[int][]string{
		'0': {"01110", "10001", "10001", "10001", "10001", "10001", "01110"},
		'1': {"00100", "01100", "00100", "00100", "00100", "00100", "01110"},
		'2': {"01110", "10001", "00001", "00110", "01000", "10000", "11111"},
		'3': {"01110", "10001", "00001", "00110", "00001", "10001", "01110"},
		'4': {"00010", "00110", "01010", "10010", "11111", "00010", "00010"},
		'5': {"11111", "10000", "11110", "00001", "00001", "10001", "01110"},
		'6': {"00110", "01000", "10000", "11110", "10001", "10001", "01110"},
		'7': {"11111", "00001", "00010", "00100", "01000", "01000", "01000"},
		'8': {"01110", "10001", "10001", "01110", "10001", "10001", "01110"},
		'9': {"01110", "10001", "10001", "01111", "00001", "00010", "01100"},
		'a': {"00000", "00000", "01110", "00001", "01111", "10001", "01111"},
		'b': {"10000", "10000", "11110", "10001", "10001", "10001", "11110"},
		'c': {"00000", "00000", "01110", "10000", "10000", "10000", "01110"},
		'd': {"00001", "00001", "01111", "10001", "10001", "10001", "01111"},
		'e': {"00000", "00000", "01110", "10001", "11111", "10000", "01110"},
		'f': {"00110", "01000", "01000", "11100", "01000", "01000", "01000"},
		'g': {"00000", "01111", "10001", "10001", "01111", "00001", "01110"},
		'h': {"10000", "10000", "11110", "10001", "10001", "10001", "10001"},
		'j': {"00010", "00000", "00110", "00010", "00010", "10010", "01100"},
		'k': {"10000", "10000", "10010", "10100", "11000", "10100", "10010"},
		'm': {"00000", "00000", "11010", "10101", "10101", "10101", "10101"},
		'n': {"00000", "00000", "11110", "10001", "10001", "10001", "10001"},
		'p': {"00000", "11110", "10001", "10001", "11110", "10000", "10000"},
		'q': {"00000", "01111", "10001", "10001", "01111", "00001", "00001"},
		'r': {"00000", "00000", "10110", "11000", "10000", "10000", "10000"},
		's': {"00000", "00000", "01110", "10000", "01110", "00001", "11110"},
		't': {"01000", "01000", "11100", "01000", "01000", "01000", "00110"},
		'u': {"00000", "00000", "10001", "10001", "10001", "10011", "01101"},
		'v': {"00000", "00000", "10001", "10001", "10001", "01010", "00100"},
		'w': {"00000", "00000", "10001", "10001", "10101", "10101", "01010"},
		'x': {"00000", "00000", "10001", "01010", "00100", "01010", "10001"},
		'y': {"00000", "10001", "10001", "10001", "01111", "00001", "01110"},
		'z': {"00000", "00000", "11111", "00010", "00100", "01000", "11111"},
	}

	if font, ok := fonts[char]; ok {
		for row, line := range font {
			for col, bit := range line {
				if bit == '1' {
					// 放大2倍
					for dy := 0; dy < 2; dy++ {
						for dx := 0; dx < 2; dx++ {
							img.Set(x+col*2+dx, y+row*2+dy, c)
						}
					}
				}
			}
		}
	}
}
