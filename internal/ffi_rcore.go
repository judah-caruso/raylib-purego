package internal

var (
	// Window-related functions
	InitWindow               = bind(&void, "InitWindow", &cint, &cint, &cstr)   // width, height, title :: Initialize window and OpenGL context
	CloseWindow              = bind(&void, "CloseWindow")                       // Close window and unload OpenGL context
	WindowShouldClose        = bind(&cbool, "WindowShouldClose")                // Check if application should close (KEY_ESCAPE pressed or windows close icon clicked)
	IsWindowReady            = bind(&cbool, "IsWindowReady")                    // Check if window has been initialized successfully
	IsWindowFullscreen       = bind(&cbool, "IsWindowFullscreen")               // Check if window is currently fullscreen
	IsWindowHidden           = bind(&cbool, "IsWindowHidden")                   // Check if window is currently hidden
	IsWindowMinimized        = bind(&cbool, "IsWindowMinimized")                // Check if window is currently minimized
	IsWindowMaximized        = bind(&cbool, "IsWindowMaximized")                // Check if window is currently maximized
	IsWindowFocused          = bind(&cbool, "IsWindowFocused")                  // Check if window is currently focused
	IsWindowResized          = bind(&cbool, "IsWindowResized")                  // Check if window has been resized last frame
	IsWindowState            = bind(&cbool, "IsWindowState", &tConfigFlags)     // flag :: Check if one specific window flag is enabled
	SetWindowState           = bind(&void, "SetWindowState", &tConfigFlags)     // flags :: Set window configuration state using flags
	ClearWindowState         = bind(&void, "ClearWindowState", &tConfigFlags)   // flags :: Clear window configuration state flags
	ToggleFullscreen         = bind(&void, "ToggleFullscreen")                  // Toggle window state: fullscreen/windowed, resizes monitor to match window resolution
	ToggleBorderlessWindowed = bind(&void, "ToggleBorderlessWindowed")          // Toggle window state: borderless windowed, resizes window to match monitor resolution
	MaximizeWindow           = bind(&void, "MaximizeWindow")                    // Set window state: maximized, if resizable
	MinimizeWindow           = bind(&void, "MinimizeWindow")                    // Set window state: minimized, if resizable
	RestoreWindow            = bind(&void, "RestoreWindow")                     // Set window state: not minimized/maximized
	SetWindowIcon            = bind(&void, "SetWindowIcon", &tImage)            // image :: Set icon for window (single image, RGBA 32bit)
	SetWindowIcons           = bind(&void, "SetWindowIcons", &tImagePtr, &cint) // images, count :: Set icon for window (multiple images, RGBA 32bit)
	SetWindowTitle           = bind(&void, "SetWindowTitle", &cstr)             // title :: Set title for window
	SetWindowPosition        = bind(&void, "SetWindowPosition", &cint, &cint)   // x, y :: Set window position on screen
	SetWindowMonitor         = bind(&void, "SetWindowMonitor", &cint)           // monitor :: Set monitor for the current window
	SetWindowMinSize         = bind(&void, "SetWindowMinSize", &cint, &cint)    // width, height :: Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
	SetWindowMaxSize         = bind(&void, "SetWindowMaxSize", &cint, &cint)    // width, height :: Set window maximum dimensions (for FLAG_WINDOW_RESIZABLE)
	SetWindowSize            = bind(&void, "SetWindowSize", &cint, &cint)       // width, height :: Set window dimensions
	SetWindowOpacity         = bind(&void, "SetWindowOpacity", &float)          // opacity :: Set window opacity [0.0f..1.0f]
	SetWindowFocused         = bind(&void, "SetWindowFocused")                  // Set window focused
	GetWindowHandle          = bind(&ptr, "GetWindowHandle")                    // Get native window handle
	GetScreenWidth           = bind(&cint, "GetScreenWidth")                    // Get current screen width
	GetScreenHeight          = bind(&cint, "GetScreenHeight")                   // Get current screen height
	GetRenderWidth           = bind(&cint, "GetRenderWidth")                    // Get current render width (it considers HiDPI)
	GetRenderHeight          = bind(&cint, "GetRenderHeight")                   // Get current render height (it considers HiDPI)
	GetMonitorCount          = bind(&cint, "GetMonitorCount")                   // Get number of connected monitors
	GetCurrentMonitor        = bind(&cint, "GetCurrentMonitor")                 // Get current monitor where window is placed
	GetMonitorPosition       = bind(&tVector2, "GetMonitorPosition", &cint)     // monitor :: Get specified monitor position
	GetMonitorWidth          = bind(&cint, "GetMonitorWidth", &cint)            // monitor :: Get specified monitor width (current video mode used by monitor)
	GetMonitorHeight         = bind(&cint, "GetMonitorHeight", &cint)           // monitor :: Get specified monitor height (current video mode used by monitor)
	GetMonitorPhysicalWidth  = bind(&cint, "GetMonitorPhysicalWidth", &cint)    // monitor :: Get specified monitor physical width in millimetres
	GetMonitorPhysicalHeight = bind(&cint, "GetMonitorPhysicalHeight", &cint)   // monitor :: Get specified monitor physical height in millimetres
	GetMonitorRefreshRate    = bind(&cint, "GetMonitorRefreshRate", &cint)      // monitor :: Get specified monitor refresh rate
	GetWindowPosition        = bind(&tVector2, "GetWindowPosition")             // Get window position XY on monitor
	GetWindowScaleDPI        = bind(&tVector2, "GetWindowScaleDPI")             // Get window scale DPI factor
	GetMonitorName           = bind(&cstr, "GetMonitorName", &cint)             // monitor :: Get the human-readable, UTF-8 encoded name of the specified monitor
	SetClipboardText         = bind(&void, "SetClipboardText", &cstr)           // text :: Set clipboard text content
	GetClipboardText         = bind(&cstr, "GetClipboardText")                  // Get clipboard text content
	GetClipboardImage        = bind(&tImage, "GetClipboardImage")               // Get clipboard image
	EnableEventWaiting       = bind(&void, "EnableEventWaiting")                // Enable waiting for events on EndDrawing(), no automatic event polling
	DisableEventWaiting      = bind(&void, "DisableEventWaiting")               // Disable waiting for events on EndDrawing(), automatic events polling

	// Cursor-related functions
	ShowCursor       = bind(&void, "ShowCursor")        // Shows cursor
	HideCursor       = bind(&void, "HideCursor")        // Hides cursor
	IsCursorHidden   = bind(&cbool, "IsCursorHidden")   // Check if cursor is not visible
	EnableCursor     = bind(&void, "EnableCursor")      // Enables cursor (unlock cursor)
	DisableCursor    = bind(&void, "DisableCursor")     // Disables cursor (lock cursor)
	IsCursorOnScreen = bind(&cbool, "IsCursorOnScreen") // Check if cursor is on the screen

	// Drawing-related functions
	ClearBackground  = bind(&void, "ClearBackground", &tColor)                     // color :: Set background color (framebuffer clear color)
	BeginDrawing     = bind(&void, "BeginDrawing")                                 // Setup canvas (framebuffer) to start drawing
	EndDrawing       = bind(&void, "EndDrawing")                                   // End canvas drawing and swap buffers (double buffering)
	BeginMode2D      = bind(&void, "BeginMode2D", &tCamera2D)                      // camera :: Begin 2D mode with custom camera (2D)
	EndMode2D        = bind(&void, "EndMode2D")                                    // Ends 2D mode with custom camera
	BeginMode3D      = bind(&void, "BeginMode3D", &tCamera3D)                      // camera :: Begin 3D mode with custom camera (3D)
	EndMode3D        = bind(&void, "EndMode3D")                                    // Ends 3D mode and returns to default 2D orthographic mode
	BeginTextureMode = bind(&void, "BeginTextureMode", &tRenderTexture2D)          // target :: Begin drawing to render texture
	EndTextureMode   = bind(&void, "EndTextureMode")                               // Ends drawing to render texture
	BeginShaderMode  = bind(&void, "BeginShaderMode", &tShader)                    // shader :: Begin custom shader drawing
	EndShaderMode    = bind(&void, "EndShaderMode")                                // End custom shader drawing (use default shader)
	BeginBlendMode   = bind(&void, "BeginBlendMode", &cint)                        // mode :: Begin blending mode (alpha, additive, multiplied, subtract, custom)
	EndBlendMode     = bind(&void, "EndBlendMode")                                 // End blending mode (reset to default: alpha blending)
	BeginScissorMode = bind(&void, "BeginScissorMode", &cint, &cint, &cint, &cint) // x, y, width, height :: Begin scissor mode (define screen area for following drawing)
	EndScissorMode   = bind(&void, "EndScissorMode")                               // End scissor mode

	BeginVrStereoMode = bind(&void, "BeginVrStereoMode", &tVrStereoConfig) // config :: Begin stereo rendering (requires VR simulator)
	EndVrStereoMode   = bind(&void, "EndVrStereoMode")                     // End stereo rendering (requires VR simulator)

	// VR stereo config functions for VR simulator
	LoadVrStereoConfig   = bind(&tVrStereoConfig, "LoadVrStereoConfig", &tVrDeviceInfo) // device :: Load VR stereo config for VR simulator device parameters
	UnloadVrStereoConfig = bind(&void, "UnloadVrStereoConfig", &tVrStereoConfig)        // config :: Unload VR stereo config

	// Shader management functions
	// NOTE: Shader functionality is not available on OpenGL 1.1
	LoadShader              = bind(&tShader, "LoadShader", &cstr, &cstr)                          // vsFileName, fsFileName :: Load shader from files and bind default locations
	LoadShaderFromMemory    = bind(&tShader, "LoadShaderFromMemory", &cstr, &cstr)                // vsCode, fsCode :: Load shader from code strings and bind default locations
	IsShaderValid           = bind(&cbool, "IsShaderValid", &tShader)                             // shader :: Check if a shader is valid (loaded on GPU)
	GetShaderLocation       = bind(&cint, "GetShaderLocation", &tShader, &cstr)                   // shader, uniformName :: Get shader uniform location
	GetShaderLocationAttrib = bind(&cint, "GetShaderLocationAttrib", &tShader, &cstr)             // shader, attribName :: Get shader attribute location
	SetShaderValue          = bind(&void, "SetShaderValue", &tShader, &cint, &ptr, &cint)         // shader, locIndex, value, uniformType :: Set shader uniform value
	SetShaderValueV         = bind(&void, "SetShaderValueV", &tShader, &cint, &ptr, &cint, &cint) // shader, locIndex, value, uniformType, count :: Set shader uniform value vector
	SetShaderValueMatrix    = bind(&void, "SetShaderValueMatrix", &tShader, &cint, &tMatrix)      // shader, locIndex, mat :: Set shader uniform value (matrix 4x4)
	SetShaderValueTexture   = bind(&void, "SetShaderValueTexture", &tShader, &cint, &tTexture2D)  // shader, locIndex, texture :: Set shader uniform value for texture (sampler2d)
	UnloadShader            = bind(&void, "UnloadShader", &tShader)                               // shader :: Unload shader from GPU memory (VRAM)

	// Screen-space-related functions
	GetScreenToWorldRay   = bind(&tRay, "GetScreenToWorldRay", &tVector2, &tCamera)                  // position, camera :: Get a ray trace from screen position (i.e mouse)
	GetScreenToWorldRayEx = bind(&tRay, "GetScreenToWorldRayEx", &tVector2, &tCamera, &cint, &cint)  // position, camera, width, height :: Get a ray trace from screen position (i.e mouse) in a viewport
	GetWorldToScreen      = bind(&tVector2, "GetWorldToScreen", &tVector3, &tCamera)                 // position, camera :: Get the screen space position for a 3d world space position
	GetWorldToScreenEx    = bind(&tVector2, "GetWorldToScreenEx", &tVector3, &tCamera, &cint, &cint) // position, camera, width, height :: Get size position for a 3d world space position
	GetWorldToScreen2D    = bind(&tVector2, "GetWorldToScreen2D", &tVector2, &tCamera2D)             // position, camera :: Get the screen space position for a 2d camera world space position
	GetScreenToWorld2D    = bind(&tVector2, "GetScreenToWorld2D", &tVector2, &tCamera2D)             // position, camera :: Get the world space position for a 2d camera screen space position
	GetCameraMatrix       = bind(&tMatrix, "GetCameraMatrix", &tCamera)                              // camera :: Get camera transform matrix (view matrix)
	GetCameraMatrix2D     = bind(&tMatrix, "GetCameraMatrix2D", &tCamera2D)                          // camera :: Get camera 2d transform matrix

	// Timing-related functions
	SetTargetFPS = bind(&void, "SetTargetFPS", &cint) // fps :: Set target FPS (maximum)
	GetFrameTime = bind(&float, "GetFrameTime")       // Get time in seconds for last frame drawn (delta time)
	GetTime      = bind(&double, "GetTime")           // Get elapsed time in seconds since InitWindow()
	GetFPS       = bind(&cint, "GetFPS")              // Get current FPS

	// Custom frame control functions
	// NOTE: Those functions are intended for advanced users that want full control over the frame processing
	// By default EndDrawing() does this job: draws everything + SwapScreenBuffer() + manage frame timing + PollInputEvents()
	// To avoid that behaviour and control frame processes manually, enable in config.h: SUPPORT_CUSTOM_FRAME_CONTROL
	SwapScreenBuffer = bind(&void, "SwapScreenBuffer")  // Swap back buffer with front buffer (screen drawing)
	PollInputEvents  = bind(&void, "PollInputEvents")   // Register all input events
	WaitTime         = bind(&void, "WaitTime", &double) // seconds :: Wait for some time (halt program execution)

	// Random values generation functions
	SetRandomSeed        = bind(&void, "SetRandomSeed", &ucint)                    // seed :: Set the seed for the random number generator
	GetRandomValue       = bind(&cint, "GetRandomValue", &cint, &cint)             // min, max :: Get a random value between min and max (both included)
	LoadRandomSequence   = bind(&cint, "LoadRandomSequence", &ucint, &cint, &cint) // count, min, max :: Load random values sequence, no values repeated
	UnloadRandomSequence = bind(&void, "UnloadRandomSequence", &ptr)               // sequence :: Unload random values sequence

	// Misc. functions
	TakeScreenshot = bind(&void, "TakeScreenshot", &cstr)         // fileName :: Takes a screenshot of current screen (filename extension defines format)
	SetConfigFlags = bind(&void, "SetConfigFlags", &tConfigFlags) // flags :: Setup init configuration flags (view FLAGS)
	OpenURL        = bind(&void, "OpenURL", &cstr)                // url :: Open URL with default system browser (if available)

	// NOTE: Following functions implemented in module [utils]
	// ------------------------------------------------------------------
	// @todo(judah): support varargs?
	// void TraceLog(int logLevel, const char *text, ...);         // Show trace log messages (LOG_DEBUG, LOG_INFO, LOG_WARNING, LOG_ERROR...)

	SetTraceLogLevel = bind(&void, "SetTraceLogLevel", &cint) // logLevel :: Set the current threshold (minimum) log level
	MemAlloc         = bind(&void, "MemAlloc", &ucint)        // size :: Internal memory allocator
	MemRealloc       = bind(&void, "MemRealloc", &ptr, &cint) // ptr, size :: Internal memory reallocator
	MemFree          = bind(&void, "MemFree", &ptr)           // Internal memory free

	// @todo(judah): callbacks

	// Set custom callbacks
	// WARNING: Callbacks setup is intended for advanced users
	// void SetTraceLogCallback(TraceLogCallback callback);         // Set custom trace log
	// void SetLoadFileDataCallback(LoadFileDataCallback callback); // Set custom file binary data loader
	// void SetSaveFileDataCallback(SaveFileDataCallback callback); // Set custom file binary data saver
	// void SetLoadFileTextCallback(LoadFileTextCallback callback); // Set custom file text data loader
	// void SetSaveFileTextCallback(SaveFileTextCallback callback); // Set custom file text data saver

	// Files management functions
	LoadFileData     = bind(&ustr, "LoadFileData", &cstr, &intPtr)           // fileName, dataSize :: Load file data as byte array (read)
	UnloadFileData   = bind(&void, "UnloadFileData", &cstr)                  // data :: Unload file data allocated by LoadFileData()
	SaveFileData     = bind(&cbool, "SaveFileData", &cstr, &ptr, &cint)      // fileName, data, dataSize :: Save data to file from byte array (write), returns true on success
	ExportDataAsCode = bind(&cbool, "ExportDataAsCode", &ustr, &cint, &cstr) // data, dataSize, fileName :: Export data to code (.h), returns true on success
	LoadFileText     = bind(&cstr, "LoadFileText", &cstr)                    // fileName :: Load text data from file (read), returns a '\0' terminated string
	UnloadFileText   = bind(&void, "UnloadFileText", &cstr)                  // text :: Unload file text data allocated by LoadFileText()
	SaveFileText     = bind(&cbool, "SaveFileText", &cstr, &cstr)            // fileName, text :: Save text data to file (write), string must be '\0' terminated, returns true on success

	// File system functions
	FileExists              = bind(&cbool, "FileExists", &cstr)                                  // fileName :: Check if file exists
	DirectoryExists         = bind(&cbool, "DirectoryExists", &cstr)                             // dirPath :: Check if a directory path exists
	IsFileExtension         = bind(&cbool, "IsFileExtension", &cstr, &cstr)                      // fileName, ext :: Check file extension (including point: .png, .wav)
	GetFileLength           = bind(&cint, "GetFileLength", &cstr)                                // fileName :: Get file length in bytes (NOTE: GetFileSize() conflicts with windows.h)
	GetFileExtension        = bind(&cstr, "GetFileExtension", &cstr)                             // fileName :: Get pointer to extension for a filename string (includes dot: '.png')
	GetFileName             = bind(&cstr, "GetFileName", &cstr)                                  // filePath :: Get pointer to filename for a path string
	GetFileNameWithoutExt   = bind(&cstr, "GetFileNameWithoutExt", &cstr)                        // filePath :: Get filename string without extension (uses static string)
	GetDirectoryPath        = bind(&cstr, "GetDirectoryPath", &cstr)                             // filePath :: Get full path for a given fileName with path (uses static string)
	GetPrevDirectoryPath    = bind(&cstr, "GetPrevDirectoryPath", &cstr)                         // dirPath :: Get previous directory path for a given path (uses static string)
	GetWorkingDirectory     = bind(&cstr, "GetWorkingDirectory")                                 // Get current working directory (uses static string)
	GetApplicationDirectory = bind(&cstr, "GetApplicationDirectory")                             // Get the directory of the running application (uses static string)
	MakeDirectory           = bind(&cint, "MakeDirectory", &cstr)                                // dirPath :: Create directories (including full path requested), returns 0 on success
	ChangeDirectory         = bind(&cbool, "ChangeDirectory", &cstr)                             // dir :: Change working directory, return true on success
	IsPathFile              = bind(&cbool, "IsPathFile", &cstr)                                  // path :: Check if a given path is a file or a directory
	IsFileNameValid         = bind(&cbool, "IsFileNameValid", &cstr)                             // fileName :: Check if fileName is valid for the platform/OS
	LoadDirectoryFiles      = bind(&tFilePathList, "LoadDirectoryFiles", &cstr)                  // dirPath :: Load directory filepaths
	LoadDirectoryFilesEx    = bind(&tFilePathList, "LoadDirectoryFilesEx", &cstr, &cstr, &cbool) // basePath, filter, scanSubdirs :: Load directory filepaths with extension filtering and recursive directory scan. Use 'DIR' in the filter string to include directories in the result
	UnloadDirectoryFiles    = bind(&void, "UnloadDirectoryFiles", &tFilePathList)                // files :: Unload filepaths
	IsFileDropped           = bind(&cbool, "IsFileDropped")                                      // Check if a file has been dropped into window
	LoadDroppedFiles        = bind(&tFilePathList, "LoadDroppedFiles")                           // Load dropped filepaths
	UnloadDroppedFiles      = bind(&void, "UnloadDroppedFiles", &tFilePathList)                  // files :: Unload dropped filepaths
	GetFileModTime          = bind(&long, "GetFileModTime", &cstr)                               // fileName :: Get file modification time (last write time)

	// Compression/Encoding functionality
	CompressData     = bind(&ustr, "CompressData", &ustr, &cint, &intPtr)     // data, dataSize, compDataSize :: Compress data (DEFLATE algorithm), memory must be MemFree()
	DecompressData   = bind(&ustr, "DecompressData", &ustr, &cint, &intPtr)   // compData, compDataSize, dataSize :: Decompress data (DEFLATE algorithm), memory must be MemFree()
	EncodeDataBase64 = bind(&cstr, "EncodeDataBase64", &ustr, &cint, &intPtr) // data, dataSize, outputSize :: Encode data to Base64 string, memory must be MemFree()
	DecodeDataBase64 = bind(&ustr, "DecodeDataBase64", &ustr, &intPtr)        // data, outputSize :: Decode Base64 string data, memory must be MemFree()
	ComputeCRC32     = bind(&ucint, "ComputeCRC32", &ustr, &cint)             // data, dataSize :: Compute CRC32 hash code
	ComputeMD5       = bind(&uintPtr, "ComputeMD5", &ustr, &cint)             // data, dataSize :: Compute MD5 hash code, returns static int[4] (16 bytes)
	ComputeSHA1      = bind(&uintPtr, "ComputeSHA1", &ustr, &cint)            // data, dataSize :: Compute SHA1 hash code, returns static int[5] (20 bytes)

	// Automation events functionality
	LoadAutomationEventList       = bind(&tAutomationEventList, "LoadAutomationEventList", &cstr)           // fileName :: Load automation events list from file, NULL for empty list, capacity = MAX_AUTOMATION_EVENTS
	UnloadAutomationEventList     = bind(&void, "UnloadAutomationEventList", &tAutomationEventList)         // list :: Unload automation events list from file
	ExportAutomationEventList     = bind(&cbool, "ExportAutomationEventList", &tAutomationEventList, &cstr) // list, fileName :: Export automation events list as text file
	SetAutomationEventList        = bind(&void, "SetAutomationEventList", &tAutomationEventListPtr)         // list :: Set automation event list to record to
	SetAutomationEventBaseFrame   = bind(&void, "SetAutomationEventBaseFrame", &cint)                       // frame :: Set automation event internal base frame to start recording
	StartAutomationEventRecording = bind(&void, "StartAutomationEventRecording")                            // Start recording automation events (AutomationEventList must be set)
	StopAutomationEventRecording  = bind(&void, "StopAutomationEventRecording")                             // Stop recording automation events
	PlayAutomationEvent           = bind(&void, "PlayAutomationEvent", &tAutomationEvent)                   // event :: Play a recorded automation event

	// Input-related functions: keyboard
	IsKeyPressed       = bind(&cbool, "IsKeyPressed", &tKeyboardKey)       // key :: Check if a key has been pressed once
	IsKeyPressedRepeat = bind(&cbool, "IsKeyPressedRepeat", &tKeyboardKey) // key :: Check if a key has been pressed again
	IsKeyDown          = bind(&cbool, "IsKeyDown", &tKeyboardKey)          // key :: Check if a key is being pressed
	IsKeyReleased      = bind(&cbool, "IsKeyReleased", &tKeyboardKey)      // key :: Check if a key has been released once
	IsKeyUp            = bind(&cbool, "IsKeyUp", &tKeyboardKey)            // key :: Check if a key is NOT being pressed
	GetKeyPressed      = bind(&tKeyboardKey, "GetKeyPressed")              // Get key pressed (keycode), call it multiple times for keys queued, returns 0 when the queue is empty
	GetCharPressed     = bind(&tKeyboardKey, "GetCharPressed")             // Get char pressed (unicode), call it multiple times for chars queued, returns 0 when the queue is empty
	SetExitKey         = bind(&void, "SetExitKey", &tKeyboardKey)          // key :: Set a custom key to exit program (default is ESC)

	// Input-related functions: gamepads
	IsGamepadAvailable      = bind(&cbool, "IsGamepadAvailable", &cint)                         // gamepad :: Check if a gamepad is available
	GetGamepadName          = bind(&cstr, "GetGamepadName", &cint)                              // gamepad :: Get gamepad internal name id
	IsGamepadButtonPressed  = bind(&cbool, "IsGamepadButtonPressed", &cint, &tGamepadButton)    // gamepad, button :: Check if a gamepad button has been pressed once
	IsGamepadButtonDown     = bind(&cbool, "IsGamepadButtonDown", &cint, &tGamepadButton)       // gamepad, button :: Check if a gamepad button is being pressed
	IsGamepadButtonReleased = bind(&cbool, "IsGamepadButtonReleased", &cint, &tGamepadButton)   // gamepad, button :: Check if a gamepad button has been released once
	IsGamepadButtonUp       = bind(&cbool, "IsGamepadButtonUp", &cint, &tGamepadButton)         // gamepad, button :: Check if a gamepad button is NOT being pressed
	GetGamepadButtonPressed = bind(&cint, "GetGamepadButtonPressed")                            // Get the last gamepad button pressed
	GetGamepadAxisCount     = bind(&cint, "GetGamepadAxisCount", &cint)                         // gamepad :: Get gamepad axis count for a gamepad
	GetGamepadAxisMovement  = bind(&float, "GetGamepadAxisMovement", &cint, &tGamepadAxis)      // gamepad, axis :: Get axis movement value for a gamepad axis
	SetGamepadMappings      = bind(&cint, "SetGamepadMappings", &cstr)                          // mappings :: Set internal gamepad mappings (SDL_GameControllerDB)
	SetGamepadVibration     = bind(&void, "SetGamepadVibration", &cint, &float, &float, &float) // gamepad, leftMotor, rightMotor, duration :: Set gamepad vibration for both motors (duration in seconds)

	// Input-related functions: mouse
	IsMouseButtonPressed  = bind(&cbool, "IsMouseButtonPressed", &tMouseButton)  // button :: Check if a mouse button has been pressed once
	IsMouseButtonDown     = bind(&cbool, "IsMouseButtonDown", &tMouseButton)     // button :: Check if a mouse button is being pressed
	IsMouseButtonReleased = bind(&cbool, "IsMouseButtonReleased", &tMouseButton) // button :: Check if a mouse button has been released once
	IsMouseButtonUp       = bind(&cbool, "IsMouseButtonUp", &tMouseButton)       // button :: Check if a mouse button is NOT being pressed
	GetMouseX             = bind(&cint, "GetMouseX")                             // Get mouse position X
	GetMouseY             = bind(&cint, "GetMouseY")                             // Get mouse position Y
	GetMousePosition      = bind(&tVector2, "GetMousePosition")                  // Get mouse position XY
	GetMouseDelta         = bind(&tVector2, "GetMouseDelta")                     // Get mouse delta between frames
	SetMousePosition      = bind(&void, "SetMousePosition", &cint, &cint)        // x, y :: Set mouse position XY
	SetMouseOffset        = bind(&void, "SetMouseOffset", &cint, &cint)          // offsetX, offsetY :: Set mouse offset
	SetMouseScale         = bind(&void, "SetMouseScale", &float, &float)         // scaleX, scaleY :: Set mouse scaling
	GetMouseWheelMove     = bind(&float, "GetMouseWheelMove")                    // Get mouse wheel movement for X or Y, whichever is larger
	GetMouseWheelMoveV    = bind(&tVector2, "GetMouseWheelMoveV")                // Get mouse wheel movement for both X and Y
	SetMouseCursor        = bind(&void, "SetMouseCursor", &tMouseCursor)         // cursor :: Set mouse cursor

	// Input-related functions: touch
	GetTouchX          = bind(&cint, "GetTouchX")                   // Get touch position X for touch point 0 (relative to screen size)
	GetTouchY          = bind(&cint, "GetTouchY")                   // Get touch position Y for touch point 0 (relative to screen size)
	GetTouchPosition   = bind(&tVector2, "GetTouchPosition", &cint) // index :: Get touch position XY for a touch point index (relative to screen size)
	GetTouchPointId    = bind(&cint, "GetTouchPointId", &cint)      // index :: Get touch point identifier for given index
	GetTouchPointCount = bind(&cint, "GetTouchPointCount")          // Get number of touch points

	//------------------------------------------------------------------------------------
	// Gestures and Touch Handling Functions (Module: rgestures)
	//------------------------------------------------------------------------------------
	SetGesturesEnabled     = bind(&void, "SetGesturesEnabled", &tGesture) // flags :: Enable a set of gestures using flags
	IsGestureDetected      = bind(&cbool, "IsGestureDetected", &tGesture) // gesture :: Check if a gesture have been detected
	GetGestureDetected     = bind(&cint, "GetGestureDetected")            // Get latest detected gesture
	GetGestureHoldDuration = bind(&float, "GetGestureHoldDuration")       // Get gesture hold time in seconds
	GetGestureDragVector   = bind(&tVector2, "GetGestureDragVector")      // Get gesture drag vector
	GetGestureDragAngle    = bind(&float, "GetGestureDragAngle")          // Get gesture drag angle
	GetGesturePinchVector  = bind(&tVector2, "GetGesturePinchVector")     // Get gesture pinch delta
	GetGesturePinchAngle   = bind(&float, "GetGesturePinchAngle")         // Get gesture pinch angle

	//------------------------------------------------------------------------------------
	// Camera System Functions (Module: rcamera)
	//------------------------------------------------------------------------------------
	UpdateCamera    = bind(&void, "UpdateCamera", &tCameraPtr, &tCameraMode)                    // camera, mode :: Update camera position for selected mode
	UpdateCameraPro = bind(&void, "UpdateCameraPro", &tCameraPtr, &tVector3, &tVector3, &float) // camera, movement, rotation, zoom :: Update camera movement/rotation
)
