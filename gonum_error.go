
package gonum

import (
    "errors"
)

var  ErrorInvalidParameter          = errors.New("Error Invalid Parameter");
var  ErrorMatrixIsEmpty             = errors.New("Error Matrix Is Empty");
var  ErrorDimUnmatched              = errors.New("Error Dimension Unmatch");
var  ErrorIndexConflict             = errors.New("Error Index Conflict");
var  ErrorElementTypeUnmatched      = errors.New("Error Element Type Unmatched");
var  ErrorElementTypeNotSet         = errors.New("Error Element Not Set");
var  ErrorElementUnknown            = errors.New("Error Element Type Unknown");
var  ErrorNotImplemented            = errors.New("Error Method Not Implemented");