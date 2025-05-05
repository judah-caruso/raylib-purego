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

// System/Window config flags
// NOTE: Every bit registers one state (use it with bit masks)
// By default all flags are set to 0
type ConfigFlags uint32

const (
	FlagVsyncHint              = ConfigFlags(0x00000040) // Set to try enabling V-Sync on GPU
	FlagFullscreenMode         = ConfigFlags(0x00000002) // Set to run program in fullscreen
	FlagWindowResizable        = ConfigFlags(0x00000004) // Set to allow resizable window
	FlagWindowUndecorated      = ConfigFlags(0x00000008) // Set to disable window decoration (frame and buttons)
	FlagWindowHidden           = ConfigFlags(0x00000080) // Set to hide window
	FlagWindowMinimized        = ConfigFlags(0x00000200) // Set to minimize window (iconify)
	FlagWindowMaximized        = ConfigFlags(0x00000400) // Set to maximize window (expanded to monitor)
	FlagWindowUnfocused        = ConfigFlags(0x00000800) // Set to window non focused
	FlagWindowTopmost          = ConfigFlags(0x00001000) // Set to window always on top
	FlagWindowAlwaysRun        = ConfigFlags(0x00000100) // Set to allow windows running while minimized
	FlagWindowTransparent      = ConfigFlags(0x00000010) // Set to allow transparent framebuffer
	FlagWindowHighdpi          = ConfigFlags(0x00002000) // Set to support HighDPI
	FlagWindowMousePassthrough = ConfigFlags(0x00004000) // Set to support mouse passthrough, only supported when FLAG_WINDOW_UNDECORATED
	FlagBorderlessWindowedMode = ConfigFlags(0x00008000) // Set to run program in borderless windowed mode
	FlagMsaa4xHint             = ConfigFlags(0x00000020) // Set to try enabling MSAA 4X
	LagInterlacedHint          = ConfigFlags(0x00010000) // Set to try enabling interlaced video format (for V3D)
)

// Pixel formats
// NOTE: Support depends on OpenGL version and platform
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

// Camera projection
type CameraProjection uint32

const (
	CameraPerspective CameraProjection = iota
	CameraOrthographic
)

// N-patch layout
type NPatchLayout uint32

const (
	NPatchNinePatch           NPatchLayout = iota // Npatch layout: 3x3 tiles
	NPatchThreePatchVertical                      // Npatch layout: 1x3 tiles
	NPatchThreePatchHoizontal                     // Npatch layout: 3x1 tiles
)

type (
	// Vector2, 2 components
	Vector2 struct {
		X float32 // Vector x component
		Y float32 // Vector y component
	}

	// Vector3, 3 components
	Vector3 struct {
		X float32 // Vector x component
		Y float32 // Vector y component
		Z float32 // Vector z component
	}

	// Vector4, 4 components
	Vector4 struct {
		X float32 // Vector x component
		Y float32 // Vector y component
		Z float32 // Vector z component
		W float32 // Vector w component
	}

	// Quaternion, 4 components (Vector4 alias)
	Quaternion = Vector4

	// @todo(judah): the bindings generator doesn't support comma separated field declarations

	// Matrix, 4x4 components, column major, OpenGL style, right-handed
	Matrix struct {
		// Matrix first row (4 components)
		M0  float32
		M4  float32
		M8  float32
		M12 float32

		// Matrix second row (4 components)
		M1  float32
		M5  float32
		M9  float32
		M13 float32

		// Matrix third row (4 components)
		M2  float32
		M6  float32
		M10 float32
		M14 float32

		// Matrix fourth row (4 components)
		M3  float32
		M7  float32
		M11 float32
		M15 float32
	}

	// Color, 4 components, R8G8B8A8 (32bit)
	Color struct {
		R uint8 // Color red value
		G uint8 // Color green value
		B uint8 // Color blue value
		A uint8 // Color alpha value
	}

	// Rectangle, 4 components
	Rectangle struct {
		X      float32 // Rectangle top-left corner position x
		Y      float32 // Rectangle top-left corner position y
		Width  float32 // Rectangle width
		Height float32 // Rectangle height
	}

	// Image, pixel data stored in CPU memory (RAM)
	Image struct {
		Data    unsafe.Pointer // Image raw data
		Width   int32          // Image base width
		Height  int32          // Image base height
		Mipmaps int32          // Mipmap levels, 1 by default
		Format  PixelFormat    // Data format (PixelFormat type)
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

	// NPatchInfo, n-patch layout info
	NPatchInfo struct {
		Source Rectangle    // Texture source rectangle
		Left   int32        // Left border offset
		Top    int32        // Top border offset
		Right  int32        // Right border offset
		Bottom int32        // Bottom border offset
		Layout NPatchLayout // Layout of the n-patch: 3x3, 1x3 or 3x1
	}

	// GlyphInfo, font characters glyphs info
	GlyphInfo struct {
		Value    int32 // Character value (Unicode)
		OffsetX  int32 // Character offset X when drawing
		OffsetY  int32 // Character offset Y when drawing
		AdvanceX int32 // Character advance position X
		Image    Image // Character image data
	}

	// Font, font texture and GlyphInfo array data
	Font struct {
		BaseSize     int32      // Base size (default chars height)
		GlyphCount   int32      // Number of glyph characters
		GlyphPadding int32      // Padding around the glyph characters
		Texture      Texture2D  // Texture atlas containing the glyphs
		Recs         *Rectangle // Rectangles in texture for the glyphs
		Glyphs       *GlyphInfo // Glyphs info data
	}

	// Camera, defines position/orientation in 3d space
	Camera3D struct {
		Position   Vector3          // Camera position
		Target     Vector3          // Camera target it looks-at
		Up         Vector3          // Camera up vector (rotation over its axis)
		FovY       float32          // Camera field-of-view aperture in Y (degrees) in perspective, used as near plane width in orthographic
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

	// Mesh, vertex data and vao/vbo
	Mesh struct {
		VertexCount   int32 // Number of vertices stored in arrays
		TriangleCount int32 // Number of triangles stored (indexed or not)

		// Vertex attributes data
		Vertices   *float32 // Vertex position (XYZ - 3 components per vertex) (shader-location = 0)
		TexCoords  *float32 // Vertex texture coordinates (UV - 2 components per vertex) (shader-location = 1)
		TexCoords2 *float32 // Vertex texture second coordinates (UV - 2 components per vertex) (shader-location = 5)
		Normals    *float32 // Vertex normals (XYZ - 3 components per vertex) (shader-location = 2)
		Tangents   *float32 // Vertex tangents (XYZW - 4 components per vertex) (shader-location = 4)
		Colors     *uint8   // Vertex colors (RGBA - 4 components per vertex) (shader-location = 3)
		Indices    *uint16  // Vertex indices (in case vertex data comes indexed)

		// Animation vertex data
		AnimVertices *float32 // Animated vertex positions (after bones transformations)
		AnimNormals  *float32 // Animated normals (after bones transformations)
		BoneIds      *uint8   // Vertex bone ids, max 255 bone ids, up to 4 bones influence by vertex (skinning) (shader-location = 6)
		BoneWeights  *float32 // Vertex bone weight, up to 4 bones influence by vertex (skinning) (shader-location = 7)
		BoneMatrices *Matrix  // Bones animated transformation matrices
		BoneCount    int32    // Number of bones

		// OpenGL identifiers
		VaoId uint32  // OpenGL Vertex Array Object id
		VboId *uint32 // OpenGL Vertex Buffer Objects id (default vertex data)
	}

	// Shader
	Shader struct {
		Id   uint32 // Shader program id
		Locs *int32 // Shader locations array (RL_MAX_SHADER_LOCATIONS)
	}

	// MaterialMap
	MaterialMap struct {
		Texture Texture2D // Material map texture
		Color   Color     // Material map color
		Value   float32   // Material map value
	}

	// Material, includes shader and maps
	Material struct {
		Shader Shader       // Material shader
		Maps   *MaterialMap // Material maps array (MAX_MATERIAL_MAPS)
		Params [4]float32   // Material generic parameters (if required)
	}

	// Transform, vertex transformation data
	Transform struct {
		Translation Vector3    // Translation
		Rotation    Quaternion // Rotation
		Scale       Vector3    // Scale
	}

	// Bone, skeletal animation bone
	BoneInfo struct {
		Name   [32]byte // Bone name
		Parent int32    // Bone parent
	}

	// Model, meshes, materials and animation data
	Model struct {
		Transform Matrix // Local transform matrix

		MeshCount     int32     // Number of meshes
		MaterialCount int32     // Number of materials
		Meshes        *Mesh     // Meshes array
		Materials     *Material // Materials array
		MeshMaterial  *int32    // Mesh material number

		// Animation data
		BoneCount int32      // Number of bones
		Bones     *BoneInfo  // Bones information (skeleton)
		BindPose  *Transform // Bones base transformation (pose)
	}

	// ModelAnimation
	ModelAnimation struct {
		BoneCount  int32       // Number of bones
		FrameCount int32       // Number of animation frames
		Bones      *BoneInfo   // Bones information (skeleton)
		FramePoses **Transform // Poses array by frame
		Name       [32]byte    // Animation name
	}

	// Ray, ray for raycasting
	Ray struct {
		Position  Vector3 // Ray position (origin)
		Direction Vector3 // Ray direction (normalized)
	}

	// RayCollision, ray hit information
	RayCollision struct {
		Hit      bool    // Did the ray hit something?
		Distance float32 // Distance to the nearest hit
		Point    Vector3 // Point of the nearest hit
		Normal   Vector3 // Surface normal of hit
	}

	// VrDeviceInfo, Head-Mounted-Display device parameters
	VrDeviceInfo struct {
		HResolution            int32      // Horizontal resolution in pixels
		VResolution            int32      // Vertical resolution in pixels
		HScreenSize            float32    // Horizontal size in meters
		VScreenSize            float32    // Vertical size in meters
		EyeToScreenDistance    float32    // Distance between eye and display in meters
		LensSeparationDistance float32    // Lens separation distance in meters
		InterpupillaryDistance float32    // IPD (distance between pupils) in meters
		LensDistortionValues   [4]float32 // Lens distortion constant parameters
		chromaAbCorrection     [4]float32 // Chromatic aberration correction parameters
	}

	// VrStereoConfig, VR stereo rendering configuration for simulator
	VrStereoConfig struct {
		Projection        [2]Matrix  // VR projection matrices (per eye)
		ViewOffset        [2]Matrix  // VR view offset matrices (per eye)
		LeftLensCenter    [2]float32 // VR left lens center
		RightLensCenter   [2]float32 // VR right lens center
		LeftScreenCenter  [2]float32 // VR left screen center
		RightScreenCenter [2]float32 // VR right screen center
		Scale             [2]float32 // VR distortion scale
		ScaleIn           [2]float32 // VR distortion scale in
	}
)
