//go:generate go run bindings/generate.go
package raylib

import (
	"unsafe"
)

const (
	VERSION_MAJOR = 5
	VERSION_MINOR = 5
	VERSION_PATCH = 0
	VERSION       = "5.5.0"
)

type PixelFormat uint32

const (
	PixelFormatUncompressedGrayscale    PixelFormat = 1 + iota // 8 bit per pixel (no alpha)
	PixelFormatUncompressedGrayAlpha                           // 8*2 bpp (2 channels)
	PixelFormatUncompressedR5g6b5                              // 16 bpp
	PixelFormatUncompressedR8g8b8                              // 24 bpp
	PixelFormatUncompressedR5g5b5a1                            // 16 bpp (1 bit alpha)
	PixelFormatUncompressedR4g4b4a4                            // 16 bpp (4 bit alpha)
	PixelFormatUncompressedR8g8b8a8                            // 32 bpp
	PixelFormatUncompressedR32                                 // 32 bpp (1 channel - float)
	PixelFormatUncompressedR32g32b32                           // 32*3 bpp (3 channels - float)
	PixelFormatUncompressedR32g32b32a32                        // 32*4 bpp (4 channels - float)
	PixelFormatUncompressedR16                                 // 16 bpp (1 channel - half float)
	PixelFormatUncompressedR16g16b16                           // 16*3 bpp (3 channels - half float)
	PixelFormatUncompressedR16g16b16a16                        // 16*4 bpp (4 channels - half float)
	PixelFormatCompressedDxt1Rgb                               // 4 bpp (no alpha)
	PixelFormatCompressedDxt1Rgba                              // 4 bpp (1 bit alpha)
	PixelFormatCompressedDxt3Rgba                              // 8 bpp
	PixelFormatCompressedDxt5Rgba                              // 8 bpp
	PixelFormatCompressedEtc1Rgb                               // 4 bpp
	PixelFormatCompressedEtc2Rgb                               // 4 bpp
	PixelFormatCompressedEtc2EacRgba                           // 8 bpp
	PixelFormatCompressedPvrtRgb                               // 4 bpp
	PixelFormatCompressedPvrtRgba                              // 4 bpp
	PixelFormatCompressedAstc4x4Rgba                           // 8 bpp
	PixelFormatCompressedAstc8x8Rgba                           // 2 bpp
)

type CameraProjection uint32

const (
	CameraPerspective CameraProjection = iota
	CameraOrthographic
)

type (
	// Vector2, 2 components
	Vector2 struct {
		X float32
		Y float32
	}

	// Vector3, 3 components
	Vector3 struct {
		X float32
		Y float32
		Z float32
	}

	// Vector4, 4 components
	Vector4 struct {
		X float32
		Y float32
		Z float32
		W float32
	}

	// Quaternion, 4 components (Vector4 alias)
	Quaternion = Vector4

	// Color, 4 components, R8G8B8A8 (32bit)
	Color struct {
		R uint8
		G uint8
		B uint8
		A uint8
	}

	// Rectangle, 4 components
	Rectangle struct {
		x      float32 // Rectangle top-left corner position x
		y      float32 // Rectangle top-left corner position y
		width  float32 // Rectangle width
		height float32 // Rectangle height
	}

	// Image, pixel data stored in CPU memory (RAM)
	Image struct {
		Data    unsafe.Pointer // Image raw data
		Width   int32          // Image base width
		Height  int32          // Image base height
		Mipmaps int32          // Mipmap levels, 1 by default
		Format  PixelFormat    // Data format (PixelFormat type)
	}

	// Camera, defines position/orientation in 3d space
	Camera3D struct {
		Position   Vector3          // Camera position
		Target     Vector3          // Camera target it looks-at
		Up         Vector3          // Camera up vector (rotation over its axis)
		Fovy       float32          // Camera field-of-view aperture in Y (degrees) in perspective, used as near plane width in orthographic
		Projection CameraProjection // Camera projection: CAMERA_PERSPECTIVE or CAMERA_ORTHOGRAPHIC
	}

	// Camera type fallback, defaults to Camera3D
	Camera = Camera3D

	// Camera2D, defines position/orientation in 2d space
	Camera2D struct {
		Offset   Vector2 // Camera offset (displacement from target)
		Target   Vector2 // Camera target (rotation and zoom origin)
		Rotation float32 // Camera rotation in degrees
		Zoom     float32 // Camera zoom (scaling), should be 1.0f by default
	}

	// Texture, tex data stored in GPU memory (VRAM)
	Texture struct {
		Id      uint32 // OpenGL texture id
		Width   int32  // Texture base width
		Height  int32  // Texture base height
		Mipmaps int32  // Mipmap levels, 1 by default
		Format  int32  // Data format (PixelFormat type)
	}

	// Texture2D, same as Texture
	Texture2D = Texture

	// TextureCubemap, same as Texture
	TextureCubemap = Texture

	// RenderTexture, fbo for texture rendering
	RenderTexture struct {
		Id      uint32  // OpenGL framebuffer object id
		Texture Texture // Color buffer attachment texture
		Depth   Texture // Depth buffer attachment texture
	}

	// RenderTexture2D, same as RenderTexture
	RenderTexture2D = RenderTexture

	// Shader
	Shader struct {
		Id   uint32 // Shader program id
		Locs *int32 // Shader locations array (RL_MAX_SHADER_LOCATIONS)
	}
)
