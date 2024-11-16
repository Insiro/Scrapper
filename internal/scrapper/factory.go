package scrapper

import (
	"Scrapper/internal/app"
	"Scrapper/internal/dto"
	"Scrapper/internal/entity/enum"
	"fmt"
	"net/url"
)

type factoryStruct struct {
	AbsScrapper
	Instance Scrapper
	Url      *url.URL
	Args     ScrapArgs
}

var _ Scrapper = (*factoryStruct)(nil)

func (f *factoryStruct) PreprocessURL(rawURL string) (*url.URL, error) {
	return f.Instance.PreprocessURL(rawURL)
}

func (f *factoryStruct) Scrap(args *ScrapArgs) (dto.ScrapCreate, error) {
	arg := args
	if args == nil {
		arg = &f.Args
	}
	result, err := f.Instance.Scrap(arg)

	return result, err
}

func Factory(urlStr string, pageType *enum.PageType, typeName *string, config *app.Config) (*factoryStruct, error) {
	var instance Scrapper
	var resolvedType enum.PageType

	if pageType == nil {
		if typeName == nil {
			hostname, _ := url.Parse(urlStr)
			resolvedType, _ = enum.FromHost(hostname.Hostname())
		} else {
			resolvedType, _ = enum.FromHost(*typeName)
		}
	} else {
		resolvedType = *pageType
	}
	base := AbsScrapper{resolvedType, config}
	switch resolvedType {
	case enum.Twitter:
		instance = &implTwitter{base}
	case enum.HoyoLab:
		instance = &implHoyolab{base}
	case enum.HoyoLink:
		base.PageType = enum.HoyoLab
		instance = &implHoyolink{implHoyolab{base}}
	case enum.Instagram:
		instance = &implInsta{base}
	default:
		return nil, fmt.Errorf("unsupported page type")
	}

	// URL 전처리 및 ScrapArgs 생성
	processedURL, err := instance.PreprocessURL(urlStr)
	if err != nil {
		return nil, err
	}
	args := instance.GenArgs(processedURL)

	return &factoryStruct{
		AbsScrapper: base,
		Instance:    instance,
		Url:         processedURL,
		Args:        args,
	}, nil
}
