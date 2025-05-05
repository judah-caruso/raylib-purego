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

	// Splines drawing functions
	drawSplineLinear                 = prep(&void, "DrawSplineLinear", &tVector2Ptr, &cint, &float, &tColor)                                    // points, pointCount, thick, color :: Draw spline: Linear, minimum 2 points
	drawSplineBasis                  = prep(&void, "DrawSplineBasis", &tVector2Ptr, &cint, &float, &tColor)                                     // points, pointCount, thick, color :: Draw spline: B-Spline, minimum 4 points
	drawSplineCatmullRom             = prep(&void, "DrawSplineCatmullRom", &tVector2Ptr, &cint, &float, &tColor)                                // points, pointCount, thick, color :: Draw spline: Catmull-Rom, minimum 4 points
	drawSplineBezierQuadratic        = prep(&void, "DrawSplineBezierQuadratic", &tVector2Ptr, &cint, &float, &tColor)                           // points, pointCount, thick, color :: Draw spline: Quadratic Bezier, minimum 3 points (1 control point): [p1, c2, p3, c4...]
	drawSplineBezierCubic            = prep(&void, "DrawSplineBezierCubic", &tVector2Ptr, &cint, &float, &tColor)                               // points, pointCount, thick, color :: Draw spline: Cubic Bezier, minimum 4 points (2 control points): [p1, c2, c3, p4, c5, c6...]
	drawSplineSegmentLinear          = prep(&void, "DrawSplineSegmentLinear", &tVector2, &tVector2, &float, &tColor)                            // p1, p2, thick, color :: Draw spline segment: Linear, 2 points
	drawSplineSegmentBasis           = prep(&void, "DrawSplineSegmentBasis", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor)       // p1, p2, p3, p4, thick, color :: Draw spline segment: B-Spline, 4 points
	drawSplineSegmentCatmullRom      = prep(&void, "DrawSplineSegmentCatmullRom", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor)  // p1, p2, p3, p4, thick, color :: Draw spline segment: Catmull-Rom, 4 points
	drawSplineSegmentBezierQuadratic = prep(&void, "DrawSplineSegmentBezierQuadratic", &tVector2, &tVector2, &tVector2, &float, &tColor)        // p1, c2, p3, thick, color :: Draw spline segment: Quadratic Bezier, 2 points, 1 control point
	drawSplineSegmentBezierCubic     = prep(&void, "DrawSplineSegmentBezierCubic", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor) // p1, c2, c3, p4, thick, color :: Draw spline segment: Cubic Bezier, 2 points, 2 control points

	// Spline segment point evaluation functions, for a given t [0.0f .. 1.0f]
	getSplinePointLinear      = prep(&tVector2, "GetSplinePointLinear", &tVector2, &tVector2, &float)                            // startPos, endPos, t :: Get (evaluate) spline point: Linear
	getSplinePointBasis       = prep(&tVector2, "GetSplinePointBasis", &tVector2, &tVector2, &tVector2, &tVector2, &float)       // p1, p2, p3, p4, t :: Get (evaluate) spline point: B-Spline
	getSplinePointCatmullRom  = prep(&tVector2, "GetSplinePointCatmullRom", &tVector2, &tVector2, &tVector2, &tVector2, &float)  // p1, p2, p3, p4, t :: Get (evaluate) spline point: Catmull-Rom
	getSplinePointBezierQuad  = prep(&tVector2, "GetSplinePointBezierQuad", &tVector2, &tVector2, &tVector2, &float)             // p1, c2, p3, t :: Get (evaluate) spline point: Quadratic Bezier
	getSplinePointBezierCubic = prep(&tVector2, "GetSplinePointBezierCubic", &tVector2, &tVector2, &tVector2, &tVector2, &float) // p1, c2, c3, p4, t :: Get (evaluate) spline point: Cubic Bezier

	// Basic shapes collision detection functions
	checkCollisionRecs          = prep(&cbool, "CheckCollisionRecs", &tRectangle, &tRectangle)                                  // rec1, rec2 :: Check collision between two rectangles
	checkCollisionCircles       = prep(&cbool, "CheckCollisionCircles", &tVector2, &float, &tVector2, &float)                   // center1, radius1, center2, radius2 :: Check collision between two circles
	checkCollisionCircleRec     = prep(&cbool, "CheckCollisionCircleRec", &tVector2, &float, &tRectangle)                       // center, radius, rec :: Check collision between circle and rectangle
	checkCollisionCircleLine    = prep(&cbool, "CheckCollisionCircleLine", &tVector2, &float, &tVector2, &tVector2)             // center, radius, p1, p2 :: Check if circle collides with a line created betweeen two points [p1] and [p2]
	checkCollisionPointRec      = prep(&cbool, "CheckCollisionPointRec", &tVector2, &tRectangle)                                // point, rec :: Check if point is inside rectangle
	checkCollisionPointCircle   = prep(&cbool, "CheckCollisionPointCircle", &tVector2, &tVector2, &float)                       // point, center, radius :: Check if point is inside circle
	checkCollisionPointTriangle = prep(&cbool, "CheckCollisionPointTriangle", &tVector2, &tVector2, &tVector2, &tVector2)       // point, p1, p2, p3 :: Check if point is inside a triangle
	checkCollisionPointLine     = prep(&cbool, "CheckCollisionPointLine", &tVector2, &tVector2, &tVector2, &cint)               // point, p1, p2, threshold :: Check if point belongs to line created between two points [p1] and [p2] with defined margin in pixels [threshold]
	checkCollisionPointPoly     = prep(&cbool, "CheckCollisionPointPoly", &tVector2, &tVector2Ptr, &cint)                       // point, points, pointCount :: Check if point is within a polygon described by array of vertices
	checkCollisionLines         = prep(&cbool, "CheckCollisionLines", &tVector2, &tVector2, &tVector2, &tVector2, &tVector2Ptr) // startPos1, endPos1, startPos2, endPos2, collisionPoint :: Check the collision between two lines defined by two points each, returns collision point by reference
	getCollisionRec             = prep(&tRectangle, "GetCollisionRec", &tRectangle, &tRectangle)                                // rec1, rec2 :: Get collision rectangle for two rectangles collision

)
