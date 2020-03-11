// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package datatypes

import (
	"base/binary"
	"base/errors"
	"datavalues"
	"fmt"
	"io"
)

const (
	DataTypeFloat32Name = "Float32"
)

type Float32DataType struct {
	DataTypeBase
}

func NewFloat32DataType() IDataType {
	return &Float32DataType{}
}

func (datatype *Float32DataType) Name() string {
	return DataTypeFloat32Name
}

func (datatype *Float32DataType) Serialize(writer *binary.Writer, v datavalues.IDataValue) error {
	if err := writer.Float32(float32(datavalues.AsFloat(v))); err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (datatype *Float32DataType) SerializeText(writer io.Writer, v datavalues.IDataValue) error {
	if _, err := writer.Write([]byte(fmt.Sprintf("%v", datavalues.AsFloat(v)))); err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (datatype *Float32DataType) Deserialize(reader *binary.Reader) (datavalues.IDataValue, error) {
	if res, err := reader.Float32(); err != nil {
		return nil, errors.Wrap(err)
	} else {
		return datavalues.MakeFloat(float64(res)), nil
	}
}
