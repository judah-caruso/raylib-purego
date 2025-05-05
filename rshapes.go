package raylib

var (
	// Set texture and rectangle to be used on shapes drawing
	// NOTE: It can be useful when using basic shapes and one single font,
	// defining a font char white rectangle would allow drawing everything in a single draw call
	setShapesTexture          = prep(&void, "SetShapesTexture", &tTexture2D, &tRectangle) // texture, source :: Set texture and rectangle to be used on shapes drawing
	getShapesTexture          = prep(&tTexture2D, "GetShapesTexture")                     // Get texture that is used for shapes drawing
	getShapesTextureRectangle = prep(&tRectangle, "GetShapesTextureRectangle")            // Get texture source rectangle that is used for shapes drawing

	// Basic shapes drawing functions
	drawPixel                   = prep(&void, "DrawPixel", &cint, &cint, &tColor)                                         // posX, posY, color :: Draw a pixel using geometry [Can be slow, use with care]
	drawPixelV                  = prep(&void, "DrawPixelV", &tVector2, &tColor)                                           // position, color :: Draw a pixel using geometry (Vector version) [Can be slow, use with care]
	drawLine                    = prep(&void, "DrawLine", &cint, &cint, &cint, &cint, &tColor)                            // startPosX, startPosY, endPosX, endPosY, color :: Draw a line
	drawLineV                   = prep(&void, "DrawLineV", &tVector2, &tVector2, &tColor)                                 // startPos, endPos, color :: Draw a line (using gl lines)
	drawLineEx                  = prep(&void, "DrawLineEx", &tVector2, &tVector2, &float, &tColor)                        // startPos, endPos, thick, color :: Draw a line (using triangles/quads)
	drawLineStrip               = prep(&void, "DrawLineStrip", &tVector2Ptr, &cint, &tColor)                              // points, pointCount, color :: Draw lines sequence (using gl lines)
	drawLineBezier              = prep(&void, "DrawLineBezier", &tVector2, &tVector2, &float, &tColor)                    // startPos, endPos, thick, color :: Draw line segment cubic-bezier in-out interpolation
	drawCircle                  = prep(&void, "DrawCircle", &cint, &cint, &float, &tColor)                                // centerX, centerY, radius, color :: Draw a color-filled circle
	drawCircleSector            = prep(&void, "DrawCircleSector", &tVector2, &float, &float, &float, &cint, &tColor)      // center, radius, startAngle, endAngle, segments, color :: Draw a piece of a circle
	drawCircleSectorLines       = prep(&void, "DrawCircleSectorLines", &tVector2, &float, &float, &float, &cint, &tColor) // center, radius, startAngle, endAngle, segments, color :: Draw circle sector outline
	drawCircleGradient          = prep(&void, "DrawCircleGradient", &cint, &cint, &float, &tColor, &tColor)               // centerX, centerY, radius, inner, outer :: Draw a gradient-filled circle
	drawCircleV                 = prep(&void, "DrawCircleV", &tVector2, &float, &tColor)                                  // center, radius, color :: Draw a color-filled circle (Vector version)
	drawCircleLines             = prep(&void, "DrawCircleLines", &cint, &cint, &float, &tColor)                           // centerX, centerY, radius, color :: Draw circle outline
	drawCircleLinesV            = prep(&void, "DrawCircleLinesV", &tVector2, &float, &tColor)                             // center, radius, color :: Draw circle outline (Vector version)
	drawEllipse                 = prep(&void, "DrawEllipse", &cint, &cint, &float, &float, &tColor)                       // centerX, centerY, radiusH, radiusV, color :: Draw ellipse
	drawEllipseLines            = prep(&void, "DrawEllipseLines", &cint, &cint, &float, &float, &tColor)                  // centerX, centerY, radiusH, radiusV, color :: Draw ellipse outline
	drawRing                    = prep(&void, "DrawRing", &tVector2, &float, &float, &float, &float, &cint, &tColor)      // center, innerRadius, outerRadius, startAngle, endAngle, segments, color :: Draw ring
	drawRingLines               = prep(&void, "DrawRingLines", &tVector2, &float, &float, &float, &float, &cint, &tColor) // center, innerRadius, outerRadius, startAngle, endAngle, segments, color :: Draw ring outline
	drawRectangle               = prep(&void, "DrawRectangle", &cint, &cint, &cint, &cint, &tColor)                       // posX, posY, width, height, color :: Draw a color-filled rectangle
	drawRectangleV              = prep(&void, "DrawRectangleV", &tVector2, &tVector2, &tColor)                            // position, size, color :: Draw a color-filled rectangle (Vector version)
	drawRectangleRec            = prep(&void, "DrawRectangleRec", &tRectangle, &tColor)                                   // rec, color :: Draw a color-filled rectangle
	drawRectanglePro            = prep(&void, "DrawRectanglePro", &tRectangle, &tVector2, &float, &tColor)                // rec, origin, rotation, color :: Draw a color-filled rectangle with pro parameters
	drawRectangleGradientV      = prep(&void, "DrawRectangleGradientV", &cint, &cint, &cint, &cint, &tColor, &tColor)     // posX, posY, width, height, top, bottom :: Draw a vertical-gradient-filled rectangle
	drawRectangleGradientH      = prep(&void, "DrawRectangleGradientH", &cint, &cint, &cint, &cint, &tColor, &tColor)     // posX, posY, width, height, left, right :: Draw a horizontal-gradient-filled rectangle
	drawRectangleGradientEx     = prep(&void, "DrawRectangleGradientEx", &tRectangle, &tColor, &tColor, &tColor, &tColor) // rec, topLeft, bottomLeft, topRight, bottomRight :: Draw a gradient-filled rectangle with custom vertex colors
	drawRectangleLines          = prep(&void, "DrawRectangleLines", &cint, &cint, &cint, &cint, &tColor)                  // posX, posY, width, height, color :: Draw rectangle outline
	drawRectangleLinesEx        = prep(&void, "DrawRectangleLinesEx", &tRectangle, &float, &tColor)                       // rec, lineThick, color :: Draw rectangle outline with extended parameters
	drawRectangleRounded        = prep(&void, "DrawRectangleRounded", &tRectangle, &float, &cint, &tColor)                // rec, roundness, segments, color :: Draw rectangle with rounded edges
	drawRectangleRoundedLines   = prep(&void, "DrawRectangleRoundedLines", &tRectangle, &float, &cint, &tColor)           // rec, roundness, segments, color :: Draw rectangle lines with rounded edges
	drawRectangleRoundedLinesEx = prep(&void, "DrawRectangleRoundedLinesEx", &tRectangle, &float, &cint, &float, &tColor) // rec, roundness, segments, lineThick, color :: Draw rectangle with rounded edges outline
	drawTriangle                = prep(&void, "DrawTriangle", &tVector2, &tVector2, &tVector2, &tColor)                   // v1, v2, v3, color :: Draw a color-filled triangle (vertex in counter-clockwise order!)
	drawTriangleLines           = prep(&void, "DrawTriangleLines", &tVector2, &tVector2, &tVector2, &tColor)              // v1, v2, v3, color :: Draw triangle outline (vertex in counter-clockwise order!)
	drawTriangleFan             = prep(&void, "DrawTriangleFan", &tVector2Ptr, &cint, &tColor)                            // points, pointCount, color :: Draw a triangle fan defined by points (first vertex is the center)
	drawTriangleStrip           = prep(&void, "DrawTriangleStrip", &tVector2Ptr, &cint, &tColor)                          // points, pointCount, color :: Draw a triangle strip defined by points
	drawPoly                    = prep(&void, "DrawPoly", &tVector2, &cint, &float, &float, &tColor)                      // center, sides, radius, rotation, color :: Draw a regular polygon (Vector version)
	drawPolyLines               = prep(&void, "DrawPolyLines", &tVector2, &cint, &float, &float, &tColor)                 // center, sides, radius, rotation, color :: Draw a polygon outline of n sides
	drawPolyLinesEx             = prep(&void, "DrawPolyLinesEx", &tVector2, &cint, &float, &float, &float, &tColor)       // center, sides, radius, rotation, lineThick, color :: Draw a polygon outline of n sides with extended parameters

)
