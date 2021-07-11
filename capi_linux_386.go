// Code generated by 'ccgo -o nk_linux_386.go nuklear/nuklear.c -export-externs X -export-fields  -export-defines  -export-enums  -export-typedefs  -export-structs  -nostdlib -pkgname nk', DO NOT EDIT.

package nk

var CAPI = map[string]struct{}{
	"nk__begin":                                  {},
	"nk__draw_begin":                             {},
	"nk__draw_end":                               {},
	"nk__draw_list_begin":                        {},
	"nk__draw_list_end":                          {},
	"nk__draw_list_next":                         {},
	"nk__draw_next":                              {},
	"nk__next":                                   {},
	"nk_begin":                                   {},
	"nk_begin_titled":                            {},
	"nk_buffer_clear":                            {},
	"nk_buffer_free":                             {},
	"nk_buffer_info":                             {},
	"nk_buffer_init":                             {},
	"nk_buffer_init_default":                     {},
	"nk_buffer_init_fixed":                       {},
	"nk_buffer_mark":                             {},
	"nk_buffer_memory":                           {},
	"nk_buffer_memory_const":                     {},
	"nk_buffer_push":                             {},
	"nk_buffer_reset":                            {},
	"nk_buffer_total":                            {},
	"nk_button_color":                            {},
	"nk_button_image":                            {},
	"nk_button_image_label":                      {},
	"nk_button_image_label_styled":               {},
	"nk_button_image_styled":                     {},
	"nk_button_image_text":                       {},
	"nk_button_image_text_styled":                {},
	"nk_button_label":                            {},
	"nk_button_label_styled":                     {},
	"nk_button_pop_behavior":                     {},
	"nk_button_push_behavior":                    {},
	"nk_button_set_behavior":                     {},
	"nk_button_symbol":                           {},
	"nk_button_symbol_label":                     {},
	"nk_button_symbol_label_styled":              {},
	"nk_button_symbol_styled":                    {},
	"nk_button_symbol_text":                      {},
	"nk_button_symbol_text_styled":               {},
	"nk_button_text":                             {},
	"nk_button_text_styled":                      {},
	"nk_chart_add_slot":                          {},
	"nk_chart_add_slot_colored":                  {},
	"nk_chart_begin":                             {},
	"nk_chart_begin_colored":                     {},
	"nk_chart_end":                               {},
	"nk_chart_push":                              {},
	"nk_chart_push_slot":                         {},
	"nk_check_flags_label":                       {},
	"nk_check_flags_text":                        {},
	"nk_check_label":                             {},
	"nk_check_text":                              {},
	"nk_checkbox_flags_label":                    {},
	"nk_checkbox_flags_text":                     {},
	"nk_checkbox_label":                          {},
	"nk_checkbox_text":                           {},
	"nk_clear":                                   {},
	"nk_color_cf":                                {},
	"nk_color_d":                                 {},
	"nk_color_dv":                                {},
	"nk_color_f":                                 {},
	"nk_color_fv":                                {},
	"nk_color_hex_rgb":                           {},
	"nk_color_hex_rgba":                          {},
	"nk_color_hsv_b":                             {},
	"nk_color_hsv_bv":                            {},
	"nk_color_hsv_f":                             {},
	"nk_color_hsv_fv":                            {},
	"nk_color_hsv_i":                             {},
	"nk_color_hsv_iv":                            {},
	"nk_color_hsva_b":                            {},
	"nk_color_hsva_bv":                           {},
	"nk_color_hsva_f":                            {},
	"nk_color_hsva_fv":                           {},
	"nk_color_hsva_i":                            {},
	"nk_color_hsva_iv":                           {},
	"nk_color_pick":                              {},
	"nk_color_picker":                            {},
	"nk_color_u32":                               {},
	"nk_colorf_hsva_f":                           {},
	"nk_colorf_hsva_fv":                          {},
	"nk_combo":                                   {},
	"nk_combo_begin_color":                       {},
	"nk_combo_begin_image":                       {},
	"nk_combo_begin_image_label":                 {},
	"nk_combo_begin_image_text":                  {},
	"nk_combo_begin_label":                       {},
	"nk_combo_begin_symbol":                      {},
	"nk_combo_begin_symbol_label":                {},
	"nk_combo_begin_symbol_text":                 {},
	"nk_combo_begin_text":                        {},
	"nk_combo_callback":                          {},
	"nk_combo_close":                             {},
	"nk_combo_end":                               {},
	"nk_combo_item_image_label":                  {},
	"nk_combo_item_image_text":                   {},
	"nk_combo_item_label":                        {},
	"nk_combo_item_symbol_label":                 {},
	"nk_combo_item_symbol_text":                  {},
	"nk_combo_item_text":                         {},
	"nk_combo_separator":                         {},
	"nk_combo_string":                            {},
	"nk_combobox":                                {},
	"nk_combobox_callback":                       {},
	"nk_combobox_separator":                      {},
	"nk_combobox_string":                         {},
	"nk_contextual_begin":                        {},
	"nk_contextual_close":                        {},
	"nk_contextual_end":                          {},
	"nk_contextual_item_image_label":             {},
	"nk_contextual_item_image_text":              {},
	"nk_contextual_item_label":                   {},
	"nk_contextual_item_symbol_label":            {},
	"nk_contextual_item_symbol_text":             {},
	"nk_contextual_item_text":                    {},
	"nk_convert":                                 {},
	"nk_draw_image":                              {},
	"nk_draw_list_add_image":                     {},
	"nk_draw_list_add_text":                      {},
	"nk_draw_list_fill_circle":                   {},
	"nk_draw_list_fill_poly_convex":              {},
	"nk_draw_list_fill_rect":                     {},
	"nk_draw_list_fill_rect_multi_color":         {},
	"nk_draw_list_fill_triangle":                 {},
	"nk_draw_list_init":                          {},
	"nk_draw_list_path_arc_to":                   {},
	"nk_draw_list_path_arc_to_fast":              {},
	"nk_draw_list_path_clear":                    {},
	"nk_draw_list_path_curve_to":                 {},
	"nk_draw_list_path_fill":                     {},
	"nk_draw_list_path_line_to":                  {},
	"nk_draw_list_path_rect_to":                  {},
	"nk_draw_list_path_stroke":                   {},
	"nk_draw_list_setup":                         {},
	"nk_draw_list_stroke_circle":                 {},
	"nk_draw_list_stroke_curve":                  {},
	"nk_draw_list_stroke_line":                   {},
	"nk_draw_list_stroke_poly_line":              {},
	"nk_draw_list_stroke_rect":                   {},
	"nk_draw_list_stroke_triangle":               {},
	"nk_draw_text":                               {},
	"nk_edit_buffer":                             {},
	"nk_edit_focus":                              {},
	"nk_edit_string":                             {},
	"nk_edit_string_zero_terminated":             {},
	"nk_edit_unfocus":                            {},
	"nk_end":                                     {},
	"nk_fill_arc":                                {},
	"nk_fill_circle":                             {},
	"nk_fill_polygon":                            {},
	"nk_fill_rect":                               {},
	"nk_fill_rect_multi_color":                   {},
	"nk_fill_triangle":                           {},
	"nk_filter_ascii":                            {},
	"nk_filter_binary":                           {},
	"nk_filter_decimal":                          {},
	"nk_filter_default":                          {},
	"nk_filter_float":                            {},
	"nk_filter_hex":                              {},
	"nk_filter_oct":                              {},
	"nk_font_atlas_add":                          {},
	"nk_font_atlas_add_compressed":               {},
	"nk_font_atlas_add_compressed_base85":        {},
	"nk_font_atlas_add_default":                  {},
	"nk_font_atlas_add_from_memory":              {},
	"nk_font_atlas_bake":                         {},
	"nk_font_atlas_begin":                        {},
	"nk_font_atlas_cleanup":                      {},
	"nk_font_atlas_clear":                        {},
	"nk_font_atlas_end":                          {},
	"nk_font_atlas_init":                         {},
	"nk_font_atlas_init_custom":                  {},
	"nk_font_atlas_init_default":                 {},
	"nk_font_chinese_glyph_ranges":               {},
	"nk_font_config":                             {},
	"nk_font_cyrillic_glyph_ranges":              {},
	"nk_font_default_glyph_ranges":               {},
	"nk_font_find_glyph":                         {},
	"nk_font_korean_glyph_ranges":                {},
	"nk_free":                                    {},
	"nk_get_null_rect":                           {},
	"nk_group_begin":                             {},
	"nk_group_begin_titled":                      {},
	"nk_group_end":                               {},
	"nk_group_get_scroll":                        {},
	"nk_group_scrolled_begin":                    {},
	"nk_group_scrolled_end":                      {},
	"nk_group_scrolled_offset_begin":             {},
	"nk_group_set_scroll":                        {},
	"nk_handle_id":                               {},
	"nk_handle_ptr":                              {},
	"nk_hsv":                                     {},
	"nk_hsv_bv":                                  {},
	"nk_hsv_f":                                   {},
	"nk_hsv_fv":                                  {},
	"nk_hsv_iv":                                  {},
	"nk_hsva":                                    {},
	"nk_hsva_bv":                                 {},
	"nk_hsva_colorf":                             {},
	"nk_hsva_colorfv":                            {},
	"nk_hsva_f":                                  {},
	"nk_hsva_fv":                                 {},
	"nk_hsva_iv":                                 {},
	"nk_image":                                   {},
	"nk_image_color":                             {},
	"nk_image_handle":                            {},
	"nk_image_id":                                {},
	"nk_image_is_subimage":                       {},
	"nk_image_ptr":                               {},
	"nk_init":                                    {},
	"nk_init_custom":                             {},
	"nk_init_default":                            {},
	"nk_init_fixed":                              {},
	"nk_input_any_mouse_click_in_rect":           {},
	"nk_input_begin":                             {},
	"nk_input_button":                            {},
	"nk_input_char":                              {},
	"nk_input_end":                               {},
	"nk_input_glyph":                             {},
	"nk_input_has_mouse_click":                   {},
	"nk_input_has_mouse_click_down_in_rect":      {},
	"nk_input_has_mouse_click_in_rect":           {},
	"nk_input_is_key_down":                       {},
	"nk_input_is_key_pressed":                    {},
	"nk_input_is_key_released":                   {},
	"nk_input_is_mouse_click_down_in_rect":       {},
	"nk_input_is_mouse_click_in_rect":            {},
	"nk_input_is_mouse_down":                     {},
	"nk_input_is_mouse_hovering_rect":            {},
	"nk_input_is_mouse_pressed":                  {},
	"nk_input_is_mouse_prev_hovering_rect":       {},
	"nk_input_is_mouse_released":                 {},
	"nk_input_key":                               {},
	"nk_input_motion":                            {},
	"nk_input_mouse_clicked":                     {},
	"nk_input_scroll":                            {},
	"nk_input_unicode":                           {},
	"nk_item_is_any_active":                      {},
	"nk_label":                                   {},
	"nk_label_colored":                           {},
	"nk_label_colored_wrap":                      {},
	"nk_label_wrap":                              {},
	"nk_layout_ratio_from_pixel":                 {},
	"nk_layout_reset_min_row_height":             {},
	"nk_layout_row":                              {},
	"nk_layout_row_begin":                        {},
	"nk_layout_row_dynamic":                      {},
	"nk_layout_row_end":                          {},
	"nk_layout_row_push":                         {},
	"nk_layout_row_static":                       {},
	"nk_layout_row_template_begin":               {},
	"nk_layout_row_template_end":                 {},
	"nk_layout_row_template_push_dynamic":        {},
	"nk_layout_row_template_push_static":         {},
	"nk_layout_row_template_push_variable":       {},
	"nk_layout_set_min_row_height":               {},
	"nk_layout_space_begin":                      {},
	"nk_layout_space_bounds":                     {},
	"nk_layout_space_end":                        {},
	"nk_layout_space_push":                       {},
	"nk_layout_space_rect_to_local":              {},
	"nk_layout_space_rect_to_screen":             {},
	"nk_layout_space_to_local":                   {},
	"nk_layout_space_to_screen":                  {},
	"nk_layout_widget_bounds":                    {},
	"nk_list_view_begin":                         {},
	"nk_list_view_end":                           {},
	"nk_menu_begin_image":                        {},
	"nk_menu_begin_image_label":                  {},
	"nk_menu_begin_image_text":                   {},
	"nk_menu_begin_label":                        {},
	"nk_menu_begin_symbol":                       {},
	"nk_menu_begin_symbol_label":                 {},
	"nk_menu_begin_symbol_text":                  {},
	"nk_menu_begin_text":                         {},
	"nk_menu_close":                              {},
	"nk_menu_end":                                {},
	"nk_menu_item_image_label":                   {},
	"nk_menu_item_image_text":                    {},
	"nk_menu_item_label":                         {},
	"nk_menu_item_symbol_label":                  {},
	"nk_menu_item_symbol_text":                   {},
	"nk_menu_item_text":                          {},
	"nk_menubar_begin":                           {},
	"nk_menubar_end":                             {},
	"nk_murmur_hash":                             {},
	"nk_option_label":                            {},
	"nk_option_text":                             {},
	"nk_plot":                                    {},
	"nk_plot_function":                           {},
	"nk_popup_begin":                             {},
	"nk_popup_close":                             {},
	"nk_popup_end":                               {},
	"nk_popup_get_scroll":                        {},
	"nk_popup_set_scroll":                        {},
	"nk_prog":                                    {},
	"nk_progress":                                {},
	"nk_property_double":                         {},
	"nk_property_float":                          {},
	"nk_property_int":                            {},
	"nk_propertyd":                               {},
	"nk_propertyf":                               {},
	"nk_propertyi":                               {},
	"nk_push_custom":                             {},
	"nk_push_scissor":                            {},
	"nk_radio_label":                             {},
	"nk_radio_text":                              {},
	"nk_rect":                                    {},
	"nk_rect_pos":                                {},
	"nk_rect_size":                               {},
	"nk_recta":                                   {},
	"nk_recti":                                   {},
	"nk_rectiv":                                  {},
	"nk_rectv":                                   {},
	"nk_rgb":                                     {},
	"nk_rgb_bv":                                  {},
	"nk_rgb_cf":                                  {},
	"nk_rgb_f":                                   {},
	"nk_rgb_fv":                                  {},
	"nk_rgb_hex":                                 {},
	"nk_rgb_iv":                                  {},
	"nk_rgba":                                    {},
	"nk_rgba_bv":                                 {},
	"nk_rgba_cf":                                 {},
	"nk_rgba_f":                                  {},
	"nk_rgba_fv":                                 {},
	"nk_rgba_hex":                                {},
	"nk_rgba_iv":                                 {},
	"nk_rgba_u32":                                {},
	"nk_select_image_label":                      {},
	"nk_select_image_text":                       {},
	"nk_select_label":                            {},
	"nk_select_symbol_label":                     {},
	"nk_select_symbol_text":                      {},
	"nk_select_text":                             {},
	"nk_selectable_image_label":                  {},
	"nk_selectable_image_text":                   {},
	"nk_selectable_label":                        {},
	"nk_selectable_symbol_label":                 {},
	"nk_selectable_symbol_text":                  {},
	"nk_selectable_text":                         {},
	"nk_slide_float":                             {},
	"nk_slide_int":                               {},
	"nk_slider_float":                            {},
	"nk_slider_int":                              {},
	"nk_spacing":                                 {},
	"nk_str_append_str_char":                     {},
	"nk_str_append_str_runes":                    {},
	"nk_str_append_str_utf8":                     {},
	"nk_str_append_text_char":                    {},
	"nk_str_append_text_runes":                   {},
	"nk_str_append_text_utf8":                    {},
	"nk_str_at_char":                             {},
	"nk_str_at_char_const":                       {},
	"nk_str_at_const":                            {},
	"nk_str_at_rune":                             {},
	"nk_str_clear":                               {},
	"nk_str_delete_chars":                        {},
	"nk_str_delete_runes":                        {},
	"nk_str_free":                                {},
	"nk_str_get":                                 {},
	"nk_str_get_const":                           {},
	"nk_str_init":                                {},
	"nk_str_init_default":                        {},
	"nk_str_init_fixed":                          {},
	"nk_str_insert_at_char":                      {},
	"nk_str_insert_at_rune":                      {},
	"nk_str_insert_str_char":                     {},
	"nk_str_insert_str_runes":                    {},
	"nk_str_insert_str_utf8":                     {},
	"nk_str_insert_text_char":                    {},
	"nk_str_insert_text_runes":                   {},
	"nk_str_insert_text_utf8":                    {},
	"nk_str_len":                                 {},
	"nk_str_len_char":                            {},
	"nk_str_remove_chars":                        {},
	"nk_str_remove_runes":                        {},
	"nk_str_rune_at":                             {},
	"nk_strfilter":                               {},
	"nk_stricmp":                                 {},
	"nk_stricmpn":                                {},
	"nk_strlen":                                  {},
	"nk_strmatch_fuzzy_string":                   {},
	"nk_strmatch_fuzzy_text":                     {},
	"nk_stroke_arc":                              {},
	"nk_stroke_circle":                           {},
	"nk_stroke_curve":                            {},
	"nk_stroke_line":                             {},
	"nk_stroke_polygon":                          {},
	"nk_stroke_polyline":                         {},
	"nk_stroke_rect":                             {},
	"nk_stroke_triangle":                         {},
	"nk_strtod":                                  {},
	"nk_strtof":                                  {},
	"nk_strtoi":                                  {},
	"nk_style_default":                           {},
	"nk_style_from_table":                        {},
	"nk_style_get_color_by_name":                 {},
	"nk_style_hide_cursor":                       {},
	"nk_style_item_color":                        {},
	"nk_style_item_hide":                         {},
	"nk_style_item_image":                        {},
	"nk_style_load_all_cursors":                  {},
	"nk_style_load_cursor":                       {},
	"nk_style_pop_color":                         {},
	"nk_style_pop_flags":                         {},
	"nk_style_pop_float":                         {},
	"nk_style_pop_font":                          {},
	"nk_style_pop_style_item":                    {},
	"nk_style_pop_vec2":                          {},
	"nk_style_push_color":                        {},
	"nk_style_push_flags":                        {},
	"nk_style_push_float":                        {},
	"nk_style_push_font":                         {},
	"nk_style_push_style_item":                   {},
	"nk_style_push_vec2":                         {},
	"nk_style_set_cursor":                        {},
	"nk_style_set_font":                          {},
	"nk_style_show_cursor":                       {},
	"nk_subimage_handle":                         {},
	"nk_subimage_id":                             {},
	"nk_subimage_ptr":                            {},
	"nk_text":                                    {},
	"nk_text_colored":                            {},
	"nk_text_wrap":                               {},
	"nk_text_wrap_colored":                       {},
	"nk_textedit_cut":                            {},
	"nk_textedit_delete":                         {},
	"nk_textedit_delete_selection":               {},
	"nk_textedit_free":                           {},
	"nk_textedit_init":                           {},
	"nk_textedit_init_default":                   {},
	"nk_textedit_init_fixed":                     {},
	"nk_textedit_paste":                          {},
	"nk_textedit_redo":                           {},
	"nk_textedit_select_all":                     {},
	"nk_textedit_text":                           {},
	"nk_textedit_undo":                           {},
	"nk_tooltip":                                 {},
	"nk_tooltip_begin":                           {},
	"nk_tooltip_end":                             {},
	"nk_tree_element_image_push_hashed":          {},
	"nk_tree_element_pop":                        {},
	"nk_tree_element_push_hashed":                {},
	"nk_tree_image_push_hashed":                  {},
	"nk_tree_pop":                                {},
	"nk_tree_push_hashed":                        {},
	"nk_tree_state_image_push":                   {},
	"nk_tree_state_pop":                          {},
	"nk_tree_state_push":                         {},
	"nk_triangle_from_direction":                 {},
	"nk_utf_at":                                  {},
	"nk_utf_decode":                              {},
	"nk_utf_encode":                              {},
	"nk_utf_len":                                 {},
	"nk_vec2":                                    {},
	"nk_vec2i":                                   {},
	"nk_vec2iv":                                  {},
	"nk_vec2v":                                   {},
	"nk_widget":                                  {},
	"nk_widget_bounds":                           {},
	"nk_widget_fitting":                          {},
	"nk_widget_has_mouse_click_down":             {},
	"nk_widget_height":                           {},
	"nk_widget_is_hovered":                       {},
	"nk_widget_is_mouse_clicked":                 {},
	"nk_widget_position":                         {},
	"nk_widget_size":                             {},
	"nk_widget_width":                            {},
	"nk_window_close":                            {},
	"nk_window_collapse":                         {},
	"nk_window_collapse_if":                      {},
	"nk_window_find":                             {},
	"nk_window_get_bounds":                       {},
	"nk_window_get_canvas":                       {},
	"nk_window_get_content_region":               {},
	"nk_window_get_content_region_max":           {},
	"nk_window_get_content_region_min":           {},
	"nk_window_get_content_region_size":          {},
	"nk_window_get_height":                       {},
	"nk_window_get_panel":                        {},
	"nk_window_get_position":                     {},
	"nk_window_get_scroll":                       {},
	"nk_window_get_size":                         {},
	"nk_window_get_width":                        {},
	"nk_window_has_focus":                        {},
	"nk_window_is_active":                        {},
	"nk_window_is_any_hovered":                   {},
	"nk_window_is_closed":                        {},
	"nk_window_is_collapsed":                     {},
	"nk_window_is_hidden":                        {},
	"nk_window_is_hovered":                       {},
	"nk_window_set_bounds":                       {},
	"nk_window_set_focus":                        {},
	"nk_window_set_position":                     {},
	"nk_window_set_scroll":                       {},
	"nk_window_set_size":                         {},
	"nk_window_show":                             {},
	"nk_window_show_if":                          {},
	"stbrp_init_target":                          {},
	"stbrp_pack_rects":                           {},
	"stbrp_setup_allow_out_of_mem":               {},
	"stbrp_setup_heuristic":                      {},
	"stbtt_BakeFontBitmap":                       {},
	"stbtt_CompareUTF8toUTF16_bigendian":         {},
	"stbtt_FindGlyphIndex":                       {},
	"stbtt_FindMatchingFont":                     {},
	"stbtt_FindSVGDoc":                           {},
	"stbtt_FreeBitmap":                           {},
	"stbtt_FreeSDF":                              {},
	"stbtt_FreeShape":                            {},
	"stbtt_GetBakedQuad":                         {},
	"stbtt_GetCodepointBitmap":                   {},
	"stbtt_GetCodepointBitmapBox":                {},
	"stbtt_GetCodepointBitmapBoxSubpixel":        {},
	"stbtt_GetCodepointBitmapSubpixel":           {},
	"stbtt_GetCodepointBox":                      {},
	"stbtt_GetCodepointHMetrics":                 {},
	"stbtt_GetCodepointKernAdvance":              {},
	"stbtt_GetCodepointSDF":                      {},
	"stbtt_GetCodepointSVG":                      {},
	"stbtt_GetCodepointShape":                    {},
	"stbtt_GetFontBoundingBox":                   {},
	"stbtt_GetFontNameString":                    {},
	"stbtt_GetFontOffsetForIndex":                {},
	"stbtt_GetFontVMetrics":                      {},
	"stbtt_GetFontVMetricsOS2":                   {},
	"stbtt_GetGlyphBitmap":                       {},
	"stbtt_GetGlyphBitmapBox":                    {},
	"stbtt_GetGlyphBitmapBoxSubpixel":            {},
	"stbtt_GetGlyphBitmapSubpixel":               {},
	"stbtt_GetGlyphBox":                          {},
	"stbtt_GetGlyphHMetrics":                     {},
	"stbtt_GetGlyphKernAdvance":                  {},
	"stbtt_GetGlyphSDF":                          {},
	"stbtt_GetGlyphSVG":                          {},
	"stbtt_GetGlyphShape":                        {},
	"stbtt_GetKerningTable":                      {},
	"stbtt_GetKerningTableLength":                {},
	"stbtt_GetNumberOfFonts":                     {},
	"stbtt_GetPackedQuad":                        {},
	"stbtt_GetScaledFontVMetrics":                {},
	"stbtt_InitFont":                             {},
	"stbtt_IsGlyphEmpty":                         {},
	"stbtt_MakeCodepointBitmap":                  {},
	"stbtt_MakeCodepointBitmapSubpixel":          {},
	"stbtt_MakeCodepointBitmapSubpixelPrefilter": {},
	"stbtt_MakeGlyphBitmap":                      {},
	"stbtt_MakeGlyphBitmapSubpixel":              {},
	"stbtt_MakeGlyphBitmapSubpixelPrefilter":     {},
	"stbtt_PackBegin":                            {},
	"stbtt_PackEnd":                              {},
	"stbtt_PackFontRange":                        {},
	"stbtt_PackFontRanges":                       {},
	"stbtt_PackFontRangesGatherRects":            {},
	"stbtt_PackFontRangesPackRects":              {},
	"stbtt_PackFontRangesRenderIntoRects":        {},
	"stbtt_PackSetOversampling":                  {},
	"stbtt_PackSetSkipMissingCodepoints":         {},
	"stbtt_Rasterize":                            {},
	"stbtt_ScaleForMappingEmToPixels":            {},
	"stbtt_ScaleForPixelHeight":                  {},
}
