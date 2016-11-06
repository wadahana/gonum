package gonum

import (
//    "fmt"
//    "reflect"
)

type Vector struct {
    matrix * Matrix;
}


func NewVector(length int, t ElementType) *Vector {
    v := Vector{};
    v.matrix = NewMatrix(length, 1, t);
    return &v;
}

func NewEmptyVector() *Vector {
    return NewVector(0, ElementUnknown);
}

func NewVectorWithData(length int, values interface{}) (*Vector, error) {
    var err error = nil;
    v := Vector{};
    v.matrix, err = NewMatrixWithData(length, 1, values);
    if err != nil {
        return nil, err;
    }
    return &v, nil;
}

func (this *Vector) String() string {
    return this.matrix.String();
}

func (this *Vector) GetElementNum() int {
    return this.matrix.GetElementNum();
}

func (this *Vector) GetElementType() ElementType {
    return this.matrix.element_type;
}

func (this *Vector) Get(index int) interface{} {
    return this.matrix.Get(index, 1);
}

func (this *Vector) Set(index int, value interface{}) error {
    return this.matrix.Set(index, 1, value);
}

// func (this *Vector) Transpose() *Vector {
//     m := this.matrix.Transpose();
// }