package raylib

var (
	// Window-related functions
	initWindow               = prep(&void, "InitWindow", &cint, &cint, &cstr)   // width, height, title :: Initialize window and OpenGL context
	closeWindow              = prep(&void, "CloseWindow")                       // Close window and unload OpenGL context
	windowShouldClose        = prep(&cbool, "WindowShouldClose")                // Check if application should close (KEY_ESCAPE pressed or windows close icon clicked)
	isWindowReady            = prep(&cbool, "IsWindowReady")                    // Check if window has been initialized successfully
	isWindowFullscreen       = prep(&cbool, "IsWindowFullscreen")               // Check if window is currently fullscreen
	isWindowHidden           = prep(&cbool, "IsWindowHidden")                   // Check if window is currently hidden
	isWindowMinimized        = prep(&cbool, "IsWindowMinimized")                // Check if window is currently minimized
	isWindowMaximized        = prep(&cbool, "IsWindowMaximized")                // Check if window is currently maximized
	isWindowFocused          = prep(&cbool, "IsWindowFocused")                  // Check if window is currently focused
	isWindowResized          = prep(&cbool, "IsWindowResized")                  // Check if window has been resized last frame
	isWindowState            = prep(&cbool, "IsWindowState", &ucint)            // flag :: Check if one specific window flag is enabled
	setWindowState           = prep(&void, "SetWindowState", &ucint)            // flags :: Set window configuration state using flags
	clearWindowState         = prep(&void, "ClearWindowState", &ucint)          // flags :: Clear window configuration state flags
	toggleFullscreen         = prep(&void, "ToggleFullscreen")                  // Toggle window state: fullscreen/windowed, resizes monitor to match window resolution
	toggleBorderlessWindowed = prep(&void, "ToggleBorderlessWindowed")          // Toggle window state: borderless windowed, resizes window to match monitor resolution
	maximizeWindow           = prep(&void, "MaximizeWindow")                    // Set window state: maximized, if resizable
	minimizeWindow           = prep(&void, "MinimizeWindow")                    // Set window state: minimized, if resizable
	restoreWindow            = prep(&void, "RestoreWindow")                     // Set window state: not minimized/maximized
	setWindowIcon            = prep(&void, "SetWindowIcon", &tImage)            // image :: Set icon for window (single image, RGBA 32bit)
	setWindowIcons           = prep(&void, "SetWindowIcons", &tImagePtr, &cint) // images, count :: Set icon for window (multiple images, RGBA 32bit)
	setWindowTitle           = prep(&void, "SetWindowTitle", &cstr)             // title :: Set title for window
	setWindowPosition        = prep(&void, "SetWindowPosition", &cint, &cint)   // x, y :: Set window position on screen
	setWindowMonitor         = prep(&void, "SetWindowMonitor", &cint)           // monitor :: Set monitor for the current window
	setWindowMinSize         = prep(&void, "SetWindowMinSize", &cint, &cint)    // width, height :: Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
	setWindowMaxSize         = prep(&void, "SetWindowMaxSize", &cint, &cint)    // width, height :: Set window maximum dimensions (for FLAG_WINDOW_RESIZABLE)
	setWindowSize            = prep(&void, "SetWindowSize", &cint, &cint)       // width, height :: Set window dimensions
	setWindowOpacity         = prep(&void, "SetWindowOpacity", &float)          // opacity :: Set window opacity [0.0f..1.0f]
	setWindowFocused         = prep(&void, "SetWindowFocused")                  // Set window focused
	getWindowHandle          = prep(&ptr, "GetWindowHandle")                    // Get native window handle
	getScreenWidth           = prep(&cint, "GetScreenWidth")                    // Get current screen width
	getScreenHeight          = prep(&cint, "GetScreenHeight")                   // Get current screen height
	getRenderWidth           = prep(&cint, "GetRenderWidth")                    // Get current render width (it considers HiDPI)
	getRenderHeight          = prep(&cint, "GetRenderHeight")                   // Get current render height (it considers HiDPI)
	getMonitorCount          = prep(&cint, "GetMonitorCount")                   // Get number of connected monitors
	getCurrentMonitor        = prep(&cint, "GetCurrentMonitor")                 // Get current monitor where window is placed
	getMonitorPosition       = prep(&tVector2, "GetMonitorPosition", &cint)     // monitor :: Get specified monitor position
	getMonitorWidth          = prep(&cint, "GetMonitorWidth", &cint)            // monitor :: Get specified monitor width (current video mode used by monitor)
	getMonitorHeight         = prep(&cint, "GetMonitorHeight", &cint)           // monitor :: Get specified monitor height (current video mode used by monitor)
	getMonitorPhysicalWidth  = prep(&cint, "GetMonitorPhysicalWidth", &cint)    // monitor :: Get specified monitor physical width in millimetres
	getMonitorPhysicalHeight = prep(&cint, "GetMonitorPhysicalHeight", &cint)   // monitor :: Get specified monitor physical height in millimetres
	getMonitorRefreshRate    = prep(&cint, "GetMonitorRefreshRate", &cint)      // monitor :: Get specified monitor refresh rate
	getWindowPosition        = prep(&tVector2, "GetWindowPosition")             // Get window position XY on monitor
	getWindowScaleDPI        = prep(&tVector2, "GetWindowScaleDPI")             // Get window scale DPI factor
	getMonitorName           = prep(&cstr, "GetMonitorName", &cint)             // monitor :: Get the human-readable, UTF-8 encoded name of the specified monitor
	setClipboardText         = prep(&void, "SetClipboardText", &cstr)           // text :: Set clipboard text content
	getClipboardText         = prep(&cstr, "GetClipboardText")                  // Get clipboard text content
	getClipboardImage        = prep(&tImage, "GetClipboardImage")               // Get clipboard image
	enableEventWaiting       = prep(&void, "EnableEventWaiting")                // Enable waiting for events on EndDrawing(), no automatic event polling
	disableEventWaiting      = prep(&void, "DisableEventWaiting")               // Disable waiting for events on EndDrawing(), automatic events polling

	// Cursor-related functions
	showCursor       = prep(&void, "ShowCursor")        // Shows cursor
	hideCursor       = prep(&void, "HideCursor")        // Hides cursor
	isCursorHidden   = prep(&cbool, "IsCursorHidden")   // Check if cursor is not visible
	enableCursor     = prep(&void, "EnableCursor")      // Enables cursor (unlock cursor)
	disableCursor    = prep(&void, "DisableCursor")     // Disables cursor (lock cursor)
	isCursorOnScreen = prep(&cbool, "IsCursorOnScreen") // Check if cursor is on the screen

	// Drawing-related functions
	clearBackground  = prep(&void, "ClearBackground", &tColor)                     // color :: Set background color (framebuffer clear color)
	beginDrawing     = prep(&void, "BeginDrawing")                                 // Setup canvas (framebuffer) to start drawing
	endDrawing       = prep(&void, "EndDrawing")                                   // End canvas drawing and swap buffers (double buffering)
	beginMode2D      = prep(&void, "BeginMode2D", &tCamera2D)                      // camera :: Begin 2D mode with custom camera (2D)
	endMode2D        = prep(&void, "EndMode2D")                                    // Ends 2D mode with custom camera
	beginMode3D      = prep(&void, "BeginMode3D", &tCamera3D)                      // camera :: Begin 3D mode with custom camera (3D)
	endMode3D        = prep(&void, "EndMode3D")                                    // Ends 3D mode and returns to default 2D orthographic mode
	beginTextureMode = prep(&void, "BeginTextureMode", &tRenderTexture2D)          // target :: Begin drawing to render texture
	endTextureMode   = prep(&void, "EndTextureMode")                               // Ends drawing to render texture
	beginShaderMode  = prep(&void, "BeginShaderMode", &tShader)                    // shader :: Begin custom shader drawing
	endShaderMode    = prep(&void, "EndShaderMode")                                // End custom shader drawing (use default shader)
	beginBlendMode   = prep(&void, "BeginBlendMode", &cint)                        // mode :: Begin blending mode (alpha, additive, multiplied, subtract, custom)
	endBlendMode     = prep(&void, "EndBlendMode")                                 // End blending mode (reset to default: alpha blending)
	beginScissorMode = prep(&void, "BeginScissorMode", &cint, &cint, &cint, &cint) // x, y, width, height :: Begin scissor mode (define screen area for following drawing)
	endScissorMode   = prep(&void, "EndScissorMode")                               // End scissor mode

	// beginVrStereoMode = prep(&void, "BeginVrStereoMode", &tVrStereoConfig)      // config :: Begin stereo rendering (requires VR simulator)
	// endVrStereoMode = prep(&void, "EndVrStereoMode") // End stereo rendering (requires VR simulator)

	// Timing-related functions
	setTargetFPS = prep(&void, "SetTargetFPS", &cint) // fps :: Set target FPS (maximum)
	getFrameTime = prep(&float, "GetFrameTime")       // Get time in seconds for last frame drawn (delta time)
	getTime      = prep(&double, "GetTime")           // Get elapsed time in seconds since InitWindow()
	getFPS       = prep(&cint, "GetFPS")              // Get current FPS
)
