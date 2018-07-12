// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gitlab.com/and07/rss2website/proto/msg.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	gitlab.com/and07/rss2website/proto/msg.proto

It has these top-level messages:
	Post
	RssList
	PostPageData
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Post struct {
	Title       string   `protobuf:"bytes,1,opt,name=Title" json:"Title,omitempty"`
	Slug        string   `protobuf:"bytes,2,opt,name=Slug" json:"Slug,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	Link        string   `protobuf:"bytes,4,opt,name=Link" json:"Link,omitempty"`
	Image       string   `protobuf:"bytes,5,opt,name=Image" json:"Image,omitempty"`
	SourceImage string   `protobuf:"bytes,6,opt,name=SourceImage" json:"SourceImage,omitempty"`
	Published   int64    `protobuf:"varint,7,opt,name=Published" json:"Published,omitempty"`
	Categories  []string `protobuf:"bytes,8,rep,name=Categories" json:"Categories,omitempty"`
	SourceTitle string   `protobuf:"bytes,9,opt,name=SourceTitle" json:"SourceTitle,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Post) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Post) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Post) GetSourceImage() string {
	if m != nil {
		return m.SourceImage
	}
	return ""
}

func (m *Post) GetPublished() int64 {
	if m != nil {
		return m.Published
	}
	return 0
}

func (m *Post) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Post) GetSourceTitle() string {
	if m != nil {
		return m.SourceTitle
	}
	return ""
}

type RssList struct {
	List map[string]string `protobuf:"bytes,1,rep,name=List" json:"List,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RssList) Reset()                    { *m = RssList{} }
func (m *RssList) String() string            { return proto.CompactTextString(m) }
func (*RssList) ProtoMessage()               {}
func (*RssList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RssList) GetList() map[string]string {
	if m != nil {
		return m.List
	}
	return nil
}

type PostPageData struct {
	PageTitle string           `protobuf:"bytes,1,opt,name=PageTitle" json:"PageTitle,omitempty"`
	PageImage string           `protobuf:"bytes,2,opt,name=PageImage" json:"PageImage,omitempty"`
	Pages     map[string]*Post `protobuf:"bytes,3,rep,name=Pages" json:"Pages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PostPageData) Reset()                    { *m = PostPageData{} }
func (m *PostPageData) String() string            { return proto.CompactTextString(m) }
func (*PostPageData) ProtoMessage()               {}
func (*PostPageData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PostPageData) GetPageTitle() string {
	if m != nil {
		return m.PageTitle
	}
	return ""
}

func (m *PostPageData) GetPageImage() string {
	if m != nil {
		return m.PageImage
	}
	return ""
}

func (m *PostPageData) GetPages() map[string]*Post {
	if m != nil {
		return m.Pages
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "msg.Post")
	proto.RegisterType((*RssList)(nil), "msg.RssList")
	proto.RegisterType((*PostPageData)(nil), "msg.PostPageData")
}

func init() { proto.RegisterFile("gitlab.com/and07/rss2website/proto/msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6a, 0xf3, 0x30,
	0x10, 0xc4, 0x71, 0x7e, 0x3e, 0x6f, 0xbe, 0x43, 0x11, 0xa5, 0x88, 0x10, 0x5a, 0x93, 0x93, 0x29,
	0xc5, 0x2e, 0xee, 0x21, 0xa5, 0xd7, 0xa4, 0x87, 0x42, 0x0e, 0xc1, 0xe9, 0x0b, 0xc8, 0x89, 0x50,
	0x44, 0xfc, 0x13, 0xbc, 0x72, 0x4b, 0x9e, 0xaf, 0xef, 0x55, 0x8a, 0xa4, 0x34, 0xd6, 0xa1, 0x17,
	0xb3, 0x3b, 0xa3, 0x1d, 0x66, 0x06, 0xc3, 0x83, 0x90, 0x6a, 0xdf, 0xe6, 0xf1, 0xb6, 0x2e, 0x13,
	0x56, 0xed, 0x1e, 0xe7, 0x49, 0x83, 0x98, 0x7e, 0xf2, 0x1c, 0xa5, 0xe2, 0xc9, 0xb1, 0xa9, 0x55,
	0x9d, 0x94, 0x28, 0x62, 0x33, 0x11, 0xbf, 0x44, 0x31, 0xfb, 0xf6, 0xa0, 0xbf, 0xae, 0x51, 0x91,
	0x6b, 0x18, 0xbc, 0x4b, 0x55, 0x70, 0xea, 0x85, 0x5e, 0x14, 0x64, 0x76, 0x21, 0x04, 0xfa, 0x9b,
	0xa2, 0x15, 0xb4, 0x67, 0x40, 0x33, 0x93, 0x10, 0xc6, 0x4b, 0x8e, 0xdb, 0x46, 0x1e, 0x95, 0xac,
	0x2b, 0xea, 0x1b, 0xca, 0x85, 0xf4, 0xd5, 0x4a, 0x56, 0x07, 0xda, 0xb7, 0x57, 0x7a, 0xd6, 0xfa,
	0x6f, 0x25, 0x13, 0x9c, 0x0e, 0xac, 0xbe, 0x59, 0xb4, 0xd6, 0xa6, 0x6e, 0x9b, 0x2d, 0xb7, 0xdc,
	0xd0, 0x6a, 0x39, 0x10, 0x99, 0x42, 0xb0, 0x6e, 0xf3, 0x42, 0xe2, 0x9e, 0xef, 0xe8, 0x28, 0xf4,
	0x22, 0x3f, 0xeb, 0x00, 0x72, 0x0b, 0xb0, 0x60, 0x8a, 0x8b, 0xba, 0x91, 0x1c, 0xe9, 0xbf, 0xd0,
	0x8f, 0x82, 0xcc, 0x41, 0x3a, 0x7d, 0x9b, 0x2d, 0x70, 0xf5, 0x0d, 0x34, 0xab, 0x60, 0x94, 0x21,
	0xae, 0x24, 0x2a, 0x72, 0xaf, 0x6d, 0xa3, 0xa2, 0x5e, 0xe8, 0x47, 0xe3, 0xf4, 0x26, 0xd6, 0x55,
	0x9d, 0xb9, 0x58, 0x7f, 0x5e, 0x2b, 0xd5, 0x9c, 0x32, 0xf3, 0x66, 0x32, 0x87, 0xe0, 0x02, 0x91,
	0x2b, 0xf0, 0x0f, 0xfc, 0x74, 0x6e, 0x4e, 0x8f, 0x3a, 0xed, 0x07, 0x2b, 0x5a, 0x7e, 0x2e, 0xce,
	0x2e, 0x2f, 0xbd, 0x67, 0x6f, 0xf6, 0xe5, 0xc1, 0x7f, 0x5d, 0xf8, 0x9a, 0x09, 0xbe, 0x64, 0x8a,
	0x99, 0x80, 0x4c, 0x70, 0xb7, 0xfc, 0x0e, 0xf8, 0x65, 0x6d, 0x3d, 0xbd, 0x8e, 0xb5, 0xe5, 0xa4,
	0x30, 0xd0, 0x0b, 0x52, 0xdf, 0x58, 0x9e, 0x1a, 0xcb, 0xae, 0x7a, 0x6c, 0x68, 0x6b, 0xdc, 0x3e,
	0x9d, 0x2c, 0x00, 0x3a, 0xf0, 0x0f, 0xeb, 0x77, 0xae, 0xf5, 0x71, 0x1a, 0x5c, 0x34, 0x9d, 0x14,
	0xf9, 0xd0, 0xfc, 0x42, 0x4f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x0a, 0xa5, 0xe0, 0x72,
	0x02, 0x00, 0x00,
}
