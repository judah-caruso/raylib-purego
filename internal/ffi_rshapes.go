package internal

var (
	// Set texture and rectangle to be used on shapes drawing
	// NOTE: It can be useful when using basic shapes and one single font,
	// defining a font char white rectangle would allow drawing everything in a single draw call
	SetShapesTexture          = bind(&void, "SetShapesTexture", &tTexture2D, &tRectangle) // texture, source :: Set texture and rectangle to be used on shapes drawing
	GetShapesTexture          = bind(&tTexture2D, "GetShapesTexture")                     // Get texture that is used for shapes drawing
	GetShapesTextureRectangle = bind(&tRectangle, "GetShapesTextureRectangle")            // Get texture source rectangle that is used for shapes drawing

	// Basic shapes drawing functions
	DrawPixel                   = bind(&void, "DrawPixel", &cint, &cint, &tColor)                                         // posX, posY, color :: Draw a pixel using geometry [Can be slow, use with care]
	DrawPixelV                  = bind(&void, "DrawPixelV", &tVector2, &tColor)                                           // position, color :: Draw a pixel using geometry (Vector version) [Can be slow, use with care]
	DrawLine                    = bind(&void, "DrawLine", &cint, &cint, &cint, &cint, &tColor)                            // startPosX, startPosY, endPosX, endPosY, color :: Draw a line
	DrawLineV                   = bind(&void, "DrawLineV", &tVector2, &tVector2, &tColor)                                 // startPos, endPos, color :: Draw a line (using gl lines)
	DrawLineEx                  = bind(&void, "DrawLineEx", &tVector2, &tVector2, &float, &tColor)                        // startPos, endPos, thick, color :: Draw a line (using triangles/quads)
	DrawLineStrip               = bind(&void, "DrawLineStrip", &tVector2Ptr, &cint, &tColor)                              // points, pointCount, color :: Draw lines sequence (using gl lines)
	DrawLineBezier              = bind(&void, "DrawLineBezier", &tVector2, &tVector2, &float, &tColor)                    // startPos, endPos, thick, color :: Draw line segment cubic-bezier in-out interpolation
	DrawCircle                  = bind(&void, "DrawCircle", &cint, &cint, &float, &tColor)                                // centerX, centerY, radius, color :: Draw a color-filled circle
	DrawCircleSector            = bind(&void, "DrawCircleSector", &tVector2, &float, &float, &float, &cint, &tColor)      // center, radius, startAngle, endAngle, segments, color :: Draw a piece of a circle
	DrawCircleSectorLines       = bind(&void, "DrawCircleSectorLines", &tVector2, &float, &float, &float, &cint, &tColor) // center, radius, startAngle, endAngle, segments, color :: Draw circle sector outline
	DrawCircleGradient          = bind(&void, "DrawCircleGradient", &cint, &cint, &float, &tColor, &tColor)               // centerX, centerY, radius, inner, outer :: Draw a gradient-filled circle
	DrawCircleV                 = bind(&void, "DrawCircleV", &tVector2, &float, &tColor)                                  // center, radius, color :: Draw a color-filled circle (Vector version)
	DrawCircleLines             = bind(&void, "DrawCircleLines", &cint, &cint, &float, &tColor)                           // centerX, centerY, radius, color :: Draw circle outline
	DrawCircleLinesV            = bind(&void, "DrawCircleLinesV", &tVector2, &float, &tColor)                             // center, radius, color :: Draw circle outline (Vector version)
	DrawEllipse                 = bind(&void, "DrawEllipse", &cint, &cint, &float, &float, &tColor)                       // centerX, centerY, radiusH, radiusV, color :: Draw ellipse
	DrawEllipseLines            = bind(&void, "DrawEllipseLines", &cint, &cint, &float, &float, &tColor)                  // centerX, centerY, radiusH, radiusV, color :: Draw ellipse outline
	DrawRing                    = bind(&void, "DrawRing", &tVector2, &float, &float, &float, &float, &cint, &tColor)      // center, innerRadius, outerRadius, startAngle, endAngle, segments, color :: Draw ring
	DrawRingLines               = bind(&void, "DrawRingLines", &tVector2, &float, &float, &float, &float, &cint, &tColor) // center, innerRadius, outerRadius, startAngle, endAngle, segments, color :: Draw ring outline
	DrawRectangle               = bind(&void, "DrawRectangle", &cint, &cint, &cint, &cint, &tColor)                       // posX, posY, width, height, color :: Draw a color-filled rectangle
	DrawRectangleV              = bind(&void, "DrawRectangleV", &tVector2, &tVector2, &tColor)                            // position, size, color :: Draw a color-filled rectangle (Vector version)
	DrawRectangleRec            = bind(&void, "DrawRectangleRec", &tRectangle, &tColor)                                   // rec, color :: Draw a color-filled rectangle
	DrawRectanglePro            = bind(&void, "DrawRectanglePro", &tRectangle, &tVector2, &float, &tColor)                // rec, origin, rotation, color :: Draw a color-filled rectangle with pro parameters
	DrawRectangleGradientV      = bind(&void, "DrawRectangleGradientV", &cint, &cint, &cint, &cint, &tColor, &tColor)     // posX, posY, width, height, top, bottom :: Draw a vertical-gradient-filled rectangle
	DrawRectangleGradientH      = bind(&void, "DrawRectangleGradientH", &cint, &cint, &cint, &cint, &tColor, &tColor)     // posX, posY, width, height, left, right :: Draw a horizontal-gradient-filled rectangle
	DrawRectangleGradientEx     = bind(&void, "DrawRectangleGradientEx", &tRectangle, &tColor, &tColor, &tColor, &tColor) // rec, topLeft, bottomLeft, topRight, bottomRight :: Draw a gradient-filled rectangle with custom vertex colors
	DrawRectangleLines          = bind(&void, "DrawRectangleLines", &cint, &cint, &cint, &cint, &tColor)                  // posX, posY, width, height, color :: Draw rectangle outline
	DrawRectangleLinesEx        = bind(&void, "DrawRectangleLinesEx", &tRectangle, &float, &tColor)                       // rec, lineThick, color :: Draw rectangle outline with extended parameters
	DrawRectangleRounded        = bind(&void, "DrawRectangleRounded", &tRectangle, &float, &cint, &tColor)                // rec, roundness, segments, color :: Draw rectangle with rounded edges
	DrawRectangleRoundedLines   = bind(&void, "DrawRectangleRoundedLines", &tRectangle, &float, &cint, &tColor)           // rec, roundness, segments, color :: Draw rectangle lines with rounded edges
	DrawRectangleRoundedLinesEx = bind(&void, "DrawRectangleRoundedLinesEx", &tRectangle, &float, &cint, &float, &tColor) // rec, roundness, segments, lineThick, color :: Draw rectangle with rounded edges outline
	DrawTriangle                = bind(&void, "DrawTriangle", &tVector2, &tVector2, &tVector2, &tColor)                   // v1, v2, v3, color :: Draw a color-filled triangle (vertex in counter-clockwise order!)
	DrawTriangleLines           = bind(&void, "DrawTriangleLines", &tVector2, &tVector2, &tVector2, &tColor)              // v1, v2, v3, color :: Draw triangle outline (vertex in counter-clockwise order!)
	DrawTriangleFan             = bind(&void, "DrawTriangleFan", &tVector2Ptr, &cint, &tColor)                            // points, pointCount, color :: Draw a triangle fan defined by points (first vertex is the center)
	DrawTriangleStrip           = bind(&void, "DrawTriangleStrip", &tVector2Ptr, &cint, &tColor)                          // points, pointCount, color :: Draw a triangle strip defined by points
	DrawPoly                    = bind(&void, "DrawPoly", &tVector2, &cint, &float, &float, &tColor)                      // center, sides, radius, rotation, color :: Draw a regular polygon (Vector version)
	DrawPolyLines               = bind(&void, "DrawPolyLines", &tVector2, &cint, &float, &float, &tColor)                 // center, sides, radius, rotation, color :: Draw a polygon outline of n sides
	DrawPolyLinesEx             = bind(&void, "DrawPolyLinesEx", &tVector2, &cint, &float, &float, &float, &tColor)       // center, sides, radius, rotation, lineThick, color :: Draw a polygon outline of n sides with extended parameters

	// Splines drawing functions
	DrawSplineLinear                 = bind(&void, "DrawSplineLinear", &tVector2Ptr, &cint, &float, &tColor)                                    // points, pointCount, thick, color :: Draw spline: Linear, minimum 2 points
	DrawSplineBasis                  = bind(&void, "DrawSplineBasis", &tVector2Ptr, &cint, &float, &tColor)                                     // points, pointCount, thick, color :: Draw spline: B-Spline, minimum 4 points
	DrawSplineCatmullRom             = bind(&void, "DrawSplineCatmullRom", &tVector2Ptr, &cint, &float, &tColor)                                // points, pointCount, thick, color :: Draw spline: Catmull-Rom, minimum 4 points
	DrawSplineBezierQuadratic        = bind(&void, "DrawSplineBezierQuadratic", &tVector2Ptr, &cint, &float, &tColor)                           // points, pointCount, thick, color :: Draw spline: Quadratic Bezier, minimum 3 points (1 control point): [p1, c2, p3, c4...]
	DrawSplineBezierCubic            = bind(&void, "DrawSplineBezierCubic", &tVector2Ptr, &cint, &float, &tColor)                               // points, pointCount, thick, color :: Draw spline: Cubic Bezier, minimum 4 points (2 control points): [p1, c2, c3, p4, c5, c6...]
	DrawSplineSegmentLinear          = bind(&void, "DrawSplineSegmentLinear", &tVector2, &tVector2, &float, &tColor)                            // p1, p2, thick, color :: Draw spline segment: Linear, 2 points
	DrawSplineSegmentBasis           = bind(&void, "DrawSplineSegmentBasis", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor)       // p1, p2, p3, p4, thick, color :: Draw spline segment: B-Spline, 4 points
	DrawSplineSegmentCatmullRom      = bind(&void, "DrawSplineSegmentCatmullRom", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor)  // p1, p2, p3, p4, thick, color :: Draw spline segment: Catmull-Rom, 4 points
	DrawSplineSegmentBezierQuadratic = bind(&void, "DrawSplineSegmentBezierQuadratic", &tVector2, &tVector2, &tVector2, &float, &tColor)        // p1, c2, p3, thick, color :: Draw spline segment: Quadratic Bezier, 2 points, 1 control point
	DrawSplineSegmentBezierCubic     = bind(&void, "DrawSplineSegmentBezierCubic", &tVector2, &tVector2, &tVector2, &tVector2, &float, &tColor) // p1, c2, c3, p4, thick, color :: Draw spline segment: Cubic Bezier, 2 points, 2 control points

	// Spline segment point evaluation functions, for a given t [0.0f .. 1.0f]
	GetSplinePointLinear      = bind(&tVector2, "GetSplinePointLinear", &tVector2, &tVector2, &float)                            // startPos, endPos, t :: Get (evaluate) spline point: Linear
	GetSplinePointBasis       = bind(&tVector2, "GetSplinePointBasis", &tVector2, &tVector2, &tVector2, &tVector2, &float)       // p1, p2, p3, p4, t :: Get (evaluate) spline point: B-Spline
	GetSplinePointCatmullRom  = bind(&tVector2, "GetSplinePointCatmullRom", &tVector2, &tVector2, &tVector2, &tVector2, &float)  // p1, p2, p3, p4, t :: Get (evaluate) spline point: Catmull-Rom
	GetSplinePointBezierQuad  = bind(&tVector2, "GetSplinePointBezierQuad", &tVector2, &tVector2, &tVector2, &float)             // p1, c2, p3, t :: Get (evaluate) spline point: Quadratic Bezier
	GetSplinePointBezierCubic = bind(&tVector2, "GetSplinePointBezierCubic", &tVector2, &tVector2, &tVector2, &tVector2, &float) // p1, c2, c3, p4, t :: Get (evaluate) spline point: Cubic Bezier

	// Basic shapes collision detection functions
	CheckCollisionRecs          = bind(&cbool, "CheckCollisionRecs", &tRectangle, &tRectangle)                                  // rec1, rec2 :: Check collision between two rectangles
	CheckCollisionCircles       = bind(&cbool, "CheckCollisionCircles", &tVector2, &float, &tVector2, &float)                   // center1, radius1, center2, radius2 :: Check collision between two circles
	CheckCollisionCircleRec     = bind(&cbool, "CheckCollisionCircleRec", &tVector2, &float, &tRectangle)                       // center, radius, rec :: Check collision between circle and rectangle
	CheckCollisionCircleLine    = bind(&cbool, "CheckCollisionCircleLine", &tVector2, &float, &tVector2, &tVector2)             // center, radius, p1, p2 :: Check if circle collides with a line created betweeen two points [p1] and [p2]
	CheckCollisionPointRec      = bind(&cbool, "CheckCollisionPointRec", &tVector2, &tRectangle)                                // point, rec :: Check if point is inside rectangle
	CheckCollisionPointCircle   = bind(&cbool, "CheckCollisionPointCircle", &tVector2, &tVector2, &float)                       // point, center, radius :: Check if point is inside circle
	CheckCollisionPointTriangle = bind(&cbool, "CheckCollisionPointTriangle", &tVector2, &tVector2, &tVector2, &tVector2)       // point, p1, p2, p3 :: Check if point is inside a triangle
	CheckCollisionPointLine     = bind(&cbool, "CheckCollisionPointLine", &tVector2, &tVector2, &tVector2, &cint)               // point, p1, p2, threshold :: Check if point belongs to line created between two points [p1] and [p2] with defined margin in pixels [threshold]
	CheckCollisionPointPoly     = bind(&cbool, "CheckCollisionPointPoly", &tVector2, &tVector2Ptr, &cint)                       // point, points, pointCount :: Check if point is within a polygon described by array of vertices
	CheckCollisionLines         = bind(&cbool, "CheckCollisionLines", &tVector2, &tVector2, &tVector2, &tVector2, &tVector2Ptr) // startPos1, endPos1, startPos2, endPos2, collisionPoint :: Check the collision between two lines defined by two points each, returns collision point by reference
	GetCollisionRec             = bind(&tRectangle, "GetCollisionRec", &tRectangle, &tRectangle)                                // rec1, rec2 :: Get collision rectangle for two rectangles collision
)
