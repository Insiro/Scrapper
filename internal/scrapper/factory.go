package scrapper

import (
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

func (f factoryStruct) PreprocessURL(rawURL string) (*url.URL, error) {
	return f.Instance.PreprocessURL(rawURL)
}

func (f factoryStruct) Scrap(args *ScrapArgs) (dto.ScrapCreate, error) {
	arg := args
	if args == nil {
		arg = &f.Args
	}
	return f.Instance.Scrap(arg)
}

func Factory(urlStr string, pageType *enum.PageType, typeName *string) (*factoryStruct, error) {
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
	scrapper := AbsScrapper{resolvedType}
	in := implHoyolab{scrapper}
	switch resolvedType {
	//	case ptype.Twitter:
	//		instance = &ImplTwitter{}
	case enum.HoyoLab:
		instance = &in
		//	case ptype.HoyoLink:
		//		instance = &ImplHoyoLink{}
		//	case ptype.Instagram:
		//		instance = &ImplInstagram{}
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
		Instance: instance,
		Url:      processedURL,
		Args:     args,
	}, nil
}
