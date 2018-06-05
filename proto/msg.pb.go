// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/and07/rss2website/proto/msg.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	github.com/and07/rss2website/proto/msg.proto

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
	Title       string `protobuf:"bytes,1,opt,name=Title" json:"Title,omitempty"`
	Slug        string `protobuf:"bytes,2,opt,name=Slug" json:"Slug,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	Link        string `protobuf:"bytes,4,opt,name=Link" json:"Link,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=Image" json:"Image,omitempty"`
	SourceImage string `protobuf:"bytes,6,opt,name=SourceImage" json:"SourceImage,omitempty"`
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

func init() { proto.RegisterFile("github.com/and07/rss2website/proto/msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x6a, 0x32, 0x31,
	0x14, 0x25, 0x8e, 0xfa, 0x31, 0xd7, 0x6f, 0x51, 0x42, 0x29, 0x41, 0x84, 0x8a, 0x2b, 0x29, 0x65,
	0xa6, 0x4c, 0x17, 0x96, 0x6e, 0x6b, 0x17, 0x05, 0x17, 0xa2, 0x7d, 0x81, 0x68, 0x43, 0x1a, 0x74,
	0x26, 0x32, 0x37, 0xd3, 0xe2, 0xcb, 0xf4, 0x65, 0xfa, 0x62, 0xe5, 0x26, 0xd6, 0xc9, 0xa2, 0x9b,
	0xe1, 0xfc, 0x24, 0x87, 0x73, 0x32, 0x70, 0xab, 0x8d, 0x7b, 0x6f, 0x36, 0xd9, 0xd6, 0x96, 0xb9,
	0xac, 0xde, 0xee, 0x66, 0x79, 0x8d, 0x58, 0x7c, 0xaa, 0x0d, 0x1a, 0xa7, 0xf2, 0x43, 0x6d, 0x9d,
	0xcd, 0x4b, 0xd4, 0x99, 0x47, 0x3c, 0x29, 0x51, 0x4f, 0xbe, 0x18, 0x74, 0x97, 0x16, 0x1d, 0xbf,
	0x84, 0xde, 0xab, 0x71, 0x7b, 0x25, 0xd8, 0x98, 0x4d, 0xd3, 0x55, 0x20, 0x9c, 0x43, 0x77, 0xbd,
	0x6f, 0xb4, 0xe8, 0x78, 0xd1, 0x63, 0x3e, 0x86, 0xc1, 0x5c, 0xe1, 0xb6, 0x36, 0x07, 0x67, 0x6c,
	0x25, 0x12, 0x6f, 0xc5, 0x12, 0xdd, 0x5a, 0x98, 0x6a, 0x27, 0xba, 0xe1, 0x16, 0x61, 0xca, 0x7f,
	0x29, 0xa5, 0x56, 0xa2, 0x17, 0xf2, 0x3d, 0xa1, 0xac, 0xb5, 0x6d, 0xea, 0xad, 0x0a, 0x5e, 0x3f,
	0x64, 0x45, 0xd2, 0xa4, 0x82, 0x7f, 0x2b, 0xc4, 0x85, 0x41, 0xc7, 0x6f, 0x28, 0x16, 0x9d, 0x60,
	0xe3, 0x64, 0x3a, 0x28, 0xae, 0x32, 0x9a, 0x72, 0xf2, 0x32, 0xfa, 0x3c, 0x57, 0xae, 0x3e, 0xae,
	0xfc, 0x99, 0xe1, 0x0c, 0xd2, 0xb3, 0xc4, 0x2f, 0x20, 0xd9, 0xa9, 0xe3, 0x69, 0x19, 0x41, 0x6a,
	0xf3, 0x21, 0xf7, 0x8d, 0x3a, 0x0d, 0x0b, 0xe4, 0xb1, 0xf3, 0xc0, 0x26, 0xdf, 0x0c, 0xfe, 0xd3,
	0x83, 0x2c, 0xa5, 0x56, 0x73, 0xe9, 0x24, 0x1f, 0x41, 0x4a, 0x38, 0x7e, 0x9c, 0x56, 0xf8, 0x75,
	0x43, 0xfd, 0x4e, 0xeb, 0x86, 0x79, 0x05, 0xf4, 0x88, 0xa0, 0x48, 0x7c, 0xe5, 0x91, 0xaf, 0x1c,
	0xa7, 0x67, 0xde, 0x0e, 0xc5, 0xc3, 0xd1, 0xe1, 0x13, 0x40, 0x2b, 0xfe, 0x51, 0xfd, 0x3a, 0xae,
	0x3e, 0x28, 0xd2, 0x73, 0x66, 0xb4, 0x62, 0xd3, 0xf7, 0xbf, 0xf8, 0xfe, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0xf7, 0x36, 0x80, 0x12, 0x02, 0x00, 0x00,
}
