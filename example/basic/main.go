package main

import (
	"fmt"
	"image"
	"image/color"
	"unsafe"

	"github.com/icexin/nk/cnk"
	"github.com/icexin/nk/libc"

	"github.com/fogleman/fauxgl"
)

type vertex struct {
	pos   [2]float32
	uv    [2]float32
	color [4]byte
}

func main() {
	tls := libc.NewTLS()

	fctx := fauxgl.NewContext(600, 480)
	fctx.Cull = fauxgl.CullNone
	fctx.ReadDepth = false
	fctx.WriteDepth = false

	null := cnk.Nk_draw_null_texture{}

	var mem [64 * 1024]byte
	var title = [...]byte{'h', 'e', 'l', 'l', 'o', 0}
	var ctx = uintptr(unsafe.Pointer(new(cnk.Nk_context)))
	cnk.Xnk_init_fixed(tls, ctx, uintptr(unsafe.Pointer(&mem[0])), 64*1024, 0)

	atlasptr := new(cnk.Nk_font_atlas)
	var atlas = uintptr(unsafe.Pointer(atlasptr))
	cnk.Xnk_font_atlas_init_default(tls, atlas)
	wptr, hptr := 0, 0
	w, h := uintptr(unsafe.Pointer(&wptr)), uintptr(unsafe.Pointer(&hptr))
	piximg := cnk.Xnk_font_atlas_bake(tls, atlas, w, h, cnk.NK_FONT_ATLAS_RGBA32)
	pix := make([]byte, wptr*hptr*4)
	copy(pix, (*[1 << 20]byte)(unsafe.Pointer(piximg))[:wptr*hptr*4])
	cnk.Xnk_font_atlas_end(tls, atlas, cnk.Xnk_handle_id(tls, 1), uintptr(unsafe.Pointer(&null)))

	img := &image.RGBA{
		Pix:    pix,
		Rect:   image.Rect(0, 0, wptr, hptr),
		Stride: wptr * 4,
	}
	tex := fauxgl.NewImageTexture(img)
	// m := fauxgl.Scale(fauxgl.V(1, -1, 0)).Translate(fauxgl.V(0, 480, 0)).Orthographic(0, 600, 0, 480, -1, 1)
	// m := fauxgl.Orthographic(0, 600, 0, 480, -1, 1)
	m := fauxgl.Orthographic(0, 600, 480, 0, -1, 1)
	shader := &glshader{
		Matrix: m,
	}
	fctx.Shader = shader

	if atlasptr.Default_font != 0 {
		default_font := (*cnk.Nk_font)(unsafe.Pointer(atlasptr.Default_font))
		cnk.Xnk_style_set_font(tls, ctx, uintptr(unsafe.Pointer(&default_font.Handle)))
	}

	cnk.Xnk_begin(tls, ctx, uintptr(unsafe.Pointer(&title[0])), cnk.Xnk_rect(tls, 50, 50, 220, 220), cnk.NK_WINDOW_BORDER|cnk.NK_WINDOW_MOVABLE|cnk.NK_WINDOW_CLOSABLE|cnk.NK_WINDOW_SCALABLE)
	cnk.Xnk_layout_row_static(tls, ctx, 30, 80, 1)
	cnk.Xnk_button_label(tls, ctx, uintptr(unsafe.Pointer(&title[0])))
	cnk.Xnk_end(tls, ctx)

	var vertex_layout = [...]cnk.Nk_draw_vertex_layout_element{
		{cnk.NK_VERTEX_POSITION, cnk.NK_FORMAT_FLOAT, cnk.Nk_size(unsafe.Offsetof(vertex{}.pos))},
		{cnk.NK_VERTEX_TEXCOORD, cnk.NK_FORMAT_FLOAT, cnk.Nk_size(unsafe.Offsetof(vertex{}.uv))},
		{cnk.NK_VERTEX_COLOR, cnk.NK_FORMAT_R8G8B8A8, cnk.Nk_size(unsafe.Offsetof(vertex{}.color))},
		{cnk.NK_VERTEX_ATTRIBUTE_COUNT, cnk.NK_FORMAT_COUNT, 0},
	}
	var cfg cnk.Nk_convert_config
	cfg.Shape_AA = cnk.NK_ANTI_ALIASING_ON
	cfg.Line_AA = cnk.NK_ANTI_ALIASING_ON
	cfg.Vertex_layout = uintptr(unsafe.Pointer(&vertex_layout[0]))
	cfg.Vertex_size = cnk.Nk_size(unsafe.Sizeof(vertex{}))
	cfg.Vertex_alignment = cnk.Nk_size(unsafe.Alignof(vertex{}))
	cfg.Circle_segment_count = 22
	cfg.Curve_segment_count = 22
	cfg.Arc_segment_count = 22
	cfg.Global_alpha = 1.0
	cfg.Null = null

	cmdbuf := uintptr(unsafe.Pointer(new(cnk.Nk_buffer)))
	vertbuf := uintptr(unsafe.Pointer(new(cnk.Nk_buffer)))
	idxbuf := uintptr(unsafe.Pointer(new(cnk.Nk_buffer)))

	cnk.Xnk_buffer_init_default(tls, cmdbuf)
	cnk.Xnk_buffer_init_default(tls, vertbuf)
	cnk.Xnk_buffer_init_default(tls, idxbuf)

	cnk.Xnk_convert(tls, ctx, cmdbuf, vertbuf, idxbuf, uintptr(unsafe.Pointer(&cfg)))
	vertptr := cnk.Xnk_buffer_memory(tls, vertbuf)
	vertsize := cnk.Xnk_buffer_total(tls, vertbuf)
	verts := ((*[1 << 20]vertex)(unsafe.Pointer(vertptr)))[:vertsize]

	idxptr := cnk.Xnk_buffer_memory(tls, idxbuf)
	idxsize := cnk.Xnk_buffer_total(tls, idxbuf)
	idxs := ((*[1 << 20]uint16)(unsafe.Pointer(idxptr)))[:idxsize]

	cmd := cnk.Xnk__draw_begin(tls, ctx, cmdbuf)
	for ; cmd != 0; cmd = cnk.Xnk__draw_next(tls, cmd, cmdbuf, ctx) {
		cmdp := (*cnk.Nk_draw_command)(unsafe.Pointer(cmd))
		if cmdp.Elem_count == 0 {
			continue
		}
		fmt.Printf("verts:%d, tex:%d\n", cmdp.Elem_count, cmdp.Texture.Ptr)
		if cmdp.Texture.Ptr != 0 {
			shader.Texture = tex
		} else {
			shader.Texture = nil
		}

		idx := idxs[:cmdp.Elem_count]
		vs := make([]fauxgl.Vertex, 0, cmdp.Elem_count)
		for _, i := range idx {
			v := verts[i]
			vs = append(vs, fauxgl.Vertex{
				Position: fauxgl.V(float64(v.pos[0]), float64(v.pos[1]), 0),
				Texture:  fauxgl.V(float64(v.uv[0]), float64(v.uv[1]), 0),
				Color:    fauxgl.MakeColor(color.RGBA{v.color[0], v.color[1], v.color[2], v.color[3]}),
			})
		}

		for i := 0; i < len(vs); i += 3 {
			t := &fauxgl.Triangle{
				V1: vs[i],
				V2: vs[i+1],
				V3: vs[i+2],
			}
			// t := (*fauxgl.Triangle)(unsafe.Pointer(&vs[i]))
			fctx.DrawTriangle(t)
		}
		idxs = idxs[cmdp.Elem_count:]
	}

	cnk.Xnk_buffer_free(tls, cmdbuf)
	cnk.Xnk_buffer_free(tls, vertbuf)
	cnk.Xnk_buffer_free(tls, idxbuf)

	fauxgl.SavePNG("out.png", fctx.Image())
}

type glshader struct {
	Matrix  fauxgl.Matrix
	Texture fauxgl.Texture
}

func (s *glshader) Vertex(v fauxgl.Vertex) fauxgl.Vertex {
	v.Output = s.Matrix.MulPositionW(v.Position)
	return v
}

func (s *glshader) Fragment(v fauxgl.Vertex) fauxgl.Color {
	if s.Texture != nil {
		return s.Texture.Sample(v.Texture.X, 1-v.Texture.Y).Mul(v.Color)
	}
	return v.Color
}
