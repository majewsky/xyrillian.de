// SPDX-FileCopyrightText: 2025 Stefan Majewsky <majewsky@gmx.net>
// SPDX-License-Identifier: GPL-3.0-only

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func MustSucceed(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MustReturn[T any](value T, err error) T {
	MustSucceed(err)
	return value
}

func UnmarshalYAMLviaJSON(buf []byte, target any) (err error) {
	// I hate dealing with YAML libraries in Go because the ecosystem is a mess,
	// so we're skipping that entire debate using <https://github.com/mikefarah/yq>
	cmd := exec.Command("yq", "-o", "json")
	cmd.Stdin = bytes.NewReader(buf)
	cmd.Stderr = os.Stderr
	buf, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("could not run `yq -o json`: " + err.Error())
	}
	return json.Unmarshal(buf, target)
}

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.Table,
		extension.Strikethrough,
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
)

func RenderMarkdown(in []byte) (string, error) {
	var buf bytes.Buffer
	err := md.Convert(in, &buf)
	return buf.String(), err
}
