package raylib

import (
	"unsafe"
)

const (
	VersionMajor = 5
	VersionMinor = 5
	VersionPatch = 0
	Version      = "5.5"
)

// Some Basic Colors
// NOTE: Custom raylib color palette for amazing visuals on WHITE background
var (
	LightGray  = Color{200, 200, 200, 255} // Light Gray
	Gray       = Color{130, 130, 130, 255} // Gray
	DarkGray   = Color{80, 80, 80, 255}    // Dark Gray
	Yellow     = Color{253, 249, 0, 255}   // Yellow
	Gold       = Color{255, 203, 0, 255}   // Gold
	Orange     = Color{255, 161, 0, 255}   // Orange
	Pink       = Color{255, 109, 194, 255} // Pink
	Red        = Color{230, 41, 55, 255}   // Red
	Maroon     = Color{190, 33, 55, 255}   // Maroon
	Green      = Color{0, 228, 48, 255}    // Green
	Lime       = Color{0, 158, 47, 255}    // Lime
	DarkGreen  = Color{0, 117, 44, 255}    // Dark Green
	SkyBlue    = Color{102, 191, 255, 255} // Sky Blue
	Blue       = Color{0, 121, 241, 255}   // Blue
	DarkBlue   = Color{0, 82, 172, 255}    // Dark Blue
	Purple     = Color{200, 122, 255, 255} // Purple
	Violet     = Color{135, 60, 190, 255}  // Violet
	DarkPurple = Color{112, 31, 126, 255}  // Dark Purple
	Beige      = Color{211, 176, 131, 255} // Beige
	Brown      = Color{127, 106, 79, 255}  // Brown
	DarkBrown  = Color{76, 63, 47, 255}    // Dark Brown
	White      = Color{255, 255, 255, 255} // White
	Black      = Color{0, 0, 0, 255}       // Black
	Blank      = Color{0, 0, 0, 0}         // Blank (Transparent)
	Magenta    = Color{255, 0, 255, 255}   // Magenta
	RayWhite   = Color{245, 245, 245, 255} // My own White (raylib logo)
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

	// BoundingBox
	BoundingBox struct {
		Min Vector3 // Minimum vertex box-corner
		Max Vector3 // Maximum vertex box-corner
	}

	// Wave, audio wave data
	Wave struct {
		FrameCount uint32         // Total number of frames (considering channels)
		SampleRate uint32         // Frequency (samples per second)
		SampleSize uint32         // Bit depth (bits per sample): 8, 16, 32 (24 not supported)
		Channels   uint32         // Number of channels (1-mono, 2-stereo, ...)
		Data       unsafe.Pointer // Buffer data pointer
	}

	// AudioStream, custom audio stream
	AudioStream struct {
		Buffer     unsafe.Pointer // Pointer to internal data used by the audio system
		Processor  unsafe.Pointer // Pointer to internal data processor, useful for audio effects
		SampleRate uint32         // Frequency (samples per second)
		SampleSize uint32         // Bit depth (bits per sample): 8, 16, 32 (24 not supported)
		Channels   uint32         // Number of channels (1-mono, 2-stereo, ...)
	}

	// Sound
	Sound struct {
		Stream     AudioStream // Audio stream
		FrameCount uint32      // Total number of frames (considering channels)
	}

	// Music, audio stream, anything longer than ~10 seconds should be streamed
	Music struct {
		Stream     AudioStream    // Audio stream
		FrameCount uint32         // Total number of frames (considering channels)
		Looping    bool           // Music looping enable
		CtxType    int32          // Type of music context (audio filetype)
		CtxData    unsafe.Pointer // Audio context data, depends on type
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

	// File path list
	FilePathList struct {
		Capacity uint32 // Filepaths max entries
		Count    uint32 // Filepaths entries count
		Paths    **byte // Filepaths entries
	}

	// Automation event
	AutomationEvent struct {
		Frame  uint32   // Event frame
		Type   uint32   // Event type (AutomationEventType)
		Params [4]int32 // Event parameters (if required)
	}

	// Automation event list
	AutomationEventList struct {
		Capacity uint32           // Events max entries (MAX_AUTOMATION_EVENTS)
		Count    uint32           // Events entries count
		Events   *AutomationEvent // Events entries
	}
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
	FlagWindowMousePassthrough = ConfigFlags(0x00004000) // Set to support mouse passthrough, only supported when FlagWindowUndecorated
	FlagBorderlessWindowedMode = ConfigFlags(0x00008000) // Set to run program in borderless windowed mode
	FlagMsaa4xHint             = ConfigFlags(0x00000020) // Set to try enabling MSAA 4X
	FlagInterlacedHint         = ConfigFlags(0x00010000) // Set to try enabling interlaced video format (for V3D)
)

// Trace log level
type TraceLogLevel uint32

const (
	LogAll     TraceLogLevel = iota // Display all logs
	LogTrace                        // Trace logging, intended for internal use only
	LogDebug                        // Debug logging, used for internal debugging, it should be disabled on release builds
	LogInfo                         // Info logging, used for program execution info
	LogWarning                      // Warning logging, used on recoverable failures
	LogError                        // Error logging, used on unrecoverable failures
	LogFatal                        // Fatal logging, used to abort program: exit(EXIT_FAILURE)
	LogNone                         // Disable logging
)

// Keyboard keys (US keyboard layout)
type KeyboardKey uint32

const (
	KeyNull = KeyboardKey(0) // Key: NULL, used for no key pressed

	// Alphanumeric keys
	KeyApostrophe   = KeyboardKey(39) // Key: '
	KeyComma        = KeyboardKey(44) // Key: ,
	KeyMinus        = KeyboardKey(45) // Key: -
	KeyPeriod       = KeyboardKey(46) // Key: .
	KeySlash        = KeyboardKey(47) // Key: /
	KeyZero         = KeyboardKey(48) // Key: 0
	KeyOne          = KeyboardKey(49) // Key: 1
	KeyTwo          = KeyboardKey(50) // Key: 2
	KeyThree        = KeyboardKey(51) // Key: 3
	KeyFour         = KeyboardKey(52) // Key: 4
	KeyFive         = KeyboardKey(53) // Key: 5
	KeySix          = KeyboardKey(54) // Key: 6
	KeySeven        = KeyboardKey(55) // Key: 7
	KeyEight        = KeyboardKey(56) // Key: 8
	KeyNine         = KeyboardKey(57) // Key: 9
	KeySemicolon    = KeyboardKey(59) // Key: ;
	KeyEqual        = KeyboardKey(61) // Key: =
	KeyA            = KeyboardKey(65) // Key: A | a
	KeyB            = KeyboardKey(66) // Key: B | b
	KeyC            = KeyboardKey(67) // Key: C | c
	KeyD            = KeyboardKey(68) // Key: D | d
	KeyE            = KeyboardKey(69) // Key: E | e
	KeyF            = KeyboardKey(70) // Key: F | f
	KeyG            = KeyboardKey(71) // Key: G | g
	KeyH            = KeyboardKey(72) // Key: H | h
	KeyI            = KeyboardKey(73) // Key: I | i
	KeyJ            = KeyboardKey(74) // Key: J | j
	KeyK            = KeyboardKey(75) // Key: K | k
	KeyL            = KeyboardKey(76) // Key: L | l
	KeyM            = KeyboardKey(77) // Key: M | m
	KeyN            = KeyboardKey(78) // Key: N | n
	KeyO            = KeyboardKey(79) // Key: O | o
	KeyP            = KeyboardKey(80) // Key: P | p
	KeyQ            = KeyboardKey(81) // Key: Q | q
	KeyR            = KeyboardKey(82) // Key: R | r
	KeyS            = KeyboardKey(83) // Key: S | s
	KeyT            = KeyboardKey(84) // Key: T | t
	KeyU            = KeyboardKey(85) // Key: U | u
	KeyV            = KeyboardKey(86) // Key: V | v
	KeyW            = KeyboardKey(87) // Key: W | w
	KeyX            = KeyboardKey(88) // Key: X | x
	KeyY            = KeyboardKey(89) // Key: Y | y
	KeyZ            = KeyboardKey(90) // Key: Z | z
	KeyLeftBracket  = KeyboardKey(91) // Key: [
	KeyBackslash    = KeyboardKey(92) // Key: '\'
	KeyRightBracket = KeyboardKey(93) // Key: ]
	KeyGrave        = KeyboardKey(96) // Key: `

	// Function keys
	KeySpace        = KeyboardKey(32)  // Key: Space
	KeyEscape       = KeyboardKey(256) // Key: Esc
	KeyEnter        = KeyboardKey(257) // Key: Enter
	KeyTab          = KeyboardKey(258) // Key: Tab
	KeyBackspace    = KeyboardKey(259) // Key: Backspace
	KeyInsert       = KeyboardKey(260) // Key: Ins
	KeyDelete       = KeyboardKey(261) // Key: Del
	KeyRight        = KeyboardKey(262) // Key: Cursor right
	KeyLeft         = KeyboardKey(263) // Key: Cursor left
	KeyDown         = KeyboardKey(264) // Key: Cursor down
	KeyUp           = KeyboardKey(265) // Key: Cursor up
	KeyPageUp       = KeyboardKey(266) // Key: Page up
	KeyPageDown     = KeyboardKey(267) // Key: Page down
	KeyHome         = KeyboardKey(268) // Key: Home
	KeyEnd          = KeyboardKey(269) // Key: End
	KeyCapsLock     = KeyboardKey(280) // Key: Caps lock
	KeyScrollLock   = KeyboardKey(281) // Key: Scroll down
	KeyNumLock      = KeyboardKey(282) // Key: Num lock
	KeyPrintScreen  = KeyboardKey(283) // Key: Print screen
	KeyPause        = KeyboardKey(284) // Key: Pause
	KeyF1           = KeyboardKey(290) // Key: F1
	KeyF2           = KeyboardKey(291) // Key: F2
	KeyF3           = KeyboardKey(292) // Key: F3
	KeyF4           = KeyboardKey(293) // Key: F4
	KeyF5           = KeyboardKey(294) // Key: F5
	KeyF6           = KeyboardKey(295) // Key: F6
	KeyF7           = KeyboardKey(296) // Key: F7
	KeyF8           = KeyboardKey(297) // Key: F8
	KeyF9           = KeyboardKey(298) // Key: F9
	KeyF10          = KeyboardKey(299) // Key: F10
	KeyF11          = KeyboardKey(300) // Key: F11
	KeyF12          = KeyboardKey(301) // Key: F12
	KeyLeftShift    = KeyboardKey(340) // Key: Shift left
	KeyLeftControl  = KeyboardKey(341) // Key: Control left
	KeyLeftAlt      = KeyboardKey(342) // Key: Alt left
	KeyLeftSuper    = KeyboardKey(343) // Key: Super left
	KeyRightShift   = KeyboardKey(344) // Key: Shift right
	KeyRightControl = KeyboardKey(345) // Key: Control right
	KeyRightAlt     = KeyboardKey(346) // Key: Alt right
	KeyRightSuper   = KeyboardKey(347) // Key: Super right
	KeyKbMenu       = KeyboardKey(348) // Key: KB menu

	// Keypad keys
	KeyKp0        = KeyboardKey(320) // Key: Keypad 0
	KeyKp1        = KeyboardKey(321) // Key: Keypad 1
	KeyKp2        = KeyboardKey(322) // Key: Keypad 2
	KeyKp3        = KeyboardKey(323) // Key: Keypad 3
	KeyKp4        = KeyboardKey(324) // Key: Keypad 4
	KeyKp5        = KeyboardKey(325) // Key: Keypad 5
	KeyKp6        = KeyboardKey(326) // Key: Keypad 6
	KeyKp7        = KeyboardKey(327) // Key: Keypad 7
	KeyKp8        = KeyboardKey(328) // Key: Keypad 8
	KeyKp9        = KeyboardKey(329) // Key: Keypad 9
	KeyKpDecimal  = KeyboardKey(330) // Key: Keypad .
	KeyKpDivide   = KeyboardKey(331) // Key: Keypad /
	KeyKpMultiply = KeyboardKey(332) // Key: Keypad *
	KeyKpSubtract = KeyboardKey(333) // Key: Keypad -
	KeyKpAdd      = KeyboardKey(334) // Key: Keypad +
	KeyKpEnter    = KeyboardKey(335) // Key: Keypad Enter
	KeyKpEqual    = KeyboardKey(336) // Key: Keypad =

	// Android key buttons
	KeyBack       = KeyboardKey(4)  // Key: Android back button
	KeyMenu       = KeyboardKey(5)  // Key: Android menu button
	KeyVolumeUp   = KeyboardKey(24) // Key: Android volume up button
	KeyVolumeDown = KeyboardKey(25) // Key: Android volume down button
)

// Mouse buttons
type MouseButton uint32

const (
	MouseButtonLeft    MouseButton = iota // Mouse button left
	MouseButtonRight                      // Mouse button right
	MouseButtonMiddle                     // Mouse button middle (pressed wheel)
	MouseButtonSide                       // Mouse button side (advanced mouse device)
	MouseButtonExtra                      // Mouse button extra (advanced mouse device)
	MouseButtonForward                    // Mouse button forward (advanced mouse device)
	MouseButtonBack                       // Mouse button back (advanced mouse device)
)

// Mouse cursor
type MouseCursor uint32

const (
	MouseCursorDefault      MouseCursor = iota // Default pointer shape
	MouseCursorArrow                           // Arrow shape
	MouseCursorIbeam                           // Text writing cursor shape
	MouseCursorCrosshair                       // Cross shape
	MouseCursorPointingHand                    // Pointing hand cursor
	MouseCursorResizeEw                        // Horizontal resize/move arrow shape
	MouseCursorResizeNs                        // Vertical resize/move arrow shape
	MouseCursorResizeNwse                      // Top-left to bottom-right diagonal resize/move arrow shape
	MouseCursorResizeNesw                      // The top-right to bottom-left diagonal resize/move arrow shape
	MouseCursorResizeAll                       // The omnidirectional resize/move cursor shape
	MouseCursorNotAllowed                      // The operation-not-allowed shape
)

// Gamepad buttons
type GamepadButton uint32

const (
	GamepadButtonUnknown        GamepadButton = iota // Unknown button, just for error checking
	GamepadButtonLeftFaceUp                          // Gamepad left DPAD up button
	GamepadButtonLeftFaceRight                       // Gamepad left DPAD right button
	GamepadButtonLeftFaceDown                        // Gamepad left DPAD down button
	GamepadButtonLeftFaceLeft                        // Gamepad left DPAD left button
	GamepadButtonRightFaceUp                         // Gamepad right button up (i.e. PS3: Triangle, Xbox: Y)
	GamepadButtonRightFaceRight                      // Gamepad right button right (i.e. PS3: Circle, Xbox: B)
	GamepadButtonRightFaceDown                       // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	GamepadButtonRightFaceLeft                       // Gamepad right button left (i.e. PS3: Square, Xbox: X)
	GamepadButtonLeftTrigger1                        // Gamepad top/back trigger left (first), it could be a trailing button
	GamepadButtonLeftTrigger2                        // Gamepad top/back trigger left (second), it could be a trailing button
	GamepadButtonRightTrigger1                       // Gamepad top/back trigger right (first), it could be a trailing button
	GamepadButtonRightTrigger2                       // Gamepad top/back trigger right (second), it could be a trailing button
	GamepadButtonMiddleLeft                          // Gamepad center buttons, left one (i.e. PS3: Select)
	GamepadButtonMiddle                              // Gamepad center buttons, middle one (i.e. PS3: PS, Xbox: XBOX)
	GamepadButtonMiddleRight                         // Gamepad center buttons, right one (i.e. PS3: Start)
	GamepadButtonLeftThumb                           // Gamepad joystick pressed button left
	GamepadButtonRightThumb                          // Gamepad joystick pressed button right
)

// Gamepad axis
type GamepadAxis uint32

const (
	GamepadAxisLeftX        GamepadAxis = iota // Gamepad left stick X axis
	GamepadAxisLeftY                           // Gamepad left stick Y axis
	GamepadAxisRightX                          // Gamepad right stick X axis
	GamepadAxisRightY                          // Gamepad right stick Y axis
	GamepadAxisLeftTrigger                     // Gamepad back trigger left, pressure level: [1..-1]
	GamepadAxisRightTrigger                    // Gamepad back trigger right, pressure level: [1..-1]
)

// Material map index
type MaterialMapIndex uint32

const (
	MaterialMapAlbedo     MaterialMapIndex = iota // Albedo material (same as: MATERIAL_MAP_DIFFUSE)
	MaterialMapMetalness                          // Metalness material (same as: MATERIAL_MAP_SPECULAR)
	MaterialMapNormal                             // Normal material
	MaterialMapRoughness                          // Roughness material
	MaterialMapOcclusion                          // Ambient occlusion material
	MaterialMapEmission                           // Emission material
	MaterialMapHeight                             // Heightmap material
	MaterialMapCubemap                            // Cubemap material (NOTE: Uses GL_TEXTURE_CUBE_MAP)
	MaterialMapIrradiance                         // Irradiance material (NOTE: Uses GL_TEXTURE_CUBE_MAP)
	MaterialMapPrefilter                          // Prefilter material (NOTE: Uses GL_TEXTURE_CUBE_MAP)
	MaterialMapBrdf                               // Brdf material

	MaterialMapDiffuse  = MaterialMapAlbedo
	MaterialMapSpecular = MaterialMapMetalness
)

// Shader location index
type ShaderLocationIndex uint32

const (
	ShaderLocVertexPosition    ShaderLocationIndex = iota // Shader location: vertex attribute: position
	ShaderLocVertexTexCoord01                             // Shader location: vertex attribute: texcoord01
	ShaderLocVertexTexCoord02                             // Shader location: vertex attribute: texcoord02
	ShaderLocVertexNormal                                 // Shader location: vertex attribute: normal
	ShaderLocVertexTangent                                // Shader location: vertex attribute: tangent
	ShaderLocVertexColor                                  // Shader location: vertex attribute: color
	ShaderLocMatrixMvp                                    // Shader location: matrix uniform: model-view-projection
	ShaderLocMatrixView                                   // Shader location: matrix uniform: view (camera transform)
	ShaderLocMatrixProjection                             // Shader location: matrix uniform: projection
	ShaderLocMatrixModel                                  // Shader location: matrix uniform: model (transform)
	ShaderLocMatrixNormal                                 // Shader location: matrix uniform: normal
	ShaderLocVectorView                                   // Shader location: vector uniform: view
	ShaderLocColorDiffuse                                 // Shader location: vector uniform: diffuse color
	ShaderLocColorSpecular                                // Shader location: vector uniform: specular color
	ShaderLocColorAmbient                                 // Shader location: vector uniform: ambient color
	ShaderLocMapAlbedo                                    // Shader location: sampler2d texture: albedo (same as: SHADER_LOC_MAP_DIFFUSE)
	ShaderLocMapMetalness                                 // Shader location: sampler2d texture: metalness (same as: SHADER_LOC_MAP_SPECULAR)
	ShaderLocMapNormal                                    // Shader location: sampler2d texture: normal
	ShaderLocMapRoughness                                 // Shader location: sampler2d texture: roughness
	ShaderLocMapOcclusion                                 // Shader location: sampler2d texture: occlusion
	ShaderLocMapEmission                                  // Shader location: sampler2d texture: emission
	ShaderLocMapHeight                                    // Shader location: sampler2d texture: height
	ShaderLocMapCubemap                                   // Shader location: samplerCube texture: cubemap
	ShaderLocMapIrradiance                                // Shader location: samplerCube texture: irradiance
	ShaderLocMapPrefilter                                 // Shader location: samplerCube texture: prefilter
	ShaderLocMapBrdf                                      // Shader location: sampler2d texture: brdf
	ShaderLocVertexBoneids                                // Shader location: vertex attribute: boneIds
	ShaderLocVertexBoneweights                            // Shader location: vertex attribute: boneWeights
	ShaderLocBoneMatrices                                 // Shader location: array of matrices uniform: boneMatrices

	ShaderLocMapDiffuse  = ShaderLocMapAlbedo
	ShaderLocMapSpecular = ShaderLocMapMetalness
)

// Shader uniform data type
type ShaderUniformDataType uint32

const (
	ShaderUniformFloat     ShaderUniformDataType = iota // Shader uniform type: float
	ShaderUniformVec2                                   // Shader uniform type: vec2 (2 float)
	ShaderUniformVec3                                   // Shader uniform type: vec3 (3 float)
	ShaderUniformVec4                                   // Shader uniform type: vec4 (4 float)
	ShaderUniformInt                                    // Shader uniform type: int
	ShaderUniformIvec2                                  // Shader uniform type: ivec2 (2 int)
	ShaderUniformIvec3                                  // Shader uniform type: ivec3 (3 int)
	ShaderUniformIvec4                                  // Shader uniform type: ivec4 (4 int)
	ShaderUniformSampler2d                              // Shader uniform type: sampler2d
)

// Shader attribute data types
type ShaderAttributeDataType uint32

const (
	ShaderAttribFloat ShaderAttributeDataType = iota // Shader attribute type: float
	ShaderAttribVec2                                 // Shader attribute type: vec2 (2 float)
	ShaderAttribVec3                                 // Shader attribute type: vec3 (3 float)
	ShaderAttribVec4                                 // Shader attribute type: vec4 (4 float)
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

// Texture parameters: filter mode
// NOTE 1: Filtering considers mipmaps if available in the texture
// NOTE 2: Filter is accordingly set for minification and magnification
type TextureFilter uint32

const (
	TextureFilterPoint          TextureFilter = iota // No filter, just pixel approximation
	TextureFilterBilinear                            // Linear filtering
	TextureFilterTrilinear                           // Trilinear filtering (linear with mipmaps)
	TextureFilterAnisotropic4x                       // Anisotropic filtering 4x
	TextureFilterAnisotropic8x                       // Anisotropic filtering 8x
	TextureFilterAnisotropic16x                      // Anisotropic filtering 16x
)

// Texture parameters: wrap mode
type TextureWrap uint32

const (
	TextureWrapRepeat       TextureWrap = iota // Repeats texture in tiled mode
	TextureWrapClamp                           // Clamps texture to edge pixel in tiled mode
	TextureWrapMirrorRepeat                    // Mirrors and repeats the texture in tiled mode
	TextureWrapMirrorClamp                     // Mirrors and clamps to border the texture in tiled mode
)

// Cubemap layouts
type CubemapLayout uint32

const (
	CubemapLayoutAutoDetect       CubemapLayout = iota // Automatically detect layout type
	CubemapLayoutLineVertical                          // Layout is defined by a vertical line with faces
	CubemapLayoutLineHorizontal                        // Layout is defined by a horizontal line with faces
	CubemapLayoutCrossThreeByFour                      // Layout is defined by a 3x4 cross with cubemap faces
	CubemapLayoutCrossFourByThree                      // Layout is defined by a 4x3 cross with cubemap faces
)

// Font type, defines generation method
type FontType uint32

const (
	FontDefault FontType = iota // Default font generation, anti-aliased
	FontBitmap                  // Bitmap font generation, no anti-aliasing
	FontSdf                     // SDF font generation, requires external shader
)

// Color blending modes (pre-defined)
type BlendMode uint32

const (
	BlendAlpha            BlendMode = iota // Blend textures considering alpha (default)
	BlendAdditive                          // Blend textures adding colors
	BlendMultiplied                        // Blend textures multiplying colors
	BlendAddColors                         // Blend textures adding colors (alternative)
	BlendSubtractColors                    // Blend textures subtracting colors (alternative)
	BlendAlphaPremultiply                  // Blend premultiplied textures considering alpha
	BlendCustom                            // Blend textures using custom src/dst factors (use rlSetBlendFactors())
	BlendCustomSeparate                    // Blend textures using custom rgb/alpha separate src/dst factors (use rlSetBlendFactorsSeparate())
)

// Gesture
// NOTE: Provided as bit-wise flags to enable only desired gestures
type Gesture uint32

const (
	GestureNone       = Gesture(0)   // No gesture
	GestureTap        = Gesture(1)   // Tap gesture
	GestureDoubletap  = Gesture(2)   // Double tap gesture
	GestureHold       = Gesture(4)   // Hold gesture
	GestureDrag       = Gesture(8)   // Drag gesture
	GestureSwipeRight = Gesture(16)  // Swipe right gesture
	GestureSwipeLeft  = Gesture(32)  // Swipe left gesture
	GestureSwipeUp    = Gesture(64)  // Swipe up gesture
	GestureSwipeDown  = Gesture(128) // Swipe down gesture
	GesturePinchIn    = Gesture(256) // Pinch in gesture
	GesturePinchOut   = Gesture(512) // Pinch out gesture
)

// Camera system modes
type CameraMode uint32

const (
	CameraCustom      CameraMode = iota // Camera custom, controlled by user (UpdateCamera() does nothing)
	CameraFree                          // Camera free mode
	CameraOrbital                       // Camera orbital, around target, zoom supported
	CameraFirstPerson                   // Camera first person
	CameraThirdPerson                   // Camera third person
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

// Callbacks

type (
// typedef void (*TraceLogCallback)(int logLevel, const char *text, va_list args)
// typedef unsigned char *(*LoadFileDataCallback)(const char *fileName, int *dataSize);
// typedef bool (*SaveFileDataCallback)(const char *fileName, void *data, int dataSize);
// typedef char *(*LoadFileTextCallback)(const char *fileName);
// typedef bool (*SaveFileTextCallback)(const char *fileName, char *text);
// typedef void (*AudioCallback)(void *bufferData, unsigned int frames);
)
