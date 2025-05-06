package raylib

import "github.com/judah-caruso/raylib-purego/internal"

var (
	// Set texture and rectangle to be used on shapes drawing
	// NOTE: It can be useful when using basic shapes and one single font,
	// defining a font char white rectangle would allow drawing everything in a single draw call
	setShapesTexture          = internal.Bind(&void, "SetShapesTexture", &tTexture2D, &tRectangle) // texture, source :: Set texture and rectangle to be used on shapes drawing
	getShapesTexture          = internal.Bind(&tTexture2D, "GetShapesTexture")                     // Get texture that is used for shapes drawing
	getShapesTextureRectangle = internal.Bind(&tRectangle, "GetShapesTextureRectangle")            // Get texture source rectangle that is used for shapes drawing

	// Basic shapes drawing functions
	drawPixel                   = internal.Bind(&void, "DrawPixel", &cint, &cint, &tColor)                                         // posX, posY, color :: Draw a pixel using geometry [Can be slow, use with care]
	drawPixelV                  = internal.Bind(&void, "DrawPixelV", &tVector2, &tColor)                                           // position, color :: Draw a pixel using geometry (Vector version) [Can be slow, use with care]
	drawLine                    = internal.Bind(&void, "DrawLine", &cint, &cint, &cint, &cint, &tColor)                            // startPosX, startPosY, endPosX, endPosY, color :: Draw a line
	drawLineV                   = internal.Bind(&void, "DrawLineV", &tVector2, &tVector2, &tColor)                                 // startPos, endPos, color :: Draw a line (using gl lines)
	drawLineEx                  = internal.Bind(&void, "DrawLineEx", &tVector2, &tVector2, &float, &tColor)                        // startPos, endPos, thick, color :: Draw a line (using triangles/quads)
	drawLineStrip               = internal.Bind(&void, "DrawLineStrip", &tVector2Ptr, &cint, &tColor)                              // points, pointCount, color :: Draw lines sequence (using gl lines)
	drawLineBezier              = internal.Bind(&void, "DrawLineBezier", &tVector2, &tVector2, &float, &tColor)                    // startPos, endPos, thick, color :: Draw line segment cubic-bezier in-out interpolation
	drawCircle                  = internal.Bind(&void, "DrawCircle", &cint, &cint, &float, &tColor)                                // centerX, centerY, radius, color :: Draw a color-filled circle
	drawCircleSector            = internal.Bind(&void, "DrawCircleSector", &tVector2, &float, &float, &float, &cint, &tColor)      // center, radius, startAngle, endAngle, segments, color :: Draw a piece of a circle
	drawCircleSectorLines       = internal.Bind(&void, "DrawCircleSectorLines", &tVector2, &float, &float, &float, &cint, &tColor) // center, radius, startAngle, endAngle, segments, color :: Draw circle sector outline
	drawCircleGradient          = internal.Bind(&void, "DrawCircleGradient", &cint, &cint, &float, &tColor, &tColor)               // centerX, centerY, radius, inner, outer :: Draw a gradient-filled circle
	drawCircleV                 = internal.Bind(&void, "DrawCircleV", &tVector2, &float, &tColor)                                  // center, radius, color :: Draw a color-filled circle (Vector version)
	drawCircleLines             = internal.Bind(&void, "DrawCircleLines", &cint, &cint, &float, &tColor)                           // centerX, centerY, radius, color :: Draw circle outline
	drawCircleLinesV            = internal.Bind(&void, "DrawCircleLinesV", &tVector2, &float, &tColor)                             // center, radius, color :: Draw circle outline (Vector version)
	drawEllipse                 = internal.Bind(&void, "DrawEllipse", &cint, &cint, &float, &float, &tColor)                       // centerX, centerY, radiusH, radiusV, color :: Draw ellipse
	drawEllipseLines            = internal.Bind(&void, "DrawEllipseLines", &cint, &cint, &float, &float, &tColor)                  // centerX, centerY, radiusH, radiusV, color :: Draw ellipse outline
	drawRing                    = internal.Bind(&void, "DrawRing", &tVector2, &float, &float, &float, &float, &cint, &tColor)      // center, innerRadius, outerRadius, startAngle, endAngle, segments, color :: Draw ring
	drawRingLines               = internal.Bind(&void, "DrawRingLines", &tVector2, &float, &float, &float, &float, &cint, &tColor) // center, innerRadius, outerRadius, startAngle, endAngle, segments, color :: Draw ring outline
	drawRectangle               = internal.Bind(&void, "DrawRectangle", &cint, &cint, &cint, &cint, &tColor)                       // posX, posY, width, height, color :: Draw a color-filled rectangle
	drawRectangleV              = internal.Bind(&void, "DrawRectangleV", &tVector2, &tVector2, &tColor)                            // position, size, color :: Draw a color-filled rectangle (Vector version)
	drawRectangleRec            = internal.Bind(&void, "DrawRectangleRec", &tRectangle, &tColor)                                   // rec, color :: Draw a color-filled rectangle
	drawRectanglePro            = internal.Bind(&void, "DrawRectanglePro", &tRectangle, &tVector2, &float, &tColor)                // rec, origin, rotation, color :: Draw a color-filled rectangle with pro parameters
	drawRectangleGradientV      = internal.Bind(&void, "DrawRectangleGradientV", &cint, &cint, &cint, &cint, &tColor, &tColor)     // posX, posY, width, height, top, bottom :: Draw a vertical-gradient-filled rectangle
	drawRectangleGradientH      = internal.Bind(&void, "DrawRectangleGradientH", &cint, &cint, &cint, &cint, &tColor, &tColor)     // posX, posY, width, height, left, right :: Draw a horizontal-gradient-filled rectangle
	drawRectangleGradientEx     = internal.Bind(&void, "DrawRectangleGradientEx", &tRectangle, &tColor, &tColor, &tColor, &tColor) // rec, topLeft, bottomLeft, topRight, bottomRight :: Draw a gradient-filled rectangle with custom vertex colors
	drawRectangleLines          = internal.Bind(&void, "DrawRectangleLines", &cint, &cint, &cint, &cint, &tColor)                  // posX, posY, width, height, color :: Draw rectangle outline
	drawRectangleLinesEx        = internal.Bind(&void, "DrawRectangleLinesEx", &tRectangle, &float, &tColor)                       // rec, lineThick, color :: Draw rectangle outline with extended parameters
	drawRectangleRounded        = internal.Bind(&void, "DrawRectangleRounded", &tRectangle, &float, &cint, &tColor)                // rec, roundness, segments, color :: Draw rectangle with rounded edges
	drawRectangleRoundedLines   = internal.Bind(&void, "DrawRectangleRoundedLines", &tRectangle, &float, &cint, &tColor)           // rec, roundness, segments, color :: Draw rectangle lines with rounded edges
	drawRectangleRoundedLinesEx = internal.Bind(&void, "DrawRectangleRoundedLinesEx", &tRectangle, &float, &cint, &float, &tColor) // rec, roundness, segments, lineThick, color :: Draw rectangle with rounded edges outline
	drawTriangle                = internal.Bind(&void, "DrawTriangle", &tVector2, &tVector2, &tVector2, &tColor)                   // v1, v2, v3, color :: Draw a color-filled triangle (vertex in counter-clockwise order!)
	drawTriangleLines           = internal.Bind(&void, "DrawTriangleLines", &tVector2, &tVector2, &tVector2, &tColor)              // v1, v2, v3, color :: Draw triangle outline (vertex in counter-clockwise order!)
	drawTriangleFan             = internal.Bind(&void, "DrawTriangleFan", &tVector2Ptr, &cint, &tColor)                            // points, pointCount, color :: Draw a triangle fan defined by points (first vertex is the center)
	drawTriangleStrip           = internal.Bind(&void, "DrawTriangleStrip", &tVector2Ptr, &cint, &tColor)                          // points, pointCount, color :: Draw a triangle strip defined by points
	drawPoly                    = internal.Bind(&void, "DrawPoly", &tVector2, &cint, &float, &float, &tColor)                      // center, sides, radius, rotation, color :: Draw a regular polygon (Vector version)
	drawPolyLines               = internal.Bind(&void, "DrawPolyLines", &tVector2, &cint, &float, &float, &tColor)                 // center, sides, radius, rotation, color :: Draw a polygon outline of n sides
	drawPolyLinesEx             = internal.Bind(&void, "DrawPolyLinesEx", &tVector2, &cint, &float, &float, &float, &tColor)       // center, sides, radius, rotation, lineThick, color :: Draw a polygon outline of n sides with extended parameters

	// Splines drawing functions
	drawSplineLinear                 = internal.Bind(&void, "DrawSplineLinear", &tVector2Ptr, &cint, &float, &tColor)                                    // points, pointCount, thick, color :: Draw spline: Linear, minimum 2 points
	drawSplineBasis                  = internal.Bind(&void, "DrawSplineBasis", &tVector2Ptr, &cint, &float, &tColor)                                     // points, pointCount, thick, color :: Draw spline: B-Spline, minimum 4 points
	drawSplineCatmullRom             = internal.Bind(&void, "DrawSplineCatmullRom", &tVector2Ptr, &cint, &float, &tColor)                                // points, pointCount, thick, color :: Draw spline: Catmull-Rom, minimum 4 points
	drawSplineBezierQuadratic        = internal.Bind(&void, "DrawSplineBezierQuadratic", &tVector2Ptr, &cint, &float, &tColor)                           // points, pointCount, thick, color :: Draw spline: Quadratic Bezier, minimum 3 points (1 control point): [p1, c2, p3, c4...]
	drawSplineBezierCubic            = internal.Bind(&void, "DrawSplineBezierCubic", &tVector2Ptr, &cint, &float, &tColor)                               // points, pointCount, thick, color :: Draw spline: Cubic Bezier, minimum 4 points (2 control points): [p1, c2, c3, p4, c5, c6...]
	drawSplineSegmentLinear          = internal.Bind(&void, "DrawSplineSegmentLinear", &tVector2, &tVector2, &float, &tColor)                            // p1, p2, thick, color :: Draw spline segment: Linear, 2 points
	drawSplineSegmentBasis           = internal.Bind(&void, "DrawSplineSegmentBasis", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor)       // p1, p2, p3, p4, thick, color :: Draw spline segment: B-Spline, 4 points
	drawSplineSegmentCatmullRom      = internal.Bind(&void, "DrawSplineSegmentCatmullRom", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor)  // p1, p2, p3, p4, thick, color :: Draw spline segment: Catmull-Rom, 4 points
	drawSplineSegmentBezierQuadratic = internal.Bind(&void, "DrawSplineSegmentBezierQuadratic", &tVector2, &tVector2, &tVector2, &float, &tColor)        // p1, c2, p3, thick, color :: Draw spline segment: Quadratic Bezier, 2 points, 1 control point
	drawSplineSegmentBezierCubic     = internal.Bind(&void, "DrawSplineSegmentBezierCubic", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor) // p1, c2, c3, p4, thick, color :: Draw spline segment: Cubic Bezier, 2 points, 2 control points

	// Spline segment point evaluation functions, for a given t [0.0f .. 1.0f]
	getSplinePointLinear      = internal.Bind(&tVector2, "GetSplinePointLinear", &tVector2, &tVector2, &float)                            // startPos, endPos, t :: Get (evaluate) spline point: Linear
	getSplinePointBasis       = internal.Bind(&tVector2, "GetSplinePointBasis", &tVector2, &tVector2, &tVector2, &tVector2, &float)       // p1, p2, p3, p4, t :: Get (evaluate) spline point: B-Spline
	getSplinePointCatmullRom  = internal.Bind(&tVector2, "GetSplinePointCatmullRom", &tVector2, &tVector2, &tVector2, &tVector2, &float)  // p1, p2, p3, p4, t :: Get (evaluate) spline point: Catmull-Rom
	getSplinePointBezierQuad  = internal.Bind(&tVector2, "GetSplinePointBezierQuad", &tVector2, &tVector2, &tVector2, &float)             // p1, c2, p3, t :: Get (evaluate) spline point: Quadratic Bezier
	getSplinePointBezierCubic = internal.Bind(&tVector2, "GetSplinePointBezierCubic", &tVector2, &tVector2, &tVector2, &tVector2, &float) // p1, c2, c3, p4, t :: Get (evaluate) spline point: Cubic Bezier

	// Basic shapes collision detection functions
	checkCollisionRecs          = internal.Bind(&cbool, "CheckCollisionRecs", &tRectangle, &tRectangle)                                  // rec1, rec2 :: Check collision between two rectangles
	checkCollisionCircles       = internal.Bind(&cbool, "CheckCollisionCircles", &tVector2, &float, &tVector2, &float)                   // center1, radius1, center2, radius2 :: Check collision between two circles
	checkCollisionCircleRec     = internal.Bind(&cbool, "CheckCollisionCircleRec", &tVector2, &float, &tRectangle)                       // center, radius, rec :: Check collision between circle and rectangle
	checkCollisionCircleLine    = internal.Bind(&cbool, "CheckCollisionCircleLine", &tVector2, &float, &tVector2, &tVector2)             // center, radius, p1, p2 :: Check if circle collides with a line created betweeen two points [p1] and [p2]
	checkCollisionPointRec      = internal.Bind(&cbool, "CheckCollisionPointRec", &tVector2, &tRectangle)                                // point, rec :: Check if point is inside rectangle
	checkCollisionPointCircle   = internal.Bind(&cbool, "CheckCollisionPointCircle", &tVector2, &tVector2, &float)                       // point, center, radius :: Check if point is inside circle
	checkCollisionPointTriangle = internal.Bind(&cbool, "CheckCollisionPointTriangle", &tVector2, &tVector2, &tVector2, &tVector2)       // point, p1, p2, p3 :: Check if point is inside a triangle
	checkCollisionPointLine     = internal.Bind(&cbool, "CheckCollisionPointLine", &tVector2, &tVector2, &tVector2, &cint)               // point, p1, p2, threshold :: Check if point belongs to line created between two points [p1] and [p2] with defined margin in pixels [threshold]
	checkCollisionPointPoly     = internal.Bind(&cbool, "CheckCollisionPointPoly", &tVector2, &tVector2Ptr, &cint)                       // point, points, pointCount :: Check if point is within a polygon described by array of vertices
	checkCollisionLines         = internal.Bind(&cbool, "CheckCollisionLines", &tVector2, &tVector2, &tVector2, &tVector2, &tVector2Ptr) // startPos1, endPos1, startPos2, endPos2, collisionPoint :: Check the collision between two lines defined by two points each, returns collision point by reference
	getCollisionRec             = internal.Bind(&tRectangle, "GetCollisionRec", &tRectangle, &tRectangle)                                // rec1, rec2 :: Get collision rectangle for two rectangles collision
)
