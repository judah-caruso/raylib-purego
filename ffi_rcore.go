package raylib

import "github.com/judah-caruso/raylib-purego/internal"

var (
	// Window-related functions
	initWindow               = internal.Bind(&void, "InitWindow", &cint, &cint, &cstr)   // width, height, title :: Initialize window and OpenGL context
	closeWindow              = internal.Bind(&void, "CloseWindow")                       // Close window and unload OpenGL context
	windowShouldClose        = internal.Bind(&cbool, "WindowShouldClose")                // Check if application should close (KEY_ESCAPE pressed or windows close icon clicked)
	isWindowReady            = internal.Bind(&cbool, "IsWindowReady")                    // Check if window has been initialized successfully
	isWindowFullscreen       = internal.Bind(&cbool, "IsWindowFullscreen")               // Check if window is currently fullscreen
	isWindowHidden           = internal.Bind(&cbool, "IsWindowHidden")                   // Check if window is currently hidden
	isWindowMinimized        = internal.Bind(&cbool, "IsWindowMinimized")                // Check if window is currently minimized
	isWindowMaximized        = internal.Bind(&cbool, "IsWindowMaximized")                // Check if window is currently maximized
	isWindowFocused          = internal.Bind(&cbool, "IsWindowFocused")                  // Check if window is currently focused
	isWindowResized          = internal.Bind(&cbool, "IsWindowResized")                  // Check if window has been resized last frame
	isWindowState            = internal.Bind(&cbool, "IsWindowState", &tConfigFlags)     // flag :: Check if one specific window flag is enabled
	setWindowState           = internal.Bind(&void, "SetWindowState", &tConfigFlags)     // flags :: Set window configuration state using flags
	clearWindowState         = internal.Bind(&void, "ClearWindowState", &tConfigFlags)   // flags :: Clear window configuration state flags
	toggleFullscreen         = internal.Bind(&void, "ToggleFullscreen")                  // Toggle window state: fullscreen/windowed, resizes monitor to match window resolution
	toggleBorderlessWindowed = internal.Bind(&void, "ToggleBorderlessWindowed")          // Toggle window state: borderless windowed, resizes window to match monitor resolution
	maximizeWindow           = internal.Bind(&void, "MaximizeWindow")                    // Set window state: maximized, if resizable
	minimizeWindow           = internal.Bind(&void, "MinimizeWindow")                    // Set window state: minimized, if resizable
	restoreWindow            = internal.Bind(&void, "RestoreWindow")                     // Set window state: not minimized/maximized
	setWindowIcon            = internal.Bind(&void, "SetWindowIcon", &tImage)            // image :: Set icon for window (single image, RGBA 32bit)
	setWindowIcons           = internal.Bind(&void, "SetWindowIcons", &tImagePtr, &cint) // images, count :: Set icon for window (multiple images, RGBA 32bit)
	setWindowTitle           = internal.Bind(&void, "SetWindowTitle", &cstr)             // title :: Set title for window
	setWindowPosition        = internal.Bind(&void, "SetWindowPosition", &cint, &cint)   // x, y :: Set window position on screen
	setWindowMonitor         = internal.Bind(&void, "SetWindowMonitor", &cint)           // monitor :: Set monitor for the current window
	setWindowMinSize         = internal.Bind(&void, "SetWindowMinSize", &cint, &cint)    // width, height :: Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
	setWindowMaxSize         = internal.Bind(&void, "SetWindowMaxSize", &cint, &cint)    // width, height :: Set window maximum dimensions (for FLAG_WINDOW_RESIZABLE)
	setWindowSize            = internal.Bind(&void, "SetWindowSize", &cint, &cint)       // width, height :: Set window dimensions
	setWindowOpacity         = internal.Bind(&void, "SetWindowOpacity", &float)          // opacity :: Set window opacity [0.0f..1.0f]
	setWindowFocused         = internal.Bind(&void, "SetWindowFocused")                  // Set window focused
	getWindowHandle          = internal.Bind(&ptr, "GetWindowHandle")                    // Get native window handle
	getScreenWidth           = internal.Bind(&cint, "GetScreenWidth")                    // Get current screen width
	getScreenHeight          = internal.Bind(&cint, "GetScreenHeight")                   // Get current screen height
	getRenderWidth           = internal.Bind(&cint, "GetRenderWidth")                    // Get current render width (it considers HiDPI)
	getRenderHeight          = internal.Bind(&cint, "GetRenderHeight")                   // Get current render height (it considers HiDPI)
	getMonitorCount          = internal.Bind(&cint, "GetMonitorCount")                   // Get number of connected monitors
	getCurrentMonitor        = internal.Bind(&cint, "GetCurrentMonitor")                 // Get current monitor where window is placed
	getMonitorPosition       = internal.Bind(&tVector2, "GetMonitorPosition", &cint)     // monitor :: Get specified monitor position
	getMonitorWidth          = internal.Bind(&cint, "GetMonitorWidth", &cint)            // monitor :: Get specified monitor width (current video mode used by monitor)
	getMonitorHeight         = internal.Bind(&cint, "GetMonitorHeight", &cint)           // monitor :: Get specified monitor height (current video mode used by monitor)
	getMonitorPhysicalWidth  = internal.Bind(&cint, "GetMonitorPhysicalWidth", &cint)    // monitor :: Get specified monitor physical width in millimetres
	getMonitorPhysicalHeight = internal.Bind(&cint, "GetMonitorPhysicalHeight", &cint)   // monitor :: Get specified monitor physical height in millimetres
	getMonitorRefreshRate    = internal.Bind(&cint, "GetMonitorRefreshRate", &cint)      // monitor :: Get specified monitor refresh rate
	getWindowPosition        = internal.Bind(&tVector2, "GetWindowPosition")             // Get window position XY on monitor
	getWindowScaleDPI        = internal.Bind(&tVector2, "GetWindowScaleDPI")             // Get window scale DPI factor
	getMonitorName           = internal.Bind(&cstr, "GetMonitorName", &cint)             // monitor :: Get the human-readable, UTF-8 encoded name of the specified monitor
	setClipboardText         = internal.Bind(&void, "SetClipboardText", &cstr)           // text :: Set clipboard text content
	getClipboardText         = internal.Bind(&cstr, "GetClipboardText")                  // Get clipboard text content
	getClipboardImage        = internal.Bind(&tImage, "GetClipboardImage")               // Get clipboard image
	enableEventWaiting       = internal.Bind(&void, "EnableEventWaiting")                // Enable waiting for events on EndDrawing(), no automatic event polling
	disableEventWaiting      = internal.Bind(&void, "DisableEventWaiting")               // Disable waiting for events on EndDrawing(), automatic events polling

	// Cursor-related functions
	showCursor       = internal.Bind(&void, "ShowCursor")        // Shows cursor
	hideCursor       = internal.Bind(&void, "HideCursor")        // Hides cursor
	isCursorHidden   = internal.Bind(&cbool, "IsCursorHidden")   // Check if cursor is not visible
	enableCursor     = internal.Bind(&void, "EnableCursor")      // Enables cursor (unlock cursor)
	disableCursor    = internal.Bind(&void, "DisableCursor")     // Disables cursor (lock cursor)
	isCursorOnScreen = internal.Bind(&cbool, "IsCursorOnScreen") // Check if cursor is on the screen

	// Drawing-related functions
	clearBackground  = internal.Bind(&void, "ClearBackground", &tColor)                     // color :: Set background color (framebuffer clear color)
	beginDrawing     = internal.Bind(&void, "BeginDrawing")                                 // Setup canvas (framebuffer) to start drawing
	endDrawing       = internal.Bind(&void, "EndDrawing")                                   // End canvas drawing and swap buffers (double buffering)
	beginMode2D      = internal.Bind(&void, "BeginMode2D", &tCamera2D)                      // camera :: Begin 2D mode with custom camera (2D)
	endMode2D        = internal.Bind(&void, "EndMode2D")                                    // Ends 2D mode with custom camera
	beginMode3D      = internal.Bind(&void, "BeginMode3D", &tCamera3D)                      // camera :: Begin 3D mode with custom camera (3D)
	endMode3D        = internal.Bind(&void, "EndMode3D")                                    // Ends 3D mode and returns to default 2D orthographic mode
	beginTextureMode = internal.Bind(&void, "BeginTextureMode", &tRenderTexture2D)          // target :: Begin drawing to render texture
	endTextureMode   = internal.Bind(&void, "EndTextureMode")                               // Ends drawing to render texture
	beginShaderMode  = internal.Bind(&void, "BeginShaderMode", &tShader)                    // shader :: Begin custom shader drawing
	endShaderMode    = internal.Bind(&void, "EndShaderMode")                                // End custom shader drawing (use default shader)
	beginBlendMode   = internal.Bind(&void, "BeginBlendMode", &cint)                        // mode :: Begin blending mode (alpha, additive, multiplied, subtract, custom)
	endBlendMode     = internal.Bind(&void, "EndBlendMode")                                 // End blending mode (reset to default: alpha blending)
	beginScissorMode = internal.Bind(&void, "BeginScissorMode", &cint, &cint, &cint, &cint) // x, y, width, height :: Begin scissor mode (define screen area for following drawing)
	endScissorMode   = internal.Bind(&void, "EndScissorMode")                               // End scissor mode

	beginVrStereoMode = internal.Bind(&void, "BeginVrStereoMode", &tVrStereoConfig) // config :: Begin stereo rendering (requires VR simulator)
	endVrStereoMode   = internal.Bind(&void, "EndVrStereoMode")                     // End stereo rendering (requires VR simulator)

	// VR stereo config functions for VR simulator
	loadVrStereoConfig   = internal.Bind(&tVrStereoConfig, "LoadVrStereoConfig", &tVrDeviceInfo) // device :: Load VR stereo config for VR simulator device parameters
	unloadVrStereoConfig = internal.Bind(&void, "UnloadVrStereoConfig", &tVrStereoConfig)        // config :: Unload VR stereo config

	// Shader management functions
	// NOTE: Shader functionality is not available on OpenGL 1.1
	loadShader              = internal.Bind(&tShader, "LoadShader", &cstr, &cstr)                          // vsFileName, fsFileName :: Load shader from files and bind default locations
	loadShaderFromMemory    = internal.Bind(&tShader, "LoadShaderFromMemory", &cstr, &cstr)                // vsCode, fsCode :: Load shader from code strings and bind default locations
	isShaderValid           = internal.Bind(&cbool, "IsShaderValid", &tShader)                             // shader :: Check if a shader is valid (loaded on GPU)
	getShaderLocation       = internal.Bind(&cint, "GetShaderLocation", &tShader, &cstr)                   // shader, uniformName :: Get shader uniform location
	getShaderLocationAttrib = internal.Bind(&cint, "GetShaderLocationAttrib", &tShader, &cstr)             // shader, attribName :: Get shader attribute location
	setShaderValue          = internal.Bind(&void, "SetShaderValue", &tShader, &cint, &ptr, &cint)         // shader, locIndex, value, uniformType :: Set shader uniform value
	setShaderValueV         = internal.Bind(&void, "SetShaderValueV", &tShader, &cint, &ptr, &cint, &cint) // shader, locIndex, value, uniformType, count :: Set shader uniform value vector
	setShaderValueMatrix    = internal.Bind(&void, "SetShaderValueMatrix", &tShader, &cint, &tMatrix)      // shader, locIndex, mat :: Set shader uniform value (matrix 4x4)
	setShaderValueTexture   = internal.Bind(&void, "SetShaderValueTexture", &tShader, &cint, &tTexture2D)  // shader, locIndex, texture :: Set shader uniform value for texture (sampler2d)
	unloadShader            = internal.Bind(&void, "UnloadShader", &tShader)                               // shader :: Unload shader from GPU memory (VRAM)

	// Screen-space-related functions
	getScreenToWorldRay   = internal.Bind(&tRay, "GetScreenToWorldRay", &tVector2, &tCamera)                  // position, camera :: Get a ray trace from screen position (i.e mouse)
	getScreenToWorldRayEx = internal.Bind(&tRay, "GetScreenToWorldRayEx", &tVector2, &tCamera, &cint, &cint)  // position, camera, width, height :: Get a ray trace from screen position (i.e mouse) in a viewport
	getWorldToScreen      = internal.Bind(&tVector2, "GetWorldToScreen", &tVector3, &tCamera)                 // position, camera :: Get the screen space position for a 3d world space position
	getWorldToScreenEx    = internal.Bind(&tVector2, "GetWorldToScreenEx", &tVector3, &tCamera, &cint, &cint) // position, camera, width, height :: Get size position for a 3d world space position
	getWorldToScreen2D    = internal.Bind(&tVector2, "GetWorldToScreen2D", &tVector2, &tCamera2D)             // position, camera :: Get the screen space position for a 2d camera world space position
	getScreenToWorld2D    = internal.Bind(&tVector2, "GetScreenToWorld2D", &tVector2, &tCamera2D)             // position, camera :: Get the world space position for a 2d camera screen space position
	getCameraMatrix       = internal.Bind(&tMatrix, "GetCameraMatrix", &tCamera)                              // camera :: Get camera transform matrix (view matrix)
	getCameraMatrix2D     = internal.Bind(&tMatrix, "GetCameraMatrix2D", &tCamera2D)                          // camera :: Get camera 2d transform matrix

	// Timing-related functions
	setTargetFPS = internal.Bind(&void, "SetTargetFPS", &cint) // fps :: Set target FPS (maximum)
	getFrameTime = internal.Bind(&float, "GetFrameTime")       // Get time in seconds for last frame drawn (delta time)
	getTime      = internal.Bind(&double, "GetTime")           // Get elapsed time in seconds since InitWindow()
	getFPS       = internal.Bind(&cint, "GetFPS")              // Get current FPS

	// Custom frame control functions
	// NOTE: Those functions are intended for advanced users that want full control over the frame processing
	// By default EndDrawing() does this job: draws everything + SwapScreenBuffer() + manage frame timing + PollInputEvents()
	// To avoid that behaviour and control frame processes manually, enable in config.h: SUPPORT_CUSTOM_FRAME_CONTROL
	swapScreenBuffer = internal.Bind(&void, "SwapScreenBuffer")  // Swap back buffer with front buffer (screen drawing)
	pollInputEvents  = internal.Bind(&void, "PollInputEvents")   // Register all input events
	waitTime         = internal.Bind(&void, "WaitTime", &double) // seconds :: Wait for some time (halt program execution)

	// Random values generation functions
	setRandomSeed        = internal.Bind(&void, "SetRandomSeed", &ucint)                    // seed :: Set the seed for the random number generator
	getRandomValue       = internal.Bind(&cint, "GetRandomValue", &cint, &cint)             // min, max :: Get a random value between min and max (both included)
	loadRandomSequence   = internal.Bind(&cint, "LoadRandomSequence", &ucint, &cint, &cint) // count, min, max :: Load random values sequence, no values repeated
	unloadRandomSequence = internal.Bind(&void, "UnloadRandomSequence", &ptr)               // sequence :: Unload random values sequence

	// Misc. functions
	takeScreenshot = internal.Bind(&void, "TakeScreenshot", &cstr)         // fileName :: Takes a screenshot of current screen (filename extension defines format)
	setConfigFlags = internal.Bind(&void, "SetConfigFlags", &tConfigFlags) // flags :: Setup init configuration flags (view FLAGS)
	openURL        = internal.Bind(&void, "OpenURL", &cstr)                // url :: Open URL with default system browser (if available)

	// NOTE: Following functions implemented in module [utils]
	// ------------------------------------------------------------------
	// @todo(judah): support varargs?
	// void TraceLog(int logLevel, const char *text, ...);         // Show trace log messages (LOG_DEBUG, LOG_INFO, LOG_WARNING, LOG_ERROR...)

	setTraceLogLevel = internal.Bind(&void, "SetTraceLogLevel", &cint) // logLevel :: Set the current threshold (minimum) log level
	memAlloc         = internal.Bind(&void, "MemAlloc", &ucint)        // size :: Internal memory allocator
	memRealloc       = internal.Bind(&void, "MemRealloc", &ptr, &cint) // ptr, size :: Internal memory reallocator
	memFree          = internal.Bind(&void, "MemFree", &ptr)           // Internal memory free

	// @todo(judah): callbacks

	// Set custom callbacks
	// WARNING: Callbacks setup is intended for advanced users
	// void SetTraceLogCallback(TraceLogCallback callback);         // Set custom trace log
	// void SetLoadFileDataCallback(LoadFileDataCallback callback); // Set custom file binary data loader
	// void SetSaveFileDataCallback(SaveFileDataCallback callback); // Set custom file binary data saver
	// void SetLoadFileTextCallback(LoadFileTextCallback callback); // Set custom file text data loader
	// void SetSaveFileTextCallback(SaveFileTextCallback callback); // Set custom file text data saver

	// Files management functions
	loadFileData     = internal.Bind(&ustr, "LoadFileData", &cstr, &intPtr)           // fileName, dataSize :: Load file data as byte array (read)
	unloadFileData   = internal.Bind(&void, "UnloadFileData", &cstr)                  // data :: Unload file data allocated by LoadFileData()
	saveFileData     = internal.Bind(&cbool, "SaveFileData", &cstr, &ptr, &cint)      // fileName, data, dataSize :: Save data to file from byte array (write), returns true on success
	exportDataAsCode = internal.Bind(&cbool, "ExportDataAsCode", &ustr, &cint, &cstr) // data, dataSize, fileName :: Export data to code (.h), returns true on success
	loadFileText     = internal.Bind(&cstr, "LoadFileText", &cstr)                    // fileName :: Load text data from file (read), returns a '\0' terminated string
	unloadFileText   = internal.Bind(&void, "UnloadFileText", &cstr)                  // text :: Unload file text data allocated by LoadFileText()
	saveFileText     = internal.Bind(&cbool, "SaveFileText", &cstr, &cstr)            // fileName, text :: Save text data to file (write), string must be '\0' terminated, returns true on success

	// File system functions
	fileExists              = internal.Bind(&cbool, "FileExists", &cstr)                                  // fileName :: Check if file exists
	directoryExists         = internal.Bind(&cbool, "DirectoryExists", &cstr)                             // dirPath :: Check if a directory path exists
	isFileExtension         = internal.Bind(&cbool, "IsFileExtension", &cstr, &cstr)                      // fileName, ext :: Check file extension (including point: .png, .wav)
	getFileLength           = internal.Bind(&cint, "GetFileLength", &cstr)                                // fileName :: Get file length in bytes (NOTE: GetFileSize() conflicts with windows.h)
	getFileExtension        = internal.Bind(&cstr, "GetFileExtension", &cstr)                             // fileName :: Get pointer to extension for a filename string (includes dot: '.png')
	getFileName             = internal.Bind(&cstr, "GetFileName", &cstr)                                  // filePath :: Get pointer to filename for a path string
	getFileNameWithoutExt   = internal.Bind(&cstr, "GetFileNameWithoutExt", &cstr)                        // filePath :: Get filename string without extension (uses static string)
	getDirectoryPath        = internal.Bind(&cstr, "GetDirectoryPath", &cstr)                             // filePath :: Get full path for a given fileName with path (uses static string)
	getPrevDirectoryPath    = internal.Bind(&cstr, "GetPrevDirectoryPath", &cstr)                         // dirPath :: Get previous directory path for a given path (uses static string)
	getWorkingDirectory     = internal.Bind(&cstr, "GetWorkingDirectory")                                 // Get current working directory (uses static string)
	getApplicationDirectory = internal.Bind(&cstr, "GetApplicationDirectory")                             // Get the directory of the running application (uses static string)
	makeDirectory           = internal.Bind(&cint, "MakeDirectory", &cstr)                                // dirPath :: Create directories (including full path requested), returns 0 on success
	changeDirectory         = internal.Bind(&cbool, "ChangeDirectory", &cstr)                             // dir :: Change working directory, return true on success
	isPathFile              = internal.Bind(&cbool, "IsPathFile", &cstr)                                  // path :: Check if a given path is a file or a directory
	isFileNameValid         = internal.Bind(&cbool, "IsFileNameValid", &cstr)                             // fileName :: Check if fileName is valid for the platform/OS
	loadDirectoryFiles      = internal.Bind(&tFilePathList, "LoadDirectoryFiles", &cstr)                  // dirPath :: Load directory filepaths
	loadDirectoryFilesEx    = internal.Bind(&tFilePathList, "LoadDirectoryFilesEx", &cstr, &cstr, &cbool) // basePath, filter, scanSubdirs :: Load directory filepaths with extension filtering and recursive directory scan. Use 'DIR' in the filter string to include directories in the result
	unloadDirectoryFiles    = internal.Bind(&void, "UnloadDirectoryFiles", &tFilePathList)                // files :: Unload filepaths
	isFileDropped           = internal.Bind(&cbool, "IsFileDropped")                                      // Check if a file has been dropped into window
	loadDroppedFiles        = internal.Bind(&tFilePathList, "LoadDroppedFiles")                           // Load dropped filepaths
	unloadDroppedFiles      = internal.Bind(&void, "UnloadDroppedFiles", &tFilePathList)                  // files :: Unload dropped filepaths
	getFileModTime          = internal.Bind(&long, "GetFileModTime", &cstr)                               // fileName :: Get file modification time (last write time)

	// Compression/Encoding functionality
	compressData     = internal.Bind(&ustr, "CompressData", &ustr, &cint, &intPtr)     // data, dataSize, compDataSize :: Compress data (DEFLATE algorithm), memory must be MemFree()
	decompressData   = internal.Bind(&ustr, "DecompressData", &ustr, &cint, &intPtr)   // compData, compDataSize, dataSize :: Decompress data (DEFLATE algorithm), memory must be MemFree()
	encodeDataBase64 = internal.Bind(&cstr, "EncodeDataBase64", &ustr, &cint, &intPtr) // data, dataSize, outputSize :: Encode data to Base64 string, memory must be MemFree()
	decodeDataBase64 = internal.Bind(&ustr, "DecodeDataBase64", &ustr, &intPtr)        // data, outputSize :: Decode Base64 string data, memory must be MemFree()
	computeCRC32     = internal.Bind(&ucint, "ComputeCRC32", &ustr, &cint)             // data, dataSize :: Compute CRC32 hash code
	computeMD5       = internal.Bind(&uintPtr, "ComputeMD5", &ustr, &cint)             // data, dataSize :: Compute MD5 hash code, returns static int[4] (16 bytes)
	computeSHA1      = internal.Bind(&uintPtr, "ComputeSHA1", &ustr, &cint)            // data, dataSize :: Compute SHA1 hash code, returns static int[5] (20 bytes)

	// Automation events functionality
	loadAutomationEventList       = internal.Bind(&tAutomationEventList, "LoadAutomationEventList", &cstr)           // fileName :: Load automation events list from file, NULL for empty list, capacity = MAX_AUTOMATION_EVENTS
	unloadAutomationEventList     = internal.Bind(&void, "UnloadAutomationEventList", &tAutomationEventList)         // list :: Unload automation events list from file
	exportAutomationEventList     = internal.Bind(&cbool, "ExportAutomationEventList", &tAutomationEventList, &cstr) // list, fileName :: Export automation events list as text file
	setAutomationEventList        = internal.Bind(&void, "SetAutomationEventList", &tAutomationEventListPtr)         // list :: Set automation event list to record to
	setAutomationEventBaseFrame   = internal.Bind(&void, "SetAutomationEventBaseFrame", &cint)                       // frame :: Set automation event internal base frame to start recording
	startAutomationEventRecording = internal.Bind(&void, "StartAutomationEventRecording")                            // Start recording automation events (AutomationEventList must be set)
	stopAutomationEventRecording  = internal.Bind(&void, "StopAutomationEventRecording")                             // Stop recording automation events
	playAutomationEvent           = internal.Bind(&void, "PlayAutomationEvent", &tAutomationEvent)                   // event :: Play a recorded automation event

	// Input-related functions: keyboard
	isKeyPressed       = internal.Bind(&cbool, "IsKeyPressed", &tKeyboardKey)       // key :: Check if a key has been pressed once
	isKeyPressedRepeat = internal.Bind(&cbool, "IsKeyPressedRepeat", &tKeyboardKey) // key :: Check if a key has been pressed again
	isKeyDown          = internal.Bind(&cbool, "IsKeyDown", &tKeyboardKey)          // key :: Check if a key is being pressed
	isKeyReleased      = internal.Bind(&cbool, "IsKeyReleased", &tKeyboardKey)      // key :: Check if a key has been released once
	isKeyUp            = internal.Bind(&cbool, "IsKeyUp", &tKeyboardKey)            // key :: Check if a key is NOT being pressed
	getKeyPressed      = internal.Bind(&tKeyboardKey, "GetKeyPressed")              // Get key pressed (keycode), call it multiple times for keys queued, returns 0 when the queue is empty
	getCharPressed     = internal.Bind(&tKeyboardKey, "GetCharPressed")             // Get char pressed (unicode), call it multiple times for chars queued, returns 0 when the queue is empty
	setExitKey         = internal.Bind(&void, "SetExitKey", &tKeyboardKey)          // key :: Set a custom key to exit program (default is ESC)

	// Input-related functions: gamepads
	isGamepadAvailable      = internal.Bind(&cbool, "IsGamepadAvailable", &cint)                         // gamepad :: Check if a gamepad is available
	getGamepadName          = internal.Bind(&cstr, "GetGamepadName", &cint)                              // gamepad :: Get gamepad internal name id
	isGamepadButtonPressed  = internal.Bind(&cbool, "IsGamepadButtonPressed", &cint, &tGamepadButton)    // gamepad, button :: Check if a gamepad button has been pressed once
	isGamepadButtonDown     = internal.Bind(&cbool, "IsGamepadButtonDown", &cint, &tGamepadButton)       // gamepad, button :: Check if a gamepad button is being pressed
	isGamepadButtonReleased = internal.Bind(&cbool, "IsGamepadButtonReleased", &cint, &tGamepadButton)   // gamepad, button :: Check if a gamepad button has been released once
	isGamepadButtonUp       = internal.Bind(&cbool, "IsGamepadButtonUp", &cint, &tGamepadButton)         // gamepad, button :: Check if a gamepad button is NOT being pressed
	getGamepadButtonPressed = internal.Bind(&cint, "GetGamepadButtonPressed")                            // Get the last gamepad button pressed
	getGamepadAxisCount     = internal.Bind(&cint, "GetGamepadAxisCount", &cint)                         // gamepad :: Get gamepad axis count for a gamepad
	getGamepadAxisMovement  = internal.Bind(&float, "GetGamepadAxisMovement", &cint, &tGamepadAxis)      // gamepad, axis :: Get axis movement value for a gamepad axis
	setGamepadMappings      = internal.Bind(&cint, "SetGamepadMappings", &cstr)                          // mappings :: Set internal gamepad mappings (SDL_GameControllerDB)
	setGamepadVibration     = internal.Bind(&void, "SetGamepadVibration", &cint, &float, &float, &float) // gamepad, leftMotor, rightMotor, duration :: Set gamepad vibration for both motors (duration in seconds)

	// Input-related functions: mouse
	isMouseButtonPressed  = internal.Bind(&cbool, "IsMouseButtonPressed", &tMouseButton)  // button :: Check if a mouse button has been pressed once
	isMouseButtonDown     = internal.Bind(&cbool, "IsMouseButtonDown", &tMouseButton)     // button :: Check if a mouse button is being pressed
	isMouseButtonReleased = internal.Bind(&cbool, "IsMouseButtonReleased", &tMouseButton) // button :: Check if a mouse button has been released once
	isMouseButtonUp       = internal.Bind(&cbool, "IsMouseButtonUp", &tMouseButton)       // button :: Check if a mouse button is NOT being pressed
	getMouseX             = internal.Bind(&cint, "GetMouseX")                             // Get mouse position X
	getMouseY             = internal.Bind(&cint, "GetMouseY")                             // Get mouse position Y
	getMousePosition      = internal.Bind(&tVector2, "GetMousePosition")                  // Get mouse position XY
	getMouseDelta         = internal.Bind(&tVector2, "GetMouseDelta")                     // Get mouse delta between frames
	setMousePosition      = internal.Bind(&void, "SetMousePosition", &cint, &cint)        // x, y :: Set mouse position XY
	setMouseOffset        = internal.Bind(&void, "SetMouseOffset", &cint, &cint)          // offsetX, offsetY :: Set mouse offset
	setMouseScale         = internal.Bind(&void, "SetMouseScale", &float, &float)         // scaleX, scaleY :: Set mouse scaling
	getMouseWheelMove     = internal.Bind(&float, "GetMouseWheelMove")                    // Get mouse wheel movement for X or Y, whichever is larger
	getMouseWheelMoveV    = internal.Bind(&tVector2, "GetMouseWheelMoveV")                // Get mouse wheel movement for both X and Y
	setMouseCursor        = internal.Bind(&void, "SetMouseCursor", &tMouseCursor)         // cursor :: Set mouse cursor

	// Input-related functions: touch
	getTouchX          = internal.Bind(&cint, "GetTouchX")                   // Get touch position X for touch point 0 (relative to screen size)
	getTouchY          = internal.Bind(&cint, "GetTouchY")                   // Get touch position Y for touch point 0 (relative to screen size)
	getTouchPosition   = internal.Bind(&tVector2, "GetTouchPosition", &cint) // index :: Get touch position XY for a touch point index (relative to screen size)
	getTouchPointId    = internal.Bind(&cint, "GetTouchPointId", &cint)      // index :: Get touch point identifier for given index
	getTouchPointCount = internal.Bind(&cint, "GetTouchPointCount")          // Get number of touch points

	//------------------------------------------------------------------------------------
	// Gestures and Touch Handling Functions (Module: rgestures)
	//------------------------------------------------------------------------------------
	setGesturesEnabled     = internal.Bind(&void, "SetGesturesEnabled", &tGesture) // flags :: Enable a set of gestures using flags
	isGestureDetected      = internal.Bind(&cbool, "IsGestureDetected", &tGesture) // gesture :: Check if a gesture have been detected
	getGestureDetected     = internal.Bind(&cint, "GetGestureDetected")            // Get latest detected gesture
	getGestureHoldDuration = internal.Bind(&float, "GetGestureHoldDuration")       // Get gesture hold time in seconds
	getGestureDragVector   = internal.Bind(&tVector2, "GetGestureDragVector")      // Get gesture drag vector
	getGestureDragAngle    = internal.Bind(&float, "GetGestureDragAngle")          // Get gesture drag angle
	getGesturePinchVector  = internal.Bind(&tVector2, "GetGesturePinchVector")     // Get gesture pinch delta
	getGesturePinchAngle   = internal.Bind(&float, "GetGesturePinchAngle")         // Get gesture pinch angle

	//------------------------------------------------------------------------------------
	// Camera System Functions (Module: rcamera)
	//------------------------------------------------------------------------------------
	updateCamera    = internal.Bind(&void, "UpdateCamera", &tCameraPtr, &tCameraMode)                    // camera, mode :: Update camera position for selected mode
	updateCameraPro = internal.Bind(&void, "UpdateCameraPro", &tCameraPtr, &tVector3, &tVector3, &float) // camera, movement, rotation, zoom :: Update camera movement/rotation
)
